// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRCGI-KPMv2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeNrcgi(nrcgi *e2sm_kpm_v2.Nrcgi) ([]byte, error) {
	nrcgiCP, err := newNrcgi(nrcgi)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNrcgi() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_NRCGI_KPMv2, unsafe.Pointer(nrcgiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNrcgi() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeNrcgi(nrcgi *e2sm_kpm_v2.Nrcgi) ([]byte, error) {
	nrcgiCP, err := newNrcgi(nrcgi)
	if err != nil {
		return nil, fmt.Errorf("perEncodeNrcgi() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_NRCGI_KPMv2, unsafe.Pointer(nrcgiCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeNrcgi() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeNrcgi(bytes []byte) (*e2sm_kpm_v2.Nrcgi, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_NRCGI_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeNrcgi((*C.NRCGI_KPMv2_t)(unsafePtr))
}

func perDecodeNrcgi(bytes []byte) (*e2sm_kpm_v2.Nrcgi, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_NRCGI_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeNrcgi((*C.NRCGI_KPMv2_t)(unsafePtr))
}

func newNrcgi(nrcgi *e2sm_kpm_v2.Nrcgi) (*C.NRCGI_KPMv2_t, error) {

	pLmnIdentityC, err := newPlmnIdentity(nrcgi.PLmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	nRcellIdentityC, err := newNrcellIdentity(nrcgi.NRcellIdentity)
	if err != nil {
		return nil, fmt.Errorf("newNrcellIdentity() %s", err.Error())
	}

	nrcgiC := C.NRCGI_KPMv2_t{
		pLMN_Identity:  *pLmnIdentityC,
		nRCellIdentity: *nRcellIdentityC,
	}

	return &nrcgiC, nil
}

func decodeNrcgi(nrcgiC *C.NRCGI_KPMv2_t) (*e2sm_kpm_v2.Nrcgi, error) {

	pLmnIdentity, err := decodePlmnIdentity(&nrcgiC.pLMN_Identity)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	nRcellIdentity, err := decodeNrcellIdentity(&nrcgiC.nRCellIdentity)
	if err != nil {
		return nil, fmt.Errorf("decodeNrcellIdentity() %s", err.Error())
	}

	nrcgi := e2sm_kpm_v2.Nrcgi{
		PLmnIdentity:   pLmnIdentity,
		NRcellIdentity: nRcellIdentity,
	}

	return &nrcgi, nil
}

func decodeNrcgiBytes(array [8]byte) (*e2sm_kpm_v2.Nrcgi, error) {
	nrcgiC := (*C.NRCGI_KPMv2_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeNrcgi(nrcgiC)
}
