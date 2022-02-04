// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/encoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerE2SmKpmIndicationHeader = "00000000  1f 21 22 23 24 18 74 78  74 00 00 03 4f 4e 46 40  |.!\"#$.txt...ONF@|\n" +
	"00000010  73 6f 6d 65 54 79 70 65  06 6f 6e 66 0c 37 34 37  |someType.onf.747|\n" +
	"00000020  00 d4 bc 08 80 30 39 20  1a 85                    |.....09 ..|"

func createE2SmKpmIndicationHeader() (*e2sm_kpm_v2_go.E2SmKpmIndicationHeader, error) {

	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	plmnID := []byte{0x37, 0x34, 0x37}
	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	var fileFormatVersion = "txt"
	var senderName = "ONF"
	var senderType = "someType"
	var vendorName = "onf"

	globalKpmNodeID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	if err != nil {
		return nil, err
	}
	globalKpmNodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmNodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(timeStamp)
	if err != nil {
		return nil, err
	}
	newE2SmKpmPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetSenderType(senderType).SetVendorName(vendorName).SetGlobalKPMnodeID(globalKpmNodeID)

	return newE2SmKpmPdu, nil
}

func Test_perEncodingE2SmKpmIndicationHeader(t *testing.T) {

	ih, err := createE2SmKpmIndicationHeader()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmIndicationHeader(ih)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationHeader PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmKpmIndicationHeader(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-KPM-IndicationHeader PER - decoded\n%v", result)
	assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetFileFormatversion(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetFileFormatversion())
	assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGNbDuId().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGNbDuId().GetValue())
	assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGNbCuUpId().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGNbCuUpId().GetValue())
	assert.DeepEqual(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue())
	assert.DeepEqual(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue())
	assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetKpmNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen())
	assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetSenderName(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetSenderName())
	assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetSenderType(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetSenderType())
	assert.DeepEqual(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetColletStartTime().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetColletStartTime().GetValue())
	assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetVendorName(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetVendorName())
}

func Test_perE2SmKpmIndicationHeaderCompareBytes(t *testing.T) {

	ih, err := createE2SmKpmIndicationHeader()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmIndicationHeader(ih)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationHeader PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerE2SmKpmIndicationHeader)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
