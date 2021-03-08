// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createGnbIDChoice() *e2sm_kpm_v2.GnbIdChoice {

	return &e2sm_kpm_v2.GnbIdChoice{
		GnbIdChoice: &e2sm_kpm_v2.GnbIdChoice_GnbId{
			GnbId: &e2sm_kpm_v2.BitString{
				Value: 0x9bcd4,
				Len:   22,
			},
		},
	}
}

func Test_xerEncodeGnbIDChoice(t *testing.T) {

	gnbIDchoice := createGnbIDChoice()

	xer, err := xerEncodeGnbIDChoice(gnbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 91, len(xer))
	t.Logf("GnbIDChoice XER\n%s", string(xer))
}

func Test_xerDecodeGnbIDChoice(t *testing.T) {

	gnbIDchoice := createGnbIDChoice()

	xer, err := xerEncodeGnbIDChoice(gnbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 91, len(xer))
	t.Logf("GnbIDChoice XER\n%s", string(xer))

	result, err := xerDecodeGnbIDChoice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GnbIDChoice XER - decoded\n%s", result)
}

func Test_perEncodeGnbIDChoice(t *testing.T) {

	gnbIDchoice := createGnbIDChoice()

	per, err := perEncodeGnbIDChoice(gnbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("GnbIDChoice PER\n%s", string(per))
}

func Test_perDecodeGnbIDChoice(t *testing.T) {

	gnbIDchoice := createGnbIDChoice()

	per, err := perEncodeGnbIDChoice(gnbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("GnbIDChoice XER\n%s", string(per))

	result, err := perDecodeGnbIDChoice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GnbIDChoice XER - decoded\n%s", result)
}
