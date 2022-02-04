// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementInfoList() (*e2sm_kpm_v2.MeasurementInfoList, error) {

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
			StartEndInd:      e2sm_kpm_v2.StartEndInd_START_END_IND_END,
		},
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	list := &e2sm_kpm_v2.MeasurementInfoItem{
		MeasType: &e2sm_kpm_v2.MeasurementType{
			MeasurementType: &e2sm_kpm_v2.MeasurementType_MeasId{
				MeasId: &e2sm_kpm_v2.MeasurementTypeId{
					Value: 21,
				},
			},
		},
		LabelInfoList: labelInfoList,
	}

	res := &e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}
	res.Value = append(res.Value, list)

	if err := labelInfoList.Validate(); err != nil {
		return nil, err
	}
	return res, nil
}

func Test_xerEncodeMeasurementInfoList(t *testing.T) {

	mii, err := createMeasurementInfoList()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementInfoList(mii)
	assert.NilError(t, err)
	//assert.Equal(t, 1181, len(xer))
	t.Logf("MeasurementInfoList XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementInfoList(t *testing.T) {

	mii, err := createMeasurementInfoList()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementInfoList(mii)
	assert.NilError(t, err)
	//assert.Equal(t, 1181, len(xer))
	t.Logf("MeasurementInfoList XER\n%s", string(xer))

	result, err := xerDecodeMeasurementInfoList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementInfoList XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementInfoList(t *testing.T) {

	mii, err := createMeasurementInfoList()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementInfoList(mii)
	assert.NilError(t, err)
	//assert.Equal(t, 42, len(per))
	t.Logf("MeasurementInfoList PER\n%v", hex.Dump(per))
}

func Test_perDecodeMeasurementInfoList(t *testing.T) {

	mii, err := createMeasurementInfoList()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementInfoList(mii)
	assert.NilError(t, err)
	//assert.Equal(t, 42, len(per))
	t.Logf("MeasurementInfoList PER\n%v", hex.Dump(per))

	result, err := perDecodeMeasurementInfoList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MeasurementInfoList PER - decoded\n%s", result)
}
