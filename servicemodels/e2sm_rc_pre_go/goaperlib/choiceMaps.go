// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package goaperlib

import (
	e2sm_rc_pre_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"reflect"
)

var rcPreChoicemap = map[string]map[int]reflect.Type{
	"cell_global_id": {
		1: reflect.TypeOf(e2sm_rc_pre_v2_go.CellGlobalId_NrCgi{}),
		2: reflect.TypeOf(e2sm_rc_pre_v2_go.CellGlobalId_EUtraCgi{}),
	},
	"arfcn": {
		1: reflect.TypeOf(e2sm_rc_pre_v2_go.Arfcn_EArfcn{}),
		2: reflect.TypeOf(e2sm_rc_pre_v2_go.Arfcn_NrArfcn{}),
	},
	"ranparameter_value": {
		1: reflect.TypeOf(e2sm_rc_pre_v2_go.RanparameterValue_ValueInt{}),
		2: reflect.TypeOf(e2sm_rc_pre_v2_go.RanparameterValue_ValueEnum{}),
		3: reflect.TypeOf(e2sm_rc_pre_v2_go.RanparameterValue_ValueBool{}),
		4: reflect.TypeOf(e2sm_rc_pre_v2_go.RanparameterValue_ValueBitS{}),
		5: reflect.TypeOf(e2sm_rc_pre_v2_go.RanparameterValue_ValueOctS{}),
		6: reflect.TypeOf(e2sm_rc_pre_v2_go.RanparameterValue_ValuePrtS{}),
	},
	"e2sm_rc_pre_event_trigger_definition_event_definition_formats": {
		1: reflect.TypeOf(e2sm_rc_pre_v2_go.EventTriggerDefinitionFormats_EventDefinitionFormat1{}),
	},
	"e2_sm_rc_pre_indication_header": {
		1: reflect.TypeOf(e2sm_rc_pre_v2_go.E2SmRcPreIndicationHeader_IndicationHeaderFormat1{}),
	},
	"e2_sm_rc_pre_indication_message": {
		1: reflect.TypeOf(e2sm_rc_pre_v2_go.E2SmRcPreIndicationMessage_RicStyleType{}),
		2: reflect.TypeOf(e2sm_rc_pre_v2_go.E2SmRcPreIndicationMessage_IndicationMessageFormat1{}),
	},
	"e2_sm_rc_pre_control_header": {
		1: reflect.TypeOf(e2sm_rc_pre_v2_go.E2SmRcPreControlHeader_ControlHeaderFormat1{}),
	},
	"e2_sm_rc_pre_control_message": {
		1: reflect.TypeOf(e2sm_rc_pre_v2_go.E2SmRcPreControlMessage_ControlMessage{}),
	},
	"e2_sm_rc_pre_control_outcome": {
		1: reflect.TypeOf(e2sm_rc_pre_v2_go.E2SmRcPreControlOutcome_ControlOutcomeFormat1{}),
	},
}
