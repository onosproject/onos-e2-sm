// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	e2smmhov2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	{{.Imports}}
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type {{.PackageName}} string

const smName = "{{.SmName}}"
const smVersion = "{{.SmVersion}}"
const moduleName = "{{.SmName}}.so.1.0.1"
const smOID = "{{.OID}}"

func (sm {{.PackageName}}) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOID,
	}
	return smData
}

func (sm {{.PackageName}}) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	{{if .TopLevelPdus.IndicationHeader.IsPresent}}
	perBytes, err := encoder.PerDecode{{.TopLevelPdus.IndicationHeader.MessageProtoName}}(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding {{.TopLevelPdus.IndicationHeader.MessageProtoName}} to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to {{.TopLevelPdus.IndicationHeader.MessageProtoName}} %s", err)
	}

	return protoBytes, nil
	{{else}}
	return nil, errors.NewInvalid("not implemented")
	{{end}}
}

func (sm {{.PackageName}}) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	{{if .TopLevelPdus.IndicationHeader.IsPresent}}
	protoObj := new({{.SmName}}.{{.TopLevelPdus.IndicationHeader.MessageProtoName}})
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to {{.TopLevelPdus.IndicationHeader.MessageProtoName}} %s", err)
	}

	perBytes, err := encoder.PerEncode{{.TopLevelPdus.IndicationHeader.MessageProtoName}}(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding {{.TopLevelPdus.IndicationHeader.MessageProtoName}} to PER %s", err)
	}

	return perBytes, nil
    {{else}}
    return nil, errors.NewInvalid("not implemented")
    {{end}}
}

func (sm {{.PackageName}}) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.IndicationHeader.IsPresent}}
	perBytes, err := encoder.PerDecodeE2SmMhoIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMhoIndicationMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMhoIndicationMessage %s", err)
	}

	return protoBytes, nil
    {{else}}
    return nil, errors.NewInvalid("not implemented")
    {{end}}
}

func (sm {{.PackageName}}) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmhov2.E2SmMhoIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMhoIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoIndicationMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMhoIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm {{.PackageName}}) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoRanFunctionDescription(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMhoRanFunctionDescription to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMhoRanFunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm {{.PackageName}}) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmhov2.E2SmMhoRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMhoRanFunctionDescription %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoRanFunctionDescription(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMhoRanFunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm {{.PackageName}}) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMhoEventTriggerDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMhoEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm {{.PackageName}}) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmhov2.E2SmMhoEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMhoEventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMhoEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm {{.PackageName}}) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MHO")
}

func (sm MhoServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MHO")
}

func (sm {{.PackageName}}) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoControlHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMhoControlHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMhoControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm {{.PackageName}}) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmhov2.E2SmMhoControlHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMhoControlHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoControlHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMhoControlHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm {{.PackageName}}) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMhoControlMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMhoControlMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMhoControlMessage %s", err)
	}

	return protoBytes, nil
}

func (sm {{.PackageName}}) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmhov2.E2SmMhoControlMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMhoControlMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMhoControlMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMhoControlMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm {{.PackageName}}) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented")
}

func (sm {{.PackageName}}) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented")
}

func (sm {{.PackageName}}) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2smmhov2.E2SmMhoRanfunctionDescription{}
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
