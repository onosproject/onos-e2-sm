// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"

	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmCCcRIcControlOutcome(msg *e2smcccv1.E2SmCCcRIcControlOutcome) ([]byte, error) {

	log.Debugf("Obtained E2SM-CCcRIcControlOutcome message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-CCcRIcControlOutcome PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "", choiceOptions.E2smCccChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-CCcRIcControlOutcome PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmCCcRIcControlOutcome(per []byte) (*e2smcccv1.E2SmCCcRIcControlOutcome, error) {

	log.Debugf("Obtained E2SM-CCcRIcControlOutcome PER bytes are\n%v", hex.Dump(per))

	result := e2smcccv1.E2SmCCcRIcControlOutcome{}
	err := aper.UnmarshalWithParams(per, &result, "", choiceOptions.E2smCccChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-CCcRIcControlOutcome from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-CCcRIcControlOutcome PDU %s", err.Error())
	}

	return &result, nil
}
