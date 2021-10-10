// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2sm_rc_pre_v2_go

import "reflect"

var RcPreChoicemap = map[string]map[int]reflect.Type{
	"cell_global_id": {
		1: reflect.TypeOf(CellGlobalId_NrCgi{}),
		2: reflect.TypeOf(CellGlobalId_EUtraCgi{}),
	},
	"arfcn": {
		1: reflect.TypeOf(Arfcn_EArfcn{}),
		2: reflect.TypeOf(Arfcn_NrArfcn{}),
	},
	"ranparameter_value": {
		1: reflect.TypeOf(RanparameterValue_ValueInt{}),
		2: reflect.TypeOf(RanparameterValue_ValueEnum{}),
		3: reflect.TypeOf(RanparameterValue_ValueBool{}),
		4: reflect.TypeOf(RanparameterValue_ValueBitS{}),
		5: reflect.TypeOf(RanparameterValue_ValueOctS{}),
		6: reflect.TypeOf(RanparameterValue_ValuePrtS{}),
	},
	"e2sm_rc_pre_event_trigger_definition_event_definition_formats": {
		1: reflect.TypeOf(EventTriggerDefinitionFormats_EventDefinitionFormat1{}),
	},
	"e2_sm_rc_pre_indication_header": {
		1: reflect.TypeOf(E2SmRcPreIndicationHeader_IndicationHeaderFormat1{}),
	},
	"e2_sm_rc_pre_indication_message": {
		1: reflect.TypeOf(E2SmRcPreIndicationMessage_RicStyleType{}),
		2: reflect.TypeOf(E2SmRcPreIndicationMessage_IndicationMessageFormat1{}),
	},
	"e2_sm_rc_pre_control_header": {
		1: reflect.TypeOf(E2SmRcPreControlHeader_ControlHeaderFormat1{}),
	},
	"e2_sm_rc_pre_control_message": {
		1: reflect.TypeOf(E2SmRcPreControlMessage_ControlMessage{}),
	},
	"e2_sm_rc_pre_control_outcome": {
		1: reflect.TypeOf(E2SmRcPreControlOutcome_ControlOutcomeFormat1{}),
	},
}
