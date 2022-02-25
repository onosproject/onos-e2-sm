// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmKpmActionDefinitionFormat1(ricStyleType int32,
	actionDefinition *e2smkpmv2.E2SmKpmActionDefinitionFormat1) (*e2smkpmv2.E2SmKpmActionDefinition, error) {

	e2SmKpmPdu := e2smkpmv2.E2SmKpmActionDefinition{
		RicStyleType: &e2smkpmv2.RicStyleType{
			Value: ricStyleType,
		},
		ActionDefinitionFormats: &e2smkpmv2.ActionDefinitionFormats{
			E2SmKpmActionDefinition: &e2smkpmv2.ActionDefinitionFormats_ActionDefinitionFormat1{
				ActionDefinitionFormat1: actionDefinition,
			},
		},
	}

	if err := e2SmKpmPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmKpmActionDefinitionFormat1(): error validating E2SmKpmActionDefinition %s", err.Error())
	}
	return &e2SmKpmPdu, nil
}

func CreateE2SmKpmActionDefinitionFormat2(ricStyleType int32,
	actionDefinitionFormat2 *e2smkpmv2.E2SmKpmActionDefinitionFormat2) (*e2smkpmv2.E2SmKpmActionDefinition, error) {

	e2SmKpmPdu := e2smkpmv2.E2SmKpmActionDefinition{
		RicStyleType: &e2smkpmv2.RicStyleType{
			Value: ricStyleType,
		},
		ActionDefinitionFormats: &e2smkpmv2.ActionDefinitionFormats{
			E2SmKpmActionDefinition: &e2smkpmv2.ActionDefinitionFormats_ActionDefinitionFormat2{
				ActionDefinitionFormat2: actionDefinitionFormat2,
			},
		},
	}

	if err := e2SmKpmPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmKpmActionDefinitionFormat2(): error validating E2SmKpmActionDefinition %s", err.Error())
	}
	return &e2SmKpmPdu, nil
}

func CreateE2SmKpmActionDefinitionFormat3(ricStyleType int32,
	actionDefinitionFormat3 *e2smkpmv2.E2SmKpmActionDefinitionFormat3) (*e2smkpmv2.E2SmKpmActionDefinition, error) {

	e2SmKpmPdu := e2smkpmv2.E2SmKpmActionDefinition{
		RicStyleType: &e2smkpmv2.RicStyleType{
			Value: ricStyleType,
		},
		ActionDefinitionFormats: &e2smkpmv2.ActionDefinitionFormats{
			E2SmKpmActionDefinition: &e2smkpmv2.ActionDefinitionFormats_ActionDefinitionFormat3{
				ActionDefinitionFormat3: actionDefinitionFormat3,
			},
		},
	}

	if err := e2SmKpmPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmKpmActionDefinitionFormat3(): error validating E2SmKpmActionDefinition %s", err.Error())
	}
	return &e2SmKpmPdu, nil
}

func CreateActionDefinitionFormat1(cellObjID string, measInfoList *e2smkpmv2.MeasurementInfoList,
	granularity int64, subID int64) (*e2smkpmv2.E2SmKpmActionDefinitionFormat1, error) {

	actionDefinitionFormat1 := e2smkpmv2.E2SmKpmActionDefinitionFormat1{
		CellObjId: &e2smkpmv2.CellObjectId{
			Value: cellObjID,
		},
		MeasInfoList: measInfoList,
		GranulPeriod: &e2smkpmv2.GranularityPeriod{
			Value: granularity,
		},
		SubscriptId: &e2smkpmv2.SubscriptionId{
			Value: subID,
		},
	}

	if err := actionDefinitionFormat1.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateActionDefinitionFormat1(): error validating E2SmKpmActionDefinitionFormat1 %s", err.Error())
	}
	return &actionDefinitionFormat1, nil
}

