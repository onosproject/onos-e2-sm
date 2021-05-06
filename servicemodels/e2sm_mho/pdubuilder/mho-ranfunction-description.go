// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
)

func CreateE2SmMhoRanfunctionDescriptionMsg(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string,
	retsl []*e2sm_mho.RicEventTriggerStyleList, rrsl []*e2sm_mho.RicReportStyleList) (*e2sm_mho.E2SmMhoRanfunctionDescription, error) {

	e2smMhoPdu := e2sm_mho.E2SmMhoRanfunctionDescription{
		RanFunctionName: &e2sm_mho.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,   //string
			RanFunctionE2SmOid:     ranFunctionE2SmOid,     //sting
			RanFunctionDescription: ranFunctionDescription, //string
			RanFunctionInstance:    -1,                     // Not valid value, indicates this item not present in message - handled later in CGo encoding
		},
		RicEventTriggerStyleList: retsl,
		RicReportStyleList:       rrsl,
	}

	if err := e2smMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &e2smMhoPdu, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) *e2sm_mho.RicEventTriggerStyleList {

	return &e2sm_mho.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2sm_mho.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2sm_mho.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2sm_mho.RicFormatType{
			Value: ricFormatType,
		},
	}
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, indHdrFormatType int32, indMsgFormatType int32) *e2sm_mho.RicReportStyleList {

	return &e2sm_mho.RicReportStyleList{
		RicReportStyleType: &e2sm_mho.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2sm_mho.RicStyleName{
			Value: ricStyleName,
		},
		RicIndicationHeaderFormatType: &e2sm_mho.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2sm_mho.RicFormatType{
			Value: indMsgFormatType,
		},
	}
}
