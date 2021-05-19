// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"encoding/hex"
	kpmctypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/kpmctypes"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var radisysBytes = []byte{
	0x70, 0x18, 0x4f, 0x52, 0x41, 0x4e, 0x2d, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x4b, 0x50, 0x4d, 0x00, 0x00, 0x18, 0x31, 0x2e,
	0x33, 0x2e, 0x36, 0x2e, 0x31, 0x2e, 0x34, 0x2e, 0x31, 0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e, 0x31, 0x2e, 0x32, 0x2e,
	0x32, 0x2e, 0x32, 0x05, 0x00, 0x4b, 0x50, 0x4d, 0x20, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x00, 0x00, 0x40, 0x00,
	0x13, 0xf1, 0x84, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x30, 0x00, 0x13, 0xf1, 0x84, 0x00, 0x00,
	0x00, 0x00, 0x10, 0x00, 0x01, 0x07, 0x00, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x20, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x00, 0x01, 0x00, 0x03, 0x09, 0x00, 0x45, 0x32, 0x20, 0x4e, 0x6f, 0x64, 0x65, 0x20, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x00, 0x01, 0x00, 0x07, 0x42, 0x60, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x53, 0x75, 0x6d, 0x00, 0x00, 0x00, 0x42, 0x80, 0x52, 0x52,
	0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x53, 0x75, 0x63, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x00,
	0x00, 0x01, 0x42, 0xa0, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41,
	0x74, 0x74, 0x2e, 0x53, 0x75, 0x6d, 0x00, 0x00, 0x02, 0x43, 0xc0, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52,
	0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x61,
	0x69, 0x6c, 0x00, 0x00, 0x03, 0x43, 0x00, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x45, 0x73, 0x74,
	0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x48, 0x4f, 0x46, 0x61, 0x69, 0x6c, 0x00, 0x00, 0x04, 0x42, 0xe0, 0x52, 0x52, 0x43,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x00, 0x00, 0x05, 0x41, 0x60, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x65, 0x61, 0x6e, 0x00, 0x00,
	0x06, 0x41, 0x40, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78, 0x00, 0x00, 0x07, 0x00, 0x01, 0x00,
	0x01}

var rfdBytesGnb1 = "70184f52414e2d4532534d2d4b504d000018312e332e362e312e342e312e35333134382e312e322e322e3205004b504d206" +
	"d6f6e69746f720000400013f18450000000010000000001300013f184000000001000010700506572696f646963207265706f7274000100030" +
	"9004532204e6f6465204d6561737572656d656e740001000742605252432e436f6e6e45737461624174742e53756d00000042805252432e436f" +
	"6e6e4573746162537563632e53756d00000142a05252432e436f6e6e526545737461624174742e53756d00000243c05252432e436f6e6e52654" +
	"5737461624174742e7265636f6e6669674661696c00000343005252432e436f6e6e526545737461624174742e484f4661696c00000442e05252" +
	"432e436f6e6e526545737461624174742e4f7468657200000541605252432e436f6e6e4d65616e00000641405252432e436f6e6e4d617800000" +
	"700010001"

var adBytes = "000300000130000710000010000110000210000310000410000510000610000700000000"

var etdBytes = "08752F"

var ihBytes = "006092A227"

var imBytes = "0E00000000013040752F000710000010000110000210000310000410000510000610000700070001000100010001000100000001000000010000000100000001000100010001"

func Test_DecodeRanFunctionDescriptionBytesRadisys(t *testing.T) {
	t.Logf("Radisys Bytes for RanFunctionDescription (Gnb=1) in HEX format are \n%v\n", hex.Dump(radisysBytes))
	rfd, err := kpmctypes.PerDecodeE2SmKpmRanfunctionDescription(rSysBytes)
	assert.NilError(t, err)
	err = rfd.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", rfd)
	assert.Equal(t, "ORAN-E2SM-KPM", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "KPM monitor", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "1.3.6.1.4.1.53148.1.2.2.2", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())

	// Printing message in XER form
	xer, err := kpmctypes.XerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("RanFunctionDescription XER is \n%s", xer)
}

func Test_DecodeRanFunctionDescriptionBytesRadisysGnb1(t *testing.T) {
	rfdPer, err := hexlib.Asn1BytesToByte(rfdBytesGnb1)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for RanFunctionDescription in HEX format are \n%v\n", hex.Dump(rfdPer))
	rfd, err := kpmctypes.PerDecodeE2SmKpmRanfunctionDescription(rfdPer)
	assert.NilError(t, err)
	err = rfd.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", rfd)
	assert.Equal(t, "ORAN-E2SM-KPM", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "KPM monitor", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "1.3.6.1.4.1.53148.1.2.2.2", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())

	// Printing message in XER form
	xer, err := kpmctypes.XerEncodeE2SmKpmRanfunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("RanFunctionDescription XER is \n%s", xer)
}

