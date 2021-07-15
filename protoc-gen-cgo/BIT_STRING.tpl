// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package {{.PackageName}}ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "BIT_STRING.h"
import "C"
import (
"encoding/binary"
"fmt"
{{.ProtoFileName}} "github.com/onosproject/onos-e2-sm/servicemodels/{{.FullPackageName}}" //ToDo - Make imports more dynamic
"math"
"unsafe"
)

func xerEncodeBitString(bs *{{.ProtoFileName}}.BitString) ([]byte, error) {
	bsC, err := newBitString(bs)
	if err != nil {
		return nil, fmt.Errorf("newBitString() %s", err.Error())
	}
	//defer freeBitString(bsC)
	bytes, err := encodeXer(&C.asn_DEF_BIT_STRING, unsafe.Pointer(bsC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func xerDecodeBitString(bytes []byte) (*{{.ProtoFileName}}, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_BIT_STRING)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeBitString((*C.BIT_STRING_t)(unsafePtr))
}

func perEncodeBitString(bs *{{.ProtoFileName}}.BitString) ([]byte, error) {
	bsC, err := newBitString(bs)
	if err != nil {
		return nil, fmt.Errorf("newBitString() %s", err.Error())
	}
	//defer freeBitString(bsC)
	bytes, err := encodePerBuffer(&C.asn_DEF_BIT_STRING, unsafe.Pointer(bsC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func perDecodeBitString(bytes []byte) (*{{.ProtoFileName}}.BitString, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_BIT_STRING)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeBitString((*C.BIT_STRING_t)(unsafePtr))
}

//func newBitString(bs *{{.ProtoFileName}}.BitString) (*C.BIT_STRING_t, error) {
//	numBytes := int(math.Ceil(float64(bs.Len) / 8.0))
//	valAsBytes := make([]byte, 8)
//	binary.LittleEndian.PutUint64(valAsBytes, bs.Value)
//	bitsUnused := numBytes*8 - int(bs.Len)
//
//	bsC := C.BIT_STRING_t{
//		buf:         (*C.uchar)(C.CBytes(valAsBytes[:numBytes])),
//		size:        C.ulong(numBytes),
//		bits_unused: C.int(bitsUnused),
//	}
//	//fmt.Printf("Bit string %+v\n", bsC)
//	return &bsC, nil
//}

// Previously newBitStringFromBytes
// Each BitString should be Octet-aligned, which puts some constraints on it.
// 1. Number of bytes in BitString.Value (of []byte) should fit all bits (defined in BitString.Len)
// For instance, we have 20 bits long BitString. It would require 3 bytes (24 bits) to fit the value.
// Then BitString.Value should be []byte{0x12, 0x34, 0x50} - 3 bytes long
// 2. To encode this in bytes for APER requires an additional 4 zeros at the end (same as bit shifting to the left by 4).
// This additional 4 zeroes shift is Octet alignment - bits should be aligned to the left of available bytes.
// Why 4 bits? 24-20 = 4.
// For more example, see examples in Test_validBitStringsOne in BIT_STRING_test.go
func newBitString(bs *{{.ProtoFileName}}.BitString) (*C.BIT_STRING_t, error) {
	//fmt.Printf("Bit String value is %x\nBitString length (size) is %v\n", bs.Value, bs.Len)
	numBytes := int(math.Ceil(float64(bs.Len) / 8.0))
	//fmt.Printf("Number of bytes is %v\n", numBytes)
	bitsUnused := numBytes*8 - int(bs.Len)
	//fmt.Printf("Number of unused bits is %v\n", bitsUnused)

	if bitsUnused > 7 {
		return nil, fmt.Errorf("bits unused (%d) is greater than 7", bitsUnused)
	}

	if len(bs.Value) < numBytes {
		return nil, fmt.Errorf("%d bytes are required for length %d, found %d", numBytes, bs.Len, len(bs.Value))
	}

	//verification
	mask := byte((1 << bitsUnused) - 1)
	if bs.Value[numBytes-1]&mask > 0 {
		return nil, fmt.Errorf("bit string is NOT octet-aligned - expecting the %d (%d-%d) unused bits to be on the right and equal 0", bitsUnused, numBytes*8, bs.Len)
	}

	bsC := C.BIT_STRING_t{
		buf:         (*C.uchar)(C.CBytes(bs.Value)),
		size:        C.ulong(numBytes),
		bits_unused: C.int(bitsUnused),
	}
	//fmt.Printf("Encoded BitString is %v\n", bsC)

	return &bsC, nil
}

func newBitStringFromArray(array [48]byte) *C.BIT_STRING_t {
	size := binary.LittleEndian.Uint64(array[8:16])
	bitsUnused := int(binary.LittleEndian.Uint32(array[16:20]))
	bytes := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[:8]))), C.int(size))

	bsC := C.BIT_STRING_t{
		buf:         (*C.uchar)(C.CBytes(bytes)),
		size:        C.ulong(size),
		bits_unused: C.int(bitsUnused),
	}

	return &bsC
}

// decodeBitString - byteString in C has 20 bytes
// 8 for a 64bit address of a buffer, 8 for the size in bytes of the buffer uint64, 4 for the unused bits
// The unused bits are at the end of the buffer
func decodeBitString(bsC *C.BIT_STRING_t) (*{{.ProtoFileName}}.BitString, error) {
	size := uint32(bsC.size)
	bitsUnused := uint32(bsC.bits_unused)
	if bitsUnused > 7 {
		return nil, fmt.Errorf("bits unused (%d) is greater than 7", bitsUnused)
	}

	bytes := C.GoBytes(unsafe.Pointer(bsC.buf), C.int(bsC.size))
	// Need to bit shift whole array to the right by bitsUnused
	//var carry byte
	//mask := byte(math.Pow(2, float64(size)) - 1)
	//for i := 0; i < int(size); i++ {
	//prevCarry := carry << (8 - bitsUnused%8)
	//carry = bytes[i] & mask
	//bytes[i] = bytes[i] >> bitsUnused%8
	//goBytes[i] = bytes[i] | prevCarry
	//}
	//fmt.Printf("bit string %x %d %d %+x %+x\n", bufAddr, size, bitsUnused, bytes, goBytes)
	//goBytes := make([]byte, 0)
	//for i := 0; i < int(size); i++ {
	//	//goBytes[i] = bytes[i]
	//	goBytes = append(goBytes, bytes[i])
	//}
	bs := &{{.ProtoFileName}}.BitString{
		Value: bytes,
		Len:   size*8 - bitsUnused,
	}

	return bs, nil
}

//func decodeBitStringBytes(array [8]byte) (*{{.ProtoFileName}}.BitString, error) {
//	bsC := (*C.BIT_STRING_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
//
//	return decodeBitString(bsC)
//}
//
//func freeBitString(bsC *C.BIT_STRING_t) {
//	C.free(unsafe.Pointer(bsC.buf))
//}

