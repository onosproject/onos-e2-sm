// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//nolint
package main

import (
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	kpmv2ctypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/kpmctypes"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdudecoder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"google.golang.org/protobuf/proto"
)

type serviceModel struct{}

const smName = "e2sm_kpm"
const smVersion = "v2"
const moduleName = "e2sm_kpm_v2.so.2.0"
const smOIDKpmV2 = "1.3.6.1.4.1.53148.1.2.2.2"

func (sm serviceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDKpmV2,
	}
	return smData
}

func (sm serviceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmIndicationHeader to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm serviceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2.E2SmKpmIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmIndicationHeader %s", err)
	}

	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmIndicationHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm serviceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmIndicationMessage to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm serviceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2.E2SmKpmIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmIndicationMessage %s", err)
	}

	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmIndicationMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm serviceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmRanfunctionDescription(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmRanfunctionDescription to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmRanfunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm serviceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2.E2SmKpmRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmRanfunctionDescription %s", err)
	}

	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmRanfunctionDescription(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm serviceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmEventTriggerDefinition to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm serviceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2.E2SmKpmEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmEventTriggerDefinition %s", err)
	}

	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm serviceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := kpmv2ctypes.PerDecodeE2SmKpmActionDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmActionDefinitio to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmActionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm serviceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2.E2SmKpmActionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmActionDefinition %s", err)
	}

	perBytes, err := kpmv2ctypes.PerEncodeE2SmKpmActionDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmActionDefinition to PER %s", err)
	}

	return perBytes, nil
}

//It is redundant so far - could be reused for future, if you need to extract something specific from RanFunctionDescription message
func (sm serviceModel) DecodeRanFunctionDescription(asn1bytes []byte) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error) {
	e2SmKpmPdu, err := kpmv2ctypes.PerDecodeE2SmKpmRanfunctionDescription(asn1bytes)
	if err != nil {
		return nil, nil, nil, err
	}
	return pdudecoder.DecodeE2SmKpmRanfunctionDescription(e2SmKpmPdu)
}

func (sm serviceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm serviceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm serviceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm serviceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm serviceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm serviceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm serviceModel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2sm_kpm_v2.E2SmKpmRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	e2Cells := request.E2Cells
	kpmV2ServiceModel := serviceModels[smOIDKpmV2]
	kpmV2ServiceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName
	kpmRANFunction := &topoapi.KPMRanFunction{}

	for _, kpmNode := range ranFunctionDescription.RicKpmNodeList {
		for _, cell := range kpmNode.CellMeasurementObjectList {
			cellObject := &topoapi.E2Cell{
				CID: cell.GetCellObjectId().GetValue(),
			}
			*e2Cells = append(*e2Cells, cellObject)
		}
	}

	for _, reportStyle := range ranFunctionDescription.GetRicReportStyleList() {
		for _, meanInfoItem := range reportStyle.GetMeasInfoActionList().GetValue() {
			kpmRANFunction.Measurements = append(kpmRANFunction.Measurements, &topoapi.KPMMeasurement{
				ID:   meanInfoItem.GetMeasId().String(),
				Name: meanInfoItem.GetMeasName().GetValue(),
			})

		}
	}
	kpmRANFunctionAny, err := gogotypes.MarshalAny(kpmRANFunction)
	kpmV2ServiceModel.RanFunctions = []*gogotypes.Any{kpmRANFunctionAny}

	return nil
}

//ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel serviceModel
