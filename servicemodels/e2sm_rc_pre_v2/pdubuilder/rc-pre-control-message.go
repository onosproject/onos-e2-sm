// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_v2/v2/e2sm-rc-pre-v2"
)

func CreateE2SmRcPreControlMessage(RANparameterID int32, RANparameterName string, RANparameterValue int32) (*e2sm_rc_pre_v2.E2SmRcPreControlMessage, error) {

	e2smRcPreMsgFormat1 := e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1{
		ParameterType: &e2sm_rc_pre_v2.RanparameterDefItem{
			RanParameterId: &e2sm_rc_pre_v2.RanparameterId{
				Value: RANparameterID,
			},
			RanParameterName: &e2sm_rc_pre_v2.RanparameterName{
				Value: RANparameterName,
			},
			RanParameterType: e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_INTEGER,
		},
		ParameterVal: &e2sm_rc_pre_v2.RanparameterValue{
			RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueInt{
				ValueInt: RANparameterValue,
			},
		},
	}
	e2smRcPrePdu := e2sm_rc_pre_v2.E2SmRcPreControlMessage{
		E2SmRcPreControlMessage: &e2sm_rc_pre_v2.E2SmRcPreControlMessage_ControlMessage{
			ControlMessage: &e2smRcPreMsgFormat1,
		},
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}
