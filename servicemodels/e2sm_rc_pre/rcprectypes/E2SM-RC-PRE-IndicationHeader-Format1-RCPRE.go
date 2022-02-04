// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-RC-PRE-IndicationHeader-Format1-RCPRE.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func XerEncodeE2SmRcPreIndicationHeaderFormat1(indicationHeaderFormat1 *e2sm_rc_pre_v2.E2SmRcPreIndicationHeaderFormat1) ([]byte, error) {
	indicationHeaderFormat1CP, err := newE2SmRcPreIndicationHeaderFormat1(indicationHeaderFormat1)
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationHeaderFormat1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_RC_PRE_IndicationHeader_Format1_RCPRE, unsafe.Pointer(indicationHeaderFormat1CP))
	if err != nil {
		return nil, fmt.Errorf("XerEncodeE2SmRcPreIndicationHeaderFormat1() %s", err.Error())
	}
	return bytes, nil
}

func XerDecodeE2SmRcPreIndicationHeaderFormat1(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreIndicationHeaderFormat1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2SM_RC_PRE_IndicationHeader_Format1_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreIndicationHeaderFormat1((*C.E2SM_RC_PRE_IndicationHeader_Format1_RCPRE_t)(unsafePtr))
}

func PerEncodeE2SmRcPreIndicationHeaderFormat1(indicationHeader *e2sm_rc_pre_v2.E2SmRcPreIndicationHeaderFormat1) ([]byte, error) {
	indicationHeaderCP, err := newE2SmRcPreIndicationHeaderFormat1(indicationHeader)
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreIndicationHeader() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2SM_RC_PRE_IndicationHeader_Format1_RCPRE, unsafe.Pointer(indicationHeaderCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeE2SmRcPreIndicationHeader() %s", err.Error())
	}
	return bytes, nil
}

func PerDecodeE2SmRcPreIndicationHeaderFormat1(bytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreIndicationHeaderFormat1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2SM_RC_PRE_IndicationHeader_Format1_RCPRE)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2SmRcPreIndicationHeaderFormat1((*C.E2SM_RC_PRE_IndicationHeader_Format1_RCPRE_t)(unsafePtr))
}

func newE2SmRcPreIndicationHeaderFormat1(indicationHeaderFormat1 *e2sm_rc_pre_v2.E2SmRcPreIndicationHeaderFormat1) (*C.E2SM_RC_PRE_IndicationHeader_Format1_RCPRE_t, error) {

	cgi, _ := newCellGlobalID(indicationHeaderFormat1.Cgi)

	indicationHeaderFormat1C := C.E2SM_RC_PRE_IndicationHeader_Format1_RCPRE_t{
		cgi: cgi,
	}

	return &indicationHeaderFormat1C, nil
}

func decodeE2SmRcPreIndicationHeaderFormat1(indicationHeaderFormat1C *C.E2SM_RC_PRE_IndicationHeader_Format1_RCPRE_t) (*e2sm_rc_pre_v2.E2SmRcPreIndicationHeaderFormat1, error) {
	indicationHeaderFormat1 := new(e2sm_rc_pre_v2.E2SmRcPreIndicationHeaderFormat1)

	if indicationHeaderFormat1C.cgi != nil { // Is optional
		cgi, err := decodeCellGlobalID(indicationHeaderFormat1C.cgi)
		if err != nil {
			return nil, fmt.Errorf("decodeE2SmRcPreIndicationHeaderFormat1() %s", err.Error())
		}
		indicationHeaderFormat1.Cgi = cgi
	}
	return indicationHeaderFormat1, nil
}

func decodeE2SmRcPreIndicationHeaderFormat1Bytes(array [8]byte) (*e2sm_rc_pre_v2.E2SmRcPreIndicationHeaderFormat1, error) {
	indicationHeaderFormat1C := (*C.E2SM_RC_PRE_IndicationHeader_Format1_RCPRE_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmRcPreIndicationHeaderFormat1(indicationHeaderFormat1C)
}
