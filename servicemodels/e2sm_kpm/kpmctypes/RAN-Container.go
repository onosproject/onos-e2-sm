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

func xerEncodeRanContainer(ranContainer *e2sm_kpm_ies.RanContainer) ([]byte, error) {

	ranContainerCP := newRanContainer(ranContainer)

	bytes, err := encodeXer(&C.asn_DEF_RAN_Container, unsafe.Pointer(ranContainerCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanContainer() %s", err.Error())
	}
	return bytes, nil
}

func newRanContainer(ranContainer *e2sm_kpm_ies.RanContainer) *C.RAN_Container_t {

	return newOctetString(string(ranContainer.GetValue()))
}

func decodeRanContainer(ranContainerC *C.RAN_Container_t) *e2sm_kpm_ies.RanContainer {
	return &e2sm_kpm_ies.RanContainer{
		Value: []byte(decodeOctetString(ranContainerC)),
	}
}
