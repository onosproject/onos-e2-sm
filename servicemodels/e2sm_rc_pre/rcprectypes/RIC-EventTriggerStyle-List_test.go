// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func createRicEventTriggerStyleItem() *e2sm_rc_pre_ies.RicEventTriggerStyleList {

	var ricEventStyleType int32 = 13
	var ricEventStyleName = "ONFevent"
	var ricEventFormatType int32 = 42

	ricEventTriggerStyleItem := &e2sm_rc_pre_ies.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2sm_rc_pre_ies.RicStyleType{
			Value: ricEventStyleType, //int32
		},
		RicEventTriggerStyleName: &e2sm_rc_pre_ies.RicStyleName{
			Value: ricEventStyleName, //string
		},
		RicEventTriggerFormatType: &e2sm_rc_pre_ies.RicFormatType{
			Value: ricEventFormatType, //int32
		},
	}

	return ricEventTriggerStyleItem
}

func Test_xerEncodeRicEventTriggerStyleItem(t *testing.T) {

	ricEventTriggerStyleItem := createRicEventTriggerStyleItem()

	xer, err := xerEncodeRicEventTriggerStyleItem(ricEventTriggerStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 259, len(xer))
	t.Logf("RIC-EventTriggerStyle-List XER\n%s", string(xer))
}

func Test_xerDecodeRicEventTriggerStyleItem(t *testing.T) {

	ricEventTriggerStyleItem := createRicEventTriggerStyleItem()

	xer, err := xerEncodeRicEventTriggerStyleItem(ricEventTriggerStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 259, len(xer))
	t.Logf("RIC-EventTriggerStyle-List XER\n%s", string(xer))

	result, err := xerDecodeRicEventTriggerStyleItem(xer)
	assert.NilError(t, err)
	assert.Equal(t, ricEventTriggerStyleItem.RicEventTriggerStyleType.Value, result.RicEventTriggerStyleType.Value, "Encoded and decoded RicStyleType values are not the same")
	assert.Equal(t, ricEventTriggerStyleItem.RicEventTriggerStyleName.Value, result.RicEventTriggerStyleName.Value, "Encoded and decoded RicStyleName values are not the same")
	assert.Equal(t, ricEventTriggerStyleItem.RicEventTriggerFormatType.Value, result.RicEventTriggerFormatType.Value, "Encoded and decoded RicFormatType values are not the same")
}

func Test_perEncodeRicEventTriggerStyleItem(t *testing.T) {

	ricEventTriggerStyleItem := createRicEventTriggerStyleItem()

	per, err := perEncodeRicEventTriggerStyleItem(ricEventTriggerStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 15, len(per))
	t.Logf("RIC-EventTriggerStyle-List PER\n%s", string(per))
}

func Test_perDecodeRicEventTriggerStyleItem(t *testing.T) {

	ricEventTriggerStyleItem := createRicEventTriggerStyleItem()

	per, err := perEncodeRicEventTriggerStyleItem(ricEventTriggerStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 15, len(per))
	t.Logf("RIC-EventTriggerStyle-List PER\n%s", string(per))

	result, err := perDecodeRicEventTriggerStyleItem(per)
	assert.NilError(t, err)
	assert.Equal(t, ricEventTriggerStyleItem.RicEventTriggerStyleType.Value, result.RicEventTriggerStyleType.Value, "Encoded and decoded RicStyleType values are not the same")
	assert.Equal(t, ricEventTriggerStyleItem.RicEventTriggerStyleName.Value, result.RicEventTriggerStyleName.Value, "Encoded and decoded RicStyleName values are not the same")
	assert.Equal(t, ricEventTriggerStyleItem.RicEventTriggerFormatType.Value, result.RicEventTriggerFormatType.Value, "Encoded and decoded RicFormatType values are not the same")
}
