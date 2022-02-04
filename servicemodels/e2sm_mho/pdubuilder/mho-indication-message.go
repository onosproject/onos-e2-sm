// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
)

func CreateE2SmMhoIndicationMsgFormat1(ueID *e2sm_mho.UeIdentity, rsrp *e2sm_mho.Rsrp) (*e2sm_mho.E2SmMhoIndicationMessage, error) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	measReport := []*e2sm_mho.E2SmMhoMeasurementReportItem{
		{
			Cgi: &e2sm_mho.CellGlobalId{
				CellGlobalId: &e2sm_mho.CellGlobalId_EUtraCgi{
					EUtraCgi: &e2sm_mho.Eutracgi{
						PLmnIdentity: &e2sm_mho.PlmnIdentity{
							Value: plmnIDBytes,
						},
						EUtracellIdentity: &e2sm_mho.EutracellIdentity{
							Value: &e2sm_mho.BitString{
								Value: 0x9bcd4ab, //uint64
								Len:   28,        //uint32
							},
						},
					},
				},
			},
			Rsrp: rsrp,
		},
	}

	E2SmMhoPdu := e2sm_mho.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: &e2sm_mho.E2SmMhoIndicationMessageFormat1{
				UeId:       ueID,
				MeasReport: measReport,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}

func CreateE2SmMhoIndicationMsgFormat2(ueID *e2sm_mho.UeIdentity, rrcStatus e2sm_mho.Rrcstatus) (*e2sm_mho.E2SmMhoIndicationMessage, error) {
	E2SmMhoPdu := e2sm_mho.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat2{
			IndicationMessageFormat2: &e2sm_mho.E2SmMhoIndicationMessageFormat2{
				UeId:      ueID,
				RrcStatus: rrcStatus,
			},
		},
	}

	if err := E2SmMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &E2SmMhoPdu, nil
}
