// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
)

func CreateE2SmRcPreIndicationHeader(cgi *e2sm_rc_pre_v2.CellGlobalId) (*e2sm_rc_pre_v2.E2SmRcPreIndicationHeader, error) {

	E2SmRcPrePdu := e2sm_rc_pre_v2.E2SmRcPreIndicationHeader{
		E2SmRcPreIndicationHeader: &e2sm_rc_pre_v2.E2SmRcPreIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: &e2sm_rc_pre_v2.E2SmRcPreIndicationHeaderFormat1{
				Cgi: cgi,
			},
		},
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmRcPrePDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}
