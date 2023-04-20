// SPDX-FileCopyrightText: 2023-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package e2sm_common_ies

import ()

func (m *RanfunctionName) SetRanFunctionInstance(ranFunctionInstance int32) *RanfunctionName {
	m.RanFunctionInstance = &ranFunctionInstance
	return m
}

func (m *UeidGNb) SetGNbCuUeF1ApIDList(gNbCuUeF1ApIDList *UeidGNbCUF1ApIDList) *UeidGNb {
	m.GNbCuUeF1ApIdList = gNbCuUeF1ApIDList
	return m
}

func (m *UeidGNb) SetGNbCuCpUeE1ApIDList(gNbCuCpUeE1ApIDList *UeidGNbCUCPE1ApIDList) *UeidGNb {
	m.GNbCuCpUeE1ApIdList = gNbCuCpUeE1ApIDList
	return m
}

func (m *UeidGNb) SetRanUeID(ranUeID *Ranueid) *UeidGNb {
	m.RanUeid = ranUeID
	return m
}

func (m *UeidGNb) SetMNgRanUeXnApID(mNgRanUeXnApID *NgRAnnodeUexnApid) *UeidGNb {
	m.MNgRanUeXnApId = mNgRanUeXnApID
	return m
}

func (m *UeidGNb) SetGlobalGnbID(globalGnbID *GlobalGnbID) *UeidGNb {
	m.GlobalGnbId = globalGnbID
	return m
}

func (m *UeidGNb) SetGlobalNgRannodeID(globalNgRannodeID *GlobalNgrannodeId) *UeidGNb {
	m.GlobalNgRannodeId = globalNgRannodeID
	return m
}

func (m *UeidGNbDU) SetRanUeID(ranUeID *Ranueid) *UeidGNbDU {
	m.RanUeid = ranUeID
	return m
}

func (m *UeidGNbCUUP) SetRanUeID(ranUeID *Ranueid) *UeidGNbCUUP {
	m.RanUeid = ranUeID
	return m
}

func (m *UeidNGENb) SetNgENbCuUeW1ApID(ngENbCuUeW1ApID *NgenbCUUEW1ApID) *UeidNGENb {
	m.NgENbCuUeW1ApId = ngENbCuUeW1ApID
	return m
}

func (m *UeidNGENb) SetMNgRanUeXnApID(mNgRanUeXnApID *NgRAnnodeUexnApid) *UeidNGENb {
	m.MNgRanUeXnApId = mNgRanUeXnApID
	return m
}

func (m *UeidNGENb) SetGlobalNgEnbID(globalNgEnbID *GlobalNgEnbID) *UeidNGENb {
	m.GlobalNgEnbId = globalNgEnbID
	return m
}

func (m *UeidNGENb) SetGlobalNgRannodeID(globalNgRannodeID *GlobalNgrannodeId) *UeidNGENb {
	m.GlobalNgRannodeId = globalNgRannodeID
	return m
}

func (m *UeidENGNb) SetMENbUeX2ApIDExtension(mENbUeX2ApIDExtension *EnbUEX2ApIDExtension) *UeidENGNb {
	m.MENbUeX2ApIdExtension = mENbUeX2ApIDExtension
	return m
}

func (m *UeidENGNb) SetGNbCuUeF1ApID(gNbCuUeF1ApID *GnbCUUEF1ApID) *UeidENGNb {
	m.GNbCuUeF1ApId = gNbCuUeF1ApID
	return m
}

func (m *UeidENGNb) SetGNbCuCpUeE1ApIDList(gNbCuCpUeE1ApIDList *UeidGNbCUCPE1ApIDList) *UeidENGNb {
	m.GNbCuCpUeE1ApIdList = gNbCuCpUeE1ApIDList
	return m
}

func (m *UeidENGNb) SetRanUeID(ranUeID *Ranueid) *UeidENGNb {
	m.RanUeid = ranUeID
	return m
}

func (m *UeidENb) SetMENbUeX2ApID(mENbUeX2ApID *EnbUEX2ApID) *UeidENb {
	m.MENbUeX2ApId = mENbUeX2ApID
	return m
}

func (m *UeidENb) SetMENbUeX2ApIDExtension(mENbUeX2ApIDExtension *EnbUEX2ApIDExtension) *UeidENb {
	m.MENbUeX2ApIdExtension = mENbUeX2ApIDExtension
	return m
}

func (m *UeidENb) SetGlobalEnbID(globalEnbID *GlobalEnbID) *UeidENb {
	m.GlobalEnbId = globalEnbID
	return m
}

func (m *SNSsai) SetSD(sD *Sd) *SNSsai {
	m.SD = sD
	return m
}

func (m *NrfrequencyInfo) SetFrequencyShift7P5Khz(frequencyShift7P5Khz NrfrequencyShift7P5Khz) *NrfrequencyInfo {
	m.FrequencyShift7P5Khz = &frequencyShift7P5Khz
	return m
}
