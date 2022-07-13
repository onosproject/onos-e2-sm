// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package choiceOptions

import (
	e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"reflect"
)

var E2smRcChoicemap = map[string]map[int]reflect.Type{
	"neighbor_cell_item": {
		1: reflect.TypeOf(e2smrcv1.NeighborCellItem_RanTypeChoiceNr{}),
		2: reflect.TypeOf(e2smrcv1.NeighborCellItem_RanTypeChoiceEutra{}),
	},
	"cell_type": {
		1: reflect.TypeOf(e2smrcv1.CellType_CellTypeChoiceIndividual{}),
		2: reflect.TypeOf(e2smrcv1.CellType_CellTypeChoiceGroup{}),
	},
	"ue_type": {
		1: reflect.TypeOf(e2smrcv1.UeType_UeTypeChoiceIndividual{}),
		2: reflect.TypeOf(e2smrcv1.UeType_UeTypeChoiceGroup{}),
	},
	"ranparameter_definition_choice": {
		1: reflect.TypeOf(e2smrcv1.RanparameterDefinitionChoice_ChoiceList{}),
		2: reflect.TypeOf(e2smrcv1.RanparameterDefinitionChoice_ChoiceStructure{}),
	},
	"ranparameter_value": {
		1: reflect.TypeOf(e2smrcv1.RanparameterValue_ValueBoolean{}),
		2: reflect.TypeOf(e2smrcv1.RanparameterValue_ValueInt{}),
		3: reflect.TypeOf(e2smrcv1.RanparameterValue_ValueReal{}),
		4: reflect.TypeOf(e2smrcv1.RanparameterValue_ValueBitS{}),
		5: reflect.TypeOf(e2smrcv1.RanparameterValue_ValueOctS{}),
		6: reflect.TypeOf(e2smrcv1.RanparameterValue_ValuePrintableString{}),
	},
	"ranparameter_value_type": {
		1: reflect.TypeOf(e2smrcv1.RanparameterValueType_RanPChoiceElementTrue{}),
		2: reflect.TypeOf(e2smrcv1.RanparameterValueType_RanPChoiceElementFalse{}),
		3: reflect.TypeOf(e2smrcv1.RanparameterValueType_RanPChoiceStructure{}),
		4: reflect.TypeOf(e2smrcv1.RanparameterValueType_RanPChoiceList{}),
	},
	"ranparameter_testing_condition": {
		1: reflect.TypeOf(e2smrcv1.RanparameterTestingCondition_RanPChoiceComparison{}),
		2: reflect.TypeOf(e2smrcv1.RanparameterTestingCondition_RanPChoicePresence{}),
	},
	"ran_parameter_type": {
		1: reflect.TypeOf(e2smrcv1.RanParameterType_RanPChoiceList{}),
		2: reflect.TypeOf(e2smrcv1.RanParameterType_RanPChoiceStructure{}),
		3: reflect.TypeOf(e2smrcv1.RanParameterType_RanPChoiceElementTrue{}),
		4: reflect.TypeOf(e2smrcv1.RanParameterType_RanPChoiceElementFalse{}),
	},
	"ric_event_trigger_formats": {
		1: reflect.TypeOf(e2smrcv1.RicEventTriggerFormats_EventTriggerFormat1{}),
		2: reflect.TypeOf(e2smrcv1.RicEventTriggerFormats_EventTriggerFormat2{}),
		3: reflect.TypeOf(e2smrcv1.RicEventTriggerFormats_EventTriggerFormat3{}),
		4: reflect.TypeOf(e2smrcv1.RicEventTriggerFormats_EventTriggerFormat4{}),
		5: reflect.TypeOf(e2smrcv1.RicEventTriggerFormats_EventTriggerFormat5{}),
	},
	"message_type_choice": {
		1: reflect.TypeOf(e2smrcv1.MessageTypeChoice_MessageTypeChoiceNi{}),
		2: reflect.TypeOf(e2smrcv1.MessageTypeChoice_MessageTypeChoiceRrc{}),
	},
	"trigger_type_choice": {
		1: reflect.TypeOf(e2smrcv1.TriggerTypeChoice_TriggerTypeChoiceRrcstate{}),
		2: reflect.TypeOf(e2smrcv1.TriggerTypeChoice_TriggerTypeChoiceUeid{}),
		3: reflect.TypeOf(e2smrcv1.TriggerTypeChoice_TriggerTypeChoiceL2State{}),
	},
	"ric_action_definition_formats": {
		1: reflect.TypeOf(e2smrcv1.RicActionDefinitionFormats_ActionDefinitionFormat1{}),
		2: reflect.TypeOf(e2smrcv1.RicActionDefinitionFormats_ActionDefinitionFormat2{}),
		3: reflect.TypeOf(e2smrcv1.RicActionDefinitionFormats_ActionDefinitionFormat3{}),
		4: reflect.TypeOf(e2smrcv1.RicActionDefinitionFormats_ActionDefinitionFormat4{}),
	},
	"ric_indication_header_formats": {
		1: reflect.TypeOf(e2smrcv1.RicIndicationHeaderFormats_IndicationHeaderFormat1{}),
		2: reflect.TypeOf(e2smrcv1.RicIndicationHeaderFormats_IndicationHeaderFormat2{}),
		3: reflect.TypeOf(e2smrcv1.RicIndicationHeaderFormats_IndicationHeaderFormat3{}),
	},
	"ric_indication_message_formats": {
		1: reflect.TypeOf(e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat1{}),
		2: reflect.TypeOf(e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat2{}),
		3: reflect.TypeOf(e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat3{}),
		4: reflect.TypeOf(e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat4{}),
		5: reflect.TypeOf(e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat5{}),
		6: reflect.TypeOf(e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat6{}),
	},
	"ric_call_process_id_formats": {
		1: reflect.TypeOf(e2smrcv1.RicCallProcessIdFormats_CallProcessIdFormat1{}),
	},
	"ric_control_header_formats": {
		1: reflect.TypeOf(e2smrcv1.RicControlHeaderFormats_ControlHeaderFormat1{}),
		2: reflect.TypeOf(e2smrcv1.RicControlHeaderFormats_ControlHeaderFormat2{}),
	},
	"ric_control_message_formats": {
		1: reflect.TypeOf(e2smrcv1.RicControlMessageFormats_ControlMessageFormat1{}),
		2: reflect.TypeOf(e2smrcv1.RicControlMessageFormats_ControlMessageFormat2{}),
	},
	"ric_control_outcome_formats": {
		1: reflect.TypeOf(e2smrcv1.RicControlOutcomeFormats_ControlOutcomeFormat1{}),
		2: reflect.TypeOf(e2smrcv1.RicControlOutcomeFormats_ControlOutcomeFormat2{}),
		3: reflect.TypeOf(e2smrcv1.RicControlOutcomeFormats_ControlOutcomeFormat3{}),
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
	"node_type": {
		1: reflect.TypeOf(e2smcommoniesv1.NodeType_GlobalEnbId{}),
		2: reflect.TypeOf(e2smcommoniesv1.NodeType_GlobalEnGnbId{}),
	},
	"group_id": {
		1: reflect.TypeOf(e2smcommoniesv1.GroupId_FiveGc{}),
		2: reflect.TypeOf(e2smcommoniesv1.GroupId_EPc{}),
	},
	"qo_sid": {
		1: reflect.TypeOf(e2smcommoniesv1.QoSid_FiveGc{}),
		2: reflect.TypeOf(e2smcommoniesv1.QoSid_EPc{}),
	},
	"rrc_type": {
		1: reflect.TypeOf(e2smcommoniesv1.RrcType_Lte{}),
		2: reflect.TypeOf(e2smcommoniesv1.RrcType_Nr{}),
	},
	"serving_cell_arfcn": {
		1: reflect.TypeOf(e2smcommoniesv1.ServingCellArfcn_NR{}),
		2: reflect.TypeOf(e2smcommoniesv1.ServingCellArfcn_EUtra{}),
	},
	"serving_cell_pci": {
		1: reflect.TypeOf(e2smcommoniesv1.ServingCellPci_NR{}),
		2: reflect.TypeOf(e2smcommoniesv1.ServingCellPci_EUtra{}),
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
		1: reflect.TypeOf(e2smcommoniesv1.EnbId_MacroENbId{}),
		2: reflect.TypeOf(e2smcommoniesv1.EnbId_HomeENbId{}),
		3: reflect.TypeOf(e2smcommoniesv1.EnbId_ShortMacroENbId{}),
		4: reflect.TypeOf(e2smcommoniesv1.EnbId_LongMacroENbId{}),
	},
	"en_gnb_id": {
		1: reflect.TypeOf(e2smcommoniesv1.EnGnbId_EnGNbId{}),
	},
	"gnb_id": {
		1: reflect.TypeOf(e2smcommoniesv1.GnbId_GNbId{}),
	},
	"ng_enb_id": {
		1: reflect.TypeOf(e2smcommoniesv1.NgEnbId_MacroNgEnbId{}),
		2: reflect.TypeOf(e2smcommoniesv1.NgEnbId_ShortMacroNgEnbId{}),
		3: reflect.TypeOf(e2smcommoniesv1.NgEnbId_LongMacroNgEnbId{}),
	},
	"global_ngrannode_id": {
		1: reflect.TypeOf(e2smcommoniesv1.GlobalNgrannodeId_GNb{}),
		2: reflect.TypeOf(e2smcommoniesv1.GlobalNgrannodeId_NgENb{}),
	},
}
