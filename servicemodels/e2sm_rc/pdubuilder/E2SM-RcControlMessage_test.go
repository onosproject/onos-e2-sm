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

func TestE2SmRcControlMessageFormat1(t *testing.T) {

	ranParameterValue, err := CreateRanparameterValueBitS(&asn1.BitString{
		Value: []byte{0xFF, 0xFF, 0xFF, 0xFF},
		Len:   32,
	})
	assert.NilError(t, err)

	ranParameterValueType, err := CreateRanparameterValueTypeChoiceElementFalse(ranParameterValue)
	assert.NilError(t, err)

	item, err := CreateE2SmRcControlMessageFormat1Item(1, ranParameterValueType)
	assert.NilError(t, err)

	ranPList := make([]*e2smrcv1.E2SmRcControlMessageFormat1Item, 0)
	ranPList = append(ranPList, item)

	msg, err := CreateE2SmRcControlMessageFormat1(ranPList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcControlMessage(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcControlMessage(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcControlMessageFormat2(t *testing.T) {

	ranParameterValue, err := CreateRanparameterValueBitS(&asn1.BitString{
		Value: []byte{0xFA, 0xAB, 0xBC, 0xCF},
		Len:   32,
	})
	assert.NilError(t, err)

	ranParameterValueType, err := CreateRanparameterValueTypeChoiceElementFalse(ranParameterValue)
	assert.NilError(t, err)

	item, err := CreateE2SmRcControlMessageFormat1Item(1, ranParameterValueType)
	assert.NilError(t, err)

	ranPList := make([]*e2smrcv1.E2SmRcControlMessageFormat1Item, 0)
	ranPList = append(ranPList, item)

	controlActionItem, err := CreateE2SmRcControlMessageFormat2ControlActionItem(3, &e2smrcv1.E2SmRcControlMessageFormat1{
		RanPList: ranPList,
	})
	assert.NilError(t, err)

	controlActionList := make([]*e2smrcv1.E2SmRcControlMessageFormat2ControlActionItem, 0)
	controlActionList = append(controlActionList, controlActionItem)

	styleItem, err := CreateE2SmRcControlMessageFormat2StyleItem(153, controlActionList)
	assert.NilError(t, err)

	styleList := make([]*e2smrcv1.E2SmRcControlMessageFormat2StyleItem, 0)
	styleList = append(styleList, styleItem)

	msg, err := CreateE2SmRcControlMessageFormat2(styleList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcControlMessage(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcControlMessage(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}
