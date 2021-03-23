// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"

	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
)

func DecodeE2SmKpmRanfunctionDescription(e2smKpmPdu *e2sm_kpm_v2.E2SmKpmRanfunctionDescription) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error) {

	if err := e2smKpmPdu.Validate(); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid E2SmKpmPdu %s", err.Error())
	}

	ranfunctionName := e2smKpmPdu.GetRanFunctionName()
	if ranfunctionName == nil {
		return nil, nil, nil, fmt.Errorf("error E2SmKpm does not have RanfunctionName")
	}
	ranFunctionName := types.RanfunctionNameDef{
		RanFunctionShortName:   types.ShortName(ranfunctionName.GetRanFunctionShortName()),
		RanFunctionE2SmOid:     types.OID(ranfunctionName.GetRanFunctionE2SmOid()),
		RanFunctionDescription: types.Description(ranfunctionName.RanFunctionDescription),
		RanFunctionInstance:    types.Instance(ranfunctionName.RanFunctionInstance),
	}

	ricEvenTriggerStyleList := make(types.RicEventTriggerList)
	ricEvenTriggerStyleIe := e2smKpmPdu.GetRicEventTriggerStyleList()
	if ricEvenTriggerStyleIe == nil {
		return &ranFunctionName, nil, nil, fmt.Errorf("error E2SnKpmPdu does not have RIC-EventTrigger list")
	}
	for _, rEtSIe := range ricEvenTriggerStyleIe {
		ricEvenTriggerStyleList[types.StyleType(rEtSIe.GetRicEventTriggerStyleType().GetValue())] = types.RicEventTriggerDef{
			RicEventStyleType:  types.StyleType(rEtSIe.GetRicEventTriggerStyleType().GetValue()),
			RicEventStyleName:  types.StyleName(rEtSIe.RicEventTriggerStyleName.GetValue()),
			RicEventFormatType: types.FormatType(rEtSIe.GetRicEventTriggerFormatType().GetValue()),
		}
	}

	ricReportStyleList := make(types.RicReportList)
	ricReportStyleIe := e2smKpmPdu.GetRicReportStyleList()
	if ricReportStyleIe == nil {
		return &ranFunctionName, &ricEvenTriggerStyleList, nil, fmt.Errorf("error E2SnKpmPdu does not have RIC-ReportStyle list")
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
