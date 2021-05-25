// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmRcPreControlMessage(t *testing.T) {
	var ranParameterName = "PCI"
	var ranParameterValue uint32 = 10
	var ranParameterID int32 = 1

	ranParameter, err := pdubuilder.CreateRanParameterValueInt(ranParameterValue)
	assert.NilError(t, err)
	e2SmRcPreControlMessage, err := pdubuilder.CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlMessage(e2SmRcPreControlMessage)
	assert.NilError(t, err)
	//assert.Equal(t, 400, len(xer))
	t.Logf("E2SM-RC-PRE-ControlMessage XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreControlMessage(t *testing.T) {
	var ranParameterName = "PCI"
	var ranParameterValue uint32 = 10
	var ranParameterID int32 = 1

	ranParameter, err := pdubuilder.CreateRanParameterValueInt(ranParameterValue)
	assert.NilError(t, err)
	e2SmRcPreControlMessage, err := pdubuilder.CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlMessage(e2SmRcPreControlMessage)
	assert.NilError(t, err)
	//assert.Equal(t, 400, len(xer))
	t.Logf("E2SM-RC-PRE-ControlMessage XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreControlMessage(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-ControlMessage decoded XER is\n%v", result)

	//assert.Equal(t, e2SmRcPreControlMessage.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}

func Test_PerDecodeE2SmRcPreControlMessage(t *testing.T) {
	var ranParameterName = "PCI"
	var ranParameterValue uint32 = 10
	var ranParameterID int32 = 1

	ranParameter, err := pdubuilder.CreateRanParameterValueInt(ranParameterValue)
	assert.NilError(t, err)
	e2SmRcPreControlMessage, err := pdubuilder.CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)

	per, err := PerEncodeE2SmRcPreControlMessage(e2SmRcPreControlMessage)
	assert.NilError(t, err)
	assert.Equal(t, 11, len(per))
	t.Logf("E2SM-RC-PRE-ControlMessage PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmRcPreControlMessage(per)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-ControlMessage decoded XER is\n%v", result)

	//assert.Equal(t, e2SmRcPreControlMessage.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}
