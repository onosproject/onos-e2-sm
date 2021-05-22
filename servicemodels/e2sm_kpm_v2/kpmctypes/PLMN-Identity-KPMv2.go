// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PLMN-Identity-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodePlmnIdentity(plmnIdentity *e2sm_kpm_v2.PlmnIdentity) ([]byte, error) {
	plmnIdentityCP, err := newPlmnIdentity(plmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("xerEncodePlmnIdentity() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_PLMN_Identity_KPMv2, unsafe.Pointer(plmnIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodePlmnIdentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodePlmnIdentity(plmnIdentity *e2sm_kpm_v2.PlmnIdentity) ([]byte, error) {
	plmnIdentityCP, err := newPlmnIdentity(plmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("perEncodePlmnIdentity() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_PLMN_Identity_KPMv2, unsafe.Pointer(plmnIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodePlmnIdentity() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodePlmnIdentity(bytes []byte) (*e2sm_kpm_v2.PlmnIdentity, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_PLMN_Identity_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodePlmnIdentity((*C.PLMN_Identity_KPMv2_t)(unsafePtr))
}

func perDecodePlmnIdentity(bytes []byte) (*e2sm_kpm_v2.PlmnIdentity, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_PLMN_Identity_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodePlmnIdentity((*C.PLMN_Identity_KPMv2_t)(unsafePtr))
}

func newPlmnIdentity(plmnIdentity *e2sm_kpm_v2.PlmnIdentity) (*C.PLMN_Identity_KPMv2_t, error) {

	plmnIdentityC, err := newOctetString(string(plmnIdentity.Value))
	if err != nil {
		return nil, err
	}

	return plmnIdentityC, nil
}

func decodePlmnIdentity(plmnIdentityC *C.PLMN_Identity_KPMv2_t) (*e2sm_kpm_v2.PlmnIdentity, error) {

	plmnID, err := decodeOctetString(plmnIdentityC)
	if err != nil {
		return nil, err
	}
	plmnIdentity := e2sm_kpm_v2.PlmnIdentity{
		Value: []byte(plmnID),
	}

	return &plmnIdentity, nil
}
