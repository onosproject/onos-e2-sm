// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package encoder

import (
	"encoding/hex"
	"github.com/google/martian/log"
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func init() {
	log.SetLevel(log.Info)
}

func PerEncodeE2SmRcPreControlHeader(ch *e2sm_rc_pre_go.E2SmRcPreControlHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader message is\n%v", ch)
	//aper.ChoiceMap = e2sm_rc_pre_go.RcPreChoicemap
	per, err := aper.MarshalWithParams(ch, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcPreControlHeader(per []byte) (*e2sm_rc_pre_go.E2SmRcPreControlHeader, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))
	//aper.ChoiceMap = e2sm_rc_pre_go.RcPreChoicemap
	result := e2sm_rc_pre_go.E2SmRcPreControlHeader{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-ControlHeader from PER is\n%v", &result)

	return &result, nil
}
