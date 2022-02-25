// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string) (*e2smrcprev2.E2SmRcPreRanfunctionDescription, error) {

	e2smRcPrePdu := e2smrcprev2.E2SmRcPreRanfunctionDescription{
		RanFunctionName: &e2smrcprev2.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,   //string
			RanFunctionE2SmOid:     ranFunctionE2SmOid,     //sting
			RanFunctionDescription: ranFunctionDescription, //string
		},
		// it's not an optional structure, so the signature should be present
		E2SmRcPreRanfunctionItem: &e2smrcprev2.E2SmRcPreRanfunctionDescription_E2SmRcPreRanfunctionItem001{},
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) *e2smrcprev2.RicEventTriggerStyleList {

	return &e2smrcprev2.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2smrcprev2.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2smrcprev2.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2smrcprev2.RicFormatType{
			Value: ricFormatType,
		},
	}
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, indHdrFormatType int32, indMsgFormatType int32) *e2smrcprev2.RicReportStyleList {

	return &e2smrcprev2.RicReportStyleList{
		RicReportStyleType: &e2smrcprev2.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2smrcprev2.RicStyleName{
			Value: ricStyleName,
		},
		RicIndicationHeaderFormatType: &e2smrcprev2.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2smrcprev2.RicFormatType{
			Value: indMsgFormatType,
		},
	}
}
