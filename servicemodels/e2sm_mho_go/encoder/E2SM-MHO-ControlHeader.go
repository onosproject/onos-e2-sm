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

func init() {
	log.SetLevel(log.Info)
}

func PerEncodeE2SmMhoControlHeader(ch *e2sm_mho_go.E2SmMhoControlHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader message is\n%v", ch)
	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(ch, "valueExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMhoControlHeader(per []byte) (*e2sm_mho_go.E2SmMhoControlHeader, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))
	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	result := e2sm_mho_go.E2SmMhoControlHeader{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-ControlHeader from PER is\n%v", &result)

	return &result, nil
}
