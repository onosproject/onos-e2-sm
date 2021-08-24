// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2sm_rsm_ies

import "reflect"

var rsmChoicemap = map[string]map[int]reflect.Type{
	"ue_identity": {
		1: reflect.TypeOf(UeIdentity_CuUeF1ApId{}),
		2: reflect.TypeOf(UeIdentity_DuUeF1ApId{}),
		3: reflect.TypeOf(UeIdentity_RanUeNgapId{}),
		4: reflect.TypeOf(UeIdentity_AmfUeNgapId{}),
		5: reflect.TypeOf(UeIdentity_EnbUeS1ApId{}),
	},
	"e2_sm_rsm_event_definition": {
		1: reflect.TypeOf(EventDefinitionFormats_EventDefinitionFormat1{}),
	},
	"e2_sm_rsm_indication_header": {
		1: reflect.TypeOf(E2SmRsmIndicationHeader_IndicationHeaderFormat1{}),
	},
	"e2_sm_rsm_indication_message": {
		1: reflect.TypeOf(E2SmRsmIndicationMessage_IndicationMessageFormat1{}),
	},
	"bearer_id": {
		1: reflect.TypeOf(BearerId_DrbId{}),
	},
	"drb_id": {
		1: reflect.TypeOf(DrbId_FourGdrbId{}),
		2: reflect.TypeOf(DrbId_FiveGdrbId{}),
	},
	"qo_sflow_level_parameters": {
		1: reflect.TypeOf(QoSflowLevelParameters_DynamicFiveQi{}),
		2: reflect.TypeOf(QoSflowLevelParameters_NonDynamicFiveQi{}),
	},
	"e2_sm_rsm_control_message": {
		1: reflect.TypeOf(E2SmRsmControlMessage_SliceCreate{}),
		2: reflect.TypeOf(E2SmRsmControlMessage_SliceUpdate{}),
		3: reflect.TypeOf(E2SmRsmControlMessage_SliceDelete{}),
		4: reflect.TypeOf(E2SmRsmControlMessage_SliceAssociate{}),
	},
}
