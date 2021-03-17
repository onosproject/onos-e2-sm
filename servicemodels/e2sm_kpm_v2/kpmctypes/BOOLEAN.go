// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "BOOLEAN.h"
import "C"

//func newBoolean(b bool) *C.BOOLEAN_t {
//
//	//ToDo - read a comment in BOOLEAN.h about possible values
//	bC := C.int(0)
//	if b {
//		bC = C.int(1)
//	}
//
//	return &bC
//}
//
//func decodeBoolean(bC *C.BOOLEAN_t) bool {
//
//	//ToDo - fix decoding
//	if int32(bC) == int32(1) {
//		return true
//	}
//	return false
//}
//
//func decodeBooleanBytes(array [8]byte) bool {
//	bC := (*C.BOOLEAN_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))
//
//	return decodeBoolean(bC)
//}