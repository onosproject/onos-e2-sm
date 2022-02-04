// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ARP-KPMv2.h"
import "C"

import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeArp(arp *e2sm_kpm_v2.Arp) ([]byte, error) {
	arpCP, err := newArp(arp)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeArp() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ARP_KPMv2, unsafe.Pointer(arpCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeArp() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeArp(arp *e2sm_kpm_v2.Arp) ([]byte, error) {
	arpCP, err := newArp(arp)
	if err != nil {
		return nil, fmt.Errorf("perEncodeArp() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ARP_KPMv2, unsafe.Pointer(arpCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeArp() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeArp(bytes []byte) (*e2sm_kpm_v2.Arp, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ARP_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeArp((*C.ARP_KPMv2_t)(unsafePtr))
}

func perDecodeArp(bytes []byte) (*e2sm_kpm_v2.Arp, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ARP_KPMv2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeArp((*C.ARP_KPMv2_t)(unsafePtr))
}

func newArp(arp *e2sm_kpm_v2.Arp) (*C.ARP_KPMv2_t, error) {

	arpC := C.long(arp.Value)

	return &arpC, nil
}

func decodeArp(arpC *C.ARP_KPMv2_t) (*e2sm_kpm_v2.Arp, error) {

	arp := e2sm_kpm_v2.Arp{
		Value: int32(*arpC),
	}

	return &arp, nil
}
