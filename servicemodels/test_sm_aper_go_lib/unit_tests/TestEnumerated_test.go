// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package unit_tests

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/testsmctypes"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"testing"
)

func createTestEnum1() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM1
}

func createTestEnum2() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM2
}

func createTestEnum3() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM3
}

func createTestEnum4() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM4
}

func createTestEnum5() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM5
}

func createTestEnum6() test_sm_ies.TestEnumerated {
	return test_sm_ies.TestEnumerated_TEST_ENUMERATED_ENUM6
}


func TestEnumerated(t *testing.T) {

	for i := 0; i < 6; i++ {

		testSM := test_sm_ies.TestEnumerated(i)
		t.Logf("Testing Test-Enumerated with value %v\n", testSM.Number())

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestEnumerated(&testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.MarshalWithParams(&testSM, "valueLB:0,valueUB:5")
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-Enumerated was successfully finished")
}
