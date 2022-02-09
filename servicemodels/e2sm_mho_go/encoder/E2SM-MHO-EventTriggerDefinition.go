// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"github.com/google/martian/log"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func PerEncodeE2SmMhoEventTriggerDefinition(etd *e2sm_mho_go.E2SmMhoEventTriggerDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-MHO-EventTriggerDefinition message is\n%v", etd)

	per, err := aper.MarshalWithParams(etd, "valueExt", e2sm_mho_go.MhoChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-MHO-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMhoEventTriggerDefinition(per []byte) (*e2sm_mho_go.E2SmMhoEventTriggerDefinition, error) {

	log.Debugf("Obtained E2SM-MHO-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2sm_mho_go.E2SmMhoEventTriggerDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_mho_go.MhoChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-MHO-EventTriggerDefinition from PER is\n%v", &result)

	return &result, nil
}
