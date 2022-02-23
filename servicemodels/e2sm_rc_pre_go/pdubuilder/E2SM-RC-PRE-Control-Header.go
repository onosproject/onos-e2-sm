// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmRcPreControlHeader() (*e2sm_rc_pre_go.E2SmRcPreControlHeader, error) {

	e2smRcPreFormat1 := e2sm_rc_pre_go.E2SmRcPreControlHeaderFormat1{
		RcCommand: e2sm_rc_pre_go.RcPreCommand_RC_PRE_COMMAND_SET_PARAMETERS,
	}

	e2smRcPrePdu := e2sm_rc_pre_go.E2SmRcPreControlHeader{
		E2SmRcPreControlHeader: &e2sm_rc_pre_go.E2SmRcPreControlHeader_ControlHeaderFormat1{
			ControlHeaderFormat1: &e2smRcPreFormat1,
		},
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}

func CreateCellGlobalIDEUTRACGI(plmnIDBytes []byte, cellID *asn1.BitString) (*e2sm_rc_pre_go.CellGlobalId, error) {

	if len(plmnIDBytes) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}
	if cellID.Len != 28 {
		return nil, fmt.Errorf("EutraCgi should be of length 28")
	}

	cgi := e2sm_rc_pre_go.CellGlobalId{
		CellGlobalId: &e2sm_rc_pre_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_rc_pre_go.Eutracgi{
				PLmnIdentity: &e2sm_rc_pre_go.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2sm_rc_pre_go.EutracellIdentity{
					Value: cellID,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &cgi, nil
}

func CreateCellGlobalIDNrCgi(plmnIDBytes []byte, cellID *asn1.BitString) (*e2sm_rc_pre_go.CellGlobalId, error) {

	if len(plmnIDBytes) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}
	if cellID.Len != 36 {
		return nil, fmt.Errorf("NrCgi should be of length 36")
	}

	cgi := e2sm_rc_pre_go.CellGlobalId{
		CellGlobalId: &e2sm_rc_pre_go.CellGlobalId_NrCgi{
			NrCgi: &e2sm_rc_pre_go.Nrcgi{
				PLmnIdentity: &e2sm_rc_pre_go.PlmnIdentity{
					Value: plmnIDBytes,
				},
				NRcellIdentity: &e2sm_rc_pre_go.NrcellIdentity{
					Value: cellID,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &cgi, nil
}
