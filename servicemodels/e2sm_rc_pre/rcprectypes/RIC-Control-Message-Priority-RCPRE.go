// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Control-Message-Priority-RCPRE.h"
import "C"

import (
	//"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	//"unsafe"
)

//func xerEncodeRicControlMessagePriority(ricControlMessagePriority *e2sm_rc_pre_v2.RicControlMessagePriority) ([]byte, error) {
//	ricControlMessagePriorityCP, err := newRicControlMessagePriority(ricControlMessagePriority)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRicControlMessagePriority() %s", err.Error())
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RIC_Control_Message_Priority_RCPRE, unsafe.Pointer(ricControlMessagePriorityCP))
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRicControlMessagePriority() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeRicControlMessagePriority(ricControlMessagePriority *e2sm_rc_pre_v2.RicControlMessagePriority) ([]byte, error) {
//	ricControlMessagePriorityCP, err := newRicControlMessagePriority(ricControlMessagePriority)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRicControlMessagePriority() %s", err.Error())
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Control_Message_Priority_RCPRE, unsafe.Pointer(ricControlMessagePriorityCP))
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeRicControlMessagePriority() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerDecodeRicControlMessagePriority(bytes []byte) (*e2sm_rc_pre_v2.RicControlMessagePriority, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Control_Message_Priority_RCPRE)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeRicControlMessagePriority((*C.RIC_Control_Message_Priority_RCPRE_t)(unsafePtr))
//}

//func perDecodeRicControlMessagePriority(bytes []byte) (*e2sm_rc_pre_v2.RicControlMessagePriority, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Control_Message_Priority_RCPRE)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRicControlMessagePriority((*C.RIC_Control_Message_Priority_RCPRE_t)(unsafePtr))
//}

func newRicControlMessagePriority(ricControlMessagePriority *e2sm_rc_pre_v2.RicControlMessagePriority) (*C.RIC_Control_Message_Priority_RCPRE_t, error) {

	ricControlMessagePriorityC := C.long(ricControlMessagePriority.Value)

	return &ricControlMessagePriorityC, nil
}

func decodeRicControlMessagePriority(ricControlMessagePriorityC *C.RIC_Control_Message_Priority_RCPRE_t) (*e2sm_rc_pre_v2.RicControlMessagePriority, error) {

	ricControlMessagePriority := e2sm_rc_pre_v2.RicControlMessagePriority{
		Value: int32(*ricControlMessagePriorityC),
	}

	return &ricControlMessagePriority, nil
}
