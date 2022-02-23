// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmKpmIndicationMessageFormat1(subscriptionID int64, measData *e2sm_kpm_v2_go.MeasurementData) (*e2sm_kpm_v2_go.E2SmKpmIndicationMessage, error) {

	e2SmKpmPdu := e2sm_kpm_v2_go.E2SmKpmIndicationMessage{
		IndicationMessageFormats: &e2sm_kpm_v2_go.IndicationMessageFormats{
			E2SmKpmIndicationMessage: &e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat1{
				IndicationMessageFormat1: &e2sm_kpm_v2_go.E2SmKpmIndicationMessageFormat1{
					SubscriptId: &e2sm_kpm_v2_go.SubscriptionId{
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

func CreateE2SmKpmIndicationMessageFormat2(subscriptionID int64, measCondUEList *e2sm_kpm_v2_go.MeasurementCondUeidList,
	measData *e2sm_kpm_v2_go.MeasurementData) (*e2sm_kpm_v2_go.E2SmKpmIndicationMessage, error) {

	e2SmKpmPdu := e2sm_kpm_v2_go.E2SmKpmIndicationMessage{
		IndicationMessageFormats: &e2sm_kpm_v2_go.IndicationMessageFormats{
			E2SmKpmIndicationMessage: &e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat2{
				IndicationMessageFormat2: &e2sm_kpm_v2_go.E2SmKpmIndicationMessageFormat2{
					SubscriptId: &e2sm_kpm_v2_go.SubscriptionId{
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

func CreateMeasurementRecordItemInteger(integer int64) *e2sm_kpm_v2_go.MeasurementRecordItem {

	return &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Integer{
			Integer: integer,
		},
	}
}

func CreateMeasurementRecordItemReal(real float64) *e2sm_kpm_v2_go.MeasurementRecordItem {

	return &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_Real{
			Real: real,
		},
	}
}

func CreateMeasurementRecordItemNoValue() *e2sm_kpm_v2_go.MeasurementRecordItem {

	return &e2sm_kpm_v2_go.MeasurementRecordItem{
		MeasurementRecordItem: &e2sm_kpm_v2_go.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
}

func CreateMeasurementCondUEIDItem(measType *e2sm_kpm_v2_go.MeasurementType, mc *e2sm_kpm_v2_go.MatchingCondList) (*e2sm_kpm_v2_go.MeasurementCondUeidItem, error) {

	measCondUEIDItem := e2sm_kpm_v2_go.MeasurementCondUeidItem{
		MeasType:     measType,
		MatchingCond: mc,
	}

	if err := measCondUEIDItem.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementCondUEIDItem(): error validating MeasurementCondUeidItem %s", err.Error())
	}
	return &measCondUEIDItem, nil
}

func CreateMatchingUEIDItem(ueID []byte) (*e2sm_kpm_v2_go.MatchingUeidItem, error) {

	mueIDi := e2sm_kpm_v2_go.MatchingUeidItem{
		UeId: &e2sm_kpm_v2_go.UeIdentity{
			Value: ueID,
		},
	}

	if err := mueIDi.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMatchingUEIDItem(): error validating MatchingUeidItem %s", err.Error())
	}
	return &mueIDi, nil
}

func CreateMeasurementDataItem(mr *e2sm_kpm_v2_go.MeasurementRecord) (*e2sm_kpm_v2_go.MeasurementDataItem, error) {

	mdi := e2sm_kpm_v2_go.MeasurementDataItem{
		MeasRecord: mr,
	}

	if err := mdi.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementDataItem(): error validating MatchingUeidItem %s", err.Error())
	}
	return &mdi, nil
}
