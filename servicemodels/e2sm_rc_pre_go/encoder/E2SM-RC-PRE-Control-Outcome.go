// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("e2sm", "rcprev2", "encoder")

func PerEncodeE2SmRcPreControlOutcome(co *e2sm_rc_pre_go.E2SmRcPreControlOutcome) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlOutcome message is\n%v", co)

	per, err := aper.MarshalWithParams(co, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-ControlOutcome PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcPreControlOutcome(per []byte) (*e2sm_rc_pre_go.E2SmRcPreControlOutcome, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlOutcome PER bytes are\n%v", hex.Dump(per))

	result := e2sm_rc_pre_go.E2SmRcPreControlOutcome{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2sm_rc_pre_go.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-ControlOutcome from PER is\n%v", &result)

	return &result, nil
}
