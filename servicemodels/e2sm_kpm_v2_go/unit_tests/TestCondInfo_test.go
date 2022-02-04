// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerTCI = "00000000  08 40 01 15                                       |.@..|"

func Test_perEncodingTestCondInfo(t *testing.T) {

	testCondInfo := &e2sm_kpm_v2_go.TestCondInfo{
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
	}

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondInfo, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondInfo PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondInfo{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("TestCondInfo PER - decoded\n%v", &result)
	assert.Equal(t, testCondInfo.GetTestValue().GetValueInt(), result.GetTestValue().GetValueInt())
	assert.Equal(t, testCondInfo.GetTestExpr().Number(), result.GetTestExpr().Number())
	assert.Equal(t, testCondInfo.GetTestType().GetAMbr().Number(), result.GetTestType().GetAMbr().Number())
}

func Test_perTestCondInfoCompareBytes(t *testing.T) {

	testCondInfo := &e2sm_kpm_v2_go.TestCondInfo{
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
	}

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondInfo, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondInfo PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCI)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_stupidExperiment2(t *testing.T) {
	perRefBytes, err := hexlib.DumpToByte(refPerTCI)
	assert.NilError(t, err)
	t.Logf("TestCondInfo PER\n%v", hex.Dump(perRefBytes))

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	result := e2sm_kpm_v2_go.TestCondInfo{}
	err = aper.UnmarshalWithParams(perRefBytes, &result, "valueExt")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("TestCondInfo PER - decoded\n%v", &result)
}
