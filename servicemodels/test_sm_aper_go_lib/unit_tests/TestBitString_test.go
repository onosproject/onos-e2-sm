// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package unit_tests

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/testsmctypes"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/asn1/testsm"
	"gotest.tools/assert"
	"math"
	"math/rand"
	"testing"
	"time"
)

func generateBitStringOfSize(n uint32) (*asn1.BitString, error) {

	bs := &asn1.BitString{
		Value: nil,
		Len:   n,
	}

	if n == 0 {
		bs = nil
	} else {
		var val []byte

		numBytes := int(math.Ceil(float64(n) / 8.0))

		var rn []rune
		for i := 1; i <= numBytes; i++ {
			rn = append(rn, createRandomAsciiChar())
		}
		val = []byte(string(rn))

		bs.Value = val

		_, err := bs.TruncateValue()
		if err != nil {
			return nil, err
		}
	}

	return bs, nil
}

func createTestBitString() (*test_sm_ies.TestBitString, error) {

	testSM := new(test_sm_ies.TestBitString)

	// Setting minimum size of BitString to 1 to avoid the case of 0-long BitString
	bs1, err := generateBitStringOfSize(uint32(rand.Intn(1000-1) + 1))
	if err != nil {
		return nil, err
	}
	testSM.AttrBs1 = bs1

	bs2, err := generateBitStringOfSize(uint32(20))
	if err != nil {
		return nil, err
	}
	testSM.AttrBs2 = bs2

	var bs3 *asn1.BitString
	//need to indicate whether next value is extensible or not
	ext := rand.Intn(1)
	if ext == 1 {
		bs3, err = generateBitStringOfSize(uint32(20) + uint32(rand.Intn(1000)))
		if err != nil {
			return nil, err
		}
		testSM.AttrBs3 = bs3
	} else {
		bs3, err = generateBitStringOfSize(uint32(20))
		if err != nil {
			return nil, err
		}
		testSM.AttrBs3 = bs3
	}

	// BitString of size 0 can't be created
	// Setting minimum size of BitString to 1 to avoid the case of 0-long BitString
	bs4, err := generateBitStringOfSize(uint32(rand.Intn(18-1) + 1))
	if err != nil {
		return nil, err
	}
	testSM.AttrBs4 = bs4

	bs5, err := generateBitStringOfSize(uint32(rand.Intn(32-22) + 22))
	if err != nil {
		return nil, err
	}
	testSM.AttrBs5 = bs5

	var bs6 *asn1.BitString
	//need to indicate whether next value is extensible or not
	ext = rand.Intn(1)
	if ext == 1 {
		bs6, err = generateBitStringOfSize(uint32(rand.Intn(32-28)+28) + uint32(rand.Intn(1000)))
		if err != nil {
			return nil, err
		}
		testSM.AttrBs6 = bs6
	} else {
		bs6, err = generateBitStringOfSize(uint32(rand.Intn(32-28) + 28))
		if err != nil {
			return nil, err
		}
		testSM.AttrBs6 = bs6
	}

	var bs7 *asn1.BitString
	//need to indicate whether next value is extensible or not
	optional := rand.Intn(1)
	if optional == 1 {
		bs7, err = generateBitStringOfSize(uint32(rand.Intn(36-22) + 22))
		if err != nil {
			return nil, err
		}
		testSM.AttrBs6 = bs7
	}

	return testSM, nil
}

func TestBitString(t *testing.T) {

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < 1000; i++ {

		testSM, err := createTestBitString()
		assert.NilError(t, err)

		t.Logf("Testing Test-BitString with following values:\n"+
			"%x, \n%x, \n%x, \n%x, \n%x, \n%x, \n%x", testSM.GetAttrBs1().GetValue(), testSM.GetAttrBs2().GetValue(),
			testSM.GetAttrBs3().GetValue(), testSM.GetAttrBs4().GetValue(), testSM.GetAttrBs5().GetValue(),
			testSM.GetAttrBs6().GetValue(), testSM.GetAttrBs7().GetValue())

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestBitString(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM, test_sm_ies.Choicemap, nil)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-BitString (with randomness) was successfully finished")
}

func TestBitStringExtensible(t *testing.T) {

	testBS, err := createTestBitString()
	assert.NilError(t, err)

	// setting BS6 to be explicitly extended (one extra bit)
	testBS.AttrBs6 = &asn1.BitString{
		Value: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x80},
		Len:   33,
	}

	// Generating APER with reference CGo approach
	perRef, err := testsmctypes.PerEncodeTestBitString(testBS)
	assert.NilError(t, err)
	// Generating APER bytes with Go APER lib
	per, err := aper.Marshal(testBS, test_sm_ies.Choicemap, nil)
	assert.NilError(t, err)

	//Comparing bytes against each other
	assert.DeepEqual(t, per, perRef)

	// verifying that we're able to decode the message back
	result := &testsm.TestBitString{}
	err = aper.Unmarshal(per, result, test_sm_ies.Choicemap, nil)
	assert.NilError(t, err)
	assert.Equal(t, testBS.String(), result.String())

	testBS1, err1 := createTestBitString()
	assert.NilError(t, err1)

	// setting BS6 to be explicitly extended (one extra bit)
	testBS1.AttrBs6 = &asn1.BitString{
		Value: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		Len:   40,
	}

	// Generating APER with reference CGo approach
	perRef1, err1 := testsmctypes.PerEncodeTestBitString(testBS1)
	assert.NilError(t, err1)
	// Generating APER bytes with Go APER lib
	per1, err1 := aper.Marshal(testBS1, test_sm_ies.Choicemap, nil)
	assert.NilError(t, err1)

	//Comparing bytes against each other
	assert.DeepEqual(t, per1, perRef1)

	// verifying that we're able to decode the message back
	result1 := &testsm.TestBitString{}
	err1 = aper.Unmarshal(per1, result1, test_sm_ies.Choicemap, nil)
	assert.NilError(t, err1)
	assert.Equal(t, testBS1.String(), result1.String())
}
