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

func newInteger(msg string) *C.INTEGER_t {

	intC := newAsnCodecsPrim(msg)

	return intC
}

func decodeInteger(intC *C.INTEGER_t) string {

	bytes := decodeAsnCodecsPrim(intC)
	return bytes
}