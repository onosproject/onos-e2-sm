// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunction-Name-RCPRE.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeRanfunctionName(ranfunctionName *e2sm_rc_pre_v2.RanfunctionName) ([]byte, error) {
	ranfunctionNameCP := newRanfunctionName(ranfunctionName)

	bytes, err := encodeXer(&C.asn_DEF_RANfunction_Name_RCPRE, unsafe.Pointer(ranfunctionNameCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanfunctionName() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRanfunctionName(ranfunctionName *e2sm_rc_pre_v2.RanfunctionName) ([]byte, error) {
	ranfunctionNameCP := newRanfunctionName(ranfunctionName)

	bytes, err := encodePerBuffer(&C.asn_DEF_RANfunction_Name_RCPRE, unsafe.Pointer(ranfunctionNameCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRanfunctionName() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRanfunctionName(bytes []byte) (*e2sm_rc_pre_v2.RanfunctionName, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANfunction_Name_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRanfunctionName((*C.RANfunction_Name_RCPRE_t)(unsafePtr)), nil
}

func perDecodeRanfunctionName(bytes []byte) (*e2sm_rc_pre_v2.RanfunctionName, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANfunction_Name_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRanfunctionName((*C.RANfunction_Name_RCPRE_t)(unsafePtr)), nil
}

func newRanfunctionName(ranfunctionName *e2sm_rc_pre_v2.RanfunctionName) *C.RANfunction_Name_RCPRE_t {

	ranFunctionShortNameC := newPrintableString(ranfunctionName.GetRanFunctionShortName())
	ranFunctionE2SmOidC := newPrintableString(ranfunctionName.GetRanFunctionE2SmOid())
	ranFunctionDescriptionC := newPrintableString(ranfunctionName.GetRanFunctionDescription())
	//ranFunctionInstanceC := (C.long)(ranfunctionName.RanFunctionInstance)

	ranfunctionNameC := C.RANfunction_Name_RCPRE_t{
		ranFunction_ShortName:   *ranFunctionShortNameC,
		ranFunction_E2SM_OID:    *ranFunctionE2SmOidC,
		ranFunction_Description: *ranFunctionDescriptionC,
		//ranFunction_Instance:    &ranFunctionInstanceC,
	}

	//instance is optional
	if ranfunctionName.GetRanFunctionInstance() != -1 {
		rfi := C.long(ranfunctionName.GetRanFunctionInstance())
		ranfunctionNameC.ranFunction_Instance = &rfi
	}

	return &ranfunctionNameC
}

func decodeRanfunctionName(ranfunctionNameC *C.RANfunction_Name_RCPRE_t) *e2sm_rc_pre_v2.RanfunctionName {

	ranFunctionShortName := decodePrintableString(&ranfunctionNameC.ranFunction_ShortName)
	ranFunctionE2SmOid := decodePrintableString(&ranfunctionNameC.ranFunction_E2SM_OID)
	ranFunctionDescription := decodePrintableString(&ranfunctionNameC.ranFunction_Description)

	ranfunctionName := e2sm_rc_pre_v2.RanfunctionName{
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
