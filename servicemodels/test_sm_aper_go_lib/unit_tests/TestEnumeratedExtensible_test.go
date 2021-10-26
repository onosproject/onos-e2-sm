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

func createTestEnum1Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM1
}

func createTestEnum2Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM2
}

func createTestEnum3Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM3
}

func createTestEnum4Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM4
}

func createTestEnum5Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM5
}

func createTestEnum6Ext() test_sm_ies.TestEnumeratedExtensible {
	return test_sm_ies.TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM6
}

func TestEnumeratedExtensible(t *testing.T) {

	for i := 0; i < 6; i++ {

		testSM := test_sm_ies.TestEnumeratedExtensible(i)
		t.Logf("Testing Test-EnumeratedExtensible with value %v\n", testSM.Number())

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestEnumeratedExtensible(&testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.MarshalWithParams(&testSM, "valueLB:0,valueUB:5,valueExt")
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-EnumeratedExtensible was successfully finished")
}
