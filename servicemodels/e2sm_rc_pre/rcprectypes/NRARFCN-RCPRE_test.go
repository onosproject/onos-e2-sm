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

func createNRARFCN() *e2sm_rc_pre_v2.Nrarfcn {

	return &e2sm_rc_pre_v2.Nrarfcn{
		Value: 253,
	}
}

func Test_xerDecodeNRARFCN(t *testing.T) {

	nrarfcn := createNRARFCN()

	xer, err := xerEncodeNRARFCN(nrarfcn)
	assert.NilError(t, err)
	//assert.Equal(t, 23, len(xer))
	t.Logf("NRARFCN XER\n%s", string(xer))

	result, err := xerDecodeNRARFCN(xer)
	assert.NilError(t, err)
	assert.Equal(t, nrarfcn.GetValue(), result.GetValue(), "Encoded and decoded NRARFCN values are not the same")
	t.Logf("Decoded NRARFCN is \n%v", result)
}

func Test_perDecodeNRARFCN(t *testing.T) {

	nrarfcn := createNRARFCN()

	per, err := perEncodeNRARFCN(nrarfcn)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("NRARFCN PER\n%v", hex.Dump(per))

	result, err := perDecodeNRARFCN(per)
	assert.NilError(t, err)
	assert.Equal(t, nrarfcn.GetValue(), result.GetValue(), "Encoded and decoded NRARFCN values are not the same")
	t.Logf("Decoded NRARFCN is \n%v", result)
}
