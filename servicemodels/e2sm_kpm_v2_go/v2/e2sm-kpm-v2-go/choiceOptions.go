// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2sm_kpm_v2_go

import "reflect"

var Choicemape2smKpm = map[string]map[int]reflect.Type{
	"cell_global_id": {
		1: reflect.TypeOf(CellGlobalId_NrCgi{}),
		2: reflect.TypeOf(CellGlobalId_EUtraCgi{}),
	},
	"measurement_type": {
		1: reflect.TypeOf(MeasurementType_MeasName{}),
		2: reflect.TypeOf(MeasurementType_MeasId{}),
	},
	"test_cond_type": {
		1: reflect.TypeOf(TestCondType_GBr{}),
		2: reflect.TypeOf(TestCondType_AMbr{}),
		3: reflect.TypeOf(TestCondType_IsStat{}),
		4: reflect.TypeOf(TestCondType_IsCatM{}),
		5: reflect.TypeOf(TestCondType_RSrp{}),
		6: reflect.TypeOf(TestCondType_RSrq{}),
	},
	"test_cond_value": {
		1: reflect.TypeOf(TestCondValue_ValueInt{}),
		2: reflect.TypeOf(TestCondValue_ValueEnum{}),
		3: reflect.TypeOf(TestCondValue_ValueBool{}),
		4: reflect.TypeOf(TestCondValue_ValueBitS{}),
		5: reflect.TypeOf(TestCondValue_ValueOctS{}),
		6: reflect.TypeOf(TestCondValue_ValuePrtS{}),
	},
	"global_kpmnode_id": {
		1: reflect.TypeOf(GlobalKpmnodeId_GNb{}),
		2: reflect.TypeOf(GlobalKpmnodeId_EnGNb{}),
		3: reflect.TypeOf(GlobalKpmnodeId_NgENb{}),
		4: reflect.TypeOf(GlobalKpmnodeId_ENb{}),
	},
	"gnb_id_choice": {
		1: reflect.TypeOf(GnbIdChoice_GnbId{}),
	},
	"engnb_id": {
		1: reflect.TypeOf(EngnbId_GNbId{}),
	},
	"enb_id_choice": {
		1: reflect.TypeOf(EnbIdChoice_EnbIdMacro{}),
		2: reflect.TypeOf(EnbIdChoice_EnbIdShortmacro{}),
		3: reflect.TypeOf(EnbIdChoice_EnbIdLongmacro{}),
	},
	"enb_id": {
		1: reflect.TypeOf(EnbId_MacroENbId{}),
		2: reflect.TypeOf(EnbId_HomeENbId{}),
	},
	"measurement_record_item": {
		1: reflect.TypeOf(MeasurementRecordItem_Integer{}),
		2: reflect.TypeOf(MeasurementRecordItem_Real{}),
		3: reflect.TypeOf(MeasurementRecordItem_NoValue{}),
	},
	"matching_cond_item": {
		1: reflect.TypeOf(MatchingCondItem_MeasLabel{}),
		2: reflect.TypeOf(MatchingCondItem_TestCondInfo{}),
	},
	"e2_sm_kpm_event_trigger_definition": {
		1: reflect.TypeOf(E2SmKpmEventTriggerDefinition_EventDefinitionFormat1{}),
	},
	"e2_sm_kpm_action_definition": {
		1: reflect.TypeOf(E2SmKpmActionDefinition_ActionDefinitionFormat1{}),
		2: reflect.TypeOf(E2SmKpmActionDefinition_ActionDefinitionFormat2{}),
		3: reflect.TypeOf(E2SmKpmActionDefinition_ActionDefinitionFormat3{}),
	},
	"e2_sm_kpm_indication_header": {
		1: reflect.TypeOf(E2SmKpmIndicationHeader_IndicationHeaderFormat1{}),
	},
	"e2_sm_kpm_indication_message": {
		1: reflect.TypeOf(E2SmKpmIndicationMessage_IndicationMessageFormat1{}),
		2: reflect.TypeOf(E2SmKpmIndicationMessage_IndicationMessageFormat2{}),
	},
}
