// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package servicemodel

import (
	"fmt"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/encoder"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies"
	"google.golang.org/protobuf/proto"
)

type RsmServiceModel string

const smName = "e2sm_rsm"
const smVersion = "v1_go"
const moduleName = "e2sm_rsm_v1_go.so.2.0"
const smOIDKpmV2 = "1.3.6.1.4.1.53148.1.1.2.102"

func (sm RsmServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDKpmV2,
	}
	return smData
}

func (sm RsmServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRsmIndicationHeader to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRsmIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rsm_ies.E2SmRsmIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRsmIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmIndicationHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRsmIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRsmIndicationMessage to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRsmIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rsm_ies.E2SmRsmIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRsmIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmIndicationMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRsmIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmRanFunctionDescription(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRsmRanFunctionDescription to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRsmRanFunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rsm_ies.E2SmRsmRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRsmRanfunctionDescription %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmRanFunctionDescription(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRsmRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRsmEventTriggerDefinition to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRsmEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rsm_ies.E2SmRsmEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRsmEventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRsmEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on RSM")
}

func (sm RsmServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on RSM")
}

//It is redundant so far - could be reused for future, if you need to extract something specific from RanFunctionDescription message
func (sm RsmServiceModel) DecodeRanFunctionDescription(asn1bytes []byte) (*e2sm_v2_ies.RanfunctionName, []*e2sm_rsm_ies.NodeSlicingCapabilityItem, error) {
	e2SmRsmPdu, err := encoder.PerDecodeE2SmRsmRanFunctionDescription(asn1bytes)
	if err != nil {
		return nil, nil, err
	}

	return e2SmRsmPdu.GetRanFunctionName(), e2SmRsmPdu.GetRicSlicingNodeCapabilityList(), nil
}

func (sm RsmServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmControlHeader(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRsmControlHeader to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRsmControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rsm_ies.E2SmRsmControlHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRsmControlHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmControlHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRsmControlHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRsmControlMessage(asn1Bytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding E2SmRsmControlMessage to PER %s", err)
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling asn1Bytes to E2SmRsmControlMessage %s", err)
	}

	return protoBytes, nil
}

func (sm RsmServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_rsm_ies.E2SmRsmControlMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmRsmControlMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRsmControlMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmRsmControlMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm RsmServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on RSM")
}

func (sm RsmServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented on RSM")
}

func (sm RsmServiceModel) OnSetup(request *types.OnSetupRequest) error {
	//protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	//if err != nil {
	//	return err
	//}
	//ranFunctionDescription := &e2sm_rsm_ies.E2SmRsmRanfunctionDescription{}
	//err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	//if err != nil {
	//	return err
	//}
	//serviceModels := request.ServiceModels
	//e2Cells := request.E2Cells
	//serviceModel := serviceModels[smOIDKpmV2]
	//serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName
	//ranFunction := &topoapi.KPMRanFunction{}
	//
	//for _, kpmNode := range ranFunctionDescription.RicKpmNodeList {
	//	for _, cell := range kpmNode.CellMeasurementObjectList {
	//		cellObject := &topoapi.E2Cell{
	//			CellObjectID: cell.GetCellObjectId().GetValue(),
	//			CellGlobalID: &topoapi.CellGlobalID{},
	//		}
	//		switch cellGlobalID := cell.CellGlobalId.GetCellGlobalId().(type) {
	//		case *e2sm_kpm_v2_go.CellGlobalId_NrCgi:
	//			cellObject.CellGlobalID.Value = fmt.Sprintf("%x", bitStringToUint64(cellGlobalID.NrCgi.NRcellIdentity.Value.Value, int(cellGlobalID.NrCgi.NRcellIdentity.Value.Len)))
	//			cellObject.CellGlobalID.Type = topoapi.CellGlobalIDType_NRCGI
	//		case *e2sm_kpm_v2_go.CellGlobalId_EUtraCgi:
	//			cellObject.CellGlobalID.Value = fmt.Sprintf("%x", bitStringToUint64(cellGlobalID.EUtraCgi.EUtracellIdentity.Value.Value, int(cellGlobalID.EUtraCgi.EUtracellIdentity.Value.Len)))
	//			cellObject.CellGlobalID.Type = topoapi.CellGlobalIDType_ECGI
	//		}
	//
	//		*e2Cells = append(*e2Cells, cellObject)
	//	}
	//}
	//
	//for _, reportStyle := range ranFunctionDescription.GetRicReportStyleList() {
	//	kpmReportStyle := &topoapi.KPMReportStyle{
	//		Name: reportStyle.RicReportStyleName.Value,
	//		Type: reportStyle.RicReportStyleType.Value,
	//	}
	//	var measurements []*topoapi.KPMMeasurement
	//	for _, meanInfoItem := range reportStyle.GetMeasInfoActionList().GetValue() {
	//		measurements = append(measurements, &topoapi.KPMMeasurement{
	//			ID:   meanInfoItem.GetMeasId().String(),
	//			Name: meanInfoItem.GetMeasName().GetValue(),
	//		})
	//	}
	//
	//	kpmReportStyle.Measurements = measurements
	//	ranFunction.ReportStyles = append(ranFunction.ReportStyles, kpmReportStyle)
	//}
	//ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	//if err != nil {
	//	return err
	//}
	//
	//serviceModel.RanFunctions = []*prototypes.Any{ranFunctionAny}
	return nil
}

//func bitStringToUint64(bitString []byte, bitCount int) uint64 {
//	var result uint64
//	for i, b := range bitString {
//		result += uint64(b) << ((len(bitString) - i - 1) * 8)
//	}
//	if bitCount%8 != 0 {
//		return result >> (8 - bitCount%8)
//	}
//	return result
//}
