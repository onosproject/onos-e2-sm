// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package unit_tests

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/testsmctypes"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"math/rand"
	"testing"
	"time"
)

func createTestOctetString() *test_sm_ies.TestOctetString {

	testSM := new(test_sm_ies.TestOctetString)

	var ps1 []rune
	// Limiting it on 1000 on purpose, otherwise test scenario would take too long
	for j := 1; j < rand.Intn(1000); j++ {
		ps1 = append(ps1, createRandomAsciiChar())
	}
	testSM.AttrOs1 = []byte(string(ps1))

	var ps2 []rune
	for j := 1; j <= 2; j++ {
		ps2 = append(ps2, createRandomAsciiChar())
	}
	testSM.AttrOs2 = []byte(string(ps2))

	var ps3 []rune
	//need to indicate whether next value is extensible or not
	ext := rand.Intn(1)
	if ext == 1 {
		for j := 1; j <= 2+rand.Int(); j++ {
			ps3 = append(ps3, createRandomAsciiChar())
		}
		testSM.AttrOs3 = []byte(string(ps3))
	} else {
		for j := 1; j <= 2; j++ {
			ps3 = append(ps3, createRandomAsciiChar())
		}
		testSM.AttrOs3 = []byte(string(ps3))
	}

	var ps4 []rune
	//for j := 1; j <= rand.Intn(3); j++ {
	//For some reason, testsmctypes encoder doesn't like anything less than 3 characters
	for j := 1; j <= 3; j++ {
		ps4 = append(ps4, createRandomAsciiChar())
	}
	testSM.AttrOs4 = []byte(string(ps4))

	var ps5 []rune
	//for j := 2; j <= 2 + rand.Intn(1); j++ {
	//For some reason, testsmctypes encoder doesn't like anything less than 3 characters
	for j := 1; j <= 3; j++ {
		ps5 = append(ps5, createRandomAsciiChar())
	}
	testSM.AttrOs5 = []byte(string(ps5))

	var ps6 []rune
	//need to indicate whether next value is extensible or not
	ext = rand.Intn(1)
	if ext == 1 {
		// Limiting it on 1000 on purpose, otherwise test scenario would take too long
		for j := 1; j <= 3+rand.Intn(1000); j++ {
			ps6 = append(ps6, createRandomAsciiChar())
		}
		testSM.AttrOs6 = []byte(string(ps6))
	} else {
		//For some reason, testsmctypes encoder doesn't like anything less than 3 characters
		for j := 1; j <= 3; j++ {
			ps6 = append(ps6, createRandomAsciiChar())
		}
		testSM.AttrOs6 = []byte(string(ps6))
	}

	var ps7 []rune
	optional := rand.Intn(1)
	if optional == 1 {
		for j := 2; j <= 2+rand.Intn(4); j++ {
			ps7 = append(ps7, createRandomAsciiChar())
		}
		testSM.AttrOs7 = []byte(string(ps7))
	}

	return testSM
}

func TestOctetString(t *testing.T) {

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < 1000; i++ {

		testSM := createTestOctetString()

		t.Logf("Testing Test-OctetString with following values:\n"+
			"%v, \n%v, \n%v, \n%v, \n%v, \n%v, \n%v", hex.Dump(testSM.GetAttrOs1()), hex.Dump(testSM.GetAttrOs2()),
			hex.Dump(testSM.GetAttrOs3()), hex.Dump(testSM.GetAttrOs4()), hex.Dump(testSM.GetAttrOs5()),
			hex.Dump(testSM.GetAttrOs6()), hex.Dump(testSM.GetAttrOs7()))

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestOctetString(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-OctetString (with randomness) was successfully finished")
}
