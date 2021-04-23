// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "EUTRACGI.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeEUTRACGI(eutracgi *e2sm_rc_pre_v2.Eutracgi) ([]byte, error) {

	eutracgiCP, err := newEUTRACGI(eutracgi)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEUTRACGI() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_EUTRACGI, unsafe.Pointer(eutracgiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEUTRACGI() %s", err.Error())
	}
	return bytes, nil
}

func newEUTRACGI(eutracgi *e2sm_rc_pre_v2.Eutracgi) (*C.EUTRACGI_t, error) {

	plmnID := newPlmnIdentity(eutracgi.PLmnIdentity)
	eutracellIdentity := newEUTRACellIdentity(eutracgi.EUtracellIdentity)

	eutracgiC := C.EUTRACGI_t{
		pLMN_Identity:     *plmnID,
		eUTRACellIdentity: *eutracellIdentity,
	}

	return &eutracgiC, nil
}

func decodeEUTRACGI(eutracgiC *C.EUTRACGI_t) (*e2sm_rc_pre_v2.Eutracgi, error) {
	eutracgi := new(e2sm_rc_pre_v2.Eutracgi)

	plmnID := decodePlmnIdentity(&eutracgiC.pLMN_Identity)
	eutracgi.PLmnIdentity = plmnID

	eutracellIdentity, err := decodeEUTRACellIdentity(&eutracgiC.eUTRACellIdentity)
	if err != nil {
		return nil, fmt.Errorf("decodeEUTRACGI() %s", err.Error())
	}
	eutracgi.EUtracellIdentity = eutracellIdentity

	return eutracgi, nil
}

func decodeEUTRACGIBytes(array [8]byte) *e2sm_rc_pre_v2.Eutracgi {
	eutracgiC := (*C.EUTRACGI_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	eutracgi, _ := decodeEUTRACGI(eutracgiC)

	return eutracgi
}
