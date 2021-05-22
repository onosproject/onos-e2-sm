// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createRicKpmnodeItem() (*e2sm_kpm_v2.RicKpmnodeItem, error) {

	res := e2sm_kpm_v2.RicKpmnodeItem{
		RicKpmnodeType: &e2sm_kpm_v2.GlobalKpmnodeId{
			GlobalKpmnodeId: &e2sm_kpm_v2.GlobalKpmnodeId_GNb{
				GNb: &e2sm_kpm_v2.GlobalKpmnodeGnbId{
					GlobalGNbId: &e2sm_kpm_v2.GlobalgNbId{
						GnbId: &e2sm_kpm_v2.GnbIdChoice{
							GnbIdChoice: &e2sm_kpm_v2.GnbIdChoice_GnbId{
								GnbId: &e2sm_kpm_v2.BitString{
									Value: 0x9bcd4,
									Len:   22,
								},
							},
						},
						PlmnId: &e2sm_kpm_v2.PlmnIdentity{
							Value: []byte("onf"),
						},
					},
					GNbCuUpId: &e2sm_kpm_v2.GnbCuUpId{
						Value: 21,
					},
					GNbDuId: &e2sm_kpm_v2.GnbDuId{
						Value: 22,
					},
				},
			},
		},
		CellMeasurementObjectList: nil,
	}

	//item := &e2sm_kpm_v2.CellMeasurementObjectItem{
	//	CellObjectId: &e2sm_kpm_v2.CellObjectId{
	//		Value: "123",
	//	},
	//	CellGlobalId: &e2sm_kpm_v2.CellGlobalId{
	//		CellGlobalId: &e2sm_kpm_v2.CellGlobalId_EUtraCgi{
	//			EUtraCgi: &e2sm_kpm_v2.Eutracgi{
	//				EUtracellIdentity: &e2sm_kpm_v2.EutracellIdentity{
	//					Value: &e2sm_kpm_v2.BitString{
	//						Value: 0x9bcd4,
	//						Len:   28,
	//					},
	//				},
	//				PLmnIdentity: &e2sm_kpm_v2.PlmnIdentity{
	//					Value: []byte("ONF"),
	//				},
	//			},
	//		},
	//	},
	//}
	//
	//res.CellMeasurementObjectList = append(res.CellMeasurementObjectList, item)

	if err := res.Validate(); err != nil {
		return nil, err
	}
	return &res, nil
}

func Test_xerEncodeRicKpmnodeItem(t *testing.T) {

	item, err := createRicKpmnodeItem()
	assert.NilError(t, err)

	xer, err := xerEncodeRicKpmnodeItem(item)
	assert.NilError(t, err)
	//assert.Equal(t, 964, len(xer))
	//assert.Equal(t, 455, len(xer)) // without CellMeasurementObjectList
	t.Logf("RicKpmNodeItem XER\n%s", string(xer))
}

func Test_xerDecodeRicKpmnnodeItem(t *testing.T) {

	item, err := createRicKpmnodeItem()
	assert.NilError(t, err)

	xer, err := xerEncodeRicKpmnodeItem(item)
	assert.NilError(t, err)
	//assert.Equal(t, 964, len(xer))
	//assert.Equal(t, 455, len(xer)) //without CellMeasurementObjectList
	t.Logf("RicKpmNodeItem XER\n%s", string(xer))

	result, err := xerDecodeRicKpmnodeItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicKpmNodeItem XER - decoded\n%s", result)
}

func Test_perEncodeRicKpmnnodeItem(t *testing.T) {

	item, err := createRicKpmnodeItem()
	assert.NilError(t, err)

	per, err := perEncodeRicKpmnodeItem(item)
	assert.NilError(t, err)
	//assert.Equal(t, 29, len(per))
	//assert.Equal(t, 13, len(per)) //without CellMeasurementObjectList
	t.Logf("RicKpmNodeItem PER\n%s", string(per))
}

func Test_perDecodeRicKpmnNodeItem(t *testing.T) {

	item, err := createRicKpmnodeItem()
	assert.NilError(t, err)

	per, err := perEncodeRicKpmnodeItem(item)
	assert.NilError(t, err)
	//assert.Equal(t, 29, len(per))
	//assert.Equal(t, 13, len(per)) //without CellMeasurementObjectList
	t.Logf("RicKpmNodeItem PER\n%s", string(per))

	result, err := perDecodeRicKpmnodeItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicKpmNodeItem PER - decode\n%s", result)
}
