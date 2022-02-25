// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smrsmies

import (
	e2smv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies"
	"reflect"
)

var RsmChoicemap = map[string]map[int]reflect.Type{
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
		2: reflect.TypeOf(E2SmRsmIndicationMessage_IndicationMessageFormat2{}),
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
	"cgi": {
		1: reflect.TypeOf(e2smv2.Cgi_NRCgi{}),
		2: reflect.TypeOf(e2smv2.Cgi_EUtraCgi{}),
	},
	"core_cpid": {
		1: reflect.TypeOf(e2smv2.CoreCpid_FiveGc{}),
		2: reflect.TypeOf(e2smv2.CoreCpid_EPc{}),
	},
	"interface_identifier": {
		1: reflect.TypeOf(e2smv2.InterfaceIdentifier_NG{}),
		2: reflect.TypeOf(e2smv2.InterfaceIdentifier_XN{}),
		3: reflect.TypeOf(e2smv2.InterfaceIdentifier_F1{}),
		4: reflect.TypeOf(e2smv2.InterfaceIdentifier_E1{}),
		5: reflect.TypeOf(e2smv2.InterfaceIdentifier_S1{}),
		6: reflect.TypeOf(e2smv2.InterfaceIdentifier_X2{}),
		7: reflect.TypeOf(e2smv2.InterfaceIdentifier_W1{}),
	},
	"node_type": {
		1: reflect.TypeOf(e2smv2.NodeType_GlobalEnbId{}),
		2: reflect.TypeOf(e2smv2.NodeType_GlobalEnGnbId{}),
	},
	"group_id": {
		1: reflect.TypeOf(e2smv2.GroupId_FiveGc{}),
		2: reflect.TypeOf(e2smv2.GroupId_EPc{}),
	},
	"qo_sid": {
		1: reflect.TypeOf(e2smv2.QoSid_FiveGc{}),
		2: reflect.TypeOf(e2smv2.QoSid_EPc{}),
	},
	"rrc_type": {
		1: reflect.TypeOf(e2smv2.RrcType_Lte{}),
		2: reflect.TypeOf(e2smv2.RrcType_Nr{}),
	},
	"serving_cell_arfcn": {
		1: reflect.TypeOf(e2smv2.ServingCellArfcn_NR{}),
		2: reflect.TypeOf(e2smv2.ServingCellArfcn_EUtra{}),
	},
	"serving_cell_pci": {
		1: reflect.TypeOf(e2smv2.ServingCellPci_NR{}),
		2: reflect.TypeOf(e2smv2.ServingCellPci_EUtra{}),
	},
	"ueid": {
		1: reflect.TypeOf(e2smv2.Ueid_GNbUeid{}),
		2: reflect.TypeOf(e2smv2.Ueid_GNbDuUeid{}),
		3: reflect.TypeOf(e2smv2.Ueid_GNbCuUpUeid{}),
		4: reflect.TypeOf(e2smv2.Ueid_NgENbUeid{}),
		5: reflect.TypeOf(e2smv2.Ueid_NgENbDuUeid{}),
		6: reflect.TypeOf(e2smv2.Ueid_EnGNbUeid{}),
		7: reflect.TypeOf(e2smv2.Ueid_ENbUeid{}),
	},
	"enb_id": {
		1: reflect.TypeOf(e2smv2.EnbId_MacroENbId{}),
		2: reflect.TypeOf(e2smv2.EnbId_HomeENbId{}),
		3: reflect.TypeOf(e2smv2.EnbId_ShortMacroENbId{}),
		4: reflect.TypeOf(e2smv2.EnbId_LongMacroENbId{}),
	},
	"en_gnb_id": {
		1: reflect.TypeOf(e2smv2.EnGnbId_EnGNbId{}),
	},
	"global_rannode_id": {
		1: reflect.TypeOf(e2smv2.GlobalRannodeId_GlobalGnbId{}),
		2: reflect.TypeOf(e2smv2.GlobalRannodeId_GlobalNgEnbId{}),
	},
	"gnb_id": {
		1: reflect.TypeOf(e2smv2.GnbId_GNbId{}),
	},
	"ng_enb_id": {
		1: reflect.TypeOf(e2smv2.NgEnbId_MacroNgEnbId{}),
		2: reflect.TypeOf(e2smv2.NgEnbId_ShortMacroNgEnbId{}),
		3: reflect.TypeOf(e2smv2.NgEnbId_LongMacroNgEnbId{}),
	},
}
