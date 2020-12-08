// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestGnbIDChoice_NewGnbIDChoice(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	choiceID := NewGnbIDChoice(NewBitString(value, length))
	assert.Equal(t, reflect.TypeOf(GnbIDChoice{}), reflect.TypeOf(*choiceID), "GnbIDChoice{} types are mismatched")
	assert.Equal(t, choiceID.ID.Value, value, "GnbIDChoice{} values are mismatched")
	assert.Equal(t, choiceID.ID.Len, length, "GnbIDChoice{} lengths are mismatched")
}

func TestGnbIDChoice_GetID(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	choiceID := NewGnbIDChoice(NewBitString(value, length))
	getID := choiceID.GetID()

	assert.Equal(t, getID.GetValue(), value, "Test_GnbIDChoice GetID values mismatch")
	assert.Equal(t, getID.GetLen(), length, "Test_GnbIDChoice GetID lengths mismatch")
}

func TestGnbIDChoice_GetGnbIDChoice(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	choiceID1 := NewGnbIDChoice(NewBitString(value, length))
	choiceID2 := choiceID1.GetGnbIDChoice()

	assert.Equal(t, choiceID1.GetID().GetValue(), choiceID2.GetID().GetValue(), "Test_GnbIDChoice GetGnbIDChoice values mismatch")
	assert.Equal(t, choiceID1.GetID().GetLen(), choiceID2.GetID().GetLen(), "Test_GnbIDChoice GetGnbIDChoice lengths mismatch")
}
