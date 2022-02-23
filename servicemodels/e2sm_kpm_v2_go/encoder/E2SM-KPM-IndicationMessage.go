// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmKpmIndicationMessage(im *e2sm_kpm_v2_go.E2SmKpmIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-KPMv2-IndicationMessage message is\n%v", im)
	if err := im.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-KPMv2-IndicationMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(im, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-KPMv2-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmKpmIndicationMessage(per []byte) (*e2sm_kpm_v2_go.E2SmKpmIndicationMessage, error) {

	log.Debugf("Obtained E2SM-KPMv2-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-KPMv2-IndicationMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-KPMv2-IndicationMessage PDU %s", err.Error())
	}

	return &result, nil
}
