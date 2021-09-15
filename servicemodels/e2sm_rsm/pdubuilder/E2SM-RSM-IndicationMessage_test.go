// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/encoder"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"gotest.tools/assert"
	"testing"
)

func TestCreateE2SmRsmIndicationMessageFormat1(t *testing.T) {

	ueID := CreateUeIDAmfUeNgapID(1)

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
	ueIDlist = append(ueIDlist, CreateUeIDAmfUeNgapID(21))
	ueIDlist = append(ueIDlist, CreateUeIDCuUeF1ApID(43))
	ueIDlist = append(ueIDlist, CreateUeIDEnbUeS1ApID(1))

	drbIDfourG, err := CreateDrbIDfourG(12, 127)
	assert.NilError(t, err)
	bearerID1 := CreateBearerIDdrb(drbIDfourG)

	flowMap := make([]*e2sm_rsm_ies.QoSflowLevelParameters, 0)
	flowMap = append(flowMap, CreateQosFlowLevelParametersDynamic(10, 62, 54))
	flowMap = append(flowMap, CreateQosFlowLevelParametersNonDynamic(11))

	drbIDfiveG, err := CreateDrbIDfiveG(32, 62, flowMap)
	assert.NilError(t, err)
	bearerID2 := CreateBearerIDdrb(drbIDfiveG)

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

	ueID := CreateUeIDAmfUeNgapID(1)

	_, err := CreateSliceMetrics(101, 100, 100, 15)
	assert.ErrorContains(t, err, "PRB utilization value should be within range 0 to 100")

	_, err = CreateSliceMetrics(99, 100, -10, 15)
	assert.ErrorContains(t, err, "BLER value should be within range 0 to 100")

	_, err = CreateSliceMetrics(99, 100, 11, 17)
	assert.ErrorContains(t, err, "CQI value should be within range 0 to 15")

	_, err = CreateE2SmRsmIndicationMessageFormat1(ueID, 27, 14, CreateEmmCaseAttach(),
		nil, nil)
	assert.ErrorContains(t, err, "Slicing Metrics list should have at least 1 item")
}
