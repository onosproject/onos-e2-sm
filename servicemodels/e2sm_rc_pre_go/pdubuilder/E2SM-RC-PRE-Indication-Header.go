// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
)

func CreateE2SmRcPreIndicationHeader(cgi *e2sm_rc_pre_go.CellGlobalId) (*e2sm_rc_pre_go.E2SmRcPreIndicationHeader, error) {

	E2SmRcPrePdu := e2sm_rc_pre_go.E2SmRcPreIndicationHeader{
		E2SmRcPreIndicationHeader: &e2sm_rc_pre_go.E2SmRcPreIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: &e2sm_rc_pre_go.E2SmRcPreIndicationHeaderFormat1{
				Cgi: cgi,
			},
		},
	}

	//if err := E2SmRcPrePdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmRcPrePDU %s", err.Error())
	//}
	return &E2SmRcPrePdu, nil
}
