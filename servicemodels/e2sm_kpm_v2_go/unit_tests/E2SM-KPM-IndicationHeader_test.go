// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerE2SmKpmIndicationHeader = "00000000  1f 21 22 23 24 18 74 78  74 00 00 03 4f 4e 46 40  |.!\"#$.txt...ONF@|" +
	"00000010  73 6f 6d 65 54 79 70 65  06 6f 6e 66 0c 37 34 37  |someType.onf.747|" +
	"00000020  00 d4 bc 08 80 30 39 20  1a 85                    |.....09 ..|"

func createE2SmKpmIndicationHeader() *e2sm_kpm_v2_go.E2SmKpmIndicationHeader {

	bs := asn1.BitString{
		Value:[]byte{0xd4, 0xbc, 0x08},
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
	globalKpmNodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmNodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmIndicationHeader(timeStamp, &fileFormatVersion, &senderName, &senderType, &vendorName, globalKpmNodeID)

	return newE2SmKpmPdu
}

func Test_perEncodingE2SmKpmIndicationHeader(t *testing.T) {

	ih := createE2SmKpmIndicationHeader()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(ih, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationHeader PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmIndicationHeader{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-IndicationHeader PER - decoded\n%v", result)
}

func Test_perE2SmKpmIndicationHeaderCompareBytes(t *testing.T) {

	ih := createE2SmKpmIndicationHeader()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(ih, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationHeader PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerE2SmKpmIndicationHeader)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
