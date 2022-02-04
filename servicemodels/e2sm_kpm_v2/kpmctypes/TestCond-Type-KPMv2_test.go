// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeTestCondType(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeGBR()

	xer, err := xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 56, len(xer))
	t.Logf("TestCondType (GBR) XER\n%s", string(xer))

	testCondType = pdubuilder.CreateTestCondTypeAMBR()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 58, len(xer))
	t.Logf("TestCondType (AMBR) XER\n%s", string(xer))

	testCondType = pdubuilder.CreateTestCondTypeIsStat()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 62, len(xer))
	t.Logf("TestCondType (IsStat) XER\n%s", string(xer))

	testCondType = pdubuilder.CreateTestCondTypeIsCatM()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 62, len(xer))
	t.Logf("TestCondType (IsCatM) XER\n%s", string(xer))

	testCondType = pdubuilder.CreateTestCondTypeRSRP()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 58, len(xer))
	t.Logf("TestCondType (RSRP) XER\n%s", string(xer))

	testCondType = pdubuilder.CreateTestCondTypeRSRQ()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 58, len(xer))
	t.Logf("TestCondType (RSRQ) XER\n%s", string(xer))
}

func Test_xerDecodeTestCondType(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeGBR()

	xer, err := xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 56, len(xer))
	t.Logf("TestCondType (GBR) XER\n%s", string(xer))

	result, err := xerDecodeTestCondType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (GBR) XER - decoded\n%s", result)

	testCondType = pdubuilder.CreateTestCondTypeAMBR()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 58, len(xer))
	t.Logf("TestCondType (AMBR) XER\n%s", string(xer))

	result, err = xerDecodeTestCondType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (AMBR) XER - decoded\n%s", result)

	testCondType = pdubuilder.CreateTestCondTypeIsStat()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 62, len(xer))
	t.Logf("TestCondType (IsStat) XER\n%s", string(xer))

	result, err = xerDecodeTestCondType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (IsStat) XER - decoded\n%s", result)

	testCondType = pdubuilder.CreateTestCondTypeIsCatM()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 62, len(xer))
	t.Logf("TestCondType (IsCatM) XER\n%s", string(xer))

	result, err = xerDecodeTestCondType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (IsCatM) XER - decoded\n%s", result)

	testCondType = pdubuilder.CreateTestCondTypeRSRP()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 58, len(xer))
	t.Logf("TestCondType (RSRP) XER\n%s", string(xer))

	result, err = xerDecodeTestCondType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (RSRP) XER - decoded\n%s", result)

	testCondType = pdubuilder.CreateTestCondTypeRSRQ()

	xer, err = xerEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	//assert.Equal(t, 58, len(xer))
	t.Logf("TestCondType (RSRQ) XER\n%s", string(xer))

	result, err = xerDecodeTestCondType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (RSRQ) XER - decoded\n%s", result)
}

func Test_perEncodeTestCondType(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeGBR()

	per, err := perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (GBR) PER\n%v", hex.Dump(per))

	testCondType = pdubuilder.CreateTestCondTypeAMBR()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (AMBR) PER\n%v", hex.Dump(per))

	testCondType = pdubuilder.CreateTestCondTypeIsStat()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (IsStat) PER\n%v", hex.Dump(per))

	testCondType = pdubuilder.CreateTestCondTypeIsCatM()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (IsCatM) PER\n%v", hex.Dump(per))

	testCondType = pdubuilder.CreateTestCondTypeRSRP()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (RSRP) PER\n%v", hex.Dump(per))

	testCondType = pdubuilder.CreateTestCondTypeRSRQ()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (RSRQ) PER\n%v", hex.Dump(per))
}

func Test_perDecodeTestCondType(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeGBR()

	per, err := perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (GBR) PER\n%v", hex.Dump(per))

	result, err := perDecodeTestCondType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (GBR) PER - decoded\n%v", result)

	testCondType = pdubuilder.CreateTestCondTypeAMBR()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (AMBR) PER\n%v", hex.Dump(per))

	result, err = perDecodeTestCondType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (AMBR) PER - decoded\n%v", result)

	testCondType = pdubuilder.CreateTestCondTypeIsStat()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (IsStat) PER\n%v", hex.Dump(per))

	result, err = perDecodeTestCondType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (IsStat) PER - decoded\n%v", result)

	testCondType = pdubuilder.CreateTestCondTypeIsCatM()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (IsCatN) PER\n%v", hex.Dump(per))

	result, err = perDecodeTestCondType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (IsCatM) PER - decoded\n%v", result)

	testCondType = pdubuilder.CreateTestCondTypeRSRP()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (RSRP) PER\n%v", hex.Dump(per))

	result, err = perDecodeTestCondType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (RSRP) PER - decoded\n%v", result)

	testCondType = pdubuilder.CreateTestCondTypeRSRQ()

	per, err = perEncodeTestCondType(testCondType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TestCondType (RSRQ) PER\n%v", hex.Dump(per))

	result, err = perDecodeTestCondType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestCondType (RSRQ) PER - decoded\n%v", result)
}
