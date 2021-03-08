// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NULL.h"
import "C"

func newNull() *C.NULL_t {
	res := C.int(0)
	return &res
}

func decodeNull() int32 {
	return int32(0)
}

//func freeAsnCodecsPrim(asnPrimTypeC *C.ASN__PRIMITIVE_TYPE_t) {
//	C.free(unsafe.Pointer(asnPrimTypeC.buf))
//}
