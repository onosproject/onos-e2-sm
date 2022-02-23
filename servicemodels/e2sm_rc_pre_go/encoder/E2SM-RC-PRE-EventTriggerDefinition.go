// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"fmt"
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func PerEncodeE2SmRcPreEventTriggerDefinition(etd *e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-EventTriggerDefinition message is\n%v", etd)
	if err := etd.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SM-RC-PRE-EventTriggerDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(etd, "valueExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcPreEventTriggerDefinition(per []byte) (*e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinition, error) {

	log.Debugf("Obtained E2SM-RC-PRE-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2sm_rc_pre_go.E2SmRcPreEventTriggerDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-EventTriggerDefinition from PER is\n%v", &result)
	if err := result.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SM-RC-PRE-EventTriggerDefinition PDU %s", err.Error())
	}

	return &result, nil
}
