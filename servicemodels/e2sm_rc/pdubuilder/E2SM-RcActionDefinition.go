// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmRcActionDefinitionFormat1(rst int32, rpIDs []int64) (*e2smrcv1.E2SmRcActionDefinition, error) {

	rrl := make([]*e2smrcv1.E2SmRcActionDefinitionFormat1Item, 0)

	for _, v := range rpIDs {
		item := &e2smrcv1.E2SmRcActionDefinitionFormat1Item{
			RanParameterId: &e2smrcv1.RanparameterId{
				Value: v,
			},
		}
		rrl = append(rrl, item)
	}

	ad := &e2smrcv1.E2SmRcActionDefinition{
		RicStyleType: &e2smcommonies.RicStyleType{
			Value: rst,
		},
		RicActionDefinitionFormats: &e2smrcv1.RicActionDefinitionFormats{
			RicActionDefinitionFormats: &e2smrcv1.RicActionDefinitionFormats_ActionDefinitionFormat1{
				ActionDefinitionFormat1: &e2smrcv1.E2SmRcActionDefinitionFormat1{
					RanPToBeReportedList: rrl,
				},
			},
		},
	}

	return ad, nil
}

func CreateE2SmRcActionDefinitionFormat2(rst int32, rpcl []*e2smrcv1.E2SmRcActionDefinitionFormat2Item) (*e2smrcv1.E2SmRcActionDefinition, error) {

	ad := &e2smrcv1.E2SmRcActionDefinition{
		RicStyleType: &e2smcommonies.RicStyleType{
			Value: rst,
		},
		RicActionDefinitionFormats: &e2smrcv1.RicActionDefinitionFormats{
			RicActionDefinitionFormats: &e2smrcv1.RicActionDefinitionFormats_ActionDefinitionFormat2{
				ActionDefinitionFormat2: &e2smrcv1.E2SmRcActionDefinitionFormat2{
					RicPolicyConditionsList: rpcl,
				},
			},
		},
	}

	return ad, nil
}

func CreateE2SmRcActionDefinitionFormat3(rst int32, riiID int32, rpIDl []int64) (*e2smrcv1.E2SmRcActionDefinition, error) {

	riil := make([]*e2smrcv1.E2SmRcActionDefinitionFormat3Item, 0)

	for _, v := range rpIDl {
		item := &e2smrcv1.E2SmRcActionDefinitionFormat3Item{
			RanParameterId: &e2smrcv1.RanparameterId{
				Value: v,
			},
		}
		riil = append(riil, item)
	}

	ad := &e2smrcv1.E2SmRcActionDefinition{
		RicStyleType: &e2smcommonies.RicStyleType{
			Value: rst,
		},
		RicActionDefinitionFormats: &e2smrcv1.RicActionDefinitionFormats{
			RicActionDefinitionFormats: &e2smrcv1.RicActionDefinitionFormats_ActionDefinitionFormat3{
				ActionDefinitionFormat3: &e2smrcv1.E2SmRcActionDefinitionFormat3{
					RicInsertIndicationId: &e2smrcv1.RicInsertIndicationId{
						Value: riiID,
					},
					RanPInsertIndicationList: riil,
				},
			},
		},
	}

	return ad, nil
}

func CreateE2SmRcActionDefinitionFormat2Item(rcaID int32, rpl []*e2smrcv1.RicPolicyActionRanparameterItem) (*e2smrcv1.E2SmRcActionDefinitionFormat2Item, error) {

	rpa := &e2smrcv1.RicPolicyAction{
		RicPolicyActionId: &e2smrcv1.RicControlActionId{
			Value: rcaID,
		},
		RanParametersList: rpl,
	}

	item := &e2smrcv1.E2SmRcActionDefinitionFormat2Item{
		RicPolicyAction: rpa,
	}

	return item, nil
}

func CreateRicPolicyActionRanParameterItem(rpID int64, rpvt *e2smrcv1.RanparameterValueType) (*e2smrcv1.RicPolicyActionRanparameterItem, error) {

	item := &e2smrcv1.RicPolicyActionRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: rpID,
		},
		RanParameterValueType: rpvt,
	}

	return item, nil
}

func CreateRanparameterValueTypeChoiceElementTrue(rpv *e2smrcv1.RanparameterValue) (*e2smrcv1.RanparameterValueType, error) {

	return &e2smrcv1.RanparameterValueType{
		RanparameterValueType: &e2smrcv1.RanparameterValueType_RanPChoiceElementTrue{
			RanPChoiceElementTrue: &e2smrcv1.RanparameterValueTypeChoiceElementTrue{
				RanParameterValue: rpv,
			},
		},
	}, nil
}

