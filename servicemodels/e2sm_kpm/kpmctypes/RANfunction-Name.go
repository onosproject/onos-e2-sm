// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunction-Name.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeRanfunctionName(ranfunctionName *e2sm_kpm_ies.RanfunctionName) ([]byte, error) {
	ranfunctionNameCP := newRanfunctionName(ranfunctionName)

	bytes, err := encodeXer(&C.asn_DEF_RANfunction_Name, unsafe.Pointer(ranfunctionNameCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanfunctionName() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRanfunctionName(ranfunctionName *e2sm_kpm_ies.RanfunctionName) ([]byte, error) {
	ranfunctionNameCP := newRanfunctionName(ranfunctionName)

	bytes, err := encodePerBuffer(&C.asn_DEF_RANfunction_Name, unsafe.Pointer(ranfunctionNameCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRanfunctionName() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRanfunctionName(bytes []byte) (*e2sm_kpm_ies.RanfunctionName, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANfunction_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRanfunctionName((*C.RANfunction_Name_t)(unsafePtr)), nil
}

func perDecodeRanfunctionName(bytes []byte) (*e2sm_kpm_ies.RanfunctionName, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANfunction_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRanfunctionName((*C.RANfunction_Name_t)(unsafePtr)), nil
}

func newRanfunctionName(ranfunctionName *e2sm_kpm_ies.RanfunctionName) *C.RANfunction_Name_t {

	ranFunctionShortNameC := newPrintableString(ranfunctionName.RanFunctionShortName)
	ranFunctionE2SmOidC := newPrintableString(ranfunctionName.RanFunctionE2SmOid)
	ranFunctionDescriptionC := newPrintableString(ranfunctionName.RanFunctionDescription)
	ranFunctionInstanceC := (C.long)(ranfunctionName.RanFunctionInstance)

	ranfunctionNameC := C.RANfunction_Name_t{
		ranFunction_ShortName:   *ranFunctionShortNameC,
		ranFunction_E2SM_OID:    *ranFunctionE2SmOidC,
		ranFunction_Description: *ranFunctionDescriptionC,
		ranFunction_Instance:    &ranFunctionInstanceC,
	}

	return &ranfunctionNameC
}

func decodeRanfunctionName(ranfunctionNameC *C.RANfunction_Name_t) *e2sm_kpm_ies.RanfunctionName {

	ranFunctionShortName := decodePrintableString(&ranfunctionNameC.ranFunction_ShortName)
	ranFunctionE2SmOid := decodePrintableString(&ranfunctionNameC.ranFunction_E2SM_OID)
	ranFunctionDescription := decodePrintableString(&ranfunctionNameC.ranFunction_Description)

	ranfunctionName := e2sm_kpm_ies.RanfunctionName{
		RanFunctionShortName:   ranFunctionShortName,
		RanFunctionE2SmOid:     ranFunctionE2SmOid,
		RanFunctionDescription: ranFunctionDescription,
	}
	// instance is optional
	if ranfunctionNameC.ranFunction_Instance != nil {
		ranfunctionName.RanFunctionInstance = int32(*ranfunctionNameC.ranFunction_Instance)
	}

	return &ranfunctionName
}
