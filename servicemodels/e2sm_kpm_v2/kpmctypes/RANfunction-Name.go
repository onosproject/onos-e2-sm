// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunction-Name.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeRanfunctionName(ranfunctionName *e2sm_kpm_v2.RanfunctionName) ([]byte, error) {
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

func perEncodeRanfunctionName(ranfunctionName *e2sm_kpm_v2.RanfunctionName) ([]byte, error) {
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

func xerDecodeRanfunctionName(bytes []byte) (*e2sm_kpm_v2.RanfunctionName, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANfunction_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRanfunctionName((*C.RANfunction_Name_t)(unsafePtr))
}

func perDecodeRanfunctionName(bytes []byte) (*e2sm_kpm_v2.RanfunctionName, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANfunction_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRanfunctionName((*C.RANfunction_Name_t)(unsafePtr))
}

func newRanfunctionName(ranfunctionName *e2sm_kpm_v2.RanfunctionName) (*C.RANfunction_Name_t, error) {

	//fmt.Printf("We're inside newRanfunctionName(), starting encoding...")
	//fmt.Printf("Encoding RanFunctionShortName: \n %v \n", ranfunctionName.RanFunctionShortName)
	ranFunctionShortNameC, err := newPrintableString(ranfunctionName.RanFunctionShortName)
	//fmt.Printf("That's the C version of encoded RanFunctionShortName: \n %v \n", ranFunctionShortNameC)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Encoding RanFunctionE2SmOid: \n %v \n", ranfunctionName.RanFunctionE2SmOid)
	ranFunctionE2SmOidC, err := newPrintableString(ranfunctionName.RanFunctionE2SmOid)
	//fmt.Printf("That's the C version of encoded RanFunctionE2SmOid: \n %v \n", ranFunctionE2SmOidC)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Encoding RanFunctionDescription: \n %v \n", ranfunctionName.RanFunctionDescription)
	ranFunctionDescriptionC, err := newPrintableString(ranfunctionName.RanFunctionDescription)
	//fmt.Printf("That's the C version of encoded RanFunctionDescription: \n %v \n", ranFunctionDescriptionC)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Encoding RanFunctionInstance: \n %v \n", ranfunctionName.RanFunctionInstance)
	ranFunctionInstanceC := C.long(ranfunctionName.RanFunctionInstance)
	//fmt.Printf("That's the C version of encoded RanFunctionInstance: \n %v \n", ranFunctionInstanceC)

	//fmt.Printf("Composing final C-structure...")
	ranfunctionNameC := C.RANfunction_Name_t{
		ranFunction_ShortName:   *ranFunctionShortNameC,
		ranFunction_E2SM_OID:    *ranFunctionE2SmOidC,
		ranFunction_Description: *ranFunctionDescriptionC,
		ranFunction_Instance:    &ranFunctionInstanceC,
	}
	//fmt.Printf("This is final C-structure: \n %v \n", ranfunctionNameC)

	return &ranfunctionNameC, nil
}

func decodeRanfunctionName(ranfunctionNameC *C.RANfunction_Name_t) (*e2sm_kpm_v2.RanfunctionName, error) {

	//fmt.Printf("We're inside decodeRanfunctionName(), starting decoding...")
	//fmt.Printf("Decoding RanFunctionShortName: \n %v \n", ranfunctionNameC.ranFunction_ShortName)
	ranFunctionShortName, err := decodePrintableString(&ranfunctionNameC.ranFunction_ShortName)
	//fmt.Printf("That's what was decoded from C: \n %v \n", ranFunctionShortName)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Decoding RanFunctionE2SmOid: \n %v \n", ranfunctionNameC.ranFunction_E2SM_OID)
	ranFunctionE2SmOid, err := decodePrintableString(&ranfunctionNameC.ranFunction_E2SM_OID)
	//fmt.Printf("That's what was decoded from C: \n %v \n", ranFunctionE2SmOid)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Decoding RanFunctionDescription: \n %v \n", ranfunctionNameC.ranFunction_Description)
	ranFunctionDescription, err := decodePrintableString(&ranfunctionNameC.ranFunction_Description)
	//fmt.Printf("That's what was decoded from C: \n %v \n", ranFunctionDescription)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Decoding RanFunctionInstance: \n %v \n", ranfunctionNameC.ranFunction_Instance)
	ranFunctionInstance := int32(*ranfunctionNameC.ranFunction_Instance)
	//fmt.Printf("That's what was decoded from C: \n %v \n", ranFunctionInstance)
	ranfunctionName := e2sm_kpm_v2.RanfunctionName{
		RanFunctionShortName:   ranFunctionShortName,
		RanFunctionE2SmOid:     ranFunctionE2SmOid,
		RanFunctionDescription: ranFunctionDescription,
		RanFunctionInstance:    ranFunctionInstance,
	}
	//fmt.Printf("This is final decoded structure: \n %v \n", ranfunctionName)

	return &ranfunctionName, nil
}
