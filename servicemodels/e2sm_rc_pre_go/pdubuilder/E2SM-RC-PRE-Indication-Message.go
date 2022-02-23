// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcPreIndicationMsgFormat1(plmnIDBytes []byte, arfcn *e2sm_rc_pre_go.Arfcn, cellSize e2sm_rc_pre_go.CellSize,
	pci int32, neighbor []*e2sm_rc_pre_go.Nrt) (*e2sm_rc_pre_go.E2SmRcPreIndicationMessage, error) {
	if len(plmnIDBytes) != 3 {
		return nil, errors.NewInvalid("error: Plmn ID should be 3 bytes, actual length is %d", len(plmnIDBytes))
	}

	e2SmIindicationMsg := e2sm_rc_pre_go.E2SmRcPreIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2sm_rc_pre_go.E2SmRcPreIndicationMessageFormat1{
			Neighbors: neighbor,
			DlArfcn:   arfcn,
			CellSize:  cellSize,
			Pci: &e2sm_rc_pre_go.Pci{
				Value: pci,
			},
		},
	}

	E2SmRcPrePdu := e2sm_rc_pre_go.E2SmRcPreIndicationMessage{
		E2SmRcPreIndicationMessage: &e2SmIindicationMsg,
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}

func CreateE2SmRcPreIndicationMsgRicStyleType(ricStyleType int32) (*e2sm_rc_pre_go.E2SmRcPreIndicationMessage, error) {

	E2SmRcPrePdu := e2sm_rc_pre_go.E2SmRcPreIndicationMessage{
		E2SmRcPreIndicationMessage: &e2sm_rc_pre_go.E2SmRcPreIndicationMessage_RicStyleType{
			RicStyleType: &e2sm_rc_pre_go.RicStyleType{
				Value: ricStyleType,
			},
		},
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}

func CreateEArfcn(value int32) (*e2sm_rc_pre_go.Arfcn, error) {

	earfcn := &e2sm_rc_pre_go.Arfcn{
		Arfcn: &e2sm_rc_pre_go.Arfcn_EArfcn{
			EArfcn: &e2sm_rc_pre_go.Earfcn{
				Value: value,
			},
		},
	}

	if err := earfcn.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEArfcn(): error validating E2SmPDU %s", err.Error())
	}
	return earfcn, nil
}

func CreateNrArfcn(value int32) (*e2sm_rc_pre_go.Arfcn, error) {

	nrarfcn := &e2sm_rc_pre_go.Arfcn{
		Arfcn: &e2sm_rc_pre_go.Arfcn_NrArfcn{
			NrArfcn: &e2sm_rc_pre_go.Nrarfcn{
				Value: value,
			},
		},
	}

	if err := nrarfcn.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrArfcn(): error validating E2SmPDU %s", err.Error())
	}
	return nrarfcn, nil
}

func CreateNrt(cgi *e2sm_rc_pre_go.CellGlobalId, arfcn *e2sm_rc_pre_go.Arfcn, cellSize e2sm_rc_pre_go.CellSize, pci *e2sm_rc_pre_go.Pci) (*e2sm_rc_pre_go.Nrt, error) {

	nrt := e2sm_rc_pre_go.Nrt{
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
