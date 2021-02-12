// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-ControlOutcome-Format1.h"
//#include "RANparameter-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func XerEncodeE2SmRcPreControlOutcomeFormat1(e2SmRcPreControlOutcomeFormat1 *e2sm_rc_pre_ies.E2SmRcPreControlOutcomeFormat1) ([]byte, error) {
	e2SmRcPreControlOutcomeFormat1CP, err := newE2SmRcPreControlOutcomeFormat1(e2SmRcPreControlOutcomeFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlOutcomeFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_ControlOutcome_Format1, unsafe.Pointer(e2SmRcPreControlOutcomeFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlOutcomeFormat1() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeE2SmRcPreControlOutcomeFormat1(e2SmRcPreControlOutcomeFormat1 *e2sm_rc_pre_ies.E2SmRcPreControlOutcomeFormat1) ([]byte, error) {
	e2SmRcPreControlOutcomeFormat1CP, err := newE2SmRcPreControlOutcomeFormat1(e2SmRcPreControlOutcomeFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreControlOutcomeFormat1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_ControlOutcome_Format1, unsafe.Pointer(e2SmRcPreControlOutcomeFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreControlOutcomeFormat1() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreControlOutcomeFormat1(bytes []byte) (*e2sm_rc_pre_ies.E2SmRcPreControlOutcomeFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_ControlOutcome_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2SmRcPreControlOutcomeFormat1((*C.E2SM_RC_PRE_ControlOutcome_Format1_t)(unsafePtr))
}

func PerDecodeE2SmRcPreControlOutcomeFormat1(bytes []byte) (*e2sm_rc_pre_ies.E2SmRcPreControlOutcomeFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_ControlOutcome_Format1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreControlOutcomeFormat1((*C.E2SM_RC_PRE_ControlOutcome_Format1_t)(unsafePtr))
}

func newE2SmRcPreControlOutcomeFormat1(e2SmRcPreControlOutcomeFormat1 *e2sm_rc_pre_ies.E2SmRcPreControlOutcomeFormat1) (*C.E2SM_RC_PRE_ControlOutcome_Format1_t, error) {

	outcomeElementListC := new(C.struct_E2SM_RC_PRE_ControlOutcome_Format1__outcomeElement_List)
	e2SmRcPreControlOutcomeFormat1C := C.E2SM_RC_PRE_ControlOutcome_Format1_t{
		outcomeElement_List: outcomeElementListC,
	}
	for _, outcomeElementListItem := range e2SmRcPreControlOutcomeFormat1.GetOutcomeElementList() {
		outcomeElementListItemC, err := newRanparameterItem(outcomeElementListItem)
		if err != nil {
			return nil, fmt.Errorf("newE2SmRcPreControlOutcomeFormat1() %s", err.Error())
		}
		//ToDo - declare pmContainersListC on top of it...
		if _, err = C.asn_sequence_add(unsafe.Pointer(outcomeElementListC), unsafe.Pointer(outcomeElementListItemC)); err != nil {
			return nil, err
		}
	}

	return &e2SmRcPreControlOutcomeFormat1C, nil
}

func decodeE2SmRcPreControlOutcomeFormat1(e2SmRcPreControlOutcomeFormat1C *C.E2SM_RC_PRE_ControlOutcome_Format1_t) (*e2sm_rc_pre_ies.E2SmRcPreControlOutcomeFormat1, error) {

	var ieCount = int(e2SmRcPreControlOutcomeFormat1C.outcomeElement_List.list.count)
	e2SmRcPreControlOutcomeFormat1 := new(e2sm_rc_pre_ies.E2SmRcPreControlOutcomeFormat1)

	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2SmRcPreControlOutcomeFormat1C.outcomeElement_List.list.array)) * uintptr(i)
		ieC := *(**C.RANparameter_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2SmRcPreControlOutcomeFormat1C.outcomeElement_List.list.array)) + offset))
		ie, err := decodeRanparameterItem(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreControlOutcomeFormat1() %s", err.Error())
		}
		e2SmRcPreControlOutcomeFormat1.OutcomeElementList = append(e2SmRcPreControlOutcomeFormat1.OutcomeElementList, ie)
	}

	return e2SmRcPreControlOutcomeFormat1, nil
}

func decodeE2SmRcPreControlOutcomeFormat1Bytes(array [8]byte) (*e2sm_rc_pre_ies.E2SmRcPreControlOutcomeFormat1, error) {
	e2SmRcPreControlOutcomeFormat1C := (*C.E2SM_RC_PRE_ControlOutcome_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	result, err := decodeE2SmRcPreControlOutcomeFormat1(e2SmRcPreControlOutcomeFormat1C)
	if err != nil {
		return nil, fmt.Errorf("decodeE2SmRcPreControlOutcomeFormat1Bytes() %s", err.Error())
	}

	return result, nil
}
