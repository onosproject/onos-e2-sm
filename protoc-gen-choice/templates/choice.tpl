// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package {{.ProtoFileName}}

import "reflect"

var Choicemap = map[string]map[int]reflect.Type{ {{ $ch := .Choices }}{{ range $fieldIndex, $field := $ch }}{{ $ie := .Items }}{{ range $fieldIndex1, $field1 := $ie }}
    "{{.ChoiceName}}":{ {{ $lf := .Leafs }}{{ range $innerFieldIndex, $innerField := $lf }}
        {{.Index}}:reflect.TypeOf({{.LeafName}}{}),{{end}}
    },{{end}}{{end}}
}
