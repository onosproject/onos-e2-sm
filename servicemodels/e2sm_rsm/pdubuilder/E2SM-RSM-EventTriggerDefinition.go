// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"fmt"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies"
)

func CreateE2SmRsmEventTriggerDefinitionFormat1(tt e2sm_rsm_ies.RsmTriggerType, ueIDlist []*e2sm_rsm_ies.UeIdentity,
	pt e2sm_rsm_ies.UeIdType, bearerList []*e2sm_rsm_ies.BearerId) (*e2sm_rsm_ies.E2SmRsmEventTriggerDefinition, error) {

	if len(ueIDlist) == 0 {
		return nil, fmt.Errorf("UE ID list should contain at least 1 element. All elements of the list should be different UE identifiers")
	}

	if len(bearerList) < 1 || len(bearerList) > 32 {
		return nil, fmt.Errorf("BearerID list should have 1 to 32 items")
	}

	return &e2sm_rsm_ies.E2SmRsmEventTriggerDefinition{
		EventDefinitionFormats: &e2sm_rsm_ies.EventDefinitionFormats{
			E2SmRsmEventDefinition: &e2sm_rsm_ies.EventDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: &e2sm_rsm_ies.E2SmRsmEventTriggerDefinitionFormat1{
					TriggerType:       tt,
					UeIdlist:          ueIDlist,
					PrefferedUeIdtype: pt,
					BearerId:          bearerList,
				},
			},
		},
	}, nil
}

func CreateRsmTriggerTypeUeAttach() e2sm_rsm_ies.RsmTriggerType {
	return e2sm_rsm_ies.RsmTriggerType_RSM_TRIGGER_TYPE_UE_ATTACH
}

func CreateRsmTriggerTypeUeDetach() e2sm_rsm_ies.RsmTriggerType {
	return e2sm_rsm_ies.RsmTriggerType_RSM_TRIGGER_TYPE_UE_DETACH
}

func CreateRsmTriggerTypeHandInUeAttach() e2sm_rsm_ies.RsmTriggerType {
	return e2sm_rsm_ies.RsmTriggerType_RSM_TRIGGER_TYPE_HAND_IN_UE_ATTACH
}

func CreateRsmTriggerTypeHandOutUeAttach() e2sm_rsm_ies.RsmTriggerType {
	return e2sm_rsm_ies.RsmTriggerType_RSM_TRIGGER_TYPE_HAND_OUT_UE_ATTACH
}

func CreateUeIDtypeCuUeF1ApID() e2sm_rsm_ies.UeIdType {
	return e2sm_rsm_ies.UeIdType_UE_ID_TYPE_CU_UE_F1_AP_ID
}

func CreateUeIDtypeDuUeF1ApID() e2sm_rsm_ies.UeIdType {
	return e2sm_rsm_ies.UeIdType_UE_ID_TYPE_DU_UE_F1_AP_ID
}

func CreateUeIDtypeRanUeNgapID() e2sm_rsm_ies.UeIdType {
	return e2sm_rsm_ies.UeIdType_UE_ID_TYPE_RAN_UE_NGAP_ID
}

func CreateUeIDtypeAmfUeNgapID() e2sm_rsm_ies.UeIdType {
	return e2sm_rsm_ies.UeIdType_UE_ID_TYPE_AMF_UE_NGAP_ID
}

func CreateUeIDtypeEnbUeS1ApID() e2sm_rsm_ies.UeIdType {
	return e2sm_rsm_ies.UeIdType_UE_ID_TYPE_ENB_UE_S1_AP_ID
}

func CreateBearerIDdrb(drbID *e2sm_rsm_ies.DrbId) *e2sm_rsm_ies.BearerId {

	return &e2sm_rsm_ies.BearerId{
		BearerId: &e2sm_rsm_ies.BearerId_DrbId{
			DrbId: drbID,
		},
	}
}

func CreateDrbIDfourG(val int32, qci int32) (*e2sm_rsm_ies.DrbId, error) {

	if qci < 0 || qci > 255 {
		return nil, fmt.Errorf("QCI value should be in range 0 to 255")
	}

	return &e2sm_rsm_ies.DrbId{
		DrbId: &e2sm_rsm_ies.DrbId_FourGdrbId{
			FourGdrbId: &e2sm_rsm_ies.FourGDrbId{
				Value: val,
				Qci: &e2sm_v2_ies.Qci{
					Value: qci,
				},
			},
		},
	}, nil
}

func CreateDrbIDfiveG(val int32, qfi int32, flowMap []*e2sm_rsm_ies.QoSflowLevelParameters) (*e2sm_rsm_ies.DrbId, error) {

	if qfi < 0 || qfi > 63 {
		return nil, fmt.Errorf("QCI value should be in range 0 to 255")
	}

	if len(flowMap) < 1 || len(flowMap) > 64 {
		return nil, fmt.Errorf("FlowsMapToDrb list should have 1 to 64 items")
	}

	return &e2sm_rsm_ies.DrbId{
		DrbId: &e2sm_rsm_ies.DrbId_FiveGdrbId{
			FiveGdrbId: &e2sm_rsm_ies.FiveGDrbId{
				Value: val,
				Qfi: &e2sm_rsm_ies.Qfi{
					Value: qfi,
				},
				FlowsMapToDrb: flowMap,
			},
		},
	}, nil
}

func CreateQosFlowLevelParametersDynamic(prlvl int32, pDelay int32, per int32) *e2sm_rsm_ies.QoSflowLevelParameters {

	return &e2sm_rsm_ies.QoSflowLevelParameters{
		QoSflowLevelParameters: &e2sm_rsm_ies.QoSflowLevelParameters_DynamicFiveQi{
			DynamicFiveQi: &e2sm_rsm_ies.DynamicFiveQi{
				PriorityLevel: prlvl,
				PacketDelayBudget: pDelay,
				PacketErrorRate: per,
			},
		},
	}
}

func CreateQosFlowLevelParametersNonDynamic(fiveQI int32) *e2sm_rsm_ies.QoSflowLevelParameters {

	return &e2sm_rsm_ies.QoSflowLevelParameters{
		QoSflowLevelParameters: &e2sm_rsm_ies.QoSflowLevelParameters_NonDynamicFiveQi{
			NonDynamicFiveQi: &e2sm_rsm_ies.NonDynamicFiveQi{
				FiveQi: &e2sm_v2_ies.FiveQi{
					Value: fiveQI,
				},
			},
		},
	}
}
