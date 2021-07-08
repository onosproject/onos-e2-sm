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

func Test_xerEncodeRicStyleName(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2.RicStyleName{
		Value: "SomeName",
	}
	xer, err := xerEncodeRicStyleName(ricStyleName)
	assert.NilError(t, err)
	//assert.Equal(t, 42, len(xer))
	t.Logf("RicStyleName XER\n%s", string(xer))
}

func Test_xerDecodeRicStyleName(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2.RicStyleName{
		Value: "SomeName",
	}
	xer, err := xerEncodeRicStyleName(ricStyleName)
	assert.NilError(t, err)
	//assert.Equal(t, 42, len(xer))
	t.Logf("RicStyleName XER\n%s", string(xer))

	result, err := xerDecodeRicStyleName(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicStyleName XER - decoded\n%s", result)
}

func Test_perEncodeRicStyleName(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2.RicStyleName{
		Value: "SomeName",
	}
	per, err := perEncodeRicStyleName(ricStyleName)
	assert.NilError(t, err)
	//assert.Equal(t, 10, len(per))
	t.Logf("RicStyleName PER\n%v", hex.Dump(per))
}

func Test_perDecodeRicStyleName(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2.RicStyleName{
		Value: "SomeName",
	}
	per, err := perEncodeRicStyleName(ricStyleName)
	assert.NilError(t, err)
	//assert.Equal(t, 10, len(per))
	t.Logf("RicStyleName PER\n%v", hex.Dump(per))

	result, err := perDecodeRicStyleName(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicStyleName PER - decoded\n%v", result)
}
