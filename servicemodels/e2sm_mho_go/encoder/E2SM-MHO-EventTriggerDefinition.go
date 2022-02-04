// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"fmt"
	"github.com/google/martian/log"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func PerEncodeE2SmMhoEventTriggerDefinition(etd *e2sm_mho_go.E2SmMhoEventTriggerDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-MHO-EventTriggerDefinition message is\n%v", etd)
	if err := etd.Validate(); err != nil {
		return nil, fmt.Errorf("E2SM-MHO-EventTriggerDefinition PDU validation has failed. Make sure your PDU correspond to the ASN1 definition\n%s", err)
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(etd, "valueExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-MHO-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMhoEventTriggerDefinition(per []byte) (*e2sm_mho_go.E2SmMhoEventTriggerDefinition, error) {

	log.Debugf("Obtained E2SM-MHO-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))
	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	result := e2sm_mho_go.E2SmMhoEventTriggerDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-MHO-EventTriggerDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, fmt.Errorf("E2SM-MHO-EventTriggerDefinition PDU validation has failed: %s", err)
	}

	return &result, nil
}
