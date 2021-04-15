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
    "encoding/binary"
    {{.ProtoFileName}} "github.com/onosproject/onos-e2-sm/servicemodels/{{.ProtoFileName}}{{.FullPackageName}}" //ToDo - Make imports more dynamic
    "unsafe"
)

func xerEncode{{.GlobalName}}({{lowCaseFirstLetter .GlobalName}} *{{.ProtoFileName}}.{{.MessageName}}) ([]byte, error) {
    {{lowCaseFirstLetter .GlobalName}}CP, err := new{{.GlobalName}}({{lowCaseFirstLetter .MessageName}})
    if err != nil {
        return nil, fmt.Errorf("xerEncode{{.GlobalName}}() %s", err.Error())
    }

    bytes, err := encodeXer(&C.asn_DEF_{{dashToUnderscore .CstructName}}, unsafe.Pointer({{lowCaseFirstLetter .GlobalName}}CP))
    if err != nil {
        return nil, fmt.Errorf("xerEncode{{.GlobalName}}() %s", err.Error())
    }
    return bytes, nil
}

func perEncode{{.GlobalName}}({{lowCaseFirstLetter .GlobalName}} *{{.ProtoFileName}}.{{.MessageName}}) ([]byte, error) {
    {{lowCaseFirstLetter .GlobalName}}CP, err := new{{.GlobalName}}({{lowCaseFirstLetter .GlobalName}})
    if err != nil {
        return nil, fmt.Errorf("perEncode{{.GlobalName}}() %s", err.Error())
    }

    bytes, err := encodePerBuffer(&C.asn_DEF_{{dashToUnderscore .CstructName}}, unsafe.Pointer({{lowCaseFirstLetter .GlobalName}}CP))
    if err != nil {
        return nil, fmt.Errorf("perEncode{{.GlobalName}}() %s", err.Error())
    }
    return bytes, nil
}

func xerDecode{{.GlobalName}}(bytes []byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
    unsafePtr, err := decodeXer(bytes, &C.asn_DEF_{{dashToUnderscore .CstructName}})
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
        return nil, fmt.Errorf("pointer decoded from XER is nil")
    }
    return decode{{.GlobalName}}((*C.{{dashToUnderscore .CstructName}}_t)(unsafePtr))
}

func perDecode{{.GlobalName}}(bytes []byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
    unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_{{dashToUnderscore .CstructName}})
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
       return nil, fmt.Errorf("pointer decoded from PER is nil")
    }
    return decode{{.GlobalName}}((*C.{{dashToUnderscore .CstructName}}_t)(unsafePtr))
}

func new{{.GlobalName}}({{lowCaseFirstLetter .GlobalName}} *{{.ProtoFileName}}.{{.MessageName}}) (*C.{{dashToUnderscore .CstructName}}_t, error) {
{{if .FieldList.SingleItem}}
{{if .Optional}}{{ template "OPTIONAL_ENCODE_SINGLE" . }}{{end}}
{{if .OneOf}}{{ template "ONEOF_ENCODE" . }}{{end}}
{{if .Repeated}}{{ template "REPEATED_ENCODE_SINGLE" . }}{{end}}
{{else}}
{{if .OneOf}}{{ template "ONEOF_ENCODE" . }}{{else}}
var err error
{{lowCaseFirstLetter .GlobalName}}C := C.{{dashToUnderscore .CstructName}}_t{}
{{if .Repeated}}{{ template "REPEATED_ENCODE" . }}{{end}}
{{if .Optional}}{{ template "OPTIONAL_ENCODE" . }}{{end}}
//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file{{ $optionalFields := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields }}
{{if not .Optionality}}{{lowCaseFirstLetter .GlobalName}}C.{{.CstructLeafName}} = {{lowCaseFirstLetter .FieldName}}C{{end}}{{end}}
{{ $repeatedFields := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields }}{{lowCaseFirstLetter .GlobalName}}C.{{.CstructLeafName}} = {{lowCaseFirstLetter .FieldName}}C
{{end}}
{{ $oneOfFields := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneOfFields }} //{{lowCaseFirstLetter .GlobalName}}C.{{.CstructLeafName}} = {{lowCaseFirstLetter .FieldName}}C{{end}}
{{end}}
{{end}}
    return &{{lowCaseFirstLetter .GlobalName}}C, nil
}

