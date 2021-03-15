// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "LabelInfoItem.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeLabelInfoItem(labelInfoItem *e2sm_kpm_v2.LabelInfoItem) ([]byte, error) {
	labelInfoItemCP, err := newLabelInfoItem(labelInfoItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeLabelInfoItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_LabelInfoItem, unsafe.Pointer(labelInfoItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeLabelInfoItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeLabelInfoItem(labelInfoItem *e2sm_kpm_v2.LabelInfoItem) ([]byte, error) {
	labelInfoItemCP, err := newLabelInfoItem(labelInfoItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeLabelInfoItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_LabelInfoItem, unsafe.Pointer(labelInfoItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeLabelInfoItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeLabelInfoItem(bytes []byte) (*e2sm_kpm_v2.LabelInfoItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_LabelInfoItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeLabelInfoItem((*C.LabelInfoItem_t)(unsafePtr))
}

func perDecodeLabelInfoItem(bytes []byte) (*e2sm_kpm_v2.LabelInfoItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_LabelInfoItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeLabelInfoItem((*C.LabelInfoItem_t)(unsafePtr))
}

func newLabelInfoItem(labelInfoItem *e2sm_kpm_v2.LabelInfoItem) (*C.LabelInfoItem_t, error) {

	measLabelC, err := newMeasurementLabel(labelInfoItem.MeasLabel)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementLabel() %s", err.Error())
	}

	labelInfoItemC := C.LabelInfoItem_t{
		measLabel: *measLabelC,
	}

	return &labelInfoItemC, nil
}

func decodeLabelInfoItem(labelInfoItemC *C.LabelInfoItem_t) (*e2sm_kpm_v2.LabelInfoItem, error) {

	measLabel, err := decodeMeasurementLabel(&labelInfoItemC.measLabel)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementLabel() %s", err.Error())
	}

	labelInfoItem := e2sm_kpm_v2.LabelInfoItem{
		MeasLabel: measLabel,
	}

	return &labelInfoItem, nil
}
