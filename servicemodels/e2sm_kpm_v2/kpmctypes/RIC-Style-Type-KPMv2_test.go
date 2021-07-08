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

func Test_xerEncodeRicStyleType(t *testing.T) {

	ricStyleType := &e2sm_kpm_v2.RicStyleType{
		Value: 111,
	}
	xer, err := xerEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	//assert.Equal(t, 37, len(xer))
	t.Logf("RicStyleType XER\n%s", string(xer))
}

func Test_xerDecodeRicStyleType(t *testing.T) {

	ricStyleType := &e2sm_kpm_v2.RicStyleType{
		Value: 111,
	}
	xer, err := xerEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	//assert.Equal(t, 37, len(xer))
	t.Logf("RicStyleType XER\n%s", string(xer))

	result, err := xerDecodeRicStyleType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicStyleType XER - decoded\n%s", result)
}

func Test_perEncodeRicStyleType(t *testing.T) {

	ricStyleType := &e2sm_kpm_v2.RicStyleType{
		Value: 111,
	}
	per, err := perEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	t.Logf("RicStyleType PER\n%v", hex.Dump(per))
}

func Test_perDecodeRicStyleType(t *testing.T) {

	ricStyleType := &e2sm_kpm_v2.RicStyleType{
		Value: 111,
	}
	per, err := perEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	t.Logf("RicStyleType PER\n%v", hex.Dump(per))

	result, err := perDecodeRicStyleType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicStyleType PER - decoded\n%v", result)
}
