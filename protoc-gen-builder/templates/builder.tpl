// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package {{.PackageName}}

import ({{.Imports}})

{{ $inst := .Instances }}{{ range $innerFieldIndex, $innerField := $inst }}
func (m *{{.Instance}}) Set{{.FunctionName}}({{.VariableName}} {{.ItemType}}) *{{.Instance}} {
	m.{{.ItemName}} = {{.VariableNamePtr}}
	return m
}
{{end}}