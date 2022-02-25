// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmKpmIndicationMessageFormat1(subscriptionID int64, measData *e2smkpmv2.MeasurementData) (*e2smkpmv2.E2SmKpmIndicationMessage, error) {

	e2SmKpmPdu := e2smkpmv2.E2SmKpmIndicationMessage{
		IndicationMessageFormats: &e2smkpmv2.IndicationMessageFormats{
			E2SmKpmIndicationMessage: &e2smkpmv2.IndicationMessageFormats_IndicationMessageFormat1{
				IndicationMessageFormat1: &e2smkpmv2.E2SmKpmIndicationMessageFormat1{
					SubscriptId: &e2smkpmv2.SubscriptionId{
						Value: subscriptionID,
					},
					MeasData: measData,
				},
			},
		},
	}

	if err := e2SmKpmPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmKpmIndicationMessageFormat1(): error validating E2SmKpmPDU %s", err.Error())
	}
	return &e2SmKpmPdu, nil
}

func CreateE2SmKpmIndicationMessageFormat2(subscriptionID int64, measCondUEList *e2smkpmv2.MeasurementCondUeidList,
	measData *e2smkpmv2.MeasurementData) (*e2smkpmv2.E2SmKpmIndicationMessage, error) {

	e2SmKpmPdu := e2smkpmv2.E2SmKpmIndicationMessage{
		IndicationMessageFormats: &e2smkpmv2.IndicationMessageFormats{
			E2SmKpmIndicationMessage: &e2smkpmv2.IndicationMessageFormats_IndicationMessageFormat2{
				IndicationMessageFormat2: &e2smkpmv2.E2SmKpmIndicationMessageFormat2{
					SubscriptId: &e2smkpmv2.SubscriptionId{
						Value: subscriptionID,
					},
					MeasCondUeidList: measCondUEList,
					MeasData:         measData,
				},
			},
		},
	}

	if err := e2SmKpmPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SmKpmPDU %s", err.Error())
	}
	return &e2SmKpmPdu, nil
}

func CreateMeasurementRecordItemInteger(integer int64) *e2smkpmv2.MeasurementRecordItem {

	return &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Integer{
			Integer: integer,
		},
	}
}

func CreateMeasurementRecordItemReal(real float64) *e2smkpmv2.MeasurementRecordItem {

	return &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_Real{
			Real: real,
		},
	}
}

func CreateMeasurementRecordItemNoValue() *e2smkpmv2.MeasurementRecordItem {

	return &e2smkpmv2.MeasurementRecordItem{
		MeasurementRecordItem: &e2smkpmv2.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
}

func CreateMeasurementCondUEIDItem(measType *e2smkpmv2.MeasurementType, mc *e2smkpmv2.MatchingCondList) (*e2smkpmv2.MeasurementCondUeidItem, error) {

	measCondUEIDItem := e2smkpmv2.MeasurementCondUeidItem{
		MeasType:     measType,
		MatchingCond: mc,
	}

	if err := measCondUEIDItem.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementCondUEIDItem(): error validating MeasurementCondUeidItem %s", err.Error())
	}
	return &measCondUEIDItem, nil
}

func CreateMatchingUEIDItem(ueID []byte) (*e2smkpmv2.MatchingUeidItem, error) {

	mueIDi := e2smkpmv2.MatchingUeidItem{
		UeId: &e2smkpmv2.UeIdentity{
			Value: ueID,
		},
	}

	if err := mueIDi.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMatchingUEIDItem(): error validating MatchingUeidItem %s", err.Error())
	}
	return &mueIDi, nil
}

func CreateMeasurementDataItem(mr *e2smkpmv2.MeasurementRecord) (*e2smkpmv2.MeasurementDataItem, error) {

	mdi := e2smkpmv2.MeasurementDataItem{
		MeasRecord: mr,
	}

	if err := mdi.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementDataItem(): error validating MatchingUeidItem %s", err.Error())
	}
	return &mdi, nil
}
