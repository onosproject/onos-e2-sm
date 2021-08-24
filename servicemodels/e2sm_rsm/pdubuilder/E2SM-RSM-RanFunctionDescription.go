// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies"
)

func CreateE2SmRsmRanFunctionDescription(rfSn string, rfE2SMoid string, rfd string) *e2sm_rsm_ies.E2SmRsmRanfunctionDescription {

	e2SmRsmPdu := e2sm_rsm_ies.E2SmRsmRanfunctionDescription{
		RanFunctionName: &e2sm_v2_ies.RanfunctionName{
			RanFunctionShortName:   rfSn,
			RanFunctionE2SmOid:     rfE2SMoid,
			RanFunctionDescription: rfd,
		},
		RicSlicingNodeCapabilityList: nil,
	}

	//if err := e2SmRsmPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmRsmRanfunctionDescription %s", err.Error())
	//}
	return &e2SmRsmPdu
}
