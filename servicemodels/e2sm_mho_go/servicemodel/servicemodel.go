// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package servicemodel

import (
	"encoding/hex"
	"fmt"
	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
	"google.golang.org/protobuf/proto"
)

type MhoServiceModel string

const smName = "e2sm_mho_go"
const smVersion = "v1"
const moduleName = "e2sm_mho_go.so.1.0.1"
const smOIDMho = "1.3.6.1.4.1.53148.1.1.2.101"

func (sm MhoServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDMho,
	}
	return smData
}

func (sm MhoServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoIndicationHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm MhoServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho_go.E2SmMhoIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoIndicationHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm MhoServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoIndicationMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm MhoServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho_go.E2SmMhoIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoIndicationMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm MhoServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoRanFunctionDescription(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoRanFunctionDescription to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoRanFunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm MhoServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho_go.E2SmMhoRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoRanFunctionDescription %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoRanFunctionDescription(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoRanFunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm MhoServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoEventTriggerDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm MhoServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho_go.E2SmMhoEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoEventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm MhoServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on MHO")
}

func (sm MhoServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on MHO")
}

func (sm MhoServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoControlHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoControlHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm MhoServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho_go.E2SmMhoControlHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoControlHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoControlHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoControlHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm MhoServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoControlMessage(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmMhoControlMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmMhoControlMessage %s", err)
	}

	return protoBytes, nil
}

func (sm MhoServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_mho_go.E2SmMhoControlMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmMhoControlMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoControlMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmMhoControlMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm MhoServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (sm MhoServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (sm MhoServiceModel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2sm_mho_go.E2SmMhoRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	serviceModel := serviceModels[smOIDMho]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName
	reportStyleList := ranFunctionDescription.GetE2SmMhoRanfunctionItem().GetRicReportStyleList()

	ranFunction := &topoapi.MHORanFunction{}
	for _, reportStyle := range reportStyleList {
		mhoReportStyle := &topoapi.MHOReportStyle{
			Name: reportStyle.RicReportStyleName.Value,
			Type: reportStyle.RicReportStyleType.Value,
		}
		ranFunction.ReportStyles = append(ranFunction.ReportStyles, mhoReportStyle)
	}

	ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	serviceModel.RanFunctions = append(serviceModel.RanFunctions, ranFunctionAny)
	return nil
}
