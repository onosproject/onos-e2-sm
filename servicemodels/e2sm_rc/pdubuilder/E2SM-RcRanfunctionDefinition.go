//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
)

func CreateE2SmRcRanfunctionDefinition(ranFunctionShortName string, ranFunctionOID string, ranFunctionDescription string) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,
			RanFunctionE2SmOid:     ranFunctionOID,
			RanFunctionDescription: ranFunctionDescription,
		},
	}

	return msg, nil
}

func CreateRanfunctionDefinitionEventTrigger(ricEventTriggerStyleList []*e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem) (*e2smrcv1.RanfunctionDefinitionEventTrigger, error) {

	msg := &e2smrcv1.RanfunctionDefinitionEventTrigger{
		RicEventTriggerStyleList: ricEventTriggerStyleList,
	}

	return msg, nil
}

func CreateRanfunctionDefinitionEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) (*e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem{
		RicEventTriggerStyleType: &e2smcommonies.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2smcommonies.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2smcommonies.RicFormatType{
			Value: ricFormatType,
		},
	}

	return msg, nil
}

func CreateL2ParametersRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.L2ParametersRanparameterItem, error) {

	msg := &e2smrcv1.L2ParametersRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateRanparameterDefinition(ranParameterDefinitionChoice *e2smrcv1.RanparameterDefinitionChoice) (*e2smrcv1.RanparameterDefinition, error) {

	msg := &e2smrcv1.RanparameterDefinition{
		RanParameterDefinitionChoice: ranParameterDefinitionChoice,
	}

	return msg, nil
}

func CreateRanparameterDefinitionChoiceList(choiceList []*e2smrcv1.RanparameterDefinitionChoiceListItem) (*e2smrcv1.RanparameterDefinitionChoice, error) {

	msg := &e2smrcv1.RanparameterDefinitionChoice{
		RanparameterDefinitionChoice: &e2smrcv1.RanparameterDefinitionChoice_ChoiceList{
			ChoiceList: &e2smrcv1.RanparameterDefinitionChoiceList{
				RanParameterList: choiceList,
			},
		},
	}

	return msg, nil
}

func CreateRanparameterDefinitionChoiceListItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.RanparameterDefinitionChoiceListItem, error) {

	msg := &e2smrcv1.RanparameterDefinitionChoiceListItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateRanparameterDefinitionChoiceStructure(ranParameterStructure []*e2smrcv1.RanparameterDefinitionChoiceStructureItem) (*e2smrcv1.RanparameterDefinitionChoice, error) {

	msg := &e2smrcv1.RanparameterDefinitionChoice{
		RanparameterDefinitionChoice: &e2smrcv1.RanparameterDefinitionChoice_ChoiceStructure{
			ChoiceStructure: &e2smrcv1.RanparameterDefinitionChoiceStructure{
				RanParameterStructure: ranParameterStructure,
			},
		},
	}

	return msg, nil
}

func CreateRanparameterDefinitionChoiceStructureItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.RanparameterDefinitionChoiceStructureItem, error) {

	msg := &e2smrcv1.RanparameterDefinitionChoiceStructureItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateRanfunctionDefinitionEventTriggerCallProcessItem(callProcessTypeID int32, callProcessTypeName string,
	callProcessBreakpointList []*e2smrcv1.RanfunctionDefinitionEventTriggerBreakpointItem) (*e2smrcv1.RanfunctionDefinitionEventTriggerCallProcessItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionEventTriggerCallProcessItem{
		CallProcessTypeId: &e2smrcv1.RicCallProcessTypeId{
			Value: callProcessTypeID,
		},
		CallProcessTypeName: &e2smrcv1.RicCallProcessTypeName{
			Value: callProcessTypeName,
		},
		CallProcessBreakpointsList: callProcessBreakpointList,
	}

	return msg, nil
}

func CreateRanfunctionDefinitionEventTriggerBreakpointItem(callProcessBreakpointID int32, callProcessBreakpointName string) (*e2smrcv1.RanfunctionDefinitionEventTriggerBreakpointItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionEventTriggerBreakpointItem{
		CallProcessBreakpointId: &e2smrcv1.RicCallProcessBreakpointId{
			Value: callProcessBreakpointID,
		},
		CallProcessBreakpointName: &e2smrcv1.RicCallProcessBreakpointName{
			Value: callProcessBreakpointName,
		},
	}

	return msg, nil
}

func CreateCallProcessBreakpointRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.CallProcessBreakpointRanparameterItem, error) {

	msg := &e2smrcv1.CallProcessBreakpointRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateUeidentificationRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.UeidentificationRanparameterItem, error) {

	msg := &e2smrcv1.UeidentificationRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateCellIdentificationRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.CellIdentificationRanparameterItem, error) {

	msg := &e2smrcv1.CellIdentificationRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateRanfunctionDefinitionReport(ricReportStyleList []*e2smrcv1.RanfunctionDefinitionReportItem) (*e2smrcv1.RanfunctionDefinitionReport, error) {

	msg := &e2smrcv1.RanfunctionDefinitionReport{
		RicReportStyleList: ricReportStyleList,
	}

	return msg, nil
}

func CreateRanfunctionDefinitionReportItem(ricReportStyleType int32, ricReportStyleName string, ricSupportedEventTriggerStyleType int32,
	ricReportActionFormatType int32, ricIndicationHeaderFormatType int32, ricIndicationMessageFormatType int32) (*e2smrcv1.RanfunctionDefinitionReportItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionReportItem{
		RicReportStyleType: &e2smcommonies.RicStyleType{
			Value: ricReportStyleType,
		},
		RicReportStyleName: &e2smcommonies.RicStyleName{
			Value: ricReportStyleName,
		},
		RicSupportedEventTriggerStyleType: &e2smcommonies.RicStyleType{
			Value: ricSupportedEventTriggerStyleType,
		},
		RicReportActionFormatType: &e2smcommonies.RicFormatType{
			Value: ricReportActionFormatType,
		},
		RicIndicationHeaderFormatType: &e2smcommonies.RicFormatType{
			Value: ricIndicationHeaderFormatType,
		},
		RicIndicationMessageFormatType: &e2smcommonies.RicFormatType{
			Value: ricIndicationMessageFormatType,
		},
	}

	return msg, nil
}

func CreateReportRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.ReportRanparameterItem, error) {

	msg := &e2smrcv1.ReportRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateRanfunctionDefinitionInsert(ricInsertStyleList []*e2smrcv1.RanfunctionDefinitionInsertItem) (*e2smrcv1.RanfunctionDefinitionInsert, error) {

	msg := &e2smrcv1.RanfunctionDefinitionInsert{
		RicInsertStyleList: ricInsertStyleList,
	}

	return msg, nil
}

func CreateRanfunctionDefinitionInsertItem(ricInsertStyleType int32, ricInsertStyleName string, ricSupportedEventTriggerStyleType int32,
	ricActionDefinitionFormatType int32, ricIndicationHeaderFormatType int32, ricIndicationMessageFormatType int32,
	ricCallProcessIDformatType int32) (*e2smrcv1.RanfunctionDefinitionInsertItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionInsertItem{
		RicInsertStyleType: &e2smcommonies.RicStyleType{
			Value: ricInsertStyleType,
		},
		RicInsertStyleName: &e2smcommonies.RicStyleName{
			Value: ricInsertStyleName,
		},
		RicSupportedEventTriggerStyleType: &e2smcommonies.RicStyleType{
			Value: ricSupportedEventTriggerStyleType,
		},
		RicActionDefinitionFormatType: &e2smcommonies.RicFormatType{
			Value: ricActionDefinitionFormatType,
		},
		RicIndicationHeaderFormatType: &e2smcommonies.RicFormatType{
			Value: ricIndicationHeaderFormatType,
		},
		RicIndicationMessageFormatType: &e2smcommonies.RicFormatType{
			Value: ricIndicationMessageFormatType,
		},
		RicCallProcessIdformatType: &e2smcommonies.RicFormatType{
			Value: ricCallProcessIDformatType,
		},
	}

	return msg, nil
}

func CreateRanfunctionDefinitionInsertIndicationItem(ricInsertIndicationID int32, ricInsertIndicationName string) (*e2smrcv1.RanfunctionDefinitionInsertIndicationItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionInsertIndicationItem{
		RicInsertIndicationId: &e2smrcv1.RicInsertIndicationId{
			Value: ricInsertIndicationID,
		},
		RicInsertIndicationName: &e2smrcv1.RicInsertIndicationName{
			Value: ricInsertIndicationName,
		},
	}

	return msg, nil
}

func CreateInsertIndicationRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.InsertIndicationRanparameterItem, error) {

	msg := &e2smrcv1.InsertIndicationRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateRanfunctionDefinitionControl(ricControlStyleList []*e2smrcv1.RanfunctionDefinitionControlItem) (*e2smrcv1.RanfunctionDefinitionControl, error) {

	msg := &e2smrcv1.RanfunctionDefinitionControl{
		RicControlStyleList: ricControlStyleList,
	}

	return msg, nil
}

func CreateRanfunctionDefinitionControlItem(ricControlStyleType int32, ricControlStyleName string, ricControlHeaderFormatType int32,
	ricControlMessageFormatType int32, ricControlOutcomeFormatType int32) (*e2smrcv1.RanfunctionDefinitionControlItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionControlItem{
		RicControlStyleType: &e2smcommonies.RicStyleType{
			Value: ricControlStyleType,
		},
		RicControlStyleName: &e2smcommonies.RicStyleName{
			Value: ricControlStyleName,
		},
		RicControlHeaderFormatType: &e2smcommonies.RicFormatType{
			Value: ricControlHeaderFormatType,
		},
		RicControlMessageFormatType: &e2smcommonies.RicFormatType{
			Value: ricControlMessageFormatType,
		},
		RicControlOutcomeFormatType: &e2smcommonies.RicFormatType{
			Value: ricControlOutcomeFormatType,
		},
	}

	return msg, nil
}

func CreateRanfunctionDefinitionControlActionItem(ricControlActionID int32, ricControlActionName string) (*e2smrcv1.RanfunctionDefinitionControlActionItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionControlActionItem{
		RicControlActionId: &e2smrcv1.RicControlActionId{
			Value: ricControlActionID,
		},
		RicControlActionName: &e2smrcv1.RicControlActionName{
			Value: ricControlActionName,
		},
	}

	return msg, nil
}

func CreateControlActionRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.ControlActionRanparameterItem, error) {

	msg := &e2smrcv1.ControlActionRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateControlOutcomeRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.ControlOutcomeRanparameterItem, error) {

	msg := &e2smrcv1.ControlOutcomeRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreateRanfunctionDefinitionPolicy(ricPolicyStyleList []*e2smrcv1.RanfunctionDefinitionPolicyItem) (*e2smrcv1.RanfunctionDefinitionPolicy, error) {

	msg := &e2smrcv1.RanfunctionDefinitionPolicy{
		RicPolicyStyleList: ricPolicyStyleList,
	}

	return msg, nil
}

func CreateRanfunctionDefinitionPolicyItem(ricPolicyStyleType int32, ricPolicyStyleName string, ricSupportedEventTriggerStyleType int32) (*e2smrcv1.RanfunctionDefinitionPolicyItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionPolicyItem{
		RicPolicyStyleType: &e2smcommonies.RicStyleType{
			Value: ricPolicyStyleType,
		},
		RicPolicyStyleName: &e2smcommonies.RicStyleName{
			Value: ricPolicyStyleName,
		},
		RicSupportedEventTriggerStyleType: &e2smcommonies.RicStyleType{
			Value: ricSupportedEventTriggerStyleType,
		},
	}

	return msg, nil
}

func CreateRanfunctionDefinitionPolicyActionItem(ricControlActionID int32, ricPolicyActionName string, ricActionDefinitionFormatType int32) (*e2smrcv1.RanfunctionDefinitionPolicyActionItem, error) {

	msg := &e2smrcv1.RanfunctionDefinitionPolicyActionItem{
		RicPolicyActionId: &e2smrcv1.RicControlActionId{
			Value: ricControlActionID,
		},
		RicPolicyActionName: &e2smrcv1.RicControlActionName{
			Value: ricPolicyActionName,
		},
		RicActionDefinitionFormatType: &e2smcommonies.RicFormatType{
			Value: ricActionDefinitionFormatType,
		},
	}

	return msg, nil
}

func CreatePolicyActionRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.PolicyActionRanparameterItem, error) {

	msg := &e2smrcv1.PolicyActionRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}

func CreatePolicyConditionRanparameterItem(ranParameterID int64, ranParameterName string) (*e2smrcv1.PolicyConditionRanparameterItem, error) {

	msg := &e2smrcv1.PolicyConditionRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: ranParameterName,
		},
	}

	return msg, nil
}
