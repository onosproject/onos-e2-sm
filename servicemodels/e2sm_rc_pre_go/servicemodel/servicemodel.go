// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/encoder"
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type RcPreServiceModel string

const smName = "e2sm_rc_pre"
const smVersion = "v2_go"
const moduleName = "e2sm_rc_pre_v2_go.so.2.0"
const smOIDRcPreV1 = "1.3.6.1.4.1.53148.1.2.2.100"

func (sm RcPreServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDRcPreV1,
	}
	return smData
}

func (sm RcPreServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRcPreIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcPreIndicationHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcPreIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm RcPreServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrcprev2.E2SmRcPreIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcPreIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcPreIndicationHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcPreIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm RcPreServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRcPreIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcPreIndicationMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcPreIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm RcPreServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrcprev2.E2SmRcPreIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcPreIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcPreIndicationMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcPreIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm RcPreServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRcPreRanFunctionDescription(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcPreRanFunctionDescription to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcPreRanFunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm RcPreServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrcprev2.E2SmRcPreRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcPreRanfunctionDescription %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcPreRanFunctionDescription(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcPreRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm RcPreServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRcPreEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcPreEventTriggerDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcPreEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm RcPreServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrcprev2.E2SmRcPreEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcPreEventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcPreEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcPreEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm RcPreServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on RC-PRE")
}

func (sm RcPreServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on RC-PRE")
}

func (sm RcPreServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRcPreControlHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcPreControlHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcPreControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm RcPreServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrcprev2.E2SmRcPreControlHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcPreControlHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcPreControlHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcPreControlHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm RcPreServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRcPreControlMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcPreControlMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcPreControlMessage %s", err)
	}

	return protoBytes, nil
}

func (sm RcPreServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrcprev2.E2SmRcPreControlMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcPreControlMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcPreControlMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcPreControlMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm RcPreServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRcPreControlOutcome(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcPreControlOutcome to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcPreControlOutcome %s", err)
	}

	return protoBytes, nil
}

func (sm RcPreServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrcprev2.E2SmRcPreControlOutcome)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcPreControlOutcome %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcPreControlOutcome(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcPreControlOutcome to PER %s", err)
	}

	return perBytes, nil
}

func (sm RcPreServiceModel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2smrcprev2.E2SmRcPreRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	serviceModel := serviceModels[smOIDRcPreV1]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName
	reportStyleList := ranFunctionDescription.GetE2SmRcPreRanfunctionItem().GetRicReportStyleList()

	ranFunction := &topoapi.RCRanFunction{}
	for _, reportStyle := range reportStyleList {
		rcReportStyle := &topoapi.RCReportStyle{
			Name: reportStyle.RicReportStyleName.Value,
			Type: reportStyle.RicReportStyleType.Value,
		}
		ranFunction.ReportStyles = append(ranFunction.ReportStyles, rcReportStyle)
	}

	ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	if err != nil {
		return err
	}
	serviceModel.RanFunctions = []*prototypes.Any{ranFunctionAny}
	return nil
}
