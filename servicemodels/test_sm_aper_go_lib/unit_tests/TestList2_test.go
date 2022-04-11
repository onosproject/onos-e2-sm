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

func generateItemExtensible() *test_sm_ies.ItemExtensible {

	item := new(test_sm_ies.ItemExtensible)

	ie1 := rand.Int31()
	item.Item1 = ie1

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	var ps1 []rune
	// Limiting it on 1000 on purpose, otherwise test scenario would take too long
	for j := 1; j <= rand.Intn(7-3)+3; j++ {
		ps1 = append(ps1, createRandomAsciiChar())
	}
	item.Item2 = []byte(string(ps1))

	return item
}

func createTestList2Msg(numItems int) (*test_sm_ies.TestList2, error) {

	list := &test_sm_ies.TestList2{
		Value: make([]*test_sm_ies.ItemExtensible, 0),
	}

	if numItems > 0 {
		for i := 1; i <= numItems; i++ {
			ie := generateItemExtensible()
			list.Value = append(list.Value, ie)
		}
	}

	return list, nil
}

func TestList2(t *testing.T) {

	for i := 0; i <= 12; i++ {

		testSM, err := createTestList2Msg(i)
		assert.NilError(t, err)
		t.Logf("Testing Test-List2 with value \n%v", testSM)

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestList2(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM, test_sm_ies.Choicemap, nil)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-List2 was successfully finished")
}
