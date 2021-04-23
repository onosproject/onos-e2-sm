// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
)

func CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string,
	ranFunctionInstance int32, ricEventStyleType int32, ricEventStyleName string, ricEventFormatType int32,
	ricReportStyleType int32, ricReportStyleName string, ricIndicationHeaderFormatType int32,
	ricIndicationMessageFormatType int32) (*e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription, error) {

	ranfunctionItem := e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription_E2SmRcPreRanfunctionItem001{
		RicEventTriggerStyleList: make([]*e2sm_rc_pre_v2.RicEventTriggerStyleList, 0),
		RicReportStyleList:       make([]*e2sm_rc_pre_v2.RicReportStyleList, 0),
	}

	ricEventTriggerStyleList := e2sm_rc_pre_v2.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2sm_rc_pre_v2.RicStyleType{
			Value: ricEventStyleType, //int32
		},
		RicEventTriggerStyleName: &e2sm_rc_pre_v2.RicStyleName{
			Value: ricEventStyleName, //string
		},
		RicEventTriggerFormatType: &e2sm_rc_pre_v2.RicFormatType{
			Value: ricEventFormatType, //int32
		},
	}
	ranfunctionItem.RicEventTriggerStyleList = append(ranfunctionItem.RicEventTriggerStyleList, &ricEventTriggerStyleList)

	ricReportStyleList := e2sm_rc_pre_v2.RicReportStyleList{
		RicReportStyleType: &e2sm_rc_pre_v2.RicStyleType{
			Value: ricReportStyleType, //int32
		},
		RicReportStyleName: &e2sm_rc_pre_v2.RicStyleName{
			Value: ricReportStyleName, //string
		},
		RicIndicationHeaderFormatType: &e2sm_rc_pre_v2.RicFormatType{
			Value: ricIndicationHeaderFormatType, //int32
		},
		RicIndicationMessageFormatType: &e2sm_rc_pre_v2.RicFormatType{
			Value: ricIndicationMessageFormatType, //int32
		},
	}
	ranfunctionItem.RicReportStyleList = append(ranfunctionItem.RicReportStyleList, &ricReportStyleList)

	e2smRcPrePdu := e2sm_rc_pre_v2.E2SmRcPreRanfunctionDescription{
		RanFunctionName: &e2sm_rc_pre_v2.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,   //string
			RanFunctionE2SmOid:     ranFunctionE2SmOid,     //sting
			RanFunctionDescription: ranFunctionDescription, //string
			RanFunctionInstance:    ranFunctionInstance,    //int32
		},
		E2SmRcPreRanfunctionItem: &ranfunctionItem,
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}
