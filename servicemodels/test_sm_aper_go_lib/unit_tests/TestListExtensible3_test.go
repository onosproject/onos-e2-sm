// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package unit_tests

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/testsmctypes"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"testing"
)

func createTestListExtensible3Msg(numItems int) *test_sm_ies.TestListExtensible3 {

	list := &test_sm_ies.TestListExtensible3{
		Value: make([]*test_sm_ies.TestFullyOptionalSequence, 0),
	}

	if numItems > 0 {
		for i := 1; i <= numItems; i++ {
			ie := createTestFullyOptionalSequence()
			list.Value = append(list.Value, ie)
		}
	}

	return list
}

func TestListExtensible3(t *testing.T) {

	for i := 0; i <= 12; i++ {

		testSM := createTestListExtensible3Msg(i)
		t.Logf("Testing Test-ListExtensible3 with value \n%v", testSM)

		// Generating APER with reference CGo approach
		perRef, err := testsmctypes.PerEncodeTestListExtensible3(testSM)
		assert.NilError(t, err)
		// Generating APER bytes with Go APER lib
		per, err := aper.Marshal(testSM, test_sm_ies.Choicemap, nil)
		assert.NilError(t, err)

		//Comparing bytes against each other
		assert.DeepEqual(t, per, perRef)
	}
	t.Logf("Testing of Test-ListExtensible3 was successfully finished")
}
