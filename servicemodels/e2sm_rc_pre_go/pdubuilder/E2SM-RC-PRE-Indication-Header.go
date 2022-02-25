// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcPreIndicationHeader(cgi *e2smrcprev2.CellGlobalId) (*e2smrcprev2.E2SmRcPreIndicationHeader, error) {

	E2SmRcPrePdu := e2smrcprev2.E2SmRcPreIndicationHeader{
		E2SmRcPreIndicationHeader: &e2smrcprev2.E2SmRcPreIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: &e2smrcprev2.E2SmRcPreIndicationHeaderFormat1{
				Cgi: cgi,
			},
		},
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmRcPrePDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}
