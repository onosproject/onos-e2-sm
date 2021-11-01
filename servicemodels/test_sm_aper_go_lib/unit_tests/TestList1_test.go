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

func generateItem() (*test_sm_ies.Item, error) {

	item := new(test_sm_ies.Item)

	ie1 := rand.Int31()
	item.Item1 = &ie1

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ie2, err := generateBitStringOfSize(uint32(rand.Intn(7-3) + 3))
	if err != nil {
		return nil, err
	}
	item.Item2 = ie2

	return item, nil
}

func createTestList1Msg(numItems int) (*test_sm_ies.TestList1, error) {

	list := &test_sm_ies.TestList1{
		Value: make([]*test_sm_ies.Item, 0),
	}

	if numItems > 0 {
		for i := 1; i <= numItems; i++ {
			ie, err := generateItem()
			if err != nil {
				return nil, err
			}
			list.Value = append(list.Value, ie)
		}
	}

	return list, nil
}

func TestList1(t *testing.T) {

	// Setting ChoiceMap to enable encoding with Go APER library (necessary prerequisite)
	aper.ChoiceMap = test_sm_ies.Choicemap

	for i := 0; i <= 12; i++ {

		testSM, err := createTestList1Msg(i)
		assert.NilError(t, err)
		t.Logf("Testing Test-List1 with value \n%v", testSM)

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestList1(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-List1 was successfully finished")
}
