// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

var kpmTestSm servicemodel

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
	assert.Equal(t, 65, len(protoBytes))

	asn1Bytes, err := kpmTestSm.IndicationHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
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
