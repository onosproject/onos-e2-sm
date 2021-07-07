// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//nolint
package main

import (
	"fmt"

	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/mhoctypes"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdudecoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"google.golang.org/protobuf/proto"
)

type servicemodel string

const smName = "e2sm_mho"
const smVersion = "v1"
const moduleName = "e2sm_mho.so.1.0.1"
const smOIDMho = "1.3.6.1.4.1.53148.1.1.2.101"

func (sm servicemodel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDMho,
	}
	return smData
}

func (sm servicemodel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := mhoctypes.PerDecodeE2SmMhoIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoIndicationHeader to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho.E2SmMhoIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoIndicationHeader %s", err)
	}

	perBytes, err := mhoctypes.PerEncodeE2SmMhoIndicationHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := mhoctypes.PerDecodeE2SmMhoIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoIndicationMessage to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho.E2SmMhoIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoIndicationMessage %s", err)
	}

	perBytes, err := mhoctypes.PerEncodeE2SmMhoIndicationMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := mhoctypes.PerDecodeE2SmMhoRanfunctionDescription(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoRanfunctionDescription to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoRanfunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho.E2SmMhoRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoRanfunctionDescription %s", err)
	}

	perBytes, err := mhoctypes.PerEncodeE2SmMhoRanfunctionDescription(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := mhoctypes.PerDecodeE2SmMhoEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoEventTriggerDefinition to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho.E2SmMhoEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoEventTriggerDefinition %s", err)
	}

	perBytes, err := mhoctypes.PerEncodeE2SmMhoEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on MHO")
}

func (sm servicemodel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on MHO")
}

func (sm servicemodel) DecodeRanFunctionDescription(asn1bytes []byte) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error) {
	e2SmMhoPdu, err := mhoctypes.PerDecodeE2SmMhoRanfunctionDescription(asn1bytes)
	if err != nil {
		return nil, nil, nil, err
	}
	return pdudecoder.DecodeE2SmMhoRanfunctionDescription(e2SmMhoPdu)
}

func (sm servicemodel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := mhoctypes.PerDecodeE2SmMhoControlHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoControlHeader to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho.E2SmMhoControlHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoControlHeader %s", err)
	}

	perBytes, err := mhoctypes.PerEncodeE2SmMhoControlHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoControlHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := mhoctypes.PerDecodeE2SmMhoControlMessage(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoControlMessage to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoControlMessage %s", err)
	}

	return protoBytes, nil
}

func (sm servicemodel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho.E2SmMhoControlMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoControlMessage %s", err)
	}

	perBytes, err := mhoctypes.PerEncodeE2SmMhoControlMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoControlMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (sm servicemodel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (sm servicemodel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2sm_mho.E2SmMhoRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	serviceModel := serviceModels[smOIDMho]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName
	reportStyleList := ranFunctionDescription.GetRicReportStyleList()

	ranFunction := &topoapi.MHORanFunction{}
	for _, reportStyle := range reportStyleList {
		mhoReportStyle := &topoapi.MHOReportStyle{
			Name: reportStyle.RicReportStyleName.Value,
			Type: reportStyle.RicReportStyleType.Value,
		}
		ranFunction.ReportStyles = append(ranFunction.ReportStyles, mhoReportStyle)
	}

	ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	serviceModel.RanFunctions = []*prototypes.Any{ranFunctionAny}
	return nil
}

// ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel
