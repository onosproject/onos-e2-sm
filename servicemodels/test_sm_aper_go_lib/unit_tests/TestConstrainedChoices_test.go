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

func createRandomIntegerInRange(min int32, max int32) int32 {
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())

	return int32(rand.Intn(int(max-min))) + min
}

func createRandomInteger64InRange(min int64, max int64) int64 {
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())

	return int64(rand.Intn(int(max-min))) + min
}

func createTestConstrainedChoicesMsg() *test_sm_ies.TestConstrainedChoices {

	var otherCAttr []rune
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < rand.Intn(50-1)+1; j++ {
		otherCAttr = append(otherCAttr, createRandomAsciiChar())
	}

	testChoices := test_sm_ies.TestConstrainedChoices{
		OtherCattr: string(otherCAttr),
		ConstrainedChoice1: &test_sm_ies.ConstrainedChoice1{
			ConstrainedChoice1: &test_sm_ies.ConstrainedChoice1_ConstrainedChoice1A{
				ConstrainedChoice1A: createRandomIntegerInRange(1, 136), // providing wider range in order to have an extensed value
			},
		},
		ConstrainedChoice4: &test_sm_ies.ConstrainedChoice4{
			ConstrainedChoice4: &test_sm_ies.ConstrainedChoice4_ConstrainedChoice4A{
				ConstrainedChoice4A: createRandomIntegerInRange(1, 256), // providing wider range in order to have an extensed value
			},
		},
	}

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ch2 := rand.Intn(1)
	switch ch2 {
	case 0:
		testChoices.ConstrainedChoice2 = &test_sm_ies.ConstrainedChoice2{
			ConstrainedChoice2: &test_sm_ies.ConstrainedChoice2_ConstrainedChoice2A{
				ConstrainedChoice2A: createRandomIntegerInRange(0, 18), // providing wider range in order to have an extensed value
			},
		}
	case 1:
		testChoices.ConstrainedChoice2 = &test_sm_ies.ConstrainedChoice2{
			ConstrainedChoice2: &test_sm_ies.ConstrainedChoice2_ConstrainedChoice2B{
				ConstrainedChoice2B: createRandomInteger64InRange(1, 4294967350),
			},
		}
	default:
		testChoices.ConstrainedChoice2 = &test_sm_ies.ConstrainedChoice2{
			ConstrainedChoice2: &test_sm_ies.ConstrainedChoice2_ConstrainedChoice2A{
				ConstrainedChoice2A: createRandomIntegerInRange(0, 15),
			},
		}
	}

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ch3 := rand.Intn(3)
	switch ch3 {
	case 0:
		testChoices.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3{
			ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3A{
				ConstrainedChoice3A: createRandomInteger(),
			},
		}
	case 1:
		testChoices.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3{
			ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3B{
				ConstrainedChoice3B: createRandomIntegerInRange(0, 15),
			},
		}
	case 2:
		testChoices.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3{
			ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3C{
				ConstrainedChoice3C: createRandomIntegerInRange(1, 2147483647),
			},
		}
	case 3:
		testChoices.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3{
			ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3D{
				ConstrainedChoice3D: createRandomIntegerInRange(1, 2147483647),
			},
		}
	default:
		testChoices.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3{
			ConstrainedChoice3: &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3A{
				ConstrainedChoice3A: createRandomInteger(),
			},
		}
	}

	return &testChoices
}

func TestConstrainedChoices(t *testing.T) {

	// Setting ChoiceMap to enable encoding with Go APER library (necessary prerequisite)
	aper.ChoiceMap = test_sm_ies.Choicemap

	for i := 1; i < 1000; i++ {

		testSM := createTestConstrainedChoicesMsg()
		t.Logf("Testing Test-ConstrainedChoices with value \n%v", testSM)

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestConstrainedChoices(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-ConstrainedChoices was successfully finished")
}