func CreateActionDefinitionFormat2(ueID []byte, actionDefinitionFormat1 *e2smkpmv2.E2SmKpmActionDefinitionFormat1) (*e2smkpmv2.E2SmKpmActionDefinitionFormat2, error) {

	actionDefinitionFormat2 := e2smkpmv2.E2SmKpmActionDefinitionFormat2{
		UeId: &e2smkpmv2.UeIdentity{
			Value: ueID,
		},
		SubscriptInfo: actionDefinitionFormat1,
	}

	if err := actionDefinitionFormat2.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateActionDefinitionFormat2(): error validating E2SmKpmActionDefinitionFormat2 %s", err.Error())
	}
	return &actionDefinitionFormat2, nil
}

func CreateActionDefinitionFormat3(cellObjID string, measCondList *e2smkpmv2.MeasurementCondList,
	granularity int64, subID int64) (*e2smkpmv2.E2SmKpmActionDefinitionFormat3, error) {

	actionDefinitionFormat3 := e2smkpmv2.E2SmKpmActionDefinitionFormat3{
		CellObjId: &e2smkpmv2.CellObjectId{
			Value: cellObjID,
		},
		MeasCondList: measCondList,
		GranulPeriod: &e2smkpmv2.GranularityPeriod{
			Value: granularity,
		},
		SubscriptId: &e2smkpmv2.SubscriptionId{
			Value: subID,
		},
	}

	if err := actionDefinitionFormat3.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateActionDefinitionFormat3(): error validating E2SmKpmActionDefinitionFormat3 %s", err.Error())
	}
	return &actionDefinitionFormat3, nil
}

func CreateMeasurementInfoItem(measType *e2smkpmv2.MeasurementType) (*e2smkpmv2.MeasurementInfoItem, error) {

	item := e2smkpmv2.MeasurementInfoItem{
		MeasType: measType,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementInfoItem(): error validating MeasurementInfoItem %s", err.Error())
	}
	return &item, nil
}

func CreateMeasurementTypeMeasID(measTypeID int32) (*e2smkpmv2.MeasurementType, error) {
	measType := e2smkpmv2.MeasurementType{
		MeasurementType: &e2smkpmv2.MeasurementType_MeasId{
			MeasId: &e2smkpmv2.MeasurementTypeId{
				Value: measTypeID,
			},
		},
	}

	if err := measType.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementTypeMeasID(): error validating MeasurementType %s", err.Error())
	}
	return &measType, nil
}

func CreateMeasurementTypeMeasName(measName string) (*e2smkpmv2.MeasurementType, error) {
	measType := e2smkpmv2.MeasurementType{
		MeasurementType: &e2smkpmv2.MeasurementType_MeasName{
			MeasName: &e2smkpmv2.MeasurementTypeName{
				Value: measName,
			},
		},
	}

	if err := measType.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementTypeMeasName(): error validating MeasurementType %s", err.Error())
	}
	return &measType, nil
}

func CreateSnssai(sst []byte) (*e2smkpmv2.Snssai, error) {

	if len(sst) != 1 {
		return nil, errors.NewInvalid("SST should be of length 1 byte, got %d", len(sst))
	}
	return &e2smkpmv2.Snssai{
		SSt: sst,
	}, nil
}

