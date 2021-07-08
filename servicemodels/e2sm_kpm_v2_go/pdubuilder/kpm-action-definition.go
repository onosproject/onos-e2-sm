// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmKpmActionDefinitionFormat1(ricStyleType int32,
	actionDefinition *e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat1) (*e2sm_kpm_v2_go.E2SmKpmActionDefinition, error) {

	e2SmKpmPdu := e2sm_kpm_v2_go.E2SmKpmActionDefinition{
		RicStyleType: &e2sm_kpm_v2_go.RicStyleType{
			Value: ricStyleType,
		},
		E2SmKpmActionDefinition: &e2sm_kpm_v2_go.E2SmKpmActionDefinition_ActionDefinitionFormat1{
			ActionDefinitionFormat1: actionDefinition,
		},
	}

	//if err := e2SmKpmPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinition %s", err.Error())
	//}
	return &e2SmKpmPdu, nil
}

func CreateE2SmKpmActionDefinitionFormat2(ricStyleType int32,
	actionDefinitionFormat2 *e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat2) (*e2sm_kpm_v2_go.E2SmKpmActionDefinition, error) {

	e2SmKpmPdu := e2sm_kpm_v2_go.E2SmKpmActionDefinition{
		RicStyleType: &e2sm_kpm_v2_go.RicStyleType{
			Value: ricStyleType,
		},
		E2SmKpmActionDefinition: &e2sm_kpm_v2_go.E2SmKpmActionDefinition_ActionDefinitionFormat2{
			ActionDefinitionFormat2: actionDefinitionFormat2,
		},
	}

	//if err := e2SmKpmPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinition %s", err.Error())
	//}
	return &e2SmKpmPdu, nil
}

func CreateE2SmKpmActionDefinitionFormat3(ricStyleType int32,
	actionDefinitionFormat3 *e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat3) (*e2sm_kpm_v2_go.E2SmKpmActionDefinition, error) {

	e2SmKpmPdu := e2sm_kpm_v2_go.E2SmKpmActionDefinition{
		RicStyleType: &e2sm_kpm_v2_go.RicStyleType{
			Value: ricStyleType,
		},
		E2SmKpmActionDefinition: &e2sm_kpm_v2_go.E2SmKpmActionDefinition_ActionDefinitionFormat3{
			ActionDefinitionFormat3: actionDefinitionFormat3,
		},
	}

	//if err := e2SmKpmPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinition %s", err.Error())
	//}
	return &e2SmKpmPdu, nil
}

func CreateActionDefinitionFormat1(cellObjID string, measInfoList *e2sm_kpm_v2_go.MeasurementInfoList,
	granularity int64, subID int64) (*e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat1, error) {

	actionDefinitionFormat1 := e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat1{
		CellObjId: &e2sm_kpm_v2_go.CellObjectId{
			Value: cellObjID,
		},
		MeasInfoList: measInfoList,
		GranulPeriod: &e2sm_kpm_v2_go.GranularityPeriod{
			Value: granularity,
		},
		SubscriptId: &e2sm_kpm_v2_go.SubscriptionId{
			Value: subID,
		},
	}

	//if err := actionDefinitionFormat1.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat1 %s", err.Error())
	//}

	return &actionDefinitionFormat1, nil
}

func CreateActionDefinitionFormat2(ueID []byte, actionDefinitionFormat1 *e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat1) (*e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat2, error) {

	actionDefinitionFormat2 := e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat2{
		UeId: &e2sm_kpm_v2_go.UeIdentity{
			Value: ueID,
		},
		SubscriptInfo: actionDefinitionFormat1,
	}

	//if err := actionDefinitionFormat2.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat2 %s", err.Error())
	//}

	return &actionDefinitionFormat2, nil
}

func CreateActionDefinitionFormat3(cellObjID string, measCondList *e2sm_kpm_v2_go.MeasurementCondList,
	granularity int64, subID int64) (*e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat3, error) {

	actionDefinitionFormat3 := e2sm_kpm_v2_go.E2SmKpmActionDefinitionFormat3{
		CellObjId: &e2sm_kpm_v2_go.CellObjectId{
			Value: cellObjID,
		},
		MeasCondList: measCondList,
		GranulPeriod: &e2sm_kpm_v2_go.GranularityPeriod{
			Value: granularity,
		},
		SubscriptId: &e2sm_kpm_v2_go.SubscriptionId{
			Value: subID,
		},
	}

	//if err := actionDefinitionFormat3.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmActionDefinitionFormat3 %s", err.Error())
	//}

	return &actionDefinitionFormat3, nil
}

func CreateMeasurementInfoItem(measType *e2sm_kpm_v2_go.MeasurementType, labelInfoList *e2sm_kpm_v2_go.LabelInfoList) (*e2sm_kpm_v2_go.MeasurementInfoItem, error) {

	item := e2sm_kpm_v2_go.MeasurementInfoItem{
		MeasType: measType,
	}

	// optional instance
	if labelInfoList != nil {
		item.LabelInfoList = labelInfoList
	}

	//if err := item.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating MeasurementInfoItem %s", err.Error())
	//}

	return &item, nil
}

