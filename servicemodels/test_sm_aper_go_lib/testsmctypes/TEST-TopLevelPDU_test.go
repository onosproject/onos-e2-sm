// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func createTestTopLevelPDU() (*test_sm_ies.TestTopLevelPdu, error) {

	opt6, err := createTestListExtensible3MsgFull()
	if err != nil {
		return nil, err
	}
	testTopLevelPdu := test_sm_ies.TestTopLevelPdu{
		Opt1: &test_sm_ies.TestUnconstrainedInt{
			AttrUciA: -153,
			AttrUciB: 21,
		},
		Opt2: &test_sm_ies.TestConstrainedReal{
			AttrCrA: 99.97,
			AttrCrB: 11.20,
			AttrCrC: -153.0,
			AttrCrD: 20.0,
			AttrCrE: 10.0,
			AttrCrF: 10.0,
		},
		Opt3: &test_sm_ies.TestNestedChoice{
			TestNestedChoice: &test_sm_ies.TestNestedChoice_Option3{
				Option3: &test_sm_ies.ConstrainedChoice4{
					ConstrainedChoice4: &test_sm_ies.ConstrainedChoice4_ConstrainedChoice4A{
						ConstrainedChoice4A: 138,
					},
				},
			},
		},
		Opt4: &test_sm_ies.TestBitString{
			AttrBs1: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x4F, 0xE0, 0x00},
				Len:   35,
			},
			AttrBs2: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x40},
				Len:   20,
			},
			AttrBs3: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x40},
				Len:   20,
			},
			//AttrBs4: &asn1.BitString{  // ToDo - there should be no Octet-Alignment - BitStrings is less than 16!!
			//	Value: []byte{0x01},
			//	Len:   1,
			//},
			AttrBs4: &asn1.BitString{
				Value: []byte{0xCC, 0xC0},
				Len:   16,
			},
			AttrBs5: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x4F, 0x00},
				Len:   32,
			},
			AttrBs6: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x4F, 0x00},
				Len:   32,
			},
			AttrBs7: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x4F, 0xE0, 0x00},
				Len:   35,
			},
		},
		Opt5: &test_sm_ies.TestOctetString{
			AttrOs1: []byte{0x12, 0x34, 0x56, 0xA4},
			AttrOs2: []byte{0xFF, 0xFF},
			AttrOs3: []byte{0xFF, 0xFF},
			AttrOs4: []byte{0xBC, 0x7D, 0xA1},             // It doesn't like anything less than 3 bytes...
			AttrOs5: []byte{0xDE, 0xC7, 0x23},             // It doesn't like anything less than 3 bytes again...
			AttrOs6: []byte{0x02, 0x31, 0xF6, 0x7D, 0x19}, // It doesn't like anything less than 3 bytes again... aand again...
			//AttrOs7: []byte{0x21, 0x44, 0xA8, 0xDF, 0x11},
		},
		Opt6: opt6,
		Opt7: test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM6,
	}

	return &testTopLevelPdu, nil
}

func createTestTopLevelPDUxclOptnl() (*test_sm_ies.TestTopLevelPdu, error) {

	opt6, err := createTestListExtensible3MsgFull()
	if err != nil {
		return nil, err
	}
	testTopLevelPdu := test_sm_ies.TestTopLevelPdu{
		Opt1: &test_sm_ies.TestUnconstrainedInt{
			AttrUciA: -153,
			AttrUciB: 21,
		},
		//Opt2: &test_sm_ies.TestConstrainedReal{
		//	AttrCrA: 99.97,
		//	AttrCrB: 11.20,
		//	AttrCrC: -153.0,
		//	AttrCrD: 20.0,
		//	AttrCrE: 10.0,
		//	AttrCrF: 10.0,
		//},
		Opt3: &test_sm_ies.TestNestedChoice{
			TestNestedChoice: &test_sm_ies.TestNestedChoice_Option3{
				Option3: &test_sm_ies.ConstrainedChoice4{
					ConstrainedChoice4: &test_sm_ies.ConstrainedChoice4_ConstrainedChoice4A{
						ConstrainedChoice4A: 138,
					},
				},
			},
		},
		Opt4: &test_sm_ies.TestBitString{
			AttrBs1: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x4F, 0xE0, 0x00},
				Len:   35,
			},
			AttrBs2: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x40},
				Len:   20,
			},
			AttrBs3: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x40},
				Len:   20,
			},
			//AttrBs4: &asn1.BitString{  // ToDo - there should be no Octet-Alignment - BitStrings is less than 16!!
			//	Value: []byte{0x01},
			//	Len:   1,
			//},
			AttrBs4: &asn1.BitString{
				Value: []byte{0xCC, 0xC0},
				Len:   16,
			},
			AttrBs5: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x4F, 0x00},
				Len:   32,
			},
			AttrBs6: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x4F, 0x00},
				Len:   32,
			},
			AttrBs7: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x4F, 0xE0, 0x00},
				Len:   35,
			},
		},
		//Opt5: &test_sm_ies.TestOctetString{
		//	AttrOs1: []byte{0x12, 0x34, 0x56, 0xA4},
		//	AttrOs2: []byte{0xFF, 0xFF},
		//	AttrOs3: []byte{0xFF, 0xFF},
		//	AttrOs4: []byte{0xBC, 0x7D, 0xA1},             // It doesn't like anything less than 3 bytes...
		//	AttrOs5: []byte{0xDE, 0xC7, 0x23},             // It doesn't like anything less than 3 bytes again...
		//	AttrOs6: []byte{0x02, 0x31, 0xF6, 0x7D, 0x19}, // It doesn't like anything less than 3 bytes again... aand again...
		//	//AttrOs7: []byte{0x21, 0x44, 0xA8, 0xDF, 0x11},
		//},
		Opt6: opt6,
		Opt7: test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM6,
	}

	return &testTopLevelPdu, nil
}

func Test_xerEncodingTestTopLevelPdu(t *testing.T) {

	testTopLevelPDU, err := createTestTopLevelPDU()
	assert.NilError(t, err, "Error creating TestTopLevelPDU PDU")

	xer, err := xerEncodeTestTopLevelPDU(testTopLevelPDU)
	assert.NilError(t, err)
	t.Logf("TestTopLevelPDU XER\n%s", string(xer))

	result, err := xerDecodeTestTopLevelPDU(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestTopLevelPDU XER - decoded\n%v", result)
	assert.Equal(t, testTopLevelPDU.GetOpt1(), result.GetOpt1())
	assert.Equal(t, testTopLevelPDU.GetOpt2(), result.GetOpt2())
	assert.Equal(t, testTopLevelPDU.GetOpt3(), result.GetOpt3())
	assert.Equal(t, testTopLevelPDU.GetOpt4(), result.GetOpt4())
	assert.Equal(t, testTopLevelPDU.GetOpt5(), result.GetOpt5())
	assert.Equal(t, testTopLevelPDU.GetOpt6(), result.GetOpt6())
	assert.Equal(t, testTopLevelPDU.GetOpt7(), result.GetOpt7())
}

func Test_perEncodingTestTopLevelPdu(t *testing.T) {

	testTopLevelPDU, err := createTestTopLevelPDU()
	assert.NilError(t, err, "Error creating TestTopLevelPDU PDU")

	per, err := perEncodeTestTopLevelPDU(testTopLevelPDU)
	assert.NilError(t, err)
	t.Logf("TestTopLevelPDU PER\n%v", hex.Dump(per))

	result, err := perDecodeTestTopLevelPDU(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestTopLevelPDU PER - decoded\n%v", result)
	assert.Equal(t, testTopLevelPDU.GetOpt1(), result.GetOpt1())
	assert.Equal(t, testTopLevelPDU.GetOpt2(), result.GetOpt2())
	assert.Equal(t, testTopLevelPDU.GetOpt3(), result.GetOpt3())
	assert.Equal(t, testTopLevelPDU.GetOpt4(), result.GetOpt4())
	assert.Equal(t, testTopLevelPDU.GetOpt5(), result.GetOpt5())
	assert.Equal(t, testTopLevelPDU.GetOpt6(), result.GetOpt6())
	assert.Equal(t, testTopLevelPDU.GetOpt7(), result.GetOpt7())
}

func Test_xerEncodingTestTopLevelPduExcludeOptional(t *testing.T) {

	testTopLevelPDU, err := createTestTopLevelPDUxclOptnl()
	assert.NilError(t, err, "Error creating TestTopLevelPDU PDU")

	xer, err := xerEncodeTestTopLevelPDU(testTopLevelPDU)
	assert.NilError(t, err)
	t.Logf("TestTopLevelPDU (w/o optional) XER\n%s", string(xer))

	result, err := xerDecodeTestTopLevelPDU(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestTopLevelPDU (w/o optional) XER - decoded\n%v", result)
	assert.Equal(t, testTopLevelPDU.GetOpt1(), result.GetOpt1())
	//assert.Equal(t, testTopLevelPDU.GetOpt2(), result.GetOpt2())
	assert.Equal(t, testTopLevelPDU.GetOpt3(), result.GetOpt3())
	assert.Equal(t, testTopLevelPDU.GetOpt4(), result.GetOpt4())
	//assert.Equal(t, testTopLevelPDU.GetOpt5(), result.GetOpt5())
	assert.Equal(t, testTopLevelPDU.GetOpt6(), result.GetOpt6())
	assert.Equal(t, testTopLevelPDU.GetOpt7(), result.GetOpt7())
}

func Test_perEncodingTestTopLevelPduExcludeOptional(t *testing.T) {

	testTopLevelPDU, err := createTestTopLevelPDUxclOptnl()
	assert.NilError(t, err, "Error creating TestTopLevelPDU PDU")

	per, err := perEncodeTestTopLevelPDU(testTopLevelPDU)
	assert.NilError(t, err)
	t.Logf("TestTopLevelPDU (w/o optional) PER\n%v", hex.Dump(per))

	result, err := perDecodeTestTopLevelPDU(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TestTopLevelPDU (w/o optional) PER - decoded\n%v", result)
	assert.Equal(t, testTopLevelPDU.GetOpt1(), result.GetOpt1())
	//assert.Equal(t, testTopLevelPDU.GetOpt2(), result.GetOpt2())
	assert.Equal(t, testTopLevelPDU.GetOpt3(), result.GetOpt3())
	assert.Equal(t, testTopLevelPDU.GetOpt4(), result.GetOpt4())
	//assert.Equal(t, testTopLevelPDU.GetOpt5(), result.GetOpt5())
	assert.Equal(t, testTopLevelPDU.GetOpt6(), result.GetOpt6())
	assert.Equal(t, testTopLevelPDU.GetOpt7(), result.GetOpt7())
}
