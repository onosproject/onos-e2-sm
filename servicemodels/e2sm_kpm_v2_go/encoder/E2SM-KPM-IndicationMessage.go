// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package encoder

import (
	"encoding/hex"
	"github.com/google/martian/log"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func PerEncodeE2SmKpmIndicationMessage(im *e2sm_kpm_v2_go.E2SmKpmIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-KPM-IndicationMessage message is\n%v", im)
	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(im, "valueExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-KPM-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmKpmIndicationMessage(per []byte) (*e2sm_kpm_v2_go.E2SmKpmIndicationMessage, error) {

	log.Debugf("Obtained E2SM-KPM-IndicationMessage PER bytes are\n%v", hex.Dump(per))
	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	result := e2sm_kpm_v2_go.E2SmKpmIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-KPM-IndicationMessage from PER is\n%v", &result)

	return &result, nil
}
