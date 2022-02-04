// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRCGI-RCPRE.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeNRCGI(nrcgi *e2sm_rc_pre_v2.Nrcgi) ([]byte, error) {

	nrcgiCP, err := newNRCGI(nrcgi)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNRCGI() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_NRCGI_RCPRE, unsafe.Pointer(nrcgiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNRCGI() %s", err.Error())
	}
	return bytes, nil
}

func newNRCGI(nrcgi *e2sm_rc_pre_v2.Nrcgi) (*C.NRCGI_RCPRE_t, error) {

	plmnID := newPlmnIdentity(nrcgi.PLmnIdentity)
	nrCellIdentity, err := newNRCellIdentity(nrcgi.GetNRcellIdentity())
	if err != nil {
		return nil, err
	}

	nrcgiC := C.NRCGI_RCPRE_t{
		pLMN_Identity:  *plmnID,
		nRCellIdentity: *nrCellIdentity,
	}

	return &nrcgiC, nil
}

func decodeNRCGI(nrcgiC *C.NRCGI_RCPRE_t) (*e2sm_rc_pre_v2.Nrcgi, error) {
	nrcgi := new(e2sm_rc_pre_v2.Nrcgi)

	plmnID := decodePlmnIdentity(&nrcgiC.pLMN_Identity)
	nrcgi.PLmnIdentity = plmnID

	nrCellIdentity, err := decodeNRCellIdentity(&nrcgiC.nRCellIdentity)
	if err != nil {
		return nil, fmt.Errorf("decodeNRCGI() %s", err.Error())
	}
	nrcgi.NRcellIdentity = nrCellIdentity

	return nrcgi, nil
}

func decodeNRCGIBytes(array [8]byte) *e2sm_rc_pre_v2.Nrcgi {
	nrcgiC := (*C.NRCGI_RCPRE_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))
	nrcgi, _ := decodeNRCGI(nrcgiC)
	return nrcgi
}
