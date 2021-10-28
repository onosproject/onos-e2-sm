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

func createTestTopLevelPDU() (*test_sm_ies.TestTopLevelPdu, error) {

	var err error
	min := -2147483648
	max := 2147483648

	ttlpdu := new(test_sm_ies.TestTopLevelPdu)

	ttlpdu.Opt1 = &test_sm_ies.TestUnconstrainedInt{}
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ttlpdu.Opt1.AttrUciA = int32(rand.Intn(max-min) + min)
	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ttlpdu.Opt1.AttrUciB = int32(rand.Intn(max-min) + min)

	// ToDo - comment out once real numbers are introduced in encoding
	// Seeding randomizer first
	//rand.Seed(time.Now().UnixNano())
	//opt1 := rand.Intn(2)
	//if opt1 == 1 {
	//	//Some routine to generate constrained real numbers for TestConstrainedReal structure
	//}

	ttlpdu.Opt3 = createTestNestedChoicesMsg()
	ttlpdu.Opt4, err = createTestBitString()
	if err != nil {
		return nil, err
	}

	ttlpdu.Opt5 = createTestOctetString()

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ttlpdu.Opt6 = createTestListExtensible3Msg(rand.Intn(6))

	// Seeding randomizer first
	rand.Seed(time.Now().UnixNano())
	ttlpdu.Opt7 = test_sm_ies.TestEnumeratedExtensible(rand.Intn(6))


	return ttlpdu, nil
}

func TestTopLevelPdu(t *testing.T) {

	// Setting ChoiceMap to enable encoding with Go APER library (necessary prerequisite)
	aper.ChoiceMap = test_sm_ies.Choicemap

	for i := 0; i < 100000; i++ {

		testSM, err := createTestTopLevelPDU()
		assert.NilError(t, err)
		t.Logf("Testing Test-TopLevelPDU with value \n%v", testSM)

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestTopLevelPDU(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.MarshalWithParams(testSM, "valueExt")
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-TopLevelPDU was successfully finished")
}
