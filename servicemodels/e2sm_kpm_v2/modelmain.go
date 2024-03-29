// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// nolint
package main

import (
	"fmt"

	prototypes "github.com/gogo/protobuf/types"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	kpmv2ctypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/kpmctypes"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"google.golang.org/protobuf/proto"
)

type servicemodel string

const smName = "e2sm_kpm"
const smVersion = "v2"
const moduleName = "e2sm_kpm_v2.so.2.0"
const smOIDKpmV2 = "1.3.6.1.4.1.53148.1.2.2.2"

func (sm servicemodel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDKpmV2,
	}
	return smData
}

func (sm servicemodel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
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

func (sm servicemodel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
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

func (sm servicemodel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
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

func (sm servicemodel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
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

func (sm servicemodel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
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

func (sm servicemodel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
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

func (sm servicemodel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
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

func (sm servicemodel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
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

func (sm servicemodel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
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

func (sm servicemodel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
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

func (sm servicemodel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm servicemodel) OnSetup(request *types.OnSetupRequest) error {
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
	serviceModel := serviceModels[smOIDKpmV2]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName
	ranFunction := &topoapi.KPMRanFunction{}

	for _, kpmNode := range ranFunctionDescription.RicKpmNodeList {
		for _, cell := range kpmNode.CellMeasurementObjectList {
			cellObject := &topoapi.E2Cell{
				CellObjectID: cell.GetCellObjectId().GetValue(),
				CellGlobalID: &topoapi.CellGlobalID{},
			}
			switch cellGlobalID := cell.CellGlobalId.GetCellGlobalId().(type) {
			case *e2sm_kpm_v2.CellGlobalId_NrCgi:
				cellObject.CellGlobalID.Value = fmt.Sprintf("%x", bitStringToUint64(cellGlobalID.NrCgi.NRcellIdentity.Value.Value, int(cellGlobalID.NrCgi.NRcellIdentity.Value.Len)))
				cellObject.CellGlobalID.Type = topoapi.CellGlobalIDType_NRCGI
			case *e2sm_kpm_v2.CellGlobalId_EUtraCgi:
				cellObject.CellGlobalID.Value = fmt.Sprintf("%x", bitStringToUint64(cellGlobalID.EUtraCgi.EUtracellIdentity.Value.Value, int(cellGlobalID.EUtraCgi.EUtracellIdentity.Value.Len)))
				cellObject.CellGlobalID.Type = topoapi.CellGlobalIDType_ECGI
			}

			*e2Cells = append(*e2Cells, cellObject)
		}
	}

	for _, reportStyle := range ranFunctionDescription.GetRicReportStyleList() {
		kpmReportStyle := &topoapi.KPMReportStyle{
			Name: reportStyle.RicReportStyleName.Value,
			Type: reportStyle.RicReportStyleType.Value,
		}
		var measurements []*topoapi.KPMMeasurement
		for _, meanInfoItem := range reportStyle.GetMeasInfoActionList().GetValue() {
			measurements = append(measurements, &topoapi.KPMMeasurement{
				ID:   meanInfoItem.GetMeasId().String(),
				Name: meanInfoItem.GetMeasName().GetValue(),
			})
		}

		kpmReportStyle.Measurements = measurements
		ranFunction.ReportStyles = append(ranFunction.ReportStyles, kpmReportStyle)
	}
	ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	serviceModel.RanFunctions = append(serviceModel.RanFunctions, ranFunctionAny)

	return nil
}

// ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel

func bitStringToUint64(bitString []byte, bitCount int) uint64 {
	var result uint64 = 0
	for i, b := range bitString {
		result += uint64(b) << ((len(bitString) - i - 1) * 8)
	}
	if bitCount%8 != 0 {
		return result >> (8 - bitCount%8)
	}
	return result
}
