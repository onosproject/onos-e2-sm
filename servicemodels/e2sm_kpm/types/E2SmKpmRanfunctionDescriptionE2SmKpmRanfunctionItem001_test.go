// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001(t *testing.T) {

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001()

	assert.Equal(t, reflect.TypeOf(E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001{}), reflect.TypeOf(*e2SmKpmRanfunctionItem001), "E2SmKpmRanfunctionItem001{} types are mismatched")
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_AddRicEventTriggerStyleListItem(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var ricFormatType int32 = 15
	ricEventTriggerStyleList := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(ricStyleType)).
		SetRicEventTriggerStyleName(NewRicStyleName(ricStyleName)).SetRicEventTriggerFormatType(NewRicFormatType(ricFormatType))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().AddRicEventTriggerStyleListItem(ricEventTriggerStyleList)

	assert.Assert(t, e2SmKpmRanfunctionItem001.RicEventTriggerStyleList != nil)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_RetrieveRicEventTriggerStyleListItemByRicStyleType(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var ricFormatType int32 = 15
	ricEventTriggerStyleListItem1 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(ricStyleType)).
		SetRicEventTriggerStyleName(NewRicStyleName(ricStyleName)).SetRicEventTriggerFormatType(NewRicFormatType(ricFormatType))
	ricEventTriggerStyleListItem2 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(21)).
		SetRicEventTriggerStyleName(NewRicStyleName("FNO")).SetRicEventTriggerFormatType(NewRicFormatType(23))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem1).AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem2)

	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicEventTriggerStyleListItemByRicStyleType(ricStyleType), ricEventTriggerStyleListItem1)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_RetrieveRicEventTriggerStyleListItemByRicStyleName(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var ricFormatType int32 = 15
	ricEventTriggerStyleListItem1 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(ricStyleType)).
		SetRicEventTriggerStyleName(NewRicStyleName(ricStyleName)).SetRicEventTriggerFormatType(NewRicFormatType(ricFormatType))
	ricEventTriggerStyleListItem2 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(21)).
		SetRicEventTriggerStyleName(NewRicStyleName("FNO")).SetRicEventTriggerFormatType(NewRicFormatType(23))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem1).AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem2)

	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicEventTriggerStyleListItemByRicStyleName("FNO"), ricEventTriggerStyleListItem2)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_RetrieveRicEventTriggerStyleListItemByRicFormatType(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var ricFormatType int32 = 15
	ricEventTriggerStyleListItem1 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(ricStyleType)).
		SetRicEventTriggerStyleName(NewRicStyleName(ricStyleName)).SetRicEventTriggerFormatType(NewRicFormatType(ricFormatType))
	ricEventTriggerStyleListItem2 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(21)).
		SetRicEventTriggerStyleName(NewRicStyleName("FNO")).SetRicEventTriggerFormatType(NewRicFormatType(23))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem1).AddRicEventTriggerStyleListItem(ricEventTriggerStyleListItem2)

	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicEventTriggerStyleListItemByRicFormatType(ricFormatType), ricEventTriggerStyleListItem1)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_AddRicReportStyleListItem(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var indHdrFormatType int32 = 15
	var indMsgFormatType int32 = 3
	ricReportStyleList := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().AddRicReportStyleListItem(ricReportStyleList)

	assert.Assert(t, e2SmKpmRanfunctionItem001.RicReportStyleList != nil)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_RetrieveRicReportStyleListItemByRicStyleType(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var indHdrFormatType int32 = 15
	var indMsgFormatType int32 = 3
	ricReportStyleListItem1 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType))
	ricReportStyleListItem2 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(21)).
		SetRicReportStyleName(NewRicStyleName("FNO")).SetRicIndicationHeaderFormatType(NewRicFormatType(7)).
		SetRicIndicationMessageFormatType(NewRicFormatType(1))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicReportStyleListItem(ricReportStyleListItem1).AddRicReportStyleListItem(ricReportStyleListItem2)

	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicReportStyleListItemByRicStyleType(ricStyleType), ricReportStyleListItem1)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_RetrieveRicReportStyleListItemByRicStyleName(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var indHdrFormatType int32 = 15
	var indMsgFormatType int32 = 3
	ricReportStyleListItem1 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType))
	ricReportStyleListItem2 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(21)).
		SetRicReportStyleName(NewRicStyleName("FNO")).SetRicIndicationHeaderFormatType(NewRicFormatType(7)).
		SetRicIndicationMessageFormatType(NewRicFormatType(1))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicReportStyleListItem(ricReportStyleListItem1).AddRicReportStyleListItem(ricReportStyleListItem2)

	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicReportStyleListItemByRicStyleName(ricStyleName), ricReportStyleListItem1)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_RetrieveRicReportStyleListItemByIndicationHeader(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var indHdrFormatType int32 = 15
	var indMsgFormatType int32 = 3
	ricReportStyleListItem1 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType))
	ricReportStyleListItem2 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(21)).
		SetRicReportStyleName(NewRicStyleName("FNO")).SetRicIndicationHeaderFormatType(NewRicFormatType(7)).
		SetRicIndicationMessageFormatType(NewRicFormatType(1))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicReportStyleListItem(ricReportStyleListItem1).AddRicReportStyleListItem(ricReportStyleListItem2)

	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicReportStyleListItemByIndicationHeader(7), ricReportStyleListItem2)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_RetrieveRicReportStyleListItemByIndicationMessage(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var indHdrFormatType int32 = 15
	var indMsgFormatType int32 = 3
	ricReportStyleListItem1 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType))
	ricReportStyleListItem2 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(21)).
		SetRicReportStyleName(NewRicStyleName("FNO")).SetRicIndicationHeaderFormatType(NewRicFormatType(7)).
		SetRicIndicationMessageFormatType(NewRicFormatType(1))

	e2SmKpmRanfunctionItem001 := NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001().
		AddRicReportStyleListItem(ricReportStyleListItem1).AddRicReportStyleListItem(ricReportStyleListItem2)

	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicReportStyleListItemByIndicationMessage(indMsgFormatType), ricReportStyleListItem1)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_GetRicEventTriggerStyleList(t *testing.T) {

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

	ricEventTriggerStyleList := e2SmKpmRanfunctionItem001.GetRicEventTriggerStyleList()

	assert.Assert(t, ricEventTriggerStyleList != nil)
	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicEventTriggerStyleListItemByRicFormatType(ricFormatType), ricEventTriggerStyleListItem1)
	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicEventTriggerStyleListItemByRicFormatType(23), ricEventTriggerStyleListItem2)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_GetRicReportStyleList(t *testing.T) {

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

	ricReportStyleList := e2SmKpmRanfunctionItem001.GetRicReportStyleList()

	assert.Assert(t, ricReportStyleList != nil)
	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicReportStyleListItemByIndicationMessage(indMsgFormatType2), ricReportStyleListItem1)
	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicReportStyleListItemByIndicationMessage(1), ricReportStyleListItem2)
}

func TestE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001_GetE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001(t *testing.T) {

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

	ricReportStyleList := e2SmKpmRanfunctionItem001.GetRicReportStyleList()
	ricEventTriggerStyleList := e2SmKpmRanfunctionItem001.GetRicEventTriggerStyleList()

	assert.Assert(t, ricReportStyleList != nil)
	assert.Assert(t, ricEventTriggerStyleList != nil)
	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicEventTriggerStyleListItemByRicFormatType(ricFormatType), ricEventTriggerStyleListItem1)
	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicEventTriggerStyleListItemByRicFormatType(23), ricEventTriggerStyleListItem2)
	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicReportStyleListItemByIndicationMessage(indMsgFormatType2), ricReportStyleListItem1)
	assert.DeepEqual(t, e2SmKpmRanfunctionItem001.RetrieveRicReportStyleListItemByIndicationMessage(1), ricReportStyleListItem2)
}
