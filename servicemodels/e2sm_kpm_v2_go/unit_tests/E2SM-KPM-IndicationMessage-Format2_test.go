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

var refPerIndMsgF2 = "00000000  08 30 38 00 00 40 10 6f  6e 66 00 01 1f ff f0 01  |.08..@.onf......|\n" +
	"00000010  02 03 40 40 01 02 03 00  17 68 18 00 1e 00 01 70  |..@@.....h.....p|\n" +
	"00000020  00 00 18 00 00 00 00 00  7a 00 01 c7 00 03 14 28  |........z......(|\n" +
	"00000030  42 00 01 15 00 00 00 06  53 6f 6d 65 55 45 00 00  |B.......SomeUE..|\n" +
	"00000040  40 03 00 15 20 09 80 d0  16 33 33 33 33 33 33 40  |@... ....333333@|"

func createIndicationMessageFormat2() (*e2sm_kpm_v2_go.E2SmKpmIndicationMessageFormat2, error) {

	muei := &e2sm_kpm_v2_go.MatchingUeidItem{
		UeId: &e2sm_kpm_v2_go.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}
	muel := &e2sm_kpm_v2_go.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2_go.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	var lmm int32 = 1
	var br int32 = 25

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

	measType := &e2sm_kpm_v2_go.MeasurementType{
		MeasurementType: &e2sm_kpm_v2_go.MeasurementType_MeasName{
			MeasName: &e2sm_kpm_v2_go.MeasurementTypeName{
				Value: "onf",
			},
		},
	}

	mcUEIDi := &e2sm_kpm_v2_go.MeasurementCondUeidItem{
		MatchingCond:     mcl,
		MeasType:         measType,
		MatchingUeidList: muel,
	}

	mcUEIDl := &e2sm_kpm_v2_go.MeasurementCondUeidList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementCondUeidItem, 0),
	}
	mcUEIDl.Value = append(mcUEIDl.Value, mcUEIDi)

	measRecord := &e2sm_kpm_v2_go.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2_go.MeasurementRecordItem, 0),
	}

	item1 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	measRecord.Value = append(measRecord.Value, item1)

	item2 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Real{
			Real: 22.2,
		},
	}
	measRecord.Value = append(measRecord.Value, item2)

	item3 := &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	measRecord.Value = append(measRecord.Value, item3)

	incf := e2sm_kpm_v2_go.IncompleteFlag_INCOMPLETE_FLAG_TRUE
	measDataItem := &e2sm_kpm_v2_go.MeasurementDataItem{
		MeasRecord:     measRecord,
		IncompleteFlag: &incf,
	}

	measData := &e2sm_kpm_v2_go.MeasurementData{
		Value: make([]*e2sm_kpm_v2_go.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	imf2 := &e2sm_kpm_v2_go.E2SmKpmIndicationMessageFormat2{
		SubscriptId: &e2sm_kpm_v2_go.SubscriptionId{
			Value: 12345,
		},
		CellObjId: &e2sm_kpm_v2_go.CellObjectId{
			Value: "123",
		},
		GranulPeriod: &e2sm_kpm_v2_go.GranularityPeriod{
			Value: 21,
		},
		MeasCondUeidList: mcUEIDl,
		MeasData:         measData,
	}

	//if err := imf2.Validate(); err != nil {
	//	return nil, err
	//}
	return imf2, nil
}

func Test_perEncodingE2SmKpmIndicationMessageFormat2(t *testing.T) {

	imf2, err := createIndicationMessageFormat2()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*imf2, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessageFormat2 PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmIndicationMessageFormat2{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SmKpmIndicationMessageFormat2 PER - decoded\n%v", result)
}

func Test_perE2SmKpmIndicationMessageFormat2CompareBytes(t *testing.T) {

	imf2, err := createIndicationMessageFormat2()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*imf2, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessageFormat2 PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerIndMsgF2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
