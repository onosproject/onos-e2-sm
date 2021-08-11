// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import "C"
import (
	"fmt"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
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

	//if err := e2SmKpmPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmPDU %s", err.Error())
	//}
	return &e2SmKpmPdu, nil
}

func SetMeasInfoList(im *e2sm_kpm_v2_go.E2SmKpmIndicationMessage, measInfoList *e2sm_kpm_v2_go.MeasurementInfoList) (*e2sm_kpm_v2_go.E2SmKpmIndicationMessage, error) {
	switch choice := im.IndicationMessageFormats.E2SmKpmIndicationMessage.(type) {
	case *e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().MeasInfoList = measInfoList
	// Left for future extensions
	//case *e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat2:
	//	im.GetIndicationMessageFormats().GetIndicationMessageFormat2().MeasInfoList = measInfoList
	default:
		return nil, fmt.Errorf("unexpected IndicationMessage format %v", choice)
	}

	//if err := im.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmPDU %s", err.Error())
	//}
	return im, nil
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

	//if err := e2SmKpmPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmPDU %s", err.Error())
	//}
	return &e2SmKpmPdu, nil
}

func SetGranularityPeriod(im *e2sm_kpm_v2_go.E2SmKpmIndicationMessage, gp int64) (*e2sm_kpm_v2_go.E2SmKpmIndicationMessage, error) {
	switch choice := im.IndicationMessageFormats.E2SmKpmIndicationMessage.(type) {
	case *e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GranulPeriod =  &e2sm_kpm_v2_go.GranularityPeriod{
			Value: gp,
		}
	case *e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat2:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GranulPeriod =  &e2sm_kpm_v2_go.GranularityPeriod{
			Value: gp,
		}
	default:
		return nil, fmt.Errorf("unexpected IndicationMessage format %v", choice)
	}

	//if err := im.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmPDU %s", err.Error())
	//}
	return im, nil
}

func SetCellObjectID(im *e2sm_kpm_v2_go.E2SmKpmIndicationMessage, cellObjID string) (*e2sm_kpm_v2_go.E2SmKpmIndicationMessage, error) {
	switch choice := im.IndicationMessageFormats.E2SmKpmIndicationMessage.(type) {
	case *e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().CellObjId =  &e2sm_kpm_v2_go.CellObjectId{
			Value: cellObjID,
		}
	case *e2sm_kpm_v2_go.IndicationMessageFormats_IndicationMessageFormat2:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat2().CellObjId =  &e2sm_kpm_v2_go.CellObjectId{
			Value: cellObjID,
		}
	default:
		return nil, fmt.Errorf("unexpected IndicationMessage format %v", choice)
	}

	//if err := im.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmPDU %s", err.Error())
	//}
	return im, nil
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

	//if err := measCondUEIDItem.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating MeasurementCondUeidItem %s", err.Error())
	//}
	return &measCondUEIDItem, nil
}

func SetMatchingUEUDlist(mci *e2sm_kpm_v2_go.MeasurementCondUeidItem, ml *e2sm_kpm_v2_go.MatchingUeidList) *e2sm_kpm_v2_go.MeasurementCondUeidItem {
	mci.MatchingUeidList = ml
	return mci
}

func CreateMatchingUEIDItem(ueID []byte) (*e2sm_kpm_v2_go.MatchingUeidItem, error) {

	mueIDi := e2sm_kpm_v2_go.MatchingUeidItem{
		UeId: &e2sm_kpm_v2_go.UeIdentity{
			Value: ueID,
		},
	}

	//if err := mueIDi.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating MatchingUeidItem %s", err.Error())
	//}
	return &mueIDi, nil
}

func CreateMeasurementDataItem(mr *e2sm_kpm_v2_go.MeasurementRecord) (*e2sm_kpm_v2_go.MeasurementDataItem, error) {

	mdi := e2sm_kpm_v2_go.MeasurementDataItem{
		MeasRecord: mr,
	}

	//if err := mdi.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating MatchingUeidItem %s", err.Error())
	//}
	return &mdi, nil
}

func SetIncompleteFlag(mdi *e2sm_kpm_v2_go.MeasurementDataItem) *e2sm_kpm_v2_go.MeasurementDataItem {
	incf := e2sm_kpm_v2_go.IncompleteFlag_INCOMPLETE_FLAG_TRUE
	mdi.IncompleteFlag = &incf
	return mdi
}
