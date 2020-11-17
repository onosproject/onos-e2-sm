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
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodePmContainersListItem(pmContainersList *e2sm_kpm_ies.PmContainersList) ([]byte, error) {

	pmContainersListCP, err := newPmContainersListItem(pmContainersList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodePmContainersListItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_PM_Containers_List, unsafe.Pointer(pmContainersListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodePmContainersListItem() %s", err.Error())
	}
	return bytes, nil
}


func newPmContainersListItem(pmContainersList *e2sm_kpm_ies.PmContainersList) (*C.PM_Containers_List_t, error) {

	pmContainersListItemCP, err := newPfContainer(pmContainersList.PerformanceContainer)
	if err != nil {
		return nil, fmt.Errorf("newPmContainersListItem() %s", err.Error())
	}

	pmContainersListC := C.PM_Containers_List_t{
		performanceContainer: pmContainersListItemCP,
		// TODO: Implement TheRancontainer
		//theRANContainer:  nil, //*C.theRANContainer,
	}

	return &pmContainersListC, nil
}


func decodePmContainersListItem(pmContainersListC *C.PM_Containers_List_t) (*e2sm_kpm_ies.PmContainersList, error) {
	pmContainersListItem := new(e2sm_kpm_ies.PmContainersList)
	performanceContainer, err := decodePfContainer(pmContainersListC.performanceContainer)
	if err != nil {
		return nil, fmt.Errorf("decodePmContainersListItem() %s", err.Error())
	}
	pmContainersListItem.PerformanceContainer = performanceContainer

	//TODO: Implement decodeTheRancontainer function
	//theRancontainer, err := decodeTheRancontainer(pmContainersListC.theRANContainer)
	//if err != nil {
	//	return nil, err
	//}
	//pmContainersListItem.TheRancontainer = theRancontainer

	return pmContainersListItem, nil
}