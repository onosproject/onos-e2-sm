// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2smrsm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeE2SmRsmControlHeader(ch *e2smrsm.E2SmRsmControlHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RSM-ControlHeader message is\n%v", ch)
	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RSM-ControlHeader PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ch, "valueExt", e2smrsm.RsmChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RSM-ControlHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRsmControlHeader(per []byte) (*e2smrsm.E2SmRsmControlHeader, error) {

	log.Debugf("Obtained E2SM-RSM-ControlHeader PER bytes are\n%v", hex.Dump(per))

	result := e2smrsm.E2SmRsmControlHeader{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smrsm.RsmChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RSM-ControlHeader from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RSM-ControlHeader PDU %s", err.Error())
	}

	return &result, nil
}
