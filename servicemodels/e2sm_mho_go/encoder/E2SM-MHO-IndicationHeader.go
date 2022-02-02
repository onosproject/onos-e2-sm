// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package encoder

import (
	"encoding/hex"
	"github.com/google/martian/log"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/goaperlib"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v1/e2sm-mho-go"
)

func PerEncodeE2SmMhoIndicationHeader(ih *e2sm_mho_go.E2SmMhoIndicationHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-MHO-IndicationHeader message is\n%v", ih)
	//aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := goaperlib.MarshalWithParams(ih, "choiceExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-MHO-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMhoIndicationHeader(per []byte) (*e2sm_mho_go.E2SmMhoIndicationHeader, error) {

	log.Debugf("Obtained E2SM-MHO-IndicationHeader PER bytes are\n%v", hex.Dump(per))
	//aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	result := e2sm_mho_go.E2SmMhoIndicationHeader{}
	err := goaperlib.UnmarshalWithParams(per, &result, "choiceExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-MHO-IndicationHeader from PER is\n%v", &result)

	return &result, nil
}
