// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smrsmies

func (m *E2SmRsmEventTriggerDefinitionFormat1) SetReportingPeriodMs(rp int32) *E2SmRsmEventTriggerDefinitionFormat1 {
	m.ReportingPeriodMs = &rp
	return m
}

func (m *E2SmRsmIndicationHeaderFormat1) SetColletStartTime(cst string) *E2SmRsmIndicationHeaderFormat1 {
	m.ColletStartTime = &cst
	return m
}

func (m *SliceConfig) SetSliceDescription(description string) *SliceConfig {
	m.SliceDescription = &description
	return m
}

func (m *SliceParameters) SetWeight(w int32) *SliceParameters {
	m.Weight = &w
	return m
}

func (m *SliceParameters) SetQoSLevel(qos int32) *SliceParameters {
	m.QosLevel = &qos
	return m
}

func (m *SliceParameters) SetScheduleInfo(schInfo *ScheduleConfig) *SliceParameters {
	m.ScheduleInfo = schInfo
	return m
}

func (m *ScheduleConfig) SetLinkAdaptation(la *LinkAdaptation) *ScheduleConfig {
	m.LinkAdaptation = la
	return m
}

func (m *ScheduleConfig) SetFeatures(f FeatureStatus) *ScheduleConfig {
	m.Features = &FeatureConfig{
		TtiBundling: &f,
	}
	return m
}

func (m *ScheduleConfig) SetCarrierAggregationLevelCap(calc CarrierAggregationLevelCap) *ScheduleConfig {
	m.CarrierAggregationCap = &calc
	return m
}

func (m *ScheduleConfig) SetUlPowerControl(ulpc *UlpowerControl) *ScheduleConfig {
	m.UlPowerControl = ulpc
	return m
}

func (m *LinkAdaptation) SetCqiCap(cqiCap int32) *LinkAdaptation {
	m.CqiCap = &cqiCap
	return m
}

func (m *LinkAdaptation) SetRiCap(riCap RiCap) *LinkAdaptation {
	m.RiCap = &riCap
	return m
}

func (m *LinkAdaptation) SetAggregationLevelCap(alc AggregationLevelCap) *LinkAdaptation {
	m.AggregationLevelCap = &alc
	return m
}

func (m *LinkAdaptation) SetTargetBlerDL(bler int32) *LinkAdaptation {
	m.TargetBlerDl = &bler
	return m
}

func (m *LinkAdaptation) SetTargetBlerUL(bler int32) *LinkAdaptation {
	m.TargetBlerUl = &bler
	return m
}

func (m *LinkAdaptation) SetMaxMcs(mcs int32) *LinkAdaptation {
	m.MaxMcs = &mcs
	return m
}

func (m *LinkAdaptation) SetMinMcs(mcs int32) *LinkAdaptation {
	m.MinMcs = &mcs
	return m
}

func (m *LinkAdaptation) SetTransmissionMode(trm TransmissionMode) *LinkAdaptation {
	m.TransmissionMode = &trm
	return m
}

func (m *LinkAdaptation) SetHarqRextCap(harq *HarqrextCap) *LinkAdaptation {
	m.HarqRetxCap = harq
	return m
}

func (m *HarqrextCap) SetDL(dl int64) *HarqrextCap {
	m.Dl = &dl
	return m
}

func (m *HarqrextCap) SetUL(ul int64) *HarqrextCap {
	m.Ul = &ul
	return m
}

func (m *FeatureConfig) SetTtiBundling(tti FeatureStatus) *FeatureConfig {
	m.TtiBundling = &tti
	return m
}

func (m *UlpowerControl) SetPUSCHtargetSNR(snr int64) *UlpowerControl {
	m.PuschTargetSnr = &snr
	return m
}

func (m *UlpowerControl) SetPUCCHtargetSNR(snr int64) *UlpowerControl {
	m.PucchTargetSnr = &snr
	return m
}

func (m *SliceAssociate) SetUplinkSliceID(sliceID int64) *SliceAssociate {
	m.UplinkSliceId = &SliceIdassoc{
		Value: sliceID,
	}
	return m
}
