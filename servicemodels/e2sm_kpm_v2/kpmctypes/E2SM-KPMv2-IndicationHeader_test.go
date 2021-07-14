// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createE2SMKPMIndicationHeader() *e2sm_kpm_v2.E2SmKpmIndicationHeader {

	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	plmnID := []byte{0x37, 0x34, 0x37}
	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	var fileFormatVersion string = "txt"
	var senderName string = "ONF"
	var senderType string = "someType"
	var vendorName string = "onf"

	globalKpmNodeID, _ := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	globalKpmNodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmNodeID.GetGNb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmIndicationHeader(timeStamp, fileFormatVersion, senderName, senderType, vendorName, globalKpmNodeID)

	return newE2SmKpmPdu
}

func Test_xerEncodeE2SmKpmIndicationHeader(t *testing.T) {

	ih := createE2SMKPMIndicationHeader()

	xer, err := XerEncodeE2SmKpmIndicationHeader(ih)
	assert.NilError(t, err)
	//assert.Equal(t, 956, len(xer))
	t.Logf("E2SmKpmIndicationHeader XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmIndicationHeader(t *testing.T) {

	ih := createE2SMKPMIndicationHeader()

	xer, err := XerEncodeE2SmKpmIndicationHeader(ih)
	assert.NilError(t, err)
	//assert.Equal(t, 956, len(xer))
	t.Logf("E2SmKpmIndicationHeader XER\n%s", string(xer))

	result, err := XerDecodeE2SmKpmIndicationHeader(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationHeader XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmIndicationHeader(t *testing.T) {

	ih := createE2SMKPMIndicationHeader()

	per, err := PerEncodeE2SmKpmIndicationHeader(ih)
	assert.NilError(t, err)
	assert.Equal(t, 42, len(per))
	t.Logf("E2SmKpmIndicationHeader PER\n%v", hex.Dump(per))
}

func Test_perDecodeE2SmKpmIndicationHeader(t *testing.T) {

	ih := createE2SMKPMIndicationHeader()

	per, err := PerEncodeE2SmKpmIndicationHeader(ih)
	assert.NilError(t, err)
	assert.Equal(t, 42, len(per))
	t.Logf("E2SmKpmIndicationHeader PER\n%v", hex.Dump(per))
	//t.Logf("E2SmKpmIndicationHeader PER\n%x", per)

	result, err := PerDecodeE2SmKpmIndicationHeader(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationHeader PER - decoded\n%s", result)
}
