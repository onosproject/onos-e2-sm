// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ENB-ID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeEnbID(enbID *e2sm_kpm_v2.EnbId) ([]byte, error) {
	enbIDCP, err := newEnbID(enbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ENB_ID, unsafe.Pointer(enbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEnbID(enbID *e2sm_kpm_v2.EnbId) ([]byte, error) {
	enbIDCP, err := newEnbID(enbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeEnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ENB_ID, unsafe.Pointer(enbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEnbID(bytes []byte) (*e2sm_kpm_v2.EnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ENB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEnbID((*C.ENB_ID_t)(unsafePtr))
}

func perDecodeEnbID(bytes []byte) (*e2sm_kpm_v2.EnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ENB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEnbID((*C.ENB_ID_t)(unsafePtr))
}

func newEnbID(enbID *e2sm_kpm_v2.EnbId) (*C.ENB_ID_t, error) {

	var pr C.ENB_ID_PR
	choiceC := [48]byte{}

	switch choice := enbID.EnbId.(type) {
	case *e2sm_kpm_v2.EnbId_MacroENbId:
		pr = C.ENB_ID_PR_macro_eNB_ID

		im, err := newBitString(choice.MacroENbId)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:8], uint64(uintptr(unsafe.Pointer(im))))
	case *e2sm_kpm_v2.EnbId_HomeENbId:
		pr = C.ENB_ID_PR_home_eNB_ID

		im, err := newBitString(choice.HomeENbId)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[8:16], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newEnbId() %T not yet implemented", choice)
	}

	enbIDC := C.ENB_ID_t{
		present: pr,
		choice:  choiceC,
	}

	return &enbIDC, nil
}

func decodeEnbID(enbIDC *C.ENB_ID_t) (*e2sm_kpm_v2.EnbId, error) {

	//This is Decoder part (OneOf)
	enbID := new(e2sm_kpm_v2.EnbId)

	switch enbIDC.present {
	case C.ENB_ID_PR_macro_eNB_ID:
		var a [8]byte
		copy(a[:], enbIDC.choice[0:8])
		enbIDstructC, err := decodeBitStringBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		enbID.EnbId = &e2sm_kpm_v2.EnbId_MacroENbId{
			MacroENbId: enbIDstructC,
		}
	case C.ENB_ID_PR_home_eNB_ID:
		var a [8]byte
		copy(a[:], enbIDC.choice[8:16])
		enbIDstructC, err := decodeBitStringBytes(a)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		enbID.EnbId = &e2sm_kpm_v2.EnbId_HomeENbId{
			HomeENbId: enbIDstructC,
		}
	default:
		return nil, fmt.Errorf("decodeEnbID() %v not yet implemented", enbIDC.present)
	}

	return enbID, nil
}

func decodeEnbIDBytes(array [48]byte) (*e2sm_kpm_v2.EnbId, error) {
	enbIDC := (*C.ENB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeEnbID(enbIDC)
}
