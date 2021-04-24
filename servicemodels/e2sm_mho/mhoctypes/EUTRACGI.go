// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "EUTRACGI.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeEutracgi(eutracgi *e2sm_mho.Eutracgi) ([]byte, error) {
	eutracgiCP, err := newEutracgi(eutracgi)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEutracgi() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_EUTRACGI, unsafe.Pointer(eutracgiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEutracgi() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEutracgi(eutracgi *e2sm_mho.Eutracgi) ([]byte, error) {
	eutracgiCP, err := newEutracgi(eutracgi)
	if err != nil {
		return nil, fmt.Errorf("perEncodeEutracgi() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_EUTRACGI, unsafe.Pointer(eutracgiCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEutracgi() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEutracgi(bytes []byte) (*e2sm_mho.Eutracgi, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_EUTRACGI)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEutracgi((*C.EUTRACGI_t)(unsafePtr))
}

func perDecodeEutracgi(bytes []byte) (*e2sm_mho.Eutracgi, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_EUTRACGI)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEutracgi((*C.EUTRACGI_t)(unsafePtr))
}

func newEutracgi(eutracgi *e2sm_mho.Eutracgi) (*C.EUTRACGI_t, error) {

	var err error
	eutracgiC := C.EUTRACGI_t{}

	pLmnIdentityC, err := newPlmnIdentity(eutracgi.PLmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	eUtracellIdentityC, err := newEutracellIDentity(eutracgi.EUtracellIdentity)
	if err != nil {
		return nil, fmt.Errorf("newEutracellIDentity() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	eutracgiC.pLMN_Identity = *pLmnIdentityC
	eutracgiC.eUTRACellIdentity = *eUtracellIdentityC

	return &eutracgiC, nil
}

func decodeEutracgi(eutracgiC *C.EUTRACGI_t) (*e2sm_mho.Eutracgi, error) {

	var err error
	eutracgi := e2sm_mho.Eutracgi{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//PLmnIdentity: pLmnIdentity,
		//EUtracellIdentity: eUtracellIdentity,

	}

	eutracgi.PLmnIdentity, err = decodePlmnIdentity(&eutracgiC.pLMN_Identity)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	eutracgi.EUtracellIdentity, err = decodeEutracellIDentity(&eutracgiC.eUTRACellIdentity)
	if err != nil {
		return nil, fmt.Errorf("decodeEutracellIDentity() %s", err.Error())
	}

	return &eutracgi, nil
}

func decodeEutracgiBytes(array [8]byte) (*e2sm_mho.Eutracgi, error) {
	eutracgiC := (*C.EUTRACGI_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeEutracgi(eutracgiC)
}
