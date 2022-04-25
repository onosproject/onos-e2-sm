//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"

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

	return cpID, nil
}
