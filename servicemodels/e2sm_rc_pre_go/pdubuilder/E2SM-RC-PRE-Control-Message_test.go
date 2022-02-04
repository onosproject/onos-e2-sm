// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/encoder"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreControlMsg(t *testing.T) {
	var ranParameterName = "PCI"
	var ranParameterValue int64 = 10
	var ranParameterID int32 = 1

	ranParameter, err := CreateRanParameterValueInt(ranParameterValue)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err := encoder.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRcPreControlMessage(per)
	assert.NilError(t, err)
	t.Logf("PER ControlMessage - decoded \n%v", result)
	assert.Equal(t, newE2SmRcPrePdu.String(), result.String())

	ranParameter = CreateRanParameterValueEnum(int32(ranParameterValue))
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err = encoder.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	result, err = encoder.PerDecodeE2SmRcPreControlMessage(per)
	assert.NilError(t, err)
	t.Logf("PER ControlMessage - decoded \n%v", result)
	assert.Equal(t, newE2SmRcPrePdu.String(), result.String())

	ranParameter = CreateRanParameterValueBool(true)
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err = encoder.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	result, err = encoder.PerDecodeE2SmRcPreControlMessage(per)
	assert.NilError(t, err)
	t.Logf("PER ControlMessage - decoded \n%v", result)
	assert.Equal(t, newE2SmRcPrePdu.String(), result.String())

	ranParameter = CreateRanParameterValueBitS(&asn1.BitString{
		Value: []byte{0x45, 0x87, 0x90},
		Len:   22,
	})
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err = encoder.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	result, err = encoder.PerDecodeE2SmRcPreControlMessage(per)
	assert.NilError(t, err)
	t.Logf("PER ControlMessage - decoded \n%v", result)
	assert.Equal(t, newE2SmRcPrePdu.String(), result.String())

	ranParameter = CreateRanParameterValueOctS("ONF")
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err = encoder.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	result, err = encoder.PerDecodeE2SmRcPreControlMessage(per)
	assert.NilError(t, err)
	t.Logf("PER ControlMessage - decoded \n%v", result)
	assert.Equal(t, newE2SmRcPrePdu.String(), result.String())

	ranParameter = CreateRanParameterValuePrtS("onf")
	newE2SmRcPrePdu, err = CreateE2SmRcPreControlMessage(ranParameterID, ranParameterName, ranParameter)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err = encoder.PerEncodeE2SmRcPreControlMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Message: \n%v", hex.Dump(per))

	result, err = encoder.PerDecodeE2SmRcPreControlMessage(per)
	assert.NilError(t, err)
	t.Logf("PER ControlMessage - decoded \n%v", result)
	assert.Equal(t, newE2SmRcPrePdu.String(), result.String())
}
