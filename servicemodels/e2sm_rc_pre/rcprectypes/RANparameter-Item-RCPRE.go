// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANparameter-Item-RCPRE.h"
import "C"

import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	//"unsafe"
)

//func xerEncodeRanparameterItem(ranparameterItem *e2sm_rc_pre_v2.RanparameterItem) ([]byte, error) {
//	ranparameterItemCP, err := newRanparameterItem(ranparameterItem)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterItem() %s", err.Error())
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RANparameter_Item_RCPRE, unsafe.Pointer(ranparameterItemCP))
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterItem() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeRanparameterItem(ranparameterItem *e2sm_rc_pre_v2.RanparameterItem) ([]byte, error) {
//	ranparameterItemCP, err := newRanparameterItem(ranparameterItem)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterItem() %s", err.Error())
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_RANparameter_Item_RCPRE, unsafe.Pointer(ranparameterItemCP))
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeRanparameterItem() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerDecodeRanparameterItem(bytes []byte) (*e2sm_rc_pre_v2.RanparameterItem, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANparameter_Item_RCPRE)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeRanparameterItem((*C.RANparameter_Item_RCPRE_t)(unsafePtr))
//}
//
//func perDecodeRanparameterItem(bytes []byte) (*e2sm_rc_pre_v2.RanparameterItem, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANparameter_Item_RCPRE)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRanparameterItem((*C.RANparameter_Item_RCPRE_t)(unsafePtr))
//}

func newRanparameterItem(ranparameterItem *e2sm_rc_pre_v2.RanparameterItem) (*C.RANparameter_Item_RCPRE_t, error) {

	ranparameterIDC, err := newRanparameterID(ranparameterItem.RanParameterId)
	if err != nil {
		return nil, fmt.Errorf("encodeRanparameterItem() %s", err.Error())
	}

	ranparameterItemC := C.RANparameter_Item_RCPRE_t{
		ranParameter_ID: *ranparameterIDC,
	}

	return &ranparameterItemC, nil
}

func decodeRanparameterItem(ranparameterItemC *C.RANparameter_Item_RCPRE_t) (*e2sm_rc_pre_v2.RanparameterItem, error) {

	ranparameterID, err := decodeRanparameterID(&ranparameterItemC.ranParameter_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeRanparameterItem() %s", err.Error())
	}

	ranparameterItem := e2sm_rc_pre_v2.RanparameterItem{
		RanParameterId: ranparameterID,
	}

	return &ranparameterItem, nil
}
