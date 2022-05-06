//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package {{.PackageName}}

// if necessary
import e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"

{{ $inst := .Instances }}{{ range $innerFieldIndex, $innerField := $inst }}
func (m *{{.Instance}}) Set{{.ItemName}}({{.VariableName}} {{.ItemType}}) *{{.Instance}} {
	m.{{.ItemName}} = {{.VariableNamePtr}}
	return m
}
{{end}}