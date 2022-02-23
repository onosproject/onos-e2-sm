// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"fmt"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
)

func CreateE2SmRsmControlHeader(command e2sm_rsm_ies.E2SmRsmCommand) (*e2sm_rsm_ies.E2SmRsmControlHeader, error) {

	ch := &e2sm_rsm_ies.E2SmRsmControlHeader{
		RsmCommand: command,
	}

	if err := ch.Validate(); err != nil {
		return nil, fmt.Errorf("CreateE2SmRsmControlHeader(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return ch, nil
}
