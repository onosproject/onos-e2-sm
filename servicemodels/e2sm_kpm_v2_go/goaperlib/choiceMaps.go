// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package goaperlib

import (
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"reflect"
)

var choicemape2smKpm = map[string]map[int]reflect.Type{
	"cell_global_id": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.CellGlobalId_NrCgi{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.CellGlobalId_EUtraCgi{}),
	},
	"measurement_type": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.MeasurementType_MeasName{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.MeasurementType_MeasId{}),
	},
	"test_cond_type": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.TestCondType_GBr{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.TestCondType_AMbr{}),
		3: reflect.TypeOf(e2sm_kpm_v2_go.TestCondType_IsStat{}),
		4: reflect.TypeOf(e2sm_kpm_v2_go.TestCondType_IsCatM{}),
		5: reflect.TypeOf(e2sm_kpm_v2_go.TestCondType_RSrp{}),
		6: reflect.TypeOf(e2sm_kpm_v2_go.TestCondType_RSrq{}),
	},
	"test_cond_value": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.TestCondValue_ValueInt{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.TestCondValue_ValueEnum{}),
		3: reflect.TypeOf(e2sm_kpm_v2_go.TestCondValue_ValueBool{}),
		4: reflect.TypeOf(e2sm_kpm_v2_go.TestCondValue_ValueBitS{}),
		5: reflect.TypeOf(e2sm_kpm_v2_go.TestCondValue_ValueOctS{}),
		6: reflect.TypeOf(e2sm_kpm_v2_go.TestCondValue_ValuePrtS{}),
	},
	"global_kpmnode_id": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.GlobalKpmnodeId_GNb{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.GlobalKpmnodeId_EnGNb{}),
		3: reflect.TypeOf(e2sm_kpm_v2_go.GlobalKpmnodeId_NgENb{}),
		4: reflect.TypeOf(e2sm_kpm_v2_go.GlobalKpmnodeId_ENb{}),
	},
	"gnb_id_choice": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.GnbIdChoice_GnbId{}),
	},
	"engnb_id": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.EngnbId_GNbId{}),
	},
	"enb_id_choice": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.EnbIdChoice_EnbIdMacro{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.EnbIdChoice_EnbIdShortmacro{}),
		3: reflect.TypeOf(e2sm_kpm_v2_go.EnbIdChoice_EnbIdLongmacro{}),
	},
	"enb_id": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.EnbId_MacroENbId{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.EnbId_HomeENbId{}),
	},
	"measurement_record_item": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.MeasurementRecordItem_Integer{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.MeasurementRecordItem_Real{}),
		3: reflect.TypeOf(e2sm_kpm_v2_go.MeasurementRecordItem_NoValue{}),
	},
	"matching_cond_item": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.MatchingCondItem_MeasLabel{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.MatchingCondItem_TestCondInfo{}),
	},
	"e2_sm_kpm_event_trigger_definition": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.EventTriggerDefinitionFormats_EventDefinitionFormat1{}),
	},
	"e2_sm_kpm_action_definition": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.ActionDefinitionFormats_ActionDefinitionFormat1{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.ActionDefinitionFormats_ActionDefinitionFormat2{}),
		3: reflect.TypeOf(e2sm_kpm_v2_go.ActionDefinitionFormats_ActionDefinitionFormat3{}),
	},
	"e2_sm_kpm_indication_header": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.IndicationHeaderFormats_IndicationHeaderFormat1{}),
	},
	"e2_sm_kpm_indication_message": {
		1: reflect.TypeOf(e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat1{}),
		2: reflect.TypeOf(e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat2{}),
	},
}
