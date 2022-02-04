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

func TestE2SmRcPreControlHeader(t *testing.T) {
	var controlMessagePriority int32 = 1
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	cellID := asn1.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90, 0x00},
		Len:   36, //uint32
	}

	cgi, err := CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader()
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
	newE2SmRcPrePdu.GetControlHeaderFormat1().SetCGI(cgi).SetRicControlMessagePriority(controlMessagePriority)

	per, err := encoder.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded ControlHeader: %v", hex.Dump(per))

	resultPer, err := encoder.PerDecodeE2SmRcPreControlHeader(per)
	assert.NilError(t, err)
	t.Logf("PER ControlHeader - decoded \n%v", resultPer)
	assert.Equal(t, newE2SmRcPrePdu.String(), resultPer.String())
}

func TestE2SmRcPreControlHeaderExcludeOptionalField(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)

	cellID := asn1.BitString{
		Value: []byte{0xab, 0xd4, 0xbc, 0x90, 0x00},
		Len:   36, //uint32
	}

	cgi, err := CreateCellGlobalIDNrCgi(plmnIDBytes, &cellID)
	assert.NilError(t, err)
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader()
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
	newE2SmRcPrePdu.GetControlHeaderFormat1().SetCGI(cgi)

	per, err := encoder.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded ControlHeader: %v", hex.Dump(per))

	resultPer, err := encoder.PerDecodeE2SmRcPreControlHeader(per)
	assert.NilError(t, err)
	t.Logf("PER ControlHeader - decoded \n%v", resultPer)
	assert.Equal(t, newE2SmRcPrePdu.String(), resultPer.String())
}

func TestE2SmRcPreControlHeaderExcludeAllOptionalFields(t *testing.T) {
	newE2SmRcPrePdu, err := CreateE2SmRcPreControlHeader()
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err := encoder.PerEncodeE2SmRcPreControlHeader(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded ControlHeader: %v", hex.Dump(per))

	resultPer, err := encoder.PerDecodeE2SmRcPreControlHeader(per)
	assert.NilError(t, err)
	t.Logf("PER ControlHeader - decoded \n%v", resultPer)
	assert.Equal(t, newE2SmRcPrePdu.String(), resultPer.String())
}
