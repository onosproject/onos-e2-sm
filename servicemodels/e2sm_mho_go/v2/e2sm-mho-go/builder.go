// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2sm_mho_go

import e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"

func (ed *E2SmMhoEventTriggerDefinitionFormat1) SetReportingPeriodInMs(rp int32) *E2SmMhoEventTriggerDefinitionFormat1 {
	ed.ReportingPeriodMs = &rp
	return ed
}

func (ed *E2SmMhoControlHeaderFormat1) SetRicControlMessagePriority(cmp int32) *E2SmMhoControlHeaderFormat1 {
	ed.RicControlMessagePriority = &RicControlMessagePriority{
		Value: cmp,
	}
	return ed
}

func (rfd *E2SmMhoRanfunctionDescription) SetRicEventTriggerStyleList(retsl []*RicEventTriggerStyleList) *E2SmMhoRanfunctionDescription {
	rfd.GetE2SmMhoRanfunctionItem().RicEventTriggerStyleList = make([]*RicEventTriggerStyleList, 0)
	rfd.GetE2SmMhoRanfunctionItem().RicEventTriggerStyleList = retsl
	return rfd
}

func (rfd *E2SmMhoRanfunctionDescription) SetRicReportStyleList(rrsl []*RicReportStyleList) *E2SmMhoRanfunctionDescription {
	rfd.GetE2SmMhoRanfunctionItem().RicReportStyleList = make([]*RicReportStyleList, 0)
	rfd.GetE2SmMhoRanfunctionItem().RicReportStyleList = rrsl
	return rfd
}

func (mri *E2SmMhoMeasurementReportItem) SetFiveQi(fiveQI int32) *E2SmMhoMeasurementReportItem {
	mri.FiveQi = &e2sm_v2_ies.FiveQi{
		Value: fiveQI,
	}
	return mri
}
