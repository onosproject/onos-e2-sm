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
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
)

/*
   servicemodel package implements interface defined in:
   https://github.com/onosproject/onos-e2t/blob/2ec06f2fde15fd765497120b15661ac8fd61d6eb/pkg/modelregistry/modelregistry.go#L43-L62
   If any new top-level PDU occurs, then this template should be correspondingly adjusted and the interface (link above) should be extended.
*/

type RCServiceModel string

const smName = "e2smrcv1"
const smVersion = "v1"
const moduleName = "e2smrcv1.so.1.0.1"
const smOID = "1.3.6.1.4.1.53148.1.1.2.3"

func (sm RCServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOID,
	}
	return smData
}

func (sm RCServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmRcIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcIndicationHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm RCServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smrcv1.E2SmRcIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcIndicationHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm RCServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmRcIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcIndicationMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm RCServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smrcv1.E2SmRcIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcIndicationMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm RCServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcRanfunctionDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcRanfunctionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm RCServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smrcv1.E2SmRcRanfunctionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcRanfunctionDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcRanfunctionDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm RCServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmRcEventTrigger(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcEventTrigger to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcEventTrigger %s", err)
	}

	return protoBytes, nil
}

func (sm RCServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smrcv1.E2SmRcEventTrigger)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcEventTrigger %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcEventTrigger(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcEventTrigger to PER %s", err)
	}

	return perBytes, nil
}

func (sm RCServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmRcActionDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcActionDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcActionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm RCServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smrcv1.E2SmRcActionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcActionDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcActionDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcActionDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm RCServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmRcControlHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcControlHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm RCServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smrcv1.E2SmRcControlHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcControlHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcControlHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcControlHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm RCServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmRcControlMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcControlMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcControlMessage %s", err)
	}

	return protoBytes, nil
}

func (sm RCServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smrcv1.E2SmRcControlMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcControlMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcControlMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcControlMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm RCServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmRcControlOutcome(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcControlOutcome to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcControlOutcome %s", err)
	}

	return protoBytes, nil
}

func (sm RCServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smrcv1.E2SmRcControlOutcome)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcControlOutcome %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcControlOutcome(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcControlOutcome to PER %s", err)
	}

	return perBytes, nil
}

func (sm RCServiceModel) CallProcessIDASN1toProto(asn1Bytes []byte) ([]byte, error) {

	perBytes, err := encoder.PerDecodeE2SmRcCallProcessId(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRcCallProcessId to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRcCallProcessId %s", err)
	}

	return protoBytes, nil
}

func (sm RCServiceModel) CallProcessIDProtoToASN1(protoBytes []byte) ([]byte, error) {

	protoObj := new(e2smrcv1.E2SmRcCallProcessId)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRcCallProcessId %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRcCallProcessId(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRcCallProcessId to PER %s", err)
	}

	return perBytes, nil
}

