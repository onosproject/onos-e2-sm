// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmRcPreIndicationHeader(ih *e2sm_rc_pre_go.E2SmRcPreIndicationHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-IndicationHeader message is\n%v", ih)
	if err := ih.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RC-PRE-IndicationHeader PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ih, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcPreIndicationHeader(per []byte) (*e2sm_rc_pre_go.E2SmRcPreIndicationHeader, error) {

	log.Debugf("Obtained E2SM-RC-PRE-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	result := e2sm_rc_pre_go.E2SmRcPreIndicationHeader{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-IndicationHeader from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RC-PRE-IndicationHeader PDU %s", err.Error())
	}

	return &result, nil
}
