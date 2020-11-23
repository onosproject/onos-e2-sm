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

//func xerEncodeInteger(integer int64) ([]byte, error) {
//	integerCP := newInteger(integer)
//
//	bytes, err := encodeXer(&C.asn_DEF_INTEGER, unsafe.Pointer(integerCP))
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeInteger() %s", err.Error())
//	}
//	return bytes, nil
//}

func newInteger(msg int64) *C.INTEGER_t {

	intC := newAsnCodecsPrim(msg)

	return intC
}

func decodeInteger(intC *C.INTEGER_t) int64 {

	bytes := decodeAsnCodecsPrim(intC)
	// There could possibly occur an issue in parsing str-to-int. It happens, output value would be 0.
	//byteToInt, _ := strconv.Atoi(string(bytes))

	return bytes //int64(byteToInt)

}