// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeQci(t *testing.T) {

	qci := &e2sm_kpm_v2.Qci{
		Value: 32,
	}
	xer, err := xerEncodeQci(qci)
	assert.NilError(t, err)
	//assert.Equal(t, 14, len(xer))
	t.Logf("QCI XER\n%s", string(xer))
}

func Test_xerDecodeQci(t *testing.T) {

	qci := &e2sm_kpm_v2.Qci{
		Value: 32,
	}
	xer, err := xerEncodeQci(qci)
	assert.NilError(t, err)
	//assert.Equal(t, 14, len(xer))
	t.Logf("QCI XER\n%s", string(xer))

	result, err := xerDecodeQci(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("QCI XER - decoded\n%s", result)
}

func Test_perEncodeQci(t *testing.T) {

	qci := &e2sm_kpm_v2.Qci{
		Value: 32,
	}
	per, err := perEncodeQci(qci)
	assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	t.Logf("QCI PER\n%v", hex.Dump(per))
}

func Test_perDecodeQci(t *testing.T) {

	qci := &e2sm_kpm_v2.Qci{
		Value: 32,
	}
	per, err := perEncodeQci(qci)
	assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	t.Logf("QCI PER\n%v", hex.Dump(per))

	result, err := perDecodeQci(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("QCI PER - decoded\n%v", result)
}
