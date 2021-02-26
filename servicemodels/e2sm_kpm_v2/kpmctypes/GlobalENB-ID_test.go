// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createGlobalEnbID() *e2sm_kpm_v2.GlobalEnbId {

	return &e2sm_kpm_v2.GlobalEnbId{
		PLmnIdentity: &e2sm_kpm_v2.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		ENbId: &e2sm_kpm_v2.EnbId{
			EnbId: &e2sm_kpm_v2.EnbId_MacroENbId{
				MacroENbId: &e2sm_kpm_v2.BitString{
					Value: 0x9bcd4,
					Len:   22,
				},
			},
		},
	}
}

func Test_xerEncodeGlobalEnbID(t *testing.T) {

	globalEnbID := createGlobalEnbID()

	xer, err := xerEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("GlobalEnbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalEnbID(t *testing.T) {

	globalEnbID := createGlobalEnbID()

	xer, err := xerEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("GlobalEnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalEnbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalEnbID(t *testing.T) {

	globalEnbID := createGlobalEnbID()

	per, err := perEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("GlobalEnbID PER\n%s", string(per))
}

func Test_perDecodeGlobalEnbID(t *testing.T) {

	globalEnbID := createGlobalEnbID()

	per, err := perEncodeGlobalEnbID(globalEnbID)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("GlobalEnbID PER\n%s", string(per))

	result, err := perDecodeGlobalEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalEnbID PER - decoded\n%s", result)
}