func decode{{.GlobalName}}({{lowCaseFirstLetter .GlobalName}}C *C.{{dashToUnderscore .CstructName}}_t) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
{{if .FieldList.SingleItem}}
{{if .Optional}}{{ template "OPTIONAL_DECODE_SINGLE" . }}{{end}}
{{if .OneOf}}{{ template "ONEOF_DECODE" . }}{{end}}
{{if .Repeated}}{{ template "REPEATED_DECODE_SINGLE" . }}{{end}}
{{else}}
{{if .OneOf}}{{ template "ONEOF_DECODE" . }}{{else}}
var err error
{{lowCaseFirstLetter .GlobalName}} := {{.ProtoFileName}}.{{.MessageName}}{
//ToDo - check whether pointers passed correctly with regard to Protobuf's definition{{ $optionalFields1 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields1 }}
//{{upperCaseFirstLetter .FieldName}}: {{lowCaseFirstLetter .FieldName}},{{end}}
{{ $repeatedFields1 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields1 }}{{upperCaseFirstLetter .FieldName}}: make([]*{{.ProtoFileName}}.{{.DataType}}, 0), //ToDo - Check if protobuf structure is implemented correctly (mainly naming)
{{end}}
{{ $oneOfFields1 := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneOfFields1 }} //{{upperCaseFirstLetter .FieldName}}: {{lowCaseFirstLetter .FieldName}}, {{end}}}
{{if .Optional}}{{ template "OPTIONAL_DECODE" . }}{{end}}
{{if .Repeated}}{{ template "REPEATED_DECODE" . }}{{end}}{{end}}
{{end}}
    return &{{lowCaseFirstLetter .GlobalName}}, nil
}

{{ template "DECODE_BYTES" . }}




{{ define "ONEOF_ENCODE" }}
{{ $fields := .FieldList.OneOfField }}
var pr C.{{dashToUnderscore .CstructName}}_PR //ToDo - verify correctness of the name
choiceC := [8]byte{} //ToDo - Check if number of bytes is sufficient

switch choice := {{lowCaseFirstLetter .GlobalName}}.{{.MessageName}}.(type) { {{ range $fieldIndex, $field := $fields }}
case *{{.ProtoFileName}}.{{.MessageName}}_{{upperCaseFirstLetter .FieldName}}:
pr = C.{{dashToUnderscore .CstructName}}_PR_{{dashToUnderscore .CstructLeafName}} //ToDo - Check if it's correct PR's name

im, err := {{encodeDataType .DataType}}(choice.{{upperCaseFirstLetter .FieldName}})
if err != nil {
return nil, fmt.Errorf("{{encodeDataType .DataType}}() %s", err.Error())
}
binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im)))){{end}}
default:
return nil, fmt.Errorf("new{{.GlobalName}}() %T not yet implemented", choice)
}

{{lowCaseFirstLetter .GlobalName}}C := C.{{dashToUnderscore .CstructName}}_t{
present: pr,
choice:  choiceC,
}{{end}}

{{ define "REPEATED_ENCODE" }}{{ $fields := .FieldList.RepeatedField }}
{{ range $fieldIndex, $field := $fields }}
{{if .Optionality}}{{template "OPTIONALITY_ENCODE_BEGIN" . }}{{end}}
{{lowCaseFirstLetter .FieldName}}C := new(C.struct_{{dashToUnderscore .CstructName}}__{{.CstructLeafName}}) //ToDo - verify correctness of the variable's name
for _, {{lowCaseFirstLetter .FieldName}}Item := range {{lowCaseFirstLetter .VariableName}}.Get{{upperCaseFirstLetter .FieldName}}() { //ToDo - Verify if GetSmth() function is called correctly
{{lowCaseFirstLetter .FieldName}}ItemC, err := {{encodeDataType .DataType}}({{lowCaseFirstLetter .FieldName}}Item)
if err != nil {
return nil, fmt.Errorf("{{encodeDataType .DataType}}() %s", err.Error())
}
if _, err = C.asn_sequence_add(unsafe.Pointer({{lowCaseFirstLetter .FieldName}}C), unsafe.Pointer({{lowCaseFirstLetter .FieldName}}ItemC)); err != nil {
return nil, err
}
}
{{if .Optionality}}{{template "OPTIONALITY_ENCODE_END" . }}{{end}}{{end}}{{end}}

