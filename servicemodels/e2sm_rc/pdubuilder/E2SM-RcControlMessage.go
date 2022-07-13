//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
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

func CreateE2SmRcControlMessageFormat2(ricControlStyleList []*e2smrcv1.E2SmRcControlMessageFormat2StyleItem) (*e2smrcv1.E2SmRcControlMessage, error) {

	ch := &e2smrcv1.E2SmRcControlMessage{
		RicControlMessageFormats: &e2smrcv1.RicControlMessageFormats{
			RicControlMessageFormats: &e2smrcv1.RicControlMessageFormats_ControlMessageFormat2{
				ControlMessageFormat2: &e2smrcv1.E2SmRcControlMessageFormat2{
					RicControlStyleList: ricControlStyleList,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlMessageFormat2() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcControlMessageFormat2StyleItem(ricStyleType int32, ricControlActionList []*e2smrcv1.E2SmRcControlMessageFormat2ControlActionItem) (*e2smrcv1.E2SmRcControlMessageFormat2StyleItem, error) {

	item := &e2smrcv1.E2SmRcControlMessageFormat2StyleItem{
		IndicatedControlStyleType: &e2smcommonies.RicStyleType{
			Value: ricStyleType,
		},
		RicControlActionList: ricControlActionList,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlMessageFormat2StyleItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateE2SmRcControlMessageFormat2ControlActionItem(ricControlActionID int32, ranPList *e2smrcv1.E2SmRcControlMessageFormat1) (*e2smrcv1.E2SmRcControlMessageFormat2ControlActionItem, error) {

	item := &e2smrcv1.E2SmRcControlMessageFormat2ControlActionItem{
		RicControlActionId: &e2smrcv1.RicControlActionId{
			Value: ricControlActionID,
		},
		RanPList: ranPList,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlMessageFormat2ControlActionItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}
