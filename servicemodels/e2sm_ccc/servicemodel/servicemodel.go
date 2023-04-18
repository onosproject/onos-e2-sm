// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/encoder"
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/protobuf/proto"
)

/*
   servicemodel package implements interface defined in:
   https://github.com/onosproject/onos-e2t/blob/2ec06f2fde15fd765497120b15661ac8fd61d6eb/pkg/modelregistry/modelregistry.go#L43-L62
   If any new top-level PDU occurs, then this template should be correspondingly adjusted and the interface (link above) should be extended.
*/

type CCCServiceModel string

const smName = "e2smcccv1"
const smVersion = "v1"
const moduleName = "e2smcccv1.so.1.0.1"
const smOID = "1.3.6.1.4.1.53148.1.1.2.103"

func (sm CCCServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOID,
	}
	return smData
}

func (sm CCCServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmCCcRIcIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmCCcRIcIndicationHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmCCcRIcIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm CCCServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smcccv1.E2SmCCcRIcIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmCCcRIcIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmCCcRIcIndicationHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmCCcRIcIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm CCCServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmCCcRIcIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmCCcRIcIndicationMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmCCcRIcIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm CCCServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smcccv1.E2SmCCcRIcIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmCCcRIcIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmCCcRIcIndicationMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmCCcRIcIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm CCCServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmCCcRAnfunctionDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmCCcRAnfunctionDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmCCcRAnfunctionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm CCCServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smcccv1.E2SmCCcRAnfunctionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmCCcRAnfunctionDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmCCcRAnfunctionDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmCCcRAnfunctionDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm CCCServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmCCcRIceventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmCCcRIceventTriggerDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmCCcRIceventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm CCCServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smcccv1.E2SmCCcRIceventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmCCcRIceventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmCCcRIceventTriggerDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmCCcRIceventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm CCCServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmCCcRIcactionDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmCCcRIcactionDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmCCcRIcactionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm CCCServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smcccv1.E2SmCCcRIcactionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmCCcRIcactionDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmCCcRIcactionDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmCCcRIcactionDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm CCCServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmCCcRIcControlHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmCCcRIcControlHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmCCcRIcControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm CCCServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smcccv1.E2SmCCcRIcControlHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmCCcRIcControlHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmCCcRIcControlHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmCCcRIcControlHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm CCCServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmCCcRIcControlMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmCCcRIcControlMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmCCcRIcControlMessage %s", err)
	}

	return protoBytes, nil
}

func (sm CCCServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smcccv1.E2SmCCcRIcControlMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmCCcRIcControlMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmCCcRIcControlMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmCCcRIcControlMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm CCCServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmCCcRIcControlOutcome(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmCCcRIcControlOutcome to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmCCcRIcControlOutcome %s", err)
	}

	return protoBytes, nil
}

func (sm CCCServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smcccv1.E2SmCCcRIcControlOutcome)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmCCcRIcControlOutcome %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmCCcRIcControlOutcome(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmCCcRIcControlOutcome to PER %s", err)
	}

	return perBytes, nil
}

func (sm CCCServiceModel) CallProcessIDASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not present on CCC (v1)")
}

func (sm CCCServiceModel) CallProcessIDProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not present on CCC (v1)")
}

// ToDo - RanFunctionDescription may vary from SM to SM and thus you'll have to adjust this function in order to extract all necessary data
func (sm CCCServiceModel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2smcccv1.E2SmCCcRAnfunctionDefinition{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	serviceModel := serviceModels[smOID]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName

	reportStyleList := ranFunctionDescription.GetE2SmCCcRAnfunctionDefinitionItem().GetRicReportStyleList()

	ranFunction := &topoapi.CCCRanFunction{}
	for _, reportStyle := range reportStyleList {
		mhoReportStyle := &topoapi.CCCReportStyle{
			Name: reportStyle.RicReportStyleName.Value,
			Type: reportStyle.RicReportStyleType.Value,
		}
		ranFunction.ReportStyles = append(ranFunction.ReportStyles, mhoReportStyle)
	}

	ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	serviceModel.RanFunctions = append(serviceModel.RanFunctions, ranFunctionAny)
	return nil
}
