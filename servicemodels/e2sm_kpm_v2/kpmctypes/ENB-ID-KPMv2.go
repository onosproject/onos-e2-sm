// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ENB-ID-KPMv2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeEnbID(enbID *e2sm_kpm_v2.EnbId) ([]byte, error) {
	enbIDCP, err := newEnbID(enbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ENB_ID_KPMv2, unsafe.Pointer(enbIDCP))
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

	bytes, err := encodePerBuffer(&C.asn_DEF_ENB_ID_KPMv2, unsafe.Pointer(enbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEnbID(bytes []byte) (*e2sm_kpm_v2.EnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ENB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEnbID((*C.ENB_ID_KPMv2_t)(unsafePtr))
}

func perDecodeEnbID(bytes []byte) (*e2sm_kpm_v2.EnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ENB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEnbID((*C.ENB_ID_KPMv2_t)(unsafePtr))
}

func newEnbID(enbID *e2sm_kpm_v2.EnbId) (*C.ENB_ID_KPMv2_t, error) {

	var pr C.ENB_ID_KPMv2_PR
	choiceC := [48]byte{}

	switch choice := enbID.EnbId.(type) {
	case *e2sm_kpm_v2.EnbId_MacroENbId:
		pr = C.ENB_ID_KPMv2_PR_macro_eNB_ID

		bsC, err := newBitString(choice.MacroENbId)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	case *e2sm_kpm_v2.EnbId_HomeENbId:
		pr = C.ENB_ID_KPMv2_PR_home_eNB_ID

		bsC, err := newBitString(choice.HomeENbId)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	default:
		return nil, fmt.Errorf("newEnbId() %T not yet implemented", choice)
	}

	enbIDC := C.ENB_ID_KPMv2_t{
		present: pr,
		choice:  choiceC,
	}

	return &enbIDC, nil
}

func decodeEnbID(enbIDC *C.ENB_ID_KPMv2_t) (*e2sm_kpm_v2.EnbId, error) {

	//This is Decoder part (OneOf)
	enbID := new(e2sm_kpm_v2.EnbId)

	switch enbIDC.present {
	case C.ENB_ID_KPMv2_PR_macro_eNB_ID:
		enbIDstructC := newBitStringFromArray(enbIDC.choice)

		enbIDdec, err := decodeBitString(enbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		enbID.EnbId = &e2sm_kpm_v2.EnbId_MacroENbId{
			MacroENbId: enbIDdec,
		}
	case C.ENB_ID_KPMv2_PR_home_eNB_ID:
		enbIDstructC := newBitStringFromArray(enbIDC.choice)

		enbIDdec, err := decodeBitString(enbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		enbID.EnbId = &e2sm_kpm_v2.EnbId_HomeENbId{
			HomeENbId: enbIDdec,
		}
	case C.ENB_ID_KPMv2_PR_NOTHING:
		return nil, fmt.Errorf("decodeEnbID() An empty EnbID has been sent %v", enbIDC.present)
	default:
		return nil, fmt.Errorf("decodeEnbID() %v not yet implemented", enbIDC.present)
	}

	return enbID, nil
}