{{ define "REPEATED_ENCODE_SINGLE" }}{{ $fields := .FieldList.RepeatedField }}
{{ range $fieldIndex, $field := $fields }}
{{lowCaseFirstLetter .VariableName}}C := new(C.{{dashToUnderscore .CstructName}}_t) //ToDo - verify correctness of the variable's name
for _, ie := range {{lowCaseFirstLetter .VariableName}}.Get{{upperCaseFirstLetter .FieldName}}() { //ToDo - Verify if GetSmth() function is called correctly
ieC, err := {{encodeDataType .DataType}}(ie)
if err != nil {
return nil, fmt.Errorf("{{encodeDataType .DataType}}() %s", err.Error())
}
if _, err = C.asn_sequence_add(unsafe.Pointer({{lowCaseFirstLetter .VariableName}}C), unsafe.Pointer(ieC)); err != nil {
return nil, err
}
}{{end}}{{end}}

{{ define "OPTIONAL_ENCODE" }}{{ $fields := .FieldList.OptionalField }}
{{ range $fieldIndex, $field := $fields }}{{if .Optionality}}{{template "OPTIONALITY_ENCODE_BEGIN" . }}{{end}}{{if checkElementaryType .DataType}}
{{doLinting .FieldName}}C := {{encodeDataType .DataType}}({{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}}){{else}}
{{doLinting .FieldName}}C, err := {{encodeDataType .DataType}}({{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}})
if err != nil {
return nil, fmt.Errorf("{{encodeDataType .DataType}}() %s", err.Error())
}
{{end}}{{if .Optionality}}{{template "OPTIONALITY_ENCODE_END" . }}{{end}}{{end}}{{end}}

{{ define "OPTIONAL_ENCODE_SINGLE" }}{{ $fields := .FieldList.OptionalField }}
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .VariableName}}C := {{encodeDataType .DataType}}({{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}}){{else}}
{{lowCaseFirstLetter .VariableName}}C, err := {{encodeDataType .DataType}}({{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}})
if err != nil {
return nil, fmt.Errorf("{{encodeDataType .DataType}}() %s", err.Error())
}
{{end}}{{end}}{{end}}

{{ define "ONEOF_DECODE" }}
{{ $fields := .FieldList.OneOfField }}
{{lowCaseFirstLetter .GlobalName}} := new({{.ProtoFileName}}.{{.MessageName}})

switch {{lowCaseFirstLetter .GlobalName}}C.present { {{ range $fieldIndex, $field := $fields }}
case C.{{dashToUnderscore .CstructName}}_PR_{{dashToUnderscore .CstructLeafName}}:
{{lowCaseFirstLetter .GlobalName}}structC, err := {{decodeDataType .DataType}}Bytes({{.VariableName}}C.choice) //ToDo - Verify if decodeSmthBytes function exists
if err != nil {
return nil, fmt.Errorf("{{decodeDataType .DataType}}Bytes() %s", err.Error())
}
{{lowCaseFirstLetter .GlobalName}}.{{upperCaseFirstLetter .FieldName}} = &{{.ProtoFileName}}.{{.MessageName}}_{{upperCaseFirstLetter .FieldName}}{
{{upperCaseFirstLetter .FieldName}}: {{lowCaseFirstLetter .GlobalName}}structC,
}{{end}}
default:
return nil, fmt.Errorf("decode{{.GlobalName}}() %v not yet implemented", {{lowCaseFirstLetter .GlobalName}}C.present)
}{{ end }}

{{ define "REPEATED_DECODE" }}
{{ $fields := .FieldList.RepeatedField }}
var ieCount int
{{ range $fieldIndex, $field := $fields }}
{{if .Optionality}}{{template "OPTIONALITY_DECODE_BEGIN" . }}{{end}}
ieCount = int({{lowCaseFirstLetter .VariableName}}C.{{.CstructLeafName}}.list.count)
for i := 0; i < ieCount; i++ {
offset := unsafe.Sizeof(unsafe.Pointer({{lowCaseFirstLetter .VariableName}}C.{{.CstructLeafName}}.list.array)) * uintptr(i)
ieC := *(**C.{{.DataType}}_t)(unsafe.Pointer(uintptr(unsafe.Pointer({{lowCaseFirstLetter .VariableName}}C.{{.CstructLeafName}}.list.array)) + offset))
ie, err := {{decodeDataType .DataType}}(ieC)
if err != nil {
return nil, fmt.Errorf("{{decodeDataType .DataType}}() %s", err.Error())
}
{{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}} = append({{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}}, ie)
}
{{if .Optionality}}{{template "OPTIONALITY_DECODE_END" . }}{{end}}{{end}}{{end}}

