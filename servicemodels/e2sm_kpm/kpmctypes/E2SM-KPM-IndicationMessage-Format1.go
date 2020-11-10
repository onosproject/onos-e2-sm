// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2SM-KPM-IndicationMessage-Format1.h"
import "C"
import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMsgFormat1 *e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1) ([]byte, error) {

	e2SmKpmIndicationMsgFormat1CP, err := newE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMsgFormat1)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_E2SM_KPM_IndicationMessage_Format1, unsafe.Pointer(e2SmKpmIndicationMsgFormat1CP))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMsgFormat1 *e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1) (*C.E2SM_KPM_IndicationMessage_Format1_t, error) {
	pmContainersListC := new(C.E2SM_KPM_IndicationMessage_Format1_t)
	for _, pmContainersItem := range e2SmKpmIndicationMsgFormat1.IndicationMessageFormat1.GetPmContainers() {
		pmContainersItemC, err := newPmContainersListItem(pmContainersItem)
		if err != nil {
			return nil, fmt.Errorf("newPmContainersListItem() %s", err.Error())
		}

		if _, err = C.asn_sequence_add(unsafe.Pointer(pmContainersListC), unsafe.Pointer(pmContainersItemC)); err != nil {
			return nil, err
		}
	}

	return pmContainersListC, nil
}

//func decodeOCUCPPFContainer(gnbIDcC *C.GNB_CU_CP_Name_t) (*e2sm_kpm_ies.GnbIdChoice, error) {
//	return nil, fmt.Errorf("decodeGnbIDChoice() not yet implemented")
//}
