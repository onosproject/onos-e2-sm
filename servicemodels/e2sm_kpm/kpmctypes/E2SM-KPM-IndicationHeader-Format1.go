// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes
//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-IndicationHeader-Format1.h"
import "C"
import (
	"encoding/binary"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeE2SmKpmIndicationHeaderFormat1(indicationHeaderFormat1 *e2sm_kpm_ies.E2SmKpmIndicationHeaderFormat1) ([]byte, error) {
	indicationHeaderFormat1CP, err := newE2SmKpmIndicationHeaderFormat1(indicationHeaderFormat1)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_IndicationHeader_Format1, unsafe.Pointer(indicationHeaderFormat1CP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newE2SmKpmIndicationHeaderFormat1(indicationHeaderFormat1 *e2sm_kpm_ies.E2SmKpmIndicationHeaderFormat1) (*C.E2SM_KPM_IndicationHeader_Format1_t, error) {

	globalKpmNodeId, err := newGlobalKpmNodeId(indicationHeaderFormat1.IdGlobalKpmnodeId)
	if err != nil {
		return nil, err
	}
	nRcgi, err := newNRCGI(indicationHeaderFormat1.NRcgi)
	if err != nil {
		return nil, err
	}
	plmnId := newPlmnIdentity(indicationHeaderFormat1.PLmnIdentity)
	sliceId := newSnssai(indicationHeaderFormat1.SliceId)
	fiveQi := indicationHeaderFormat1.FiveQi
	qCi := indicationHeaderFormat1.Qci

	fiveQiC := C.long(fiveQi)
	qCiC := C.long(qCi)

	indicationHeaderFormat1C := C.E2SM_KPM_IndicationHeader_Format1_t{
		id_GlobalKPMnode_ID: globalKpmNodeId,
		nRCGI: nRcgi,
		pLMN_Identity: plmnId,
		sliceID: sliceId,
		fiveQI: &fiveQiC,
		qci: &qCiC,
	}

	return &indicationHeaderFormat1C, nil
}

func decodeE2SmKpmIndicationHeaderFormat1(indicationHeaderFormat1C *C.E2SM_KPM_IndicationHeader_Format1_t) (*e2sm_kpm_ies.E2SmKpmIndicationHeaderFormat1, error) {
	indicationHeaderFormat1 := new(e2sm_kpm_ies.E2SmKpmIndicationHeaderFormat1)

	globalKpmNodeId, err := decodeGlobalKpmNodeId(indicationHeaderFormat1C.id_GlobalKPMnode_ID)
	if err != nil {
		return nil, err
	}
	indicationHeaderFormat1.IdGlobalKpmnodeId = globalKpmNodeId

	nRcgi, err := decodeNRCGI(indicationHeaderFormat1C.nRCGI)
	if err != nil {
		return nil, err
	}
	indicationHeaderFormat1.NRcgi = nRcgi

	plmnId := decodePlmnIdentity(indicationHeaderFormat1C.pLMN_Identity)
	indicationHeaderFormat1.PLmnIdentity = plmnId

	sliceId := decodeSnssai(indicationHeaderFormat1C.sliceID)
	indicationHeaderFormat1.SliceId = sliceId

	indicationHeaderFormat1.FiveQi = int32(*indicationHeaderFormat1C.fiveQI)
	indicationHeaderFormat1.Qci = int32(*indicationHeaderFormat1C.qci)

	return indicationHeaderFormat1, nil
}

func decodeE2SmKpmIndicationHeaderFormat1Bytes(array [8]byte) (*e2sm_kpm_ies.E2SmKpmIndicationHeaderFormat1, error) {
	indicationHeaderFormat1C := (*C.E2SM_KPM_IndicationHeader_Format1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2SmKpmIndicationHeaderFormat1(indicationHeaderFormat1C)
}

