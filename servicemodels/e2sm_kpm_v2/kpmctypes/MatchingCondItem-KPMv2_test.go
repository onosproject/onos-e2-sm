// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createMatchingCondItem1() (*e2sm_kpm_v2.MatchingCondItem, error) {

	measLabel := &e2sm_kpm_v2.MeasurementLabel{
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
	}

	mci, _ := pdubuilder.CreateMatchingCondItemMeasLabel(measLabel)

	if err := mci.Validate(); err != nil {
		return nil, err
	}
	return mci, nil
}

func createMatchingCondItem2() (*e2sm_kpm_v2.MatchingCondItem, error) {

	testCondInfo := &e2sm_kpm_v2.TestCondInfo{
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
	}

	mci, _ := pdubuilder.CreateMatchingCondItemTestCondInfo(testCondInfo)

	if err := mci.Validate(); err != nil {
		return nil, err
	}
	return mci, nil
}

func Test_xerEncodeMatchingCondItem(t *testing.T) {

	mci, err := createMatchingCondItem1()
	assert.NilError(t, err)

	xer, err := xerEncodeMatchingCondItem(mci)
	assert.NilError(t, err)
	//assert.Equal(t, 681, len(xer))
	t.Logf("MatchingCondItem (MeasurementLabel) XER\n%s", string(xer))

	mci, err = createMatchingCondItem2()
	assert.NilError(t, err)

	xer, err = xerEncodeMatchingCondItem(mci)
	assert.NilError(t, err)
	//assert.Equal(t, 271, len(xer))
	t.Logf("MatchingCondItem (TestCondInfo) XER\n%s", string(xer))
}

func Test_xerDecodeMatchingCondItem(t *testing.T) {

	mci, err := createMatchingCondItem1()
	assert.NilError(t, err)

	xer, err := xerEncodeMatchingCondItem(mci)
	assert.NilError(t, err)
	//assert.Equal(t, 681, len(xer))
	t.Logf("MatchingCondItem (MeasurementLabel) XER\n%s", string(xer))

	result, err := xerDecodeMatchingCondItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, 1, len(result.GetValue()))
	t.Logf("MatchingCondItem (MeasurementLabel) XER - decoded\n%s", result)

	mci, err = createMatchingCondItem2()
	assert.NilError(t, err)

	xer, err = xerEncodeMatchingCondItem(mci)
	assert.NilError(t, err)
	//assert.Equal(t, 271, len(xer))
	t.Logf("MatchingCondItem (TestCondInfo) XER\n%s", string(xer))

	result, err = xerDecodeMatchingCondItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, 1, len(result.GetValue()))
	t.Logf("MatchingCondItem (TestCondInfo) XER - decoded\n%s", result)
}

func Test_perEncodeMatchingCondItem(t *testing.T) {

	mci, err := createMatchingCondItem1()
	assert.NilError(t, err)

	per, err := perEncodeMatchingCondItem(mci)
	assert.NilError(t, err)
	//assert.Equal(t, 36, len(per))
	t.Logf("MatchingCondItem (MeasurementLabel) PER\n%v", hex.Dump(per))

	mci, err = createMatchingCondItem2()
	assert.NilError(t, err)

	per, err = perEncodeMatchingCondItem(mci)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("MatchingCondItem (TestCondInfo) PER\n%v", hex.Dump(per))
}

func Test_perDecodeMatchingCondItem(t *testing.T) {

	mci, err := createMatchingCondItem1()
	assert.NilError(t, err)

	per, err := perEncodeMatchingCondItem(mci)
	assert.NilError(t, err)
	//assert.Equal(t, 36, len(per))
	t.Logf("MatchingCondItem (MeasurementLabel) PER\n%v", hex.Dump(per))

	result, err := perDecodeMatchingCondItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, 1, len(result.GetValue()))
	t.Logf("MatchingCondItem PER - decoded\n%v", result)

	mci, err = createMatchingCondItem2()
	assert.NilError(t, err)

	per, err = perEncodeMatchingCondItem(mci)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("MatchingCondItem (TestCondInfo) PER\n%v", hex.Dump(per))

	result, err = perDecodeMatchingCondItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, 1, len(result.GetValue()))
	t.Logf("MatchingCondItem PER - decoded\n%v", result)
}
