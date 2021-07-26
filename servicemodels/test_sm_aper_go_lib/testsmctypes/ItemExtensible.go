// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ItemExtensible.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeItemExtensible(itemExtensible *test_sm_ies.ItemExtensible) ([]byte, error) {
	itemExtensibleCP, err := newItemExtensible(itemExtensible)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeItemExtensible() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ItemExtensible, unsafe.Pointer(itemExtensibleCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeItemExtensible() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeItemExtensible(itemExtensible *test_sm_ies.ItemExtensible) ([]byte, error) {
	itemExtensibleCP, err := newItemExtensible(itemExtensible)
	if err != nil {
		return nil, fmt.Errorf("perEncodeItemExtensible() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ItemExtensible, unsafe.Pointer(itemExtensibleCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeItemExtensible() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeItemExtensible(bytes []byte) (*test_sm_ies.ItemExtensible, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ItemExtensible)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeItemExtensible((*C.ItemExtensible_t)(unsafePtr))
}

func perDecodeItemExtensible(bytes []byte) (*test_sm_ies.ItemExtensible, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ItemExtensible)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeItemExtensible((*C.ItemExtensible_t)(unsafePtr))
}

func newItemExtensible(itemExtensible *test_sm_ies.ItemExtensible) (*C.ItemExtensible_t, error) {

	itemExtensibleC := C.ItemExtensible_t{}

	item1C := C.long(itemExtensible.Item1)
	itemExtensibleC.item1 = item1C

	if itemExtensible.Item2 != nil {
		item2C, err := newOctetString(itemExtensible.Item2)
		if err != nil {
			return nil, err
		}
		itemExtensibleC.item2 = item2C
	}

	return &itemExtensibleC, nil
}

func decodeItemExtensible(itemExtensibleC *C.ItemExtensible_t) (*test_sm_ies.ItemExtensible, error) {

	var err error
	itemExtensible := test_sm_ies.ItemExtensible{}

	itemExtensible.Item1 = int32(itemExtensibleC.item1)

	if itemExtensibleC.item2 != nil {
		itemExtensible.Item2, err = decodeOctetString(itemExtensibleC.item2)
		if err != nil {
			return nil, err
		}
	}

	return &itemExtensible, nil
}
