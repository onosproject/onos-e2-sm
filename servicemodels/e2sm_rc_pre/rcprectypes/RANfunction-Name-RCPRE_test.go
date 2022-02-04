// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

import (
	"encoding/hex"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func createRanfunctionName() *e2sm_rc_pre_v2.RanfunctionName {

	var ranFunctionShortName = "ONF"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 3

	ranfunctionName := &e2sm_rc_pre_v2.RanfunctionName{
		RanFunctionShortName:   ranFunctionShortName,
		RanFunctionE2SmOid:     ranFunctionE2SmOid,
		RanFunctionDescription: ranFunctionDescription,
		RanFunctionInstance:    ranFunctionInstance,
	}

	return ranfunctionName
}

func Test_xerEncodeRanfunctionName(t *testing.T) {

	ranfunctionName := createRanfunctionName()

	xer, err := xerEncodeRanfunctionName(ranfunctionName)
	assert.NilError(t, err)
	//assert.Equal(t, 268, len(xer))
	t.Logf("RANfunction-Name XER\n%s", string(xer))
}

func Test_xerDecodeRanfunctionName(t *testing.T) {

	ranfunctionName := createRanfunctionName()

	xer, err := xerEncodeRanfunctionName(ranfunctionName)
	assert.NilError(t, err)
	//assert.Equal(t, 268, len(xer))
	t.Logf("RANfunction-Name XER\n%s", string(xer))

	result, err := xerDecodeRanfunctionName(xer)
	assert.NilError(t, err)
	t.Logf("RANfunction-Name XER\n%v", result)
	assert.Equal(t, ranfunctionName.RanFunctionShortName, result.RanFunctionShortName, "Encoded and decoded RanFunctionShortName values are not the same")
	assert.Equal(t, ranfunctionName.RanFunctionE2SmOid, result.RanFunctionE2SmOid, "Encoded and decoded RanFunctionE2SmOid values are not the same")
	assert.Equal(t, ranfunctionName.RanFunctionDescription, result.RanFunctionDescription, "Encoded and decoded RanFunctionDescription values are not the same")
	assert.Equal(t, ranfunctionName.RanFunctionInstance, result.RanFunctionInstance, "Encoded and decoded RanFunctionInstance values are not the same")
}

func Test_perEncodeRanfunctionName(t *testing.T) {

	ranfunctionName := createRanfunctionName()

	per, err := perEncodeRanfunctionName(ranfunctionName)
	assert.NilError(t, err)
	assert.Equal(t, 29, len(per))
	t.Logf("RANfunction-Name PER\n%v", hex.Dump(per))
}

func Test_perDecodeRanfunctionName(t *testing.T) {

	ranfunctionName := createRanfunctionName()

	per, err := perEncodeRanfunctionName(ranfunctionName)
	assert.NilError(t, err)
	assert.Equal(t, 29, len(per))
	t.Logf("RANfunction-Name PER\n%v", hex.Dump(per))

	result, err := perDecodeRanfunctionName(per)
	assert.NilError(t, err)
	t.Logf("RANfunction-Name PER\n%v", result)
	assert.Equal(t, ranfunctionName.RanFunctionShortName, result.RanFunctionShortName, "Encoded and decoded RanFunctionShortName values are not the same")
	assert.Equal(t, ranfunctionName.RanFunctionE2SmOid, result.RanFunctionE2SmOid, "Encoded and decoded RanFunctionE2SmOid values are not the same")
	assert.Equal(t, ranfunctionName.RanFunctionDescription, result.RanFunctionDescription, "Encoded and decoded RanFunctionDescription values are not the same")
	assert.Equal(t, ranfunctionName.RanFunctionInstance, result.RanFunctionInstance, "Encoded and decoded RanFunctionInstance values are not the same")
}
