// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"fmt"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
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

	//if err := e2smMhoPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	//}
	return &e2smMhoPdu, nil
}

func CreateCellGlobalIDEUTRACGI(plmnIDBytes []byte, cellID *asn1.BitString) (*e2sm_mho_go.CellGlobalId, error) {

	if len(plmnIDBytes) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}
	if cellID.Len != 28 {
		return nil, fmt.Errorf("EutraCgi should be of length 28")
	}

	cgi := e2sm_mho_go.CellGlobalId{
		CellGlobalId: &e2sm_mho_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_mho_go.Eutracgi{
				PLmnIdentity: &e2sm_mho_go.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_mho_go.EutracellIdentity{
					Value: cellID,
				},
			},
		},
	}

	//if err := cgi.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	//}
	return &cgi, nil
}

func CreateCellGlobalIDNrCgi(plmnIDBytes []byte, cellID *asn1.BitString) (*e2sm_mho_go.CellGlobalId, error) {

	if len(plmnIDBytes) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}
	if cellID.Len != 36 {
		return nil, fmt.Errorf("NrCgi should be of length 36")
	}

	cgi := e2sm_mho_go.CellGlobalId{
		CellGlobalId: &e2sm_mho_go.CellGlobalId_NrCgi{
			NrCgi: &e2sm_mho_go.Nrcgi{
				PLmnIdentity: &e2sm_mho_go.PlmnIdentity{
					Value: plmnIDBytes,
				},
				NRcellIdentity: &e2sm_mho_go.NrcellIdentity{
					Value: cellID,
				},
			},
		},
	}

	//if err := cgi.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	//}
	return &cgi, nil
}
