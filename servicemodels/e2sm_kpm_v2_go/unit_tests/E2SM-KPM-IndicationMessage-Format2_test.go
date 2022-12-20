// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

// var refPerIndMsgF2 = "00000000  68 30 38 00 00 03 31 32  33 00 14 00 00 40 10 6f  |h08...123....@.o|\n" +
//
//	"00000010  6e 66 00 01 1f ff f0 21  22 23 40 40 01 02 03 00  |nf.....!\"#@@....|\n" +
//	"00000020  17 68 18 00 1e 00 01 70  00 00 18 00 00 00 00 00  |.h.....p........|\n" +
//	"00000030  7a 00 01 c7 00 03 14 28  42 00 01 15 00 00 00 06  |z......(B.......|\n" +
//	"00000040  53 6f 6d 65 55 45 00 00  40 03 00 15 20 09 80 d0  |SomeUE..@... ...|\n" +
//	"00000050  16 33 33 33 33 33 33 40                           |.333333@|"
var refPerIndMsgF2noReal = "00000000  68 30 38 00 00 03 31 32  33 00 14 00 00 40 10 6f  |h08...123....@.o|\n" +
	"00000010  6e 66 00 01 1f ff f0 21  22 23 40 40 01 02 03 00  |nf.....!\"#@@....|\n" +
	"00000020  17 68 18 00 1e 00 01 70  00 00 18 00 00 00 00 00  |.h.....p........|\n" +
	"00000030  7a 00 01 c7 00 03 14 28  42 00 01 15 00 00 00 06  |z......(B.......|\n" +
	"00000040  53 6f 6d 65 55 45 00 00  40 02 00 15 40           |SomeUE..@...@|"

