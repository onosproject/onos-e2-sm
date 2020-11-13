// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PF-Container.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodePFContainer(pfContainer *e2sm_kpm_ies.PfContainer) ([]byte, error) {

	pfContainerCP, err := newPfContainer(pfContainer)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_PF_Container, unsafe.Pointer(pfContainerCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newPfContainer(pfContainer *e2sm_kpm_ies.PfContainer) (*C.PF_Container_t, error) {
	var pr C.PF_Container_PR
	choiceC := [8]byte{}

	switch choice := pfContainer.PfContainer.(type) {
	case *e2sm_kpm_ies.PfContainer_OCuCp:
		pr = C.PF_Container_PR_oCU_CP

		im, err := newOCuCpPfContainer(choice.OCuCp)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newPfContainer() %T not yet implemented", choice)
	}

	pfContainerC := C.PF_Container_t{
		present: pr,
		choice:  choiceC,
	}

	return &pfContainerC, nil
}

func decodePfContainer(pfContainerC *C.PF_Container_t) (*e2sm_kpm_ies.PfContainer, error) {
	pfContainer := new(e2sm_kpm_ies.PfContainer)

	switch pfContainerC.present {
	case C.PF_Container_PR_oCU_CP:
		oCuCpPfContainer := decodeOCuCpContainerBytes(pfContainerC.choice)

		pfContainer.PfContainer = &e2sm_kpm_ies.PfContainer_OCuCp{
			OCuCp: oCuCpPfContainer,
		}

	default:
		return nil, fmt.Errorf("decodePfContainer() %v not yet implemented", pfContainerC.present)
	}

	return pfContainer, nil
}
