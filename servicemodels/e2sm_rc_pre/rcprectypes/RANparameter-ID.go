// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANparameter-ID.h"
import "C"

import (
	//"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	//"unsafe"
)

//func xerEncodeRanparameterId(ranparameterId *e2sm_rc_pre_ies.RanparameterId) ([]byte, error) {
//	ranparameterIDCP, err := newRanparameterID(ranparameterId)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterId() %s", err.Error())
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RANparameter_ID, unsafe.Pointer(ranparameterIDCP))
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterId() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeRanparameterId(ranparameterId *e2sm_rc_pre_ies.RanparameterId) ([]byte, error) {
//	ranparameterIDCP, err := newRanparameterID(ranparameterId)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterId() %s", err.Error())
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_RANparameter_ID, unsafe.Pointer(ranparameterIDCP))
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeRanparameterId() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerdecodeRanparameterID(bytes []byte) (*e2sm_rc_pre_ies.RanparameterId, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANparameter_ID)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeRanparameterID((*C.RANparameter_ID_t)(unsafePtr))
//}
//
//func perdecodeRanparameterID(bytes []byte) (*e2sm_rc_pre_ies.RanparameterId, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANparameter_ID)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRanparameterID((*C.RANparameter_ID_t)(unsafePtr))
//}

func newRanparameterID(ranparameterID *e2sm_rc_pre_ies.RanparameterId) (*C.RANparameter_ID_t, error) {

	ranparameterIDC := C.long(ranparameterID.Value)

	return &ranparameterIDC, nil
}

func decodeRanparameterID(ranparameterIDC *C.RANparameter_ID_t) (*e2sm_rc_pre_ies.RanparameterId, error) {

	ranparameterID := e2sm_rc_pre_ies.RanparameterId{
		Value: int32(*ranparameterIDC),
	}

	return &ranparameterID, nil
}
