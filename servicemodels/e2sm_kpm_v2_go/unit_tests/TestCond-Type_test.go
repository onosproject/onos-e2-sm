// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerTCTgbr = "00000000  00                                                |.|"
var refPerTCTambr = "00000000  10                                                |.|"
var refPerTCTisstat = "00000000  20                                                | |"
var refPerTCTiscatm = "00000000  30                                                |0|"
var refPerTCTrsrp = "00000000  40                                                |@|"
var refPerTCTrsrq = "00000000  50                                                |P|"

func Test_perEncodingTestCondTypeGBR(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeGBR()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (GBR) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("TestCondType (GBR) PER - decoded\n%v", &result)
	assert.Equal(t, testCondType.GetGBr().Number(), testCondType.GetGBr().Number())
}

func Test_perTestCondTypeGBRCompareBytes(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeGBR()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (GBR) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCTgbr)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondTypeAMBR(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeAMBR()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (AMBR) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("TestCondType (AMBR) PER - decoded\n%v", &result)
	assert.Equal(t, testCondType.GetAMbr().Number(), testCondType.GetAMbr().Number())
}

func Test_perTestCondTypeAMBRCompareBytes(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeAMBR()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (AMBR) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCTambr)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondTypeIsStat(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeIsStat()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (IsStat) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("TestCondType (IsStat) PER - decoded\n%v", &result)
	assert.Equal(t, testCondType.GetIsStat().Number(), testCondType.GetIsStat().Number())
}

func Test_perTestCondTypeIsStatCompareBytes(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeIsStat()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (IsStat) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCTisstat)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondTypeIsCatM(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeIsCatM()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (IsCatM) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("TestCondType (IsCatM) PER - decoded\n%v", &result)
	assert.Equal(t, testCondType.GetIsCatM().Number(), testCondType.GetIsCatM().Number())
}

func Test_perTestCondTypeIsCatMCompareBytes(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeIsCatM()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (IsCatM) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCTiscatm)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondTypeRSRP(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeRSRP()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (RSRP) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("TestCondType (RSRP) PER - decoded\n%v", &result)
	assert.Equal(t, testCondType.GetRSrp().Number(), testCondType.GetRSrp().Number())
}

func Test_perTestCondTypeRSRPCompareBytes(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeRSRP()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (RSRP) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCTrsrp)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondTypeRSRQ(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeRSRQ()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (RSRQ) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("TestCondType (RSRQ) PER - decoded\n%v", &result)
	assert.Equal(t, testCondType.GetRSrq().Number(), testCondType.GetRSrq().Number())
}

func Test_perTestCondTypeRSRQCompareBytes(t *testing.T) {

	testCondType := pdubuilder.CreateTestCondTypeRSRQ()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(testCondType, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondType (RSRQ) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCTrsrq)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
