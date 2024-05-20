// SPDX-License-Identifier: Apache-2.0
// Copyright 2023-present Intel Corporation

package pdubuilder

import (
	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateListOfCellsForRanfunctionDefinition(value []*e2smcccv1.CellForRanfunctionDefinition) (*e2smcccv1.ListOfCellsForRanfunctionDefinition, error) {

	msg := &e2smcccv1.ListOfCellsForRanfunctionDefinition{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfCellsForRanfunctionDefinition() error validating PDU %s", err.Error())
	}

	return msg, nil
}
