// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RRCStatus.h"
import "C"
import (
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeRrcstatus(rrcstatus *e2sm_mho.Rrcstatus) ([]byte, error) {
	rrcstatusCP, err := newRrcstatus(rrcstatus)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RRCStatus, unsafe.Pointer(rrcstatusCP)) //ToDo - change name of C-encoder tag
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRrcstatus() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRrcstatus(rrcstatus *e2sm_mho.Rrcstatus) ([]byte, error) {
	rrcstatusCP, err := newRrcstatus(rrcstatus)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RRCStatus, unsafe.Pointer(rrcstatusCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRrcstatus() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRrcstatus(bytes []byte) (*e2sm_mho.Rrcstatus, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RRCStatus)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRrcstatus((*C.RRCStatus_t)(unsafePtr)) //ToDo - change name of C-struct
}

func perDecodeRrcstatus(bytes []byte) (*e2sm_mho.Rrcstatus, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RRCStatus)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRrcstatus((*C.RRCStatus_t)(unsafePtr))
}

func newRrcstatus(rrcstatus *e2sm_mho.Rrcstatus) (*C.RRCStatus_t, error) {
	var ret C.RRCStatus_t
	switch *rrcstatus {
	case e2sm_mho.Rrcstatus_RRCSTATUS_CONNECTED:
		ret = C.RRCStatus_connected //ToDo - double-check correctness of the name
	case e2sm_mho.Rrcstatus_RRCSTATUS_INACTIVE:
		ret = C.RRCStatus_inactive //ToDo - double-check correctness of the name
	case e2sm_mho.Rrcstatus_RRCSTATUS_IDLE:
		ret = C.RRCStatus_idle //ToDo - double-check correctness of the name
	default:
		return nil, fmt.Errorf("unexpected Rrcstatus %v", rrcstatus)
	}

	return &ret, nil
}

func decodeRrcstatus(rrcstatusC *C.RRCStatus_t) (*e2sm_mho.Rrcstatus, error) {

	//ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
	rrcstatus := e2sm_mho.Rrcstatus(int32(*rrcstatusC))

	return &rrcstatus, nil
}
