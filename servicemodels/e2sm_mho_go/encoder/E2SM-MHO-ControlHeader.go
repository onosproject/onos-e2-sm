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

func init() {
	log.SetLevel(log.Info)
}

func PerEncodeE2SmMhoControlHeader(ch *e2sm_mho_go.E2SmMhoControlHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader message is\n%v", ch)
	//aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := goaperlib.MarshalWithParams(ch, "choiceExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMhoControlHeader(per []byte) (*e2sm_mho_go.E2SmMhoControlHeader, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))
	//aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	result := e2sm_mho_go.E2SmMhoControlHeader{}
	err := goaperlib.UnmarshalWithParams(per, &result, "choiceExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-ControlHeader from PER is\n%v", &result)

	return &result, nil
}
