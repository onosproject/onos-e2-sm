// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	{{.Imports}} "github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/protobuf/proto"
)

/*
    servicemodel package implements interface defined in:
    https://github.com/onosproject/onos-e2t/blob/2ec06f2fde15fd765497120b15661ac8fd61d6eb/pkg/modelregistry/modelregistry.go#L43-L62
    If any new top-level PDU occurs, then this template should be correspondingly adjusted and the interface (link above) should be extended.
*/

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

	return protoBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
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

	return perBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.IndicationMessage.IsPresent}}
	perBytes, err := encoder.PerDecode{{.TopLevelPdus.IndicationMessage.MessageProtoName}}(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding {{.TopLevelPdus.IndicationMessage.MessageProtoName}} to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to {{.TopLevelPdus.IndicationMessage.MessageProtoName}} %s", err)
	}

	return protoBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	{{if .TopLevelPdus.IndicationHeader.IsPresent}}
	protoObj := new({{.SmName}}.{{.TopLevelPdus.IndicationMessage.MessageProtoName}})
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to {{.TopLevelPdus.IndicationMessage.MessageProtoName}} %s", err)
	}

	perBytes, err := encoder.PerEncode{{.TopLevelPdus.IndicationMessage.MessageProtoName}}(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding {{.TopLevelPdus.IndicationMessage.MessageProtoName}} to PER %s", err)
	}

	return perBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.RanFunctionDescription.IsPresent}}
	perBytes, err := encoder.PerDecode{{.TopLevelPdus.RanFunctionDescription.MessageProtoName}}(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding {{.TopLevelPdus.RanFunctionDescription.MessageProtoName}} to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to {{.TopLevelPdus.RanFunctionDescription.MessageProtoName}} %s", err)
	}

	return protoBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.RanFunctionDescription.IsPresent}}
	protoObj := new({{.SmName}}.{{.TopLevelPdus.RanFunctionDescription.MessageProtoName}})
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to {{.TopLevelPdus.RanFunctionDescription.MessageProtoName}} %s", err)
	}

	perBytes, err := encoder.PerEncode{{.TopLevelPdus.RanFunctionDescription.MessageProtoName}}(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding {{.TopLevelPdus.RanFunctionDescription.MessageProtoName}} to PER %s", err)
	}

	return perBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.EventTriggerDefinition.IsPresent}}
	perBytes, err := encoder.PerDecode{{.TopLevelPdus.EventTriggerDefinition.MessageProtoName}}(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding {{.TopLevelPdus.EventTriggerDefinition.MessageProtoName}} to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to {{.TopLevelPdus.EventTriggerDefinition.MessageProtoName}} %s", err)
	}

	return protoBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.EventTriggerDefinition.IsPresent}}
	protoObj := new({{.SmName}}.{{.TopLevelPdus.EventTriggerDefinition.MessageProtoName}})
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to {{.TopLevelPdus.EventTriggerDefinition.MessageProtoName}} %s", err)
	}

	perBytes, err := encoder.PerEncode{{.TopLevelPdus.EventTriggerDefinition.MessageProtoName}}(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding {{.TopLevelPdus.EventTriggerDefinition.MessageProtoName}} to PER %s", err)
	}

	return perBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.ActionDefinition.IsPresent}}
	perBytes, err := encoder.PerDecode{{.TopLevelPdus.ActionDefinition.MessageProtoName}}(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding {{.TopLevelPdus.ActionDefinition.MessageProtoName}} to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to {{.TopLevelPdus.ActionDefinition.MessageProtoName}} %s", err)
	}

	return protoBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm MhoServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.ActionDefinition.IsPresent}}
    protoObj := new({{.SmName}}.{{.TopLevelPdus.ActionDefinition.MessageProtoName}})
    if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
        return nil, errors.NewInvalid("error unmarshalling protoBytes to {{.TopLevelPdus.ActionDefinition.MessageProtoName}} %s", err)
    }

    perBytes, err := encoder.PerEncode{{.TopLevelPdus.ActionDefinition.MessageProtoName}}(protoObj)
    if err != nil {
        return nil, errors.NewInvalid("error encoding {{.TopLevelPdus.ActionDefinition.MessageProtoName}} to PER %s", err)
    }

    return perBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.ControlHeader.IsPresent}}
	perBytes, err := encoder.PerDecode{{.TopLevelPdus.ControlHeader.MessageProtoName}}(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding {{.TopLevelPdus.ControlHeader.MessageProtoName}} to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to {{.TopLevelPdus.ControlHeader.MessageProtoName}} %s", err)
	}

	return protoBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.ControlHeader.IsPresent}}
	protoObj := new({{.SmName}}.{{.TopLevelPdus.ControlHeader.MessageProtoName}})
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to {{.TopLevelPdus.ControlHeader.MessageProtoName}} %s", err)
	}

	perBytes, err := encoder.PerEncode{{.TopLevelPdus.ControlHeader.MessageProtoName}}(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding {{.TopLevelPdus.ControlHeader.MessageProtoName}} to PER %s", err)
	}

	return perBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.ControlMessage.IsPresent}}
	perBytes, err := encoder.PerDecode{{.TopLevelPdus.ControlMessage.MessageProtoName}}(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding {{.TopLevelPdus.ControlMessage.MessageProtoName}} to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to {{.TopLevelPdus.ControlMessage.MessageProtoName}} %s", err)
	}

	return protoBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.ControlMessage.IsPresent}}
	protoObj := new({{.SmName}}.{{.TopLevelPdus.ControlMessage.MessageProtoName}})
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to {{.TopLevelPdus.ControlMessage.MessageProtoName}} %s", err)
	}

	perBytes, err := encoder.PerEncode{{.TopLevelPdus.ControlMessage.MessageProtoName}}(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding {{.TopLevelPdus.ControlMessage.MessageProtoName}} to PER %s", err)
	}

	return perBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.ControlOutcome.IsPresent}}
	perBytes, err := encoder.PerDecode{{.TopLevelPdus.ControlOutcome.MessageProtoName}}(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding {{.TopLevelPdus.ControlOutcome.MessageProtoName}} to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to {{.TopLevelPdus.ControlOutcome.MessageProtoName}} %s", err)
	}

	return protoBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.ControlOutcome.IsPresent}}
	protoObj := new({{.SmName}}.{{.TopLevelPdus.ControlOutcome.MessageProtoName}})
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to {{.TopLevelPdus.ControlOutcome.MessageProtoName}} %s", err)
	}

	perBytes, err := encoder.PerEncode{{.TopLevelPdus.ControlOutcome.MessageProtoName}}(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding {{.TopLevelPdus.ControlOutcome.MessageProtoName}} to PER %s", err)
	}

	return perBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) CallProcessIDASN1toProto(asn1Bytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.CallProcessID.IsPresent}}
	perBytes, err := encoder.PerDecode{{.TopLevelPdus.CallProcessID.MessageProtoName}}(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding {{.TopLevelPdus.CallProcessID.MessageProtoName}} to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to {{.TopLevelPdus.CallProcessID.MessageProtoName}} %s", err)
	}

	return protoBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

