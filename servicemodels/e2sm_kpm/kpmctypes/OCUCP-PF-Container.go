// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "OCUCP-PF-Container.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeOCuCpPfContainer(oCuCpContainer *e2sm_kpm_ies.OcucpPfContainer) ([]byte, error) {
	oCuCpContainerCP, err := newOCuCpPfContainer(oCuCpContainer)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeOCuCpPfContainer() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_OCUCP_PF_Container, unsafe.Pointer(oCuCpContainerCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeOCuCpPfContainer() %s", err.Error())
	}
	return bytes, nil
}

func newOCuCpPfContainer(oCuCpC *e2sm_kpm_ies.OcucpPfContainer) (*C.OCUCP_PF_Container_t, error) {
	var nUEs = C.long(oCuCpC.GetCuCpResourceStatus().NumberOfActiveUes)

	gnbNameCP, err := newGnbCuCpName(oCuCpC.GNbCuCpName)
	if err != nil {
		return nil, fmt.Errorf("newOCuCpPfContainer() %s", err.Error())
	}

	// TODO: Fix an issue with Go pointers (see output of 'go test -v ./...' command)
	ocucpC := C.OCUCP_PF_Container_t{
		gNB_CU_CP_Name: gnbNameCP,
		cu_CP_Resource_Status: C.struct_OCUCP_PF_Container__cu_CP_Resource_Status{
			numberOfActive_UEs: &nUEs,
		},
	}

	return &ocucpC, nil
}

func decodeOCuCpPfContainer(oCuCpContainerC *C.OCUCP_PF_Container_t) *e2sm_kpm_ies.OcucpPfContainer {
	oCuCpContainer := new(e2sm_kpm_ies.OcucpPfContainer)

	oCuCpContainer.GNbCuCpName = decodeGnbCuCpName(oCuCpContainerC.gNB_CU_CP_Name)
	//TODO: Implement decodeCuCpResourceStatus
	nUes := int32(*oCuCpContainerC.cu_CP_Resource_Status.numberOfActive_UEs)
	oCuCpContainer.CuCpResourceStatus = &e2sm_kpm_ies.OcucpPfContainer_CuCpResourceStatus001{
		NumberOfActiveUes: nUes,
	}

	return oCuCpContainer
}

func decodeOCuCpContainerBytes(array [8]byte) *e2sm_kpm_ies.OcucpPfContainer {
	oCuCpPfContainerC := (*C.OCUCP_PF_Container_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeOCuCpPfContainer(oCuCpPfContainerC)
}
