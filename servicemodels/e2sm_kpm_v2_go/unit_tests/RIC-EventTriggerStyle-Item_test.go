// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerRicEventTriggerStyleItem = "00000000  00 15 01 00 6f 6e 66 00  16                       |....onf..|"

func createRicEventTriggerStyleItem() *e2smkpmv2.RicEventTriggerStyleItem {

	return &e2smkpmv2.RicEventTriggerStyleItem{
		RicEventTriggerStyleType: &e2smkpmv2.RicStyleType{
			Value: 21,
		},
		RicEventTriggerStyleName: &e2smkpmv2.RicStyleName{
			Value: "onf",
		},
		RicEventTriggerFormatType: &e2smkpmv2.RicFormatType{
			Value: 22,
		},
	}
}

func Test_perEncodingRicEventTriggerStyleItem(t *testing.T) {

	item := createRicEventTriggerStyleItem()

	per, err := aper.MarshalWithParams(item, "valueExt", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-EventTriggerStyle-Item PER\n%v", hex.Dump(per))

	result := e2smkpmv2.RicEventTriggerStyleItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("RIC-EventTriggerStyle-Item - decoded\n%v", &result)
	assert.Equal(t, item.GetRicEventTriggerStyleType().GetValue(), result.GetRicEventTriggerStyleType().GetValue())
	assert.Equal(t, item.GetRicEventTriggerStyleName().GetValue(), result.GetRicEventTriggerStyleName().GetValue())
	assert.Equal(t, item.GetRicEventTriggerFormatType().GetValue(), result.GetRicEventTriggerFormatType().GetValue())
}

func Test_perRicEventTriggerStyleItemCompareBytes(t *testing.T) {

	item := createRicEventTriggerStyleItem()

	per, err := aper.MarshalWithParams(item, "valueExt", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-EventTriggerStyle-Item PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicEventTriggerStyleItem)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
