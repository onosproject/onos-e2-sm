// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GNB-CU-CP-Name.h"
import "C"
import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
)

func newGnbCuCpName(gnbCuCpNamec *e2sm_kpm_ies.GnbCuCpName) (*C.GNB_CU_CP_Name_t, error) {

	gnbName := newPrintableString(gnbCuCpNamec.GetValue())
	gnbcucpcnameC := (*C.GNB_CU_CP_Name_t)(gnbName)

	return gnbcucpcnameC, nil
}

func decodeGnbCuCpName(gnbCuCpNameC *C.GNB_CU_CP_Name_t) *e2sm_kpm_ies.GnbCuCpName {
	if gnbCuCpNameC == nil {
		return nil
	}

	gnbNamePS := decodePrintableString(gnbCuCpNameC)

	gnbName := &e2sm_kpm_ies.GnbCuCpName{
		Value: gnbNamePS,
	}

	return gnbName
}
