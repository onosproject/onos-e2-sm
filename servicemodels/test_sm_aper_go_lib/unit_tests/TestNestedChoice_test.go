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

func createChoice3() *test_sm_ies.Choice3 {

	choice3 := new(test_sm_ies.Choice3)
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ch3 := rand.Intn(2)
	switch ch3 {
	case 0:
		choice3.Choice3 = &test_sm_ies.Choice3_Choice3A{
			Choice3A: createRandomInteger(),
		}
	case 1:
		choice3.Choice3 = &test_sm_ies.Choice3_Choice3B{
			Choice3B: createRandomInteger(),
		}
	case 2:
		choice3.Choice3 = &test_sm_ies.Choice3_Choice3C{
			Choice3C: createRandomInteger(),
		}
	default:
		choice3.Choice3 = &test_sm_ies.Choice3_Choice3B{
			Choice3B: createRandomInteger(),
		}
	}

	return choice3
}

func createConstrainedChoice3() *test_sm_ies.ConstrainedChoice3 {

	constrainedChoice3 := new(test_sm_ies.ConstrainedChoice3)
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	conCh3 := rand.Intn(3)
	switch conCh3 {
	case 0:
		constrainedChoice3.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3A{
			ConstrainedChoice3A: createRandomInteger(),
		}
	case 1:
		constrainedChoice3.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3B{
			ConstrainedChoice3B: createRandomIntegerInRange(0, 18), // providing wider range in order to have an extensed value
		}
	case 2:
		constrainedChoice3.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3C{
			ConstrainedChoice3C: createRandomIntegerInRange(1, 2147483647),
		}
	case 3:
		constrainedChoice3.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3D{
			ConstrainedChoice3D: createRandomIntegerInRange(1, 4),
		}
	default:
		constrainedChoice3.ConstrainedChoice3 = &test_sm_ies.ConstrainedChoice3_ConstrainedChoice3C{
			ConstrainedChoice3C: createRandomIntegerInRange(1, 2147483647),
		}
	}

	return constrainedChoice3
}

func createConstrainedChoice4() *test_sm_ies.ConstrainedChoice4 {

	constrainedChoice4 := test_sm_ies.ConstrainedChoice4{
		ConstrainedChoice4: &test_sm_ies.ConstrainedChoice4_ConstrainedChoice4A{
			ConstrainedChoice4A: createRandomIntegerInRange(1, 140), // providing wider range in order to have an extensed value
		},
	}

	return &constrainedChoice4
}

func createTestNestedChoicesMsg() *test_sm_ies.TestNestedChoice {

	testNestedChoice := new(test_sm_ies.TestNestedChoice)

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	option := rand.Intn(2)
	switch option {
	case 0:
		testNestedChoice.TestNestedChoice = &test_sm_ies.TestNestedChoice_Option1{
			Option1: createChoice3(),
		}
	case 1:
		testNestedChoice.TestNestedChoice = &test_sm_ies.TestNestedChoice_Option2{
			Option2: createConstrainedChoice3(),
		}
	case 2:
		testNestedChoice.TestNestedChoice = &test_sm_ies.TestNestedChoice_Option3{
			Option3: createConstrainedChoice4(),
		}
	default:
		testNestedChoice.TestNestedChoice = &test_sm_ies.TestNestedChoice_Option2{
			Option2: createConstrainedChoice3(),
		}
	}

	return testNestedChoice
}

func TestNestedChoices(t *testing.T) {

	// Setting ChoiceMap to enable encoding with Go APER library (necessary prerequisite)
	aper.ChoiceMap = test_sm_ies.Choicemap

	for i := 1; i < 1000; i++ {

		testSM := createTestNestedChoicesMsg()
		t.Logf("Testing Test-NestedChoice with value \n%v", testSM)

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestNestedChoice(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.MarshalWithParams(testSM, "choiceExt")
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-NestedChoice was successfully finished")
}
