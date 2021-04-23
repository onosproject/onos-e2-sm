// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Style-Name.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeRicStyleName(ricStyleName *e2sm_rc_pre_v2.RicStyleName) ([]byte, error) {
	ricStyleNameCP := newRicStyleName(ricStyleName)

	bytes, err := encodeXer(&C.asn_DEF_RIC_Style_Name, unsafe.Pointer(ricStyleNameCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicStyleName() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicStyleName(ricStyleName *e2sm_rc_pre_v2.RicStyleName) ([]byte, error) {
	ricStyleNameCP := newRicStyleName(ricStyleName)

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Style_Name, unsafe.Pointer(ricStyleNameCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicStyleName() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicStyleName(bytes []byte) (*e2sm_rc_pre_v2.RicStyleName, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Style_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicStyleName((*C.RIC_Style_Name_t)(unsafePtr)), nil
}

func perDecodeRicStyleName(bytes []byte) (*e2sm_rc_pre_v2.RicStyleName, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Style_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicStyleName((*C.RIC_Style_Name_t)(unsafePtr)), nil
}

func newRicStyleName(ricStyleName *e2sm_rc_pre_v2.RicStyleName) *C.RIC_Style_Name_t {
	ricStyleNameCP := newPrintableString(ricStyleName.Value)

	return ricStyleNameCP
}

func decodeRicStyleName(ricStyleNameC *C.RIC_Style_Name_t) *e2sm_rc_pre_v2.RicStyleName {

	ricStyleName := e2sm_rc_pre_v2.RicStyleName{
		Value: decodePrintableString(ricStyleNameC),
	}
	return &ricStyleName
}
