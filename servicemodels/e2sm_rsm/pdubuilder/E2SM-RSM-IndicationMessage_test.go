// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/encoder"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"gotest.tools/assert"
	"testing"
)

func TestCreateE2SmRsmIndicationMessageFormat1(t *testing.T) {

	ueID, err := CreateUeIDAmfUeNgapID(1)
	assert.NilError(t, err)

	ulSm := make([]*e2sm_rsm_ies.SliceMetrics, 0)
	ulM1, err := CreateSliceMetrics(100, 100, 100, 15)
	assert.NilError(t, err)
	ulSm = append(ulSm, ulM1)
	ulM2, err := CreateSliceMetrics(10, 1, 50, 7)
	assert.NilError(t, err)
	ulSm = append(ulSm, ulM2)
	ulM3, err := CreateSliceMetrics(47, 16, 12, 1)
	assert.NilError(t, err)
	ulSm = append(ulSm, ulM3)

	dlSm := make([]*e2sm_rsm_ies.SliceMetrics, 0)
	dlM1, err := CreateSliceMetrics(91, 31, 54, 11)
	assert.NilError(t, err)
	dlSm = append(dlSm, dlM1)
	dlM2, err := CreateSliceMetrics(84, 37, 41, 6)
	assert.NilError(t, err)
	dlSm = append(dlSm, dlM2)

	im, err := CreateE2SmRsmIndicationMessageFormat1(ueID, 27, 14, CreateEmmCaseAttach(),
		ulSm, dlSm)
	assert.NilError(t, err)
	assert.Assert(t, im != nil)
	t.Logf("Created E2SM-RSM-IndicationMessage is \n%v", im)

	// APER encoding validation
	per, err := encoder.PerEncodeE2SmRsmIndicationMessage(im)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-IndicationMessage PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmIndicationMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-IndicationMessage PER - decoded\n%v", result)
	assert.DeepEqual(t, im.String(), result.String())
}

func TestCreateE2SmRsmIndicationMessageFormat2(t *testing.T) {

	ueIDlist := make([]*e2sm_rsm_ies.UeIdentity, 0)
	ngap, err := CreateUeIDAmfUeNgapID(21)
	assert.NilError(t, err)
	ueIDlist = append(ueIDlist, ngap)
	f1ap, err := CreateUeIDCuUeF1ApID(43)
	assert.NilError(t, err)
	ueIDlist = append(ueIDlist, f1ap)
	s1ap, err := CreateUeIDEnbUeS1ApID(1)
	assert.NilError(t, err)
	ueIDlist = append(ueIDlist, s1ap)

	drbIDfourG, err := CreateDrbIDfourG(12, 127)
	assert.NilError(t, err)
	bearerID1, err := CreateBearerIDdrb(drbIDfourG)
	assert.NilError(t, err)

	flowMap := make([]*e2sm_rsm_ies.QoSflowLevelParameters, 0)
	dqos, err := CreateQosFlowLevelParametersDynamic(10, 62, 54)
	assert.NilError(t, err)
	flowMap = append(flowMap, dqos)
	ndqos, err := CreateQosFlowLevelParametersNonDynamic(11)
	assert.NilError(t, err)
	flowMap = append(flowMap, ndqos)

	drbIDfiveG, err := CreateDrbIDfiveG(32, 62, flowMap)
	assert.NilError(t, err)
	bearerID2, err := CreateBearerIDdrb(drbIDfiveG)
	assert.NilError(t, err)

	bearerList := make([]*e2sm_rsm_ies.BearerId, 0)
	bearerList = append(bearerList, bearerID1)
	bearerList = append(bearerList, bearerID2)

	im, err := CreateE2SmRsmIndicationMessageFormat2(CreateRsmEmmTriggerTypeUeAttach(), ueIDlist, CreateUeIDtypeAmfUeNgapID(), bearerList)
	assert.NilError(t, err)
	assert.Assert(t, im != nil)
	t.Logf("Created E2SM-RSM-IndicationMessage is \n%v", im)

	// APER encoding validation
	per, err := encoder.PerEncodeE2SmRsmIndicationMessage(im)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-IndicationMessage PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmIndicationMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-IndicationMessage PER - decoded\n%v", result)
	assert.DeepEqual(t, im.String(), result.String())
}

func TestCreateE2SmRsmIndicationMessageErrors(t *testing.T) {

	ueID, err := CreateUeIDAmfUeNgapID(1)
	assert.NilError(t, err)

	_, err = CreateSliceMetrics(101, 100, 100, 15)
	assert.ErrorContains(t, err, "PRB utilization value should be within range 0 to 100")

	_, err = CreateSliceMetrics(99, 100, -10, 15)
	assert.ErrorContains(t, err, "BLER value should be within range 0 to 100")

	_, err = CreateSliceMetrics(99, 100, 11, 17)
	assert.ErrorContains(t, err, "CQI value should be within range 0 to 15")

	_, err = CreateE2SmRsmIndicationMessageFormat1(ueID, 27, 14, CreateEmmCaseAttach(),
		nil, nil)
	assert.ErrorContains(t, err, "Slicing Metrics list should have at least 1 item")
}

func TestWoojong(t *testing.T) {

	//bytes := []byte{0x50, 0x81, 0xe6, 0xf5, 0x14, 0xe6, 0xf5, 0x48, 0xb7, 0xbc, 0x2f, 0x10, 0x01, 0x00, 0x09} //UeAttach case
	bytes := []byte{0x52, 0x12, 0xb7, 0xbc, 0x2f, 0x40, 0x01, 0x00, 0x09} // UeDetach
	t.Logf("PER bytes are \n%v", hex.Dump(bytes))

	res, err := encoder.PerDecodeE2SmRsmIndicationMessage(bytes)
	assert.NilError(t, err)
	t.Logf("Decoded message is \n%v", res)
}

func TestKushal(t *testing.T) {

	ueIDlist := make([]*e2sm_rsm_ies.UeIdentity, 0)
	cuf1ap, err := CreateUeIDCuUeF1ApID(59125)
	assert.NilError(t, err)
	ueIDlist = append(ueIDlist, cuf1ap)
	duf1ap, err := CreateUeIDDuUeF1ApID(59125)
	assert.NilError(t, err)
	ueIDlist = append(ueIDlist, duf1ap)
	s1ap, err := CreateUeIDEnbUeS1ApID(12041263)
	assert.NilError(t, err)
	ueIDlist = append(ueIDlist, s1ap)

	drbIDfourG, err := CreateDrbIDfourG(5, 9)
	assert.NilError(t, err)
	bearerID1, err := CreateBearerIDdrb(drbIDfourG)
	assert.NilError(t, err)

	bearerList := make([]*e2sm_rsm_ies.BearerId, 0)
	bearerList = append(bearerList, bearerID1)

	im, err := CreateE2SmRsmIndicationMessageFormat2(CreateRsmEmmTriggerTypeUeAttach(), ueIDlist, CreateUeIDtypeDuUeF1ApID(), bearerList)
	assert.NilError(t, err)
	assert.Assert(t, im != nil)
	t.Logf("Created E2SM-RSM-IndicationMessage is \n%v", im)

	// APER encoding validation
	per, err := encoder.PerEncodeE2SmRsmIndicationMessage(im)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-IndicationMessage PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmIndicationMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-IndicationMessage PER - decoded\n%v", result)
	assert.DeepEqual(t, im.String(), result.String())
}
