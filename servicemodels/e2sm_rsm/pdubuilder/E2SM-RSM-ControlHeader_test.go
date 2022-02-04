// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/encoder"
	"gotest.tools/assert"
	"testing"
)

func Test_E2SmRsmControlHeader(t *testing.T) {

	ch1 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandSliceCreate())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch1)

	// APER validation
	per1, err := encoder.PerEncodeE2SmRsmControlHeader(ch1)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-ControlHeader (Slice Create) PER\n%v", hex.Dump(per1))

	result1, err := encoder.PerDecodeE2SmRsmControlHeader(per1)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("E2SM-RSM-ControlHeader (Slice Create) PER - decoded\n%v", result1)
	assert.DeepEqual(t, ch1.String(), result1.String())

	ch2 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandSliceUpdate())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch2)

	// APER validation
	per2, err := encoder.PerEncodeE2SmRsmControlHeader(ch2)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-ControlHeader (Slice Update) PER\n%v", hex.Dump(per2))

	result2, err := encoder.PerDecodeE2SmRsmControlHeader(per2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("E2SM-RSM-ControlHeader (Slice Update) PER - decoded\n%v", result2)
	assert.DeepEqual(t, ch2.String(), result2.String())

	ch3 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandSliceDelete())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch3)

	// APER validation
	per3, err := encoder.PerEncodeE2SmRsmControlHeader(ch3)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-ControlHeader (Slice Delete) PER\n%v", hex.Dump(per3))

	result3, err := encoder.PerDecodeE2SmRsmControlHeader(per3)
	assert.NilError(t, err)
	assert.Assert(t, result3 != nil)
	t.Logf("E2SM-RSM-ControlHeader (Slice Delete) PER - decoded\n%v", result3)
	assert.DeepEqual(t, ch3.String(), result3.String())

	ch4 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandUeAssociate())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch4)

	// APER validation
	per4, err := encoder.PerEncodeE2SmRsmControlHeader(ch4)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-ControlHeader (UE Associate) PER\n%v", hex.Dump(per4))

	result4, err := encoder.PerDecodeE2SmRsmControlHeader(per4)
	assert.NilError(t, err)
	assert.Assert(t, result4 != nil)
	t.Logf("E2SM-RSM-ControlHeader (UE Associate) PER - decoded\n%v", result4)
	assert.DeepEqual(t, ch4.String(), result4.String())

	ch5 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandEventTriggers())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch5)

	// APER validation
	per5, err := encoder.PerEncodeE2SmRsmControlHeader(ch5)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-ControlHeader (Event Triggers) PER\n%v", hex.Dump(per5))

	result5, err := encoder.PerDecodeE2SmRsmControlHeader(per5)
	assert.NilError(t, err)
	assert.Assert(t, result5 != nil)
	t.Logf("E2SM-RSM-ControlHeader (Event Triggers) PER - decoded\n%v", result5)
	assert.DeepEqual(t, ch5.String(), result5.String())
}
