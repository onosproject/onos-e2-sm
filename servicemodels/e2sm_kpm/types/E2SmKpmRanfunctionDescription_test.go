// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestE2SmKpmRanfunctionDescription_NewE2SmKpmRanfunctionDescription(t *testing.T) {

	e2SmKpmRanfunctionDescription := NewE2SmKpmRanfunctionDescription()

	assert.Equal(t, reflect.TypeOf(E2SmKpmRanfunctionDescription{}), reflect.TypeOf(*e2SmKpmRanfunctionDescription), "E2SmKpmRanfunctionDescription{} types are mismatched")
}

func TestE2SmKpmRanfunctionDescription_SetRanfunctionName(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName := NewRanfunctionName(shortName, oid, description, instance)

	e2SmKpmRanfunctionDescription := NewE2SmKpmRanfunctionDescription().SetRanfunctionName(ranfunctionName)

	assert.DeepEqual(t, e2SmKpmRanfunctionDescription.RanFunctionName, ranfunctionName)
}

func TestE2SmKpmRanfunctionDescription_SetE2SmKpmRanfunctionItem(t *testing.T) {

	var ricStyleType1 int32 = 14
	var ricStyleName1 string = "ONF"
	var ricFormatType int32 = 15
	ricEventTriggerStyleListItem1 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(ricStyleType1)).
		SetRicEventTriggerStyleName(NewRicStyleName(ricStyleName1)).SetRicEventTriggerFormatType(NewRicFormatType(ricFormatType))
	ricEventTriggerStyleListItem2 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(21)).
		SetRicEventTriggerStyleName(NewRicStyleName("FNO")).SetRicEventTriggerFormatType(NewRicFormatType(23))

	var ricStyleType2 int32 = 14
	var ricStyleName2 string = "ONF"
	var indHdrFormatType2 int32 = 15
	var indMsgFormatType2 int32 = 3
	ricReportStyleListItem1 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType2)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName2)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType2)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType2))
	ricReportStyleListItem2 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(21)).
		SetRicReportStyleName(NewRicStyleName("FNO")).SetRicIndicationHeaderFormatType(NewRicFormatType(7)).
		SetRicIndicationMessageFormatType(NewRicFormatType(1))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem1).AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem2).
		AddRicReportStyleListItem(ricReportStyleListItem1).AddRicReportStyleListItem(ricReportStyleListItem2)

	e2SmKpmRanfunctionDescription := NewE2SmKpmRanfunctionDescription().SetE2SmKpmRanfunctionItem(e2SmKpmRanfunctionItem001)

	assert.DeepEqual(t, e2SmKpmRanfunctionDescription.E2SmKpmRanfunctionItem, e2SmKpmRanfunctionItem001)
}

func TestE2SmKpmRanfunctionDescription_GetRanfunctionName(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName := NewRanfunctionName(shortName, oid, description, instance)

	var ricStyleType1 int32 = 14
	var ricStyleName1 string = "ONF"
	var ricFormatType int32 = 15
	ricEventTriggerStyleListItem1 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(ricStyleType1)).
		SetRicEventTriggerStyleName(NewRicStyleName(ricStyleName1)).SetRicEventTriggerFormatType(NewRicFormatType(ricFormatType))
	ricEventTriggerStyleListItem2 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(21)).
		SetRicEventTriggerStyleName(NewRicStyleName("FNO")).SetRicEventTriggerFormatType(NewRicFormatType(23))

	var ricStyleType2 int32 = 14
	var ricStyleName2 string = "ONF"
	var indHdrFormatType2 int32 = 15
	var indMsgFormatType2 int32 = 3
	ricReportStyleListItem1 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType2)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName2)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType2)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType2))
	ricReportStyleListItem2 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(21)).
		SetRicReportStyleName(NewRicStyleName("FNO")).SetRicIndicationHeaderFormatType(NewRicFormatType(7)).
		SetRicIndicationMessageFormatType(NewRicFormatType(1))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem1).AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem2).
		AddRicReportStyleListItem(ricReportStyleListItem1).AddRicReportStyleListItem(ricReportStyleListItem2)

	e2SmKpmRanfunctionDescription := NewE2SmKpmRanfunctionDescription().SetRanfunctionName(ranfunctionName).SetE2SmKpmRanfunctionItem(e2SmKpmRanfunctionItem001)

	assert.DeepEqual(t, e2SmKpmRanfunctionDescription.GetRanfunctionName(), ranfunctionName)
}