func CreateRanparameterValueTypeChoiceElementFalse(rpv *e2smrcv1.RanparameterValue) (*e2smrcv1.RanparameterValueType, error) {

	return &e2smrcv1.RanparameterValueType{
		RanparameterValueType: &e2smrcv1.RanparameterValueType_RanPChoiceElementFalse{
			RanPChoiceElementFalse: &e2smrcv1.RanparameterValueTypeChoiceElementFalse{
				RanParameterValue: rpv,
			},
		},
	}, nil
}

func CreateRanparameterValueTypeChoiceStructure(rps *e2smrcv1.RanparameterStructure) (*e2smrcv1.RanparameterValueType, error) {

	return &e2smrcv1.RanparameterValueType{
		RanparameterValueType: &e2smrcv1.RanparameterValueType_RanPChoiceStructure{
			RanPChoiceStructure: &e2smrcv1.RanparameterValueTypeChoiceStructure{
				RanParameterStructure: rps,
			},
		},
	}, nil
}

func CreateRanparameterValueTypeChoiceList(rpl *e2smrcv1.RanparameterList) (*e2smrcv1.RanparameterValueType, error) {

	return &e2smrcv1.RanparameterValueType{
		RanparameterValueType: &e2smrcv1.RanparameterValueType_RanPChoiceList{
			RanPChoiceList: &e2smrcv1.RanparameterValueTypeChoiceList{
				RanParameterList: rpl,
			},
		},
	}, nil
}

func CreateRanparameterStructureItem(rpID int64, rpvt *e2smrcv1.RanparameterValueType) (*e2smrcv1.RanparameterStructureItem, error) {

	item := &e2smrcv1.RanparameterStructureItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: rpID,
		},
		RanParameterValueType: rpvt,
	}

	return item, nil
}

func CreateRanParameterStructure(srp []*e2smrcv1.RanparameterStructureItem) (*e2smrcv1.RanparameterStructure, error) {

	item := &e2smrcv1.RanparameterStructure{
		SequenceOfRanParameters: srp,
	}

	return item, nil
}

func CreateRanparameterValueBoolean(val bool) (*e2smrcv1.RanparameterValue, error) {

	return &e2smrcv1.RanparameterValue{
		RanparameterValue: &e2smrcv1.RanparameterValue_ValueBoolean{
			ValueBoolean: val,
		},
	}, nil
}

func CreateRanparameterValueInt(val int64) (*e2smrcv1.RanparameterValue, error) {

	return &e2smrcv1.RanparameterValue{
		RanparameterValue: &e2smrcv1.RanparameterValue_ValueInt{
			ValueInt: val,
		},
	}, nil
}

func CreateRanparameterValueReal(val float32) (*e2smrcv1.RanparameterValue, error) {

	return &e2smrcv1.RanparameterValue{
		RanparameterValue: &e2smrcv1.RanparameterValue_ValueReal{
			ValueReal: val,
		},
	}, nil
}

func CreateRanparameterValueBitS(val *asn1.BitString) (*e2smrcv1.RanparameterValue, error) {

	return &e2smrcv1.RanparameterValue{
		RanparameterValue: &e2smrcv1.RanparameterValue_ValueBitS{
			ValueBitS: val,
		},
	}, nil
}

func CreateRanparameterValueOctS(val []byte) (*e2smrcv1.RanparameterValue, error) {

	return &e2smrcv1.RanparameterValue{
		RanparameterValue: &e2smrcv1.RanparameterValue_ValueOctS{
			ValueOctS: val,
		},
	}, nil
}

func CreateRanparameterValuePrintableString(val string) (*e2smrcv1.RanparameterValue, error) {

	return &e2smrcv1.RanparameterValue{
		RanparameterValue: &e2smrcv1.RanparameterValue_ValuePrintableString{
			ValuePrintableString: val,
		},
	}, nil
}

func CreateRanparameterTestingItem(rpID int64, rpt *e2smrcv1.RanParameterType) (*e2smrcv1.RanparameterTestingItem, error) {

	item := &e2smrcv1.RanparameterTestingItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: rpID,
		},
		RanParameterType: rpt,
	}

	return item, nil
}

func CreateRanParameterTypeChoiceList(rptl *e2smrcv1.RanparameterTestingList) (*e2smrcv1.RanParameterType, error) {

	return &e2smrcv1.RanParameterType{
		RanParameterType: &e2smrcv1.RanParameterType_RanPChoiceList{
			RanPChoiceList: &e2smrcv1.RanparameterTestingItemChoiceList{
				RanParameterList: rptl,
			},
		},
	}, nil
}

