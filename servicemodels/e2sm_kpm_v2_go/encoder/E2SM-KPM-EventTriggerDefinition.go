// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encoder

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/prometheus/common/log"
)

func PerEncodeE2SmKpmEventTriggerDefinition(etd *e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-KPM-EventTriggerDefinition message is\n%v", etd)
	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*etd, "valueExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-KPM-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmKpmEventTriggerDefinition(per []byte) (*e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition, error) {

	log.Debugf("Obtained E2SM-KPM-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))
	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	result := e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-KPM-EventTriggerDefinition from PER is\n%v", result)

	return &result, nil
}
