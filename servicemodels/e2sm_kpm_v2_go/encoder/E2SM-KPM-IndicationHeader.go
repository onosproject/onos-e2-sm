// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmKpmIndicationHeader(ih *e2smkpmv2.E2SmKpmIndicationHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-KPMv2-IndicationHeader message is\n%v", ih)
	if err := ih.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-KPMv2-IndicationHeader PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ih, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-KPMv2-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmKpmIndicationHeader(per []byte) (*e2smkpmv2.E2SmKpmIndicationHeader, error) {

	log.Debugf("Obtained E2SM-KPMv2-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	result := e2smkpmv2.E2SmKpmIndicationHeader{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-KPMv2-IndicationHeader from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-KPMv2-IndicationHeader PDU %s", err.Error())
	}

	return &result, nil
}