func TestE2SmKpmRanfunctionDescription_GetE2SmKpmRanfunctionItem(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName := NewRanfunctionName(shortName, oid, description, instance)

	var ricStyleType1 int32 = 14
	var ricStyleName1 string = "ONF"
	var ricFormatType int32 = 15
	ricEventTriggerStyleListItem1 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(ricStyleType1)).
		SetRicEventTriggerStyleName(NewRicStyleName(ricStyleName1)).SetRicEventTriggerFormatType(NewRicFormatType(ricFormatType))
	ricEventTriggerStyleListItem2 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(21)).
		SetRicEventTriggerStyleName(NewRicStyleName("FNO")).SetRicEventTriggerFormatType(NewRicFormatType(23))

	var ricStyleType2 int32 = 14
	var ricStyleName2 string = "ONF"
	var indHdrFormatType2 int32 = 15
	var indMsgFormatType2 int32 = 3
	ricReportStyleListItem1 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType2)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName2)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType2)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType2))
	ricReportStyleListItem2 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(21)).
		SetRicReportStyleName(NewRicStyleName("FNO")).SetRicIndicationHeaderFormatType(NewRicFormatType(7)).
		SetRicIndicationMessageFormatType(NewRicFormatType(1))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem1).AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem2).
		AddRicReportStyleListItem(ricReportStyleListItem1).AddRicReportStyleListItem(ricReportStyleListItem2)

	e2SmKpmRanfunctionDescription := NewE2SmKpmRanfunctionDescription().SetRanfunctionName(ranfunctionName).SetE2SmKpmRanfunctionItem(e2SmKpmRanfunctionItem001)

	assert.DeepEqual(t, e2SmKpmRanfunctionDescription.GetE2SmKpmRanfunctionItem(), e2SmKpmRanfunctionItem001)
}

func TestE2SmKpmRanfunctionDescription_GetE2SmKpmRanfunctionDescription(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName := NewRanfunctionName(shortName, oid, description, instance)

	var ricStyleType1 int32 = 14
	var ricStyleName1 string = "ONF"
	var ricFormatType int32 = 15
	ricEventTriggerStyleListItem1 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(ricStyleType1)).
		SetRicEventTriggerStyleName(NewRicStyleName(ricStyleName1)).SetRicEventTriggerFormatType(NewRicFormatType(ricFormatType))
	ricEventTriggerStyleListItem2 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(21)).
		SetRicEventTriggerStyleName(NewRicStyleName("FNO")).SetRicEventTriggerFormatType(NewRicFormatType(23))

	var ricStyleType2 int32 = 14
	var ricStyleName2 string = "ONF"
	var indHdrFormatType2 int32 = 15
	var indMsgFormatType2 int32 = 3
	ricReportStyleListItem1 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType2)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName2)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType2)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType2))
	ricReportStyleListItem2 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(21)).
		SetRicReportStyleName(NewRicStyleName("FNO")).SetRicIndicationHeaderFormatType(NewRicFormatType(7)).
		SetRicIndicationMessageFormatType(NewRicFormatType(1))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem1).AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem2).
		AddRicReportStyleListItem(ricReportStyleListItem1).AddRicReportStyleListItem(ricReportStyleListItem2)

	e2SmKpmRanfunctionDescription1 := NewE2SmKpmRanfunctionDescription().SetRanfunctionName(ranfunctionName).SetE2SmKpmRanfunctionItem(e2SmKpmRanfunctionItem001)
	e2SmKpmRanfunctionDescription2 := e2SmKpmRanfunctionDescription1.GetE2SmKpmRanfunctionDescription()

	assert.DeepEqual(t, e2SmKpmRanfunctionDescription1, e2SmKpmRanfunctionDescription2)
}
