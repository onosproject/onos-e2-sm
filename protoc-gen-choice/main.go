// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// ToDo: delete it (logs generation) once plugin is complete - necessity only in debugging
// Later on log files are stored to the /tmp/report.txt file (basically, it's ALMOST output of the Execute() function in module.go)

package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/onosproject/onos-e2-sm/protoc-gen-choice/generic"
)

func main() {
	g := pgs.Init(pgs.DebugMode())
	g.RegisterModule(generic.NewModule())
	g.RegisterPostProcessor(pgsgo.GoFmt())
	g.Render()
}
