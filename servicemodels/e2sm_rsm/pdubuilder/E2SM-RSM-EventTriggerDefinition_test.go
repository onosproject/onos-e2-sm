// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_E2SmRsmEventTriggerDefinition(t *testing.T) {

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

	drbIDfiveG, err := CreateDrbIDfiveG(37, 62, flowMap)
	assert.NilError(t, err)
	bearerID2 := CreateBearerIDdrb(drbIDfiveG)

	bearerList := make([]*e2sm_rsm_ies.BearerId, 0)
	bearerList = append(bearerList, bearerID1)
	bearerList = append(bearerList, bearerID2)

	etd, err := CreateE2SmRsmEventTriggerDefinitionFormat1(CreateRsmTriggerTypeUeAttach(), ueIDlist, CreateUeIDtypeAmfUeNgapID(), bearerList)
	assert.NilError(t, err)
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", etd)

	//ToDo - embed APER validation
}