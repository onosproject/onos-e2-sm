// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
)

func CreateE2SmRcPreIndicationMsgFormat1(plmnIDBytes []byte, arfcn *e2sm_rc_pre_v2.Arfcn, cellSize e2sm_rc_pre_v2.CellSize,
	pci int32, neighbor []*e2sm_rc_pre_v2.Nrt) (*e2sm_rc_pre_v2.E2SmRcPreIndicationMessage, error) {
	if len(plmnIDBytes) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 bytes, actual length is %d", len(plmnIDBytes))
	}

	e2SmIindicationMsg := e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2sm_rc_pre_v2.E2SmRcPreIndicationMessageFormat1{
			Neighbors: neighbor,
			DlArfcn:   arfcn,
			CellSize:  cellSize,
			Pci: &e2sm_rc_pre_v2.Pci{
				Value: pci,
			},
		},
	}

	//e2SmIindicationMsg.IndicationMessageFormat1.Neighbors = append(e2SmIindicationMsg.IndicationMessageFormat1.Neighbors, neighbors)

	E2SmRcPrePdu := e2sm_rc_pre_v2.E2SmRcPreIndicationMessage{
		E2SmRcPreIndicationMessage: &e2SmIindicationMsg,
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}

func CreateE2SmRcPreIndicationMsgRicStyleType(ricStyleType int32, plmnIDBytes []byte) (*e2sm_rc_pre_v2.E2SmRcPreIndicationMessage, error) {
	if len(plmnIDBytes) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 bytes, actual length is %d", len(plmnIDBytes))
	}

	E2SmRcPrePdu := e2sm_rc_pre_v2.E2SmRcPreIndicationMessage{
		E2SmRcPreIndicationMessage: &e2sm_rc_pre_v2.E2SmRcPreIndicationMessage_RicStyleType{
			RicStyleType: &e2sm_rc_pre_v2.RicStyleType{
				Value: ricStyleType,
			},
		},
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}

func CreateEArfcn(value int32) *e2sm_rc_pre_v2.Arfcn {

	return &e2sm_rc_pre_v2.Arfcn{
		Arfcn: &e2sm_rc_pre_v2.Arfcn_EArfcn{
			EArfcn: &e2sm_rc_pre_v2.Earfcn{
				Value: value,
			},
		},
	}
}

func CreateNrArfcn(value int32) *e2sm_rc_pre_v2.Arfcn {

	return &e2sm_rc_pre_v2.Arfcn{
		Arfcn: &e2sm_rc_pre_v2.Arfcn_NrArfcn{
			NrArfcn: &e2sm_rc_pre_v2.Nrarfcn{
				Value: value,
			},
		},
	}
}

func CreateNrt(cgi *e2sm_rc_pre_v2.CellGlobalId, arfcn *e2sm_rc_pre_v2.Arfcn, cellSize e2sm_rc_pre_v2.CellSize, pci *e2sm_rc_pre_v2.Pci) (*e2sm_rc_pre_v2.Nrt, error) {

	nrt := e2sm_rc_pre_v2.Nrt{
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
