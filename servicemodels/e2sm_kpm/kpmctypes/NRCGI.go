// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NRCGI.h"
import "C"
import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeNRCGI(nrcgi *e2sm_kpm_ies.Nrcgi) ([]byte, error) {

	nrcgiCP, err := newNRCGI(nrcgi)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_NRCGI, unsafe.Pointer(nrcgiCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newNRCGI(nrcgi *e2sm_kpm_ies.Nrcgi) (*C.NRCGI_t, error) {

	plmnID := newPlmnIdentity(nrcgi.PLmnIdentity)
	nrCellIdentity := newNRCellIdentity(nrcgi.NRcellIdentity)

	nrcgiC := C.NRCGI_t{
		pLMN_Identity: *plmnID,
		nRCellIdentity:  *nrCellIdentity,
	}

	return &nrcgiC, nil
}

func decodeNRCGI(nrcgiC *C.NRCGI_t) (*e2sm_kpm_ies.Nrcgi, error) {
	nrcgi := new(e2sm_kpm_ies.Nrcgi)

	plmnID := decodePlmnIdentity(&nrcgiC.pLMN_Identity)
	nrcgi.PLmnIdentity = plmnID

	nrCellIdentity, err := decodeNRCellIdentity(&nrcgiC.nRCellIdentity)
	if err != nil {
		return nil, err
	}
	nrcgi.NRcellIdentity = nrCellIdentity

	return nrcgi, nil
}