func CreateLabelInfoItem(plmnID []byte, sst []byte, sd []byte, fiveQI *int32, qfi *int32, qci *int32, qciMax *int32,
	qciMin *int32, arpMax *int32, arpMin *int32, br *int32, lmm *int32, sum *e2smkpmv2.SUM, dbx *int32, dby *int32,
	dbz *int32, plo *e2smkpmv2.PreLabelOverride, seind *e2smkpmv2.StartEndInd) (*e2smkpmv2.LabelInfoItem, error) {

	labelInfoItem := e2smkpmv2.LabelInfoItem{
		MeasLabel: &e2smkpmv2.MeasurementLabel{},
	}

	if plmnID != nil {
		if len(plmnID) != 3 {
			return nil, errors.NewInvalid("error: Plmn ID should be 3 chars")
		}
		labelInfoItem.MeasLabel.PlmnId = &e2smkpmv2.PlmnIdentity{
			Value: plmnID,
		}
	}
	if sst != nil {
		if len(sst) != 1 {
			return nil, errors.NewInvalid("error: SST should be 1 char")
		}
		labelInfoItem.MeasLabel.SliceId = &e2smkpmv2.Snssai{
			SSt: sst,
		}
		if sd != nil {
			if len(sd) != 3 {
				return nil, errors.NewInvalid("error: SD should be 3 chars")
			}
			labelInfoItem.MeasLabel.SliceId.SD = sd
		}
	}
	if fiveQI != nil {
		labelInfoItem.MeasLabel.FiveQi = &e2smkpmv2.FiveQi{
			Value: *fiveQI,
		}
	}
	if qfi != nil {
		labelInfoItem.MeasLabel.QFi = &e2smkpmv2.Qfi{
			Value: *qfi,
		}
	}
	if qci != nil {
		labelInfoItem.MeasLabel.QCi = &e2smkpmv2.Qci{
			Value: *qci,
		}
	}
	if qciMax != nil {
		labelInfoItem.MeasLabel.QCimax = &e2smkpmv2.Qci{
			Value: *qciMax,
		}
	}
	if qciMin != nil {
		labelInfoItem.MeasLabel.QCimin = &e2smkpmv2.Qci{
			Value: *qciMin,
		}
	}
	if arpMax != nil {
		labelInfoItem.MeasLabel.ARpmax = &e2smkpmv2.Arp{
			Value: *arpMax,
		}
	}
	if arpMin != nil {
		labelInfoItem.MeasLabel.ARpmin = &e2smkpmv2.Arp{
			Value: *arpMin,
		}
	}
	if br != nil {
		labelInfoItem.MeasLabel.BitrateRange = br
	}
	if lmm != nil {
		labelInfoItem.MeasLabel.LayerMuMimo = lmm
	}
	if sum != nil {
		labelInfoItem.MeasLabel.SUm = sum
	}
	if dbx != nil {
		labelInfoItem.MeasLabel.DistBinX = dbx
	}
	if dby != nil {
		labelInfoItem.MeasLabel.DistBinY = dby
	}
	if dbz != nil {
		labelInfoItem.MeasLabel.DistBinZ = dbz
	}
	if plo != nil {
		labelInfoItem.MeasLabel.PreLabelOverride = plo
	}
	if seind != nil {
		labelInfoItem.MeasLabel.StartEndInd = seind
	}

	if err := labelInfoItem.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateLabelInfoItem(): error validating LabelInfoItem %s", err.Error())
	}
	return &labelInfoItem, nil
}

func CreateMeasurementLabelEmpty() *e2smkpmv2.MeasurementLabel {
	return new(e2smkpmv2.MeasurementLabel)
}

func CreateMeasurementCondItem(measType *e2smkpmv2.MeasurementType, measCondList *e2smkpmv2.MatchingCondList) (*e2smkpmv2.MeasurementCondItem, error) {

	measCondItem := e2smkpmv2.MeasurementCondItem{
		MeasType:     measType,
		MatchingCond: measCondList,
	}

	if err := measCondItem.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementCondItem(): error validating MeasurementCondItem %s", err.Error())
	}
	return &measCondItem, nil
}

func CreateMatchingCondItemMeasLabel(measLabel *e2smkpmv2.MeasurementLabel) (*e2smkpmv2.MatchingCondItem, error) {

	res := e2smkpmv2.MatchingCondItem{
		MatchingCondItem: &e2smkpmv2.MatchingCondItem_MeasLabel{
			MeasLabel: measLabel,
		},
	}

	if err := res.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMatchingCondItemMeasLabel(): error validating MatchingCondItem (MeasLabel) %s", err.Error())
	}
	return &res, nil
}

func CreateMatchingCondItemTestCondInfo(testCondInfo *e2smkpmv2.TestCondInfo) (*e2smkpmv2.MatchingCondItem, error) {

	res := e2smkpmv2.MatchingCondItem{
		MatchingCondItem: &e2smkpmv2.MatchingCondItem_TestCondInfo{
			TestCondInfo: testCondInfo,
		},
	}

	if err := res.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMatchingCondItemTestCondInfo(): error validating MatchingCondItem (TestCondInfo) %s", err.Error())
	}
	return &res, nil
}

