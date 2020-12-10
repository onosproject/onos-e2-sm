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

func TestRicEventTriggerStyleList_NewRicEventTriggerStyleList(t *testing.T) {

	ricEventTriggerStyleList := NewRicEventTriggerStyleList()

	assert.Equal(t, reflect.TypeOf(RicEventTriggerStyleList{}), reflect.TypeOf(*ricEventTriggerStyleList), "RicEventTriggerStyleList{} types are mismatched")
	assert.DeepEqual(t, *ricEventTriggerStyleList, RicEventTriggerStyleList{})
}

func TestRicEventTriggerStyleList_SetRicEventTriggerStyleType(t *testing.T) {

	var value int32 = 14
	ricEventTriggerStyleList := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(value))

	assert.Equal(t, ricEventTriggerStyleList.RicEventTriggerStyleType.GetValue(), value, "Test_RicEventTriggerStyleList SetRicStyleType values mismatch")
}

func TestRicEventTriggerStyleList_SetRicEventTriggerStyleName(t *testing.T) {

	var value string = "ONF"
	ricEventTriggerStyleList := NewRicEventTriggerStyleList().SetRicEventTriggerStyleName(NewRicStyleName(value))

	assert.Equal(t, ricEventTriggerStyleList.RicEventTriggerStyleName.GetValue(), value, "Test_RicEventTriggerStyleList SetRicStyleName values mismatch")
}

func TestRicEventTriggerStyleList_SetRicEventTriggerFormatType(t *testing.T) {

	var value int32 = 15
	ricEventTriggerStyleList := NewRicEventTriggerStyleList().SetRicEventTriggerFormatType(NewRicFormatType(value))

	assert.Equal(t, ricEventTriggerStyleList.RicEventTriggerFormatType.GetValue(), value, "Test_RicEventTriggerStyleList SetRicFormatType values mismatch")
}

func TestRicEventTriggerStyleList_GetRicEventTriggerStyleType(t *testing.T) {

	var value int32 = 14
	ricEventTriggerStyleList := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(value))

	assert.Equal(t, ricEventTriggerStyleList.GetRicEventTriggerStyleType().GetValue(), value, "Test_RicEventTriggerStyleList GetRicStyleType values mismatch")
}

func TestRicEventTriggerStyleList_GetRicEventTriggerStyleName(t *testing.T) {

	var value string = "ONF"
	ricEventTriggerStyleList := NewRicEventTriggerStyleList().SetRicEventTriggerStyleName(NewRicStyleName(value))

	assert.Equal(t, ricEventTriggerStyleList.GetRicEventTriggerStyleName().GetValue(), value, "Test_RicEventTriggerStyleList GetRicStyleName values mismatch")
}

func TestRicEventTriggerStyleList_GetRicEventTriggerFormatType(t *testing.T) {

	var value int32 = 15
	ricEventTriggerStyleList := NewRicEventTriggerStyleList().SetRicEventTriggerFormatType(NewRicFormatType(value))

	assert.Equal(t, ricEventTriggerStyleList.GetRicEventTriggerFormatType().GetValue(), value, "Test_RicEventTriggerStyleList GetRicFormatType values mismatch")
}

func TestRicEventTriggerStyleList_GetRicEventTriggerStyleList(t *testing.T) {

	var ricStyleType int32 = 14
	var ricStyleName string = "ONF"
	var ricFormatType int32 = 15
	ricEventTriggerStyleList1 := NewRicEventTriggerStyleList().SetRicEventTriggerStyleType(NewRicStyleType(ricStyleType)).
		SetRicEventTriggerStyleName(NewRicStyleName(ricStyleName)).SetRicEventTriggerFormatType(NewRicFormatType(ricFormatType))
	ricEventTriggerStyleList2 := ricEventTriggerStyleList1.GetRicEventTriggerStyleList()

	assert.DeepEqual(t, ricEventTriggerStyleList1, ricEventTriggerStyleList2)
}
