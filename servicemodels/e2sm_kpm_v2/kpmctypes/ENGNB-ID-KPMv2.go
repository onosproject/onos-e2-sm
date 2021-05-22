// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ENGNB-ID-KPMv2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeEngnbID(engnbID *e2sm_kpm_v2.EngnbId) ([]byte, error) {
	engnbIDCP, err := newEngnbID(engnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEngnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ENGNB_ID_KPMv2, unsafe.Pointer(engnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEngnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEngnbID(engnbID *e2sm_kpm_v2.EngnbId) ([]byte, error) {
	engnbIDCP, err := newEngnbID(engnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeEngnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ENGNB_ID_KPMv2, unsafe.Pointer(engnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEngnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEngnbID(bytes []byte) (*e2sm_kpm_v2.EngnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ENGNB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEngnbID((*C.ENGNB_ID_KPMv2_t)(unsafePtr))
}

func perDecodeEngnbID(bytes []byte) (*e2sm_kpm_v2.EngnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ENGNB_ID_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEngnbID((*C.ENGNB_ID_KPMv2_t)(unsafePtr))
}

func newEngnbID(engnbID *e2sm_kpm_v2.EngnbId) (*C.ENGNB_ID_KPMv2_t, error) {

	var pr C.ENGNB_ID_KPMv2_PR
	choiceC := [48]byte{}

	switch choice := engnbID.EngnbId.(type) {
	case *e2sm_kpm_v2.EngnbId_GNbId:
		pr = C.ENGNB_ID_KPMv2_PR_gNB_ID

		bsC, err := newBitString(choice.GNbId)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	default:
		return nil, fmt.Errorf("newEngnbID() %T not yet implemented", choice)
	}

	engnbIDC := C.ENGNB_ID_KPMv2_t{
		present: pr,
		choice:  choiceC,
	}

	return &engnbIDC, nil
}

func decodeEngnbID(engnbIDC *C.ENGNB_ID_KPMv2_t) (*e2sm_kpm_v2.EngnbId, error) {

	//This is Decoder part (OneOf)
	engnbID := new(e2sm_kpm_v2.EngnbId)

	switch engnbIDC.present {
	case C.ENGNB_ID_KPMv2_PR_gNB_ID:
		engnbIDstructC := newBitStringFromArray(engnbIDC.choice)

		engnbIDCdec, err := decodeBitString(engnbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		engnbID.EngnbId = &e2sm_kpm_v2.EngnbId_GNbId{
			GNbId: engnbIDCdec,
		}
	case C.ENGNB_ID_KPMv2_PR_NOTHING:
		return nil, fmt.Errorf("decodeEngnbID() An empty en-Gnb-ID has been sent %v", engnbIDC.present)
	default:
		return nil, fmt.Errorf("decodeEngnbID() %v not yet implemented", engnbIDC.present)
	}

	return engnbID, nil
}
