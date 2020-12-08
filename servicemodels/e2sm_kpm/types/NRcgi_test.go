// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestNRcgi_NewNRcgi(t *testing.T) {

	nrcgi := NewNRcgi()

	assert.Equal(t, reflect.TypeOf(NRcgi{}), reflect.TypeOf(*nrcgi), "NRcgi{} types are mismatched")
}

func TestNRcgi_SetPlmnID(t *testing.T) {

	value := []byte{0x22, 0x21, 0x20}
	plmnID, err := NewPlmnID(value)
	assert.NilError(t, err)
	nrcgi := NewNRcgi().SetPlmnID(plmnID)

	assert.DeepEqual(t, nrcgi.PlmnID.GetValue(), value)
}

func TestNRcgi_SetNRcellID(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	nrcellIdentity := NewNrcellIdentity(NewBitString(value, length))
	nrcgi := NewNRcgi().SetNRcellID(nrcellIdentity)

	assert.Equal(t, nrcgi.NRcellID.GetValue().GetValue(), value, "Test_NRcgi() SetNRcellID values mismatch")
	assert.Equal(t, nrcgi.NRcellID.GetValue().GetLen(), length, "Test_GlobalGnbID() SetNRcellID lengths mismatch")
	assert.DeepEqual(t, nrcgi.NRcellID.GetNrcellIdentity(), nrcellIdentity.GetNrcellIdentity())
}

func TestNRcgi_GetPlmnID(t *testing.T) {

	value := []byte{0x22, 0x21, 0x20}
	plmnID, err := NewPlmnID(value)
	assert.NilError(t, err)
	nrcgi := NewNRcgi().SetPlmnID(plmnID)

	assert.DeepEqual(t, nrcgi.GetPlmnID().GetValue(), plmnID.GetValue())
	assert.DeepEqual(t, nrcgi.GetPlmnID(), plmnID)
}

func TestNRcgi_GetNRcellID(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	nrcellIdentity := NewNrcellIdentity(NewBitString(value, length))
	nrcgi := NewNRcgi().SetNRcellID(nrcellIdentity)

	assert.Equal(t, nrcgi.GetNRcellID().GetValue().GetValue(), value, "Test_GlobalGnbID() GetNRcellID values mismatch")
	assert.Equal(t, nrcgi.GetNRcellID().GetValue().GetLen(), length, "Test_GlobalGnbID() GetNRcellID lengths mismatch")
	assert.DeepEqual(t, nrcgi.GetNRcellID(), nrcellIdentity)
}

func TestNRcgi_GetNRcgi(t *testing.T) {

	var value uint64 = 0x9bcd4
	var length uint32 = 22
	bytes := []byte{0x22, 0x21, 0x20}
	plmnID, err := NewPlmnID(bytes)
	assert.NilError(t, err)
	nrcgi1 := NewNRcgi().SetPlmnID(plmnID).SetNRcellID(NewNrcellIdentity(NewBitString(value, length)))
	nrcgi2 := nrcgi1.GetNRcgi()

	assert.DeepEqual(t, nrcgi1, nrcgi2)
}
