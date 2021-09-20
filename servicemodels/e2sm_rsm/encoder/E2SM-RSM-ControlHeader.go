// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package encoder

import (
	"encoding/hex"
	"github.com/google/martian/log"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func PerEncodeE2SmRsmControlHeader(ch *e2sm_rsm_ies.E2SmRsmControlHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RSM-ControlHeader message is\n%v", ch)
	aper.ChoiceMap = e2sm_rsm_ies.RsmChoicemap
	per, err := aper.MarshalWithParams(ch, "valueExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RSM-ControlHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRsmControlHeader(per []byte) (*e2sm_rsm_ies.E2SmRsmControlHeader, error) {

	log.Debugf("Obtained E2SM-RSM-ControlHeader PER bytes are\n%v", hex.Dump(per))
	aper.ChoiceMap = e2sm_rsm_ies.RsmChoicemap
	result := e2sm_rsm_ies.E2SmRsmControlHeader{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RSM-ControlHeader from PER is\n%v", &result)

	return &result, nil
}
