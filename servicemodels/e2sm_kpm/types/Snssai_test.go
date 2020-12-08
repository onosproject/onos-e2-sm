// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestSnssai_NewSnssai(t *testing.T) {

	sst := []byte{0x12}
	sd := []byte{0x22, 0x21, 0x20}
	snssai, err := NewSnssai(sst, sd)
	assert.NilError(t, err)

	assert.Assert(t, reflect.TypeOf(Snssai{}) == reflect.TypeOf(*snssai), "Snssai{} types are mismatched")
	assert.DeepEqual(t, snssai.SSt, sst)
	assert.DeepEqual(t, snssai.SD, sd)
}

func TestSnssai_GetSSt(t *testing.T) {

	sst := []byte{0x12}
	sd := []byte{0x22, 0x21, 0x20}
	snssai, err := NewSnssai(sst, sd)
	assert.NilError(t, err)

	assert.DeepEqual(t, snssai.GetSSt(), sst)
}

func TestSnssai_GetSD(t *testing.T) {

	sst := []byte{0x12}
	sd := []byte{0x22, 0x21, 0x20}
	snssai, err := NewSnssai(sst, sd)
	assert.NilError(t, err)

	assert.DeepEqual(t, snssai.GetSD(), sd)
}

func TestSnssai_GetSnssai(t *testing.T) {

	sst := []byte{0x12}
	sd := []byte{0x22, 0x21, 0x20}
	snssai1, err := NewSnssai(sst, sd)
	assert.NilError(t, err)
	snssai2 := snssai1.GetSnssai()

	assert.DeepEqual(t, snssai1, snssai2)
}
