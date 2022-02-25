// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcPreIndicationMsgFormat1(plmnIDBytes []byte, arfcn *e2smrcprev2.Arfcn, cellSize e2smrcprev2.CellSize,
	pci int32, neighbor []*e2smrcprev2.Nrt) (*e2smrcprev2.E2SmRcPreIndicationMessage, error) {
	if len(plmnIDBytes) != 3 {
		return nil, errors.NewInvalid("error: Plmn ID should be 3 bytes, actual length is %d", len(plmnIDBytes))
	}

	e2SmIindicationMsg := e2smrcprev2.E2SmRcPreIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2smrcprev2.E2SmRcPreIndicationMessageFormat1{
			Neighbors: neighbor,
			DlArfcn:   arfcn,
			CellSize:  cellSize,
			Pci: &e2smrcprev2.Pci{
				Value: pci,
			},
		},
	}

	E2SmRcPrePdu := e2smrcprev2.E2SmRcPreIndicationMessage{
		E2SmRcPreIndicationMessage: &e2SmIindicationMsg,
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}

func CreateE2SmRcPreIndicationMsgRicStyleType(ricStyleType int32) (*e2smrcprev2.E2SmRcPreIndicationMessage, error) {

	E2SmRcPrePdu := e2smrcprev2.E2SmRcPreIndicationMessage{
		E2SmRcPreIndicationMessage: &e2smrcprev2.E2SmRcPreIndicationMessage_RicStyleType{
			RicStyleType: &e2smrcprev2.RicStyleType{
				Value: ricStyleType,
			},
		},
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}

func CreateEArfcn(value int32) (*e2smrcprev2.Arfcn, error) {

	earfcn := &e2smrcprev2.Arfcn{
		Arfcn: &e2smrcprev2.Arfcn_EArfcn{
			EArfcn: &e2smrcprev2.Earfcn{
				Value: value,
			},
		},
	}

	if err := earfcn.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEArfcn(): error validating E2SmPDU %s", err.Error())
	}
	return earfcn, nil
}

func CreateNrArfcn(value int32) (*e2smrcprev2.Arfcn, error) {

	nrarfcn := &e2smrcprev2.Arfcn{
		Arfcn: &e2smrcprev2.Arfcn_NrArfcn{
			NrArfcn: &e2smrcprev2.Nrarfcn{
				Value: value,
			},
		},
	}

	if err := nrarfcn.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrArfcn(): error validating E2SmPDU %s", err.Error())
	}
	return nrarfcn, nil
}

func CreateNrt(cgi *e2smrcprev2.CellGlobalId, arfcn *e2smrcprev2.Arfcn, cellSize e2smrcprev2.CellSize, pci *e2smrcprev2.Pci) (*e2smrcprev2.Nrt, error) {

	nrt := e2smrcprev2.Nrt{
		Cgi:      cgi,
		DlArfcn:  arfcn,
		CellSize: cellSize,
		Pci:      pci,
	}

	if err := nrt.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrt(): error validating E2SmPDU %s", err.Error())
	}
	return &nrt, nil
}
