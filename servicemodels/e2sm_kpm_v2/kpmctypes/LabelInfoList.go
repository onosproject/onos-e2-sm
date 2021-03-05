// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "LabelInfoList.h"
//#include "LabelInfoItem.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeLabelInfoList(labelInfoList *e2sm_kpm_v2.LabelInfoList) ([]byte, error) {
	labelInfoListCP, err := newLabelInfoList(labelInfoList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeLabelInfoList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_LabelInfoList, unsafe.Pointer(labelInfoListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeLabelInfoList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeLabelInfoList(labelInfoList *e2sm_kpm_v2.LabelInfoList) ([]byte, error) {
	labelInfoListCP, err := newLabelInfoList(labelInfoList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeLabelInfoList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_LabelInfoList, unsafe.Pointer(labelInfoListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeLabelInfoList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeLabelInfoList(bytes []byte) (*e2sm_kpm_v2.LabelInfoList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_LabelInfoList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeLabelInfoList((*C.LabelInfoList_t)(unsafePtr))
}

func perDecodeLabelInfoList(bytes []byte) (*e2sm_kpm_v2.LabelInfoList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_LabelInfoList)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeLabelInfoList((*C.LabelInfoList_t)(unsafePtr))
}

func newLabelInfoList(labelInfoList *e2sm_kpm_v2.LabelInfoList) (*C.LabelInfoList_t, error) {

	labelInfoListC := new(C.LabelInfoList_t)
	for _, item := range labelInfoList.GetValue() {
		itemC, err := newLabelInfoItem(item)
		if err != nil {
			return nil, fmt.Errorf("newLabelInfoItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(labelInfoListC), unsafe.Pointer(itemC)); err != nil {
			return nil, err
		}
	}

	return labelInfoListC, nil
}

func decodeLabelInfoList(labelInfoListC *C.LabelInfoList_t) (*e2sm_kpm_v2.LabelInfoList, error) {

	labelInfoList := new(e2sm_kpm_v2.LabelInfoList)
	var ieCount = int(labelInfoListC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(labelInfoListC.list.array)) * uintptr(i)
		ieC := *(**C.LabelInfoItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(labelInfoListC.list.array)) + offset))
		ie, err := decodeLabelInfoItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeLabelInfoItem() %s", err.Error())
		}
		labelInfoList.Value = append(labelInfoList.Value, ie)
	}

	return labelInfoList, nil
}
