// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PCI.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func xerEncodePCI(pci *e2sm_rc_pre_ies.Pci) ([]byte, error) {
	pciCP := newPCI(pci)

	bytes, err := encodeXer(&C.asn_DEF_PCI, unsafe.Pointer(pciCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodePCI() %s", err.Error())
	}
	return bytes, nil
}

func perEncodePCI(pci *e2sm_rc_pre_ies.Pci) ([]byte, error) {
	pciCP := newPCI(pci)

	bytes, err := encodePerBuffer(&C.asn_DEF_PCI, unsafe.Pointer(pciCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodePCI() %s", err.Error())
	}
	return bytes, nil
}

func perDecodePCI(bytes []byte) (*e2sm_rc_pre_ies.Pci, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_PCI)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodePCI((*C.PCI_t)(unsafePtr)), nil
}

func newPCI(pci *e2sm_rc_pre_ies.Pci) *C.PCI_t {
	pciC := C.long(pci.Value)
	return &pciC
}

func decodePCI(pciC *C.PCI_t) *e2sm_rc_pre_ies.Pci {
	return &e2sm_rc_pre_ies.Pci{
		Value: int32(*pciC),
	}
}
