// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementInfoItem() (*e2sm_kpm_v2.MeasurementInfoItem, error) {

	labelInfoList := &e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}

	labelInfoItem := &e2sm_kpm_v2.LabelInfoItem{
		MeasLabel: &e2sm_kpm_v2.MeasurementLabel{
			PlmnId: &e2sm_kpm_v2.PlmnIdentity{
				Value: []byte{0x1, 0x2, 0x3},
			},
			SliceId: &e2sm_kpm_v2.Snssai{
				SD:  []byte{0x01, 0x02, 0x03},
				SSt: []byte{0x01},
			},
			FiveQi: &e2sm_kpm_v2.FiveQi{
				Value: 23,
			},
			QFi: &e2sm_kpm_v2.Qfi{
				Value: 62,
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
				Value: 15,
			},
			ARpmin: &e2sm_kpm_v2.Arp{
				Value: 1,
			},
			BitrateRange:     25,
			LayerMuMimo:      1,
			SUm:              e2sm_kpm_v2.SUM_SUM_TRUE,
			DistBinX:         123,
			DistBinY:         456,
			DistBinZ:         789,
			PreLabelOverride: e2sm_kpm_v2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE,
			StartEndInd:      e2sm_kpm_v2.StartEndInd_START_END_IND_START,
		},
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	res := &e2sm_kpm_v2.MeasurementInfoItem{
		MeasType: &e2sm_kpm_v2.MeasurementType{
			MeasurementType: &e2sm_kpm_v2.MeasurementType_MeasId{
				MeasId: &e2sm_kpm_v2.MeasurementTypeId{
					Value: 21,
				},
			},
		},
		LabelInfoList: labelInfoList,
	}
	if err := res.Validate(); err != nil {
		return nil, err
	}

	return res, nil
}

func Test_xerEncodeMeasurementInfoItem(t *testing.T) {

	mii, err := createMeasurementInfoItem()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementInfoItem(mii)
	assert.NilError(t, err)
	//assert.Equal(t, 1014, len(xer))
	//assert.Equal(t, 104, len(xer)) // without LabelInfoList
	t.Logf("MeasurementInfoItem XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementInfoItem(t *testing.T) {

	mii, err := createMeasurementInfoItem()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementInfoItem(mii)
	assert.NilError(t, err)
	//assert.Equal(t, 1014, len(xer))
	//assert.Equal(t, 104, len(xer)) // without LabelInfoList
	t.Logf("MeasurementInfoItem XER\n%s", string(xer))

	result, err := xerDecodeMeasurementInfoItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementInfoItem XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementInfoItem(t *testing.T) {

	mii, err := createMeasurementInfoItem()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementInfoItem(mii)
	assert.NilError(t, err)
	//assert.Equal(t, 40, len(per))
	//assert.Equal(t, 3, len(per)) // without LabelInfoList
	t.Logf("MeasurementInfoItem PER\n%v", hex.Dump(per))
}

func Test_perDecodeMeasurementInfoItem(t *testing.T) {

	mii, err := createMeasurementInfoItem()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementInfoItem(mii)
	assert.NilError(t, err)
	//assert.Equal(t, 40, len(per))
	//assert.Equal(t, 3, len(per)) // without LabelInfoList
	t.Logf("MeasurementInfoItem PER\n%v", hex.Dump(per))

	result, err := perDecodeMeasurementInfoItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementInfoItem PER - decoded\n%s", result)
}
