// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func createRanparameterValueValueInteger() *e2sm_rc_pre_ies.RanparameterValue {
	return &e2sm_rc_pre_ies.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValueInt{
			ValueInt: 44,
		},
	}
}

func createRanparameterValueValueEnum() *e2sm_rc_pre_ies.RanparameterValue {
	return &e2sm_rc_pre_ies.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValueEnum{
			ValueEnum: 98,
		},
	}
}

//func createRanparameterValueValueOctS() *e2sm_rc_pre_ies.RanparameterValue {
//	return &e2sm_rc_pre_ies.RanparameterValue{
//		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValueOctS{
//			ValueOctS: "onf",
//		},
//	}
//}
//
//func createRanparameterValueValuePrtS() *e2sm_rc_pre_ies.RanparameterValue {
//	return &e2sm_rc_pre_ies.RanparameterValue{
//		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValuePrtS{
//			ValuePrtS: "onf",
//		},
//	}
//}
//
//func createRanparameterValueValueBitS() *e2sm_rc_pre_ies.RanparameterValue {
//	return &e2sm_rc_pre_ies.RanparameterValue{
//		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValueBitS{
//			ValueBitS: &e2sm_rc_pre_ies.BitString{
//				Value: 0x9bcd4,
//				Len:   22,
//			},
//		},
//	}
//}

func Test_xerEncodeRanparameterValue(t *testing.T) {

	rpv := createRanparameterValueValueInteger()

	xer, err := xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 71, len(xer))
	t.Logf("RANparameterValue (Integer) XER\n%s", string(xer))

	rpv = createRanparameterValueValueEnum()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 73, len(xer))
	t.Logf("RANparameterValue (Enum) XER\n%s", string(xer))

	//rpv = createRanparameterValueValueOctS()
	//
	//xer, err = xerEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	//t.Logf("RANparameterValue (OctetString) XER\n%s", string(xer))
	//
	//rpv = createRanparameterValueValuePrtS()
	//
	//xer, err = xerEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	//t.Logf("RANparameterValue (PrintableString) XER\n%s", string(xer))

	//rpv = createRanparameterValueValueBitS()
	//
	//xer, err = xerEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	//t.Logf("RANparameterValue (BitString) XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementRecordItem(t *testing.T) {

	rpv := createRanparameterValueValueInteger()

	xer, err := xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 71, len(xer))
	t.Logf("RANparameterValue (Integer) XER\n%s", string(xer))

	result, err := xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Integer) XER - decoded\n%s", result)

	rpv = createRanparameterValueValueEnum()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 73, len(xer))
	t.Logf("RANparameterValue (Enum) XER\n%s", string(xer))

	result, err = xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Enum) XER - decoded\n%s", result)

	//rpv = createRanparameterValueValueOctS()
	//
	//xer, err = xerEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	//t.Logf("RANparameterValue (OctetString) XER\n%s", string(xer))
	//
	//result, err = xerDecodeRanparameterValue(xer)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	//t.Logf("RANparameterValue (OctetString) XER - decoded\n%s", result)
	//
	//rpv = createRanparameterValueValuePrtS()
	//
	//xer, err = xerEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	//t.Logf("RANparameterValue (PrintableString) XER\n%s", string(xer))
	//
	//result, err = xerDecodeRanparameterValue(xer)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	//t.Logf("RANparameterValue (PrintableString) XER - decoded\n%s", result)

	//rpv = createRanparameterValueValueBitS()
	//
	//xer, err = xerEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	//t.Logf("RANparameterValue (BitString) XER\n%s", string(xer))
	//
	//result, err = xerDecodeRanparameterValue(xer)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
}

func Test_perEncodeMeasurementRecordItem(t *testing.T) {

	rpv := createRanparameterValueValueInteger()

	per, err := perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("RANparameterValue (Integer) PER\n%s", string(per))

	rpv = createRanparameterValueValueEnum()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("RANparameterValue (Enum) PER\n%s", string(per))

	//rpv = createRanparameterValueValueOctS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	//t.Logf("RANparameterValue (OctetString) PER\n%s", string(per))
	//
	//rpv = createRanparameterValueValuePrtS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	//t.Logf("RANparameterValue (PrintableString) PER\n%s", string(per))

	//rpv = createRanparameterValueValueBitS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	//t.Logf("RANparameterValue (BitString) PER\n%s", string(per))
}

func Test_perDecodeMeasurementRecordItem(t *testing.T) {

	rpv := createRanparameterValueValueInteger()

	per, err := perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("RANparameterValue (Integer) PER\n%s", string(per))

	result, err := perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Integer) PER - decoded\n%s", result)

	rpv = createRanparameterValueValueEnum()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("RANparameterValue (Enum) PER\n%s", string(per))

	result, err = perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Enum) PER - decoded\n%s", result)

	//rpv = createRanparameterValueValueOctS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	//t.Logf("RANparameterValue (OctetString) PER\n%s", string(per))
	//
	//result, err = perDecodeRanparameterValue(per)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	//
	//rpv = createRanparameterValueValuePrtS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	//t.Logf("RANparameterValue (PrintableString) PER\n%s", string(per))
	//
	//result, err = perDecodeRanparameterValue(per)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)

	//rpv = createRanparameterValueValueBitS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	//t.Logf("RANparameterValue (BitString) PER\n%s", string(per))
	//
	//result, err = perDecodeRanparameterValue(per)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
}
