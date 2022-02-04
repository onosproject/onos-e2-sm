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

func createLabelInfoItem() *e2sm_kpm_v2.LabelInfoItem {

	return &e2sm_kpm_v2.LabelInfoItem{
		MeasLabel: &e2sm_kpm_v2.MeasurementLabel{
			PlmnId: &e2sm_kpm_v2.PlmnIdentity{
				Value: []byte{0x01, 0x02, 0x03},
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
}

func Test_xerEncodeLabelInfoItem(t *testing.T) {

	lii := createLabelInfoItem()

	xer, err := xerEncodeLabelInfoItem(lii)
	assert.NilError(t, err)
	//assert.Equal(t, 675, len(xer))
	t.Logf("LabelInfoItem XER\n%s", string(xer))
}

func Test_xerDecodeLabelInfoItem(t *testing.T) {

	lii := createLabelInfoItem()

	xer, err := xerEncodeLabelInfoItem(lii)
	assert.NilError(t, err)
	//assert.Equal(t, 675, len(xer))
	t.Logf("LabelInfoItem XER\n%s", string(xer))

	result, err := xerDecodeLabelInfoItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("LabelInfoItem XER - decoded\n%s", result)
}

func Test_perEncodeLabelInfoItem(t *testing.T) {

	lii := createLabelInfoItem()

	per, err := perEncodeLabelInfoItem(lii)
	assert.NilError(t, err)
	//assert.Equal(t, 36, len(per))
	t.Logf("LabelInfoItem PER\n%v", hex.Dump(per))
}

func Test_perDecodeLabelInfoItem(t *testing.T) {

	lii := createLabelInfoItem()

	per, err := perEncodeLabelInfoItem(lii)
	assert.NilError(t, err)
	//assert.Equal(t, 36, len(per))
	t.Logf("LabelInfoItem PER\n%v", hex.Dump(per))

	result, err := perDecodeLabelInfoItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("LabelInfoItem PER - decoded\n%v", result)
}
