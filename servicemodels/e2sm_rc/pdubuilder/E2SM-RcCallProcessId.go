//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcCallProcessIDFormat1(ranCallProcessID int32) (*e2smrcv1.E2SmRcCallProcessId, error) {

	cpID := &e2smrcv1.E2SmRcCallProcessId{
		RicCallProcessIdFormats: &e2smrcv1.RicCallProcessIdFormats{
			RicCallProcessIdFormats: &e2smrcv1.RicCallProcessIdFormats_CallProcessIdFormat1{
				CallProcessIdFormat1: &e2smrcv1.E2SmRcCallProcessIdFormat1{
					RicCallProcessId: &e2smrcv1.RanCallProcessId{
						Value: ranCallProcessID,
					},
				},
			},
		},
	}

	if err := cpID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcCallProcessIDFormat1() error validating E2SM-RC PDU %s", err.Error())
	}

	return cpID, nil
}
