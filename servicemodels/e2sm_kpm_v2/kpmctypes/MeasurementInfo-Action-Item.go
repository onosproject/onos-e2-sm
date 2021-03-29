// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementInfo-Action-Item.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeMeasurementInfoActionItem(measurementInfoActionItem *e2sm_kpm_v2.MeasurementInfoActionItem) ([]byte, error) {
	measurementInfoActionItemCP, err := newMeasurementInfoActionItem(measurementInfoActionItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementInfoActionItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementInfo_Action_Item, unsafe.Pointer(measurementInfoActionItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementInfoActionItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementInfoActionItem(measurementInfoActionItem *e2sm_kpm_v2.MeasurementInfoActionItem) ([]byte, error) {
	measurementInfoActionItemCP, err := newMeasurementInfoActionItem(measurementInfoActionItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementInfoActionItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementInfo_Action_Item, unsafe.Pointer(measurementInfoActionItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementInfoActionItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementInfoActionItem(bytes []byte) (*e2sm_kpm_v2.MeasurementInfoActionItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementInfo_Action_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementInfoActionItem((*C.MeasurementInfo_Action_Item_t)(unsafePtr))
}

func perDecodeMeasurementInfoActionItem(bytes []byte) (*e2sm_kpm_v2.MeasurementInfoActionItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementInfo_Action_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementInfoActionItem((*C.MeasurementInfo_Action_Item_t)(unsafePtr))
}

func newMeasurementInfoActionItem(measurementInfoActionItem *e2sm_kpm_v2.MeasurementInfoActionItem) (*C.MeasurementInfo_Action_Item_t, error) {

	measNameC, err := newMeasurementTypeName(measurementInfoActionItem.MeasName)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementTypeName() %s", err.Error())
	}

	measurementInfoActionItemC := C.MeasurementInfo_Action_Item_t{
		measName: *measNameC,
		//measID:   measIDC,
	}

	if measurementInfoActionItem.MeasId != nil {
		measurementInfoActionItemC.measID, err = newMeasurementTypeID(measurementInfoActionItem.MeasId)
		if err != nil {
			return nil, fmt.Errorf("newMeasurementTypeId() %s", err.Error())
		}
	}

	return &measurementInfoActionItemC, nil
}

func decodeMeasurementInfoActionItem(measurementInfoActionItemC *C.MeasurementInfo_Action_Item_t) (*e2sm_kpm_v2.MeasurementInfoActionItem, error) {

	measName, err := decodeMeasurementTypeName(&measurementInfoActionItemC.measName)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementTypeName() %s", err.Error())
	}

	measurementInfoActionItem := e2sm_kpm_v2.MeasurementInfoActionItem{
		MeasName: measName,
		//MeasId:   measID,
	}

	if measurementInfoActionItemC.measID != nil {
		measurementInfoActionItem.MeasId, err = decodeMeasurementTypeID(measurementInfoActionItemC.measID)
		if err != nil {
			return nil, fmt.Errorf("decodeMeasurementTypeId() %s", err.Error())
		}
	}

	return &measurementInfoActionItem, nil
}
