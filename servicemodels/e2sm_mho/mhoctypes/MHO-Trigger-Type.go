// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MHO-Trigger-Type.h"
import "C"
import (
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
)

//func xerEncodeMhoTriggerType(mhoTriggerType *e2sm_mho.MhoTriggerType) ([]byte, error) {
//	mhoTriggerTypeCP, err := newMhoTriggerType(mhoTriggerType)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_MHO_Trigger_Type, unsafe.Pointer(mhoTriggerTypeCP)) //ToDo - change name of C-encoder tag
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeMhoTriggerType() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeMhoTriggerType(mhoTriggerType *e2sm_mho.MhoTriggerType) ([]byte, error) {
//	mhoTriggerTypeCP, err := newMhoTriggerType(mhoTriggerType)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_MHO_Trigger_Type, unsafe.Pointer(mhoTriggerTypeCP))
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeMhoTriggerType() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerDecodeMhoTriggerType(bytes []byte) (*e2sm_mho.MhoTriggerType, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MHO_Trigger_Type)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeMhoTriggerType((*C.MHO_Trigger_Type_t)(unsafePtr)) //ToDo - change name of C-struct
//}
//
//func perDecodeMhoTriggerType(bytes []byte) (*e2sm_mho.MhoTriggerType, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MHO_Trigger_Type)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeMhoTriggerType((*C.MHO_Trigger_Type_t)(unsafePtr))
//}

func newMhoTriggerType(mhoTriggerType *e2sm_mho.MhoTriggerType) (*C.MHO_Trigger_Type_t, error) {
	var ret C.MHO_Trigger_Type_t
	switch *mhoTriggerType {
	case e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_PERIODIC:
		ret = C.MHO_Trigger_Type_periodic //ToDo - double-check correctness of the name
	case e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_RCV_MEAS_REPORT:
		ret = C.MHO_Trigger_Type_upon_rcv_meas_report //ToDo - double-check correctness of the name
	case e2sm_mho.MhoTriggerType_MHO_TRIGGER_TYPE_UPON_CHANGE_RRC_STATUS:
		ret = C.MHO_Trigger_Type_upon_change_rrc_status //ToDo - double-check correctness of the name
	default:
		return nil, fmt.Errorf("unexpected MhoTriggerType %v", mhoTriggerType)
	}

	return &ret, nil
}

func decodeMhoTriggerType(mhoTriggerTypeC *C.MHO_Trigger_Type_t) (*e2sm_mho.MhoTriggerType, error) {

	//ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
	mhoTriggerType := e2sm_mho.MhoTriggerType(int32(*mhoTriggerTypeC))

	return &mhoTriggerType, nil
}

//func decodeMhoTriggerTypeBytes(array [8]byte) (*e2sm_mho.MhoTriggerType, error) { //ToDo - Check addressing correct structure in Protobuf
//	mhoTriggerTypeC := (*C.MHO_Trigger_Type_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))
//
//	return decodeMhoTriggerType(mhoTriggerTypeC)
//}