func CreateTestCondInfo(tct *e2smkpmv2.TestCondType, tce e2smkpmv2.TestCondExpression, tcv *e2smkpmv2.TestCondValue) (*e2smkpmv2.TestCondInfo, error) {

	tci := e2smkpmv2.TestCondInfo{
		TestValue: tcv,
		TestExpr:  tce,
		TestType:  tct,
	}

	if err := tci.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTestCondInfo(): error validating TestCondInfo (TestCondInfo) %s", err.Error())
	}
	return &tci, nil
}

func CreateTestCondTypeGBR() *e2smkpmv2.TestCondType {

	return &e2smkpmv2.TestCondType{
		TestCondType: &e2smkpmv2.TestCondType_GBr{
			GBr: e2smkpmv2.GBR_GBR_TRUE,
		},
	}
}

func CreateTestCondTypeAMBR() *e2smkpmv2.TestCondType {

	return &e2smkpmv2.TestCondType{
		TestCondType: &e2smkpmv2.TestCondType_AMbr{
			AMbr: e2smkpmv2.AMBR_AMBR_TRUE,
		},
	}
}

func CreateTestCondTypeIsStat() *e2smkpmv2.TestCondType {

	return &e2smkpmv2.TestCondType{
		TestCondType: &e2smkpmv2.TestCondType_IsStat{
			IsStat: e2smkpmv2.ISSTAT_ISSTAT_TRUE,
		},
	}
}

func CreateTestCondTypeIsCatM() *e2smkpmv2.TestCondType {

	return &e2smkpmv2.TestCondType{
		TestCondType: &e2smkpmv2.TestCondType_IsCatM{
			IsCatM: e2smkpmv2.ISCATM_ISCATM_TRUE,
		},
	}
}

func CreateTestCondTypeRSRP() *e2smkpmv2.TestCondType {

	return &e2smkpmv2.TestCondType{
		TestCondType: &e2smkpmv2.TestCondType_RSrp{
			RSrp: e2smkpmv2.RSRP_RSRP_TRUE,
		},
	}
}

func CreateTestCondTypeRSRQ() *e2smkpmv2.TestCondType {

	return &e2smkpmv2.TestCondType{
		TestCondType: &e2smkpmv2.TestCondType_RSrq{
			RSrq: e2smkpmv2.RSRQ_RSRQ_TRUE,
		},
	}
}

func CreateTestCondValueInt(val int64) *e2smkpmv2.TestCondValue {

	return &e2smkpmv2.TestCondValue{
		TestCondValue: &e2smkpmv2.TestCondValue_ValueInt{
			ValueInt: val,
		},
	}
}

func CreateTestCondValueEnum(val int64) *e2smkpmv2.TestCondValue {

	return &e2smkpmv2.TestCondValue{
		TestCondValue: &e2smkpmv2.TestCondValue_ValueEnum{
			ValueEnum: val,
		},
	}
}

func CreateTestCondValueBool(val bool) *e2smkpmv2.TestCondValue {

	return &e2smkpmv2.TestCondValue{
		TestCondValue: &e2smkpmv2.TestCondValue_ValueBool{
			ValueBool: val,
		},
	}
}

func CreateTestCondValueBitS(val *asn1.BitString) *e2smkpmv2.TestCondValue {

	return &e2smkpmv2.TestCondValue{
		TestCondValue: &e2smkpmv2.TestCondValue_ValueBitS{
			ValueBitS: val,
		},
	}
}

func CreateTestCondValueOctS(val []byte) *e2smkpmv2.TestCondValue {

	return &e2smkpmv2.TestCondValue{
		TestCondValue: &e2smkpmv2.TestCondValue_ValueOctS{
			ValueOctS: val,
		},
	}
}

func CreateTestCondValuePrtS(val string) *e2smkpmv2.TestCondValue {

	return &e2smkpmv2.TestCondValue{
		TestCondValue: &e2smkpmv2.TestCondValue_ValuePrtS{
			ValuePrtS: val,
		},
	}
}
