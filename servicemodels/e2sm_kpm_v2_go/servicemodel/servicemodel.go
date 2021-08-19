// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package servicemodel

import (
	"fmt"
	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/encoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdudecoder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"google.golang.org/protobuf/proto"
)

type Kpm2ServiceModel string

const smName = "e2sm_kpm"
const smVersion = "v2_go"
const moduleName = "e2sm_kpm_v2_go.so.2.0"
const smOIDKpmV2 = "1.3.6.1.4.1.53148.1.2.2.2"

func (sm Kpm2ServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDKpmV2,
	}
	return smData
}

func (sm Kpm2ServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmKpmIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmIndicationHeader to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm Kpm2ServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2_go.E2SmKpmIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmKpmIndicationHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm Kpm2ServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmKpmIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmIndicationMessage to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm Kpm2ServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2_go.E2SmKpmIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmKpmIndicationMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm Kpm2ServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmRanFunctionDescription to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmRanFunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm Kpm2ServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmRanfunctionDescription %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmKpmRanFunctionDescription(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm Kpm2ServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmKpmEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmEventTriggerDefinition to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm Kpm2ServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmEventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmKpmEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm Kpm2ServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmKpmActionDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmKpmActionDefinitio to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmKpmActionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm Kpm2ServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_v2_go.E2SmKpmActionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmActionDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmKpmActionDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmActionDefinition to PER %s", err)
	}

	return perBytes, nil
}

//It is redundant so far - could be reused for future, if you need to extract something specific from RanFunctionDescription message
func (sm Kpm2ServiceModel) DecodeRanFunctionDescription(asn1bytes []byte) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error) {
	e2SmKpmPdu, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(asn1bytes)
	if err != nil {
		return nil, nil, nil, err
	}
	return pdudecoder.DecodeE2SmKpmRanFunctionDescription(e2SmKpmPdu)
}

func (sm Kpm2ServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm Kpm2ServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm Kpm2ServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm Kpm2ServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm Kpm2ServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm Kpm2ServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on KPM")
}

func (sm Kpm2ServiceModel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription{}
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
			case *e2sm_kpm_v2_go.CellGlobalId_NrCgi:
				cellObject.CellGlobalID.Value = fmt.Sprintf("%x", bitStringToUint64(cellGlobalID.NrCgi.NRcellIdentity.Value.Value, int(cellGlobalID.NrCgi.NRcellIdentity.Value.Len)))
				cellObject.CellGlobalID.Type = topoapi.CellGlobalIDType_NRCGI
			case *e2sm_kpm_v2_go.CellGlobalId_EUtraCgi:
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
	if err != nil {
		return err
	}

	serviceModel.RanFunctions = []*prototypes.Any{ranFunctionAny}
	return nil
}

func bitStringToUint64(bitString []byte, bitCount int) uint64 {
	var result uint64
	for i, b := range bitString {
		result += uint64(b) << ((len(bitString) - i - 1) * 8)
	}
	if bitCount%8 != 0 {
		return result >> (8 - bitCount%8)
	}
	return result
}
