// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

//nolint
package main

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	kpmv2ctypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/kpmctypes"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdudecoder"
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

//It is redundant so far - could be reused for future, if you need to extract something specific from RanFunctionDescription message
func (sm servicemodel) DecodeRanFunctionDescription(asn1bytes []byte) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error) {
	e2SmKpmPdu, err := kpmv2ctypes.PerDecodeE2SmKpmRanfunctionDescription(asn1bytes)
	if err != nil {
		return nil, nil, nil, err
	}
	return pdudecoder.DecodeE2SmKpmRanfunctionDescription(e2SmKpmPdu)
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

// This function was created for debugging purposes
// It takes as an input string of asn1 bytes (in HEXadecimal format)
// and converts them to the array of bytes ([]byte) which could be passed
// to the input of *ASN1toProto function
// Input data could be obtained in a following way:
// t.Logf("E2SM-KPM-ActionDefinition (Format3) asn1Bytes are \n%x", asn1Bytes)
func (sm servicemodel) Asn1BytesToByte(str string) ([]byte, error) {

	return hex.DecodeString(str)
}

// This function was created for debugging purposes.
// It takes as an input output of hex.Dump() for asn1 bytes
// and converts them to the array of bytes ([]byte)
// which could be passed to the input of *ASN1toProto function
// Input data could be obtained in a following way:
// t.Logf("E2SM-KPM-ActionDefinition (Format3) asn1Bytes are \n%v", hex.Dump(asn1Bytes))
func (sm servicemodel) HexDumpToByte(str string) ([]byte, error) {

	r, err := regexp.Compile("([\t\n\f\r ][0-9a-f]{2})")
	if err != nil {
		return nil, err
	}
	out := r.FindAllString(str, -1)

	res := ""
	escapeElement := 16
	for i, slice := range out {
		postprcs := strings.ReplaceAll(slice, " ", "")
		postprcss := strings.ReplaceAll(postprcs, "  ", "")
		if i != escapeElement {
			res = res + fmt.Sprintf("%v", strings.ReplaceAll(postprcss, "\n", ""))
		} else {
			escapeElement = escapeElement + 17
		}
	}

	b, err := hex.DecodeString(res)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel
