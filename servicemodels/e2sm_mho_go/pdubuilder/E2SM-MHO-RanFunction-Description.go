// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
)

func CreateE2SmMhoRanfunctionDescriptionMsg(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string) (*e2sm_mho_go.E2SmMhoRanfunctionDescription, error) {

	e2smMhoPdu := e2sm_mho_go.E2SmMhoRanfunctionDescription{
		RanFunctionName: &e2sm_mho_go.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,   //string
			RanFunctionE2SmOid:     ranFunctionE2SmOid,     //sting
			RanFunctionDescription: ranFunctionDescription, //string
		},
		E2SmMhoRanfunctionItem: &e2sm_mho_go.E2SmMhoRanfunctionDescription_E2SmMhoRanfunctionItem001{},
	}

	//if err := e2smMhoPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	//}
	return &e2smMhoPdu, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) *e2sm_mho_go.RicEventTriggerStyleList {

	return &e2sm_mho_go.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2sm_mho_go.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2sm_mho_go.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2sm_mho_go.RicFormatType{
			Value: ricFormatType,
		},
	}
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, indHdrFormatType int32, indMsgFormatType int32) *e2sm_mho_go.RicReportStyleList {

	return &e2sm_mho_go.RicReportStyleList{
		RicReportStyleType: &e2sm_mho_go.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2sm_mho_go.RicStyleName{
			Value: ricStyleName,
		},
		RicIndicationHeaderFormatType: &e2sm_mho_go.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2sm_mho_go.RicFormatType{
			Value: indMsgFormatType,
		},
	}
}
