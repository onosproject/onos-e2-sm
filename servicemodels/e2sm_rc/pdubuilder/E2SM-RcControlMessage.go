//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcControlMessageFormat1(rpl []*e2smrcv1.E2SmRcControlMessageFormat1Item) (*e2smrcv1.E2SmRcControlMessage, error) {

	ch := &e2smrcv1.E2SmRcControlMessage{
		RicControlMessageFormats: &e2smrcv1.RicControlMessageFormats{
			RicControlMessageFormats: &e2smrcv1.RicControlMessageFormats_ControlMessageFormat1{
				ControlMessageFormat1: &e2smrcv1.E2SmRcControlMessageFormat1{
					RanPList: rpl,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlMessageFormat1() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcControlMessageFormat1Item(rpID int64, rpvt *e2smrcv1.RanparameterValueType) (*e2smrcv1.E2SmRcControlMessageFormat1Item, error) {

	item := &e2smrcv1.E2SmRcControlMessageFormat1Item{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: rpID,
		},
		RanParameterValueType: rpvt,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlMessageFormat1Item() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}
