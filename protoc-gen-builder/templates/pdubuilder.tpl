// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
    {{.Imports}}
    "github.com/onosproject/onos-lib-go/pkg/errors"
)

{{if .MessagePresence}}{{ $inst := .Messages }}{{ range $fieldIndex, $field := $inst }}
func Create{{.FunctionName}}({{ $inst := .Items }}{{ range $innerFieldIndex, $innerField := $inst }}{{.VariableName}} {{.ItemType}}, {{end}}) (*{{.PackageName}}.{{.FunctionOutputType}}, error) {

    msg := &{{.PackageName}}.{{.FunctionOutputType}}{}
    {{ $inst := .Items }}{{ range $innerFieldIndex, $innerField := $inst }}msg.{{.FieldName}} = {{.VariableName}}
    {{end}}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("Create{{.FunctionName}}() error validating PDU %s", err.Error())
	}

	return msg, nil
}
{{end}}{{end}}

{{if .OneOfPresence}}{{ $inst := .OneOfs }}{{ range $fieldIndex, $field := $inst }}
func Create{{.FunctionName}}({{.VariableName}} {{.VariableType}}) (*{{.PackageName}}.{{.FunctionOutputType}}, error) {

    item := &{{.PackageName}}.{{.FunctionOutputType}}{
        {{.FunctionOutputType}}: &{{.PackageName}}.{{.FunctionOutputType}}_{{.ItemName}}{
            {{.ItemName}}: {{.VariableName}},
        },
    }

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("Create{{.FunctionName}}() error validating PDU %s", err.Error())
	}

    return item, nil
}{{end}}{{end}}

{{if .EnumPresence}}{{ $inst := .Enums }}{{ range $fieldIndex, $field := $inst }}
func Create{{.FunctionName}}() {{.FunctionOutputType}} {
    return {{.PackageName}}.{{.Item}}
}{{end}}{{end}}