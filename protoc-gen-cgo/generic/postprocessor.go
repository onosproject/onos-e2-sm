// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package generic

import (
	"fmt"
	"github.com/lyft/protoc-gen-star"
)

// New returns a PostProcessor that adds a copyright comment to the top
// of all generated files.
func NewPostProcessor() pgs.PostProcessor {
	return copyrightPostProcessor{}
}

type copyrightPostProcessor struct{}

// Match returns true only for Custom and Generated files (including templates).
func (cpp copyrightPostProcessor) Match(a pgs.Artifact) bool {
	fmt.Printf("This is POST-Processor!\n")
	switch f := a.(type) {
	case pgs.GeneratorFile, pgs.GeneratorTemplateFile,
		pgs.CustomFile, pgs.CustomTemplateFile:
		// This is what later call Process() method
		fmt.Printf("Match found!")
		fmt.Printf("---------------------------------------->\n%v", f)
		return true
	default:
		fmt.Printf("FALSE ONE ---------------------------------------->\n%v", f)
		return false
	}
}

// Process attaches the copyright header to the top of the input bytes
func (cpp copyrightPostProcessor) Process(in []byte) (out []byte, err error) {
	cmt := fmt.Sprintf("// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>\n//\n" +
		"// SPDX-License-Identifier: Apache-2.0\n")

	return append([]byte(cmt), in...), nil
}
