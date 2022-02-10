// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2sm_v2_ies

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
		PLmnidentity: &PlmnIdentity{
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
