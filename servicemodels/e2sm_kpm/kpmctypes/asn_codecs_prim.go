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
	"encoding/binary"
	"unsafe"
)

// TODO: Change the argument to a []byte
func newAsnCodecsPrim(msg int64) *C.ASN__PRIMITIVE_TYPE_t {

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(msg))
	msgBytes := C.CBytes(b)
	//fmt.Printf("newAsnCodecPrim -- msg -- %v\n", msg)
	//msgBytes := C.CBytes([]byte(string(msg)))
	//fmt.Printf("newAsnCodecPrim -- msgBytes -- %v\n", msgBytes)
	//fmt.Printf("newAsnCodecPrim -- b -- %v\n", b)
	//fmt.Printf("newAsnCodecPrim -- msgBytes -- %v\n", msgBytes)
	//fmt.Printf("newAsnCodecPrim -- len(b) -- %v\n", len(b))
	asnPrimTypeC := C.ASN__PRIMITIVE_TYPE_t{
		buf:  (*C.uchar)(msgBytes),
		size: C.ulong(len(b)),
	}

	return &asnPrimTypeC
}

func decodeAsnCodecsPrim(asnPrimTypeC *C.ASN__PRIMITIVE_TYPE_t) int64 {

	//if asnPrimTypeC == nil {
	//	return nil
	//}
	//fmt.Printf("decodeAsnCodecsPrim -- asnPrimTypeC -- %v\n", asnPrimTypeC)
	bytes := C.GoBytes(unsafe.Pointer(asnPrimTypeC.buf), C.int(asnPrimTypeC.size))
	//fmt.Printf("decodeAsnCodecsPrim -- bytes -- %v\n", bytes)
	//fmt.Printf("decodeAsnCodecsPrim -- string(bytes) -- %v\n", string(bytes))
	result := binary.LittleEndian.Uint64(bytes)
	//fmt.Printf("decodeAsnCodecsPrim -- result -- %v\n", result)
	//fmt.Printf("decodeAsnCodecsPrim -- int64(result) -- %v\n", int64(result))

	return int64(result)
}
