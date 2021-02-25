// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func createRanparameterValue_ValueInteger() *e2sm_rc_pre_ies.RanparameterValue {
	return &e2sm_rc_pre_ies.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValueInt{
			ValueInt: 44,
		},
	}
}

func createRanparameterValue_ValueEnum() *e2sm_rc_pre_ies.RanparameterValue {
	return &e2sm_rc_pre_ies.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValueEnum{
			ValueEnum: 98,
		},
	}
}

func createRanparameterValue_ValueOctS() *e2sm_rc_pre_ies.RanparameterValue {
	return &e2sm_rc_pre_ies.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValueOctS{
			ValueOctS: "onf",
		},
	}
}

func createRanparameterValue_ValuePrtS() *e2sm_rc_pre_ies.RanparameterValue {
	return &e2sm_rc_pre_ies.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValuePrtS{
			ValuePrtS: "onf",
		},
	}
}

func createRanparameterValue_ValueBitS() *e2sm_rc_pre_ies.RanparameterValue {
	return &e2sm_rc_pre_ies.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_ies.RanparameterValue_ValueBitS{
			ValueBitS: &e2sm_rc_pre_ies.BitString{
				Value: 0x9bcd4,
				Len:   22,
			},
		},
	}
}

func Test_xerEncodeRanparameterValue(t *testing.T) {

	rpv := createRanparameterValue_ValueInteger()

	xer, err := xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 71, len(xer))
	t.Logf("RANparameterValue (Integer) XER\n%s", string(xer))

	rpv = createRanparameterValue_ValueEnum()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 83, len(xer))
	t.Logf("RANparameterValue (Enum) XER\n%s", string(xer))

	//rpv = createRanparameterValue_ValueOctS()
	//
	//xer, err = xerEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	//t.Logf("RANparameterValue (OctetString) XER\n%s", string(xer))
	//
	//rpv = createRanparameterValue_ValuePrtS()
	//
	//xer, err = xerEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	//t.Logf("RANparameterValue (PrintableString) XER\n%s", string(xer))

	//rpv = createRanparameterValue_ValueBitS()
	//
	//xer, err = xerEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	//t.Logf("RANparameterValue (BitString) XER\n%s", string(xer))
}

func Test_xerDecodeMeasurementRecordItem(t *testing.T) {

	rpv := createRanparameterValue_ValueInteger()

	xer, err := xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 71, len(xer))
	t.Logf("RANparameterValue (Integer) XER\n%s", string(xer))

	result, err := xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Integer) XER - decoded\n%s", result)

	rpv = createRanparameterValue_ValueEnum()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 83, len(xer))
	t.Logf("RANparameterValue (Enum) XER\n%s", string(xer))

	result, err = xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Enum) XER - decoded\n%s", result)

	//rpv = createRanparameterValue_ValueOctS()
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
	//rpv = createRanparameterValue_ValuePrtS()
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

	//rpv = createRanparameterValue_ValueBitS()
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

	rpv := createRanparameterValue_ValueInteger()

	per, err := perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 8, len(per))
	t.Logf("RANparameterValue (Integer) PER\n%s", string(per))

	rpv = createRanparameterValue_ValueEnum()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("RANparameterValue (Enum) PER\n%s", string(per))

	//rpv = createRanparameterValue_ValueOctS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	//t.Logf("RANparameterValue (OctetString) PER\n%s", string(per))
	//
	//rpv = createRanparameterValue_ValuePrtS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	//t.Logf("RANparameterValue (PrintableString) PER\n%s", string(per))

	//rpv = createRanparameterValue_ValueBitS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 4, len(per))
	//t.Logf("RANparameterValue (BitString) PER\n%s", string(per))
}

func Test_perDecodeMeasurementRecordItem(t *testing.T) {

	rpv := createRanparameterValue_ValueInteger()

	per, err := perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("RANparameterValue (Integer) PER\n%s", string(per))

	result, err := perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Integer) PER - decoded\n%s", result)

	rpv = createRanparameterValue_ValueEnum()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 7, len(per))
	t.Logf("RANparameterValue (Enum) PER\n%s", string(per))

	result, err = perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Enum) PER - decoded\n%s", result)

	//rpv = createRanparameterValue_ValueOctS()
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
	//rpv = createRanparameterValue_ValuePrtS()
	//
	//per, err = perEncodeRanparameterValue(rpv)
	//assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	//t.Logf("RANparameterValue (PrintableString) PER\n%s", string(per))
	//
	//result, err = perDecodeRanparameterValue(per)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)

	//rpv = createRanparameterValue_ValueBitS()
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
