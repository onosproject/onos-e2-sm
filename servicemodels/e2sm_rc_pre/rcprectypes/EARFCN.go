// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "EARFCN.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeEARFCN(earfcn *e2sm_rc_pre_v2.Earfcn) ([]byte, error) {
	earfcnCP := newEARFCN(earfcn)

	bytes, err := encodeXer(&C.asn_DEF_EARFCN, unsafe.Pointer(earfcnCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEARFCN() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEARFCN(bytes []byte) (*e2sm_rc_pre_v2.Earfcn, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_EARFCN)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEARFCN((*C.EARFCN_t)(unsafePtr)), nil
}

func perEncodeEARFCN(earfcn *e2sm_rc_pre_v2.Earfcn) ([]byte, error) {
	earfcnCP := newEARFCN(earfcn)

	bytes, err := encodePerBuffer(&C.asn_DEF_EARFCN, unsafe.Pointer(earfcnCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEARFCN() %s", err.Error())
	}
	return bytes, nil
}

func perDecodeEARFCN(bytes []byte) (*e2sm_rc_pre_v2.Earfcn, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_EARFCN)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEARFCN((*C.EARFCN_t)(unsafePtr)), nil
}

func newEARFCN(earfcn *e2sm_rc_pre_v2.Earfcn) *C.EARFCN_t {
	earfcnC := C.long(earfcn.Value)
	return &earfcnC
}

func decodeEARFCN(earfcnC *C.EARFCN_t) *e2sm_rc_pre_v2.Earfcn {
	return &e2sm_rc_pre_v2.Earfcn{
		Value: int32(*earfcnC),
	}
}

func decodeEARFCNBytes(array [8]byte) *e2sm_rc_pre_v2.Earfcn {
	earfcnC := (C.EARFCN_t)(binary.LittleEndian.Uint32(array[0:4]))
	earfcn := decodeEARFCN(&earfcnC)

	return earfcn
}
