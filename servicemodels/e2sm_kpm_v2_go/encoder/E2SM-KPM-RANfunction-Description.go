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

func PerEncodeE2SmKpmRanFunctionDescription(rfd *e2smkpmv2.E2SmKpmRanfunctionDescription) ([]byte, error) {

	log.Debugf("Obtained E2SM-KPMv2-RANfunctionDescription message is\n%v", rfd)
	if err := rfd.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-KPMv2-RANfunctionDescription PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(rfd, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-KPMv2-RANfunctionDescription PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmKpmRanFunctionDescription(per []byte) (*e2smkpmv2.E2SmKpmRanfunctionDescription, error) {

	log.Debugf("Obtained E2SM-KPMv2-RANfunctionDescription PER bytes are\n%v", hex.Dump(per))

	result := e2smkpmv2.E2SmKpmRanfunctionDescription{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-KPMv2-RANfunctionDescription from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-KPMv2-RANfunctionDescription PDU %s", err.Error())
	}

	return &result, nil
}
