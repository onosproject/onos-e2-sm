// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createIndicationMessageFormat2() (*e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2, error) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}
	muel := &e2sm_kpm_v2.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	mci1 := &e2sm_kpm_v2.MatchingCondItem{
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

	if err := mci1.Validate(); err != nil {
		return nil, err
	}

	mci2 := &e2sm_kpm_v2.MatchingCondItem{
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

	if err := mci2.Validate(); err != nil {
		return nil, err
	}

	mcl := &e2sm_kpm_v2.MatchingCondList{
		Value: make([]*e2sm_kpm_v2.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci1)
	mcl.Value = append(mcl.Value, mci2)

	measType := &e2sm_kpm_v2.MeasurementType{
		MeasurementType: &e2sm_kpm_v2.MeasurementType_MeasName{
			MeasName: &e2sm_kpm_v2.MeasurementTypeName{
				Value: "onf",
			},
		},
	}

	mcUEIDi := &e2sm_kpm_v2.MeasurementCondUeidItem{
		MatchingCond:     mcl,
		MeasType:         measType,
		MatchingUeidList: muel,
	}

	mcUEIDl := &e2sm_kpm_v2.MeasurementCondUeidList{
		Value: make([]*e2sm_kpm_v2.MeasurementCondUeidItem, 0),
	}
	mcUEIDl.Value = append(mcUEIDl.Value, mcUEIDi)

	measRecord := &e2sm_kpm_v2.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2.MeasurementRecordItem, 0),
	}

	item1 := &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_Integer{
			Integer: 21,
		},
	}
	measRecord.Value = append(measRecord.Value, item1)

	item2 := &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_Real{
			Real: 22.2,
		},
	}
	measRecord.Value = append(measRecord.Value, item2)

	item3 := &e2sm_kpm_v2.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
	measRecord.Value = append(measRecord.Value, item3)

	measDataItem := &e2sm_kpm_v2.MeasurementDataItem{
		MeasRecord:     measRecord,
		IncompleteFlag: e2sm_kpm_v2.IncompleteFlag_INCOMPLETE_FLAG_TRUE,
	}

	measData := &e2sm_kpm_v2.MeasurementData{
		Value: make([]*e2sm_kpm_v2.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	imf2 := &e2sm_kpm_v2.E2SmKpmIndicationMessageFormat2{
		SubscriptId: &e2sm_kpm_v2.SubscriptionId{
			Value: 12345,
		},
		//CellObjId: &e2sm_kpm_v2.CellObjectId{
		//	Value: "123",
		//},
		//GranulPeriod: &e2sm_kpm_v2.GranularityPeriod{
		//	Value: 21,
		//},
		MeasCondUeidList: mcUEIDl,
		MeasData:         measData,
	}

	if err := imf2.Validate(); err != nil {
		return nil, err
	}
	return imf2, nil
}

func Test_xerEncodeE2SmKpmIndicationMessageFormat2(t *testing.T) {

	imf2, err := createIndicationMessageFormat2()
	assert.NilError(t, err)

	xer, err := xerEncodeE2SmKpmIndicationMessageFormat2(imf2)
	assert.NilError(t, err)
	//assert.Equal(t, 2412, len(xer))
	//assert.Equal(t, 2345, len(xer)) // without GranulPeriod and CellObjID
	t.Logf("E2SmKpmIndicationMessageFormat2 XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmIndicationMessageFormat2(t *testing.T) {

	imf2, err := createIndicationMessageFormat2()
	assert.NilError(t, err)

	xer, err := xerEncodeE2SmKpmIndicationMessageFormat2(imf2)
	assert.NilError(t, err)
	//assert.Equal(t, 2412, len(xer))
	//assert.Equal(t, 2345, len(xer)) // without GranulPeriod and CellObjID
	t.Logf("E2SmKpmIndicationMessageFormat2 XER\n%s", string(xer))

	result, err := xerDecodeE2SmKpmIndicationMessageFormat2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationMessageFormat2 XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmIndicationMessageFormat2(t *testing.T) {

	imf2, err := createIndicationMessageFormat2()
	assert.NilError(t, err)

	per, err := perEncodeE2SmKpmIndicationMessageFormat2(imf2)
	assert.NilError(t, err)
	//assert.Equal(t, 89, len(per))
	assert.Equal(t, 80, len(per)) // without GranulPeriod and CellObjID
	t.Logf("E2SmKpmIndicationMessageFormat2 PER\n%v", hex.Dump(per))
}

func Test_perDecodeE2SmKpmIndicationMessageFormat2(t *testing.T) {

	imf2, err := createIndicationMessageFormat2()
	assert.NilError(t, err)

	per, err := perEncodeE2SmKpmIndicationMessageFormat2(imf2)
	assert.NilError(t, err)
	//assert.Equal(t, 89, len(per))
	assert.Equal(t, 80, len(per)) // without GranulPeriod and CellObjID
	t.Logf("E2SmKpmIndicationMessageFormat2 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2SmKpmIndicationMessageFormat2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationMessageFormat2 PER - decoded\n%v", result)
}
