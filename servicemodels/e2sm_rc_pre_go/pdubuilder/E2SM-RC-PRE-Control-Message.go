// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcPreControlMessage(RANparameterID int32, RANparameterName string, ranParameter *e2smrcprev2.RanparameterValue) (*e2smrcprev2.E2SmRcPreControlMessage, error) {

	e2smRcPreMsgFormat1 := e2smrcprev2.E2SmRcPreControlMessageFormat1{
		ParameterType: &e2smrcprev2.RanparameterDefItem{
			RanParameterId: &e2smrcprev2.RanparameterId{
				Value: RANparameterID,
			},
			RanParameterName: &e2smrcprev2.RanparameterName{
				Value: RANparameterName,
			},
		},
	}

	switch ranParameter.RanparameterValue.(type) {
	case *e2smrcprev2.RanparameterValue_ValueInt:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2smrcprev2.RanparameterType_RANPARAMETER_TYPE_INTEGER
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2smrcprev2.RanparameterValue_ValueEnum:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2smrcprev2.RanparameterType_RANPARAMETER_TYPE_ENUMERATED
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2smrcprev2.RanparameterValue_ValueBool:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2smrcprev2.RanparameterType_RANPARAMETER_TYPE_BOOLEAN
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2smrcprev2.RanparameterValue_ValueBitS:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2smrcprev2.RanparameterType_RANPARAMETER_TYPE_BIT_STRING
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2smrcprev2.RanparameterValue_ValueOctS:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2smrcprev2.RanparameterType_RANPARAMETER_TYPE_OCTET_STRING
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	case *e2smrcprev2.RanparameterValue_ValuePrtS:
		e2smRcPreMsgFormat1.ParameterType.RanParameterType = e2smrcprev2.RanparameterType_RANPARAMETER_TYPE_PRINTABLE_STRING
		e2smRcPreMsgFormat1.ParameterVal = ranParameter

	default:
		return nil, errors.NewInvalid("CreateE2SmRcPreControlMessage() unknown type of RANparameterValue %T", ranParameter)
	}

	e2smRcPrePdu := e2smrcprev2.E2SmRcPreControlMessage{
		E2SmRcPreControlMessage: &e2smrcprev2.E2SmRcPreControlMessage_ControlMessage{
			ControlMessage: &e2smRcPreMsgFormat1,
		},
	}

	if err := e2smRcPrePdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmPDU %s", err.Error())
	}
	return &e2smRcPrePdu, nil
}

func CreateRanParameterValueInt(value int64) (*e2smrcprev2.RanparameterValue, error) {
	if value > 4294967295 || value < 0 {
		return nil, errors.NewInvalid("Value should be within range 0 to 4294967295")
	}
	return &e2smrcprev2.RanparameterValue{
		RanparameterValue: &e2smrcprev2.RanparameterValue_ValueInt{
			ValueInt: value,
		},
	}, nil
}

func CreateRanParameterValueEnum(value int32) *e2smrcprev2.RanparameterValue {
	return &e2smrcprev2.RanparameterValue{
		RanparameterValue: &e2smrcprev2.RanparameterValue_ValueEnum{
			ValueEnum: value,
		},
	}
}

func CreateRanParameterValueBool(value bool) *e2smrcprev2.RanparameterValue {
	return &e2smrcprev2.RanparameterValue{
		RanparameterValue: &e2smrcprev2.RanparameterValue_ValueBool{
			ValueBool: value,
		},
	}
}

func CreateRanParameterValueBitS(value *asn1.BitString) *e2smrcprev2.RanparameterValue {
	return &e2smrcprev2.RanparameterValue{
		RanparameterValue: &e2smrcprev2.RanparameterValue_ValueBitS{
			ValueBitS: value,
		},
	}
}

func CreateRanParameterValueOctS(value string) *e2smrcprev2.RanparameterValue {
	return &e2smrcprev2.RanparameterValue{
		RanparameterValue: &e2smrcprev2.RanparameterValue_ValueOctS{
			ValueOctS: value,
		},
	}
}

func CreateRanParameterValuePrtS(value string) *e2smrcprev2.RanparameterValue {
	return &e2smrcprev2.RanparameterValue{
		RanparameterValue: &e2smrcprev2.RanparameterValue_ValuePrtS{
			ValuePrtS: value,
		},
	}
}
