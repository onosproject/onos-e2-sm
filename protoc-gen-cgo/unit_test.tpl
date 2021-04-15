// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package {{.PackageName}}ctypes

import (
"encoding/hex"
"fmt"
pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/{{.ProtoFileName}}/pdubuilder"
{{.ProtoFileName}} "github.com/onosproject/onos-e2-sm/servicemodels/{{.ProtoFileName}}{{.FullPackageName}}" //ToDo - Make imports more dynamic
"gotest.tools/assert"
"testing"
)

func create{{.GlobalName}}Msg() (*{{.ProtoFileName}}.{{.MessageName}}, error) {

// {{lowCaseFirstLetter .GlobalName}} := pdubuilder.Create{{.GlobalName}}() //ToDo - fill in arguments here(if this function exists

{{lowCaseFirstLetter .GlobalName}} := {{.ProtoFileName}}.{{.MessageName}}{ {{if .Optional}}{{ $optionalFields1 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields1 }}
{{upperCaseFirstLetter .FieldName}}: nil,{{end}}
{{ $repeatedFields1 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields1 }}{{upperCaseFirstLetter .FieldName}}: make([]*{{.ProtoFileName}}.{{.DataType}}, 0), //ToDo - Check if protobuf structure is implemented correctly (mainly naming)
{{end}}
{{end}}{{if .OneOf}}{{ $oneofFields1 := .FieldList.OneOfField }}
{{ range $fieldIndex, $field := $oneofFields1 }}{{upperCaseFirstLetter .FieldName}}: nil,
{{.MessageName}}: &{{.ProtoFileName}}.{{.MessageName}}_{{.FieldName}} {
{{upperCaseFirstLetter .FieldName}}: &{{.ProtoFileName}}.{{lowCaseFirstLetter .FieldName}} {
    Value: nil,
},
},
{{end}}
{{end}}
}

if err := {{lowCaseFirstLetter .GlobalName}}.Validate(); err != nil {
return nil, fmt.Errorf("error validating {{.MessageName}} %s", err.Error())
}
return &{{lowCaseFirstLetter .GlobalName}}, nil
}

func Test_xerEncoding{{.GlobalName}}(t *testing.T) {

{{lowCaseFirstLetter .GlobalName}}, err := create{{.GlobalName}}Msg()
assert.NilError(t, err, "Error creating {{.MessageName}} PDU")

xer, err := xerEncode{{.GlobalName}}({{lowCaseFirstLetter .GlobalName}})
assert.NilError(t, err)
assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
t.Logf("{{.GlobalName}} XER\n%s", string(xer))

result, err := xerDecode{{.GlobalName}}(xer)
assert.NilError(t, err)
assert.Assert(t, result != nil)
t.Logf("{{.GlobalName}} XER - decoded\n%v", result)
//ToDo - adjust field's verification {{ $optionalFields1 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields1 }}
assert.Equal(t, {{lowCaseFirstLetter .GlobalName}}.Get{{upperCaseFirstLetter .FieldName}}(), result.Get{{upperCaseFirstLetter .FieldName}}()){{end}}
{{ $repeatedFields1 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields1 }}
assert.Equal(t, 1, len(result.Get{{upperCaseFirstLetter .FieldName}}())) //ToDo - adjust length of a list
assert.DeepEqual(t, {{lowCaseFirstLetter .GlobalName}}.Get{{upperCaseFirstLetter .FieldName}}(), result.Get{{upperCaseFirstLetter .FieldName}}()){{end}}
{{if .OneOf}}//This one if for OneOf fields
{{ $oneofFields1 := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneofFields1 }}
assert.DeepEqual(t, {{lowCaseFirstLetter .GlobalName}}.Get{{upperCaseFirstLetter .FieldName}}().GetValue(), result.Get{{upperCaseFirstLetter .FieldName}}().GetValue()){{end}}
{{end}}
}

func Test_perEncoding{{.GlobalName}}(t *testing.T) {

{{lowCaseFirstLetter .GlobalName}}, err := create{{.GlobalName}}Msg()
assert.NilError(t, err, "Error creating {{.MessageName}} PDU")

per, err := perEncode{{.GlobalName}}({{lowCaseFirstLetter .GlobalName}})
assert.NilError(t, err)
assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
t.Logf("{{.GlobalName}} PER\n%v", hex.Dump(per))

result, err := perDecode{{.GlobalName}}(per)
assert.NilError(t, err)
assert.Assert(t, result != nil)
t.Logf("{{.GlobalName}} PER - decoded\n%v", result)
//ToDo - adjust field's verification {{ $optionalFields2 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields2 }}
assert.Equal(t, {{lowCaseFirstLetter .GlobalName}}.Get{{upperCaseFirstLetter .FieldName}}(), result.Get{{upperCaseFirstLetter .FieldName}}()){{end}}
{{ $repeatedFields2 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields2 }}
assert.Equal(t, 1, len(result.Get{{upperCaseFirstLetter .FieldName}}())) //ToDo - adjust length of a list
assert.DeepEqual(t, {{lowCaseFirstLetter .GlobalName}}.Get{{upperCaseFirstLetter .FieldName}}(), result.Get{{upperCaseFirstLetter .FieldName}}()){{end}}
{{if .OneOf}}//This one if for OneOf fields
{{ $oneofFields1 := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneofFields1 }}
assert.DeepEqual(t, {{lowCaseFirstLetter .GlobalName}}.Get{{upperCaseFirstLetter .FieldName}}().GetValue(), result.Get{{upperCaseFirstLetter .FieldName}}().GetValue()){{end}}
{{end}}
}
