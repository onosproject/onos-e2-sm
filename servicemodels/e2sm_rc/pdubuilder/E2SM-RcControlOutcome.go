//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcControlOutcomeFormat1(rpl []*e2smrcv1.E2SmRcControlOutcomeFormat1Item) (*e2smrcv1.E2SmRcControlOutcome, error) {

	ch := &e2smrcv1.E2SmRcControlOutcome{
		RicControlOutcomeFormats: &e2smrcv1.RicControlOutcomeFormats{
			RicControlOutcomeFormats: &e2smrcv1.RicControlOutcomeFormats_ControlOutcomeFormat1{
				ControlOutcomeFormat1: &e2smrcv1.E2SmRcControlOutcomeFormat1{
					RanPList: rpl,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat1() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcControlOutcomeFormat1Item(rpID int64, rpv *e2smrcv1.RanparameterValue) (*e2smrcv1.E2SmRcControlOutcomeFormat1Item, error) {

	item := &e2smrcv1.E2SmRcControlOutcomeFormat1Item{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: rpID,
		},
		RanParameterValue: rpv,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat1Item() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateE2SmRcControlOutcomeFormat2(ricControlStyleList []*e2smrcv1.E2SmRcControlOutcomeFormat2StyleItem) (*e2smrcv1.E2SmRcControlOutcome, error) {

	ch := &e2smrcv1.E2SmRcControlOutcome{
		RicControlOutcomeFormats: &e2smrcv1.RicControlOutcomeFormats{
			RicControlOutcomeFormats: &e2smrcv1.RicControlOutcomeFormats_ControlOutcomeFormat2{
				ControlOutcomeFormat2: &e2smrcv1.E2SmRcControlOutcomeFormat2{
					RicControlStyleList: ricControlStyleList,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat2() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcControlOutcomeFormat2StyleItem(ricStyleType int32, ricControlOutcomeList []*e2smrcv1.E2SmRcControlOutcomeFormat2ControlOutcomeItem) (*e2smrcv1.E2SmRcControlOutcomeFormat2StyleItem, error) {

	item := &e2smrcv1.E2SmRcControlOutcomeFormat2StyleItem{
		IndicatedControlStyleType: &e2smcommonies.RicStyleType{
			Value: ricStyleType,
		},
		RicControlOutcomeList: ricControlOutcomeList,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat2StyleItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateE2SmRcControlOutcomeFormat2ControlOutcomeItem(ricControlActionID int32, ranPList []*e2smrcv1.E2SmRcControlOutcomeFormat2RanpItem) (*e2smrcv1.E2SmRcControlOutcomeFormat2ControlOutcomeItem, error) {

	item := &e2smrcv1.E2SmRcControlOutcomeFormat2ControlOutcomeItem{
		RicControlActionId: &e2smrcv1.RicControlActionId{
			Value: ricControlActionID,
		},
		RanPList: ranPList,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat2ControlOutcomeItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateE2SmRcControlOutcomeFormat2RanpItem(ranParameterID int64, ranParameterValue *e2smrcv1.RanparameterValue) (*e2smrcv1.E2SmRcControlOutcomeFormat2RanpItem, error) {

	item := &e2smrcv1.E2SmRcControlOutcomeFormat2RanpItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterValue: ranParameterValue,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat2RanpItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateE2SmRcControlOutcomeFormat3(ranPList []*e2smrcv1.E2SmRcControlOutcomeFormat3Item) (*e2smrcv1.E2SmRcControlOutcome, error) {

	ch := &e2smrcv1.E2SmRcControlOutcome{
		RicControlOutcomeFormats: &e2smrcv1.RicControlOutcomeFormats{
			RicControlOutcomeFormats: &e2smrcv1.RicControlOutcomeFormats_ControlOutcomeFormat3{
				ControlOutcomeFormat3: &e2smrcv1.E2SmRcControlOutcomeFormat3{
					RanPList: ranPList,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat3() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcControlOutcomeFormat3Item(ranParameterID int64, ranParametervalueType *e2smrcv1.RanparameterValueType) (*e2smrcv1.E2SmRcControlOutcomeFormat3Item, error) {

	item := &e2smrcv1.E2SmRcControlOutcomeFormat3Item{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterValueType: ranParametervalueType,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat3Item() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}
