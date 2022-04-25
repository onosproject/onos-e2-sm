//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcActionDefinitionFormat1(t *testing.T) {

	ranParameterIDs := make([]int64, 0)
	ranParameterIDs = append(ranParameterIDs, 1)
	ranParameterIDs = append(ranParameterIDs, 2)
	ranParameterIDs = append(ranParameterIDs, 3)
	ranParameterIDs = append(ranParameterIDs, 12)
	ranParameterIDs = append(ranParameterIDs, 41)
	ranParameterIDs = append(ranParameterIDs, 153)
	ranParameterIDs = append(ranParameterIDs, 7)

	msg, err := CreateE2SmRcActionDefinitionFormat1(11, ranParameterIDs)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcActionDefinition(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcActionDefinition(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcActionDefinitionFormat2(t *testing.T) {

	actionDefinitionItemFormat2List := make([]*e2smrcv1.E2SmRcActionDefinitionFormat2Item, 0)

	ricPolicyActionList := make([]*e2smrcv1.RicPolicyActionRanparameterItem, 0)

	ranParameterValue, err := CreateRanparameterValueBitS(&asn1.BitString{
		Value: []byte{0xC0},
		Len:   8,
	})
	assert.NilError(t, err)
	ranParameterValueType, err := CreateRanparameterValueTypeChoiceElementFalse(ranParameterValue)
	assert.NilError(t, err)

	actionItem1, err := CreateRicPolicyActionRanParameterItem(12, ranParameterValueType)
	assert.NilError(t, err)
	ricPolicyActionList = append(ricPolicyActionList, actionItem1)

	item1, err := CreateE2SmRcActionDefinitionFormat2Item(25, ricPolicyActionList)
	assert.NilError(t, err)
	actionDefinitionItemFormat2List = append(actionDefinitionItemFormat2List, item1)

	ranParameterTestingList := &e2smrcv1.RanparameterTestingList{
		Value: make([]*e2smrcv1.RanparameterTestingItem, 0),
	}

	innerRanParameterValue, err := CreateRanparameterValueInt(37)
	assert.NilError(t, err)

	ranParameterTestingCondition, err := CreateRanparameterTestingConditionComparison(CreateRanPChoiceComparisonStartsWith())
	assert.NilError(t, err)

	innerRanParameterType, err := CreateRanParameterTypeChoiceElementFalse(ranParameterTestingCondition)
	assert.NilError(t, err)
	innerRanParameterType.GetRanPChoiceElementFalse().SetRanParameterValue(innerRanParameterValue).SetLogicalOr(CreateLogicalOrFalse())

	innerRanParameterTestingItem, err := CreateRanparameterTestingItem(94, innerRanParameterType)
	assert.NilError(t, err)
	ranParameterTestingList.Value = append(ranParameterTestingList.Value, innerRanParameterTestingItem)

	ranParameterType, err := CreateRanParameterTypeChoiceList(ranParameterTestingList)
	assert.NilError(t, err)

	ranParameterTestingItem, err := CreateRanparameterTestingItem(69, ranParameterType)
	assert.NilError(t, err)

	ranParameterTesting := &e2smrcv1.RanparameterTesting{
		Value: make([]*e2smrcv1.RanparameterTestingItem, 0),
	}
	ranParameterTesting.Value = append(ranParameterTesting.Value, ranParameterTestingItem)

	actionItem2, err := CreateE2SmRcActionDefinitionFormat2Item(27, ricPolicyActionList)
	assert.NilError(t, err)
	actionItem2.SetRicPolicyConditionDefinition(ranParameterTesting)
	actionDefinitionItemFormat2List = append(actionDefinitionItemFormat2List, actionItem2)

	msg, err := CreateE2SmRcActionDefinitionFormat2(11, actionDefinitionItemFormat2List)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcActionDefinition(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcActionDefinition(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcActionDefinitionFormat3(t *testing.T) {

	ranParameterInsertIndicationList := make([]int64, 0)
	ranParameterInsertIndicationList = append(ranParameterInsertIndicationList, 72)
	ranParameterInsertIndicationList = append(ranParameterInsertIndicationList, 16)
	ranParameterInsertIndicationList = append(ranParameterInsertIndicationList, 29)

	msg, err := CreateE2SmRcActionDefinitionFormat3(5, 100, ranParameterInsertIndicationList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcActionDefinition(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcActionDefinition(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}
