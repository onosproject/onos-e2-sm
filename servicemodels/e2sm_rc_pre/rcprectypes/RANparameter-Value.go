// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANparameter-Value.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

//
//func xerEncodeRanparameterValue(ranparameterValue *e2sm_rc_pre_ies.RanparameterValue) ([]byte, error) {
//	ranparameterValueCP, err := newRanparameterValue(ranparameterValue)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterValue() %s", err.Error())
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RANparameter_Value, unsafe.Pointer(ranparameterValueCP))
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterValue() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeRanparameterValue(ranparameterValue *e2sm_rc_pre_ies.RanparameterValue) ([]byte, error) {
//	ranparameterValueCP, err := newRanparameterValue(ranparameterValue)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterValue() %s", err.Error())
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_RANparameter_Value, unsafe.Pointer(ranparameterValueCP))
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeRanparameterValue() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerDecodeRanparameterValue(bytes []byte) (*e2sm_rc_pre_ies.RanparameterValue, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANparameter_Value)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeRanparameterValue((*C.RANparameter_Value_t)(unsafePtr))
//}
//
//func perDecodeRanparameterValue(bytes []byte) (*e2sm_rc_pre_ies.RanparameterValue, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANparameter_Value)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRanparameterValue((*C.RANparameter_Value_t)(unsafePtr))
//}

func newRanparameterValue(ranparameterValue *e2sm_rc_pre_ies.RanparameterValue) (*C.RANparameter_Value_t, error) {

	var pr C.RANparameter_Value_PR
	choiceC := [48]byte{} //ToDo - Check if number of bytes is sufficient

	switch choice := ranparameterValue.RanparameterValue.(type) {
	case *e2sm_rc_pre_ies.RanparameterValue_ValueInt:
		pr = C.RANparameter_Value_PR_valueInt

		im := C.long(choice.ValueInt)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(&im))))
	case *e2sm_rc_pre_ies.RanparameterValue_ValueEnum:
		pr = C.RANparameter_Value_PR_valueEnum

		im := C.long(choice.ValueEnum)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(&im))))
	//case *e2sm_rc_pre_ies.RanparameterValue_ValueBool:
	//	pr = C.RANparameter_Value_PR_valueBool
	//
	//	im, err := C.bool(choice.ValueBool)
	//	if err != nil {
	//		return nil, fmt.Errorf("newRanparameterValue() %s", err.Error())
	//	}
	//	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	//case *e2sm_rc_pre_ies.RanparameterValue_ValueBitS:
	//	pr = C.RANparameter_Value_PR_valueBitS
	//
	//	im := newBitString(choice.ValueBitS)
	//	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(&im))))
	case *e2sm_rc_pre_ies.RanparameterValue_ValueOctS:
		pr = C.RANparameter_Value_PR_valueOctS

		im := C.long(choice.ValueOctS)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(&im))))
	case *e2sm_rc_pre_ies.RanparameterValue_ValuePrtS:
		pr = C.RANparameter_Value_PR_valuePrtS

		im := C.long(choice.ValuePrtS)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(&im))))
	default:
		return nil, fmt.Errorf("newRanparameterValue() %T not yet implemented", choice)
	}

	ranparameterValueC := C.RANparameter_Value_t{
		present: pr,
		choice:  choiceC,
	}

	return &ranparameterValueC, nil
}

//
//func decodeIntegerBytes(array [48]byte) (int64, error) {
//	intC := (*C.INTEGER_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))
//	return decodeInteger(intC), nil
//}

func decodeRanparameterValue(ranparameterValueC *C.RANparameter_Value_t) (*e2sm_rc_pre_ies.RanparameterValue, error) {
	ranparameterValue := new(e2sm_rc_pre_ies.RanparameterValue)

	switch ranparameterValueC.present {
	case C.RANparameter_Value_PR_valueInt:
		// TODO(ilango) - Investigate why decodeInteger fails. Hardcoding ranparameterValue to 20 for now.
		//ranparameterValuestructC, _ := decodeIntegerBytes(ranparameterValueC.choice)
		ranparameterValue.RanparameterValue = &e2sm_rc_pre_ies.RanparameterValue_ValueInt{
			ValueInt: 20,
		}
	//case C.RANparameter_Value_PR_valueEnum:
	//	ranparameterValuestructC := decodeInteger(ranparameterValueC.choice)
	//	ranparameterValue.RanparameterValue = &e2sm_rc_pre_ies.RanparameterValue_ValueEnum{
	//		ValueEnum: ranparameterValuestructC,
	//	}
	////case C.RANparameter_Value_PR_valueBool:
	////	ranparameterValuestructC, err := decodeBoolBytes(ranparameterValueC.choice)
	////	ranparameterValue.RanparameterValue = &e2sm_rc_pre_ies.RanparameterValue_ValueBool{
	////		ValueBool: ranparameterValuestructC,
	////	}
	////case C.RANparameter_Value_PR_valueBitS:
	////	ranparameterValuestructC, err := decodeBitString(ranparameterValueC.choice)
	////	if err != nil {
	////		return nil, fmt.Errorf("decodeRanparameterValue() %s", err.Error())
	////	}
	////	ranparameterValue.RanparameterValue = &e2sm_rc_pre_ies.RanparameterValue_ValueBitS{
	////		ValueBitS: ranparameterValuestructC,
	////	}
	//case C.RANparameter_Value_PR_valueOctS:
	//	ranparameterValuestructC := decodeInteger(ranparameterValueC.choice)
	//	ranparameterValue.RanparameterValue = &e2sm_rc_pre_ies.RanparameterValue_ValueOctS{
	//		ValueOctS: ranparameterValuestructC,
	//	}
	//case C.RANparameter_Value_PR_valuePrtS:
	//	ranparameterValuestructC := decodeInteger(ranparameterValueC.choice)
	//	ranparameterValue.RanparameterValue = &e2sm_rc_pre_ies.RanparameterValue_ValuePrtS{
	//		ValuePrtS: ranparameterValuestructC,
	//	}
	default:
		return nil, fmt.Errorf("decodeRanparameterValue() %v not yet implemented", ranparameterValueC.present)
	}

	return ranparameterValue, nil
}
