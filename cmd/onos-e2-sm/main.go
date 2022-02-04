// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/rogpeppe/go-internal/modfile"
	"github.com/rogpeppe/go-internal/module"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	modFile       = "go.mod"
	modVersionSep = "@"
)

func main() {
	if err := getCmd().Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "onos-e2-sm",
	}
	cmd.AddCommand(getGenDepsCmd())
	return cmd
}

func getGenDepsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "gen-deps",
		Short:        "Generates a go.mod file for the given service model module",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			model := args[0]
			target, _ := cmd.Flags().GetString("target")

			resolver := newModuleResolver()
			mod, err := resolver.resolve(model, target)
			if err != nil {
				return err
			}
			bytes, err := mod.Format()
			if err != nil {
				return err
			}
			os.Stdout.Write(bytes)
			return nil
		},
	}
	cmd.Flags().StringP("target", "t", "", "the target Go module")
	return cmd
}

// newModuleResolver creates a new module resolver
func newModuleResolver() *moduleResolver {
	return &moduleResolver{}
}

// moduleResolver is a module resolver
type moduleResolver struct{}

func (r *moduleResolver) exec(dir string, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GO111MODULE=on", "CGO_ENABLED=1")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (r *moduleResolver) getGoEnv() (goEnv, error) {
	wd, err := os.Getwd()
	if err != nil {
		return goEnv{}, err
	}

	envJSON, err := r.exec(wd, "go", "env", "-json", "GOPATH", "GOMODCACHE")
	if err != nil {
		return goEnv{}, err
	}
	env := goEnv{}
	if err := json.Unmarshal([]byte(envJSON), &env); err != nil {
		return goEnv{}, err
	}
	return env, nil
}

func (r *moduleResolver) getGoModCacheDir() (string, error) {
	env, err := r.getGoEnv()
	if err != nil {
		return "", err
	}
	modCache := env.GOMODCACHE
	if modCache == "" {
		// For Go 1.14 and older.
		return filepath.Join(env.GOPATH, "pkg", "mod"), nil
	}
	return modCache, nil
}

func (r *moduleResolver) getModName(model string) string {
	return fmt.Sprintf("github.com/onosproject/onos-e2-sm/servicemodels/%s", model)
}

// resolve resolves the module info for the target module
func (r *moduleResolver) resolve(model, target string) (*modfile.File, error) {
	mod, err := r.fetchMod(target)
	if err != nil {
		return nil, err
	}
	if err := mod.AddModuleStmt(r.getModName(model)); err != nil {
		return nil, err
	}
	return mod, nil
}

func (r *moduleResolver) fetchMod(target string) (*modfile.File, error) {
	localModPath := filepath.Join(target, "go.mod")
	if _, err := os.Stat(localModPath); !os.IsNotExist(err) {
		localModBytes, err := ioutil.ReadFile(localModPath)
		if err != nil {
			return nil, err
		}
		return modfile.Parse(localModPath, localModBytes, nil)
	}

	var replace string
	if !strings.HasPrefix(target, "github.com/onosproject/") {
		replace = target
		elems := strings.SplitN(target, "/", 3)
		target = fmt.Sprintf("github.com/onosproject/%s", elems[2])
	}

	if target == "" {
		err := errors.NewInvalid("no target module configured")
		return nil, err
	}

	targetPath, _ := splitModPathVersion(target)

	fakeModDir, err := ioutil.TempDir("", "onos-e2-sm-plugin")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(fakeModDir)

	// Generate a temporary module with which to pull the target module
	fakeMod := []byte("module m\n")
	if replace != "" {
		replacePath, replaceVersion := splitModPathVersion(replace)
		fakeMod = append(fakeMod, []byte(fmt.Sprintf("replace %s => %s %s\n", targetPath, replacePath, replaceVersion))...)
	}

	// Write the temporary module file
	fakeModPath := filepath.Join(fakeModDir, modFile)
	if err := ioutil.WriteFile(fakeModPath, fakeMod, 0666); err != nil {
		return nil, err
	}

	// Add the target dependency to the temporary module and download the target module
	if _, err := r.exec(fakeModDir, "go", "get", "-d", target); err != nil {
		return nil, err
	}

	// Read the updated go.mod for the temporary module
	fakeMod, err = ioutil.ReadFile(fakeModPath)
	if err != nil {
		return nil, err
	}

	// Parse the updated go.mod for the temporary module
	tmpModFile, err := modfile.Parse(fakeModPath, fakeMod, nil)
	if err != nil {
		return nil, err
	}

	// Determine the path/version for the target module
	var modPath string
	var modVersion string
	if replace == "" {
		for _, require := range tmpModFile.Require {
			if require.Mod.Path == targetPath {
				modPath = require.Mod.Path
				modVersion = require.Mod.Version
				break
			}
		}
	} else {
		for _, replace := range tmpModFile.Replace {
			if replace.Old.Path == targetPath {
				modPath = replace.New.Path
				modVersion = replace.New.Version
				break
			}
		}
	}

	// Encode the target dependency path
	encPath, err := module.EncodePath(modPath)
	if err != nil {
		return nil, err
	}
	modPath = encPath

	// Lookup the Go cache from the environment
	modCache, err := r.getGoModCacheDir()
	if err != nil {
		return nil, err
	}

	// Read the target go.mod from the cache
	cacheModPath := filepath.Join(modCache, "cache", "download", modPath, "@v", modVersion+".mod")
	modBytes, err := ioutil.ReadFile(cacheModPath)
	if err != nil {
		return nil, err
	}

	// Parse the target go.mod
	targetModFile, err := modfile.Parse(cacheModPath, modBytes, nil)
	if err != nil {
		return nil, err
	}
	return targetModFile, nil
}

func splitModPathVersion(mod string) (string, string) {
	if i := strings.Index(mod, modVersionSep); i >= 0 {
		return mod[:i], mod[i+1:]
	}
	return mod, ""
}

type goEnv struct {
	GOPATH     string
	GOMODCACHE string
}
