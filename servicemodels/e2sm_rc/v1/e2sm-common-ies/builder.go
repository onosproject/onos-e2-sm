// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smcommonies

import "github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

func (m *RanfunctionName) SetRanFunctionInstance(rfi int32) *RanfunctionName {
	m.RanFunctionInstance = &rfi
	return m
}

func (m *UeidGnb) SetGNbCuUeF1ApIDList(list []int64) *UeidGnb {
	m.GNbCuUeF1ApIdList = &UeidGnbCuF1ApIdList{
		Value: make([]*UeidGnbCuCpF1ApIdItem, 0),
	}

	for _, val := range list {
		item := &UeidGnbCuCpF1ApIdItem{
			GNbCuUeF1ApId: &GnbCuUeF1ApId{
				Value: val,
			},
		}
		m.GNbCuUeF1ApIdList.Value = append(m.GNbCuUeF1ApIdList.Value, item)
	}

	return m
}

func (m *UeidGnb) SetGNbCuCpUeE1ApIDList(list []int64) *UeidGnb {
	m.GNbCuCpUeE1ApIdList = &UeidGnbCuCpE1ApIdList{
		Value: make([]*UeidGnbCuCpE1ApIdItem, 0),
	}

	for _, val := range list {
		item := &UeidGnbCuCpE1ApIdItem{
			GNbCuCpUeE1ApId: &GnbCuCpUeE1ApId{
				Value: val,
			},
		}
		m.GNbCuCpUeE1ApIdList.Value = append(m.GNbCuCpUeE1ApIdList.Value, item)
	}

	return m
}

func (m *UeidGnb) SetRanUeID(ranUeID []byte) *UeidGnb {
	m.RanUeid = &Ranueid{
		Value: ranUeID,
	}

	return m
}

func (m *UeidGnb) SetMNgRanUeXnApID(ngRanNodeID int64) *UeidGnb {
	m.MNgRanUeXnApId = &NgRannodeUexnApid{
		Value: ngRanNodeID,
	}

	return m
}

func (m *UeidGnb) SetGlobalGnbID(plmnID []byte, gnbID *asn1.BitString) *UeidGnb {
	m.GlobalGnbId = &GlobalGnbId{
		PLmnidentity: &Plmnidentity{
			Value: plmnID,
		},
		GNbId: &GnbId{
			GnbId: &GnbId_GNbId{
				GNbId: gnbID,
			},
		},
	}

	return m
}

func (m *UeidGnbDu) SetRanUeID(ranUeID []byte) *UeidGnbDu {
	m.RanUeid = &Ranueid{
		Value: ranUeID,
	}

	return m
}

func (m *UeidGnbCuUp) SetRanUeID(ranUeID []byte) *UeidGnbCuUp {
	m.RanUeid = &Ranueid{
		Value: ranUeID,
	}

	return m
}

func (m *UeidNgEnb) SetNgENbCuUeW1ApID(ngENbCuUeW1ApID int64) *UeidNgEnb {
	m.NgENbCuUeW1ApId = &NgenbCuUeW1ApId{
		Value: ngENbCuUeW1ApID,
	}

	return m
}

func (m *UeidNgEnb) SetMNgRanUeXnApID(mNgRanUeXnApId int64) *UeidNgEnb {
	m.MNgRanUeXnApId = &NgRannodeUexnApid{
		Value: mNgRanUeXnApId,
	}

	return m
}

func (m *UeidNgEnb) SetGlobalNgEnbID(plmnID []byte, ngEnbID *NgEnbId) *UeidNgEnb {
	m.GlobalNgEnbId = &GlobalNgEnbId{
		PLmnidentity: &Plmnidentity{
			Value: plmnID,
		},
		NgEnbId: ngEnbID,
	}

	return m
}

func (m *UeidNgEnb) SetGlobalNgRannodeID(globalNgRannodeID *GlobalNgrannodeId) *UeidNgEnb {
	m.GlobalNgRannodeId = globalNgRannodeID

	return m
}

func (m *UeidEnGnb) SetMENbUeX2ApIDExtension(mENbUeX2ApIDExtension int32) *UeidEnGnb {
	m.MENbUeX2ApIdExtension = &EnbUeX2ApIdExtension{
		Value: mENbUeX2ApIDExtension,
	}

	return m
}

func (m *UeidEnGnb) SetGNbCuUeF1ApID(gNbCuUeF1ApID int64) *UeidEnGnb {
	m.GNbCuUeF1ApId = &GnbCuUeF1ApId{
		Value: gNbCuUeF1ApID,
	}

	return m
}

func (m *UeidEnGnb) SetGNbCuCpUeE1ApIDList(gNbCuCpUeE1ApIDList *UeidGnbCuCpE1ApIdList) *UeidEnGnb {
	m.GNbCuCpUeE1ApIdList = gNbCuCpUeE1ApIDList

	return m
}

func (m *UeidEnGnb) SetRanUeID(ranUeID []byte) *UeidEnGnb {
	m.RanUeid = &Ranueid{
		Value: ranUeID,
	}

	return m
}

func (m *UeidEnb) SetMENbUeX2ApID(mENbUeX2ApID int32) *UeidEnb {
	m.MENbUeX2ApId = &EnbUeX2ApId{
		Value: mENbUeX2ApID,
	}

	return m
}

func (m *UeidEnb) SetMENbUeX2ApIDExtension(mENbUeX2ApID int32) *UeidEnb {
	m.MENbUeX2ApIdExtension = &EnbUeX2ApIdExtension{
		Value: mENbUeX2ApID,
	}

	return m
}

func (m *UeidEnb) SetGlobalEnbID(plmnID []byte, enbID *EnbId) *UeidEnb {
	m.GlobalEnbId = &GlobalEnbId{
		PLmnidentity: &Plmnidentity{
			Value: plmnID,
		},
		ENbId: enbID,
	}

	return m
}

func (m *SNssai) SetSD(sd []byte) *SNssai {
	m.SD = &Sd{
		Value: sd,
	}

	return m
}

func (m *NrfrequencyInfo) SetFrequencyShift7P5Khz(fShift NrfrequencyShift7P5Khz) *NrfrequencyInfo {
	m.FrequencyShift7P5Khz = &fShift

	return m
}
