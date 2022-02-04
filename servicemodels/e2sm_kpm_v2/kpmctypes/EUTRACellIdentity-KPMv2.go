// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "EUTRACellIdentity-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeEutracellIdentity(eutracellIdentity *e2sm_kpm_v2.EutracellIdentity) ([]byte, error) {
	eutracellIdentityCP, err := newEutracellIdentity(eutracellIdentity)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEutracellIdentity() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_EUTRACellIdentity_KPMv2, unsafe.Pointer(eutracellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEutracellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEutracellIdentity(eutracellIdentity *e2sm_kpm_v2.EutracellIdentity) ([]byte, error) {
	eutracellIdentityCP, err := newEutracellIdentity(eutracellIdentity)
	if err != nil {
		return nil, fmt.Errorf("perEncodeEutracellIdentity() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_EUTRACellIdentity_KPMv2, unsafe.Pointer(eutracellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEutracellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEutracellIdentity(bytes []byte) (*e2sm_kpm_v2.EutracellIdentity, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_EUTRACellIdentity_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEutracellIdentity((*C.EUTRACellIdentity_KPMv2_t)(unsafePtr))
}

func perDecodeEutracellIdentity(bytes []byte) (*e2sm_kpm_v2.EutracellIdentity, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_EUTRACellIdentity_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEutracellIdentity((*C.EUTRACellIdentity_KPMv2_t)(unsafePtr))
}

func newEutracellIdentity(eutracellIdentity *e2sm_kpm_v2.EutracellIdentity) (*C.EUTRACellIdentity_KPMv2_t, error) {

	eutracellIdentityC, err := newBitString(eutracellIdentity.Value)
	if err != nil {
		return nil, fmt.Errorf("newBitString() %s", err.Error())
	}

	return eutracellIdentityC, nil
}

func decodeEutracellIdentity(eutracellIdentityC *C.EUTRACellIdentity_KPMv2_t) (*e2sm_kpm_v2.EutracellIdentity, error) {

	eutracellIdentity := new(e2sm_kpm_v2.EutracellIdentity)
	eutracellIdentityValue, err := decodeBitString(eutracellIdentityC)
	if err != nil {
		return nil, fmt.Errorf("decodeBitString() %s", err.Error())
	}
	eutracellIdentity.Value = eutracellIdentityValue

	return eutracellIdentity, nil
}
