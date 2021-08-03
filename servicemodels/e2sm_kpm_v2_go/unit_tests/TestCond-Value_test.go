// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerTCVint string = "00000000  00 01 15                                          |...|"
var refPerTCVenum string = "00000000  10 01 36                                          |..6|"
var refPerTCVbool string = "00000000  28                                                |(|"
var refPerTCVbs string = "00000000  30 16 d4 bc 08                                    |0....|"
var refPerTCVos string = "00000000  40 03 6f 6e 66                                    |@.onf|"
var refPerTCVps string = "00000000  50 03 4f 4e 46                                    |P.ONF|"

func Test_perEncodingTestCondValueInt(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueInt(21)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (Integer) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondValue{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("TestCondValue (Integer) PER - decoded\n%v", result)
}

func Test_perTestCondValueIntCompareBytes(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueInt(21)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (Integer) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCVint)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondValueEnum(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueEnum(54)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (Enumerator) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondValue{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("TestCondValue (Enumerator) PER - decoded\n%v", result)
}

func Test_perTestCondValueEnumCompareBytes(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueEnum(54)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (Enumerator) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCVenum)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondValueBool(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueBool(true)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (Boolean) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondValue{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("TestCondValue (Boolean) PER - decoded\n%v", result)
}

func Test_perTestCondValueBoolCompareBytes(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueBool(true)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (Boolean) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCVbool)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondValueBS(t *testing.T) {

	bs := &asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	testCondValue := pdubuilder.CreateTestCondValueBitS(bs)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (BitString) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondValue{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("TestCondValue (BitString) PER - decoded\n%v", result)
}

func Test_perTestCondValueBSCompareBytes(t *testing.T) {

	bs := &asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	testCondValue := pdubuilder.CreateTestCondValueBitS(bs)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (BitString) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCVbs)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondValueOS(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueOctS([]byte("onf"))

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (OctetString) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondValue{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("TestCondValue (OctetString) PER - decoded\n%v", result)
}

func Test_perTestCondValueOSCompareBytes(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValueOctS([]byte("onf"))

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (OctetString) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCVos)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingTestCondValuePS(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValuePrtS("ONF")

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (PrintableString) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.TestCondValue{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("TestCondValue (PrintableString) PER - decoded\n%v", result)
}

func Test_perTestCondValuePSCompareBytes(t *testing.T) {

	testCondValue := pdubuilder.CreateTestCondValuePrtS("ONF")

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*testCondValue, "valueExt")
	assert.NilError(t, err)
	t.Logf("TestCondValue (PrintableString) PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerTCVps)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
