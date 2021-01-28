// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// To build new protobuf plugin run following command within onos-e2-sm/protoc-gen-cgo/ directory:
// go build -v -o ./protoc-gen-cgo

// I've put imports (required in *.proto files) in the same directory as a proto file.
// Run from onos-e2-sm/protoc-gen-cgo/ folder following command:
// protoc -I="." --plugin="./protoc-gen-cgo" --cgo_out="./generated" *.proto
// Or more specifically:
// protoc -I="." --plugin="./protoc-gen-cgo" --cgo_out="./generated" e2sm_fb_ies_v1.proto
// -I="..." specifies path to the imports in the .proto file(s)
// --plugin="..." specifies path to your custom plugin if it is not placed into one of the folders in $PATH
// --cgo_out="..." specifies path where to store generated files
// *.proto is a path to the source .proto files to process
//
// ToDo: delete it (logs generation) once plugin is complete - necessity only in debugging
// Later on log files are stored to the /tmp/report.txt file (basically, it's ALMOST output of the Execute() function in module.go)

package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/onosproject/onos-e2-sm/protoc-gen-cgo/generic"
)

func main() {
	g := pgs.Init(pgs.DebugMode())
	g.RegisterModule(generic.NewModule())
	//g.RegisterPostProcessor(generic.NewPostProcessor()) // Probably wouldn't need in the end
	g.RegisterPostProcessor(pgsgo.GoFmt())
	g.Render()
}
