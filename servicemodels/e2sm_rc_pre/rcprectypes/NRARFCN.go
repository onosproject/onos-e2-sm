// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRARFCN.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeNRARFCN(nrarfcn *e2sm_rc_pre_v2.Nrarfcn) ([]byte, error) {

	nrarfcnCP := newNRARFCN(nrarfcn)

	bytes, err := encodeXer(&C.asn_DEF_NRARFCN, unsafe.Pointer(nrarfcnCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNRARFCN() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeNRARFCN(nrarfcn *e2sm_rc_pre_v2.Nrarfcn) ([]byte, error) {
	nrarfcnCP := newNRARFCN(nrarfcn)

	bytes, err := encodePerBuffer(&C.asn_DEF_NRARFCN, unsafe.Pointer(nrarfcnCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeNRARFCN() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeNRARFCN(bytes []byte) (*e2sm_rc_pre_v2.Nrarfcn, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_NRARFCN)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeNRARFCN((*C.NRARFCN_t)(unsafePtr)), nil
}

func perDecodeNRARFCN(bytes []byte) (*e2sm_rc_pre_v2.Nrarfcn, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_NRARFCN)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeNRARFCN((*C.NRARFCN_t)(unsafePtr)), nil
}

func newNRARFCN(nrarfcn *e2sm_rc_pre_v2.Nrarfcn) *C.NRARFCN_t {
	nrarfcnC := C.long(nrarfcn.Value)
	return &nrarfcnC
}

func decodeNRARFCN(nrarfcnC *C.NRARFCN_t) *e2sm_rc_pre_v2.Nrarfcn {
	return &e2sm_rc_pre_v2.Nrarfcn{
		Value: int32(*nrarfcnC),
	}
}

func decodeNRARFCNBytes(array [8]byte) *e2sm_rc_pre_v2.Nrarfcn {
	nrarfcnC := (C.NRARFCN_t)(binary.LittleEndian.Uint32(array[0:4]))
	nrarfcn := decodeNRARFCN(&nrarfcnC)

	return nrarfcn
}