func CreateMeasurementTypeMeasID(measTypeID int32) (*e2sm_kpm_v2_go.MeasurementType, error) {
	measType := e2sm_kpm_v2_go.MeasurementType{
		MeasurementType: &e2sm_kpm_v2_go.MeasurementType_MeasId{
			MeasId: &e2sm_kpm_v2_go.MeasurementTypeId{
				Value: measTypeID,
			},
		},
	}

	//if err := measType.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating MeasurementType %s", err.Error())
	//}

	return &measType, nil
}

func CreateMeasurementTypeMeasName(measName string) (*e2sm_kpm_v2_go.MeasurementType, error) {
	measType := e2sm_kpm_v2_go.MeasurementType{
		MeasurementType: &e2sm_kpm_v2_go.MeasurementType_MeasName{
			MeasName: &e2sm_kpm_v2_go.MeasurementTypeName{
				Value: measName,
			},
		},
	}

	//if err := measType.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating MeasurementType %s", err.Error())
	//}

	return &measType, nil
}

func CreateLabelInfoItem(plmnID []byte, sst []byte, sd []byte, fiveQI *int32, qfi *int32, qci *int32, qciMax *int32,
	qciMin *int32, arpMin *int32, arpMax *int32, br *int32, lmm *int32, sum *e2sm_kpm_v2_go.SUM, dbx *int32, dby *int32,
	dbz *int32, plo *e2sm_kpm_v2_go.PreLabelOverride, seind *e2sm_kpm_v2_go.StartEndInd) (*e2sm_kpm_v2_go.LabelInfoItem, error) {

	labelInfoItem := e2sm_kpm_v2_go.LabelInfoItem{
		MeasLabel: &e2sm_kpm_v2_go.MeasurementLabel{},
	}

	if plmnID != nil {
		if len(plmnID) != 3 {
			return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
		}
		labelInfoItem.MeasLabel.PlmnId = &e2sm_kpm_v2_go.PlmnIdentity{
			Value: plmnID,
		}
	}
	if sst != nil {
		if len(sst) != 1 {
			return nil, fmt.Errorf("error: SST should be 1 char")
		}
		labelInfoItem.MeasLabel.SliceId = &e2sm_kpm_v2_go.Snssai{
			SSt: sst,
		}
		if sd != nil {
			if len(sd) != 3 {
				return nil, fmt.Errorf("error: SD should be 3 chars")
			}
			labelInfoItem.MeasLabel.SliceId.SD = sd
		}
	}
	if fiveQI != nil {
		labelInfoItem.MeasLabel.FiveQi = &e2sm_kpm_v2_go.FiveQi{
			Value: *fiveQI,
		}
	}
	if qfi != nil {
		labelInfoItem.MeasLabel.QFi = &e2sm_kpm_v2_go.Qfi{
			Value: *qfi,
		}
	}
	if qci != nil {
		labelInfoItem.MeasLabel.QCi = &e2sm_kpm_v2_go.Qci{
			Value: *qci,
		}
	}
	if qciMax != nil {
		labelInfoItem.MeasLabel.QCimax = &e2sm_kpm_v2_go.Qci{
			Value: *qciMax,
		}
	}
	if qciMin != nil {
		labelInfoItem.MeasLabel.QCimin = &e2sm_kpm_v2_go.Qci{
			Value: *qciMin,
		}
	}
	if arpMin != nil {
		labelInfoItem.MeasLabel.ARpmin = &e2sm_kpm_v2_go.Arp{
			Value: *arpMin,
		}
	}
	if arpMax != nil {
		labelInfoItem.MeasLabel.ARpmax = &e2sm_kpm_v2_go.Arp{
			Value: *arpMax,
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

	//if err := labelInfoItem.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating LabelInfoItem %s", err.Error())
	//}

	return &labelInfoItem, nil
}
func CreateMeasurementCondItem(measType *e2sm_kpm_v2_go.MeasurementType, measCondList *e2sm_kpm_v2_go.MatchingCondList) (*e2sm_kpm_v2_go.MeasurementCondItem, error) {

	measCondItem := e2sm_kpm_v2_go.MeasurementCondItem{
		MeasType:     measType,
		MatchingCond: measCondList,
	}

	//if err := measCondItem.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating MeasurementCondItem %s", err.Error())
	//}
	return &measCondItem, nil
}

func CreateMatchingCondItemMeasLabel(measLabel *e2sm_kpm_v2_go.MeasurementLabel) (*e2sm_kpm_v2_go.MatchingCondItem, error) {

	res := e2sm_kpm_v2_go.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2_go.MatchingCondItem_MeasLabel{
			MeasLabel: measLabel,
		},
	}

	//if err := res.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating MatchingCondItem (MeasLabel) %s", err.Error())
	//}
	return &res, nil
}

func CreateMatchingCondItemTestCondInfo(testCondInfo *e2sm_kpm_v2_go.TestCondInfo) (*e2sm_kpm_v2_go.MatchingCondItem, error) {

	res := e2sm_kpm_v2_go.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2_go.MatchingCondItem_TestCondInfo{
			TestCondInfo: testCondInfo,
		},
	}

	//if err := res.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating MatchingCondItem (TestCondInfo) %s", err.Error())
	//}
	return &res, nil
}

