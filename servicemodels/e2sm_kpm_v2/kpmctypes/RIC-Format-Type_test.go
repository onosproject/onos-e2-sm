// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeRicFormatType(t *testing.T) {

	ricFormatType := &e2sm_kpm_v2.RicFormatType{
		Value: 32,
	}
	xer, err := xerEncodeRicFormatType(ricFormatType)
	assert.NilError(t, err)
	assert.Equal(t, 38, len(xer))
	t.Logf("RicFormatType XER\n%s", string(xer))
}

func Test_xerDecodeRicFormatType(t *testing.T) {

	ricFormatType := &e2sm_kpm_v2.RicFormatType{
		Value: 32,
	}
	xer, err := xerEncodeRicFormatType(ricFormatType)
	assert.NilError(t, err)
	assert.Equal(t, 38, len(xer))
	t.Logf("RicFormatType XER\n%s", string(xer))

	result, err := xerDecodeRicFormatType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicFormatType XER - decoded\n%s", result)
}

func Test_perEncodeRicFormatType(t *testing.T) {

	ricFormatType := &e2sm_kpm_v2.RicFormatType{
		Value: 32,
	}
	per, err := perEncodeRicFormatType(ricFormatType)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("RicFormatType PER\n%s", string(per))
}

func Test_perDecodeRicFormatType(t *testing.T) {

	ricFormatType := &e2sm_kpm_v2.RicFormatType{
		Value: 32,
	}
	per, err := perEncodeRicFormatType(ricFormatType)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("RicFormatType PER\n%s", string(per))

	result, err := perDecodeRicFormatType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicFormatType PER - decoded\n%s", result)
}
