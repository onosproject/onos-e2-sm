// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package {{.PackageName}}ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "{{underscoreToDash .CstructName}}.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
{{if .Repeated}}//#include ".h" //ToDo - include correct .h file for corresponding C-struct of "Repeated" field or other anonymous structure defined in .h file
import "C"{{else}}import "C"
{{end}}
import (
    "fmt"
    {{.ProtoFileName}} "github.com/onosproject/onos-e2-sm/servicemodels/{{.FullPackageName}}" //ToDo - Make imports more dynamic
    "unsafe"
)

func xerEncode{{.MessageName}}({{lowCaseFirstLetter .MessageName}} *{{.ProtoFileName}}.{{.MessageName}}) ([]byte, error) {
    {{lowCaseFirstLetter .MessageName}}CP, err := new{{.MessageName}}({{lowCaseFirstLetter .MessageName}})
    if err != nil {
        return nil, fmt.Errorf("xerEncode{{.MessageName}}() %s", err.Error())
    }

    bytes, err := encodeXer(&C.asn_DEF_{{dashToUnderscore .CstructName}}, unsafe.Pointer({{lowCaseFirstLetter .MessageName}}CP))
    if err != nil {
        return nil, fmt.Errorf("xerEncode{{.MessageName}}() %s", err.Error())
    }
    return bytes, nil
}

func perEncode{{.MessageName}}({{lowCaseFirstLetter .MessageName}} *{{.ProtoFileName}}.{{.MessageName}}) ([]byte, error) {
    {{lowCaseFirstLetter .MessageName}}CP, err := new{{.MessageName}}({{lowCaseFirstLetter .MessageName}})
    if err != nil {
        return nil, fmt.Errorf("xerEncode{{.MessageName}}() %s", err.Error())
    }

    bytes, err := encodePerBuffer(&C.asn_DEF_{{dashToUnderscore .CstructName}}, unsafe.Pointer({{lowCaseFirstLetter .MessageName}}CP))
    if err != nil {
        return nil, fmt.Errorf("perEncode{{.MessageName}}() %s", err.Error())
    }
    return bytes, nil
}

func xerDecode{{.MessageName}}(bytes []byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
    unsafePtr, err := decodeXer(bytes, &C.asn_DEF_{{dashToUnderscore .CstructName}})
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
        return nil, fmt.Errorf("pointer decoded from XER is nil")
    }
    return decode{{.MessageName}}((*C.{{dashToUnderscore .CstructName}}_t)(unsafePtr))
}

func perDecode{{.MessageName}}(bytes []byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
    unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_{{dashToUnderscore .CstructName}})
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
       return nil, fmt.Errorf("pointer decoded from PER is nil")
    }
    return decode{{.MessageName}}((*C.{{dashToUnderscore .CstructName}}_t)(unsafePtr))
}

func new{{.MessageName}}({{lowCaseFirstLetter .MessageName}} *{{.ProtoFileName}}.{{.MessageName}}) (*C.{{dashToUnderscore .CstructName}}_t, error) {
{{if .FieldList.SingleItem}}
{{if .Optional}}{{ template "OPTIONAL_ENCODE_SINGLE" . }}{{end}}
{{if .OneOf}}{{ template "ONEOF_ENCODE" . }}{{end}}
{{if .Repeated}}{{ template "REPEATED_ENCODE" . }}{{end}}
{{else}}
{{if .OneOf}}{{ template "ONEOF_ENCODE" . }}{{else}}
{{if .Repeated}}{{ template "REPEATED_ENCODE" . }}{{end}}
{{if .Optional}}{{ template "OPTIONAL_ENCODE" . }}{{end}}
{{lowCaseFirstLetter .MessageName}}C := C.{{dashToUnderscore .CstructName}}_t{
//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file{{ $optionalFields := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields }}
{{.CstructLeafName}}: {{lowCaseFirstLetter .FieldName}}C,{{end}}
{{ $repeatedFields := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields }}{{.CstructLeafName}}: *{{lowCaseFirstLetter .FieldName}}C,
{{end}}
{{ $oneOfFields := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneOfFields }} //{{.CstructLeafName}}: {{lowCaseFirstLetter .FieldName}}C,{{end}}
}{{end}}
{{end}}
    return &{{lowCaseFirstLetter .MessageName}}C, nil
}

