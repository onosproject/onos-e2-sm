// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/choiceOptions"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"

	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeE2SmRcEventTrigger(msg *e2smrcv1.E2SmRcEventTrigger) ([]byte, error) {

	log.Debugf("Obtained E2SM-RcEventTrigger message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcEventTrigger PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RcEventTrigger PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcEventTrigger(per []byte) (*e2smrcv1.E2SmRcEventTrigger, error) {

	log.Debugf("Obtained E2SM-RcEventTrigger PER bytes are\n%v", hex.Dump(per))

	result := e2smrcv1.E2SmRcEventTrigger{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RcEventTrigger from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcEventTrigger PDU %s", err.Error())
	}

	return &result, nil
}
