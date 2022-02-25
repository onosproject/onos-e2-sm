// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2smmhov2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmMhoControlMessage(cm *e2smmhov2.E2SmMhoControlMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-MHO-ControlMessage message is\n%v", cm)
	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MHO-ControlMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(cm, "choiceExt", e2smmhov2.MhoChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-MHO-ControlMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMhoControlMessage(per []byte) (*e2smmhov2.E2SmMhoControlMessage, error) {

	log.Debugf("Obtained E2SM-MHO-ControlMessage PER bytes are\n%v", hex.Dump(per))

	result := e2smmhov2.E2SmMhoControlMessage{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2smmhov2.MhoChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-MHO-ControlMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MHO-ControlMessage PDU %s", err.Error())
	}

	return &result, nil
}
