// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//nolint
package main

import (
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"google.golang.org/protobuf/proto"
)

type servicemodel string

const smname = "e2sm_rc_pre"
const smversion = "v1"
const modulename = "e2sm_rc_pre.so.1.0.1"

func (sm servicemodel) ServiceModelData() (string, string, string) {
	return smname, smversion, modulename
}

func (sm servicemodel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := rcprectypes.PerDecodeE2SmRcPreIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRcPreIndicationHeader to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRcPreIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rc_pre_ies.E2SmRcPreIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRcPreIndicationHeader %s", err)
	}

	perBytes, err := rcprectypes.PerEncodeE2SmRcPreIndicationHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRcPreIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := rcprectypes.PerDecodeE2SmRcPreIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRcPreIndicationMessage to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRcPreIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rc_pre_ies.E2SmRcPreIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRcPreIndicationMessage %s", err)
	}

	perBytes, err := rcprectypes.PerEncodeE2SmRcPreIndicationMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRcPreIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := rcprectypes.PerDecodeE2SmRcPreEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRcPreEventTriggerDefinition to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRcPreEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRcPreEventTriggerDefinition %s", err)
	}

	perBytes, err := rcprectypes.PerEncodeE2SmRcPreEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRcPreEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := rcprectypes.PerDecodeE2SmRcPreActionDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRcPreActionDefinitio to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRcPreActionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rc_pre_ies.E2SmRcPreActionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRcPreActionDefinition %s", err)
	}

	perBytes, err := rcprectypes.PerEncodeE2SmRcPreActionDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRcPreActionDefinition to PER %s", err)
	}

	return perBytes, nil
}

// ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel
