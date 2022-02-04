// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
)

func CreateE2SmMhoIndicationHeader(cgi *e2sm_mho.CellGlobalId) (*e2sm_mho.E2SmMhoIndicationHeader, error) {

	E2SmMhoPdu := e2sm_mho.E2SmMhoIndicationHeader{
		E2SmMhoIndicationHeader: &e2sm_mho.E2SmMhoIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: &e2sm_mho.E2SmMhoIndicationHeaderFormat1{
				Cgi: cgi,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmMhoPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}
