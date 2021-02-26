// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementInfoItem() *e2sm_kpm_v2.MeasurementInfoItem {

	labelInfoList := &e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}

	labelInfoItem := &e2sm_kpm_v2.LabelInfoItem{
		MeasLabel: &e2sm_kpm_v2.MeasurementLabel{
			PlmnId: &e2sm_kpm_v2.PlmnIdentity{
				Value: []byte{0x1, 0x2, 0x3},
			},
			SliceId: &e2sm_kpm_v2.Snssai{
				SD:  21,
				SSt: 22,
			},
			FiveQi: &e2sm_kpm_v2.FiveQi{
				Value: 23,
			},
			QCi: &e2sm_kpm_v2.Qci{
				Value: 24,
			},
			QCimax: &e2sm_kpm_v2.Qci{
				Value: 30,
			},
			QCimin: &e2sm_kpm_v2.Qci{
				Value: 1,
			},
			ARpmax: &e2sm_kpm_v2.Arp{
				Value: 50,
			},
			ARpmin: &e2sm_kpm_v2.Arp{
				Value: 1,
			},
			BitrateRange:     25,
			LayerMuMimo:      1,
			SUm:              11,
			DistBinX:         123,
			DistBinY:         456,
			DistBinZ:         789,
			PreLabelOverride: 2,
			StartEndInd:      1,
		},
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	return &e2sm_kpm_v2.MeasurementInfoItem{
		MeasType: &e2sm_kpm_v2.MeasurementType{
			MeasurementType: &e2sm_kpm_v2.MeasurementType_MeasId{
				MeasId: &e2sm_kpm_v2.MeasurementTypeId{
					Value: 21,
				},
			},
		},
		LabelInfoList: labelInfoList,
	}
}

func Test_xerEncodeMeasurementInfoItem(t *testing.T) {

	mii := createMeasurementInfoItem()

	xer, err := xerEncodeMeasurementInfoItem(mii)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementInfoItem XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementInfoItem(t *testing.T) {

	mii := createMeasurementInfoItem()

	xer, err := xerEncodeMeasurementInfoItem(mii)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("MeasurementInfoItem XER\n%s", string(xer))

	result, err := xerDecodeMeasurementInfoItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeMeasurementInfoItem(t *testing.T) {

	mii := createMeasurementInfoItem()

	per, err := perEncodeMeasurementInfoItem(mii)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementInfoItem PER\n%s", string(per))
}

func Test_perDecodeMeasurementInfoItem(t *testing.T) {

	mii := createMeasurementInfoItem()

	per, err := perEncodeMeasurementInfoItem(mii)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("MeasurementInfoItem PER\n%s", string(per))

	result, err := perDecodeMeasurementInfoItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
