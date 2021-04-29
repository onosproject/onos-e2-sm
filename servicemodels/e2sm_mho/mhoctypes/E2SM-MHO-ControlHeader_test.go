// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmMhoControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1

	e2SmMhoControlHeader, err := pdubuilder.CreateE2SmMhoControlHeader(controlMessagePriority)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmMhoControlHeader(e2SmMhoControlHeader)
	assert.NilError(t, err)
	assert.Equal(t, 232, len(xer))
	t.Logf("E2SM-MHO-ControlHeader XER\n%s", string(xer))
}

func Test_XerDecodeE2SmMhoControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1

	e2SmMhoControlHeader, err := pdubuilder.CreateE2SmMhoControlHeader(controlMessagePriority)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmMhoControlHeader(e2SmMhoControlHeader)
	assert.NilError(t, err)
	assert.Equal(t, 232, len(xer))
	t.Logf("E2SM-MHO-ControlHeader XER\n%s", string(xer))

	result, err := XerDecodeE2SmMhoControlHeader(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-MHO-ControlHeader decoded XER is\n%v", result)

	//assert.Equal(t, e2SmMhoControlHeader.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}

func Test_PerDecodeE2SmMhoControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1

	e2SmMhoControlHeader, err := pdubuilder.CreateE2SmMhoControlHeader(controlMessagePriority)
	assert.NilError(t, err)

	per, err := PerEncodeE2SmMhoControlHeader(e2SmMhoControlHeader)
	assert.NilError(t, err)
	t.Logf("E2SM-MHO-ControlHeader PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmMhoControlHeader(per)
	assert.NilError(t, err)
	t.Logf("E2SM-MHO-ControlHeader decoded PER is\n%v", result)

	//assert.Equal(t, e2SmMhoControlHeader.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}
