// SPDX-FileCopyrightText: (C) 2023 Intel Corporation
// SPDX-License-Identifier: LicenseRef-Intel

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-Extension2.h"
import "C"

import (
	"fmt"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"unsafe"
)

func xerEncodeTestExtension2(testExt1 *test_sm_ies.TestExtension2) ([]byte, error) {
	testExtension2CP, err := newTestExtension2(testExt1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestExtension2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TEST_Extension2, unsafe.Pointer(testExtension2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTestExtension2() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeTestExtension2(testExt1 *test_sm_ies.TestExtension2) ([]byte, error) {
	testExt1CP, err := newTestExtension2(testExt1)
	if err != nil {
		return nil, fmt.Errorf("PerEncodeTestExtension2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TEST_Extension2, unsafe.Pointer(testExt1CP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeTestExtension2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTestExtension2(bytes []byte) (*test_sm_ies.TestExtension2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TEST_Extension2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTestExtension2((*C.TEST_Extension2_t)(unsafePtr))
}

func PerDecodeTestExtension2(bytes []byte) (*test_sm_ies.TestExtension2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TEST_Extension2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTestExtension2((*C.TEST_Extension2_t)(unsafePtr))
}

func newTestExtension2(testExtension2 *test_sm_ies.TestExtension2) (*C.TEST_Extension2_t, error) {

	testExtension2C := C.TEST_Extension2_t{
		item1: C.long(testExtension2.GetItem1()),
	}

	if testExtension2.GetItem2() != nil {
		item2C, err := newOctetString(testExtension2.GetItem2())
		if err != nil {
			return nil, err
		}
		testExtension2C.item2 = item2C
	}

	if testExtension2.Ext1 != nil {
		ext1C, err := newOctetString(testExtension2.GetExt1())
		if err != nil {
			return nil, err
		}
		testExtension2C.ext1 = ext1C
	}

	if testExtension2.Ext2 != nil {
		ext2C, err := newOctetString(testExtension2.GetExt2())
		if err != nil {
			return nil, err
		}
		testExtension2C.ext2 = ext2C
	}

	if testExtension2.Ext3 != nil {
		ext3C, err := newOctetString(testExtension2.GetExt3())
		if err != nil {
			return nil, err
		}
		testExtension2C.ext3 = ext3C
	}

	if testExtension2.Ext4 != nil {
		ext4C, err := newOctetString(testExtension2.GetExt4())
		if err != nil {
			return nil, err
		}
		testExtension2C.ext4 = ext4C
	}

	if testExtension2.Ext5 != nil {
		ext5C, err := newOctetString(testExtension2.GetExt5())
		if err != nil {
			return nil, err
		}
		testExtension2C.ext5 = ext5C
	}

	if testExtension2.Ext6 != nil {
		ext6C, err := newOctetString(testExtension2.GetExt6())
		if err != nil {
			return nil, err
		}
		testExtension2C.ext6 = ext6C
	}

	if testExtension2.Ext7 != nil {
		ext7C, err := newOctetString(testExtension2.GetExt7())
		if err != nil {
			return nil, err
		}
		testExtension2C.ext7 = ext7C
	}

	if testExtension2.Ext8 != nil {
		ext8C, err := newOctetString(testExtension2.GetExt8())
		if err != nil {
			return nil, err
		}
		testExtension2C.ext8 = ext8C
	}

	if testExtension2.Ext9 != nil {
		ext9C, err := newOctetString(testExtension2.GetExt9())
		if err != nil {
			return nil, err
		}

		testExtension2C.ext9 = ext9C
	}

	return &testExtension2C, nil
}

func decodeTestExtension2(testExtension2C *C.TEST_Extension2_t) (*test_sm_ies.TestExtension2, error) {

	var err error
	testExtension2 := test_sm_ies.TestExtension2{}

	testExtension2.Item1 = int32(testExtension2C.item1)

	if testExtension2C.item2 != nil {
		testExtension2.Item2, err = decodeOctetString(testExtension2C.item2)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext1 != nil {
		testExtension2.Ext1, err = decodeOctetString(testExtension2C.ext1)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext2 != nil {
		testExtension2.Ext2, err = decodeOctetString(testExtension2C.ext2)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext3 != nil {
		testExtension2.Ext3, err = decodeOctetString(testExtension2C.ext3)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext4 != nil {
		testExtension2.Ext4, err = decodeOctetString(testExtension2C.ext4)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext5 != nil {
		testExtension2.Ext5, err = decodeOctetString(testExtension2C.ext5)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext6 != nil {
		testExtension2.Ext6, err = decodeOctetString(testExtension2C.ext6)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext7 != nil {
		testExtension2.Ext7, err = decodeOctetString(testExtension2C.ext7)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext8 != nil {
		testExtension2.Ext8, err = decodeOctetString(testExtension2C.ext8)
		if err != nil {
			return nil, err
		}
	}

	if testExtension2C.ext9 != nil {
		testExtension2.Ext9, err = decodeOctetString(testExtension2C.ext9)
		if err != nil {
			return nil, err
		}
	}

	return &testExtension2, nil
}