func CreateRanParameterTypeChoiceStructure(rpts *e2smrcv1.RanparameterTestingStructure) (*e2smrcv1.RanParameterType, error) {

	return &e2smrcv1.RanParameterType{
		RanParameterType: &e2smrcv1.RanParameterType_RanPChoiceStructure{
			RanPChoiceStructure: &e2smrcv1.RanparameterTestingItemChoiceStructure{
				RanParameterStructure: rpts,
			},
		},
	}, nil
}

func CreateRanParameterTypeChoiceElementTrue(rpv *e2smrcv1.RanparameterValue) (*e2smrcv1.RanParameterType, error) {

	return &e2smrcv1.RanParameterType{
		RanParameterType: &e2smrcv1.RanParameterType_RanPChoiceElementTrue{
			RanPChoiceElementTrue: &e2smrcv1.RanparameterTestingItemChoiceElementTrue{
				RanParameterValue: rpv,
			},
		},
	}, nil
}

func CreateRanParameterTypeChoiceElementFalse(rptc *e2smrcv1.RanparameterTestingCondition) (*e2smrcv1.RanParameterType, error) {

	return &e2smrcv1.RanParameterType{
		RanParameterType: &e2smrcv1.RanParameterType_RanPChoiceElementFalse{
			RanPChoiceElementFalse: &e2smrcv1.RanparameterTestingItemChoiceElementFalse{
				RanParameterTestCondition: rptc,
			},
		},
	}, nil
}

func CreateRanparameterTestingConditionComparison(rpcc e2smrcv1.RanPChoiceComparison) (*e2smrcv1.RanparameterTestingCondition, error) {

	return &e2smrcv1.RanparameterTestingCondition{
		RanparameterTestingCondition: &e2smrcv1.RanparameterTestingCondition_RanPChoiceComparison{
			RanPChoiceComparison: rpcc,
		},
	}, nil
}

func CreateRanparameterTestingConditionPresence(rpcc e2smrcv1.RanPChoicePresence) (*e2smrcv1.RanparameterTestingCondition, error) {

	return &e2smrcv1.RanparameterTestingCondition{
		RanparameterTestingCondition: &e2smrcv1.RanparameterTestingCondition_RanPChoicePresence{
			RanPChoicePresence: rpcc,
		},
	}, nil
}

func CreateLogicalOrTrue() e2smrcv1.LogicalOr {
	return e2smrcv1.LogicalOr_LOGICAL_OR_TRUE
}

func CreateLogicalOrFalse() e2smrcv1.LogicalOr {
	return e2smrcv1.LogicalOr_LOGICAL_OR_FALSE
}

func CreateRanPChoiceComparisonEqual() e2smrcv1.RanPChoiceComparison {
	return e2smrcv1.RanPChoiceComparison_RAN_P_CHOICE_COMPARISON_EQUAL
}

func CreateRanPChoiceComparisonDifference() e2smrcv1.RanPChoiceComparison {
	return e2smrcv1.RanPChoiceComparison_RAN_P_CHOICE_COMPARISON_DIFFERENCE
}

func CreateRanPChoiceComparisonGreaterthan() e2smrcv1.RanPChoiceComparison {
	return e2smrcv1.RanPChoiceComparison_RAN_P_CHOICE_COMPARISON_GREATERTHAN
}

func CreateRanPChoiceComparisonLessthan() e2smrcv1.RanPChoiceComparison {
	return e2smrcv1.RanPChoiceComparison_RAN_P_CHOICE_COMPARISON_LESSTHAN
}

func CreateRanPChoiceComparisonContains() e2smrcv1.RanPChoiceComparison {
	return e2smrcv1.RanPChoiceComparison_RAN_P_CHOICE_COMPARISON_CONTAINS
}

func CreateRanPChoiceComparisonStartsWith() e2smrcv1.RanPChoiceComparison {
	return e2smrcv1.RanPChoiceComparison_RAN_P_CHOICE_COMPARISON_STARTS_WITH
}

func CreateRanPChoicePresencePresent() e2smrcv1.RanPChoicePresence {
	return e2smrcv1.RanPChoicePresence_RAN_P_CHOICE_PRESENCE_PRESENT
}

func CreateRanPChoicePresenceConfigured() e2smrcv1.RanPChoicePresence {
	return e2smrcv1.RanPChoicePresence_RAN_P_CHOICE_PRESENCE_CONFIGURED
}

func CreateRanPChoicePresenceRollover() e2smrcv1.RanPChoicePresence {
	return e2smrcv1.RanPChoicePresence_RAN_P_CHOICE_PRESENCE_ROLLOVER
}

func CreateRanPChoicePresenceNonzero() e2smrcv1.RanPChoicePresence {
	return e2smrcv1.RanPChoicePresence_RAN_P_CHOICE_PRESENCE_NONZERO
}
