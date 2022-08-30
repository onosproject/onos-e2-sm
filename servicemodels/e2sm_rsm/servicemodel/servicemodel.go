// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/encoder"
	e2smrsm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type RsmServiceModel string

const smName = "e2sm_rsm"
const smVersion = "v1_go"
const moduleName = "e2sm_rsm_v1_go.so.2.0"
const smOIDRsmV1 = "1.3.6.1.4.1.53148.1.1.2.102"

func (sm RsmServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDRsmV1,
	}
	return smData
}

func (sm RsmServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRsmIndicationHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRsmIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrsm.E2SmRsmIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRsmIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmIndicationHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRsmIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRsmIndicationMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRsmIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrsm.E2SmRsmIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRsmIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmIndicationMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRsmIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmRanFunctionDescription(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRsmRanFunctionDescription to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRsmRanFunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrsm.E2SmRsmRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRsmRanfunctionDescription %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmRanFunctionDescription(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRsmRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRsmEventTriggerDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRsmEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrsm.E2SmRsmEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRsmEventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRsmEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on RSM")
}

func (sm RsmServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on RSM")
}

func (sm RsmServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmControlHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRsmControlHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRsmControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrsm.E2SmRsmControlHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRsmControlHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmControlHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRsmControlHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmControlMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRsmControlMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRsmControlMessage %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrsm.E2SmRsmControlMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRsmControlMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmControlMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRsmControlMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on RSM")
}

func (sm RsmServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on RSM")
}

func (sm RsmServiceModel) CallProcessIDASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on RSM")
}

func (sm RsmServiceModel) CallProcessIDProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on RSM")
}

func (sm RsmServiceModel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2smrsm.E2SmRsmRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	serviceModel := serviceModels[smOIDRsmV1]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName
	sliceNodeCapList := ranFunctionDescription.GetRicSlicingNodeCapabilityList()

	ranFunction := &topoapi.RSMRanFunction{}
	for _, sliceNodeCap := range sliceNodeCapList {
		rsmConfigItems := make([]*topoapi.RSMSupportedSlicingConfigItem, 0)
		for _, item := range sliceNodeCap.SupportedConfig {
			rsmConfig := &topoapi.RSMSupportedSlicingConfigItem{
				SlicingConfigType: topoapi.E2SmRsmCommand(item.GetSlicingConfigType()),
			}
			rsmConfigItems = append(rsmConfigItems, rsmConfig)
		}

		rsmNodeSlicingCapItem := &topoapi.RSMNodeSlicingCapabilityItem{
			MaxNumberOfSlicesDl:    sliceNodeCap.MaxNumberOfSlicesDl,
			MaxNumberOfSlicesUl:    sliceNodeCap.MaxNumberOfSlicesUl,
			SlicingType:            topoapi.RSMSlicingType(sliceNodeCap.SlicingType),
			MaxNumberOfUesPerSlice: sliceNodeCap.MaxNumberOfUesPerSlice,
			SupportedConfig:        rsmConfigItems,
		}
		ranFunction.RicSlicingNodeCapabilityList = append(ranFunction.RicSlicingNodeCapabilityList, rsmNodeSlicingCapItem)
	}

	ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	if err != nil {
		return err
	}
	serviceModel.RanFunctions = append(serviceModel.RanFunctions, ranFunctionAny)
	return nil
}
