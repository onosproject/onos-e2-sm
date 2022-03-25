// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2appdudescriptionsv2 "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"

	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeE2ApPdu(msg *e2appdudescriptionsv2.E2ApPdu) ([]byte, error) {

	log.Debugf("Obtained E2AP message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2AP PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "valueExt", e2appdudescriptionsv2.E2apChoicemap, e2appdudescriptionsv2.E2apCanonicalChoicemap)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2AP PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2ApPdu(per []byte) (*e2appdudescriptionsv2.E2ApPdu, error) {

	log.Debugf("Obtained E2AP PER bytes are\n%v", hex.Dump(per))

	result := e2appdudescriptionsv2.E2ApPdu{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2appdudescriptionsv2.E2apChoicemap, e2appdudescriptionsv2.E2apCanonicalChoicemap)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2AP from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2AP PDU %s", err.Error())
	}

	return &result, nil
}
