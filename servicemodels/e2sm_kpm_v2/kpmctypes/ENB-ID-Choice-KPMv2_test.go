// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createEnbIDChoiceMacro() *e2sm_kpm_v2.EnbIdChoice {

	return &e2sm_kpm_v2.EnbIdChoice{
		EnbIdChoice: &e2sm_kpm_v2.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: &e2sm_kpm_v2.BitString{
				Value: 0x9bcd4,
				Len:   20,
			},
		},
	}
}

func createEnbIDChoiceShortMacro() *e2sm_kpm_v2.EnbIdChoice {

	return &e2sm_kpm_v2.EnbIdChoice{
		EnbIdChoice: &e2sm_kpm_v2.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: &e2sm_kpm_v2.BitString{
				Value: 0x9bcd4,
				Len:   18,
			},
		},
	}
}

func createEnbIDChoiceLongMacro() *e2sm_kpm_v2.EnbIdChoice {

	return &e2sm_kpm_v2.EnbIdChoice{
		EnbIdChoice: &e2sm_kpm_v2.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: &e2sm_kpm_v2.BitString{
				Value: 0x9bcd4,
				Len:   21,
			},
		},
	}
}

func Test_xerEncodeEnbIDChoice(t *testing.T) {

	enbIDchoice := createEnbIDChoiceMacro()

	xer, err := xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 101, len(xer))
	t.Logf("EnbIDChoice (Macro) XER\n%s", string(xer))

	enbIDchoice = createEnbIDChoiceShortMacro()

	xer, err = xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 109, len(xer))
	t.Logf("EnbIDChoice (ShortMacro) XER\n%s", string(xer))

	enbIDchoice = createEnbIDChoiceLongMacro()

	xer, err = xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 110, len(xer))
	t.Logf("EnbIDChoice (LongMacro) XER\n%s", string(xer))
}

func Test_xerDecodeEnbIDChoice(t *testing.T) {

	enbIDchoice := createEnbIDChoiceMacro()

	xer, err := xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 101, len(xer))
	t.Logf("EnbIDChoice (Macro) XER\n%s", string(xer))

	result, err := xerDecodeEnbIDChoice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (Macro) XER - decoded\n%s", result)

	enbIDchoice = createEnbIDChoiceShortMacro()

	xer, err = xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 109, len(xer))
	t.Logf("EnbIDChoice (ShortMacro) XER\n%s", string(xer))

	result, err = xerDecodeEnbIDChoice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (ShortMacro) XER - decoded\n%s", result)

	enbIDchoice = createEnbIDChoiceLongMacro()

	xer, err = xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 110, len(xer))
	t.Logf("EnbIDChoice (LongMacro) XER\n%s", string(xer))

	result, err = xerDecodeEnbIDChoice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (LongMacro) XER - decoded\n%s", result)
}

func Test_perEncodeEnbIDChoice(t *testing.T) {

	enbIDchoice := createEnbIDChoiceMacro()

	per, err := perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (Macro) PER\n%s", string(per))

	enbIDchoice = createEnbIDChoiceShortMacro()

	per, err = perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (ShortMacro)PER\n%s", string(per))

	enbIDchoice = createEnbIDChoiceLongMacro()

	per, err = perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (LongMacro) PER\n%s", string(per))
}

func Test_perDecodeEnbIDChoice(t *testing.T) {

	enbIDchoice := createEnbIDChoiceMacro()

	per, err := perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (Macro) PER\n%s", string(per))

	result, err := perDecodeEnbIDChoice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (Macro) PER - decoded\n%s", result)

	enbIDchoice = createEnbIDChoiceShortMacro()

	per, err = perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (ShortMacro) PER\n%s", string(per))

	result, err = perDecodeEnbIDChoice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (ShortMacro) PER - decoded\n%s", result)

	enbIDchoice = createEnbIDChoiceLongMacro()

	per, err = perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (LongMacro) PER\n%s", string(per))

	result, err = perDecodeEnbIDChoice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (LongMacro) PER - decoded\n%s", result)
}
