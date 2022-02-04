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

func PerEncodeE2SmRcPreControlMessage(cm *e2sm_rc_pre_go.E2SmRcPreControlMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlMessage message is\n%v", cm)
	//aper.ChoiceMap = e2sm_rc_pre_go.RcPreChoicemap
	per, err := aper.MarshalWithParams(cm, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-ControlMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcPreControlMessage(per []byte) (*e2sm_rc_pre_go.E2SmRcPreControlMessage, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlMessage PER bytes are\n%v", hex.Dump(per))
	//aper.ChoiceMap = e2sm_rc_pre_go.RcPreChoicemap
	result := e2sm_rc_pre_go.E2SmRcPreControlMessage{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-ControlMessage from PER is\n%v", &result)

	return &result, nil
}
