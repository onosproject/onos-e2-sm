// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeMatchingUeIDItem(t *testing.T) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}

	xer, err := xerEncodeMatchingUeIDItem(muei)
	assert.NilError(t, err)
	//assert.Equal(t, 74, len(xer))
	t.Logf("MatchingUEidItem XER\n%s", string(xer))
}

func Test_xerDecodeMatchingUeIDItem(t *testing.T) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}

	xer, err := xerEncodeMatchingUeIDItem(muei)
	assert.NilError(t, err)
	//assert.Equal(t, 74, len(xer))
	t.Logf("MatchingUEidItem XER\n%s", string(xer))

	result, err := xerDecodeMatchingUeIDItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MatchingUEidItem XER - decoded\n%s", result)
}

func Test_perEncodeMatchingUeIDItem(t *testing.T) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}

	per, err := perEncodeMatchingUeIDItem(muei)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("MatchingUEidItem PER\n%s", string(per))
}

func Test_perDecodeMatchingUeIDItem(t *testing.T) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}

	per, err := perEncodeMatchingUeIDItem(muei)
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("MatchingUEidItem PER\n%s", string(per))

	result, err := perDecodeMatchingUeIDItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MatchingUEidItem PER - decoded\n%v", result)
}
