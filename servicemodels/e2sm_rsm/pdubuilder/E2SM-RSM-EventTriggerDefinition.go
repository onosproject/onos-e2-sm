// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrsm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	e2smv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRsmEventTriggerDefinitionFormat1(tt e2smrsm.RsmRicindicationTriggerType) (*e2smrsm.E2SmRsmEventTriggerDefinition, error) {

	etd := &e2smrsm.E2SmRsmEventTriggerDefinition{
		EventDefinitionFormats: &e2smrsm.EventDefinitionFormats{
			E2SmRsmEventDefinition: &e2smrsm.EventDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: &e2smrsm.E2SmRsmEventTriggerDefinitionFormat1{
					TriggerType: tt,
				},
			},
		},
	}

	if err := etd.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmEventTriggerDefinitionFormat1(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return etd, nil
}

func CreateRsmRicindicationTriggerTypePeriodicMetrics() e2smrsm.RsmRicindicationTriggerType {
	return e2smrsm.RsmRicindicationTriggerType_RSM_RICINDICATION_TRIGGER_TYPE_PERIODIC_METRICS
}

func CreateRsmRicindicationTriggerTypeUponEmmEvent() e2smrsm.RsmRicindicationTriggerType {
	return e2smrsm.RsmRicindicationTriggerType_RSM_RICINDICATION_TRIGGER_TYPE_UPON_EMM_EVENT
}

func CreateRsmEmmTriggerTypeUeAttach() e2smrsm.RsmEmmTriggerType {
	return e2smrsm.RsmEmmTriggerType_RSM_EMM_TRIGGER_TYPE_UE_ATTACH
}

func CreateRsmEmmTriggerTypeUeDetach() e2smrsm.RsmEmmTriggerType {
	return e2smrsm.RsmEmmTriggerType_RSM_EMM_TRIGGER_TYPE_UE_DETACH
}

func CreateRsmEmmTriggerTypeHandInUeAttach() e2smrsm.RsmEmmTriggerType {
	return e2smrsm.RsmEmmTriggerType_RSM_EMM_TRIGGER_TYPE_HAND_IN_UE_ATTACH
}

func CreateRsmEmmTriggerTypeHandOutUeAttach() e2smrsm.RsmEmmTriggerType {
	return e2smrsm.RsmEmmTriggerType_RSM_EMM_TRIGGER_TYPE_HAND_OUT_UE_ATTACH
}

func CreateUeIDtypeCuUeF1ApID() e2smrsm.UeIdType {
	return e2smrsm.UeIdType_UE_ID_TYPE_CU_UE_F1_AP_ID
}

func CreateUeIDtypeDuUeF1ApID() e2smrsm.UeIdType {
	return e2smrsm.UeIdType_UE_ID_TYPE_DU_UE_F1_AP_ID
}

func CreateUeIDtypeRanUeNgapID() e2smrsm.UeIdType {
	return e2smrsm.UeIdType_UE_ID_TYPE_RAN_UE_NGAP_ID
}

func CreateUeIDtypeAmfUeNgapID() e2smrsm.UeIdType {
	return e2smrsm.UeIdType_UE_ID_TYPE_AMF_UE_NGAP_ID
}

func CreateUeIDtypeEnbUeS1ApID() e2smrsm.UeIdType {
	return e2smrsm.UeIdType_UE_ID_TYPE_ENB_UE_S1_AP_ID
}

func CreateBearerIDdrb(drbID *e2smrsm.DrbId) (*e2smrsm.BearerId, error) {

	drb := &e2smrsm.BearerId{
		BearerId: &e2smrsm.BearerId_DrbId{
			DrbId: drbID,
		},
	}

	if err := drb.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateBearerIDdrb(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return drb, nil
}

func CreateDrbIDfourG(val int32, qci int32) (*e2smrsm.DrbId, error) {

	if qci < 0 || qci > 255 {
		return nil, errors.NewInvalid("QCI value should be in range 0 to 255")
	}

	fourg := &e2smrsm.DrbId{
		DrbId: &e2smrsm.DrbId_FourGdrbId{
			FourGdrbId: &e2smrsm.FourGDrbId{
				Value: val,
				Qci: &e2smv2.Qci{
					Value: qci,
				},
			},
		},
	}

	if err := fourg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateDrbIDfourG(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return fourg, nil
}

func CreateDrbIDfiveG(val int32, qfi int32, flowMap []*e2smrsm.QoSflowLevelParameters) (*e2smrsm.DrbId, error) {

	if val < 1 || val > 32 {
		return nil, errors.NewInvalid("FiveGdrbID value should be in range 1 to 32")
	}

	if qfi < 0 || qfi > 63 {
		return nil, errors.NewInvalid("QCI value should be in range 0 to 255")
	}

	if len(flowMap) < 1 || len(flowMap) > 64 {
		return nil, errors.NewInvalid("FlowsMapToDrb list should have 1 to 64 items")
	}

	fiveg := &e2smrsm.DrbId{
		DrbId: &e2smrsm.DrbId_FiveGdrbId{
			FiveGdrbId: &e2smrsm.FiveGDrbId{
				Value: val,
				Qfi: &e2smrsm.Qfi{
					Value: qfi,
				},
				FlowsMapToDrb: flowMap,
			},
		},
	}

	if err := fiveg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateDrbIDfiveG(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return fiveg, nil
}

func CreateQosFlowLevelParametersDynamic(prlvl int32, pDelay int32, per int32) (*e2smrsm.QoSflowLevelParameters, error) {

	qos := &e2smrsm.QoSflowLevelParameters{
		QoSflowLevelParameters: &e2smrsm.QoSflowLevelParameters_DynamicFiveQi{
			DynamicFiveQi: &e2smrsm.DynamicFiveQi{
				PriorityLevel:     prlvl,
				PacketDelayBudget: pDelay,
				PacketErrorRate:   per,
			},
		},
	}

	if err := qos.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateQosFlowLevelParametersDynamic(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return qos, nil
}

func CreateQosFlowLevelParametersNonDynamic(fiveQI int32) (*e2smrsm.QoSflowLevelParameters, error) {

	qos := &e2smrsm.QoSflowLevelParameters{
		QoSflowLevelParameters: &e2smrsm.QoSflowLevelParameters_NonDynamicFiveQi{
			NonDynamicFiveQi: &e2smrsm.NonDynamicFiveQi{
				FiveQi: &e2smv2.FiveQi{
					Value: fiveQI,
				},
			},
		},
	}

	if err := qos.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateQosFlowLevelParametersNonDynamic(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return qos, nil
}
