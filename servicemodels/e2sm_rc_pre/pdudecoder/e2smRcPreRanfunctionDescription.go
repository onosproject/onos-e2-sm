// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"

	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
)

func DecodeE2SmRcPreRanfunctionDescription(e2smRcPrePdu *e2sm_rc_pre_ies.E2SmRcPreRanfunctionDescription) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error) {

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid E2SmRcPrePdu %s", err.Error())
	}

	ranfunctionName := e2smRcPrePdu.GetRanFunctionName()
	if ranfunctionName == nil {
		return nil, nil, nil, fmt.Errorf("error E2SmRcPre does not have RanfunctionName")
	}
	ranFunctionName := types.RanfunctionNameDef{
		RanFunctionShortName:   types.ShortName(ranfunctionName.GetRanFunctionShortName()),
		RanFunctionE2SmOid:     types.OID(ranfunctionName.GetRanFunctionE2SmOid()),
		RanFunctionDescription: types.Description(ranfunctionName.RanFunctionDescription),
		RanFunctionInstance:    types.Instance(ranfunctionName.RanFunctionInstance),
	}

	ricEvenTriggerStyleList := make(types.RicEventTriggerList)
	ricEvenTriggerStyleIe := e2smRcPrePdu.GetE2SmRcPreRanfunctionItem().GetRicEventTriggerStyleList()
	if ricEvenTriggerStyleIe == nil {
		return &ranFunctionName, nil, nil, fmt.Errorf("error E2SnRcPrePdu does not have RIC-EventTrigger list")
	}
	for _, rEtSIe := range ricEvenTriggerStyleIe {
		ricEvenTriggerStyleList[types.StyleType(rEtSIe.GetRicEventTriggerStyleType().GetValue())] = types.RicEventTriggerDef{
			RicEventStyleType:  types.StyleType(rEtSIe.GetRicEventTriggerStyleType().GetValue()),
			RicEventStyleName:  types.StyleName(rEtSIe.RicEventTriggerStyleName.GetValue()),
			RicEventFormatType: types.FormatType(rEtSIe.GetRicEventTriggerFormatType().GetValue()),
		}
	}

	ricReportStyleList := make(types.RicReportList)
	ricReportStyleIe := e2smRcPrePdu.GetE2SmRcPreRanfunctionItem().GetRicReportStyleList()
	if ricReportStyleIe == nil {
		return &ranFunctionName, &ricEvenTriggerStyleList, nil, fmt.Errorf("error E2SnRcPrePdu does not have RIC-ReportStyle list")
	}
	for _, rRsIe := range ricReportStyleIe {
		ricReportStyleList[types.StyleType(rRsIe.GetRicReportStyleType().GetValue())] = types.RicReportStyleDef{
			RicReportStyleType:             types.StyleType(rRsIe.GetRicReportStyleType().GetValue()),
			RicReportStyleName:             types.StyleName(rRsIe.GetRicReportStyleName().GetValue()),
			RicIndicationHeaderFormatType:  types.FormatType(rRsIe.GetRicIndicationHeaderFormatType().GetValue()),
			RicIndicationMessageFormatType: types.FormatType(rRsIe.GetRicIndicationMessageFormatType().GetValue()),
		}
	}

	return &ranFunctionName, &ricEvenTriggerStyleList, &ricReportStyleList, nil
}
