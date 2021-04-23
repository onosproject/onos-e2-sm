// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RC-PRE-Trigger-Type.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	//"unsafe"
)

//func xerEncodeRcPreTriggerType(rcPreTriggerType *e2sm_rc_pre_v2.RcPreTriggerType) ([]byte, error) {
//	rcPreTriggerTypeCP, err := newRcPreTriggerType(rcPreTriggerType)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RC_PRE_Trigger_Type, unsafe.Pointer(rcPreTriggerTypeCP)) //ToDo - change name of C-encoder tag
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRcPreTriggerType() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeRcPreTriggerType(rcPreTriggerType *e2sm_rc_pre_v2.RcPreTriggerType) ([]byte, error) {
//	rcPreTriggerTypeCP, err := newRcPreTriggerType(rcPreTriggerType)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_RC_PRE_Trigger_Type, unsafe.Pointer(rcPreTriggerTypeCP)) //ToDo - change name of C-encoder tag
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeRcPreTriggerType() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerDecodeRcPreTriggerType(bytes []byte) (*e2sm_rc_pre_v2.RcPreTriggerType, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RC_PRE_Trigger_Type)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeRcPreTriggerType((*C.RC_PRE_Trigger_Type_t)(unsafePtr)) //ToDo - change name of C-struct
//}
//
//func perDecodeRcPreTriggerType(bytes []byte) (*e2sm_rc_pre_v2.RcPreTriggerType, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RC_PRE_Trigger_Type)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRcPreTriggerType((*C.RC_PRE_Trigger_Type_t)(unsafePtr)) //ToDo - change name of C-struct
//}

func newRcPreTriggerType(rcPreTriggerType *e2sm_rc_pre_v2.RcPreTriggerType) (*C.RC_PRE_Trigger_Type_t, error) { //ToDo - change name of C-struct
	var ret C.RC_PRE_Trigger_Type_t
	switch *rcPreTriggerType {
	case e2sm_rc_pre_v2.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_UPON_CHANGE:
		ret = C.RC_PRE_Trigger_Type_upon_change
	case e2sm_rc_pre_v2.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_PERIODIC:
		ret = C.RC_PRE_Trigger_Type_periodic
	default:
		return nil, fmt.Errorf("unexpected RC-PRE-Trigger-Type %v", rcPreTriggerType)
	}

	return &ret, nil
}

func decodeRcPreTriggerType(rcPreTriggerTypeC *C.RC_PRE_Trigger_Type_t) (*e2sm_rc_pre_v2.RcPreTriggerType, error) { //ToDo - change name of C-struct

	//ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
	rcPreTriggerType := e2sm_rc_pre_v2.RcPreTriggerType(int32(*rcPreTriggerTypeC))

	return &rcPreTriggerType, nil
}
