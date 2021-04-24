// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRCGI.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeNrcgi(nrcgi *e2sm_mho.Nrcgi) ([]byte, error) {
	nrcgiCP, err := newNrcgi(nrcgi)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNrcgi() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_NRCGI, unsafe.Pointer(nrcgiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNrcgi() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeNrcgi(nrcgi *e2sm_mho.Nrcgi) ([]byte, error) {
	nrcgiCP, err := newNrcgi(nrcgi)
	if err != nil {
		return nil, fmt.Errorf("perEncodeNrcgi() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_NRCGI, unsafe.Pointer(nrcgiCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeNrcgi() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeNrcgi(bytes []byte) (*e2sm_mho.Nrcgi, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_NRCGI)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeNrcgi((*C.NRCGI_t)(unsafePtr))
}

func perDecodeNrcgi(bytes []byte) (*e2sm_mho.Nrcgi, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_NRCGI)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeNrcgi((*C.NRCGI_t)(unsafePtr))
}

func newNrcgi(nrcgi *e2sm_mho.Nrcgi) (*C.NRCGI_t, error) {

	var err error
	nrcgiC := C.NRCGI_t{}

	pLmnIDentityC, err := newPlmnIdentity(nrcgi.PLmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	nRcellIDentityC, err := newNrcellIDentity(nrcgi.NRcellIdentity)
	if err != nil {
		return nil, fmt.Errorf("newNrcellIDentity() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	nrcgiC.pLMN_Identity = pLmnIdentityC
	nrcgiC.nRCellIdentity = nRcellIdentityC

	return &nrcgiC, nil
}

func decodeNrcgi(nrcgiC *C.NRCGI_t) (*e2sm_mho.Nrcgi, error) {

	var err error
	nrcgi := e2sm_mho.Nrcgi{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//PLmnIdentity: pLmnIdentity,
		//NRcellIdentity: nRcellIdentity,

	}

	nrcgi.PLmnIdentity, err = decodePlmnIdentity(nrcgiC.pLMN_Identity)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	nrcgi.NRcellIdentity, err = decodeNrcellIDentity(nrcgiC.nRCellIdentity)
	if err != nil {
		return nil, fmt.Errorf("decodeNrcellIDentity() %s", err.Error())
	}

	return &nrcgi, nil
}

func decodeNrcgiBytes(array [8]byte) (*e2sm_mho.Nrcgi, error) {
	nrcgiC := (*C.NRCGI_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeNrcgi(nrcgiC)
}
