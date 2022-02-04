// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-IndicationMessage-Format1.h"
//#include "PM-Containers-List.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMsgFormat1 *e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1) ([]byte, error) {

	e2SmKpmIndicationMsgFormat1CP, err := newE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMsgFormat1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessageFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_IndicationMessage_Format1, unsafe.Pointer(e2SmKpmIndicationMsgFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2SmKpmIndicationMessageFormat1() %s", err.Error())
	}
	return bytes, nil
}

func newE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMsgFormat1 *e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1) (*C.E2SM_KPM_IndicationMessage_Format1_t, error) {
	pmContainersListC := new(C.E2SM_KPM_IndicationMessage_Format1_t)
	for _, pmContainersItem := range e2SmKpmIndicationMsgFormat1.IndicationMessageFormat1.GetPmContainers() {
		pmContainersItemC, err := newPmContainersListItem(pmContainersItem)
		if err != nil {
			return nil, fmt.Errorf("newE2SmKpmIndicationMessageFormat1() %s", err.Error())
		}

		if _, err = C.asn_sequence_add(unsafe.Pointer(pmContainersListC), unsafe.Pointer(pmContainersItemC)); err != nil {
			return nil, err
		}
	}

	return pmContainersListC, nil
}

func decodeE2SmKpmIndicationMessageFormat1(e2SmIindicationMsgFormat1C *C.E2SM_KPM_IndicationMessage_Format1_t) (*e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1, error) {
	e2SmIindicationMsgFormat1 := &e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2sm_kpm_ies.E2SmKpmIndicationMessageFormat1{
			PmContainers: make([]*e2sm_kpm_ies.PmContainersList, 0),
		},
	}

	ieCount := int(e2SmIindicationMsgFormat1C.pm_Containers.list.count)
	//fmt.Printf("RanFunctionListC %T List %T %v Array %T %v Deref %v\n", rflC, rflC.list, rflC.list, rflC.list.array, *rflC.list.array, *(rflC.list.array))
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*e2SmIindicationMsgFormat1C.pm_Containers.list.array)) * uintptr(i)
		pmContainersListItemC := *(**C.PM_Containers_List_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmIindicationMsgFormat1C.pm_Containers.list.array)) + offset))
		//fmt.Printf("Value %T %p %v\n", rfiIeC, rfiIeC, rfiIeC)
		pmContainersListItem, err := decodePmContainersListItem(pmContainersListItemC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmKpmIndicationMessageFormat1() %s", err.Error())
		}
		e2SmIindicationMsgFormat1.IndicationMessageFormat1.PmContainers = append(e2SmIindicationMsgFormat1.IndicationMessageFormat1.PmContainers, pmContainersListItem)
	}

	return e2SmIindicationMsgFormat1, nil
}

func decodeE2SmKpmIndicationMessageFormat1Bytes(array [8]byte) (*e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1, error) {
	eSmKpmIndicationMsgFormat1C := (*C.E2SM_KPM_IndicationMessage_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	result, err := decodeE2SmKpmIndicationMessageFormat1(eSmKpmIndicationMsgFormat1C)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmKpmIndicationMessageFormat1Bytes() %s", err.Error())
	}

	return result, nil
}
