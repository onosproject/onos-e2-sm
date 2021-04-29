// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	"encoding/binary"
	"encoding/hex"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodingBolean(t *testing.T) {

	xer, err := xerEncodeBoolean(true)
	assert.NilError(t, err)
	assert.Equal(t, 27, len(xer))
	t.Logf("BOOLEAN (True) XER\n%s", string(xer))

	result, err := xerDecodeBoolean(xer)
	assert.NilError(t, err)
	assert.Assert(t, result == true)
	t.Logf("BOOLEAN (True) XER - decoded\n%v", result)

	xer, err = xerEncodeBoolean(false)
	assert.NilError(t, err)
	assert.Equal(t, 28, len(xer))
	t.Logf("BOOLEAN (False) XER\n%s", string(xer))

	result, err = xerDecodeBoolean(xer)
	assert.NilError(t, err)
	assert.Assert(t, result == false)
	t.Logf("BOOLEAN (False) XER - decoded\n%v", result)

	choiceC := [8]byte{}
	bC := newBoolean(true)
	binary.LittleEndian.PutUint32(choiceC[:], uint32(*bC))

	var a [8]byte
	copy(a[:], choiceC[:])
	b := decodeBooleanBytes(a)
	t.Logf("Decoded Boolean from bytes is %v\n", b)

}

func Test_perEncodingBoolean(t *testing.T) {

	per, err := perEncodeBoolean(false)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("BOOLEAN (False) PER\n%s", hex.Dump(per))

	result, err := perDecodeBoolean(per)
	assert.NilError(t, err)
	assert.Assert(t, result == false)
	t.Logf("BOOLEAN (False) PER - decoded\n%v", result)

	per, err = perEncodeBoolean(true)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("BOOLEAN (True) PER\n%s", hex.Dump(per))

	result, err = perDecodeBoolean(per)
	assert.NilError(t, err)
	assert.Assert(t, result == true)
	t.Logf("BOOLEAN (True) PER - decoded\n%v", result)
}
