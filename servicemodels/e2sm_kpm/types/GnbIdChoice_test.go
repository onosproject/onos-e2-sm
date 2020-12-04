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
	assert.Assert(t, reflect.TypeOf(GnbIDChoice{}) == reflect.TypeOf(*choiceID), "GnbIdChoice{} types are mismatched")

}

func TestGnbIdChoice_SetId(t *testing.T) {

	id := BitString{
		Value: 0x9bcd4,
		Len:   22,
	}

	choiceID := NewGnbIDChoice()
	choiceID.SetID(id)

	assert.Assert(t, choiceID.ID.Value == 0x9bcd4, "Mismatch of GnbIdChoice ID values")
	assert.Assert(t, choiceID.ID.Len == 22, "Mismatch of GnbIdChoice ID lengths")

}

//func TestGnbIdChoice_SetTypeGnb(t *testing.T) {
//
//	gnbType := "GNb"
//
//	choiceId := NewGnbIdChoice()
//	choiceId.SetTypeGNb()
//
//	assert.Assert(t, choiceId.Type == gnbType, "Mismatch of GnbIdChoice type (GNb)")
//
//}
//
//func TestGnbIdChoice_SetTypeNgGnb(t *testing.T) {
//
//	gnbType := "NgGNb"
//
//	choiceId := NewGnbIdChoice()
//	choiceId.SetTypeNgGNb()
//
//	assert.Assert(t, choiceId.Type == gnbType, "Mismatch of GnbIdChoice type (NgGNb)")
//
//}

func TestGnbIdChoice_GetId(t *testing.T) {

	id := BitString{
		Value: 0x9bcd4,
		Len:   22,
	}

	choiceID := NewGnbIDChoice()
	choiceID.SetID(id)

	getID := choiceID.GetID()

	assert.Assert(t, getID.GetValue() == 0x9bcd4, "Mismatch of GnbIdChoice ID values")
	assert.Assert(t, getID.GetLen() == 22, "Mismatch of GnbIdChoice ID lengths")

}

//func TestGnbIdChoice_GetType(t *testing.T) {
//
//	choiceId := NewGnbIdChoice()
//	choiceId.SetTypeGNb()
//	assert.Assert(t, choiceId.GetType() == "GNb", "Mismatch of GnbIdChoice type (GNb)")
//
//	choiceId.SetTypeNgGNb()
//	assert.Assert(t, choiceId.GetType() == "NgGNb", "Mismatch of GnbIdChoice type (NgGNb)")
//
//}

func TestGnbIdChoice_GetGnbIdChoice(t *testing.T) {

	id := BitString{
		Value: 0x9bcd4,
		Len:   22,
	}

	choiceID1 := NewGnbIDChoice()
	choiceID1.SetID(id)
	//choiceId1.SetTypeGNb()

	choiceID2 := choiceID1.GnbIDChoice()

	assert.Assert(t, choiceID1.GetID().GetValue() == choiceID2.GetID().GetValue(), "Mismatch of GnbIdChoice ID values")
	assert.Assert(t, choiceID1.GetID().GetLen() == choiceID2.GetID().GetLen(), "Mismatch of GnbIdChoice ID lengths")
	//assert.Assert(t, choiceId1.GetType() == choiceId2.GetType(), "Mismatch of GnbIdChoice type (GNb)")

}
