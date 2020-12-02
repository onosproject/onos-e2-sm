// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	e2smkpmies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var kpmTestSm servicemodel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	var plmnID = "ONF"
	var gNbCuUpID int64 = 0
	var gNbDuID int64 = 0
	var plmnIDnrcgi = "onf"
	var sst = "1"
	var sd = "SD1"
	var fiveQi int32 = 0
	var qCi int32 = 0

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(plmnID, gNbCuUpID, gNbDuID, plmnIDnrcgi, sst, sd, fiveQi, qCi)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmKpmPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationHeader to bytes")
	assert.Equal(t, 68, len(protoBytes))

	asn1Bytes, err := kpmTestSm.IndicationHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 31, len(asn1Bytes))
}

func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	var plmnID = "ONF"
	var cellIdentityValue uint64 = 0
	var cellIdentityLen uint32 = 0
	var ulTotalAvlblProbs int32 = 0
	var dlTotalAvlblProbs int32 = 0
	var fiveQi int32 = 0
	var dlPrbusage int32 = 0
	var ulPrbusage int32 = 0
	var qCi int32 = 0
	var qciDlPrbusage int32 = 0
	var qciUlPrbusage int32 = 0
	var sst = "1"
	var sd = "SD1"
	var gNbCuName = "OpenNetworking"
	var cuCpNumberActvts int32 = 0
	var ranContainer = "ranContainer"
	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationMsg(plmnID, cellIdentityValue, cellIdentityLen, dlTotalAvlblProbs,
		ulTotalAvlblProbs, fiveQi, dlPrbusage, ulPrbusage, qCi, qciDlPrbusage, qciUlPrbusage, sst,
		sd, gNbCuName, cuCpNumberActvts, ranContainer)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmKpmPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationMessage to bytes")
	assert.Equal(t, 44, len(protoBytes)) //with ODU it will raise on 110

	asn1Bytes, err := kpmTestSm.IndicationMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
}

func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-KPM-Indication-Header.xml
	indicationHeaderAsn1Bytes := []byte{0x3F, 0x08, 0x37, 0x34, 0x37, 0x38, 0xB5, 0xC6, 0x77, 0x88, 0x02, 0x37, 0x34, 0x37, 0x22, 0x5B,
		0xD6, 0x00, 0x70, 0x37, 0x34, 0x37, 0x98, 0x80, 0x31, 0x30, 0x30, 0x09, 0x09}

	protoBytes, err := kpmTestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 74, len(protoBytes))
	testIH := e2smkpmies.E2SmKpmIndicationHeader{}
	err = proto.Unmarshal(protoBytes, &testIH)
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte{0x37, 0x34, 0x37}, testIH.GetIndicationHeaderFormat1().GetPLmnIdentity().GetValue())
}

func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {
	// This message is taken from Shad
	indicationMessageAsn1 := []byte{0x40, 0x00, 0x00, 0x4B, 0x01, 0x00}

	protoBytes, err := kpmTestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 10, len(protoBytes))
	testIM := e2smkpmies.E2SmKpmIndicationMessage{}
	err = proto.Unmarshal(protoBytes, &testIM)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(testIM.GetIndicationMessageFormat1().GetPmContainers()))
	pm1 := testIM.GetIndicationMessageFormat1().GetPmContainers()[0]
	assert.Equal(t, 0, int(pm1.GetPerformanceContainer().GetOCuCp().GetCuCpResourceStatus().GetNumberOfActiveUes()))
}
