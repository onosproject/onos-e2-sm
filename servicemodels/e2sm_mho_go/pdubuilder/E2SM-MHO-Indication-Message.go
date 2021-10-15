// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmMhoIndicationMsgFormat1(ueID *e2sm_mho_go.UeIdentity, rsrp *e2sm_mho_go.Rsrp) (*e2sm_mho_go.E2SmMhoIndicationMessage, error) {

	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)

	if err != nil {
		return nil, err
	}

	measReport := []*e2sm_mho_go.E2SmMhoMeasurementReportItem{
		{
			Cgi: &e2sm_mho_go.CellGlobalId{
				CellGlobalId: &e2sm_mho_go.CellGlobalId_EUtraCgi{
					EUtraCgi: &e2sm_mho_go.Eutracgi{
						PLmnIdentity: &e2sm_mho_go.PlmnIdentity{
							Value: plmnIDBytes,
						},
						EUtracellIdentity: &e2sm_mho_go.EutracellIdentity{
							Value: &asn1.BitString{
								Value: []byte{0x9b, 0xcd, 0x4a, 0xb0},
								Len:   28, //uint32
							},
						},
					},
				},
			},
			Rsrp: rsrp,
		},
	}

	E2SmMhoPdu := e2sm_mho_go.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho_go.E2SmMhoIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: &e2sm_mho_go.E2SmMhoIndicationMessageFormat1{
				UeId:       ueID,
				MeasReport: measReport,
			},
		},
	}

	//if err := E2SmMhoPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	//}
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

	//if err := E2SmMhoPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	//}
	return &E2SmMhoPdu, nil
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
