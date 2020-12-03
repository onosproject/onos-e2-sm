// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-Style-Name.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeRicStyleName(ricStyleName *e2sm_kpm_ies.RicStyleName) ([]byte, error) {
	gnbDuIDCP := newRicStyleName(ricStyleName)

	bytes, err := encodeXer(&C.asn_DEF_RIC_Style_Name, unsafe.Pointer(gnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicStyleName() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicStyleName(ricStyleName *e2sm_kpm_ies.RicStyleName) ([]byte, error) {
	gnbDuIDCP := newRicStyleName(ricStyleName)

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_Style_Name, unsafe.Pointer(gnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicStyleName() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicStyleName(bytes []byte) (*e2sm_kpm_ies.RicStyleName, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_Style_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicStyleName((*C.RIC_Style_Name_t)(unsafePtr)), nil
}

func perDecodeRicStyleName(bytes []byte) (*e2sm_kpm_ies.RicStyleName, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_Style_Name)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicStyleName((*C.RIC_Style_Name_t)(unsafePtr)), nil
}

func newRicStyleName(ricStyleName *e2sm_kpm_ies.RicStyleName) *C.RIC_Style_Name_t {
	ricStyleNameCP := newPrintableString(ricStyleName.Value)

	return ricStyleNameCP
}

func decodeRicStyleName(ricStyleNameC *C.RIC_Style_Name_t) *e2sm_kpm_ies.RicStyleName {

	ricStyleName := e2sm_kpm_ies.RicStyleName{
		Value: decodePrintableString(ricStyleNameC),
	}
	return &ricStyleName
}