func createIndicationMessageFormat2() (*e2smkpmv2.E2SmKpmIndicationMessageFormat2, error) {

	muei := &e2smkpmv2.MatchingUeidItem{
		UeId: &e2smkpmv2.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}
	muel := &e2smkpmv2.MatchingUeidList{
		Value: make([]*e2smkpmv2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	seind := e2smkpmv2.StartEndInd_START_END_IND_END
	plo := e2smkpmv2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	sum := e2smkpmv2.SUM_SUM_TRUE
	var lmm int32 = 1
	var br int32 = 25

	mci1 := &e2smkpmv2.MatchingCondItem{
		MatchingCondItem: &e2smkpmv2.MatchingCondItem_MeasLabel{
			MeasLabel: &e2smkpmv2.MeasurementLabel{
				PlmnId: &e2smkpmv2.PlmnIdentity{
					Value: []byte{0x21, 0x22, 0x23},
				},
				SliceId: &e2smkpmv2.Snssai{
					SD:  []byte{0x01, 0x02, 0x03},
					SSt: []byte{0x01},
				},
				FiveQi: &e2smkpmv2.FiveQi{
					Value: 23,
				},
				QFi: &e2smkpmv2.Qfi{
					Value: 52,
				},
				QCi: &e2smkpmv2.Qci{
					Value: 24,
				},
				QCimax: &e2smkpmv2.Qci{
					Value: 30,
				},
				QCimin: &e2smkpmv2.Qci{
					Value: 1,
				},
				ARpmax: &e2smkpmv2.Arp{
					Value: 15,
				},
				ARpmin: &e2smkpmv2.Arp{
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

	mci2 := &e2smkpmv2.MatchingCondItem{
		MatchingCondItem: &e2smkpmv2.MatchingCondItem_TestCondInfo{
			TestCondInfo: &e2smkpmv2.TestCondInfo{
				TestType: &e2smkpmv2.TestCondType{
					TestCondType: &e2smkpmv2.TestCondType_AMbr{
						AMbr: e2smkpmv2.AMBR_AMBR_TRUE,
					},
				},
				TestExpr: e2smkpmv2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
				TestValue: &e2smkpmv2.TestCondValue{
					TestCondValue: &e2smkpmv2.TestCondValue_ValueInt{
						ValueInt: 21,
					},
				},
			},
		},
	}

	//if err := mci2.Validate(); err != nil {
	//	return nil, err
	//}

	mcl := &e2smkpmv2.MatchingCondList{
		Value: make([]*e2smkpmv2.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci1)
	mcl.Value = append(mcl.Value, mci2)

	measType := &e2smkpmv2.MeasurementType{
		MeasurementType: &e2smkpmv2.MeasurementType_MeasName{
			MeasName: &e2smkpmv2.MeasurementTypeName{
				Value: "onf",
			},
		},
	}

	mcUEIDi := &e2smkpmv2.MeasurementCondUeidItem{
		MatchingCond:     mcl,
		MeasType:         measType,
		MatchingUeidList: muel,
	}

	mcUEIDl := &e2smkpmv2.MeasurementCondUeidList{
		Value: make([]*e2smkpmv2.MeasurementCondUeidItem, 0),
	}
	mcUEIDl.Value = append(mcUEIDl.Value, mcUEIDi)

	measRecord := &e2smkpmv2.MeasurementRecord{
		Value: make([]*e2smkpmv2.MeasurementRecordItem, 0),
	}

	item1 := &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	measRecord.Value = append(measRecord.Value, item1)

	//item2 := &e2smkpmv2.MeasurementRecordItem{
	//	MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Real{
	//		Real: 22.2,
	//	},
	//}
	//measRecord.Value = append(measRecord.Value, item2)

	item3 := &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	measRecord.Value = append(measRecord.Value, item3)

	incf := e2smkpmv2.IncompleteFlag_INCOMPLETE_FLAG_TRUE
	measDataItem := &e2smkpmv2.MeasurementDataItem{
		MeasRecord:     measRecord,
		IncompleteFlag: &incf,
	}

	measData := &e2smkpmv2.MeasurementData{
		Value: make([]*e2smkpmv2.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	imf2 := &e2smkpmv2.E2SmKpmIndicationMessageFormat2{
		SubscriptId: &e2smkpmv2.SubscriptionId{
			Value: 12345,
		},
		CellObjId: &e2smkpmv2.CellObjectId{
			Value: "123",
		},
		GranulPeriod: &e2smkpmv2.GranularityPeriod{
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

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(imf2, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessageFormat2 PER\n%v", hex.Dump(per))

	result := e2smkpmv2.E2SmKpmIndicationMessageFormat2{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("E2SmKpmIndicationMessageFormat2 PER - decoded\n%v", &result)
	assert.Equal(t, imf2.GetGranulPeriod().GetValue(), result.GetGranulPeriod().GetValue())
	assert.Equal(t, imf2.GetSubscriptId().GetValue(), result.GetSubscriptId().GetValue())
	assert.Equal(t, imf2.GetCellObjId().GetValue(), result.GetCellObjId().GetValue())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMeasType().GetMeasName().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMeasType().GetMeasName().GetValue())
	assert.DeepEqual(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestValue().GetValueInt(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestValue().GetValueInt())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestType().GetAMbr().Number(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestType().GetAMbr().Number())
	assert.Equal(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestExpr().Number(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestExpr().Number())
	assert.DeepEqual(t, imf2.GetMeasCondUeidList().GetValue()[0].GetMatchingUeidList().GetValue()[0].GetUeId().GetValue(), result.GetMeasCondUeidList().GetValue()[0].GetMatchingUeidList().GetValue()[0].GetUeId().GetValue())
	assert.Equal(t, imf2.GetMeasData().GetValue()[0].GetIncompleteFlag().Number(), result.GetMeasData().GetValue()[0].GetIncompleteFlag().Number())
	assert.Equal(t, imf2.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger(), result.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[0].GetInteger())
	//assert.Equal(t, imf2.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetReal(), result.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetReal())
	assert.Equal(t, imf2.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue(), result.GetMeasData().GetValue()[0].GetMeasRecord().GetValue()[1].GetNoValue())
}

func Test_perE2SmKpmIndicationMessageFormat2CompareBytes(t *testing.T) {

	imf2, err := createIndicationMessageFormat2()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(imf2, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("E2SmKpmIndicationMessageFormat2 PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerIndMsgF2noReal)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
