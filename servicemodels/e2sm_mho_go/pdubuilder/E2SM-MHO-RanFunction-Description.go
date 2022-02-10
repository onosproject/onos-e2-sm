// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
)

func CreateE2SmMhoRanfunctionDescriptionMsg(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string) (*e2sm_mho_go.E2SmMhoRanfunctionDescription, error) {

	e2smMhoPdu := e2sm_mho_go.E2SmMhoRanfunctionDescription{
		RanFunctionName: &e2sm_v2_ies.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,   //string
			RanFunctionE2SmOid:     ranFunctionE2SmOid,     //sting
			RanFunctionDescription: ranFunctionDescription, //string
		},
		E2SmMhoRanfunctionItem: &e2sm_mho_go.E2SmMhoRanfunctionDescription_E2SmMhoRanfunctionItem001{},
	}

	if err := e2smMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("CreateE2SmMhoRanfunctionDescriptionMsg(): error validating E2SmPDU %s", err.Error())
	}
	return &e2smMhoPdu, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) (*e2sm_mho_go.RicEventTriggerStyleList, error) {

	res := &e2sm_mho_go.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2sm_v2_ies.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2sm_v2_ies.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2sm_v2_ies.RicFormatType{
			Value: ricFormatType,
		},
	}

	if err := res.Validate(); err != nil {
		return nil, fmt.Errorf("CreateRicEventTriggerStyleItem(): error validationg E2SmPDU %s", err)
	}

	return res, nil
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, indHdrFormatType int32, indMsgFormatType int32) (*e2sm_mho_go.RicReportStyleList, error) {

	res := &e2sm_mho_go.RicReportStyleList{
		RicReportStyleType: &e2sm_v2_ies.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2sm_v2_ies.RicStyleName{
			Value: ricStyleName,
		},
		RicIndicationHeaderFormatType: &e2sm_v2_ies.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2sm_v2_ies.RicFormatType{
			Value: indMsgFormatType,
		},
	}

	if err := res.Validate(); err != nil {
		return nil, fmt.Errorf("CreateRicReportStyleItem(): error validationg E2SmPDU %s", err)
	}

	return res, nil
}
