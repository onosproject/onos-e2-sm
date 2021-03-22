// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package {{.ProtoFileName}}

type {{.MessageName}}Builder interface {
New{{.MessageName}}()
Get{{.MessageName}}()
//ToDo - check whether generated Set & Get functions correspond to actual Protobuf structure
{{if .Optional}}{{ $optionalFields1 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields1 }}
Set{{upperCaseFirstLetter .FieldName}}()
Get{{upperCaseFirstLetter .FieldName}}()
{{end}}{{ $repeatedFields1 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields1 }}
Set{{upperCaseFirstLetter .FieldName}}()
Get{{upperCaseFirstLetter .FieldName}}()
{{end}}{{end}}{{if .OneOf}}{{ $oneofFields1 := .FieldList.OneOfField }}{{ range $fieldIndex, $field := $oneofFields1 }}
Set{{upperCaseFirstLetter .FieldName}}()
Get{{upperCaseFirstLetter .FieldName}}()
{{end}}{{end}}
}

func New{{.MessageName}}() *{{.MessageName}} {
return &{{.MessageName}}{}
}

func (b *{{.MessageName}}) Get{{.MessageName}}() *{{.MessageName}} {
return b
}

//ToDo - check whether generated Set & Get functions correspond to actual Protobuf structure
{{if .Optional}}{{ $optionalFields1 := .FieldList.OptionalField }}{{ range $fieldIndex, $field := $optionalFields1 }}
func (b *{{.MessageName}}) Set{{upperCaseFirstLetter .FieldName}}(input *{{upperCaseFirstLetter .FieldName}}) *{{.MessageName}} {
b.{{upperCaseFirstLetter .FieldName}} = input
return b
}

func (b *{{.MessageName}}) Get{{upperCaseFirstLetter .FieldName}}() *{{.MessageName}} {
return b.{{upperCaseFirstLetter .FieldName}}
}
{{end}}
{{ $repeatedFields1 := .FieldList.RepeatedField }}{{ range $fieldIndex, $field := $repeatedFields1 }}
func (b *{{.MessageName}}) Set{{upperCaseFirstLetter .FieldName}}(input *{{upperCaseFirstLetter .FieldName}}) *{{.MessageName}} {
b.{{upperCaseFirstLetter .FieldName}} = input
return b
}

func (b *{{.MessageName}}) Get{{upperCaseFirstLetter .FieldName}}() *{{.MessageName}} {
return b.{{upperCaseFirstLetter .FieldName}}
}
{{end}}
{{end}}{{if .OneOf}}{{ $oneofFields1 := .FieldList.OneOfField }}
{{ range $fieldIndex, $field := $oneofFields1 }}
func (b *{{.MessageName}}) Set{{upperCaseFirstLetter .FieldName}}(input *{{upperCaseFirstLetter .FieldName}}) *{{.MessageName}} {
b.{{upperCaseFirstLetter .FieldName}} = input
return b
}

func (b *{{.MessageName}}) Get{{upperCaseFirstLetter .FieldName}}() *{{.MessageName}} {
return b.{{upperCaseFirstLetter .FieldName}}
}
{{end}}
{{end}}






