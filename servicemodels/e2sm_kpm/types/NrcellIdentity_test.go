// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestNrcellIdentity_NewNrcellIdentity(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	nrcellIdentity := NewNrcellIdentity(NewBitString(value, length))
	assert.Equal(t, reflect.TypeOf(NrcellIdentity{}), reflect.TypeOf(*nrcellIdentity), "NrcellIdentity{} types are mismatched")
	assert.Equal(t, nrcellIdentity.Value.Value, value, "NrcellIdentity{} values are mismatched")
	assert.Equal(t, nrcellIdentity.Value.Len, length, "NrcellIdentity{} lengths are mismatched")
}

func TestNrcellIdentity_GetValue(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	nrcellIdentity := NewNrcellIdentity(NewBitString(value, length))
	nrValue := nrcellIdentity.GetValue()

	assert.Equal(t, nrValue.GetValue(), value, "Test_NrcellIdentity GetValue values mismatch")
	assert.Equal(t, nrValue.GetLen(), length, "Test_NrcellIdentity GetValue lengths mismatch")
}

func TestNrcellIdentity_GetNrcellIdentity(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	nrcellIdentity1 := NewNrcellIdentity(NewBitString(value, length))
	nrcellIdentity2 := nrcellIdentity1.GetNrcellIdentity()

	assert.Equal(t, nrcellIdentity1.GetValue().GetValue(), nrcellIdentity2.GetValue().GetValue(), "Test_NrcellIdentity GetNrcellIdentity values mismatch")
	assert.Equal(t, nrcellIdentity1.GetValue().GetLen(), nrcellIdentity2.GetValue().GetLen(), "Test_NrcellIdentity GetNrcellIdentity lengths mismatch")
}