func decode{{.MessageName}}({{lowCaseFirstLetter .MessageName}}C *C.{{dashToUnderscore .CstructName}}_t) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
{{if .FieldList.SingleItem}}
{{if .Optional}}{{ template "OPTIONAL_DECODE_SINGLE" . }}{{end}}
{{if .OneOf}}{{ template "ONEOF_DECODE" . }}{{end}}
{{if .Repeated}}{{ template "REPEATED_DECODE" . }}{{end}}
{{else}}
{{if .OneOf}}{{ template "ONEOF_DECODE" . }}{{else}}
{{if .Optional}}{{ template "OPTIONAL_DECODE" . }}{{end}}
{{lowCaseFirstLetter .MessageName}} := {{.ProtoFileName}}.{{.MessageName}}{
//ToDo - check whether pointers passed correctly with regard to Protobuf's definition{{ $optionalFields1 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields1 }}
{{upperCaseFirstLetter .FieldName}}: {{lowCaseFirstLetter .FieldName}},{{end}}
{{ $repeatedFields1 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields1 }}{{upperCaseFirstLetter .FieldName}}: make([]*{{.ProtoFileName}}.{{.DataType}}, 0), //ToDo - Check if protobuf structure is implemented correctly (mainly naming)
{{end}}
{{ $oneOfFields1 := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneOfFields1 }} //{{upperCaseFirstLetter.FieldName}}: {{lowCaseFirstLetter .FieldName}}, {{end}}}
{{if .Repeated}}{{ template "REPEATED_DECODE" . }}{{end}}{{end}}
{{end}}
    return &{{lowCaseFirstLetter .MessageName}}, nil
}

{{if .Repeated}}{{ template "DECODE_BYTES" . }}{{end}}




{{ define "ONEOF_ENCODE" }}
{{ $fields := .FieldList.OneOfField }}
var pr C.{{dashToUnderscore .CstructName}}_PR
choiceC := [8]byte{} //ToDo - Check if number of bytes is sufficient

switch choice := {{lowCaseFirstLetter .MessageName}}.{{.MessageName}}.(type) { {{ range $fieldIndex, $field := $fields }}
case {{.ProtoFileName}}.{{.MessageName}}_{{.FieldName}}:
pr = C.{{dashToUnderscore .CstructName}}_PR_{{dashToUnderscore .CstructLeafName}} //ToDo - Check if it's correct PR

im, err := {{encodeDataType .DataType}}(choice.{{.CstructLeafName}})
if err != nil {
return nil, fmt.Errorf("new{{.MessageName}}() %s", err.Error())
}
binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im)))){{end}}
default:
return nil, fmt.Errorf("new{{.MessageName}}() %T not yet implemented", choice)
}

{{lowCaseFirstLetter .MessageName}}C := C.{{dashToUnderscore .CstructName}}_t{
present: pr,
choice:  choiceC,
}{{end}}

{{ define "REPEATED_ENCODE" }}
{{ $fields := .FieldList.RepeatedField }}
{{ range $fieldIndex, $field := $fields }}
{{lowCaseFirstLetter .FieldName}}C := new(C.struct_{{dashToUnderscore .CstructName}}_{{.CstructLeafName}}) //ToDo - double-check if struct has correct name
for _, {{lowCaseFirstLetter .FieldName}}Item := range {{.ProtoFileName}}.{{.MessageName}}.Get{{.FieldName}}() { //ToDo - Verify if GetSmth() function is called correctly
{{lowCaseFirstLetter .FieldName}}ItemC, err := {{encodeDataType .DataType}}({{lowCaseFirstLetter .FieldName}}Item)
if err != nil {
return nil, fmt.Errorf("new{{.MessageName}}() %s", err.Error())
}
//ToDo - declare pmContainersListC on top of it...
if _, err = C.asn_sequence_add(unsafe.Pointer({{lowCaseFirstLetter .FieldName}}C), unsafe.Pointer({{lowCaseFirstLetter .FieldName}}ItemC)); err != nil {
return nil, err
}
}{{end}}
{{end}}

{{ define "OPTIONAL_ENCODE" }}{{ $fields := .FieldList.OptionalField }}
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .FieldName}}C := {{encodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}.{{upperCaseFirstLetter .FieldName}}){{else}}
{{lowCaseFirstLetter .FieldName}}C, err := {{encodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}.{{upperCaseFirstLetter .FieldName}})
if err != nil {
return nil, fmt.Errorf("new{{.MessageName}}() %s", err.Error())
}
{{end}}{{end}}{{end}}

{{ define "OPTIONAL_ENCODE_SINGLE" }}{{ $fields := .FieldList.OptionalField }}
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .MessageName}}C := {{encodeDataType .DataType}}({{lowCaseFirstLetter .MessageName}}.{{upperCaseFirstLetter .FieldName}}){{else}}
{{lowCaseFirstLetter .MessageName}}C, err := {{encodeDataType .DataType}}({{lowCaseFirstLetter .MessageName}}.{{upperCaseFirstLetter .FieldName}})
if err != nil {
return nil, fmt.Errorf("new{{.MessageName}}() %s", err.Error())
}
{{end}}{{end}}{{end}}

