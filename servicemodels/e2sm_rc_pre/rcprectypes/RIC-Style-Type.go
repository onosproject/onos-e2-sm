// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Style-Type.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeRicStyleType(ricStyleType *e2sm_rc_pre_v2.RicStyleType) ([]byte, error) {
	ricStyleTypeCP := newRicStyleType(ricStyleType)

	bytes, err := encodeXer(&C.asn_DEF_RIC_Style_Type, unsafe.Pointer(ricStyleTypeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicStyleType() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicStyleType(ricStyleType *e2sm_rc_pre_v2.RicStyleType) ([]byte, error) {
	ricStyleTypeCP := newRicStyleType(ricStyleType)

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Style_Type, unsafe.Pointer(ricStyleTypeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicStyleType() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicStyleType(bytes []byte) (*e2sm_rc_pre_v2.RicStyleType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Style_Type)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicStyleType((*C.RIC_Style_Type_t)(unsafePtr)), nil
}

func perDecodeRicStyleType(bytes []byte) (*e2sm_rc_pre_v2.RicStyleType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Style_Type)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicStyleType((*C.RIC_Style_Type_t)(unsafePtr)), nil
}

func newRicStyleType(ricStyleType *e2sm_rc_pre_v2.RicStyleType) *C.RIC_Style_Type_t {

	ricStyleTypeC := C.long(ricStyleType.Value)

	return &ricStyleTypeC
}

func decodeRicStyleType(ricStyleTypeC *C.RIC_Style_Type_t) *e2sm_rc_pre_v2.RicStyleType {

	ricStyleType := e2sm_rc_pre_v2.RicStyleType{
		Value: int32(*ricStyleTypeC),
	}
	return &ricStyleType
}

func decodeRicStyleTypeBytes(array [8]byte) *e2sm_rc_pre_v2.RicStyleType {
	ricStyleTypeC := (*C.RIC_Style_Type_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	ricStyleType := decodeRicStyleType(ricStyleTypeC)

	return ricStyleType
}
