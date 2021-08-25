// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"

func CreateE2SmRsmControlHeader(command e2sm_rsm_ies.E2SmRsmCommand) *e2sm_rsm_ies.E2SmRsmControlHeader {

	return &e2sm_rsm_ies.E2SmRsmControlHeader{
		RsmCommand: command,
	}
}
