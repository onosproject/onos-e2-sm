// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smv2ies

import "reflect"

var E2SmChoicemap = map[string]map[int]reflect.Type{
	"cgi": {
		1: reflect.TypeOf(Cgi_NRCgi{}),
		2: reflect.TypeOf(Cgi_EUtraCgi{}),
	},
	"core_cpid": {
		1: reflect.TypeOf(CoreCpid_FiveGc{}),
		2: reflect.TypeOf(CoreCpid_EPc{}),
	},
	"interface_identifier": {
		1: reflect.TypeOf(InterfaceIdentifier_NG{}),
		2: reflect.TypeOf(InterfaceIdentifier_XN{}),
		3: reflect.TypeOf(InterfaceIdentifier_F1{}),
		4: reflect.TypeOf(InterfaceIdentifier_E1{}),
		5: reflect.TypeOf(InterfaceIdentifier_S1{}),
		6: reflect.TypeOf(InterfaceIdentifier_X2{}),
		7: reflect.TypeOf(InterfaceIdentifier_W1{}),
	},
	"node_type": {
		1: reflect.TypeOf(NodeType_GlobalEnbId{}),
		2: reflect.TypeOf(NodeType_GlobalEnGnbId{}),
	},
	"group_id": {
		1: reflect.TypeOf(GroupId_FiveGc{}),
		2: reflect.TypeOf(GroupId_EPc{}),
	},
	"qo_sid": {
		1: reflect.TypeOf(QoSid_FiveGc{}),
		2: reflect.TypeOf(QoSid_EPc{}),
	},
	"rrc_type": {
		1: reflect.TypeOf(RrcType_Lte{}),
		2: reflect.TypeOf(RrcType_Nr{}),
	},
	"serving_cell_arfcn": {
		1: reflect.TypeOf(ServingCellArfcn_NR{}),
		2: reflect.TypeOf(ServingCellArfcn_EUtra{}),
	},
	"serving_cell_pci": {
		1: reflect.TypeOf(ServingCellPci_NR{}),
		2: reflect.TypeOf(ServingCellPci_EUtra{}),
	},
	"ueid": {
		1: reflect.TypeOf(Ueid_GNbUeid{}),
		2: reflect.TypeOf(Ueid_GNbDuUeid{}),
		3: reflect.TypeOf(Ueid_GNbCuUpUeid{}),
		4: reflect.TypeOf(Ueid_NgENbUeid{}),
		5: reflect.TypeOf(Ueid_NgENbDuUeid{}),
		6: reflect.TypeOf(Ueid_EnGNbUeid{}),
		7: reflect.TypeOf(Ueid_ENbUeid{}),
	},
	"enb_id": {
		1: reflect.TypeOf(EnbId_MacroENbId{}),
		2: reflect.TypeOf(EnbId_HomeENbId{}),
		3: reflect.TypeOf(EnbId_ShortMacroENbId{}),
		4: reflect.TypeOf(EnbId_LongMacroENbId{}),
	},
	"en_gnb_id": {
		1: reflect.TypeOf(EnGnbId_EnGNbId{}),
	},
	"global_rannode_id": {
		1: reflect.TypeOf(GlobalRannodeId_GlobalGnbId{}),
		2: reflect.TypeOf(GlobalRannodeId_GlobalNgEnbId{}),
	},
	"gnb_id": {
		1: reflect.TypeOf(GnbId_GNbId{}),
	},
	"ng_enb_id": {
		1: reflect.TypeOf(NgEnbId_MacroNgEnbId{}),
		2: reflect.TypeOf(NgEnbId_ShortMacroNgEnbId{}),
		3: reflect.TypeOf(NgEnbId_LongMacroNgEnbId{}),
	},
}
