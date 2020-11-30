// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeRicStyleType(t *testing.T) {

	var ricType int32 = 11
	ricStyleType := &e2sm_kpm_ies.RicStyleType{
		Value: ricType,
	}

	xer, err := xerEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	assert.Equal(t, 36, len(xer))
	t.Logf("RIC-Style-Type XER\n%s", string(xer))
}

func Test_xerDecodeRicStyleType(t *testing.T) {

	var ricType int32 = 11
	ricStyleType := &e2sm_kpm_ies.RicStyleType{
		Value: ricType,
	}

	xer, err := xerEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	assert.Equal(t, 36, len(xer))
	t.Logf("RIC-Style-Type XER\n%s", string(xer))

	result, err := xerDecodeRicStyleType(xer)
	assert.NilError(t, err)
	assert.Equal(t, ricStyleType.Value, result.Value, "Encoded and decoded values are not the same")
}

func Test_perEncodeRicStyleType(t *testing.T) {

	var ricType int32 = 11
	ricStyleType := &e2sm_kpm_ies.RicStyleType{
		Value: ricType,
	}

	per, err := perEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("RIC-Style-Type PER\n%s", string(per))
}

func Test_perDecodeRicStyleType(t *testing.T) {

	var ricType int32 = 11
	ricStyleType := &e2sm_kpm_ies.RicStyleType{
		Value: ricType,
	}

	per, err := perEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("RIC-Style-Type PER\n%s", string(per))

	result, err := perDecodeRicStyleType(per)
	assert.NilError(t, err)
	assert.Equal(t, ricStyleType.Value, result.Value, "Encoded and decoded values are not the same")
}
