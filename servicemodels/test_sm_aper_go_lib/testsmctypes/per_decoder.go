// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TEST-TopLevelPDU.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func decodePer(bytes []byte, size int, valueType *C.asn_TYPE_descriptor_t) (unsafe.Pointer, error) {
	var result unsafe.Pointer

	cBytes := C.CBytes(bytes)
	defer C.free(cBytes)
	decRetVal, err := C.aper_decode_complete(nil, valueType, &result, cBytes, C.ulong(size))
	if err != nil {
		return nil, err
	}
	switch decRetVal.code {
	case C.RC_OK:
		return result, nil
	case C.RC_WMORE:
		return nil, fmt.Errorf("unhandled - want more. Consumed %v", decRetVal.consumed)
	case C.RC_FAIL:
		return nil, fmt.Errorf("failed to decode. Consumed %v", decRetVal.consumed)
	default:
		return nil, fmt.Errorf("unexpected return code %v", decRetVal.code)
	}
}
