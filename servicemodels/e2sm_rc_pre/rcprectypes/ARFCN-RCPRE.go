// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ARFCN-RCPRE.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeARFCN(arfcn *e2sm_rc_pre_v2.Arfcn) ([]byte, error) {

	arfcnCP, err := newARFCN(arfcn)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeArfcn() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ARFCN_RCPRE, unsafe.Pointer(arfcnCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeArfcn() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeARFCN(arfcn *e2sm_rc_pre_v2.Arfcn) ([]byte, error) {
	arfcnCP, err := newARFCN(arfcn)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ARFCN_RCPRE, unsafe.Pointer(arfcnCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeArfcn() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeARFCN(bytes []byte) (*e2sm_rc_pre_v2.Arfcn, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ARFCN_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeARFCN((*C.ARFCN_RCPRE_t)(unsafePtr))
}

func perDecodeARFCN(bytes []byte) (*e2sm_rc_pre_v2.Arfcn, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ARFCN_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeARFCN((*C.ARFCN_RCPRE_t)(unsafePtr))
}

func newARFCN(arfcn *e2sm_rc_pre_v2.Arfcn) (*C.ARFCN_RCPRE_t, error) {
	var pr C.ARFCN_RCPRE_PR
	choiceC := [8]byte{}

	switch choice := arfcn.Arfcn.(type) {
	case *e2sm_rc_pre_v2.Arfcn_EArfcn:
		pr = C.ARFCN_RCPRE_PR_eARFCN

		im := newEARFCN(choice.EArfcn)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(*im))
	case *e2sm_rc_pre_v2.Arfcn_NrArfcn:
		pr = C.ARFCN_RCPRE_PR_nrARFCN

		im := newNRARFCN(choice.NrArfcn)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(*im))
	default:
		return nil, fmt.Errorf("newARFCN() %T not yet implemented", choice)
	}

	arfcnC := C.ARFCN_RCPRE_t{
		present: pr,
		choice:  choiceC,
	}

	return &arfcnC, nil
}

func decodeARFCN(arfcnC *C.ARFCN_RCPRE_t) (*e2sm_rc_pre_v2.Arfcn, error) {
	arfcn := new(e2sm_rc_pre_v2.Arfcn)

	switch arfcnC.present {
	case C.ARFCN_RCPRE_PR_eARFCN:
		earfcn := decodeEARFCNBytes(arfcnC.choice)
		arfcn.Arfcn = &e2sm_rc_pre_v2.Arfcn_EArfcn{
			EArfcn: earfcn,
		}
	case C.ARFCN_RCPRE_PR_nrARFCN:
		nrarfcn := decodeNRARFCNBytes(arfcnC.choice)
		arfcn.Arfcn = &e2sm_rc_pre_v2.Arfcn_NrArfcn{
			NrArfcn: nrarfcn,
		}
	default:
		return nil, fmt.Errorf("decodeArfcn() %v not yet implemented", arfcnC.present)
	}

	return arfcn, nil
}
