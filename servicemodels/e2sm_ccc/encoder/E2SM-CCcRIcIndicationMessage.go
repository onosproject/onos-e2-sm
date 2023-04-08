// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/choiceOptions"
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"

	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmCCcRIcIndicationMessage(msg *e2smcccv1.E2SmCCcRIcIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-CCcRIcIndicationMessage message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-CCcRIcIndicationMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "", choiceOptions.E2smCccChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-CCcRIcIndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmCCcRIcIndicationMessage(per []byte) (*e2smcccv1.E2SmCCcRIcIndicationMessage, error) {

	log.Debugf("Obtained E2SM-CCcRIcIndicationMessage PER bytes are\n%v", hex.Dump(per))

	result := e2smcccv1.E2SmCCcRIcIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "", choiceOptions.E2smCccChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-CCcRIcIndicationMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-CCcRIcIndicationMessage PDU %s", err.Error())
	}

	return &result, nil
}
