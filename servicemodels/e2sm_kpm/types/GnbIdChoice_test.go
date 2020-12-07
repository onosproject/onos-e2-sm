// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestGnbIdChoice_NewGnbIdChoice(t *testing.T) {

	choiceID := NewGnbIDChoice()
	assert.Equal(t, reflect.TypeOf(GnbIDChoice{}), reflect.TypeOf(*choiceID), "GnbIdChoice{} types are mismatched")

}

func TestGnbIdChoice_SetId(t *testing.T) {

	var value uint64 = 0x9bcd4
	var len uint32 = 22

	choiceID := NewGnbIDChoice().SetID(NewBitString().SetValue(value).SetLen(len))

	assert.Equal(t, choiceID.ID.Value, value, "Mismatch of GnbIdChoice ID values")
	assert.Equal(t, choiceID.ID.Len, len, "Mismatch of GnbIdChoice ID lengths")

}

func TestGnbIdChoice_GetId(t *testing.T) {

	var value uint64 = 0x9bcd4
	var len uint32 = 22

	choiceID := NewGnbIDChoice().SetID(NewBitString().SetValue(value).SetLen(len))

	getID := choiceID.GetID()

	assert.Equal(t, getID.GetValue(), value, "Mismatch of GnbIdChoice ID values")
	assert.Equal(t, getID.GetLen(), len, "Mismatch of GnbIdChoice ID lengths")

}

func TestGnbIdChoice_GetGnbIdChoice(t *testing.T) {

	var value uint64 = 0x9bcd4
	var len uint32 = 22

	choiceID1 := NewGnbIDChoice().SetID(NewBitString().SetValue(value).SetLen(len))

	choiceID2 := choiceID1.GetGnbIDChoice()

	assert.Equal(t, choiceID1.GetID().GetValue(), choiceID2.GetID().GetValue(), "Mismatch of GnbIdChoice ID values")
	assert.Equal(t, choiceID1.GetID().GetLen(), choiceID2.GetID().GetLen(), "Mismatch of GnbIdChoice ID lengths")

}
