// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerGlobalEnbID1 = "00000000  00 21 22 23 00 d4 bc 00                           |.!\"#....|"
var refPerGlobalEnbID2 = "00000000  00 21 22 23 40 d4 bc 09  00                       |.!\"#@....|"

func createGlobalEnbID1() *e2sm_kpm_v2_go.GlobalEnbId {

	return &e2sm_kpm_v2_go.GlobalEnbId{
		PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		ENbId: &e2sm_kpm_v2_go.EnbId{
			EnbId: &e2sm_kpm_v2_go.EnbId_MacroENbId{
				MacroENbId: &asn1.BitString{
					Value: []byte{0xd4, 0xbc, 0x00},
					Len:   20,
				},
			},
		},
	}
}

func createGlobalEnbID2() *e2sm_kpm_v2_go.GlobalEnbId {

	return &e2sm_kpm_v2_go.GlobalEnbId{
		PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		ENbId: &e2sm_kpm_v2_go.EnbId{
			EnbId: &e2sm_kpm_v2_go.EnbId_HomeENbId{
				HomeENbId: &asn1.BitString{
					Value: []byte{0xd4, 0xbc, 0x09, 0x00},
					Len:   28,
				},
			},
		},
	}
}

func Test_perEncodingGlobalEnbID1(t *testing.T) {

	globalEnbID1 := createGlobalEnbID1()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per1, err := aper.MarshalWithParams(*globalEnbID1, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalEnbID (Macro) PER\n%v", hex.Dump(per1))

	result1 := e2sm_kpm_v2_go.GlobalEnbId{}
	err = aper.UnmarshalWithParams(per1, &result1, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result1 != nil)
	t.Logf("GlobalEnbID (Macro) PER - decoded\n%v", result1)
}

func Test_perGlobalEnbID1CompareBytes(t *testing.T) {

	globalEnbID1 := createGlobalEnbID1()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per1, err := aper.MarshalWithParams(*globalEnbID1, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalEnbID (Macro) PER\n%v", hex.Dump(per1))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalEnbID1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per1, perRefBytes)
}

func Test_perEncodingGlobalEnbID2(t *testing.T) {
	globalEnbID2 := createGlobalEnbID2()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per2, err := aper.MarshalWithParams(*globalEnbID2, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalEnbID (Home) PER\n%v", hex.Dump(per2))

	result2 := e2sm_kpm_v2_go.GlobalEnbId{}
	err = aper.UnmarshalWithParams(per2, &result2, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result2 != nil)
	t.Logf("GlobalEnbID (Home) PER - decoded\n%v", result2)
}

func Test_perGlobalEnbID2CompareBytes(t *testing.T) {
	globalEnbID2 := createGlobalEnbID2()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per2, err := aper.MarshalWithParams(*globalEnbID2, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalEnbID (Home) PER\n%v", hex.Dump(per2))

	//Comparing with reference bytes
	perRefBytes2, err := hexlib.DumpToByte(refPerGlobalEnbID2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per2, perRefBytes2)
}
