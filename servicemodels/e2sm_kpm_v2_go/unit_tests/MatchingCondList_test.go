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

var refPerMCL string = "00000000  00 01 1f ff f0 01 02 03  40 40 01 02 03 00 17 68  |........@@.....h|\n" +
	"00000010  18 00 1e 00 01 70 00 00  18 00 00 00 00 00 7a 00  |.....p........z.|\n" +
	"00000020  01 c7 00 03 14 28 42 00  01 15                    |.....(B...|"

func createMatchingCondList() (*e2sm_kpm_v2_go.MatchingCondList, error) {

	var br int32 = 25
	var lmm int32 = 1
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END

	mci1 := &e2sm_kpm_v2_go.MatchingCondItem{
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

	//if err := mci1.Validate(); err != nil {
	//	return nil, err
	//}

	mci2 := &e2sm_kpm_v2_go.MatchingCondItem{
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

	//if err := mci2.Validate(); err != nil {
	//	return nil, err
	//}

	mcl := &e2sm_kpm_v2_go.MatchingCondList{
		Value: make([]*e2sm_kpm_v2_go.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci1)
	mcl.Value = append(mcl.Value, mci2)

	//if err := mcl.Validate(); err != nil {
	//	return nil, err
	//}
	return mcl, nil
}

func Test_perEncodeMatchingCondList(t *testing.T) {

	mcl, err := createMatchingCondList()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*mcl, "")
	assert.NilError(t, err)
	t.Logf("MatchingCondItem PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MatchingCondList{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MatchingCondList PER - decoded\n%v", result)
}

func Test_perMatchingCondListCompareBytes(t *testing.T) {

	mcl, err := createMatchingCondList()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*mcl, "")
	assert.NilError(t, err)
	t.Logf("MatchingCondItem PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMCL)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
