// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2sm_mho_go

import "reflect"

var MhoChoicemap = map[string]map[int]reflect.Type{
	"cell_global_id": {
		1: reflect.TypeOf(CellGlobalId_NrCgi{}),
		2: reflect.TypeOf(CellGlobalId_EUtraCgi{}),
	},
	"e2_sm_mho_event_trigger_definition": {
		1: reflect.TypeOf(MhoEventTriggerDefinitionFormats_EventDefinitionFormat1{}),
	},
	"e2_sm_mho_indication_header": {
		1: reflect.TypeOf(E2SmMhoIndicationHeader_IndicationHeaderFormat1{}),
	},
	"e2_sm_mho_indication_message": {
		1: reflect.TypeOf(E2SmMhoIndicationMessage_IndicationMessageFormat1{}),
		2: reflect.TypeOf(E2SmMhoIndicationMessage_IndicationMessageFormat2{}),
	},
	"e2_sm_mho_control_header": {
		1: reflect.TypeOf(E2SmMhoControlHeader_ControlHeaderFormat1{}),
	},
	"e2_sm_mho_control_message": {
		1: reflect.TypeOf(E2SmMhoControlMessage_ControlMessageFormat1{}),
	},
}
