// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package {{.ProtoFileName}}

type {{.GlobalName}}Builder interface {
New{{.GlobalName}}()
Get{{.GlobalName}}()
//ToDo - check whether generated Set & Get functions correspond to actual Protobuf structure
{{if .Optional}}{{ $optionalFields1 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields1 }}
Set{{doLinting .FieldName}}()
Get{{doLinting .FieldName}}()
{{end}}{{ $repeatedFields1 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields1 }}
Set{{doLinting .FieldName}}()
Get{{doLinting .FieldName}}()
{{end}}{{end}}{{if .OneOf}}{{ $oneofFields1 := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneofFields1 }}
Set{{doLinting .FieldName}}()
Get{{doLinting .FieldName}}()
{{end}}{{end}}
}

func New{{.GlobalName}}() *{{.MessageName}} {
return &{{.MessageName}}{}
}

func (b *{{.MessageName}}) Get{{.GlobalName}}() *{{.MessageName}} {
return b
}

//ToDo - check whether generated Set & Get functions correspond to actual Protobuf structure
{{if .Optional}}{{ $optionalFields1 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields1 }}
func (b *{{.MessageName}}) Set{{doLinting .FieldName}}(input *{{upperCaseFirstLetter .FieldName}}) *{{.MessageName}} {
b.{{upperCaseFirstLetter .FieldName}} = input
return b
}

func (b *{{.MessageName}}) Get{{doLinting .FieldName}}() *{{.MessageName}} {
return b.{{upperCaseFirstLetter .FieldName}}
}
{{end}}
{{ $repeatedFields1 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields1 }}
func (b *{{.MessageName}}) Set{{doLinting .FieldName}}(input *{{upperCaseFirstLetter .FieldName}}) *{{.MessageName}} {
b.{{upperCaseFirstLetter .FieldName}} = input
return b
}

func (b *{{.MessageName}}) Get{{doLinting .FieldName}}() *{{.MessageName}} {
return b.{{upperCaseFirstLetter .FieldName}}
}
{{end}}
{{end}}{{if .OneOf}}{{ $oneofFields1 := .FieldList.OneOfField }}
{{ range $fieldIndex, $field := $oneofFields1 }}
func (b *{{.MessageName}}) Set{{doLinting .FieldName}}(input *{{upperCaseFirstLetter .FieldName}}) *{{.MessageName}} {
b.{{upperCaseFirstLetter .FieldName}} = input
return b
}

func (b *{{.MessageName}}) Get{{doLinting .FieldName}}() *{{.MessageName}} {
return b.{{upperCaseFirstLetter .FieldName}}
}
{{end}}
{{end}}






