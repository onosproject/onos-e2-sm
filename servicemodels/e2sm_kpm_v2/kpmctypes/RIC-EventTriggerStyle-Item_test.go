// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createRicEventTriggerStyleItem() *e2sm_kpm_v2.RicEventTriggerStyleItem {

	return &e2sm_kpm_v2.RicEventTriggerStyleItem{
		RicEventTriggerStyleType: &e2sm_kpm_v2.RicStyleType{
			Value: 21,
		},
		RicEventTriggerStyleName: &e2sm_kpm_v2.RicStyleName{
			Value: "onf",
		},
		RicEventTriggerFormatType: &e2sm_kpm_v2.RicFormatType{
			Value: 22,
		},
	}
}

func Test_xerEncodeRicEventTriggerStyleItem(t *testing.T) {

	item := createRicEventTriggerStyleItem()

	xer, err := xerEncodeRicEventTriggerStyleItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("RicEventTriggerStyleItem XER\n%s", string(xer))
}

func Test_xerDecodeRicEventTriggerStyleItem(t *testing.T) {

	item := createRicEventTriggerStyleItem()

	xer, err := xerEncodeRicEventTriggerStyleItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(xer))
	t.Logf("RicEventTriggerStyleItem XER\n%s", string(xer))

	result, err := perDecodeRicEventTriggerStyleItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}

func Test_perEncodeRicEventTriggerStyleItem(t *testing.T) {

	item := createRicEventTriggerStyleItem()

	per, err := perEncodeRicEventTriggerStyleItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("RicEventTriggerStyleItem PER\n%s", string(per))
}

func Test_perDecodeRicEventTriggerStyleItem(t *testing.T) {

	item := createRicEventTriggerStyleItem()

	per, err := perEncodeRicEventTriggerStyleItem(item)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("RicEventTriggerStyleItem PER\n%s", string(per))

	result, err := perDecodeRicEventTriggerStyleItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
