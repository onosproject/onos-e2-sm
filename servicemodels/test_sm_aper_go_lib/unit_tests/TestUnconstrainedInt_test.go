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

//func TestUnconstrainedIntAllOption1(t *testing.T) {
//
//	testSM := new(test_sm_ies.TestUnconstrainedInt)
//
//	for i := -2147483647; i < 2147483648; i++ {
//		testSM.AttrUciA = int32(i)
//		for j := -2147483647; j < 2147483648; j++ {
//			testSM.AttrUciB = int32(j)
//			t.Logf("Testing Test-UnconstrainedInt with values %v and %v", i, j)
//
//			perRef, err := testsmctypes.PerEncodeTestUnconstrainedInt(testSM)
//			assert.NilError(t, err)
//			t.Logf("CGo APER\n%v", hex.Dump(perRef))
//			per, err := aper.Marshal(testSM)
//			assert.NilError(t, err)
//
//			t.Logf("Go APER\n%v", hex.Dump(per))
//			assert.DeepEqual(t, per, perRef)
//		}
//	}
//	t.Logf("Testing of Test-UnconstrainedInt was successfully finished")
//}

func TestUnconstrainedIntAllOption2(t *testing.T) {

	testSM := new(test_sm_ies.TestUnconstrainedInt)

	min := -2147483648
	max := 2147483648

	for i := 1; i < 1000; i++ {
		// Seeding randomizer first
		rand.Seed(time.Now().UnixNano())
		// Generating random numbers
		testSM.AttrUciA = int32(rand.Intn(max-min) + min)
		testSM.AttrUciB = int32(rand.Intn(max-min) + min)
		t.Logf("Testing Test-UnconstrainedInt with values %v and %v", testSM.GetAttrUciA(), testSM.GetAttrUciB())

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestUnconstrainedInt(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-UnconstrainedInt (with randomness) was successfully finished")
}
