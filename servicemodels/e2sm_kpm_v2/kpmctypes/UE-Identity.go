// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "UE-Identity.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeUeIdentity(ueIdentity *e2sm_kpm_v2.UeIdentity) ([]byte, error) {
	ueIdentityCP, err := newUeIdentity(ueIdentity)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeUeIdentity() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_UE_Identity, unsafe.Pointer(ueIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeUeIdentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeUeIdentity(ueIdentity *e2sm_kpm_v2.UeIdentity) ([]byte, error) {
	ueIdentityCP, err := newUeIdentity(ueIdentity)
	if err != nil {
		return nil, fmt.Errorf("perEncodeUeIdentity() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_UE_Identity, unsafe.Pointer(ueIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeUeIdentity() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeUeIdentity(bytes []byte) (*e2sm_kpm_v2.UeIdentity, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_UE_Identity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeUeIdentity((*C.UE_Identity_t)(unsafePtr))
}

func perDecodeUeIdentity(bytes []byte) (*e2sm_kpm_v2.UeIdentity, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_UE_Identity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeUeIdentity((*C.UE_Identity_t)(unsafePtr))
}

func newUeIdentity(ueIdentity *e2sm_kpm_v2.UeIdentity) (*C.UE_Identity_t, error) {

	ueIdentityC, err := newOctetString(ueIdentity.Value)
	if err != nil {
		return nil, fmt.Errorf("decodeOctetString() %s", err.Error())
	}

	return ueIdentityC, nil
}

func decodeUeIdentity(ueIdentityC *C.UE_Identity_t) (*e2sm_kpm_v2.UeIdentity, error) {

	ueIdentity := new(e2sm_kpm_v2.UeIdentity)
	ueIdentityValue, err := decodeOctetString(ueIdentityC)
	if err != nil {
		return nil, fmt.Errorf("decodeOctetString() %s", err.Error())
	}
	ueIdentity.Value = ueIdentityValue

	return ueIdentity, nil
}
