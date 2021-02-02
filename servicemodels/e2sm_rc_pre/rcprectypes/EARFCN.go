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
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func xerEncodeEARFCN(earfcn *e2sm_rc_pre_ies.Earfcn) ([]byte, error) {
	earfcnCP := newEARFCN(earfcn)

	bytes, err := encodeXer(&C.asn_DEF_EARFCN, unsafe.Pointer(earfcnCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEARFCN() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEARFCN(earfcn *e2sm_rc_pre_ies.Earfcn) ([]byte, error) {
	earfcnCP := newEARFCN(earfcn)

	bytes, err := encodePerBuffer(&C.asn_DEF_EARFCN, unsafe.Pointer(earfcnCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEARFCN() %s", err.Error())
	}
	return bytes, nil
}

func perDecodeEARFCN(bytes []byte) (*e2sm_rc_pre_ies.Earfcn, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_EARFCN)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEARFCN((*C.EARFCN_t)(unsafePtr)), nil
}

func newEARFCN(earfcn *e2sm_rc_pre_ies.Earfcn) *C.EARFCN_t {
	earfcnC := C.long(earfcn.Value)
	return &earfcnC
}

func decodeEARFCN(earfcnC *C.EARFCN_t) *e2sm_rc_pre_ies.Earfcn {
	return &e2sm_rc_pre_ies.Earfcn{
		Value: int32(*earfcnC),
	}
}
