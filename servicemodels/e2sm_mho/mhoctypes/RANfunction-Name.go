// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunction-Name.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeRanfunctionName(ranfunctionName *e2sm_mho.RanfunctionName) ([]byte, error) {
	ranfunctionNameCP, err := newRanfunctionName(ranfunctionName)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanfunctionName() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RANfunction_Name, unsafe.Pointer(ranfunctionNameCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanfunctionName() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRanfunctionName(ranfunctionName *e2sm_mho.RanfunctionName) ([]byte, error) {
	ranfunctionNameCP, err := newRanfunctionName(ranfunctionName)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRanfunctionName() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RANfunction_Name, unsafe.Pointer(ranfunctionNameCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRanfunctionName() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRanfunctionName(bytes []byte) (*e2sm_mho.RanfunctionName, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANfunction_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRanfunctionName((*C.RANfunction_Name_t)(unsafePtr))
}

func perDecodeRanfunctionName(bytes []byte) (*e2sm_mho.RanfunctionName, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANfunction_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRanfunctionName((*C.RANfunction_Name_t)(unsafePtr))
}

func newRanfunctionName(ranfunctionName *e2sm_mho.RanfunctionName) (*C.RANfunction_Name_t, error) {

	var err error
	ranfunctionNameC := C.RANfunction_Name_t{}

	ranFunctionShortNameC, err := newPrintableString(ranfunctionName.RanFunctionShortName)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	ranFunctionE2SmOidC, err := newPrintableString(ranfunctionName.RanFunctionE2SmOid)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	ranFunctionDescriptionC, err := newPrintableString(ranfunctionName.RanFunctionDescription)
	if err != nil {
		return nil, fmt.Errorf("newPrintableString() %s", err.Error())
	}

	//instance is optional
	if ranfunctionName.RanFunctionInstance != nil {

		ranFunctionInstanceC := C.long(ranfunctionName.RanFunctionInstance)
		ranfunctionNameC.ranFunction_Instance = ranFunctionInstanceC
	}
	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	ranfunctionNameC.ranFunction_ShortName = ranFunctionShortNameC
	ranfunctionNameC.ranFunction_E2SM_OID = ranFunctionE2SmOidC
	ranfunctionNameC.ranFunction_Decsription = ranFunctionDescriptionC

	return &ranfunctionNameC, nil
}

func decodeRanfunctionName(ranfunctionNameC *C.RANfunction_Name_t) (*e2sm_mho.RanfunctionName, error) {

	var err error
	ranfunctionName := e2sm_mho.RanfunctionName{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//RanFunctionShortName: ranFunctionShortName,
		//RanFunctionE2SmOid: ranFunctionE2SmOid,
		//RanFunctionDescription: ranFunctionDescription,
		//RanFunctionInstance: ranFunctionInstance,

	}

	ranfunctionName.RanFunctionShortName, err = decodePrintableString(ranfunctionNameC.ranFunction_ShortName)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	ranfunctionName.RanFunctionE2SmOid, err = decodePrintableString(ranfunctionNameC.ranFunction_E2SM_OID)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	ranfunctionName.RanFunctionDescription, err = decodePrintableString(ranfunctionNameC.ranFunction_Decsription)
	if err != nil {
		return nil, fmt.Errorf("decodePrintableString() %s", err.Error())
	}

	//instance is optional
	if ranfunctionNameC.ranFunction_Instance != nil {

		ranfunctionName.RanFunctionInstance = int32(ranfunctionNameC.ranFunction_Instance)
	}

	return &ranfunctionName, nil
}

func decodeRanfunctionNameBytes(array [8]byte) (*e2sm_mho.RanfunctionName, error) {
	ranfunctionNameC := (*C.RANfunction_Name_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRanfunctionName(ranfunctionNameC)
}
