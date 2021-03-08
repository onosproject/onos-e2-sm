// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodePlmnIdentity(t *testing.T) {

	plmnID := &e2sm_kpm_v2.PlmnIdentity{
		Value: []byte{0x21, 0x22, 0x23},
	}
	xer, err := xerEncodePlmnIdentity(plmnID)
	assert.NilError(t, err)
	assert.Equal(t, 40, len(xer))
	t.Logf("PlmnIdentity XER\n%s", string(xer))
}

func Test_xerDecodePlmnIdentity(t *testing.T) {

	plmnID := &e2sm_kpm_v2.PlmnIdentity{
		Value: []byte{0x21, 0x22, 0x23},
	}
	xer, err := xerEncodePlmnIdentity(plmnID)
	assert.NilError(t, err)
	assert.Equal(t, 40, len(xer))
	t.Logf("PlmnIdentity XER\n%s", string(xer))

	result, err := xerDecodePlmnIdentity(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("PlmnIdentity XER - decoded\n%s", result)
}

func Test_perEncodePlmnIdentity(t *testing.T) {

	plmnID := &e2sm_kpm_v2.PlmnIdentity{
		Value: []byte{0x21, 0x22, 0x23},
	}
	per, err := perEncodePlmnIdentity(plmnID)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("PlmnIdentity PER\n%s", string(per))
}

func Test_perDecodePlmnIdentity(t *testing.T) {

	plmnID := &e2sm_kpm_v2.PlmnIdentity{
		Value: []byte{0x21, 0x22, 0x23},
	}
	per, err := perEncodePlmnIdentity(plmnID)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("PlmnIdentity PER\n%s", string(per))

	result, err := perDecodePlmnIdentity(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("PlmnIdentity PER - decoded\n%v", result)
}
