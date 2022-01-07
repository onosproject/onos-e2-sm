// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package encoder

import (
	"encoding/hex"
	"github.com/google/martian/log"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func PerEncodeE2SmMhoIndicationMessage(im *e2sm_mho_go.E2SmMhoIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-MHO-IndicationMessage message is\n%v", im)
	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(im, "choiceExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-MHO-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMhoIndicationMessage(per []byte) (*e2sm_mho_go.E2SmMhoIndicationMessage, error) {

	log.Debugf("Obtained E2SM-MHO-IndicationMessage PER bytes are\n%v", hex.Dump(per))
	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	result := e2sm_mho_go.E2SmMhoIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-MHO-IndicationMessage from PER is\n%v", &result)

	return &result, nil
}
