// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctype

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "{{underscoreToDash .CstructName}}.h" //ToDo: if there is a Repeted option, it would require linking additional structure file definition (the one above/before)
import "C"
import (
    "fmt"
            {{.ProtoFileName}} "github.com/onosproject/onos-e2-sm/servicemodels/{{.ProtoFileName}}/v1beta1/{{underscoreToDash .ProtoFileName}}" //ToDo: Make imports more dynamic
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

{{if .OneOf}}{{ template "ONEOF_ENCODE" . }}{{end}}
{{if .Repeated}}{{ template "REPEATED_ENCODE" . }}{{end}}
{{if .Optional}}    {{ template "OPTIONAL_ENCODE" . }}{{end}}
{{lowCaseFirstLetter .MessageName}}C := C.{{dashToUnderscore .CstructName}}_t{
//ToDo: check whether pointers passed correctly with regard to C-struct's definition .h file{{ $optionalFields := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields }}
{{.CstructLeafName}}: {{lowCaseFirstLetter .FieldName}}C,{{end}}
{{ $repeatedFields := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields }}
{{.CstructLeafName}}: {{lowCaseFirstLetter .FieldName}}C,{{end}}
{{ $oneOfFields := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneOfFields }}
{{.CstructLeafName}}: {{lowCaseFirstLetter .FieldName}}C,{{end}}
}
    return &{{lowCaseFirstLetter .MessageName}}C, nil
}

func decode{{.MessageName}}({{lowCaseFirstLetter .MessageName}}C *C.{{dashToUnderscore .CstructName}}_t) (*{{.ProtoFileName}}.{{.MessageName}}, error) {

{{if .OneOf}}{{ template "ONEOF_DECODE" . }}{{end}}
{{if .Repeated}}{{ template "REPEATED_DECODE" . }}{{end}}
{{if .Optional}}{{ template "OPTIONAL_DECODE" . }}{{end}}

{{lowCaseFirstLetter .MessageName}} := {{.ProtoFileName}}.{{.MessageName}}{
//ToDo: check whether pointers passed correctly with regard to Protobuf's definition{{ $optionalFields1 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields1 }}
{{.FieldName}}: {{lowCaseFirstLetter .FieldName}},{{end}}
{{ $repeatedFields1 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields1 }}
{{.FieldName}}: {{lowCaseFirstLetter .FieldName}},{{end}}
{{ $oneOfFields1 := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneOfFields1 }}
{{.FieldName}}: {{lowCaseFirstLetter .FieldName}},{{end}}
}

    return &{{lowCaseFirstLetter .MessageName}}, nil
}

{{ define "ONEOF_ENCODE" }}
{{ $fields := .FieldList.OneOfField }}
//This is Encoder part (OneOf)
{{ range $fieldIndex, $field := $fields }}
OneOf field #{{ $fieldIndex }}: {{ $field.FieldName }}{{end}}
{{end}}

{{ define "REPEATED_ENCODE" }}
{{ $fields := .FieldList.RepeatedField }}
//This is Encoder part (Repeated)
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .FieldName}} := {{encodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}.{{.FieldName}}){{else}}
{{lowCaseFirstLetter .FieldName}}C, err := {{encodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}.{{.FieldName}})
    if err != nil {
        return nil, fmt.Errorf("encode{{.MessageName}}() %s", err.Error())
    }{{end}}{{end}}{{end}}

{{ define "OPTIONAL_ENCODE" }}{{ $fields := .FieldList.OptionalField }}
//This is Encoder part (Optional)
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .FieldName}}C := {{encodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}.{{.FieldName}}){{else}}
{{lowCaseFirstLetter .FieldName}}C, err := {{encodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}.{{.FieldName}})
    if err != nil {
        return nil, fmt.Errorf("encode{{.MessageName}}() %s", err.Error())
    }{{end}}{{end}}{{end}}

{{ define "ONEOF_DECODE" }}
{{ $fields := .FieldList.OneOfField }}
//This is Decoder part (OneOf)
{{ range $fieldIndex, $field := $fields }}
OneOf field #{{ $fieldIndex }}: {{ $field.FieldName }}{{end}}
{{ end }}

{{ define "REPEATED_DECODE" }}
{{ $fields := .FieldList.RepeatedField }}
//This is Decoder part (Repeated)
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .FieldName}} := {{decodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}C.{{.CstructLeafName}}) {{else}}
{{lowCaseFirstLetter .FieldName}}, err := {{decodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}C.{{.CstructLeafName}})
    if err != nil {
        return nil, fmt.Errorf("decode{{.MessageName}}() %s", err.Error())
    }{{end}}{{end}}{{end}}

{{ define "OPTIONAL_DECODE" }}
{{ $fields := .FieldList.OptionalField }}
//This is Decoder part (Optional)
{{ range $fieldIndex, $field := $fields }}{{if checkElementaryType .DataType}}
{{lowCaseFirstLetter .FieldName}} := {{decodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}C.{{.CstructLeafName}}){{else}}
{{lowCaseFirstLetter .FieldName}}, err := {{decodeDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}C.{{.CstructLeafName}})
    if err != nil {
        return nil, fmt.Errorf("decode{{.MessageName}}() %s", err.Error())
    }{{end}}{{end}}{{end}}