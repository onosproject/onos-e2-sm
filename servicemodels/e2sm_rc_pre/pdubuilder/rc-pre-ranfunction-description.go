// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
)

func CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string,
	retsl []*e2sm_rc_pre_v2.RicEventTriggerStyleList, rrsl []*e2sm_rc_pre_v2.RicReportStyleList) (*e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription, error) {

	e2smRcPrePdu := e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription{
		RanFunctionName: &e2sm_rc_pre_v2.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,   //string
			RanFunctionE2SmOid:     ranFunctionE2SmOid,     //sting
			RanFunctionDescription: ranFunctionDescription, //string
			RanFunctionInstance:    -1,                     // Not valid value, indicates this item not present in message - handled later in CGo encoding
		},
		//E2SmRcPreRanfunctionItem: &ranfunctionItem,
	}

	if retsl != nil || rrsl != nil {
		ranfunctionItem := e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription_E2SmRcPreRanfunctionItem001{
			//RicEventTriggerStyleList: retsl,
			//RicReportStyleList:       rrsl,
		}
		if retsl != nil {
			ranfunctionItem.RicEventTriggerStyleList = retsl
		}

		if rrsl != nil {
			ranfunctionItem.RicReportStyleList = rrsl
		}
		e2smRcPrePdu.E2SmRcPreRanfunctionItem = &ranfunctionItem
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) *e2sm_rc_pre_v2.RicEventTriggerStyleList {

	return &e2sm_rc_pre_v2.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2sm_rc_pre_v2.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2sm_rc_pre_v2.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2sm_rc_pre_v2.RicFormatType{
			Value: ricFormatType,
		},
	}
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, indHdrFormatType int32, indMsgFormatType int32) *e2sm_rc_pre_v2.RicReportStyleList {

	return &e2sm_rc_pre_v2.RicReportStyleList{
		RicReportStyleType: &e2sm_rc_pre_v2.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2sm_rc_pre_v2.RicStyleName{
			Value: ricStyleName,
		},
		RicIndicationHeaderFormatType: &e2sm_rc_pre_v2.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2sm_rc_pre_v2.RicFormatType{
			Value: indMsgFormatType,
		},
	}
}
