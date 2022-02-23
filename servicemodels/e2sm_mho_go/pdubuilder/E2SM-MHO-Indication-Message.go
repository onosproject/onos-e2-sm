// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMhoIndicationMsgFormat1(ueID *e2sm_v2_ies.Ueid, measReport []*e2sm_mho_go.E2SmMhoMeasurementReportItem) (*e2sm_mho_go.E2SmMhoIndicationMessage, error) {

	E2SmMhoPdu := e2sm_mho_go.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho_go.E2SmMhoIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: &e2sm_mho_go.E2SmMhoIndicationMessageFormat1{
				UeId:       ueID,
				MeasReport: measReport,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMhoIndicationMsgFormat1(): error validating E2SmPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}

func CreateE2SmMhoIndicationMsgFormat2(ueID *e2sm_v2_ies.Ueid, rrcStatus e2sm_mho_go.Rrcstatus) (*e2sm_mho_go.E2SmMhoIndicationMessage, error) {
	E2SmMhoPdu := e2sm_mho_go.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho_go.E2SmMhoIndicationMessage_IndicationMessageFormat2{
			IndicationMessageFormat2: &e2sm_mho_go.E2SmMhoIndicationMessageFormat2{
				UeId:      ueID,
				RrcStatus: rrcStatus,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMhoIndicationMsgFormat2(): error validating E2SmPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}

func CreateMeasurementRecordItem(cgi *e2sm_v2_ies.Cgi, rsrp *e2sm_mho_go.Rsrp) (*e2sm_mho_go.E2SmMhoMeasurementReportItem, error) {
	res := &e2sm_mho_go.E2SmMhoMeasurementReportItem{
		Cgi:  cgi,
		Rsrp: rsrp,
	}

	if err := res.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementRecordItem(): error validationg E2SmPDU %s", err)
	}

	return res, nil
}

func CreateCgiNrCGI(plmnID []byte, nrCGI *asn1.BitString) (*e2sm_v2_ies.Cgi, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("CreateCellGlobalIDNrCGI(): PlmnID should contain only 3 bytes, got %v", len(plmnID))
	}

	if nrCGI.Len != uint32(36) {
		return nil, errors.NewInvalid("CreateCellGlobalIDNrCGI(): EutraCellIdentity should be of length 36 bits, got %v", nrCGI.Len)
	}
	if len(nrCGI.Value) != 5 {
		return nil, errors.NewInvalid("CreateCellGlobalIDNrCGI(): EutraCellIdentity should be of length 36 bits (5 bytes), got %v bytes", len(nrCGI.Value))
	}

	cgi := &e2sm_v2_ies.Cgi{
		Cgi: &e2sm_v2_ies.Cgi_NRCgi{
			NRCgi: &e2sm_v2_ies.NrCgi{
				PLmnidentity: &e2sm_v2_ies.PlmnIdentity{
					Value: plmnID,
				},
				NRcellIdentity: &e2sm_v2_ies.NrcellIdentity{
					Value: nrCGI,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellGlobalIDNrCGI(): error validationg E2SmPDU %s", err)
	}

	return cgi, nil
}

func CreateCgiEutraCGI(plmnID []byte, eutraCGI *asn1.BitString) (*e2sm_v2_ies.Cgi, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("CreateCellGlobalIDEutraCGI(): PlmnID should contain only 3 bytes, got %v", len(plmnID))
	}
	if eutraCGI.Len != uint32(28) {
		return nil, errors.NewInvalid("CreateCellGlobalIDEutraCGI(): EutraCellIdentity should be of length 28 bits, got %v", eutraCGI.Len)
	}
	if len(eutraCGI.Value) != 4 {
		return nil, errors.NewInvalid("CreateCellGlobalIDEutraCGI(): EutraCellIdentity should be of length 28 bits (4 bytes), got %v bytes", len(eutraCGI.Value))
	}

	cgi := &e2sm_v2_ies.Cgi{
		Cgi: &e2sm_v2_ies.Cgi_EUtraCgi{
			EUtraCgi: &e2sm_v2_ies.EutraCgi{
				PLmnidentity: &e2sm_v2_ies.PlmnIdentity{
					Value: plmnID,
				},
				EUtracellIdentity: &e2sm_v2_ies.EutracellIdentity{
					Value: eutraCGI,
				},
			},
		},
	}

	if err := cgi.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellGlobalIDEutraCGI(): error validationg E2SmPDU %s", err)
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
