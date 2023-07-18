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

func PerEncodeE2SmCCcRIceventTriggerDefinition(msg *e2smcccv1.E2SmCCcRIceventTriggerDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-CCcRIceventTriggerDefinition message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-CCcRIceventTriggerDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "", choiceOptions.E2smCccChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-CCcRIceventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmCCcRIceventTriggerDefinition(per []byte) (*e2smcccv1.E2SmCCcRIceventTriggerDefinition, error) {

	log.Debugf("Obtained E2SM-CCcRIceventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2smcccv1.E2SmCCcRIceventTriggerDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "", choiceOptions.E2smCccChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-CCcRIceventTriggerDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-CCcRIceventTriggerDefinition PDU %s", err.Error())
	}

	return &result, nil
}
