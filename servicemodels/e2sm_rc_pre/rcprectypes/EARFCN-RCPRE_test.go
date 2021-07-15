// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeEARFCN(t *testing.T) {

	var earfcn int32 = 253

	c := &e2sm_rc_pre_v2.Earfcn{
		Value: earfcn,
	}

	xer, err := xerEncodeEARFCN(c)
	assert.NilError(t, err)
	t.Logf("EARFCN XER\n%s", string(xer))
}

func Test_xerDecodeEARFCN(t *testing.T) {

	var earfcn int32 = 253

	c := &e2sm_rc_pre_v2.Earfcn{
		Value: earfcn,
	}

	xer, err := xerEncodeEARFCN(c)
	assert.NilError(t, err)
	t.Logf("EARFCN XER\n%s", string(xer))

	result, err := xerDecodeEARFCN(xer)
	assert.NilError(t, err)
	assert.Equal(t, earfcn, result.GetValue(), "Something went wrong, comparison is incorrect")
	t.Logf("EARFCN XER - decoded\n%v", result)
}

func Test_perEncodeEARFCN(t *testing.T) {

	var earfcn int32 = 253

	c := &e2sm_rc_pre_v2.Earfcn{
		Value: earfcn,
	}

	earfcnPer, err := perEncodeEARFCN(c)
	assert.NilError(t, err)
	assert.Assert(t, earfcnPer != nil)
	t.Logf("EARFCN PER\n%v", hex.Dump(earfcnPer))
}

func Test_perDecodeEARFCN(t *testing.T) {

	var earfcn int32 = 253

	c := &e2sm_rc_pre_v2.Earfcn{
		Value: earfcn,
	}

	earfcnPer, err := perEncodeEARFCN(c)
	assert.NilError(t, err)
	assert.Assert(t, earfcnPer != nil)
	t.Logf("EARFCN PER\n%v", hex.Dump(earfcnPer))

	result, err := perDecodeEARFCN(earfcnPer)
	assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	t.Logf("EARFCN PER decoded is \n%v", result)

	assert.Equal(t, earfcn, result.Value, "Encoded and decoded values are not the same")
}
