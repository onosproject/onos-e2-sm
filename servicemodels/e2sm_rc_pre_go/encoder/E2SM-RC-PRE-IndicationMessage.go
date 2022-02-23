// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"fmt"
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func PerEncodeE2SmRcPreIndicationMessage(im *e2sm_rc_pre_go.E2SmRcPreIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-IndicationMessage message is\n%v", im)
	if err := im.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SM-RC-PRE-IndicationMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(im, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcPreIndicationMessage(per []byte) (*e2sm_rc_pre_go.E2SmRcPreIndicationMessage, error) {

	log.Debugf("Obtained E2SM-RC-PRE-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	result := e2sm_rc_pre_go.E2SmRcPreIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-IndicationMessage from PER is\n%v", &result)
	if err := result.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SM-RC-PRE-IndicationMessage PDU %s", err.Error())
	}

	return &result, nil
}
