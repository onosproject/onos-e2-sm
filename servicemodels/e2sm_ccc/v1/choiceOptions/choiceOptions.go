// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package choiceOptions

import (
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-common-ies"
	"reflect"
)

var E2smCccChoicemap = map[string]map[int]reflect.Type{
	"ran_configuration_structure": {
		1: reflect.TypeOf(e2smcccv1.RanConfigurationStructure_OGnbCuCpFunction{}),
		2: reflect.TypeOf(e2smcccv1.RanConfigurationStructure_OGnbCuUpFunction{}),
		3: reflect.TypeOf(e2smcccv1.RanConfigurationStructure_OGnbDuFunction{}),
		4: reflect.TypeOf(e2smcccv1.RanConfigurationStructure_ONrCellCu{}),
		5: reflect.TypeOf(e2smcccv1.RanConfigurationStructure_ONrCellDu{}),
		6: reflect.TypeOf(e2smcccv1.RanConfigurationStructure_ORrmpolicyRatio{}),
		7: reflect.TypeOf(e2smcccv1.RanConfigurationStructure_OBwp{}),
	},
	"indication_header_format": {
		1: reflect.TypeOf(e2smcccv1.IndicationHeaderFormat_E2SmCccIndicationHeaderFormat1{}),
	},
	"indication_message_format": {
		1: reflect.TypeOf(e2smcccv1.IndicationMessageFormat_E2SmCccIndicationMessageFormat1{}),
		2: reflect.TypeOf(e2smcccv1.IndicationMessageFormat_E2SmCccIndicationMessageFormat2{}),
	},
	"cell_global_id": {
		1: reflect.TypeOf(e2smcccv1.CellGlobalId_NRCgi{}),
		2: reflect.TypeOf(e2smcccv1.CellGlobalId_EUtraCgi{}),
	},
	"control_header_format": {
		1: reflect.TypeOf(e2smcccv1.ControlHeaderFormat_E2SmCccControlHeaderFormat1{}),
	},
	"control_message_format": {
		1: reflect.TypeOf(e2smcccv1.ControlMessageFormat_E2SmCccControlMessageFormat1{}),
		2: reflect.TypeOf(e2smcccv1.ControlMessageFormat_E2SmCccControlMessageFormat2{}),
	},
	"event_trigger_definition_format": {
		1: reflect.TypeOf(e2smcccv1.EventTriggerDefinitionFormat_E2SmCccEventTriggerDefinitionFormat1{}),
		2: reflect.TypeOf(e2smcccv1.EventTriggerDefinitionFormat_E2SmCccEventTriggerDefinitionFormat2{}),
		3: reflect.TypeOf(e2smcccv1.EventTriggerDefinitionFormat_E2SmCccEventTriggerDefinitionFormat3{}),
	},
	"action_definition_format": {
		1: reflect.TypeOf(e2smcccv1.ActionDefinitionFormat_E2SmCccActionDefinitionFormat1{}),
		2: reflect.TypeOf(e2smcccv1.ActionDefinitionFormat_E2SmCccActionDefinitionFormat2{}),
	},
	"cgi": {
		1: reflect.TypeOf(e2smcommoniesv1.Cgi_NRCgi{}),
		2: reflect.TypeOf(e2smcommoniesv1.Cgi_EUtraCgi{}),
	},
	"core_cpid": {
		1: reflect.TypeOf(e2smcommoniesv1.CoreCpid_FiveGc{}),
		2: reflect.TypeOf(e2smcommoniesv1.CoreCpid_EPc{}),
	},
	"interface_identifier": {
		1: reflect.TypeOf(e2smcommoniesv1.InterfaceIdentifier_NG{}),
		2: reflect.TypeOf(e2smcommoniesv1.InterfaceIdentifier_XN{}),
		3: reflect.TypeOf(e2smcommoniesv1.InterfaceIdentifier_F1{}),
		4: reflect.TypeOf(e2smcommoniesv1.InterfaceIdentifier_E1{}),
		5: reflect.TypeOf(e2smcommoniesv1.InterfaceIdentifier_S1{}),
		6: reflect.TypeOf(e2smcommoniesv1.InterfaceIdentifier_X2{}),
		7: reflect.TypeOf(e2smcommoniesv1.InterfaceIdentifier_W1{}),
	},
	"node_type_interface_id_x2": {
		1: reflect.TypeOf(e2smcommoniesv1.NodeTypeInterfaceIdX2_GlobalENbId{}),
		2: reflect.TypeOf(e2smcommoniesv1.NodeTypeInterfaceIdX2_GlobalEnGNbId{}),
	},
	"group_id": {
		1: reflect.TypeOf(e2smcommoniesv1.GroupId_FiveGc{}),
		2: reflect.TypeOf(e2smcommoniesv1.GroupId_EPc{}),
	},
	"qo_sid": {
		1: reflect.TypeOf(e2smcommoniesv1.QoSid_FiveGc{}),
		2: reflect.TypeOf(e2smcommoniesv1.QoSid_EPc{}),
	},
	"rrc_type_rrc_message_id": {
		1: reflect.TypeOf(e2smcommoniesv1.RrcTypeRrcMessageId_LTe{}),
		2: reflect.TypeOf(e2smcommoniesv1.RrcTypeRrcMessageId_NR{}),
	},
	"serving_cell_arfcn": {
		1: reflect.TypeOf(e2smcommoniesv1.ServingCellARfcn_NR{}),
		2: reflect.TypeOf(e2smcommoniesv1.ServingCellARfcn_EUtra{}),
	},
	"serving_cell_pci": {
		1: reflect.TypeOf(e2smcommoniesv1.ServingCellPCi_NR{}),
		2: reflect.TypeOf(e2smcommoniesv1.ServingCellPCi_EUtra{}),
	},
	"ueid": {
		1: reflect.TypeOf(e2smcommoniesv1.Ueid_GNbUeid{}),
		2: reflect.TypeOf(e2smcommoniesv1.Ueid_GNbDuUeid{}),
		3: reflect.TypeOf(e2smcommoniesv1.Ueid_GNbCuUpUeid{}),
		4: reflect.TypeOf(e2smcommoniesv1.Ueid_NgENbUeid{}),
		5: reflect.TypeOf(e2smcommoniesv1.Ueid_NgENbDuUeid{}),
		6: reflect.TypeOf(e2smcommoniesv1.Ueid_EnGNbUeid{}),
		7: reflect.TypeOf(e2smcommoniesv1.Ueid_ENbUeid{}),
	},
	"enb_id": {
		1: reflect.TypeOf(e2smcommoniesv1.EnbID_MacroENbId{}),
		2: reflect.TypeOf(e2smcommoniesv1.EnbID_HomeENbId{}),
		3: reflect.TypeOf(e2smcommoniesv1.EnbID_ShortMacroENbId{}),
		4: reflect.TypeOf(e2smcommoniesv1.EnbID_LongMacroENbId{}),
	},
	"en_gnb_id": {
		1: reflect.TypeOf(e2smcommoniesv1.EnGNbID_EnGNbId{}),
	},
	"gnb_id": {
		1: reflect.TypeOf(e2smcommoniesv1.GnbID_GNbId{}),
	},
	"ng_enb_id": {
		1: reflect.TypeOf(e2smcommoniesv1.NgEnbID_MacroNgEnbId{}),
		2: reflect.TypeOf(e2smcommoniesv1.NgEnbID_ShortMacroNgEnbId{}),
		3: reflect.TypeOf(e2smcommoniesv1.NgEnbID_LongMacroNgEnbId{}),
	},
	"global_ngrannode_id": {
		1: reflect.TypeOf(e2smcommoniesv1.GlobalNgrannodeId_GNb{}),
		2: reflect.TypeOf(e2smcommoniesv1.GlobalNgrannodeId_NgENb{}),
	},
}