{{ define "REPEATED_DECODE_SINGLE" }}
{{ $fields := .FieldList.RepeatedField }}
var ieCount int
{{ range $fieldIndex, $field := $fields }}
{{lowCaseFirstLetter .VariableName}} := {{.ProtoFileName}}.{{.MessageName}}{}
{{if .Optionality}}{{template "OPTIONALITY_DECODE_BEGIN" . }}{{end}}
ieCount = int({{lowCaseFirstLetter .VariableName}}C.list.count)
for i := 0; i < ieCount; i++ {
offset := unsafe.Sizeof(unsafe.Pointer({{lowCaseFirstLetter .VariableName}}C.list.array)) * uintptr(i)
ieC := *(**C.{{.DataType}}_t)(unsafe.Pointer(uintptr(unsafe.Pointer({{lowCaseFirstLetter .VariableName}}C.list.array)) + offset))
ie, err := {{decodeDataType .DataType}}(ieC)
if err != nil {
return nil, fmt.Errorf("{{decodeDataType .DataType}}() %s", err.Error())
}
{{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}} = append({{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}}, ie)
}
{{if .Optionality}}{{template "OPTIONALITY_DECODE_END" . }}{{end}}{{end}}{{end}}

{{ define "DECODE_BYTES" }}
func decode{{.GlobalName}}Bytes(array [8]byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
{{lowCaseFirstLetter .GlobalName}}C := (*C.{{dashToUnderscore .CstructName}}_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

return decode{{.GlobalName}}({{lowCaseFirstLetter .GlobalName}}C)
}
{{end}}

{{ define "OPTIONAL_DECODE" }}
{{ $fields := .FieldList.OptionalField }}
{{ range $fieldIndex, $field := $fields }}{{if .Optionality}}{{template "OPTIONALITY_DECODE_BEGIN" . }}{{end}}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .GlobalName}}.{{upperCaseFirstLetter .FieldName}} = {{decodeDataType .DataType}}({{lowCaseFirstLetter .VariableName}}C.{{.CstructLeafName}}){{else}}
{{lowCaseFirstLetter .GlobalName}}.{{upperCaseFirstLetter .FieldName}}, err = {{decodeDataType .DataType}}({{lowCaseFirstLetter .VariableName}}C.{{.CstructLeafName}})
if err != nil {
return nil, fmt.Errorf("{{decodeDataType .DataType}}() %s", err.Error())
}
{{end}}{{if .Optionality}}{{template "OPTIONALITY_DECODE_END" . }}{{end}}{{end}}{{end}}

{{ define "OPTIONAL_DECODE_SINGLE" }}
{{ $fields := .FieldList.OptionalField }}
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .VariableName}} := {{.ProtoFileName}}.{{.MessageName}}{
{{upperCaseFirstLetter .FieldName}}: {{decodeDataType .DataType}}({{lowCaseFirstLetter .VariableName}}C),
}{{else}}
{{lowCaseFirstLetter .VariableName}} := new({{.ProtoFileName}}.{{.MessageName}})
{{lowCaseFirstLetter .VariableName}}Value, err := {{decodeDataType .DataType}}({{lowCaseFirstLetter .VariableName}}C)
if err != nil {
return nil, fmt.Errorf("{{decodeDataType .DataType}}() %s", err.Error())
}
{{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}} = {{lowCaseFirstLetter .VariableName}}Value
{{end}}{{end}}{{end}}

{{ define "OPTIONALITY_ENCODE_BEGIN" }}
//instance is optional
if {{lowCaseFirstLetter .VariableName}}.{{upperCaseFirstLetter .FieldName}} != nil {
{{end}}

{{ define "OPTIONALITY_ENCODE_END" }}
{{lowCaseFirstLetter .VariableName}}C.{{.CstructLeafName}} = {{doLinting .FieldName}}C
}{{end}}

{{ define "OPTIONALITY_DECODE_BEGIN" }}
//instance is optional
if {{lowCaseFirstLetter .VariableName}}C.{{.CstructLeafName}} != nil {
{{end}}

{{ define "OPTIONALITY_DECODE_END" }}
}{{end}}
