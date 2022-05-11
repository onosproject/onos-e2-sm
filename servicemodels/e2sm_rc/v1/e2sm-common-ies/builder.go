// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smcommoniesv1

import (
e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
)


func (m *RanfunctionName) SetRanFunctionInstance(ranFunctionInstance int32) *RanfunctionName {
	m.RanFunctionInstance = &ranFunctionInstance
	return m
}

func (m *UeidGnb) SetGNbCuUeF1ApIDList(gNbCuUeF1ApIDList *e2smcommoniesv1.UeidGnbCuF1ApIdList) *UeidGnb {
	m.GNbCuUeF1ApIdList = gNbCuUeF1ApIDList
	return m
}

func (m *UeidGnb) SetGNbCuCpUeE1ApIDList(gNbCuCpUeE1ApIDList *e2smcommoniesv1.UeidGnbCuCpE1ApIdList) *UeidGnb {
	m.GNbCuCpUeE1ApIdList = gNbCuCpUeE1ApIDList
	return m
}

func (m *UeidGnb) SetRanUeID(ranUeID *e2smcommoniesv1.Ranueid) *UeidGnb {
	m.RanUeid = ranUeID
	return m
}

func (m *UeidGnb) SetMNgRanUeXnApID(mNgRanUeXnApID *e2smcommoniesv1.NgRannodeUexnApid) *UeidGnb {
	m.MNgRanUeXnApId = mNgRanUeXnApID
	return m
}

func (m *UeidGnb) SetGlobalGnbID(globalGnbID *e2smcommoniesv1.GlobalGnbId) *UeidGnb {
	m.GlobalGnbId = globalGnbID
	return m
}

func (m *UeidGnb) SetGlobalNgRannodeID(globalNgRannodeID *e2smcommoniesv1.GlobalNgrannodeId) *UeidGnb {
	m.GlobalNgRannodeId = globalNgRannodeID
	return m
}

func (m *UeidGnbDu) SetRanUeID(ranUeID *e2smcommoniesv1.Ranueid) *UeidGnbDu {
	m.RanUeid = ranUeID
	return m
}

func (m *UeidGnbCuUp) SetRanUeID(ranUeID *e2smcommoniesv1.Ranueid) *UeidGnbCuUp {
	m.RanUeid = ranUeID
	return m
}

func (m *UeidNgEnb) SetNgENbCuUeW1ApID(ngENbCuUeW1ApID *e2smcommoniesv1.NgenbCuUeW1ApId) *UeidNgEnb {
	m.NgENbCuUeW1ApId = ngENbCuUeW1ApID
	return m
}

func (m *UeidNgEnb) SetMNgRanUeXnApID(mNgRanUeXnApID *e2smcommoniesv1.NgRannodeUexnApid) *UeidNgEnb {
	m.MNgRanUeXnApId = mNgRanUeXnApID
	return m
}

func (m *UeidNgEnb) SetGlobalNgEnbID(globalNgEnbID *e2smcommoniesv1.GlobalNgEnbId) *UeidNgEnb {
	m.GlobalNgEnbId = globalNgEnbID
	return m
}

func (m *UeidNgEnb) SetGlobalNgRannodeID(globalNgRannodeID *e2smcommoniesv1.GlobalNgrannodeId) *UeidNgEnb {
	m.GlobalNgRannodeId = globalNgRannodeID
	return m
}

func (m *UeidEnGnb) SetMENbUeX2ApIDExtension(mENbUeX2ApIDExtension *e2smcommoniesv1.EnbUeX2ApIdExtension) *UeidEnGnb {
	m.MENbUeX2ApIdExtension = mENbUeX2ApIDExtension
	return m
}

func (m *UeidEnGnb) SetGNbCuUeF1ApID(gNbCuUeF1ApID *e2smcommoniesv1.GnbCuUeF1ApId) *UeidEnGnb {
	m.GNbCuUeF1ApId = gNbCuUeF1ApID
	return m
}

func (m *UeidEnGnb) SetGNbCuCpUeE1ApIDList(gNbCuCpUeE1ApIDList *e2smcommoniesv1.UeidGnbCuCpE1ApIdList) *UeidEnGnb {
	m.GNbCuCpUeE1ApIdList = gNbCuCpUeE1ApIDList
	return m
}

func (m *UeidEnGnb) SetRanUeID(ranUeID *e2smcommoniesv1.Ranueid) *UeidEnGnb {
	m.RanUeid = ranUeID
	return m
}

func (m *UeidEnb) SetMENbUeX2ApID(mENbUeX2ApID *e2smcommoniesv1.EnbUeX2ApId) *UeidEnb {
	m.MENbUeX2ApId = mENbUeX2ApID
	return m
}

func (m *UeidEnb) SetMENbUeX2ApIDExtension(mENbUeX2ApIDExtension *e2smcommoniesv1.EnbUeX2ApIdExtension) *UeidEnb {
	m.MENbUeX2ApIdExtension = mENbUeX2ApIDExtension
	return m
}

func (m *UeidEnb) SetGlobalEnbID(globalEnbID *e2smcommoniesv1.GlobalEnbId) *UeidEnb {
	m.GlobalEnbId = globalEnbID
	return m
}

func (m *SNssai) SetSD(sD *e2smcommoniesv1.Sd) *SNssai {
	m.SD = sD
	return m
}

func (m *NrfrequencyInfo) SetFrequencyShift7P5Khz(frequencyShift7P5Khz NrfrequencyShift7P5Khz) *NrfrequencyInfo {
	m.FrequencyShift7P5Khz = &frequencyShift7P5Khz
	return m
}
