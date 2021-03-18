// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeTestCondValue(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueInt(21)

	xer, err := xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 63, len(xer))
	t.Logf("TestCondValue (Integer) XER\n%s", string(xer))

	testCondValue = pdubuilder.CreateTestCondValueEnum(54)

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 65, len(xer))
	t.Logf("TestCondValue (Enum) XER\n%s", string(xer))

	testCondValue = pdubuilder.CreateTestCondValueBool(true)

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 70, len(xer))
	t.Logf("TestCondValue (Boolean) XER\n%s", string(xer))

	bs := &e2sm_kpm_v2.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	testCondValue = pdubuilder.CreateTestCondValueBitS(bs)

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 99, len(xer))
	t.Logf("TestCondValue (BitString) XER\n%s", string(xer))

	testCondValue = pdubuilder.CreateTestCondValueOctS("onf")

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 71, len(xer))
	t.Logf("TestCondValue (OctetString) XER\n%s", string(xer))

	testCondValue = pdubuilder.CreateTestCondValuePrtS("ONF")

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 66, len(xer))
	t.Logf("TestCondValue (PrintableString) XER\n%s", string(xer))
}

func Test_xerDecodeTestCondValue(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueInt(21)

	xer, err := xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 63, len(xer))
	t.Logf("TestCondValue (Integer) XER\n%s", string(xer))

	result, err := xerDecodeTestCondValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (Integer) XER - decoded\n%s", result)

	testCondValue = pdubuilder.CreateTestCondValueEnum(54)

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 65, len(xer))
	t.Logf("TestCondValue (Enum) XER\n%s", string(xer))

	result, err = xerDecodeTestCondValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (Enum) XER - decoded\n%s", result)

	testCondValue = pdubuilder.CreateTestCondValueBool(true)

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 70, len(xer))
	t.Logf("TestCondValue (Boolean) XER\n%s", string(xer))

	result, err = xerDecodeTestCondValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (Boolean) XER - decoded\n%s", result)

	bs := &e2sm_kpm_v2.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	testCondValue = pdubuilder.CreateTestCondValueBitS(bs)

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 99, len(xer))
	t.Logf("TestCondValue (BitString) XER\n%s", string(xer))

	result, err = xerDecodeTestCondValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (BitString) XER - decoded\n%s", result)

	testCondValue = pdubuilder.CreateTestCondValueOctS("onf")

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 71, len(xer))
	t.Logf("TestCondValue (OctetString) XER\n%s", string(xer))

	result, err = xerDecodeTestCondValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (OctetString) XER - decoded\n%s", result)

	testCondValue = pdubuilder.CreateTestCondValuePrtS("ONF")

	xer, err = xerEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 66, len(xer))
	t.Logf("TestCondValue (PrintableString) XER\n%s", string(xer))

	result, err = xerDecodeTestCondValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (PrintableString) XER - decoded\n%s", result)
}

func Test_perEncodeTestCondValue(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueInt(21)

	per, err := perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("TestCondValue (Integer) PER\n%s", string(per))

	testCondValue = pdubuilder.CreateTestCondValueEnum(54)

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("TestCondValue (Enum) PER\n%s", string(per))

	testCondValue = pdubuilder.CreateTestCondValueBool(true)

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondValue (Boolean) PER\n%s", string(per))

	bs := &e2sm_kpm_v2.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	testCondValue = pdubuilder.CreateTestCondValueBitS(bs)

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("TestCondValue (BitString) PER\n%s", string(per))

	testCondValue = pdubuilder.CreateTestCondValueOctS("onf")

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("TestCondValue (OctetString) PER\n%s", string(per))

	testCondValue = pdubuilder.CreateTestCondValuePrtS("ONF")

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("TestCondValue (PrintableString) PER\n%s", string(per))
}

func Test_perDecodeTestCondValue(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueInt(21)

	per, err := perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("TestCondValue (Integer) PER\n%s", string(per))

	result, err := perDecodeTestCondValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (Integer) PER - decoded\n%v", result)

	testCondValue = pdubuilder.CreateTestCondValueEnum(54)

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("TestCondValue (Enum) PER\n%s", string(per))

	result, err = perDecodeTestCondValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (Enum) PER - decoded\n%v", result)

	testCondValue = pdubuilder.CreateTestCondValueBool(false)

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondValue (Boolean) PER\n%s", string(per))

	result, err = perDecodeTestCondValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (Boolean) PER - decoded\n%v", result)

	bs := &e2sm_kpm_v2.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	testCondValue = pdubuilder.CreateTestCondValueBitS(bs)

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("TestCondValue (BitString) PER\n%s", string(per))

	result, err = perDecodeTestCondValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (BitString) PER - decoded\n%v", result)

	testCondValue = pdubuilder.CreateTestCondValueOctS("onf")

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("TestCondValue (OctetString) PER\n%s", string(per))

	result, err = perDecodeTestCondValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (OctetString) PER - decoded\n%v", result)

	testCondValue = pdubuilder.CreateTestCondValuePrtS("ONF")

	per, err = perEncodeTestCondValue(testCondValue)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("TestCondValue (PrintableString) PER\n%s", string(per))

	result, err = perDecodeTestCondValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondValue (PrintableString) PER - decoded\n%v", result)
}
