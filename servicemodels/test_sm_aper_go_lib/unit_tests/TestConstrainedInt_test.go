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

func TestConstrainedInt(t *testing.T) {

	testSM := new(test_sm_ies.TestConstrainedInt)

	min := -2147483648
	max := 2147483648

	for i := 1; i < 100000; i++ {
		// Seeding randomizer first
		rand.Seed(time.Now().UnixNano())
		// Generating random numbers
		testSM.AttrCiA = int32(rand.Intn(100-10)) + 10
		testSM.AttrCiB = int32(rand.Intn(65535-255)) + 255
		testSM.AttrCiC = int32(rand.Intn(max-10)) + 10
		testSM.AttrCiD = int32(rand.Intn(100-min) + min)
		testSM.AttrCiE = int32(rand.Intn(20-10)) + 10
		testSM.AttrCiF = 10
		testSM.AttrCiG = 10 + int32(rand.Intn(max-10))

		t.Logf("Testing Test-ConstrainedInt with values:\n" +
			"%v, %v, %v, %v, %v, %v, %v", testSM.GetAttrCiA(), testSM.GetAttrCiB(), testSM.GetAttrCiC(),
			testSM.GetAttrCiD(), testSM.GetAttrCiE(), testSM.GetAttrCiF(), testSM.GetAttrCiG())

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestConstrainedInt(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-ConstrainedInt (with randomness) was successfully finished")
}
