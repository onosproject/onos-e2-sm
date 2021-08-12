// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package encoder

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("e2sm_kpm_v2_go")

func PerEncodeE2SmKpmActionDefinition(ad *e2sm_kpm_v2_go.E2SmKpmActionDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-KPM-ActionDefinition message is\n%v", ad)
	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(ad, "valueExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-KPM-ActionDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmKpmActionDefinition(per []byte) (*e2sm_kpm_v2_go.E2SmKpmActionDefinition, error) {

	log.Debugf("Obtained E2SM-KPM-ActionDefinition PER bytes are\n%v", hex.Dump(per))
	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	result := e2sm_kpm_v2_go.E2SmKpmActionDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-KPM-ActionDefinition from PER is\n%v", &result)

	return &result, nil
}
