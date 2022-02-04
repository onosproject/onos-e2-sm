// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANparameter-ID-RCPRE.h"
import "C"

import (
	//"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	//"unsafe"
)

//func xerEncodeRanparameterId(ranparameterId *e2sm_rc_pre_v2.RanparameterId) ([]byte, error) {
//	ranparameterIDCP, err := newRanparameterID(ranparameterId)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterId() %s", err.Error())
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RANparameter_ID_RCPRE, unsafe.Pointer(ranparameterIDCP))
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterId() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeRanparameterId(ranparameterId *e2sm_rc_pre_v2.RanparameterId) ([]byte, error) {
//	ranparameterIDCP, err := newRanparameterID(ranparameterId)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterId() %s", err.Error())
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_RANparameter_ID_RCPRE, unsafe.Pointer(ranparameterIDCP))
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeRanparameterId() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerdecodeRanparameterID(bytes []byte) (*e2sm_rc_pre_v2.RanparameterId, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANparameter_ID_RCPRE)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeRanparameterID((*C.RANparameter_ID_RCPRE_t)(unsafePtr))
//}
//
//func perdecodeRanparameterID(bytes []byte) (*e2sm_rc_pre_v2.RanparameterId, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANparameter_ID_RCPRE)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRanparameterID((*C.RANparameter_ID_RCPRE_t)(unsafePtr))
//}

func newRanparameterID(ranparameterID *e2sm_rc_pre_v2.RanparameterId) (*C.RANparameter_ID_RCPRE_t, error) {

	ranparameterIDC := C.long(ranparameterID.Value)

	return &ranparameterIDC, nil
}

func decodeRanparameterID(ranparameterIDC *C.RANparameter_ID_RCPRE_t) (*e2sm_rc_pre_v2.RanparameterId, error) {

	ranparameterID := e2sm_rc_pre_v2.RanparameterId{
		Value: int32(*ranparameterIDC),
	}

	return &ranparameterID, nil
}
