// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string) (*e2sm_rc_pre_go.E2SmRcPreRanfunctionDescription, error) {

	e2smRcPrePdu := e2sm_rc_pre_go.E2SmRcPreRanfunctionDescription{
		RanFunctionName: &e2sm_rc_pre_go.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,   //string
			RanFunctionE2SmOid:     ranFunctionE2SmOid,     //sting
			RanFunctionDescription: ranFunctionDescription, //string
		},
		// it's not an optional structure, so the signature should be present
		E2SmRcPreRanfunctionItem: &e2sm_rc_pre_go.E2SmRcPreRanfunctionDescription_E2SmRcPreRanfunctionItem001{},
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) *e2sm_rc_pre_go.RicEventTriggerStyleList {

	return &e2sm_rc_pre_go.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2sm_rc_pre_go.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2sm_rc_pre_go.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2sm_rc_pre_go.RicFormatType{
			Value: ricFormatType,
		},
	}
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, indHdrFormatType int32, indMsgFormatType int32) *e2sm_rc_pre_go.RicReportStyleList {

	return &e2sm_rc_pre_go.RicReportStyleList{
		RicReportStyleType: &e2sm_rc_pre_go.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2sm_rc_pre_go.RicStyleName{
			Value: ricStyleName,
		},
		RicIndicationHeaderFormatType: &e2sm_rc_pre_go.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2sm_rc_pre_go.RicFormatType{
			Value: indMsgFormatType,
		},
	}
}
