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

func PerEncodeE2SmMhoControlMessage(cm *e2sm_mho_go.E2SmMhoControlMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-MHO-ControlMessage message is\n%v", cm)

	per, err := aper.MarshalWithParams(cm, "choiceExt", e2sm_mho_go.MhoChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-MHO-ControlMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMhoControlMessage(per []byte) (*e2sm_mho_go.E2SmMhoControlMessage, error) {

	log.Debugf("Obtained E2SM-MHO-ControlMessage PER bytes are\n%v", hex.Dump(per))

	result := e2sm_mho_go.E2SmMhoControlMessage{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2sm_mho_go.MhoChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-MHO-ControlMessage from PER is\n%v", &result)

	return &result, nil
}
