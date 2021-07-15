// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreControlMsg(t *testing.T) {
	var ranParameterName = "PCI"
	var ranParameterValue uint32 = 10
	var ranParameterID int32 = 1

	ranParameter, err := CreateRanParameterValueInt(ranParameterValue)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err := rcprectypes.XerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Control Message: \n%s", string(xer))

	per, err := rcprectypes.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	ranParameter = CreateRanParameterValueEnum(int32(ranParameterValue))
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err = rcprectypes.XerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Control Message: \n%s", string(xer))

	per, err = rcprectypes.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	ranParameter = CreateRanParameterValueBool(true)
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err = rcprectypes.XerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Control Message: \n%s", string(xer))

	per, err = rcprectypes.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	ranParameter = CreateRanParameterValueBitS(&e2sm_rc_pre_v2.BitString{
		Value: []byte{0x45, 0x87, 0x90},
		Len:   22,
	})
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err = rcprectypes.XerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Control Message: \n%s", string(xer))

	per, err = rcprectypes.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	ranParameter = CreateRanParameterValueOctS("ONF")
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err = rcprectypes.XerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Control Message: \n%s", string(xer))

	per, err = rcprectypes.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	ranParameter = CreateRanParameterValuePrtS("onf")
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	xer, err = rcprectypes.XerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Control Message: \n%s", string(xer))

	per, err = rcprectypes.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))
}
