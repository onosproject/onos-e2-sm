// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package unit_tests

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/testsmctypes"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"math/rand"
	"testing"
	"time"
)

func createRandomAsciiChar() rune {
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())

	return rune(rand.Intn(127))
}

func TestPrintableString(t *testing.T) {

	testSM := new(test_sm_ies.TestPrintableString)

	for i := 1; i < 100000; i++ {

		var ps1 []rune
		// Limiting it on 1000 on purpose, otherwise test scenario would take too long
		for j := 1; j < rand.Intn(1000); j++ {
			ps1 = append(ps1, createRandomAsciiChar())
		}
		testSM.AttrPs1 = string(ps1)

		var ps2 []rune
		for j := 1; j <= 2; j++ {
			ps2 = append(ps2, createRandomAsciiChar())
		}
		testSM.AttrPs2 = string(ps2)

		var ps3 []rune
		//need to indicate whether next value is extensible or not
		ext := rand.Intn(1)
		if ext == 1 {
			for j := 1; j <= 2 + rand.Int(); j++ {
				ps3 = append(ps3, createRandomAsciiChar())
			}
			testSM.AttrPs3 = string(ps3)
		} else {
			for j := 1; j <= 2; j++ {
				ps3 = append(ps3, createRandomAsciiChar())
			}
			testSM.AttrPs3 = string(ps3)
		}

		var ps4 []rune
		//for j := 1; j <= rand.Intn(3); j++ {
		//For some reason, testsmctypes encoder doesn't like anything less than 3 characters
		for j := 1; j <= 3; j++ {
			ps4 = append(ps4, createRandomAsciiChar())
		}
		testSM.AttrPs4 = string(ps4)

		var ps5 []rune
		//for j := 2; j <= 2 + rand.Intn(1); j++ {
		//For some reason, testsmctypes encoder doesn't like anything less than 3 characters
		for j := 1; j <= 3; j++ {
			ps5 = append(ps5, createRandomAsciiChar())
		}
		testSM.AttrPs5 = string(ps5)

		var ps6 []rune
		//need to indicate whether next value is extensible or not
		ext = rand.Intn(1)
		if ext == 1 {
			// Limiting it on 1000 on purpose, otherwise test scenario would take too long
			for j := 1; j <= 3 + rand.Intn(1000); j++ {
				ps6 = append(ps6, createRandomAsciiChar())
			}
			testSM.AttrPs6 = string(ps6)
		} else {
			//For some reason, testsmctypes encoder doesn't like anything less than 3 characters
			for j := 1; j <= 3; j++ {
				ps6 = append(ps6, createRandomAsciiChar())
			}
			testSM.AttrPs6 = string(ps6)
		}

		var ps7 []rune
		optional := rand.Intn(1)
		if optional == 1 {
			for j := 2; j <= 2 + rand.Intn(4); j++ {
				ps7 = append(ps7, createRandomAsciiChar())
			}
			str := string(ps7)
			testSM.AttrPs7 = &str
		}

		t.Logf("Testing Test-PrintableString with following values:\n"+
			"%v, \n%v, \n%v, \n%v, \n%v, \n%v, \n%v", testSM.GetAttrPs1(), testSM.GetAttrPs2(), testSM.GetAttrPs3(),
			testSM.GetAttrPs4(), testSM.GetAttrPs5(), testSM.GetAttrPs6(), testSM.GetAttrPs7())

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestPrintableString(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-PrintableString (with randomness) was successfully finished")
}
