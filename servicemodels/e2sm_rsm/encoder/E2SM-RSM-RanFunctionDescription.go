// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2smrsm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmRsmRanFunctionDescription(rfd *e2smrsm.E2SmRsmRanfunctionDescription) ([]byte, error) {

	log.Debugf("Obtained E2SM-RSM-RanFunctionDescription message is\n%v", rfd)
	if err := rfd.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RSM-RanFunctionDescription PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(rfd, "valueExt", e2smrsm.RsmChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RSM-RanFunctionDescription PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRsmRanFunctionDescription(per []byte) (*e2smrsm.E2SmRsmRanfunctionDescription, error) {

	log.Debugf("Obtained E2SM-RSM-RanFunctionDescription PER bytes are\n%v", hex.Dump(per))

	result := e2smrsm.E2SmRsmRanfunctionDescription{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smrsm.RsmChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RSM-RanFunctionDescription from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RSM-RanFunctionDescription PDU %s", err.Error())
	}

	return &result, nil
}