// OnSetup extract the required information from the RAN function description. RanFunctionDescription may vary from SM to SM and thus you'll have to adjust this function in order to extract all necessary data
func (sm RCServiceModel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2smrcv1.E2SmRcRanfunctionDefinition{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	serviceModel := serviceModels[smOID]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName

	// Extract report style information
	reportStyleList := ranFunctionDescription.GetRanFunctionDefinitionReport().GetRicReportStyleList()
	ranFunction := &topoapi.RCRanFunction{}
	for _, reportStyle := range reportStyleList {
		reportStyleRANParameters := reportStyle.RanReportParametersList
		rcReportStyle := &topoapi.RCReportStyle{
			Name: reportStyle.RicReportStyleName.Value,
			Type: reportStyle.RicReportStyleType.Value,
		}
		ranParameters := make([]*topoapi.RANParameter, 0)
		for _, ranParameter := range reportStyleRANParameters {
			ranParameters = append(ranParameters, &topoapi.RANParameter{
				ID:   ranParameter.RanParameterId.Value,
				Name: ranParameter.RanParameterName.Value,
			})
		}
		rcReportStyle.RanParameters = ranParameters
		ranFunction.ReportStyles = append(ranFunction.ReportStyles, rcReportStyle)
	}

	// Extracts insert style information
	insertStyleList := ranFunctionDescription.GetRanFunctionDefinitionInsert().GetRicInsertStyleList()
	for _, insertStyle := range insertStyleList {
		rcInsertStyle := &topoapi.RCInsertStyle{
			Name: insertStyle.RicInsertStyleName.Value,
			Type: insertStyle.RicInsertStyleType.Value,
		}
		insertIndicationList := make([]*topoapi.InsertIndication, 0)
		for _, insertIndication := range insertStyle.GetRicInsertIndicationList() {
			rcInsertIndication := &topoapi.InsertIndication{
				ID:   insertIndication.RicInsertIndicationId.Value,
				Name: insertIndication.RicInsertIndicationName.Value,
			}

			insertIndicationRANParameters := make([]*topoapi.RANParameter, 0)
			for _, indicationParameter := range insertIndication.RanInsertIndicationParametersList {
				insertIndicationRANParameters = append(insertIndicationRANParameters, &topoapi.RANParameter{
					ID:   indicationParameter.RanParameterId.Value,
					Name: indicationParameter.RanParameterName.Value,
				})
			}
			rcInsertIndication.RanParameters = insertIndicationRANParameters
			insertIndicationList = append(insertIndicationList, rcInsertIndication)
		}
		rcInsertStyle.InsertIndications = insertIndicationList
		ranFunction.InsertStyles = append(ranFunction.InsertStyles, rcInsertStyle)

	}
	eventTriggerStyleList := ranFunctionDescription.GetRanFunctionDefinitionEventTrigger().GetRicEventTriggerStyleList()
	for _, eventTriggerStyle := range eventTriggerStyleList {
		rcEventTriggerStyle := &topoapi.RCEventTriggerStyle{
			Name:       eventTriggerStyle.RicEventTriggerStyleName.Value,
			Type:       eventTriggerStyle.RicEventTriggerStyleType.Value,
			FormatType: eventTriggerStyle.RicEventTriggerFormatType.Value,
		}
		ranFunction.EventTriggerStyles = append(ranFunction.EventTriggerStyles, rcEventTriggerStyle)
	}

	policyStyleList := ranFunctionDescription.GetRanFunctionDefinitionPolicy().GetRicPolicyStyleList()
	for _, policyStyle := range policyStyleList {
		policyActionList := make([]*topoapi.PolicyAction, 0)
		for _, policyAction := range policyStyle.GetRicPolicyActionList() {
			actionRANParams := make([]*topoapi.RANParameter, 0)
			conditionRANParams := make([]*topoapi.RANParameter, 0)
			for _, actionRANParam := range policyAction.RanPolicyActionParametersList {
				ranParam := &topoapi.RANParameter{
					ID:   actionRANParam.RanParameterId.Value,
					Name: actionRANParam.RanParameterName.Value,
				}
				actionRANParams = append(actionRANParams, ranParam)
			}

			for _, conditionRANParam := range policyAction.RanPolicyConditionParametersList {
				ranParam := &topoapi.RANParameter{
					ID:   conditionRANParam.RanParameterId.Value,
					Name: conditionRANParam.RanParameterName.Value,
				}
				conditionRANParams = append(conditionRANParams, ranParam)
			}

			rcPolicyAction := &topoapi.PolicyAction{
				ID:                           policyAction.RicPolicyActionId.Value,
				Name:                         policyAction.RicPolicyActionName.Value,
				PolicyActionRanParameters:    actionRANParams,
				PolicyConditionRanParameters: conditionRANParams,
			}
			policyActionList = append(policyActionList, rcPolicyAction)
		}

		rcPolicyStyle := &topoapi.RCPolicyStyle{
			Name:          policyStyle.RicPolicyStyleName.Value,
			Type:          policyStyle.RicPolicyStyleType.Value,
			PolicyActions: policyActionList,
		}
		ranFunction.PolicyStyles = append(ranFunction.PolicyStyles, rcPolicyStyle)
	}

	controlStyleList := ranFunctionDescription.GetRanFunctionDefinitionControl().GetRicControlStyleList()
	for _, controlStyle := range controlStyleList {
		rcControlStyle := &topoapi.RCControlStyle{
			Name:                     controlStyle.RicControlStyleName.Value,
			Type:                     controlStyle.RicControlStyleType.Value,
			HeaderFormatType:         controlStyle.RicControlHeaderFormatType.Value,
			MessageFormatType:        controlStyle.RicControlMessageFormatType.Value,
			ControlOutcomeFormatType: controlStyle.RicControlOutcomeFormatType.Value,
		}

		for _, action := range controlStyle.RicControlActionList {
			controlAction := &topoapi.ControlAction{
				Name: action.RicControlActionName.Value,
				ID:   action.RicControlActionId.Value,
			}
			for _, ranControlParameter := range action.RanControlActionParametersList {
				controlAction.RanParameters = append(controlAction.RanParameters, &topoapi.RANParameter{
					Name: ranControlParameter.RanParameterName.Value,
					ID:   ranControlParameter.RanParameterId.Value,
				})
			}
			rcControlStyle.ControlActions = append(rcControlStyle.ControlActions, controlAction)
		}

		ranFunction.ControlStyles = append(ranFunction.ControlStyles, rcControlStyle)
	}

	ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	if err != nil {
		return err
	}
	serviceModel.RanFunctions = append(serviceModel.RanFunctions, ranFunctionAny)
	return nil
}
