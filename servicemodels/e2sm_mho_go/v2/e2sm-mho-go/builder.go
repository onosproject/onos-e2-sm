// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2sm_mho_go

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

func (rfn *RanfunctionName) SetRanFunctionInstance(rfi int32) *RanfunctionName {
	rfn.RanFunctionInstance = &rfi
	return rfn
}

func (mri *E2SmMhoMeasurementReportItem) SetFiveQi(fiveQI int32) *E2SmMhoMeasurementReportItem {
	mri.FiveQi = &FiveQi{
		Value: fiveQI,
	}
	return mri
}
