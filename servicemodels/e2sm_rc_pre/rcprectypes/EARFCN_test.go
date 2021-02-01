// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeEARFCN(t *testing.T) {

	var earfcn int32 = 253

	c := &e2sm_rc_pre_ies.Earfcn{
		Value: earfcn,
	}

	xer, err := xerEncodeEARFCN(c)
	assert.NilError(t, err)
	t.Logf("EARFCN XER\n%s", string(xer))
}

func Test_decodeEARFCN(t *testing.T) {

	var earfcn int32 = 253

	c := &e2sm_rc_pre_ies.Earfcn{
		Value: earfcn,
	}

	earfcnC := newEARFCN(c)
	result := decodeEARFCN(earfcnC)
	assert.Equal(t, earfcn, result.Value, "Something went wrong, comparison is incorrect")
}

func Test_perEncodeEARFCN(t *testing.T) {

	var earfcn int32 = 253

	c := &e2sm_rc_pre_ies.Earfcn{
		Value: earfcn,
	}

	earfcnPer, err := perEncodeEARFCN(c)
	assert.NilError(t, err)
	assert.Assert(t, earfcnPer != nil)
	t.Logf("EARFCN PER\n%v", earfcnPer)
}

func Test_perDecodeEARFCN(t *testing.T) {

	var earfcn int32 = 253

	c := &e2sm_rc_pre_ies.Earfcn{
		Value: earfcn,
	}

	earfcnPer, err := perEncodeEARFCN(c)
	assert.NilError(t, err)
	assert.Assert(t, earfcnPer != nil)
	t.Logf("EARFCN PER\n%v", earfcnPer)

	result, err := perDecodeEARFCN(earfcnPer)
	assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	t.Logf("EARFCN PER decoded is \n%v", result)

	assert.Equal(t, earfcn, result.Value, "Encoded and decoded values are not the same")
}
