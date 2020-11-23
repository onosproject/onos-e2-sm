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
	"strconv"
)

func newInteger(msg []byte) *C.INTEGER_t {

	intC := newAsnCodecsPrim(msg)

	return intC
}

func decodeInteger(intC *C.INTEGER_t) uint64 {

	bytes := decodeAsnCodecsPrim(intC)
	// There could possibly occur an issue in parsing str-to-int. Iit happens, output value would be 0.
	byteToInt, _ := strconv.Atoi(string(bytes))

	return uint64(byteToInt)

}