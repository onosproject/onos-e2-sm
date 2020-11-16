// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes
//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "INTEGER.h"
import "C"
import (
	"unsafe"
)

// TODO: Change the argument to a []byte
func newInteger(msg string) *C.asn_INTEGER_enum_map_t {

	intC := C.asn_INTEGER_enum_map_t{
		nat_value: C.ulong(msg),
		enum_len: C.ulong(len(msg)),
		enum_name: C.CString(msg),
	}

	return &intC
}

func decodeInteger(intC *C.asn_INTEGER_enum_map_t) string {

	bytes := C.GoBytes(unsafe.Pointer(intC.enum_len), C.int(intC.nat_value))
	return string(bytes)
}