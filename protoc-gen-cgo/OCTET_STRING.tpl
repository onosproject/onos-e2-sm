// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package {{.PackageName}}ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "OCTET_STRING.h"
import "C"
import (
	"unsafe"
)

// TODO: Change the argument to a []byte
func newOctetString(msg string) (*C.OCTET_STRING_t, error) {
	msgBytes := C.CBytes([]byte(msg))
	octStrC := C.OCTET_STRING_t{
		buf:  (*C.uchar)(msgBytes),
		size: C.ulong(len(msg)),
	}
	return &octStrC, nil
}

func decodeOctetString(octC *C.OCTET_STRING_t) (string, error) {

	if octC == nil {
		return "", nil //ToDo - should it return error that OctetString is empty one?
	}
	bytes := C.GoBytes(unsafe.Pointer(octC.buf), C.int(octC.size))
	return string(bytes), nil
}

//func freeOctetString(octC *C.OCTET_STRING_t) {
//	C.free(unsafe.Pointer(octC))
//}
