// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
)

func CreateE2SmMhoControlMessage(servingCgi *e2sm_mho_go.CellGlobalId, uedID *e2sm_mho_go.UeIdentity, targetCgi *e2sm_mho_go.CellGlobalId) (*e2sm_mho_go.E2SmMhoControlMessage, error) {

	e2smMhoMsgFormat1 := e2sm_mho_go.E2SmMhoControlMessageFormat1{
		ServingCgi: servingCgi,
		UedId:      uedID,
		TargetCgi:  targetCgi,
	}

	e2smMhoPdu := e2sm_mho_go.E2SmMhoControlMessage{
		E2SmMhoControlMessage: &e2sm_mho_go.E2SmMhoControlMessage_ControlMessageFormat1{
			ControlMessageFormat1: &e2smMhoMsgFormat1,
		},
	}

	//if err := e2smMhoPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	//}
	return &e2smMhoPdu, nil
}
