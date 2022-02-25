// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcPreControlHeader() (*e2smrcprev2.E2SmRcPreControlHeader, error) {

	e2smRcPreFormat1 := e2smrcprev2.E2SmRcPreControlHeaderFormat1{
		RcCommand: e2smrcprev2.RcPreCommand_RC_PRE_COMMAND_SET_PARAMETERS,
	}

	e2smRcPrePdu := e2smrcprev2.E2SmRcPreControlHeader{
		E2SmRcPreControlHeader: &e2smrcprev2.E2SmRcPreControlHeader_ControlHeaderFormat1{
			ControlHeaderFormat1: &e2smRcPreFormat1,
		},
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}

func CreateCellGlobalIDEUTRACGI(plmnIDBytes []byte, cellID *asn1.BitString) (*e2smrcprev2.CellGlobalId, error) {

	if len(plmnIDBytes) != 3 {
		return nil, errors.NewInvalid("error: Plmn ID should be 3 chars")
	}
	if cellID.Len != 28 {
		return nil, errors.NewInvalid("EutraCgi should be of length 28")
	}

	cgi := e2smrcprev2.CellGlobalId{
		CellGlobalId: &e2smrcprev2.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2smrcprev2.Eutracgi{
				PLmnIdentity: &e2smrcprev2.PlmnIdentity{
					Value: plmnIDBytes,
				},
				EUtracellIdentity: &e2smrcprev2.EutracellIdentity{
					Value: cellID,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &cgi, nil
}

func CreateCellGlobalIDNrCgi(plmnIDBytes []byte, cellID *asn1.BitString) (*e2smrcprev2.CellGlobalId, error) {

	if len(plmnIDBytes) != 3 {
		return nil, errors.NewInvalid("error: Plmn ID should be 3 chars")
	}
	if cellID.Len != 36 {
		return nil, errors.NewInvalid("NrCgi should be of length 36")
	}

	cgi := e2smrcprev2.CellGlobalId{
		CellGlobalId: &e2smrcprev2.CellGlobalId_NrCgi{
			NrCgi: &e2smrcprev2.Nrcgi{
				PLmnIdentity: &e2smrcprev2.PlmnIdentity{
					Value: plmnIDBytes,
				},
				NRcellIdentity: &e2smrcprev2.NrcellIdentity{
					Value: cellID,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &cgi, nil
}
