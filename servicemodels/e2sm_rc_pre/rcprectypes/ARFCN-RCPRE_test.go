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

func createNrarfcn() *e2sm_rc_pre_v2.Arfcn {

	return &e2sm_rc_pre_v2.Arfcn{
		Arfcn: &e2sm_rc_pre_v2.Arfcn_NrArfcn{
			NrArfcn: &e2sm_rc_pre_v2.Nrarfcn{
				Value: 253,
			},
		},
	}
}

func createEarfcn() *e2sm_rc_pre_v2.Arfcn {

	return &e2sm_rc_pre_v2.Arfcn{
		Arfcn: &e2sm_rc_pre_v2.Arfcn_EArfcn{
			EArfcn: &e2sm_rc_pre_v2.Earfcn{
				Value: 198,
			},
		},
	}
}

func Test_xerDecodeARFCN(t *testing.T) {

	nrarfcn := createNrarfcn()

	xer, err := xerEncodeARFCN(nrarfcn)
	assert.NilError(t, err)
	//assert.Equal(t, 44, len(xer))
	t.Logf("ARFCN (NRARFCN) XER\n%s", string(xer))

	result, err := xerDecodeARFCN(xer)
	assert.NilError(t, err)
	t.Logf("Decoded ARFCN (NRARFCN) is \n%v", result)
	//assert.DeepEqual(t, nrarfcn, result)

	earfcn := createEarfcn()

	xer, err = xerEncodeARFCN(earfcn)
	assert.NilError(t, err)
	//assert.Equal(t, 32, len(xer))
	t.Logf("ARFCN (EARFCN) XER\n%s", string(xer))

	result, err = xerDecodeARFCN(xer)
	assert.NilError(t, err)
	t.Logf("Decoded ARFCN (EARFCN) is \n%v", result)
	//assert.DeepEqual(t, nrarfcn, result)
}

func Test_perDecodeARFCN(t *testing.T) {

	nrarfcn := createNrarfcn()

	per, err := perEncodeARFCN(nrarfcn)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("ARFCN (NRARFCN) PER\n%v", hex.Dump(per))

	result, err := perDecodeARFCN(per)
	assert.NilError(t, err)
	t.Logf("Decoded ARFCN (NRARFCN) is \n%v", result)
	//assert.DeepEqual(t, nrarfcn, result)

	earfcn := createEarfcn()

	per, err = perEncodeARFCN(earfcn)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("ARFCN (EARFCN) PER\n%v", hex.Dump(per))

	result, err = perDecodeARFCN(per)
	assert.NilError(t, err)
	t.Logf("Decoded ARFCN (EARFCN) is \n%v", result)
	//assert.DeepEqual(t, nrarfcn, result)
}
