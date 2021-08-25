// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/encoder"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"gotest.tools/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log := logging.GetLogger("asn1")
	log.SetLevel(logging.DebugLevel)
	os.Exit(m.Run())
}

func Test_E2SmRsmControlMessageSliceCreate(t *testing.T) {

	parameters := CreateSliceParameters(CreateSchedulerTypeQosBased())

	config := CreateSliceConfig(1, parameters).SetSliceDescription("IoT")

	cm := CreateE2SmRsmControlMessageSliceCreate(config)
	t.Logf("Created E2SM-RSM-ControlMessage (Create Slice) is \n%v", cm)

	//ToDo - embed APER validation
	per, err := encoder.PerEncodeE2SmRsmControlMessage(cm)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-ControlMessage (Create Slice) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmControlMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-ControlMessage (Create Slice) PER - decoded\n%v", result)
	assert.DeepEqual(t, cm.String(), result.String())
}

func Test_E2SmRsmControlMessageSliceUpdate(t *testing.T) {

	parameters := CreateSliceParameters(CreateSchedulerTypeQosBased()).
		SetWeight(21).SetQoSLevel(129).SetScheduleInfo(CreateScheduleConfigEmpty().
			SetCarrierAggregationLevelCap(CreateCarrierAggregationLevelCapThree()).
			SetUlPowerControl(CreateUlPowerControlEmpty().
				SetPUSCHtargetSNR(117).SetPUCCHtargetSNR(84)).
			SetFeatures(CreateFeatureStatusEnable()).SetLinkAdaptation(CreateLinkAdaptationEmpty().
				SetCqiCap(11).SetRiCap(CreateRiCapTwo()).SetAggregationLevelCap(CreateAggregationLevelCapEight()).
				SetTargetBlerDL(91).SetTargetBlerUL(89).SetMaxMcs(28).
				SetMinMcs(0).SetTransmissionMode(CreateTransmissionModeThree()).
				SetHarqRextCap(CreateHarqRextCapEmpty().
					SetUL(21).SetDL(58))))

	config := CreateSliceConfig(1, parameters).SetSliceDescription("Automotive")

	cm := CreateE2SmRsmControlMessageSliceUpdate(config)
	t.Logf("Created E2SM-RSM-ControlMessage (Update Slice) is \n%v", cm)

	//ToDo - embed APER validation
	per, err := encoder.PerEncodeE2SmRsmControlMessage(cm)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-ControlMessage (Update Slice) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmControlMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-ControlMessage (Update Slice) PER - decoded\n%v", result)
	assert.DeepEqual(t, cm.String(), result.String())
}

func Test_E2SmRsmControlMessageSliceDelete(t *testing.T) {

	cm := CreateE2SmRsmControlMessageSliceDelete(3)
	t.Logf("Created E2SM-RSM-ControlMessage (Delete Slice) is \n%v", cm)

	//ToDo - embed APER validation
	per, err := encoder.PerEncodeE2SmRsmControlMessage(cm)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-ControlMessage (Delete Slice) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmControlMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-ControlMessage (Delete Slice) PER - decoded\n%v", result)
	assert.DeepEqual(t, cm.String(), result.String())
}

func Test_E2SmRsmControlMessageSliceAssociate(t *testing.T) {

	ueID := CreateUeIDRanUeNgapID(7)

	ueIDlist := make([]*e2sm_rsm_ies.UeIdentity, 0)
	ueIDlist = append(ueIDlist, CreateUeIDAmfUeNgapID(21))
	ueIDlist = append(ueIDlist, CreateUeIDCuUeF1ApID(43))
	ueIDlist = append(ueIDlist, CreateUeIDRanUeNgapID(7))
	ueIDlist = append(ueIDlist, CreateUeIDDuUeF1ApID(59))
	ueIDlist = append(ueIDlist, CreateUeIDEnbUeS1ApID(1))

	drbIDfourG, err := CreateDrbIDfourG(12, 127)
	assert.NilError(t, err)
	bearerID1 := CreateBearerIDdrb(drbIDfourG)

	flowMap := make([]*e2sm_rsm_ies.QoSflowLevelParameters, 0)
	flowMap = append(flowMap, CreateQosFlowLevelParametersDynamic(10, 62, 54))
	flowMap = append(flowMap, CreateQosFlowLevelParametersNonDynamic(11))

	drbIDfiveG, err := CreateDrbIDfiveG(37, 62, flowMap)
	assert.NilError(t, err)
	bearerID2 := CreateBearerIDdrb(drbIDfiveG)

	bearerList := make([]*e2sm_rsm_ies.BearerId, 0)
	bearerList = append(bearerList, bearerID1)
	bearerList = append(bearerList, bearerID2)

	config, err := CreateSliceAssociate(ueID, bearerList, 7)
	assert.NilError(t, err)
	config.SetUplinkSliceID(19)

	cm := CreateE2SmRsmControlMessageSliceAssociate(config)
	t.Logf("Created E2SM-RSM-ControlMessage (Associate Slice) is \n%v", cm)

	//ToDo - embed APER validation
	per, err := encoder.PerEncodeE2SmRsmControlMessage(cm)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-ControlMessage (Associate Slice) PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmControlMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-ControlMessage (Associate Slice) PER - decoded\n%v", result)
	assert.DeepEqual(t, cm.String(), result.String())
}
