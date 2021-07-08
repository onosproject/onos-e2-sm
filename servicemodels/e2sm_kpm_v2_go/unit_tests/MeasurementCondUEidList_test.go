// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerMeasCondUeIDL string = "00000000  00 00 40 10 6f 6e 66 00  01 1f ff f0 01 02 03 40  |..@.onf........@|\n" +
	"00000010  40 01 02 03 00 17 68 18  00 1e 00 01 70 00 00 18  |@.....h.....p...|\n" +
	"00000020  00 00 00 00 00 7a 00 01  c7 00 03 14 28 42 00 01  |.....z......(B..|\n" +
	"00000030  15 00 00 00 06 53 6f 6d  65 55 45                 |.....SomeUE|"

func createMeasurementCondUEIDList() (*e2sm_kpm_v2_go.MeasurementCondUeidList, error) {

	muei := &e2sm_kpm_v2_go.MatchingUeidItem{
		UeId: &e2sm_kpm_v2_go.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}
	muel := &e2sm_kpm_v2_go.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2_go.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	plmnID := []byte{0x01, 0x02, 0x03}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	var qfi int32 = 62
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2sm_kpm_v2_go.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	labelInfoItem, err := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	if err != nil {
		return nil, err
	}

	mci1, err := pdubuilder.CreateMatchingCondItemMeasLabel(labelInfoItem.GetMeasLabel())
	if err != nil {
		return nil, err
	}

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

	//if err := mcUEIDl.Validate(); err != nil {
	//	return nil, err
	//}
	return mcUEIDl, nil
}

func Test_perEncodeMeasurementCondUEIDList(t *testing.T) {

	mcueIDl, err := createMeasurementCondUEIDList()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*mcueIDl, "")
	assert.NilError(t, err)
	t.Logf("MeasurementCondUEIDList PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementCondUeidList{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementCondUEIDList PER - decoded\n%v", result)
}

func Test_perMeasurementCondUEIDListCompareBytes(t *testing.T) {

	mcueIDl, err := createMeasurementCondUEIDList()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*mcueIDl, "")
	assert.NilError(t, err)
	t.Logf("MeasurementCondUEIDList PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasCondUeIDL)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
