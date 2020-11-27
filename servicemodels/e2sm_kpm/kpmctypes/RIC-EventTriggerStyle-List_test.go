// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createSampleVariable() *e2sm_kpm_ies.RicEventTriggerStyleList {

	var ricEventStyleType int32 = 13
	var ricEventStyleName = "ONFevent"
	var ricEventFormatType int32 = 42

	ricEventTriggerStyleList := &e2sm_kpm_ies.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2sm_kpm_ies.RicStyleType{
			Value: ricEventStyleType, //int32
		},
		RicEventTriggerStyleName: &e2sm_kpm_ies.RicStyleName{
			Value: ricEventStyleName, //string
		},
		RicEventTriggerFormatType: &e2sm_kpm_ies.RicFormatType{
			Value: ricEventFormatType, //int32
		},
	}

	return ricEventTriggerStyleList
}

func Test_xerEncodeRicEventTriggerStyleItem(t *testing.T) {

	ricEventTriggerStyleList := createSampleVariable()

	xer, err := xerEncodeRicEventTriggerStyleItem(ricEventTriggerStyleList)
	assert.NilError(t, err)
	assert.Equal(t, 259, len(xer))
	t.Logf("RIC-EventTriggerStyle-List XER\n%s", string(xer))
}

func Test_xerDecodeRicEventTriggerStyleItem(t *testing.T) {

	ricEventTriggerStyleList := createSampleVariable()

	xer, err := xerEncodeRicEventTriggerStyleItem(ricEventTriggerStyleList)
	assert.NilError(t, err)
	assert.Equal(t, 259, len(xer))
	t.Logf("RIC-EventTriggerStyle-List XER\n%s", string(xer))

	result, err := xerDecodeRicEventTriggerStyleItem(xer)
	assert.NilError(t, err)
	assert.Equal(t, ricEventTriggerStyleList.RicEventTriggerStyleType.Value, result.RicEventTriggerStyleType.Value, "Encoded and decoded RicStyleType values are not the same")
	assert.Equal(t, ricEventTriggerStyleList.RicEventTriggerStyleName.Value, result.RicEventTriggerStyleName.Value, "Encoded and decoded RicStyleName values are not the same")
	assert.Equal(t, ricEventTriggerStyleList.RicEventTriggerFormatType.Value, result.RicEventTriggerFormatType.Value, "Encoded and decoded RicFormatType values are not the same")
}

func Test_perEncodeRicEventTriggerStyleItem(t *testing.T) {

	ricEventTriggerStyleList := createSampleVariable()

	per, err := perEncodeRicEventTriggerStyleItem(ricEventTriggerStyleList)
	assert.NilError(t, err)
	assert.Equal(t, 15, len(per))
	t.Logf("RIC-EventTriggerStyle-List PER\n%s", string(per))
}

func Test_perDecodeRicEventTriggerStyleItem(t *testing.T) {

	ricEventTriggerStyleList := createSampleVariable()

	per, err := perEncodeRicEventTriggerStyleItem(ricEventTriggerStyleList)
	assert.NilError(t, err)
	assert.Equal(t, 15, len(per))
	t.Logf("RIC-EventTriggerStyle-List PER\n%s", string(per))

	result, err := perDecodeRicEventTriggerStyleItem(per)
	assert.NilError(t, err)
	assert.Equal(t, ricEventTriggerStyleList.RicEventTriggerStyleType.Value, result.RicEventTriggerStyleType.Value, "Encoded and decoded RicStyleType values are not the same")
	assert.Equal(t, ricEventTriggerStyleList.RicEventTriggerStyleName.Value, result.RicEventTriggerStyleName.Value, "Encoded and decoded RicStyleName values are not the same")
	assert.Equal(t, ricEventTriggerStyleList.RicEventTriggerFormatType.Value, result.RicEventTriggerFormatType.Value, "Encoded and decoded RicFormatType values are not the same")
}
