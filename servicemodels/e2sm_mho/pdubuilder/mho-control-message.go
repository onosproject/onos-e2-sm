// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
)

func CreateE2SmMhoControlMessage(servingCgi *e2sm_mho.CellGlobalId, uedID *e2sm_mho.UeIdentity, targetCgi *e2sm_mho.CellGlobalId) (*e2sm_mho.E2SmMhoControlMessage, error) {

	e2smMhoMsgFormat1 := e2sm_mho.E2SmMhoControlMessageFormat1{
		ServingCgi: servingCgi,
		UedId:      uedID,
		TargetCgi:  targetCgi,
	}

	e2smMhoPdu := e2sm_mho.E2SmMhoControlMessage{
		E2SmMhoControlMessage: &e2sm_mho.E2SmMhoControlMessage_ControlMessageFormat1{
			ControlMessageFormat1: &e2smMhoMsgFormat1,
		},
	}

	if err := e2smMhoPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &e2smMhoPdu, nil
}
