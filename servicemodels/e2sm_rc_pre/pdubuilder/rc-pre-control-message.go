// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
)

func CreateE2SmRcPreControlMessage(RANparameterID int32, RANparameterName string, ranParameter *e2sm_rc_pre_v2.RanparameterValue) (*e2sm_rc_pre_v2.E2SmRcPreControlMessage, error) {

	e2smRcPreMsgFormat1 := e2sm_rc_pre_v2.E2SmRcPreControlMessageFormat1{
		ParameterType: &e2sm_rc_pre_v2.RanparameterDefItem{
			RanParameterId: &e2sm_rc_pre_v2.RanparameterId{
				Value: RANparameterID,
			},
			RanParameterName: &e2sm_rc_pre_v2.RanparameterName{
				Value: RANparameterName,
			},
		},
	}

	switch ranParameter.RanparameterValue.(type) {
	case *e2sm_rc_pre_v2.RanparameterValue_ValueInt:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_INTEGER
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2sm_rc_pre_v2.RanparameterValue_ValueEnum:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_ENUMERATED
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2sm_rc_pre_v2.RanparameterValue_ValueBool:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_BOOLEAN
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2sm_rc_pre_v2.RanparameterValue_ValueBitS:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_BIT_STRING
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2sm_rc_pre_v2.RanparameterValue_ValueOctS:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_OCTET_STRING
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2sm_rc_pre_v2.RanparameterValue_ValuePrtS:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_PRINTABLE_STRING
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	default:
		return nil, fmt.Errorf("CreateE2SmRcPreControlMessage() unknown type of RANparameterValue %T", ranParameter)
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

func CreateRanParameterValueInt(value uint32) (*e2sm_rc_pre_v2.RanparameterValue, error) {
	if value > 4294967295 {
		return nil, fmt.Errorf("Value should be within range 0 to 4294967295")
	}
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueInt{
			ValueInt: value,
		},
	}, nil
}

func CreateRanParameterValueEnum(value int32) *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueEnum{
			ValueEnum: value,
		},
	}
}

func CreateRanParameterValueBool(value bool) *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueBool{
			ValueBool: value,
		},
	}
}

func CreateRanParameterValueBitS(value *e2sm_rc_pre_v2.BitString) *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueBitS{
			ValueBitS: value,
		},
	}
}

func CreateRanParameterValueOctS(value string) *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueOctS{
			ValueOctS: value,
		},
	}
}

func CreateRanParameterValuePrtS(value string) *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValuePrtS{
			ValuePrtS: value,
		},
	}
}
