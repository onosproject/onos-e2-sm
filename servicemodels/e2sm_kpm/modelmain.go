// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
)

type servicemodel string

const smname = "e2sm_kpm"
const smversion = "v1beta1"
const modulename = "e2sm_kpm.so.1.0.1"

func (sm servicemodel) ServiceModelData() (string, string, string) {
	return smname, smversion, modulename
}

func (sm servicemodel) IndicationHeaderASN1toProto([]byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_ies.E2SmKpmIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmIndicationHeader %s", err)
	}

	// TODO: change this to Per when Per is fixed
	perBytes, err := kpmctypes.XerEncodeE2SmKpmIndicationHeader(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) IndicationMessageASN1toProto([]byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2sm_kpm_ies.E2SmKpmIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling protoBytes to E2SmKpmIndicationMessage %s", err)
	}

	// TODO: change this to Per when Per is fixed
	perBytes, err := kpmctypes.XerEncodeE2SmKpmIndicationMessage(protoObj)
	if err != nil {
		return nil, fmt.Errorf("error encoding E2SmKpmIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm servicemodel) RanFuncDescriptionASN1toProto([]byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (sm servicemodel) RanFuncDescriptionProtoToASN1([]byte) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

// ServiceModel is the exported symbol that gives an entry point to this shared module
var ServiceModel servicemodel
