// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"

	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	
)

func PerEncodeE2SmRcRanfunctionDefinition(msg *e2smrcv1.E2SmRcRanfunctionDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-RcRanfunctionDefinition message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcRanfunctionDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RcRanfunctionDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcRanfunctionDefinition(per []byte) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	log.Debugf("Obtained E2SM-RcRanfunctionDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2smrcv1.E2SmRcRanfunctionDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RcRanfunctionDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcRanfunctionDefinition PDU %s", err.Error())
	}

	return &result, nil
}