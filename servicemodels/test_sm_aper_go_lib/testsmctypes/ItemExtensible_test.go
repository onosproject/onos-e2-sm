// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createItemExtensibleMsg() (*test_sm_ies.ItemExtensible, error) {

	// itemExtensible := pdubuilder.CreateItemExtensible() //ToDo - fill in arguments here(if this function exists

	itemExtensible := test_sm_ies.ItemExtensible{
		Item1: 32,
		Item2: []byte{0xab, 0xcd, 0xef},
	}

	//if err := itemExtensible.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating ItemExtensible %s", err.Error())
	//}
	return &itemExtensible, nil
}

func Test_xerEncodingItemExtensible(t *testing.T) {

	itemExtensible, err := createItemExtensibleMsg()
	assert.NilError(t, err, "Error creating ItemExtensible PDU")

	xer, err := xerEncodeItemExtensible(itemExtensible)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("ItemExtensible XER\n%s", string(xer))

	result, err := xerDecodeItemExtensible(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ItemExtensible XER - decoded\n%v", result)
	assert.Equal(t, itemExtensible.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, itemExtensible.GetItem2(), result.GetItem2())
}

func Test_perEncodingItemExtensible(t *testing.T) {

	itemExtensible, err := createItemExtensibleMsg()
	assert.NilError(t, err, "Error creating ItemExtensible PDU")

	per, err := perEncodeItemExtensible(itemExtensible)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("ItemExtensible PER\n%v", hex.Dump(per))

	result, err := perDecodeItemExtensible(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ItemExtensible PER - decoded\n%v", result)
	assert.Equal(t, itemExtensible.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, itemExtensible.GetItem2(), result.GetItem2())
}
