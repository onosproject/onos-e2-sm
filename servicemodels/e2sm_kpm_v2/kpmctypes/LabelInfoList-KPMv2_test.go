// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createLabelInfoList() (*e2sm_kpm_v2.LabelInfoList, error) {

	labelInfoList := &e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}

	item := &e2sm_kpm_v2.LabelInfoItem{
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
	labelInfoList.Value = append(labelInfoList.Value, item)
	labelInfoList.Value = append(labelInfoList.Value, item)

	if err := labelInfoList.Validate(); err != nil {
		return nil, fmt.Errorf("error validating labelInfoList %s", err.Error())
	}
	return labelInfoList, nil
}

func Test_xerEncodeLabelInfoList(t *testing.T) {

	lil, err := createLabelInfoList()
	assert.NilError(t, err)

	xer, err := xerEncodeLabelInfoList(lil)
	assert.NilError(t, err)
	//assert.Equal(t, 1575, len(xer))
	t.Logf("LabelInfoList XER\n%s", string(xer))
}

func Test_xerDecodeLabelInfoList(t *testing.T) {

	lil, err := createLabelInfoList()
	assert.NilError(t, err)

	xer, err := xerEncodeLabelInfoList(lil)
	assert.NilError(t, err)
	//assert.Equal(t, 1575, len(xer))
	t.Logf("LabelInfoList XER\n%s", string(xer))

	result, err := xerDecodeLabelInfoList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("LabelInfoList XER - decoded\n%s", result)
}

func Test_perEncodeLabelInfoList(t *testing.T) {

	lil, err := createLabelInfoList()
	assert.NilError(t, err)

	per, err := perEncodeLabelInfoList(lil)
	assert.NilError(t, err)
	//assert.Equal(t, 72, len(per))
	t.Logf("LabelInfoList PER\n%v", hex.Dump(per))
}

func Test_perDecodeLabelInfoList(t *testing.T) {

	lil, err := createLabelInfoList()
	assert.NilError(t, err)

	per, err := perEncodeLabelInfoList(lil)
	assert.NilError(t, err)
	//assert.Equal(t, 72, len(per))
	t.Logf("LabelInfoList PER\n%v", hex.Dump(per))

	result, err := perDecodeLabelInfoList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("LabelInfoList PER - decoded\n%v", result)
}
