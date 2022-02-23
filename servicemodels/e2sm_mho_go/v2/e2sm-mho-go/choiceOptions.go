// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smmho

import (
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"reflect"
)

var MhoChoicemap = map[string]map[int]reflect.Type{
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
	"cgi": {
		1: reflect.TypeOf(e2sm_v2_ies.Cgi_NRCgi{}),
		2: reflect.TypeOf(e2sm_v2_ies.Cgi_EUtraCgi{}),
	},
	"core_cpid": {
		1: reflect.TypeOf(e2sm_v2_ies.CoreCpid_FiveGc{}),
		2: reflect.TypeOf(e2sm_v2_ies.CoreCpid_EPc{}),
	},
	"interface_identifier": {
		1: reflect.TypeOf(e2sm_v2_ies.InterfaceIdentifier_NG{}),
		2: reflect.TypeOf(e2sm_v2_ies.InterfaceIdentifier_XN{}),
		3: reflect.TypeOf(e2sm_v2_ies.InterfaceIdentifier_F1{}),
		4: reflect.TypeOf(e2sm_v2_ies.InterfaceIdentifier_E1{}),
		5: reflect.TypeOf(e2sm_v2_ies.InterfaceIdentifier_S1{}),
		6: reflect.TypeOf(e2sm_v2_ies.InterfaceIdentifier_X2{}),
		7: reflect.TypeOf(e2sm_v2_ies.InterfaceIdentifier_W1{}),
	},
	"node_type": {
		1: reflect.TypeOf(e2sm_v2_ies.NodeType_GlobalEnbId{}),
		2: reflect.TypeOf(e2sm_v2_ies.NodeType_GlobalEnGnbId{}),
	},
	"group_id": {
		1: reflect.TypeOf(e2sm_v2_ies.GroupId_FiveGc{}),
		2: reflect.TypeOf(e2sm_v2_ies.GroupId_EPc{}),
	},
	"qo_sid": {
		1: reflect.TypeOf(e2sm_v2_ies.QoSid_FiveGc{}),
		2: reflect.TypeOf(e2sm_v2_ies.QoSid_EPc{}),
	},
	"rrc_type": {
		1: reflect.TypeOf(e2sm_v2_ies.RrcType_Lte{}),
		2: reflect.TypeOf(e2sm_v2_ies.RrcType_Nr{}),
	},
	"serving_cell_arfcn": {
		1: reflect.TypeOf(e2sm_v2_ies.ServingCellArfcn_NR{}),
		2: reflect.TypeOf(e2sm_v2_ies.ServingCellArfcn_EUtra{}),
	},
	"serving_cell_pci": {
		1: reflect.TypeOf(e2sm_v2_ies.ServingCellPci_NR{}),
		2: reflect.TypeOf(e2sm_v2_ies.ServingCellPci_EUtra{}),
	},
	"ueid": {
		1: reflect.TypeOf(e2sm_v2_ies.Ueid_GNbUeid{}),
		2: reflect.TypeOf(e2sm_v2_ies.Ueid_GNbDuUeid{}),
		3: reflect.TypeOf(e2sm_v2_ies.Ueid_GNbCuUpUeid{}),
		4: reflect.TypeOf(e2sm_v2_ies.Ueid_NgENbUeid{}),
		5: reflect.TypeOf(e2sm_v2_ies.Ueid_NgENbDuUeid{}),
		6: reflect.TypeOf(e2sm_v2_ies.Ueid_EnGNbUeid{}),
		7: reflect.TypeOf(e2sm_v2_ies.Ueid_ENbUeid{}),
	},
	"enb_id": {
		1: reflect.TypeOf(e2sm_v2_ies.EnbId_MacroENbId{}),
		2: reflect.TypeOf(e2sm_v2_ies.EnbId_HomeENbId{}),
		3: reflect.TypeOf(e2sm_v2_ies.EnbId_ShortMacroENbId{}),
		4: reflect.TypeOf(e2sm_v2_ies.EnbId_LongMacroENbId{}),
	},
	"en_gnb_id": {
		1: reflect.TypeOf(e2sm_v2_ies.EnGnbId_EnGNbId{}),
	},
	"global_rannode_id": {
		1: reflect.TypeOf(e2sm_v2_ies.GlobalRannodeId_GlobalGnbId{}),
		2: reflect.TypeOf(e2sm_v2_ies.GlobalRannodeId_GlobalNgEnbId{}),
	},
	"gnb_id": {
		1: reflect.TypeOf(e2sm_v2_ies.GnbId_GNbId{}),
	},
	"ng_enb_id": {
		1: reflect.TypeOf(e2sm_v2_ies.NgEnbId_MacroNgEnbId{}),
		2: reflect.TypeOf(e2sm_v2_ies.NgEnbId_ShortMacroNgEnbId{}),
		3: reflect.TypeOf(e2sm_v2_ies.NgEnbId_LongMacroNgEnbId{}),
	},
}
