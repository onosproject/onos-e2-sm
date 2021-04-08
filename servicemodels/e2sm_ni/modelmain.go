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
)

type servicemodel string

const smName = "e2sm_ni"
const smVersion = "v1beta1"
const moduleName = "e2sm_ni.so.1.0.1"
const smOIDNiV1 = "1.3.6.1.4.1.53148.1.1.2.1"

func (sm servicemodel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDNiV1,
	}
	return smData
}

func (sm servicemodel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}
func (sm servicemodel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}
func (sm servicemodel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}
func (sm servicemodel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}
func (sm servicemodel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) DecodeRanFunctionDescription(asn1bytes []byte) (*types.RanfunctionNameDef, *types.RicEventTriggerList, *types.RicReportList, error) {
	return nil, nil, nil, fmt.Errorf("not yet implemented")
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

// ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel
