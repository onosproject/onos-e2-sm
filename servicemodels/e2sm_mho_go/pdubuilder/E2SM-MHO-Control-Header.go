// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"fmt"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
)

func CreateE2SmMhoControlHeader(controlMessagePriority int32) (*e2sm_mho_go.E2SmMhoControlHeader, error) {

	e2smMhoFormat1 := e2sm_mho_go.E2SmMhoControlHeaderFormat1{
		RcCommand: e2sm_mho_go.MhoCommand_MHO_COMMAND_INITIATE_HANDOVER,
		RicControlMessagePriority: &e2sm_mho_go.RicControlMessagePriority{
			Value: controlMessagePriority,
		},
	}
	e2smMhoPdu := e2sm_mho_go.E2SmMhoControlHeader{
		E2SmMhoControlHeader: &e2sm_mho_go.E2SmMhoControlHeader_ControlHeaderFormat1{
			ControlHeaderFormat1: &e2smMhoFormat1,
		},
	}

	if err := e2smMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("CreateE2SmMhoControlHeader(): error validating E2SmPDU %s", err.Error())
	}
	return &e2smMhoPdu, nil
}
