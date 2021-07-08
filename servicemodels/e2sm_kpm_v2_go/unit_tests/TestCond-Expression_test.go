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

var refPerTestCondExpr string = "00000000  40                                                |@|"

func Test_perEncodingTestCondExpression(t *testing.T) {

	testCondExpression := e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_PRESENT

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondExpression, "valueExt,valueLB:0,valueUB:4")
	assert.NilError(t, err)
	t.Logf("TestCondExpression PER\n%v", hex.Dump(per))

	var result int32
	err = aper.UnmarshalWithParams(per, &result, "valueExt,valueLB:0,valueUB:4")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("TestCondExpression PER - decoded\n%v", result)
}

func Test_perTestCondExpressionCompareBytes(t *testing.T) {

	testCondExpression := e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_PRESENT

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondExpression, "valueExt,valueLB:0,valueUB:4")
	assert.NilError(t, err)
	t.Logf("TestCondExpression PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTestCondExpr)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
