// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "asn_codecs_prim.h"
import "C"
import (
	"unsafe"
)

// TODO: Change the argument to a []byte
func newAsnCodecsPrim(msg string) *C.ASN__PRIMITIVE_TYPE_t {
	msgBytes := C.CBytes([]byte(msg))
	asnPrimTypeC := C.ASN__PRIMITIVE_TYPE_t{
		buf:  (*C.uchar)(msgBytes),
		size: C.ulong(len(msg)),
	}
	return &asnPrimTypeC
}

func decodeAsnCodecsPrim(asnPrimTypeC *C.ASN__PRIMITIVE_TYPE_t) string {

	bytes := C.GoBytes(unsafe.Pointer(asnPrimTypeC.buf), C.int(asnPrimTypeC.size))
	return string(bytes)
}
