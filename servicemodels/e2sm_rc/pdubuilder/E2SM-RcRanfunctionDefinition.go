//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
)

func CreateE2SmRcRanfunctionDefinition(ranFunctionShortName string, ranFunctionOID string, ranFunctionDescription string) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,
			RanFunctionE2SmOid:     ranFunctionOID,
			RanFunctionDescription: ranFunctionDescription,
		},
	}

	return msg, nil
}
