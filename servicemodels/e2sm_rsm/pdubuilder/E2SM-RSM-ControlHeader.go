// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrsm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRsmControlHeader(command e2smrsm.E2SmRsmCommand) (*e2smrsm.E2SmRsmControlHeader, error) {

	ch := &e2smrsm.E2SmRsmControlHeader{
		RsmCommand: command,
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmControlHeader(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return ch, nil
}
