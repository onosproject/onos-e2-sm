// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerRicEventTriggerStyleItem = "00000000  00 15 01 00 6f 6e 66 00  16                       |....onf..|"

func createRicEventTriggerStyleItem() *e2sm_kpm_v2_go.RicEventTriggerStyleItem {

	return &e2sm_kpm_v2_go.RicEventTriggerStyleItem{
		RicEventTriggerStyleType: &e2sm_kpm_v2_go.RicStyleType{
			Value: 21,
		},
		RicEventTriggerStyleName: &e2sm_kpm_v2_go.RicStyleName{
			Value: "onf",
		},
		RicEventTriggerFormatType: &e2sm_kpm_v2_go.RicFormatType{
			Value: 22,
		},
	}
}

func Test_perEncodingRicEventTriggerStyleItem(t *testing.T) {

	item := createRicEventTriggerStyleItem()

	per, err := aper.MarshalWithParams(item, "valueExt")
	assert.NilError(t, err)
	t.Logf("RIC-EventTriggerStyle-Item PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RicEventTriggerStyleItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("RIC-EventTriggerStyle-Item - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicEventTriggerStyleItem)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
