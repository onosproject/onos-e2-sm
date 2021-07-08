// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerMeasCL string = "00000000  00 00 20 00 7a 00 01 42  10 01 15 1f ff f0 01 02  |.. .z..B........|\n" +
	"00000010  03 40 40 01 02 03 00 17  68 18 00 1e 00 01 70 00  |.@@.....h.....p.|\n" +
	"00000020  00 18 00 00 00 00 00 7a  00 01 c7 00 03 14 20     |.......z...... |"

func createMeasurementCondList() (*e2sm_kpm_v2_go.MeasurementCondList, error) {

	measCondItem := &e2sm_kpm_v2_go.MeasurementCondItem{
		MeasType: &e2sm_kpm_v2_go.MeasurementType{
			MeasurementType: &e2sm_kpm_v2_go.MeasurementType_MeasId{
				MeasId: &e2sm_kpm_v2_go.MeasurementTypeId{
					Value: 123,
				},
			},
		},
		MatchingCond: &e2sm_kpm_v2_go.MatchingCondList{
			Value: make([]*e2sm_kpm_v2_go.MatchingCondItem, 0),
		},
	}

	mci1 := &e2sm_kpm_v2_go.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2_go.MatchingCondItem_TestCondInfo{
			TestCondInfo: &e2sm_kpm_v2_go.TestCondInfo{
				TestType: &e2sm_kpm_v2_go.TestCondType{
					TestCondType: &e2sm_kpm_v2_go.TestCondType_AMbr{
						AMbr: e2sm_kpm_v2_go.AMBR_AMBR_TRUE,
					},
				},
				TestExpr: e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
				TestValue: &e2sm_kpm_v2_go.TestCondValue{
					TestCondValue: &e2sm_kpm_v2_go.TestCondValue_ValueInt{
						ValueInt: 21,
					},
				},
			},
		},
	}
	measCondItem.MatchingCond.Value = append(measCondItem.MatchingCond.Value, mci1)

	var br int32 = 25
	var lmm int32 = 1
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END

	mci2 := &e2sm_kpm_v2_go.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2_go.MatchingCondItem_MeasLabel{
			MeasLabel: &e2sm_kpm_v2_go.MeasurementLabel{
				PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
					Value: []byte{0x1, 0x2, 0x3},
				},
				SliceId: &e2sm_kpm_v2_go.Snssai{
					SD:  []byte{0x01, 0x02, 0x03},
					SSt: []byte{0x01},
				},
				FiveQi: &e2sm_kpm_v2_go.FiveQi{
					Value: 23,
				},
				QFi: &e2sm_kpm_v2_go.Qfi{
					Value: 52,
				},
				QCi: &e2sm_kpm_v2_go.Qci{
					Value: 24,
				},
				QCimax: &e2sm_kpm_v2_go.Qci{
					Value: 30,
				},
				QCimin: &e2sm_kpm_v2_go.Qci{
					Value: 1,
				},
				ARpmax: &e2sm_kpm_v2_go.Arp{
					Value: 15,
				},
				ARpmin: &e2sm_kpm_v2_go.Arp{
					Value: 1,
				},
				BitrateRange:     &br,
				LayerMuMimo:      &lmm,
				SUm:              &sum,
				DistBinX:         &dbx,
				DistBinY:         &dby,
				DistBinZ:         &dbz,
				PreLabelOverride: &plo,
				StartEndInd:      &seind,
			},
		},
	}
	measCondItem.MatchingCond.Value = append(measCondItem.MatchingCond.Value, mci2)

	measCondList := &e2sm_kpm_v2_go.MeasurementCondList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementCondItem, 0),
	}
	measCondList.Value = append(measCondList.Value, measCondItem)
	//if err := measCondList.Validate(); err != nil {
	//	return nil, err
	//}
	return measCondList, nil
}

func Test_perEncodeMeasurementCondList(t *testing.T) {

	mcl, err := createMeasurementCondList()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*mcl, "")
	assert.NilError(t, err)
	t.Logf("MeasurementCondList PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementCondList{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementCondList PER - decoded\n%v", result)
}

func Test_perMeasurementCondListCompareBytes(t *testing.T) {

	mcl, err := createMeasurementCondList()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*mcl, "")
	assert.NilError(t, err)
	t.Logf("MeasurementCondList PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasCL)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
