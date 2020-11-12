// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "OCTET_STRING.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// TODO: Change the argument to a []byte
func newOctetString(msg string) *C.OCTET_STRING_t {
	msgBytes := C.CBytes([]byte(msg))
	octStrC := C.OCTET_STRING_t{
		buf:  (*C.uchar)(msgBytes),
		size: C.ulong(len(msg)),
	}
	return &octStrC
}

func decodeOctetString(octC *C.OCTET_STRING_t) (string, error) {
	size := uint64(octC.size)
	//TODO: Doublecheck with specification
	if size > 8 {
		return "", fmt.Errorf("max size is 8 bytes (64 bits) got %d", size)
	}

	bytes := C.GoBytes(unsafe.Pointer(octC.buf), C.int(octC.size))
	return string(bytes), nil
}
