// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMeasurementCondItem() (*e2sm_kpm_v2.MeasurementCondItem, error) {

	measCondItem := &e2sm_kpm_v2.MeasurementCondItem{
		MeasType: &e2sm_kpm_v2.MeasurementType{
			MeasurementType: &e2sm_kpm_v2.MeasurementType_MeasId{
				MeasId: &e2sm_kpm_v2.MeasurementTypeId{
					Value: 123,
				},
			},
		},
		MatchingCond: &e2sm_kpm_v2.MatchingCondList{
			Value: make([]*e2sm_kpm_v2.MatchingCondItem, 0),
		},
	}

	mci1 := &e2sm_kpm_v2.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2.MatchingCondItem_TestCondInfo{
			TestCondInfo: &e2sm_kpm_v2.TestCondInfo{
				TestType: &e2sm_kpm_v2.TestCondType{
					TestCondType: &e2sm_kpm_v2.TestCondType_AMbr{
						AMbr: e2sm_kpm_v2.AMBR_AMBR_TRUE,
					},
				},
				TestExpr: e2sm_kpm_v2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
				TestValue: &e2sm_kpm_v2.TestCondValue{
					TestCondValue: &e2sm_kpm_v2.TestCondValue_ValueInt{
						ValueInt: 21,
					},
				},
			},
		},
	}
	measCondItem.MatchingCond.Value = append(measCondItem.MatchingCond.Value, mci1)

	mci2 := &e2sm_kpm_v2.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2.MatchingCondItem_MeasLabel{
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
					Value: 52,
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
		},
	}
	measCondItem.MatchingCond.Value = append(measCondItem.MatchingCond.Value, mci2)

	if err := measCondItem.Validate(); err != nil {
		return nil, err
	}
	return measCondItem, nil
}

func Test_xerEncodeMeasurementCondItem(t *testing.T) {

	mci, err := createMeasurementCondItem()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementCondItem(mci)
	assert.NilError(t, err)
	assert.Equal(t, 1284, len(xer))
	t.Logf("MeasurementCondItem XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementCondItem(t *testing.T) {

	mci, err := createMeasurementCondItem()
	assert.NilError(t, err)

	xer, err := xerEncodeMeasurementCondItem(mci)
	assert.NilError(t, err)
	assert.Equal(t, 1284, len(xer))
	t.Logf("MeasurementCondItem XER\n%s", string(xer))

	result, err := xerDecodeMeasurementCondItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//measRecord := result.GetValue()[0]
	//assert.Equal(t, 3, len(measRecord.GetValue()))
	t.Logf("MeasurementCondItem XER - decoded\n%s", result)
}

func Test_perEncodeMeasurementCondItem(t *testing.T) {

	mci, err := createMeasurementCondItem()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementCondItem(mci)
	assert.NilError(t, err)
	assert.Equal(t, 45, len(per))
	t.Logf("MeasurementCondItem PER\n%s", string(per))
}

func Test_perDecodeMeasurementCondItem(t *testing.T) {

	mci, err := createMeasurementCondItem()
	assert.NilError(t, err)

	per, err := perEncodeMeasurementCondItem(mci)
	assert.NilError(t, err)
	assert.Equal(t, 45, len(per))
	t.Logf("MeasurementCondItem PER\n%s", string(per))

	result, err := perDecodeMeasurementCondItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, 1, len(result.GetValue()))
	t.Logf("MeasurementCondItem PER - decoded\n%v", result)
}
