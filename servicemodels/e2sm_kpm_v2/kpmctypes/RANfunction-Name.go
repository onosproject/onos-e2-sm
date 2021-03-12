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
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
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

	ranFunctionShortNameC, err := newPrintableString(ranfunctionName.RanFunctionShortName)
	if err != nil {
		return nil, err
	}
	ranFunctionE2SmOidC, err := newPrintableString(ranfunctionName.RanFunctionE2SmOid)
	if err != nil {
		return nil, err
	}
	ranFunctionDescriptionC, err := newPrintableString(ranfunctionName.RanFunctionDescription)
	if err != nil {
		return nil, err
	}
	ranFunctionInstanceC := C.long(ranfunctionName.RanFunctionInstance)
	ranfunctionNameC := C.RANfunction_Name_t{
		ranFunction_ShortName:   *ranFunctionShortNameC,
		ranFunction_E2SM_OID:    *ranFunctionE2SmOidC,
		ranFunction_Description: *ranFunctionDescriptionC,
		ranFunction_Instance:    &ranFunctionInstanceC,
	}

	return &ranfunctionNameC, nil
}

func decodeRanfunctionName(ranfunctionNameC *C.RANfunction_Name_t) (*e2sm_kpm_v2.RanfunctionName, error) {

	ranFunctionShortName, err := decodePrintableString(&ranfunctionNameC.ranFunction_ShortName)
	if err != nil {
		return nil, err
	}
	ranFunctionE2SmOid, err := decodePrintableString(&ranfunctionNameC.ranFunction_E2SM_OID)
	if err != nil {
		return nil, err
	}
	ranFunctionDescription, err := decodePrintableString(&ranfunctionNameC.ranFunction_Description)
	if err != nil {
		return nil, err
	}
	ranFunctionInstance := int32(*ranfunctionNameC.ranFunction_Instance)
	ranfunctionName := e2sm_kpm_v2.RanfunctionName{
		RanFunctionShortName:   ranFunctionShortName,
		RanFunctionE2SmOid:     ranFunctionE2SmOid,
		RanFunctionDescription: ranFunctionDescription,
		RanFunctionInstance:    ranFunctionInstance,
	}

	return &ranfunctionName, nil
}
