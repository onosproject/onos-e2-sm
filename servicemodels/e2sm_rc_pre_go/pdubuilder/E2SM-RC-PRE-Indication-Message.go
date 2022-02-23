// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
)

func CreateE2SmRcPreIndicationMsgFormat1(plmnIDBytes []byte, arfcn *e2sm_rc_pre_go.Arfcn, cellSize e2sm_rc_pre_go.CellSize,
	pci int32, neighbor []*e2sm_rc_pre_go.Nrt) (*e2sm_rc_pre_go.E2SmRcPreIndicationMessage, error) {
	if len(plmnIDBytes) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 bytes, actual length is %d", len(plmnIDBytes))
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
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
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
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}

func CreateEArfcn(value int32) *e2sm_rc_pre_go.Arfcn {

	return &e2sm_rc_pre_go.Arfcn{
		Arfcn: &e2sm_rc_pre_go.Arfcn_EArfcn{
			EArfcn: &e2sm_rc_pre_go.Earfcn{
				Value: value,
			},
		},
	}
}

func CreateNrArfcn(value int32) *e2sm_rc_pre_go.Arfcn {

	return &e2sm_rc_pre_go.Arfcn{
		Arfcn: &e2sm_rc_pre_go.Arfcn_NrArfcn{
			NrArfcn: &e2sm_rc_pre_go.Nrarfcn{
				Value: value,
			},
		},
	}
}

func CreateNrt(cgi *e2sm_rc_pre_go.CellGlobalId, arfcn *e2sm_rc_pre_go.Arfcn, cellSize e2sm_rc_pre_go.CellSize, pci *e2sm_rc_pre_go.Pci) (*e2sm_rc_pre_go.Nrt, error) {

	nrt := e2sm_rc_pre_go.Nrt{
		Cgi:      cgi,
		DlArfcn:  arfcn,
		CellSize: cellSize,
		Pci:      pci,
	}

	if err := nrt.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &nrt, nil
}
