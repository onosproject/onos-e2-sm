// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "OCTET_STRING.h"
import "C"
import (
	"encoding/binary"
	"unsafe"
)

// TODO: Change the argument to a []byte
func newOctetString(msg string) (*C.OCTET_STRING_t, error) {
	//fmt.Printf("We're inside newOctetString(), starting encoding...")
	//fmt.Printf("Encoding OctetString: \n %v \n", msg)
	msgBytes := C.CBytes([]byte(msg))
	octStrC := C.OCTET_STRING_t{
		buf:  (*C.uchar)(msgBytes),
		size: C.ulong(len(msg)),
	}
	//C.free(unsafe.Pointer(msgBytes))
	//fmt.Printf("That's the C version of encoded OctetString: \n %v \n", octStrC)
	return &octStrC, nil
}

func decodeOctetString(octC *C.OCTET_STRING_t) (string, error) {

	//fmt.Printf("We're inside decodeOctetString(), starting decoding...")
	//fmt.Printf("Decoding OctetString: \n %v \n", octC)
	if octC == nil {
		return "", nil //ToDo - should it return error that OctetString is empty one?
	}
	bytes := C.GoBytes(unsafe.Pointer(octC.buf), C.int(octC.size))
	//freeOctetString(octC)
	//fmt.Printf("That's what was decoded from C: \n %v \n", bytes)
	return string(bytes), nil
}

//func freeOctetString(octC *C.OCTET_STRING_t) {
//	C.free(unsafe.Pointer(octC))
//}

func newOctetStringFromArray(array [16]byte) *C.OCTET_STRING_t {
	//fmt.Printf("We're inside newOctetStringFromArray(), starting encoding...")
	//fmt.Printf("Creating OctetString from bytes: \n %v \n", array)
	size := binary.LittleEndian.Uint64(array[8:16])
	bytes := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[:8]))), C.int(size))

	//fmt.Printf("Creating OctetString from bytes: \n %v of size %v \n", bytes, size)
	octStrC := C.OCTET_STRING_t{
		buf:  (*C.uchar)(C.CBytes(bytes)),
		size: C.ulong(size),
	}
	//C.free(unsafe.Pointer(bytes))
	//fmt.Printf("That's the C version of OctetString encoded from bytes: \n %v \n", octStrC)

	return &octStrC
}

func decodeOctetStringBytes(array [16]byte) (string, error) {
	octSC := newOctetStringFromArray(array)

	return decodeOctetString(octSC)
}