{{ define "ONEOF_DECODE" }}
{{ $fields := .FieldList.OneOfField }}
//This is Decoder part (OneOf)
{{lowCaseFirstLetter .MessageName}} := new({{.ProtoFileName}}.{{.MessageName}})

switch {{lowCaseFirstLetter .MessageName}}C.present { {{ range $fieldIndex, $field := $fields }}
case C.{{dashToUnderscore .CstructName}}_PR_{{dashToUnderscore .CstructLeafName}}:
{{lowCaseFirstLetter .MessageName}}structC, err := {{decodeDataType .DataType}}Bytes({{lowCaseFirstLetter .MessageName}}C.choice) //ToDo - Verify if decodeSmthBytes function exists
if err != nil {
return nil, fmt.Errorf("decode{{.MessageName}}() %s", err.Error())
}
{{lowCaseFirstLetter .MessageName}}.{{.MessageName}} = &{{.ProtoFileName}}.{{.MessageName}}_{{upperCaseFirstLetter .FieldName}}{
{{upperCaseFirstLetter .FieldName}}: {{lowCaseFirstLetter .MessageName}}structC,
}{{end}}
default:
return nil, fmt.Errorf("decode{{.MessageName}}() %v not yet implemented", {{lowCaseFirstLetter .MessageName}}C.present)
}{{ end }}

{{ define "REPEATED_DECODE" }}
{{ $fields := .FieldList.RepeatedField }}
var ieCount int
{{ range $fieldIndex, $field := $fields }}ieCount = int({{lowCaseFirstLetter .MessageName}}C.{{.CstructLeafName}}.list.count)
for i := 0; i < ieCount; i++ {
offset := unsafe.Sizeof(unsafe.Pointer({{lowCaseFirstLetter .MessageName}}C.{{.CstructLeafName}}.list.array)) * uintptr(i)
ieC := *(**C.{{.DataType}}_t)(unsafe.Pointer(uintptr(unsafe.Pointer({{lowCaseFirstLetter .MessageName}}C.{{.CstructLeafName}}.list.array)) + offset))
ie, err := {{decodeDataType .DataType}}(ieC)
if err != nil {
return nil, fmt.Errorf("decode{{.MessageName}}() %s", err.Error())
}
{{lowCaseFirstLetter .MessageName}}.{{upperCaseFirstLetter .FieldName}} = append({{lowCaseFirstLetter .MessageName}}.{{upperCaseFirstLetter .FieldName}}, ie)
}
{{end}}{{end}}

{{ define "DECODE_BYTES" }}
func decode{{.MessageName}}Bytes(array [8]byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) { //ToDo - Check addressing correct structure in Protobuf
{{lowCaseFirstLetter .MessageName}}C := (*C.{{dashToUnderscore .MessageName}}_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
result, err := decode{{.MessageName}}({{lowCaseFirstLetter .MessageName}}C)
if err != nil {
return nil, fmt.Errorf("decode{{.MessageName}}Bytes() %s", err.Error())
}

return result, nil
}
{{end}}

{{ define "OPTIONAL_DECODE" }}
{{ $fields := .FieldList.OptionalField }}
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .FieldName}} := {{decodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}C.{{.CstructLeafName}}){{else}}
{{lowCaseFirstLetter .FieldName}}, err := {{decodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}C.{{.CstructLeafName}})
if err != nil {
return nil, fmt.Errorf("decode{{.MessageName}}() %s", err.Error())
}
{{end}}{{end}}{{end}}

{{ define "OPTIONAL_DECODE_SINGLE" }}
{{ $fields := .FieldList.OptionalField }}
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .MessageName}} := {{.ProtoFileName}}.{{.MessageName}}{
{{upperCaseFirstLetter .FieldName}}: {{decodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}C),
}{{else}}
{{lowCaseFirstLetter .MessageName}} := new({{.ProtoFileName}}.{{.MessageName}})
{{lowCaseFirstLetter .MessageName}}Value, err := {{decodeDataType .DataType}}({{lowCaseFirstLetter .MessageName}}C)
if err != nil {
return nil, fmt.Errorf("decode{{.MessageName}}() %s", err.Error())
}
{{lowCaseFirstLetter .MessageName}}.{{upperCaseFirstLetter .FieldName}} = {{lowCaseFirstLetter .MessageName}}Value
{{end}}{{end}}{{end}}