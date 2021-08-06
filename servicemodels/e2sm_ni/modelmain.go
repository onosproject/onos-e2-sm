// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//nolint
package main

import (
	"fmt"

	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
)

type servicemodel string

const smName = "e2sm_ni"
const smVersion = "v1beta1"
const moduleName = "e2sm_ni.so.1.0.1"
const smOIDNiV1 = "1.3.6.1.4.1.53148.1.1.2.1"

func (sm servicemodel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDNiV1,
	}
	return smData
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

func (sm servicemodel) DecodeRanFunctionDescription(asn1bytes []byte) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error) {
	return nil, nil, nil, fmt.Errorf("not yet implemented")
}

// ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel
