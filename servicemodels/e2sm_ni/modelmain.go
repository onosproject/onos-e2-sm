// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//nolint
package main

import (
	"fmt"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
)

type servicemodel string

const smname = "e2sm_ni"
const smversion = "v1beta1"
const modulename = "e2sm_ni.so.1.0.1"

func (sm servicemodel) ServiceModelData() (string, string, string) {
	return smname, smversion, modulename
}

func (sm servicemodel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}
func (sm servicemodel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}
func (sm servicemodel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}
func (sm servicemodel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}
func (sm servicemodel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) DecodeRanFunctionDescription(asn1bytes []byte) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error){
	return nil, nil, nil, fmt.Errorf("not yet implemented")
}

// ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel
