//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcControlOutcomeFormat1(t *testing.T) {

	ranParameterValue, err := CreateRanparameterValuePrintableString("Ran Parameter Value")
	assert.NilError(t, err)

	item, err := CreateE2SmRcControlOutcomeFormat1Item(1, ranParameterValue)
	assert.NilError(t, err)

	ranPList := make([]*e2smrcv1.E2SmRcControlOutcomeFormat1Item, 0)
	ranPList = append(ranPList, item)

	msg, err := CreateE2SmRcControlOutcomeFormat1(ranPList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcControlOutcome(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcControlOutcome(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcControlOutcomeFormat2(t *testing.T) {

	ranParameterValue, err := CreateRanparameterValuePrintableString("Ran Parameter Value")
	assert.NilError(t, err)

	item, err := CreateE2SmRcControlOutcomeFormat2RanpItem(1, ranParameterValue)
	assert.NilError(t, err)

	ranPList := make([]*e2smrcv1.E2SmRcControlOutcomeFormat2RanpItem, 0)
	ranPList = append(ranPList, item)

	controlOutcomeItem, err := CreateE2SmRcControlOutcomeFormat2ControlOutcomeItem(12, ranPList)
	assert.NilError(t, err)

	controlOutcomeList := make([]*e2smrcv1.E2SmRcControlOutcomeFormat2ControlOutcomeItem, 0)
	controlOutcomeList = append(controlOutcomeList, controlOutcomeItem)

	styleItem, err := CreateE2SmRcControlOutcomeFormat2StyleItem(21, controlOutcomeList)
	assert.NilError(t, err)

	ricControlStyleList := make([]*e2smrcv1.E2SmRcControlOutcomeFormat2StyleItem, 0)
	ricControlStyleList = append(ricControlStyleList, styleItem)

	msg, err := CreateE2SmRcControlOutcomeFormat2(ricControlStyleList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcControlOutcome(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcControlOutcome(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcControlOutcomeFormat3(t *testing.T) {

	ranParameterValue, err := CreateRanparameterValuePrintableString("Ran Parameter Value")
	assert.NilError(t, err)

	ranParameterValueType, err := CreateRanparameterValueTypeChoiceElementFalse(ranParameterValue)
	assert.NilError(t, err)

	item, err := CreateE2SmRcControlOutcomeFormat3Item(1, ranParameterValueType)
	assert.NilError(t, err)

	ranPList := make([]*e2smrcv1.E2SmRcControlOutcomeFormat3Item, 0)
	ranPList = append(ranPList, item)

	msg, err := CreateE2SmRcControlOutcomeFormat3(ranPList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcControlOutcome(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcControlOutcome(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}