func (sm {{.PackageName}}) CallProcessIDProtoToASN1(protoBytes []byte) ([]byte, error) {
    {{if .TopLevelPdus.CallProcessID.IsPresent}}
	protoObj := new({{.SmName}}.{{.TopLevelPdus.CallProcessID.MessageProtoName}})
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to {{.TopLevelPdus.CallProcessID.MessageProtoName}} %s", err)
	}

	perBytes, err := encoder.PerEncode{{.TopLevelPdus.CallProcessID.MessageProtoName}}(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding {{.TopLevelPdus.CallProcessID.MessageProtoName}} to PER %s", err)
	}

	return perBytes, nil{{else}}return nil, errors.NewInvalid("not present on {{.E2SmName}} ({{.SmVersion}})"){{end}}
}

// ToDo - RanFunctionDescription may vary from SM to SM and thus you'll have to adjust this function in order to extract all necessary data
func (sm {{.PackageName}}) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &{{.SmName}}.{{.TopLevelPdus.RanFunctionDescription.MessageProtoName}}{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	serviceModel := serviceModels[smOID]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName

	reportStyleList := ranFunctionDescription.Get{{.TopLevelPdus.RanFunctionDescription.MessageProtoName}}Item().GetRicReportStyleList()

	ranFunction := &topoapi.{{.E2SmName}}RanFunction{}
	for _, reportStyle := range reportStyleList {
		mhoReportStyle := &topoapi.{{.E2SmName}}ReportStyle{
			Name: reportStyle.RicReportStyleName.Value,
			Type: reportStyle.RicReportStyleType.Value,
		}
		ranFunction.ReportStyles = append(ranFunction.ReportStyles, mhoReportStyle)
	}

	ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	serviceModel.RanFunctions = append(serviceModel.RanFunctions, ranFunctionAny)
	return nil
}
