// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes
//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "OCUCP-PF-Container.h"
import "C"
import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeOCUCPPFContainer(oCuCpContainer *e2sm_kpm_ies.OcucpPfContainer) ([]byte, error) {
	oCuCpContainerCP, err := newOCUCPPFContainer(oCuCpContainer)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_OCUCP_PF_Container, unsafe.Pointer(oCuCpContainerCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newOCUCPPFContainer(oCuCpC *e2sm_kpm_ies.OcucpPfContainer) (*C.OCUCP_PF_Container_t, error) {
	//var pr C.OCUCP_PF_Container
	var nUEs = C.long(oCuCpC.GetCuCpResourceStatus().NumberOfActiveUes)

	//gnbNameCP, err := newGnbCuCpName(oCuCpC.GNbCuCpName)

	//if err != nil {
	//	return nil, err
	//}

	ocucpC := C.OCUCP_PF_Container_t{
		//gNB_CU_CP_Name: gnbNameCP,
		cu_CP_Resource_Status: C.struct_OCUCP_PF_Container__cu_CP_Resource_Status{
			numberOfActive_UEs: &nUEs,
		},
	}

	return &ocucpC, nil
}

//func decodeOCUCPPFContainer(gnbIDcC *C.GNB_CU_CP_Name_t) (*e2sm_kpm_ies.GnbIdChoice, error) {
//	return nil, fmt.Errorf("decodeGnbIDChoice() not yet implemented")
//}