// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PrintableString.h"
import "C"

import (
	"fmt"
)

// TODO: Change the argument to a []byte
func newPrintableString(msg string) (*C.PrintableString_t, error) {
	// PrintableString is defined via OctetString --> see PrintableString.h
	//fmt.Printf("We're inside newPrintableString(), starting encoding...")
	//fmt.Printf("Encoding PrintableString: \n %v \n", msg)
	prntStrC, err := newOctetString(msg)
	//fmt.Printf("That's the C version of encoded PrintableString: \n %v \n", prntStrC)
	if err != nil {
		return nil, fmt.Errorf("newOctetString() %s", err.Error())
	}

	return prntStrC, nil
}

func decodePrintableString(octC *C.PrintableString_t) (string, error) {

	//fmt.Printf("We're inside decodePrintableString(), starting decoding...")
	//fmt.Printf("Decoding PrintableString: \n %v \n", octC)
	bytes, err := decodeOctetString(octC)
	//fmt.Printf("That's what was decoded from C: \n %v \n", bytes)
	if err != nil {
		return "", fmt.Errorf("decodeOctetString() %s", err.Error())
	}

	return bytes, nil
}

func newPrintableStringFromArray(array [16]byte) *C.PrintableString_t {

	//fmt.Printf("We're inside newPrintableStringFromArray(), starting encoding...")
	//fmt.Printf("Creating PrintableString from bytes: \n %v \n", array)
	prtStrC := newOctetStringFromArray(array)
	//fmt.Printf("That's the C version of PrintableString encoded from bytes: \n %v \n", prtStrC)

	return prtStrC
}

func decodePrintableStringBytes(array [16]byte) (string, error) {
	prtStrC := newPrintableStringFromArray(array)

	return decodePrintableString(prtStrC)
}
