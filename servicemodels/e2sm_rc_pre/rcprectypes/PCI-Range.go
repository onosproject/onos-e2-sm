// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PCI-Range.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

func xerEncodePCIRange(pciRange *e2sm_rc_pre_ies.PciRange) ([]byte, error) {

	pciRangeCP, err := newPciRange(pciRange)
	if err != nil {
		return nil, fmt.Errorf("xerEncodePciRange() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_PCI_Range, unsafe.Pointer(pciRangeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodePciRange() %s", err.Error())
	}
	return bytes, nil
}

func newPciRange(pciRange *e2sm_rc_pre_ies.PciRange) (*C.PCI_Range_t, error) {

	lowerPci := newPCI(pciRange.LowerPci)
	upperPci := newPCI(pciRange.UpperPci)

	pciRangeC := C.PCI_Range_t{
		lower_pci: *lowerPci,
		upper_pci: *upperPci,
	}

	return &pciRangeC, nil
}

func decodePciRange(pciRangeC *C.PCI_Range_t) (*e2sm_rc_pre_ies.PciRange, error) {
	pciRange := new(e2sm_rc_pre_ies.PciRange)

	lowerPci := decodePCI(&pciRangeC.lower_pci)
	pciRange.LowerPci = lowerPci

	upperPci := decodePCI(&pciRangeC.upper_pci)
	pciRange.UpperPci = upperPci

	return pciRange, nil
}
