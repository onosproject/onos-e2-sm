// SPDX-FileCopyrightText: 2023-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"

	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmCCcRIcactionDefinition(msg *e2smcccv1.E2SmCCcRIcactionDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-CCcRIcactionDefinition message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-CCcRIcactionDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "", choiceOptions.E2smCccChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-CCcRIcactionDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmCCcRIcactionDefinition(per []byte) (*e2smcccv1.E2SmCCcRIcactionDefinition, error) {

	log.Debugf("Obtained E2SM-CCcRIcactionDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2smcccv1.E2SmCCcRIcactionDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "", choiceOptions.E2smCccChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-CCcRIcactionDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-CCcRIcactionDefinition PDU %s", err.Error())
	}

	return &result, nil
}
