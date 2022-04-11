// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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

func createTestFullyOptionalSequence() *test_sm_ies.TestFullyOptionalSequence {

	test := new(test_sm_ies.TestFullyOptionalSequence)

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	opt1 := rand.Intn(2)
	if opt1 == 1 {
		ie1 := rand.Int31()
		test.Item1 = &ie1
	}

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	opt2 := rand.Intn(2)
	if opt2 == 1 {
		// Seeding randomizer first
		rand.Seed(time.Now().UnixNano())
		var ps1 []rune
		// Limiting it on 1000 on purpose, otherwise test scenario would take too long
		for j := 1; j <= rand.Intn(7-3)+3; j++ {
			ps1 = append(ps1, createRandomAsciiChar())
		}
		test.Item2 = []byte(string(ps1))
	}

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	opt3 := rand.Intn(2)
	if opt3 == 1 {
		var b bool
		ie3 := rand.Intn(2)
		if ie3 == 0 {
			b = false
		}

		if ie3 == 1 {
			b = true
		}
		test.Item3 = &b
	}

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	opt4 := rand.Intn(2)
	if opt4 == 1 {
		// Seeding randomizer first
		rand.Seed(time.Now().UnixNano())
		enum := test_sm_ies.TestFullyOptionalSequenceItem4(rand.Intn(2))
		test.Item4 = &enum
	}

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	opt5 := rand.Intn(2)
	if opt5 == 1 {
		ie5 := int32(0)
		test.Item5 = &ie5
	}

	return test
}

func TestFullyOptionalSequence(t *testing.T) {

	for i := 0; i < 1000; i++ {

		testSM := createTestFullyOptionalSequence()
		t.Logf("Testing Test-FullyOptionalSequence with value \n%v", testSM)

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestFullyOptionalSequence(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.MarshalWithParams(testSM, "valueExt", test_sm_ies.Choicemap, nil)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-FullyOptionalSequence was successfully finished")
}