func Test_DecodeActionDefinitionBytesRadisys(t *testing.T) {
	adPer, err := hexlib.Asn1BytesToByte(adBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for ActionDefinition in HEX format are \n%v\n", hex.Dump(adPer))
	ad, err := kpmctypes.PerDecodeE2SmKpmActionDefinition(adPer)
	assert.NilError(t, err)
	err = ad.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", ad)
	assert.Equal(t, int32(3), ad.GetRicStyleType().GetValue())
	assert.Equal(t, "0", ad.GetActionDefinitionFormat1().GetCellObjId().GetValue())
	assert.Equal(t, int32(5), ad.GetActionDefinitionFormat1().GetMeasInfoList().GetValue()[4].GetMeasType().GetMeasId().GetValue())
	assert.Equal(t, uint32(1), ad.GetActionDefinitionFormat1().GetGranulPeriod().GetValue())
	assert.Equal(t, int64(1), ad.GetActionDefinitionFormat1().GetSubscriptId().GetValue())

	// Printing message in XER form
	xer, err := kpmctypes.XerEncodeE2SmKpmActionDefinition(ad)
	assert.NilError(t, err)
	t.Logf("ActionDefinition XER message is \n%s", xer)
}

func Test_DecodeEventTriggerDefinitionBytesRadisys(t *testing.T) {
	etdPer, err := hexlib.Asn1BytesToByte(etdBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for EventTriggerDefinition in HEX format are \n%v\n", hex.Dump(etdPer))
	etd, err := kpmctypes.PerDecodeE2SmKpmEventTriggerDefinition(etdPer)
	assert.NilError(t, err)
	err = etd.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", etd)
	assert.Equal(t, uint32(30000), etd.GetEventDefinitionFormat1().GetReportingPeriod())

	// Printing message in XER form
	//xer, err := kpmctypes.XerEncodeE2SmKpmEventTriggerDefinition(etd)
	//assert.NilError(t, err)
	//t.Logf("EventTriggerDefinition XER is \n%s", xer)
}

func Test_DecodeIndicationHeaderBytesRadisys(t *testing.T) {
	ihPer, err := hexlib.Asn1BytesToByte(ihBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for IndicationHeader in HEX format are \n%v\n", hex.Dump(ihPer))
	ih, err := kpmctypes.PerDecodeE2SmKpmIndicationHeader(ihPer)
	assert.NilError(t, err)
	err = ih.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", ih)
	assert.DeepEqual(t, []byte{0x60, 0x92, 0xA2, 0x27}, ih.GetIndicationHeaderFormat1().GetColletStartTime().GetValue())

	// Printing message in XER form
	xer, err := kpmctypes.XerEncodeE2SmKpmIndicationHeader(ih)
	assert.NilError(t, err)
	t.Logf("IndicationHeader XER is \n%s", xer)
}

func Test_DecodeIndicationMessageBytesRadisys(t *testing.T) {
	imPer, err := hexlib.Asn1BytesToByte(imBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for IndicationMessage in HEX format are \n%v\n", hex.Dump(imPer))
	im, err := kpmctypes.PerDecodeE2SmKpmIndicationMessage(imPer)
	assert.NilError(t, err)
	err = im.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", im)
	assert.Equal(t, int64(1), im.GetIndicationMessageFormat1().GetSubscriptId().GetValue())
	assert.Equal(t, "0", im.GetIndicationMessageFormat1().GetCellObjId().GetValue())
	assert.Equal(t, int32(1), im.GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[0].GetMeasType().GetMeasId().GetValue())
	assert.Equal(t, int32(6), im.GetIndicationMessageFormat1().GetMeasInfoList().GetValue()[5].GetMeasType().GetMeasId().GetValue())
	assert.Equal(t, e2sm_kpm_v2.IncompleteFlag_INCOMPLETE_FLAG_TRUE, im.GetIndicationMessageFormat1().GetMeasData().GetValue()[0].GetIncompleteFlag())
	assert.Equal(t, e2sm_kpm_v2.IncompleteFlag_INCOMPLETE_FLAG_TRUE, im.GetIndicationMessageFormat1().GetMeasData().GetValue()[6].GetIncompleteFlag())

	// Printing message in XER form
	xer, err := kpmctypes.XerEncodeE2SmKpmIndicationMessage(im)
	assert.NilError(t, err)
	t.Logf("IndicationMessage XER is \n%s", xer)
}
