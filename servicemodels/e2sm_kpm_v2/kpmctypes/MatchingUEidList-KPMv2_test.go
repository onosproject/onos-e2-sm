// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeMatchingUeIDList(t *testing.T) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}
	muel := &e2sm_kpm_v2.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	xer, err := xerEncodeMatchingUeIDList(muel)
	assert.NilError(t, err)
	//assert.Equal(t, 125, len(xer))
	t.Logf("MatchingUEidList XER\n%s", string(xer))
}

func Test_xerDecodeMatchingUeIDList(t *testing.T) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}
	muel := &e2sm_kpm_v2.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	xer, err := xerEncodeMatchingUeIDList(muel)
	assert.NilError(t, err)
	//assert.Equal(t, 125, len(xer))
	t.Logf("MatchingUEidList XER\n%s", string(xer))

	result, err := xerDecodeMatchingUeIDList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MatchingUEidList XER - decoded\n%s", result)
}

func Test_perEncodeMatchingUeIDList(t *testing.T) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}
	muel := &e2sm_kpm_v2.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	per, err := perEncodeMatchingUeIDList(muel)
	assert.NilError(t, err)
	//assert.Equal(t, 10, len(per))
	t.Logf("MatchingUEidList PER\n%s", string(per))
}

func Test_perDecodeMatchingUeIDList(t *testing.T) {

	muei := &e2sm_kpm_v2.MatchingUeidItem{
		UeId: &e2sm_kpm_v2.UeIdentity{
			Value: "SomeUE",
		},
	}
	muel := &e2sm_kpm_v2.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	per, err := perEncodeMatchingUeIDList(muel)
	assert.NilError(t, err)
	//assert.Equal(t, 10, len(per))
	t.Logf("MatchingUEidList PER\n%s", string(per))

	result, err := perDecodeMatchingUeIDList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("MatchingUEidList PER - decoded\n%v", result)
}
