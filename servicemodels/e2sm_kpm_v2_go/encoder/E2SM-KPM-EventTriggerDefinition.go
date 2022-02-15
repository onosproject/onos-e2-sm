// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func PerEncodeE2SmKpmEventTriggerDefinition(etd *e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-KPM-EventTriggerDefinition message is\n%v", etd)

	per, err := aper.MarshalWithParams(etd, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-KPM-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmKpmEventTriggerDefinition(per []byte) (*e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition, error) {

	log.Debugf("Obtained E2SM-KPM-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-KPM-EventTriggerDefinition from PER is\n%v", &result)

	return &result, nil
}
