// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementRecordItem.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeMeasurementRecordItem(measurementRecordItem *e2sm_kpm_v2.MeasurementRecordItem) ([]byte, error) {
	measurementRecordItemCP, err := newMeasurementRecordItem(measurementRecordItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementRecordItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementRecordItem, unsafe.Pointer(measurementRecordItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementRecordItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementRecordItem(measurementRecordItem *e2sm_kpm_v2.MeasurementRecordItem) ([]byte, error) {
	measurementRecordItemCP, err := newMeasurementRecordItem(measurementRecordItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementRecordItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementRecordItem, unsafe.Pointer(measurementRecordItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementRecordItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementRecordItem(bytes []byte) (*e2sm_kpm_v2.MeasurementRecordItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementRecordItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementRecordItem((*C.MeasurementRecordItem_t)(unsafePtr))
}

func perDecodeMeasurementRecordItem(bytes []byte) (*e2sm_kpm_v2.MeasurementRecordItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementRecordItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementRecordItem((*C.MeasurementRecordItem_t)(unsafePtr))
}

func newMeasurementRecordItem(measurementRecordItem *e2sm_kpm_v2.MeasurementRecordItem) (*C.MeasurementRecordItem_t, error) {

	var pr C.MeasurementRecordItem_PR
	choiceC := [16]byte{}

	switch choice := measurementRecordItem.MeasurementRecordItem.(type) {
	case *e2sm_kpm_v2.MeasurementRecordItem_Integer:
		pr = C.MeasurementRecordItem_PR_integer

		im, err := newInteger(choice.Integer)
		if err != nil {
			return nil, fmt.Errorf("newInteger() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_kpm_v2.MeasurementRecordItem_Real:
		pr = C.MeasurementRecordItem_PR_real

		im, err := newReal(choice.Real)
		if err != nil {
			return nil, fmt.Errorf("newReal() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_kpm_v2.MeasurementRecordItem_NoValue:
		pr = C.MeasurementRecordItem_PR_noValue

		im := newNull()
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newMeasurementRecordItem() %T not yet implemented", choice)
	}

	measurementRecordItemC := C.MeasurementRecordItem_t{
		present: pr,
		choice:  choiceC,
	}

	return &measurementRecordItemC, nil
}

func decodeMeasurementRecordItem(measurementRecordItemC *C.MeasurementRecordItem_t) (*e2sm_kpm_v2.MeasurementRecordItem, error) {

	//This is Decoder part (OneOf)
	measurementRecordItem := new(e2sm_kpm_v2.MeasurementRecordItem)

	switch measurementRecordItemC.present {
	case C.MeasurementRecordItem_PR_integer:
		var a [8]byte
		copy(a[:], measurementRecordItemC.choice[0:8])
		measRecordItem, err := decodeIntegerBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeIntegerBytes() %s", err.Error())
		}
		measurementRecordItem.MeasurementRecordItem = &e2sm_kpm_v2.MeasurementRecordItem_Integer{
			Integer: measRecordItem,
		}
	case C.MeasurementRecordItem_PR_real:
		var a [8]byte
		copy(a[:], measurementRecordItemC.choice[0:8])
		measRecordItem, err := decodeRealBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeRealBytes() %s", err.Error())
		}
		measurementRecordItem.MeasurementRecordItem = &e2sm_kpm_v2.MeasurementRecordItem_Real{
			Real: measRecordItem,
		}
	case C.MeasurementRecordItem_PR_noValue:
		null := decodeNull()
		measurementRecordItem.MeasurementRecordItem = &e2sm_kpm_v2.MeasurementRecordItem_NoValue{
			NoValue: null,
		}
	default:
		return nil, fmt.Errorf("decodeMeasurementRecordItem() %v not yet implemented", measurementRecordItemC.present)
	}

	return measurementRecordItem, nil
}

func decodeMeasurementRecordItemBytes(array [16]byte) (*e2sm_kpm_v2.MeasurementRecordItem, error) {
	measurementRecordItemC := (*C.MeasurementRecordItem_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))

	return decodeMeasurementRecordItem(measurementRecordItemC)
}
