// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PLMN-Identity.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodePlmnIdentity(plmnIdentity *e2sm_mho.PlmnIdentity) ([]byte, error) {
	plmnIdentityCP, err := newPlmnIdentity(plmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("xerEncodePlmnIdentity() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_PLMN_Identity, unsafe.Pointer(plmnIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodePlmnIdentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodePlmnIdentity(plmnIdentity *e2sm_mho.PlmnIdentity) ([]byte, error) {
	plmnIdentityCP, err := newPlmnIdentity(plmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("perEncodePlmnIdentity() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_PLMN_Identity, unsafe.Pointer(plmnIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodePlmnIdentity() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodePlmnIdentity(bytes []byte) (*e2sm_mho.PlmnIdentity, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_PLMN_Identity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodePlmnIdentity((*C.PLMN_Identity_t)(unsafePtr))
}

func perDecodePlmnIdentity(bytes []byte) (*e2sm_mho.PlmnIdentity, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_PLMN_Identity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodePlmnIdentity((*C.PLMN_Identity_t)(unsafePtr))
}

func newPlmnIdentity(plmnIdentity *e2sm_mho.PlmnIdentity) (*C.PLMN_Identity_t, error) {

	plmnIdentityC := decodeOctetString(plmnIdentity.Value)

	return &plmnIdentityC, nil
}

func decodePlmnIdentity(plmnIdentityC *C.PLMN_Identity_t) (*e2sm_mho.PlmnIdentity, error) {

	plmnIdentity := e2sm_mho.PlmnIdentity{
		Value: newOctetString(plmnIdentityC),
	}

	return &plmnIdentity, nil
}

func decodePlmnIdentityBytes(array [8]byte) (*e2sm_mho.PlmnIdentity, error) {
	plmnIdentityC := (*C.PLMN_Identity_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodePlmnIdentity(plmnIdentityC)
}
