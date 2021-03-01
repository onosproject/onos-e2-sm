// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
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
}

func Test_xerEncodeLabelInfoItem(t *testing.T) {

	lii := createLabelInfoItem()

	xer, err := xerEncodeLabelInfoItem(lii)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("LabelInfoItem XER\n%s", string(xer))
}

func Test_xerDecodeLabelInfoItem(t *testing.T) {

	lii := createLabelInfoItem()

	xer, err := xerEncodeLabelInfoItem(lii)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("LabelInfoItem XER\n%s", string(xer))

	result, err := xerDecodeLabelInfoItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeLabelInfoItem(t *testing.T) {

	lii := createLabelInfoItem()

	per, err := perEncodeLabelInfoItem(lii)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("LabelInfoItem PER\n%s", string(per))
}

func Test_perDecodeLabelInfoItem(t *testing.T) {

	lii := createLabelInfoItem()

	per, err := perEncodeLabelInfoItem(lii)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("LabelInfoItem PER\n%s", string(per))

	result, err := perDecodeLabelInfoItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
