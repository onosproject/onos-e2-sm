// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2sm_mho_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeE2SmMhoControlHeader(ch *e2sm_mho_go.E2SmMhoControlHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader message is\n%v", ch)
	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MHO-ControlHeader PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ch, "choiceExt", e2sm_mho_go.MhoChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMhoControlHeader(per []byte) (*e2sm_mho_go.E2SmMhoControlHeader, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))

	result := e2sm_mho_go.E2SmMhoControlHeader{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2sm_mho_go.MhoChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-ControlHeader from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MHO-ControlHeader PDU %s", err.Error())
	}

	return &result, nil
}
