// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2sm_rc_pre_v2_go

func (ed *E2SmRcPreEventTriggerDefinitionFormat1) SetReportingPeriodInMs(rp int64) *E2SmRcPreEventTriggerDefinitionFormat1 {
	ed.ReportingPeriodMs = &rp
	return ed
}

func (ihf1 *E2SmRcPreIndicationHeaderFormat1) SetCGI(cgi *CellGlobalId) *E2SmRcPreIndicationHeaderFormat1 {
	ihf1.Cgi = cgi
	return ihf1
}

func (chf1 *E2SmRcPreControlHeaderFormat1) SetCGI(cgi *CellGlobalId) *E2SmRcPreControlHeaderFormat1 {
	chf1.Cgi = cgi
	return chf1
}

func (chf1 *E2SmRcPreControlHeaderFormat1) SetRicControlMessagePriority(rcmp int32) *E2SmRcPreControlHeaderFormat1 {
	chf1.RicControlMessagePriority = &RicControlMessagePriority{
		Value: rcmp,
	}
	return chf1
}

func (cof1 *E2SmRcPreControlOutcomeFormat1) SetOutcomeElementList(oel []*RanparameterItem) *E2SmRcPreControlOutcomeFormat1 {
	cof1.OutcomeElementList = oel
	return cof1
}

func (rfd *E2SmRcPreRanfunctionDescription) SetRicEventTriggerStyleList(retsl []*RicEventTriggerStyleList) *E2SmRcPreRanfunctionDescription {
	rfd.GetE2SmRcPreRanfunctionItem().RicEventTriggerStyleList = make([]*RicEventTriggerStyleList, 0)
	rfd.GetE2SmRcPreRanfunctionItem().RicEventTriggerStyleList = retsl
	return rfd
}

func (rfd *E2SmRcPreRanfunctionDescription) SetRicReportStyleList(rrsl []*RicReportStyleList) *E2SmRcPreRanfunctionDescription {
	rfd.GetE2SmRcPreRanfunctionItem().RicReportStyleList = make([]*RicReportStyleList, 0)
	rfd.GetE2SmRcPreRanfunctionItem().RicReportStyleList = rrsl
	return rfd
}

func (rfn *RanfunctionName) SetRanFunctionInstance(rfi int32) *RanfunctionName {
	rfn.RanFunctionInstance = &rfi
	return rfn
}
