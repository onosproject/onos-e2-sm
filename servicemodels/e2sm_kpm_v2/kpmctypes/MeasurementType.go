// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementType.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeMeasurementType(measurementType *e2sm_kpm_v2.MeasurementType) ([]byte, error) {
	measurementTypeCP, err := newMeasurementType(measurementType)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementType() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementType, unsafe.Pointer(measurementTypeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementType() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementType(measurementType *e2sm_kpm_v2.MeasurementType) ([]byte, error) {
	measurementTypeCP, err := newMeasurementType(measurementType)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementType() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementType, unsafe.Pointer(measurementTypeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementType() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementType(bytes []byte) (*e2sm_kpm_v2.MeasurementType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementType)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementType((*C.MeasurementType_t)(unsafePtr))
}

func perDecodeMeasurementType(bytes []byte) (*e2sm_kpm_v2.MeasurementType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementType)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementType((*C.MeasurementType_t)(unsafePtr))
}

func newMeasurementType(measurementType *e2sm_kpm_v2.MeasurementType) (*C.MeasurementType_t, error) {

	var pr C.MeasurementType_PR
	choiceC := [40]byte{}

	switch choice := measurementType.GetMeasurementType().(type) {
	case *e2sm_kpm_v2.MeasurementType_MeasName:
		pr = C.MeasurementType_PR_measName

		im, err := newMeasurementTypeName(choice.MeasName)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementTypeName() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(uintptr(unsafe.Pointer(im.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:16], uint64(im.size))
	case *e2sm_kpm_v2.MeasurementType_MeasId:
		pr = C.MeasurementType_PR_measID

		im, err := newMeasurementTypeID(choice.MeasId)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementTypeID() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(*im))
	default:
		return nil, fmt.Errorf("newMeasurementType() %T not yet implemented", choice)
	}

	measurementTypeC := C.MeasurementType_t{
		present: pr,
		choice:  choiceC,
	}

	return &measurementTypeC, nil
}

func decodeMeasurementType(measurementTypeC *C.MeasurementType_t) (*e2sm_kpm_v2.MeasurementType, error) {

	measurementType := new(e2sm_kpm_v2.MeasurementType)

	switch measurementTypeC.present {
	case C.MeasurementType_PR_measName:
		var a [16]byte
		copy(a[:], measurementTypeC.choice[0:])
		measurementTypestructC, err := decodeMeasurementTypeNameBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementTypeNameBytes() %s", err.Error())
		}
		measurementType.MeasurementType = &e2sm_kpm_v2.MeasurementType_MeasName{
			MeasName: measurementTypestructC,
		}
	case C.MeasurementType_PR_measID:
		var a [8]byte
		copy(a[:], measurementTypeC.choice[0:])
		measurementTypestructC, err := decodeMeasurementTypeIDBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementTypeIDBytes() %s", err.Error())
		}
		measurementType.MeasurementType = &e2sm_kpm_v2.MeasurementType_MeasId{
			MeasId: measurementTypestructC,
		}
	default:
		return nil, fmt.Errorf("decodeMeasurementType() %v not yet implemented", measurementTypeC.present)
	}

	return measurementType, nil
}
