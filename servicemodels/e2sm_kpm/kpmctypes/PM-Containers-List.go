// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes
//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PM-Containers-List.h"
import "C"
import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodePmContainersListItem(pmContainersList *e2sm_kpm_ies.PmContainersList) ([]byte, error) {

	pmContainersListCP, err := newPmContainersListItem(pmContainersList)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_PM_Containers_List, unsafe.Pointer(pmContainersListCP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}


func newPmContainersListItem(pmContainersList *e2sm_kpm_ies.PmContainersList) (*C.PM_Containers_List_t, error) {

	pmContainersListItemCP, err := newPfContainer(pmContainersList.PerformanceContainer)

	if err != nil {
		return nil, err
	}

	pmContainersListC := C.PM_Containers_List_t{
		performanceContainer: pmContainersListItemCP,
		// TODO: Implement TheRancontainer
		//theRANContainer:  nil, //*C.theRANContainer,
	}

	return &pmContainersListC, nil
}


//func decodeOCUCPPFContainer(gnbIDcC *C.GNB_CU_CP_Name_t) (*e2sm_kpm_ies.GnbIdChoice, error) {
//	return nil, fmt.Errorf("decodeGnbIDChoice() not yet implemented")
//}