func CreateTestCondInfo(tct *e2sm_kpm_v2_go.TestCondType, tce e2sm_kpm_v2_go.TestCondExpression, tcv *e2sm_kpm_v2_go.TestCondValue) (*e2sm_kpm_v2_go.TestCondInfo, error) {

	tci := e2sm_kpm_v2_go.TestCondInfo{
		TestValue: tcv,
		TestExpr:  tce,
		TestType:  tct,
	}

	//if err := tci.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating TestCondInfo (TestCondInfo) %s", err.Error())
	//}
	return &tci, nil
}

func CreateTestCondTypeGBR() *e2sm_kpm_v2_go.TestCondType {

	return &e2sm_kpm_v2_go.TestCondType{
		TestCondType: &e2sm_kpm_v2_go.TestCondType_GBr{
			GBr: e2sm_kpm_v2_go.GBR_GBR_TRUE,
		},
	}
}

func CreateTestCondTypeAMBR() *e2sm_kpm_v2_go.TestCondType {

	return &e2sm_kpm_v2_go.TestCondType{
		TestCondType: &e2sm_kpm_v2_go.TestCondType_AMbr{
			AMbr: e2sm_kpm_v2_go.AMBR_AMBR_TRUE,
		},
	}
}

func CreateTestCondTypeIsStat() *e2sm_kpm_v2_go.TestCondType {

	return &e2sm_kpm_v2_go.TestCondType{
		TestCondType: &e2sm_kpm_v2_go.TestCondType_IsStat{
			IsStat: e2sm_kpm_v2_go.ISSTAT_ISSTAT_TRUE,
		},
	}
}

func CreateTestCondTypeIsCatM() *e2sm_kpm_v2_go.TestCondType {

	return &e2sm_kpm_v2_go.TestCondType{
		TestCondType: &e2sm_kpm_v2_go.TestCondType_IsCatM{
			IsCatM: e2sm_kpm_v2_go.ISCATM_ISCATM_TRUE,
		},
	}
}

func CreateTestCondTypeRSRP() *e2sm_kpm_v2_go.TestCondType {

	return &e2sm_kpm_v2_go.TestCondType{
		TestCondType: &e2sm_kpm_v2_go.TestCondType_RSrp{
			RSrp: e2sm_kpm_v2_go.RSRP_RSRP_TRUE,
		},
	}
}

func CreateTestCondTypeRSRQ() *e2sm_kpm_v2_go.TestCondType {

	return &e2sm_kpm_v2_go.TestCondType{
		TestCondType: &e2sm_kpm_v2_go.TestCondType_RSrq{
			RSrq: e2sm_kpm_v2_go.RSRQ_RSRQ_TRUE,
		},
	}
}

func CreateTestCondValueInt(val int64) *e2sm_kpm_v2_go.TestCondValue {

	return &e2sm_kpm_v2_go.TestCondValue{
		TestCondValue: &e2sm_kpm_v2_go.TestCondValue_ValueInt{
			ValueInt: val,
		},
	}
}

func CreateTestCondValueEnum(val int64) *e2sm_kpm_v2_go.TestCondValue {

	return &e2sm_kpm_v2_go.TestCondValue{
		TestCondValue: &e2sm_kpm_v2_go.TestCondValue_ValueEnum{
			ValueEnum: val,
		},
	}
}

func CreateTestCondValueBool(val bool) *e2sm_kpm_v2_go.TestCondValue {

	return &e2sm_kpm_v2_go.TestCondValue{
		TestCondValue: &e2sm_kpm_v2_go.TestCondValue_ValueBool{
			ValueBool: val,
		},
	}
}

func CreateTestCondValueBitS(val *asn1.BitString) *e2sm_kpm_v2_go.TestCondValue {

	return &e2sm_kpm_v2_go.TestCondValue{
		TestCondValue: &e2sm_kpm_v2_go.TestCondValue_ValueBitS{
			ValueBitS: val,
		},
	}
}

func CreateTestCondValueOctS(val []byte) *e2sm_kpm_v2_go.TestCondValue {

	return &e2sm_kpm_v2_go.TestCondValue{
		TestCondValue: &e2sm_kpm_v2_go.TestCondValue_ValueOctS{
			ValueOctS: val,
		},
	}
}

func CreateTestCondValuePrtS(val string) *e2sm_kpm_v2_go.TestCondValue {

	return &e2sm_kpm_v2_go.TestCondValue{
		TestCondValue: &e2sm_kpm_v2_go.TestCondValue_ValuePrtS{
			ValuePrtS: val,
		},
	}
}
