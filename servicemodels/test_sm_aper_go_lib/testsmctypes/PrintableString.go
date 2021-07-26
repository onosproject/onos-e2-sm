// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package testsmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PrintableString.h"
import "C"

import "fmt"

func newPrintableString(msg string) (*C.PrintableString_t, error) {
	// PrintableString is defined via OctetString --> see PrintableString.h
	prntStrC, err := newOctetString([]byte(msg))
	if err != nil {
		return nil, fmt.Errorf("newOctetString() %s", err.Error())
	}

	return prntStrC, nil
}

func decodePrintableString(octC *C.PrintableString_t) (string, error) {

	bytes, err := decodeOctetString(octC)
	if err != nil {
		return "", fmt.Errorf("decodeOctetString() %s", err.Error())
	}

	return string(bytes), nil
}

func newPrintableStringFromArray(array [16]byte) *C.PrintableString_t {

	prtStrC := newOctetStringFromArray(array)
	return prtStrC
}

func decodePrintableStringBytes(array [16]byte) (string, error) {
	prtStrC := newPrintableStringFromArray(array)

	return decodePrintableString(prtStrC)
}
