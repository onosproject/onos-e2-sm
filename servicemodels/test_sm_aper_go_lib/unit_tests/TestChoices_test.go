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

func createRandomInteger() int32 {
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())

	return rand.Int31()
}

func createTestChoicesMsg() *test_sm_ies.TestChoices {

	var otherAttr []rune
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	// Limiting it on 1000 on purpose, otherwise test scenario would take too long
	for j := 1; j < rand.Intn(1000); j++ {
		otherAttr = append(otherAttr, createRandomAsciiChar())
	}

	testChoices := test_sm_ies.TestChoices{
		OtherAttr: string(otherAttr),
		Choice1: &test_sm_ies.Choice1{
			Choice1: &test_sm_ies.Choice1_Choice1A{
				Choice1A: createRandomInteger(),
			},
		},
		Choice4: &test_sm_ies.Choice4{
			Choice4: &test_sm_ies.Choice4_Choice4A{
				Choice4A: createRandomInteger(),
			},
		},
	}

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ch2 := rand.Intn(1)
	switch ch2 {
	case 0:
		testChoices.Choice2 = &test_sm_ies.Choice2{
				Choice2: &test_sm_ies.Choice2_Choice2A{
					Choice2A: createRandomInteger(),
				},
			}
	case 1:
		testChoices.Choice2 = &test_sm_ies.Choice2{
			Choice2: &test_sm_ies.Choice2_Choice2B{
				Choice2B: createRandomInteger(),
			},
		}
	default:
		testChoices.Choice2 = &test_sm_ies.Choice2{
			Choice2: &test_sm_ies.Choice2_Choice2A{
				Choice2A: createRandomInteger(),
			},
		}
	}

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ch3 := rand.Intn(2)
	switch ch3 {
	case 0:
		testChoices.Choice3 = &test_sm_ies.Choice3{
			Choice3: &test_sm_ies.Choice3_Choice3A{
				Choice3A: createRandomInteger(),
			},
		}
	case 1:
		testChoices.Choice3 = &test_sm_ies.Choice3{
			Choice3: &test_sm_ies.Choice3_Choice3B{
				Choice3B: createRandomInteger(),
			},
		}
	case 2:
		testChoices.Choice3 = &test_sm_ies.Choice3{
			Choice3: &test_sm_ies.Choice3_Choice3C{
				Choice3C: createRandomInteger(),
			},
		}
	default:
		testChoices.Choice3 = &test_sm_ies.Choice3{
			Choice3: &test_sm_ies.Choice3_Choice3C{
				Choice3C: createRandomInteger(),
			},
		}
	}

	return &testChoices
}

func TestChoices(t *testing.T) {

	// Setting ChoiceMap to enable encoding with Go APER library (necessary prerequisite)
	aper.ChoiceMap = test_sm_ies.Choicemap

	for i := 1; i < 100000; i++ {

		testSM := createTestChoicesMsg()
		t.Logf("Testing Test-Choices with value \n%v", testSM)

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestChoices(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-Choices was successfully finished")
}
