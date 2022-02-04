// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmMhoIndicationMsgFormat1(ueID *e2sm_mho_go.UeIdentity, measReport []*e2sm_mho_go.E2SmMhoMeasurementReportItem) (*e2sm_mho_go.E2SmMhoIndicationMessage, error) {

	E2SmMhoPdu := e2sm_mho_go.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho_go.E2SmMhoIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: &e2sm_mho_go.E2SmMhoIndicationMessageFormat1{
				UeId:       ueID,
				MeasReport: measReport,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("CreateE2SmMhoIndicationMsgFormat1(): error validating E2SmPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}

func CreateE2SmMhoIndicationMsgFormat2(ueID *e2sm_mho_go.UeIdentity, rrcStatus e2sm_mho_go.Rrcstatus) (*e2sm_mho_go.E2SmMhoIndicationMessage, error) {
	E2SmMhoPdu := e2sm_mho_go.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho_go.E2SmMhoIndicationMessage_IndicationMessageFormat2{
			IndicationMessageFormat2: &e2sm_mho_go.E2SmMhoIndicationMessageFormat2{
				UeId:      ueID,
				RrcStatus: rrcStatus,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("CreateE2SmMhoIndicationMsgFormat2(): error validating E2SmPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}

func CreateMeasurementRecordItem(cgi *e2sm_mho_go.CellGlobalId, rsrp *e2sm_mho_go.Rsrp) (*e2sm_mho_go.E2SmMhoMeasurementReportItem, error) {
	res := &e2sm_mho_go.E2SmMhoMeasurementReportItem{
		Cgi:  cgi,
		Rsrp: rsrp,
	}

	if err := res.Validate(); err != nil {
		return nil, fmt.Errorf("CreateMeasurementRecordItem(): error validationg E2SmPDU %s", err)
	}

	return res, nil
}

func CreateCellGlobalIDNrCGI(plmnID []byte, nrCGI *asn1.BitString) (*e2sm_mho_go.CellGlobalId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("CreateCellGlobalIDNrCGI(): PlmnID should contain only 3 bytes, got %v", len(plmnID))
	}

	if nrCGI.Len != uint32(36) {
		return nil, fmt.Errorf("CreateCellGlobalIDNrCGI(): EutraCellIdentity should be of length 36 bits, got %v", nrCGI.Len)
	}
	if len(nrCGI.Value) != 5 {
		return nil, fmt.Errorf("CreateCellGlobalIDNrCGI(): EutraCellIdentity should be of length 36 bits (5 bytes), got %v bytes", len(nrCGI.Value))
	}

	cgi := &e2sm_mho_go.CellGlobalId{
		CellGlobalId: &e2sm_mho_go.CellGlobalId_NrCgi{
			NrCgi: &e2sm_mho_go.Nrcgi{
				PLmnIdentity: &e2sm_mho_go.PlmnIdentity{
					Value: plmnID,
				},
				NRcellIdentity: &e2sm_mho_go.NrcellIdentity{
					Value: nrCGI,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
		return nil, fmt.Errorf("CreateCellGlobalIDNrCGI(): error validationg E2SmPDU %s", err)
	}

	return cgi, nil
}

func CreateCellGlobalIDEutraCGI(plmnID []byte, eutraCGI *asn1.BitString) (*e2sm_mho_go.CellGlobalId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("CreateCellGlobalIDEutraCGI(): PlmnID should contain only 3 bytes, got %v", len(plmnID))
	}

	if eutraCGI.Len != uint32(28) {
		return nil, fmt.Errorf("CreateCellGlobalIDEutraCGI(): EutraCellIdentity should be of length 28 bits, got %v", eutraCGI.Len)
	}
	if len(eutraCGI.Value) != 4 {
		return nil, fmt.Errorf("CreateCellGlobalIDEutraCGI(): EutraCellIdentity should be of length 28 bits (4 bytes), got %v bytes", len(eutraCGI.Value))
	}

	cgi := &e2sm_mho_go.CellGlobalId{
		CellGlobalId: &e2sm_mho_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_mho_go.Eutracgi{
				PLmnIdentity: &e2sm_mho_go.PlmnIdentity{
					Value: plmnID,
				},
				EUtracellIdentity: &e2sm_mho_go.EutracellIdentity{
					Value: eutraCGI,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
		return nil, fmt.Errorf("CreateCellGlobalIDEutraCGI(): error validationg E2SmPDU %s", err)
	}

	return cgi, nil
}

func CreateRrcStatusConnected() e2sm_mho_go.Rrcstatus {
	return e2sm_mho_go.Rrcstatus_RRCSTATUS_CONNECTED
}

func CreateRrcStatusInactive() e2sm_mho_go.Rrcstatus {
	return e2sm_mho_go.Rrcstatus_RRCSTATUS_INACTIVE
}

func CreateRrcStatusIdle() e2sm_mho_go.Rrcstatus {
	return e2sm_mho_go.Rrcstatus_RRCSTATUS_IDLE
}
