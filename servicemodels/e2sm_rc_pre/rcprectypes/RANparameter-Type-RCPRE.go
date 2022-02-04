// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANparameter-Type-RCPRE.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	//"unsafe"
)

//func xerEncodeRanparameterType(ranparameterType *e2sm_rc_pre_v2.RanparameterType) ([]byte, error) {
//	ranparameterTypeCP, err := newRanparameterType(ranparameterType)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RANparameter_Type_RCPRE, unsafe.Pointer(ranparameterTypeCP))
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterType() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeRanparameterType(ranparameterType *e2sm_rc_pre_v2.RanparameterType) ([]byte, error) {
//	ranparameterTypeCP, err := newRanparameterType(ranparameterType)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_RANparameter_Type_RCPRE, unsafe.Pointer(ranparameterTypeCP))
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeRanparameterType() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerDecodeRanparameterType(bytes []byte) (*e2sm_rc_pre_v2.RanparameterType, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANparameter_Type_RCPRE)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeRanparameterType((*C.RANparameter_Type_RCPRE_t)(unsafePtr))
//}
//
//func perDecodeRanparameterType(bytes []byte) (*e2sm_rc_pre_v2.RanparameterType, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANparameter_Type_RCPRE)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRanparameterType((*C.RANparameter_Type_RCPRE_t)(unsafePtr))
//}

func newRanparameterType(ranparameterType *e2sm_rc_pre_v2.RanparameterType) (*C.RANparameter_Type_RCPRE_t, error) {
	var ret C.RANparameter_Type_RCPRE_t
	switch *ranparameterType {
	case e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_INTEGER:
		ret = C.RANparameter_Type_RCPRE_integer
	case e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_ENUMERATED:
		ret = C.RANparameter_Type_RCPRE_enumerated
	case e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_BOOLEAN:
		ret = C.RANparameter_Type_RCPRE_boolean
	case e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_BIT_STRING:
		ret = C.RANparameter_Type_RCPRE_bit_string
	case e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_OCTET_STRING:
		ret = C.RANparameter_Type_RCPRE_octet_string
	case e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_PRINTABLE_STRING:
		ret = C.RANparameter_Type_RCPRE_printable_string
	default:
		return nil, fmt.Errorf("unexpected RANparameter-Type %v", ranparameterType)
	}

	return &ret, nil
}

func decodeRanparameterType(ranparameterTypeC *C.RANparameter_Type_RCPRE_t) (*e2sm_rc_pre_v2.RanparameterType, error) {

	ranparameterType := e2sm_rc_pre_v2.RanparameterType(int32(*ranparameterTypeC))

	return &ranparameterType, nil
}
