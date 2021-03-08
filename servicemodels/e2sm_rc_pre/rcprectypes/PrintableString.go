// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PrintableString.h"
import "C"

//import "unsafe"

// TODO: Change the argument to a []byte
func newPrintableString(msg string) *C.PrintableString_t {
	// PrintableString is defined via OctetString --> see PrintableString.h
	prntStrC := newOctetString(msg)

	return prntStrC
}

func decodePrintableString(octC *C.PrintableString_t) string {

	bytes := decodeOctetString(octC)
	return bytes
}

//func decodePrintableStringBytes(array [16]byte) (string, error) {
//	intC := (*C.PrintableString_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))
//	return decodePrintableString(intC), nil
//}
