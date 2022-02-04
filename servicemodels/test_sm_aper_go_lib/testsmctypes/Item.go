// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "Item.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeItem(item *test_sm_ies.Item) ([]byte, error) {
	itemCP, err := newItem(item)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_Item, unsafe.Pointer(itemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeItem(item *test_sm_ies.Item) ([]byte, error) {
	itemCP, err := newItem(item)
	if err != nil {
		return nil, fmt.Errorf("perEncodeItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_Item, unsafe.Pointer(itemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeItem(bytes []byte) (*test_sm_ies.Item, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeItem((*C.Item_t)(unsafePtr))
}

func perDecodeItem(bytes []byte) (*test_sm_ies.Item, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeItem((*C.Item_t)(unsafePtr))
}

func newItem(item *test_sm_ies.Item) (*C.Item_t, error) {

	var err error
	itemC := C.Item_t{}

	if item.Item1 != nil {
		item1C := C.long(*item.Item1)
		itemC.item1 = &item1C
	}

	item2C, err := newBitString(item.Item2)
	if err != nil {
		return nil, fmt.Errorf("new.asn1.v1.BitString() %s", err.Error())
	}

	itemC.item2 = *item2C

	return &itemC, nil
}

func decodeItem(itemC *C.Item_t) (*test_sm_ies.Item, error) {

	var err error
	item := test_sm_ies.Item{}

	if itemC.item1 != nil {
		ie1 := int32(*itemC.item1)
		item.Item1 = &ie1
	}
	item.Item2, err = decodeBitString(&itemC.item2)
	if err != nil {
		return nil, fmt.Errorf("decode.asn1.v1.BitString() %s", err.Error())
	}

	return &item, nil
}
