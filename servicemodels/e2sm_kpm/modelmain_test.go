// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	"gotest.tools/assert"
	"testing"
	"fmt"
)

var kpmTestSm servicemodel
//Taken as an output from TestServicemodel_IndicationHeaderProtoToASN1 function
//const indicationHeaderAsn1Bytes = "6312797870021218880480481111101022391712121880797870152128836849000000" +
//	"000000000000000000000000000000000000000000000000000000000000" +
//	"00000000000000000000000000000000000000000000000000000000000000000000000000000" +
//	"0000000000000000000000000000000000000000000000000000000000000000000000000000"


func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	var plmnID = "ONF"
	var gNbCuUpId int64 = 0
	var gNbDuId int64 = 0
	var plmnIDnrcgi = "onf"
	var sst = "1"
	var sd = "SD1"
	var fiveQi int32 = 0
	var qCi int32 = 0

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationHeader(plmnID, gNbCuUpId, gNbDuId, plmnIDnrcgi, sst, sd, fiveQi, qCi)
	assert.NilError(t, err, "error creating E2SmPDU")

	err = newE2SmKpmPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmKpmPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmKpmIndicationHeader to bytes")
	assert.Equal(t, 68, len(protoBytes))

	asn1Bytes, err := kpmTestSm.IndicationHeaderProtoToASN1(protoBytes)
	fmt.Printf("ASN1 bytes: %x\n", asn1Bytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, asn1Bytes != nil)
	assert.Equal(t, 248, len(asn1Bytes))
}

//TODO: Fix it
func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {

	//TODO: Replace with a correct IndicationHeader. If it won't work, dig inside the decoding part of ASN structures
	//indicationHeaderAsn1Bytes := []byte{0x68, 0x65, 0x61, 0x64, 0x65, 0x72}

	// This message is taken from the function above (asn1Bytes variable)
	// For some reason, during decoding, it detects GnbID of type "GlobalKPMnode_ID_PR_eNB",
	// which is not implemented (yet). Two reasons - either parsing of the message over string doesn't work,
	// or something wrong in a decoding process..
	indicationHeaderAsn1Bytes := "3f0c4f4e4600d4bc08003000306f6e66efabd4bc004f4e469880534431"

	protoBytes, err := kpmTestSm.IndicationHeaderASN1toProto([]byte(indicationHeaderAsn1Bytes))
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 67, len(protoBytes))
}

func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	var plmnID = "ONF"
	var cellIdentityValue uint64 = 0x9bcd4abef
	var cellIdentityLen uint32 = 36
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
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, asn1Bytes != nil)
}

func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {
	indicationMessageAsn1 := []byte{0x40, 0x00, 0x00, 0x4B, 0x01, 0x00}

	protoBytes, err := kpmTestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	assert.Equal(t, 12, len(protoBytes))
}