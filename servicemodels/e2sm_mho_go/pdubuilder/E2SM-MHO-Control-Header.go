// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smmhov2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMhoControlHeader(controlMessagePriority int32) (*e2smmhov2.E2SmMhoControlHeader, error) {

	e2smMhoFormat1 := e2smmhov2.E2SmMhoControlHeaderFormat1{
		RcCommand: e2smmhov2.MhoCommand_MHO_COMMAND_INITIATE_HANDOVER,
		RicControlMessagePriority: &e2smmhov2.RicControlMessagePriority{
			Value: controlMessagePriority,
		},
	}
	e2smMhoPdu := e2smmhov2.E2SmMhoControlHeader{
		E2SmMhoControlHeader: &e2smmhov2.E2SmMhoControlHeader_ControlHeaderFormat1{
			ControlHeaderFormat1: &e2smMhoFormat1,
		},
	}

	if err := e2smMhoPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMhoControlHeader(): error validating E2SmPDU %s", err.Error())
	}
	return &e2smMhoPdu, nil
}
