// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package {{.PackageName}}

import (
    "reflect"
    {{.Imports}}
)

var {{.MapName}}Choicemap = map[string]map[int]reflect.Type{ {{ $ch := .Choices }}{{ range $fieldIndex, $field := $ch }}{{ $ie := .Items }}{{ range $fieldIndex1, $field1 := $ie }}
    "{{.ChoiceName}}":{ {{ $lf := .Leafs }}{{ range $innerFieldIndex, $innerField := $lf }}
    {{.Index}}:reflect.TypeOf({{.ProtoFileName}}.{{.LeafName}}{}),{{end}}
    },{{end}}{{end}}
}
