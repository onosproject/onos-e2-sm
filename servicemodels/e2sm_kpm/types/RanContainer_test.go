// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestRanContainer_NewRanContainer(t *testing.T) {

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)

	assert.Assert(t, reflect.TypeOf(RanContainer{}) == reflect.TypeOf(*ranContainer), "RanContainer{} types are mismatched")
	assert.DeepEqual(t, ranContainer.Value, value)
}

func TestRanContainer_GetValue(t *testing.T) {

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)

	assert.DeepEqual(t, ranContainer.GetValue(), value)
}

func TestRanContainer_GetRanContainer(t *testing.T) {

	value := []byte{0x22, 0x21}
	ranContainer1 := NewRanContainer(value)
	ranContainer2 := ranContainer1.GetRanContainer()

	assert.DeepEqual(t, ranContainer1, ranContainer2)
}
