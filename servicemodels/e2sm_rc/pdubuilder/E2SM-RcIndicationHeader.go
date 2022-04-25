//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRcIndicationHeaderFormat1() (*e2smrcv1.E2SmRcIndicationHeader, error) {

	msg := &e2smrcv1.E2SmRcIndicationHeader{
		RicIndicationHeaderFormats: &e2smrcv1.RicIndicationHeaderFormats{
			RicIndicationHeaderFormats: &e2smrcv1.RicIndicationHeaderFormats_IndicationHeaderFormat1{
				IndicationHeaderFormat1: &e2smrcv1.E2SmRcIndicationHeaderFormat1{
					RicEventTriggerConditionId: &e2smrcv1.RicEventTriggerConditionId{
						//the only item is OPTIONAL
					},
				},
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationHeaderFormat1() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationHeaderFormat2(ueID *e2smcommonies.Ueid, ricInsertStyleType int32, ricInsertIndicationID int32) (*e2smrcv1.E2SmRcIndicationHeader, error) {

	msg := &e2smrcv1.E2SmRcIndicationHeader{
		RicIndicationHeaderFormats: &e2smrcv1.RicIndicationHeaderFormats{
			RicIndicationHeaderFormats: &e2smrcv1.RicIndicationHeaderFormats_IndicationHeaderFormat2{
				IndicationHeaderFormat2: &e2smrcv1.E2SmRcIndicationHeaderFormat2{
					UeId: ueID,
					RicInsertStyleType: &e2smcommonies.RicStyleType{
						Value: ricInsertStyleType,
					},
					RicInsertIndicationId: &e2smrcv1.RicInsertIndicationId{
						Value: ricInsertIndicationID,
					},
				},
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationHeaderFormat2() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}
