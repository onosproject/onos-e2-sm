// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func createItemMsg() (*test_sm_ies.Item, error) {

	// item := pdubuilder.CreateItem() //ToDo - fill in arguments here(if this function exists

	var ie1 int32 = 32
	item := test_sm_ies.Item{
		Item1: &ie1,
		Item2: &asn1.BitString{
			Value: []byte{0x00, 0x00, 0x40},
			Len: 20,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating Item %s", err.Error())
	//}
	return &item, nil
}

func Test_xerEncodingItem(t *testing.T) {

	item, err := createItemMsg()
	assert.NilError(t, err, "Error creating Item PDU")

	xer, err := xerEncodeItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("Item XER\n%s", string(xer))

	result, err := xerDecodeItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Item XER - decoded\n%v", result)
	assert.Equal(t, item.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, item.GetItem2().GetValue(), result.GetItem2().GetValue())
	assert.Equal(t, item.GetItem2().GetLen(), result.GetItem2().GetLen())
}

func Test_perEncodingItem(t *testing.T) {

	item, err := createItemMsg()
	assert.NilError(t, err, "Error creating Item PDU")

	per, err := perEncodeItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("Item PER\n%v", hex.Dump(per))

	result, err := perDecodeItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Item PER - decoded\n%v", result)
	assert.Equal(t, item.GetItem1(), result.GetItem1())
	assert.DeepEqual(t, item.GetItem2().GetValue(), result.GetItem2().GetValue())
	assert.Equal(t, item.GetItem2().GetLen(), result.GetItem2().GetLen())
}
