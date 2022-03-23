// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package generic

import (
	"bytes"
	"fmt"
	pgs "github.com/lyft/protoc-gen-star"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

const moduleName = "choice"

var templateDir = os.Getenv("GOPATH")
var templates = template.Must(template.ParseGlob(filepath.Join(templateDir, "src/github.com/onosproject/onos-e2-sm/protoc-gen-choice/templates/choice.tpl")))

type encoderStruct struct {
	Logger                  bool   // this is to indicate of logger was initialized - one structure should be randomly picked and assigned a logger flag
	Imports                 string // this is to hold all necessary imports, mainly the Protobuf which contains top-level PDUs (they're defined as PDUs with multiple formats wrapped as a nested CHOICE inside SEQUENCE)
	ProtoName               string // this is to hold ProtobufName to refer to the same package
	MessageName             string // top-level PDU message name
	MessageNameInLogging    string // this is to hold message name in the logs
	ChoiceMapName           string // this is to hold choice map name
	CanonicalChoiceMapName  string // this is to hold canonical choice map name
	CanonicalChoicePresence bool   // indicates whether Canonical ordering for choices is present and the map should be passed
	Parameters              string // this is to hold all necessary parameters to pass to the message - mostly it is "choiceExt" or "valueExt"
}

// Defines data structure to pass to enum template
type choiceStruct struct {
	PackageName             string
	Imports                 string
	MapName                 string
	Choices                 []choiceMsg
	CanonicalChoices        []canonicalChoiceMsg
	CanonicalChoicePresence bool
}

type choiceMsg struct {
	MsgName string
	Items   []choiceItem
}

type choiceItem struct {
	Leafs      []leaf
	ChoiceName string
}

type leaf struct {
	Index         int
	LeafName      string
	ProtoFileName string
}

type canonicalChoiceMsg struct {
	MsgName string
	Items   []canonicalChoiceItem
}

type canonicalChoiceItem struct {
	Leafs      []canonicalLeaf
	ChoiceName string
}

type canonicalLeaf struct {
	Index         string
	LeafName      string
	ProtoFileName string
	PackageName   string
	ItemType      string // This is to store type of the CHOICE item (for correct referencing in index)
}

//type canonicalChoicesList struct {
//	Name string
//}

// ReportModule creates a report of all the target messages generated by the
// protoc run, writing the file into the /tmp directory.
type reportModule struct {
	*pgs.ModuleBase
}

// NewModule configures the module with an instance of ModuleBase
func NewModule() pgs.Module {
	return &reportModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

// Name is the identifier used to identify the module. This value is
// automatically attached to the BuildContext associated with the ModuleBase.
func (m *reportModule) Name() string {
	return moduleName
}

// Execute is passed the target files as well as its dependencies in the pkgs
// map. The implementation should return a slice of Artifacts that represent
// the files to be generated. In this case, "/tmp/report.txt" will be created
// outside of the normal protoc flow.
func (m *reportModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}

	dir, err := os.Getwd()
	// handle err
	if err != nil {
		return nil
	}
	//printFiles(path)
	_, err = fmt.Fprintf(buf, "Working directory is %v\n", dir)
	if err != nil {
		return nil
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	for _, file := range files {
		_, err = fmt.Fprintf(buf, "Found file %v\n", file.Name())
		if err != nil {
			return nil
		}
	}

	choices := choiceStruct{
		Choices:          make([]choiceMsg, 0),
		CanonicalChoices: make([]canonicalChoiceMsg, 0),
	}

	canonicalChoices := make([]string, 0)

	for _, f := range targets { // Input .proto files
		for _, msg := range f.AllMessages() {
			_, err = fmt.Fprintf(buf, "Message name is %v\n", msg.Name().String())
			if err != nil {
				return nil
			}
			for _, item := range msg.Fields() {
				_, err = fmt.Fprintf(buf, "Item's name is %v\n", getItemType(item.Descriptor().GetTypeName()))
				if err != nil {
					return nil
				}
				_, err = fmt.Fprintf(buf, "Comments for this message are \n%v\n", item.SourceCodeInfo().LeadingComments())
				if err != nil {
					return nil
				}
				if strings.Contains(item.SourceCodeInfo().LeadingComments(), "canonicalOrder") {
					_, err = fmt.Fprintf(buf, "----------------- Adding %v to the Canonical CHOICE map\n", getItemType(item.Descriptor().GetTypeName()))
					if err != nil {
						return nil
					}
					canonicalChoices = append(canonicalChoices, getItemType(item.Descriptor().GetTypeName()))
				}
			}
		}
	}

	_, err = fmt.Fprintf(buf, "Messages with CanonicalChoice ordering are:\n")
	if err != nil {
		return nil
	}
	for l, iitem := range canonicalChoices {
		_, err = fmt.Fprintf(buf, "%v - message name is - %v\n", l, iitem)
		if err != nil {
			return nil
		}
	}

	if len(canonicalChoices) > 0 {
		choices.CanonicalChoicePresence = true
	}

	_, err = fmt.Fprintf(buf, "Canonical CHOICE presence is %v\n", choices.CanonicalChoicePresence)
	if err != nil {
		return nil
	}

	// In case of canonical choice presence we should read out constants in order to compose map correctly
	constantsList := make([]string, 0)
	if choices.CanonicalChoicePresence {
		for _, f := range targets {
			// fix on .proto file with constants
			if strings.Contains(f.Name().Split()[0], "constants") {
				for _, msg := range f.AllMessages() {
					// constant has only single item
					structType := msg.Fields()[0].SourceCodeInfo().TrailingComments()
					if structType != "" {
						name := msg.Name().String()
						_, err = fmt.Fprintf(buf, "Name is %v\n", name)
						if err != nil {
							return nil
						}
						_, err = fmt.Fprintf(buf, "StructType is %v\n", structType)
						if err != nil {
							return nil
						}

						constant := adjustCanonicalChoiceIndex(name, structType)
						_, err = fmt.Fprintf(buf, "Constant is %v\n", constant)
						if err != nil {
							return nil
						}
						constantsList = append(constantsList, constant)
					}
				}
				break
			}
		}
	}

	// Figuring out name of the package
	packageName := ""
	// We are assuming that all .proto files are located at the same directory (the last directory usually contains version number)
	for _, f := range targets {
		_, err = fmt.Fprintf(buf, "Obtained input path is \n%v\n", f.InputPath().Dir().String())
		if err != nil {
			return nil
		}

		packageName = extractDirectoryName(f.InputPath().Dir().String())
		// Input .proto files
		if packageName != "" {
			choices.PackageName = dashToUnderscore(packageName)
		} else {
			choices.PackageName = adjustPackageName(adjustProtoFileName(extractProtoFileName(f.Name().Split()[0])), f.File().InputPath().Dir().String())
		}
		break
	}

	for _, f := range targets { // Input .proto files
		m.Push(f.Name().String()).Debug("reporting")
		targetPath := ""
		_, err := fmt.Fprintf(buf, "Processing following Proto file %v\n", f.File().InputPath().BaseName())
		if err != nil {
			return nil
		}
		targetPath = f.File().InputPath().Dir().String()
		_, err = fmt.Fprintf(buf, "Target file Proto path pattern is %v\n", targetPath)
		if err != nil {
			return nil
		}

		protoFilePath := ""
		err = filepath.Walk(dir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if strings.Contains(path, f.File().InputPath().BaseName()+".pb.go") && !strings.Contains(path, "/_") && strings.Contains(path, targetPath) {
					protoFilePath = path
					_, err = fmt.Fprintf(buf, "Hooray! Found file %v with path %v\n", info.Name(), path)
					if err != nil {
						return nil
					}
				}

				return nil
			})
		if err != nil {
			_, err = fmt.Fprintf(buf, "Something went wrong in searching for the file path.. %v\n", err)
			if err != nil {
				return nil
			}
		}

		if protoFilePath == "" {
			_, err = fmt.Fprintf(buf, "ProtoFilePath is empty %v\n", protoFilePath)
			if err != nil {
				return nil
			}
			m.OverwriteCustomFile(
				"/tmp/report.txt",
				buf.String(),
				0644,
			)

			return m.Artifacts()
		} else {
			//Composing protoFilePath to correspond to the correct input for Golang imports
			index := strings.Index(protoFilePath, "github.com/")
			if index == -1 {
				_, err = fmt.Fprintf(buf, "Something went wrong in searching for the file path for the import..\n%v", protoFilePath)
				if err != nil {
					return nil
				}
			}
			protoFilePath = protoFilePath[index:]

			//Taking out file name from the path
			indexx := strings.LastIndex(protoFilePath, "/")
			if indexx == -1 {
				_, err = fmt.Fprintf(buf, "Something went wrong in searching for the file path for the import..\n%v", protoFilePath)
				if err != nil {
					return nil
				}
			}
			protoFilePath = protoFilePath[:indexx]

			_, err = fmt.Fprintf(buf, "Proto file path is %v\n", protoFilePath)
			if err != nil {
				return nil
			}
		}

		choices.Imports = choices.Imports + adjustPackageName(adjustProtoFileName(extractProtoFileName(f.Name().Split()[0])), f.File().InputPath().Dir().String()) + " \"" + protoFilePath + "\"" + "\n"
		if choices.MapName == "" {
			choices.MapName = adjustMapVariableName(extractPackageName(f.Name().Split()[0]))
		}

		_, err = fmt.Fprintf(buf, "MapVariableName is %v\n", choices.MapName)
		if err != nil {
			return nil
		}

		for _, msg := range f.AllMessages() {
			_, err = fmt.Fprintf(buf, "Message name is %v\n", msg.Name().String())
			if err != nil {
				return nil
			}
			for _, item := range msg.Fields() {
				_, err = fmt.Fprintf(buf, "Item's name is %v\n", getItemType(item.Descriptor().GetTypeName()))
				if err != nil {
					return nil
				}
				_, err = fmt.Fprintf(buf, "Comments for this message are \n%v\n", item.SourceCodeInfo().LeadingComments())
				if err != nil {
					return nil
				}
				if strings.Contains(item.SourceCodeInfo().LeadingComments(), "canonicalOrder") {
					_, err = fmt.Fprintf(buf, "----------------- Adding %v to the Canonical CHOICE map\n", getItemType(item.Descriptor().GetTypeName()))
					if err != nil {
						return nil
					}
					canonicalChoices = append(canonicalChoices, getItemType(item.Descriptor().GetTypeName()))
				}
			}
		}

		for _, msg := range f.AllMessages() {
			_, err = fmt.Fprintf(buf, "Message name is %v\n", msg.Name().String())
			if err != nil {
				return nil
			}
			if msg.OneOfs() != nil {
				if choices.CanonicalChoicePresence && presentInCanonicalChoiceList(msg.Name().String(), canonicalChoices) {
					//ToDo - compose canonical choice struct
					_, err = fmt.Fprintf(buf, "------------------- Filling in canonical choice structure\n")
					if err != nil {
						return nil
					}

					_, err = fmt.Fprintf(buf, "---- Message name is %v\n", msg.Name().String())
					if err != nil {
						return nil
					}

					chMsg := canonicalChoiceMsg{
						MsgName: adjustOneOfStructName(msg.Name().String()),
						Items:   make([]canonicalChoiceItem, 0),
					}

					for i, plg := range msg.OneOfs() {
						if !strings.HasPrefix(plg.Name().String(), "_") {
							_, err = fmt.Fprintf(buf, "%v OneOf name is %v\n", i+1, plg.Name())
							if err != nil {
								return nil
							}
							chItem := canonicalChoiceItem{
								ChoiceName: plg.Name().String(),
								Leafs:      make([]canonicalLeaf, 0),
							}
							for j, field := range plg.Fields() {
								_, err = fmt.Fprintf(buf, "OneOf field (item) type (child structure) is %v\n", getItemType(field.Descriptor().GetTypeName()))
								if err != nil {
									return nil
								}

								_, err = fmt.Fprintf(buf, "%v, OneOf field is %v\n", j+1, field.Name().String())
								if err != nil {
									return nil
								}
								lf := canonicalLeaf{
									PackageName: extractPackageNameForCanonicalChoices(f.File().InputPath().Dir().String()),
									Index:       lookUpCanonicalChoiceIndex(strings.ReplaceAll(field.Name().String(), "_", ""), constantsList),
									LeafName:    msg.Name().String() + "_" + adjustOneOfLeafName(field.Name().String()),
								}
								lf.ProtoFileName = adjustPackageName(adjustProtoFileName(extractProtoFileName(f.Name().Split()[0])), f.File().InputPath().Dir().String())
								chItem.Leafs = append(chItem.Leafs, lf)
							}
							chMsg.Items = append(chMsg.Items, chItem)
						}
					}

					_, err = fmt.Fprintf(buf, "Obtained OneOf item is \n%v\n", chMsg)
					if err != nil {
						return nil
					}

					choices.CanonicalChoices = append(choices.CanonicalChoices, chMsg)
				} else {
					chMsg := choiceMsg{
						MsgName: adjustOneOfStructName(msg.Name().String()),
						Items:   make([]choiceItem, 0),
					}

					for i, plg := range msg.OneOfs() {
						if !strings.HasPrefix(plg.Name().String(), "_") {
							_, err = fmt.Fprintf(buf, "%v OneOf name is %v\n", i+1, plg.Name())
							if err != nil {
								return nil
							}
							chItem := choiceItem{
								ChoiceName: plg.Name().String(),
								Leafs:      make([]leaf, 0),
							}
							for j, field := range plg.Fields() {
								_, err = fmt.Fprintf(buf, "OneOf field (item) type (child structure) is %v\n", getItemType(field.Descriptor().GetTypeName()))
								if err != nil {
									return nil
								}

								_, err = fmt.Fprintf(buf, "%v, OneOf field is %v\n", j+1, field.Name())
								if err != nil {
									return nil
								}
								lf := leaf{
									Index:    j + 1,
									LeafName: msg.Name().String() + "_" + adjustOneOfLeafName(field.Name().String()),
								}
								lf.ProtoFileName = adjustPackageName(adjustProtoFileName(extractProtoFileName(f.Name().Split()[0])), f.File().InputPath().Dir().String())
								chItem.Leafs = append(chItem.Leafs, lf)
							}
							chMsg.Items = append(chMsg.Items, chItem)
						}
					}

					_, err = fmt.Fprintf(buf, "Obtained OneOf item is \n%v\n", chMsg)
					if err != nil {
						return nil
					}

					choices.Choices = append(choices.Choices, chMsg)
				}
			}
		}

		m.Pop()
		_, err = fmt.Fprintf(buf, "-----------------------------------------------------------------------------------------------\n")
		if err != nil {
			return nil
		}
	}

	_, err = fmt.Fprintf(buf, "Messages with CanonicalChoice ordering are:\n")
	if err != nil {
		return nil
	}
	for l, iitem := range canonicalChoices {
		_, err = fmt.Fprintf(buf, "%v - message name is - %v\n", l, iitem)
		if err != nil {
			return nil
		}
	}

	if len(canonicalChoices) > 0 {
		choices.CanonicalChoicePresence = true
	}

	_, err = fmt.Fprintf(buf, "Canonical CHOICE presence is %v\n", choices.CanonicalChoicePresence)
	if err != nil {
		return nil
	}

	//Generating new .go file
	m.OverwriteGeneratorTemplateFile("choiceOptions.go", templates.Lookup("choice.tpl"), choices)

	sm, _ := m.Parameters().Bool("sm")

	if sm {
		//ToDo - generate servicemodel package

	}

	m.OverwriteCustomFile(
		"/tmp/report.txt",
		buf.String(),
		0644,
	)

	return m.Artifacts()
}

/////////////////////////////////
/// Here is necessary tooling ///
/////////////////////////////////

func upperCaseFirstLetter(str string) string {

	for i, ch := range str {
		return string(unicode.ToUpper(ch)) + str[i+1:]
	}
	return ""
}

func adjustOneOfLeafName(leafName string) string {

	var res string
	for _, element := range strings.Split(leafName, "_") {
		res = res + upperCaseFirstLetter(element)
	}

	return res
}

func adjustOneOfStructName(msgName string) string {

	var res string
	for i, r := range msgName {
		if unicode.IsUpper(r) {
			if i > 1 {
				res = res + "_" + strings.ToLower(string(r))
			} else {
				res = strings.ToLower(string(r))
			}
		} else {
			res = res + string(r)
		}
	}

	return res
}

func extractProtoFileName(proto string) string {

	if strings.LastIndex(proto, "/") != -1 {
		return proto[strings.LastIndex(proto, "/")+1:]
	}
	return proto
}

func extractPackageName(pckg string) string {

	if strings.Contains(pckg, "/") {
		return pckg[:strings.Index(pckg, "/")]
	}
	return pckg
}

func extractPackageNameForCanonicalChoices(path string) string {

	re := regexp.MustCompile(`v\d{1}`)
	res := re.FindString(path)
	return res
}

func adjustProtoFileName(filename string) string {

	res := dashToUnderscore(filename)
	// space for future adjustments
	return res
}

func adjustPackageName(filename string, path string) string {

	// remove redundant patterns
	res := strings.ReplaceAll(filename, "_go", "")

	//extract version number
	index := strings.LastIndex(path, "/v")
	version := path[index+1:]

	// clarify if the version is not already present in the filename
	present := strings.Contains(res, version)

	//check if another versioning number presents in the naming - relevant for external dependencies
	re := regexp.MustCompile(`v\d{1}`)
	externalVersion := re.MatchString(res)

	if version != "" && !present && !externalVersion {
		res = res + version
	}

	//remove all underscores
	res = strings.ReplaceAll(res, "_", "")
	//remove all camel cases
	res = strings.ToLower(res)

	return res
}

func adjustMapVariableName(mapName string) string {

	flag := true
	for flag {
		index := strings.IndexAny(mapName, "_")
		if index == -1 {
			flag = false
		} else {
			// Assumption - string doesn't end with "_"
			mapName = mapName[:index] + strings.ToUpper(mapName[index+1:index+2]) + mapName[index+2:]
		}
	}
	mapName = strings.ReplaceAll(mapName, "Go", "")
	//ToDo - make a workaround through regexp
	//re := regexp.MustCompile(`v\d{1}`)
	//mapName = re.ReplaceAllString(mapName, "")
	mapName = strings.ReplaceAll(mapName, "V1", "")
	mapName = strings.ReplaceAll(mapName, "v1", "")
	mapName = strings.ReplaceAll(mapName, "V2", "")
	mapName = strings.ReplaceAll(mapName, "v2", "")
	mapName = strings.ReplaceAll(mapName, "V3", "")
	mapName = strings.ReplaceAll(mapName, "v3", "")
	mapName = strings.ReplaceAll(mapName, "V4", "")
	mapName = strings.ReplaceAll(mapName, "v4", "")

	return strings.ToUpper(mapName[:1]) + mapName[1:]
}

func dashToUnderscore(str string) string {

	return strings.ReplaceAll(str, "-", "_")
}

func extractDirectoryName(str string) string {

	res := ""
	index := strings.LastIndex(str, "/")
	if index != -1 {
		res = str[index+1:]
	} else {
		res = str
	}

	return res
}

func getItemType(str string) string {

	res := ""
	index := strings.LastIndex(str, ".")
	if index != -1 {
		res = str[index+1:]
	} else {
		res = "ErrorInParsing"
	}

	return res
}

func presentInCanonicalChoiceList(target string, list []string) bool {

	flag := false
	for _, item := range list {
		if strings.EqualFold(item, target) {
			flag = true
		}
	}

	return flag
}

func adjustCanonicalChoiceIndex(name string, structType string) string {

	res := ""
	// Assuming that all names start from Id which should be cut out
	newName := name[2:]
	//newName := strings.ReplaceAll(name, "Id", "")
	newStructType := strings.ReplaceAll(structType, "-", "") // could contain "ID"

	tmp := newStructType + newName
	if !strings.Contains(strings.ToLower(newStructType), "id") {
		tmp = newStructType + "ID" + newName
	}

	res = strings.ReplaceAll(tmp, "\n", "")
	res = strings.ReplaceAll(res, "IE", "Ie")
	res = strings.ReplaceAll(res, "Id", "ID")
	res = strings.ReplaceAll(res, " ", "")
	return res
}

func lookUpCanonicalChoiceIndex(name string, list []string) string {

	res := ""
	for _, item := range list {
		if strings.Contains(strings.ToLower(item), strings.ToLower(name)) {
			res = item
			break
		}
	}

	if res == "" {
		res = "NoIndexDefined"
	}
	return res
}
