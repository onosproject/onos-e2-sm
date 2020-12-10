// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	//"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	"gotest.tools/assert"
	//"io/ioutil"
	"reflect"
	"testing"
)

func TestRicReportStyleList_RicReportStyleList(t *testing.T) {

	ricReportStyleList := NewRicReportStyleList()

	assert.Equal(t, reflect.TypeOf(RicReportStyleList{}), reflect.TypeOf(*ricReportStyleList), "RicReportStyleList{} types are mismatched")
	assert.DeepEqual(t, *ricReportStyleList, RicReportStyleList{})
}

func TestRicReportStyleList_SetRicReportStyleType(t *testing.T) {

	var value int32 = 14
	ricReportStyleList := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(value))

	assert.Equal(t, ricReportStyleList.RicReportStyleType.GetValue(), value, "Test_RicReportStyleList SetRicReportStyleType values mismatch")
}

func TestRicReportStyleList_SetRicReportStyleName(t *testing.T) {

	var value string = "ONF"
	ricReportStyleList := NewRicReportStyleList().SetRicReportStyleName(NewRicStyleName(value))

	assert.Equal(t, ricReportStyleList.RicReportStyleName.GetValue(), value, "Test_RicReportStyleList SetRicStyleName values mismatch")
}

func TestRicReportStyleList_SetRicReportIndicationHeaderFormatType(t *testing.T) {

	var value int32 = 15
	ricReportStyleList := NewRicReportStyleList().SetRicIndicationHeaderFormatType(NewRicFormatType(value))

	assert.Equal(t, ricReportStyleList.RicIndicationHeaderFormatType.GetValue(), value, "Test_RicReportStyleList SetRicIndicationHeaderFormatType values mismatch")
}

func TestRicReportStyleList_SetRicReportIndicationMessageFormatType(t *testing.T) {

	var value int32 = 3
	ricReportStyleList := NewRicReportStyleList().SetRicIndicationMessageFormatType(NewRicFormatType(value))

	assert.Equal(t, ricReportStyleList.RicIndicationMessageFormatType.GetValue(), value, "Test_RicReportStyleList SetRicIndicationMessageFormatType values mismatch")
}

func TestRicReportStyleList_GetRicReportStyleType(t *testing.T) {

	var value int32 = 14
	ricReportStyleList := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(value))

	assert.Equal(t, ricReportStyleList.GetRicReportStyleType().GetValue(), value, "Test_RicReportStyleList GetRicStyleType values mismatch")
}

func TestRicReportStyleList_GetRicReportStyleName(t *testing.T) {

	var value string = "ONF"
	ricReportStyleList := NewRicReportStyleList().SetRicReportStyleName(NewRicStyleName(value))

	assert.Equal(t, ricReportStyleList.GetRicReportStyleName().GetValue(), value, "Test_RicReportStyleList GetRicStyleName values mismatch")
}

func TestRicReportStyleList_GetRicIndicationHeaderFormatType(t *testing.T) {

	var value int32 = 15
	ricReportStyleList := NewRicReportStyleList().SetRicIndicationHeaderFormatType(NewRicFormatType(value))

	assert.Equal(t, ricReportStyleList.GetRicIndicationHeaderFormatType().GetValue(), value, "Test_RicReportStyleList GetRicIndicationHeaderFormatType values mismatch")
}

func TestRicReportStyleList_GetRicIndicationMessageFormatType(t *testing.T) {

	var value int32 = 3
	ricReportStyleList := NewRicReportStyleList().SetRicIndicationMessageFormatType(NewRicFormatType(value))

	assert.Equal(t, ricReportStyleList.GetRicIndicationMessageFormatType().GetValue(), value, "Test_RicReportStyleList GetRicIndicationMessageFormatType values mismatch")
}

func TestRicReportStyleList_GetRicReportStyleList(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var indHdrFormatType int32 = 15
	var indMsgFormatType int32 = 3
	ricReportStyleList1 := NewRicReportStyleList().SetRicReportStyleType(NewRicStyleType(ricStyleType)).
		SetRicReportStyleName(NewRicStyleName(ricStyleName)).SetRicIndicationHeaderFormatType(NewRicFormatType(indHdrFormatType)).
		SetRicIndicationMessageFormatType(NewRicFormatType(indMsgFormatType))
	ricReportStyleList2 := ricReportStyleList1.GetRicReportStyleList()

	assert.DeepEqual(t, ricReportStyleList1, ricReportStyleList2)
}
