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

var refPerGlobalenGnbID = "00000000  00 21 22 23 00 d4 bc 08                           |.!\"#....|"
var refPerGlobalenGnbIDlen31 = "00000000  00 21 22 23 48 d4 bc 08  fe                       |.!\"#H....|"
var refPerGlobalenGnbIDlen32 = "00000000  00 21 22 23 50 d4 bc 08  ff                       |.!\"#P....|"

func createGlobalenGnbID() *e2sm_kpm_v2_go.GlobalenGnbId {

	return &e2sm_kpm_v2_go.GlobalenGnbId{
		PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		GNbId: &e2sm_kpm_v2_go.EngnbId{
			EngnbId: &e2sm_kpm_v2_go.EngnbId_GNbId{
				GNbId: &asn1.BitString{
					Value: []byte{0xd4, 0xbc, 0x08},
					Len:   22,
				},
			},
		},
	}
}

func createGlobalenGnbIDlen31() *e2sm_kpm_v2_go.GlobalenGnbId {

	return &e2sm_kpm_v2_go.GlobalenGnbId{
		PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		GNbId: &e2sm_kpm_v2_go.EngnbId{
			EngnbId: &e2sm_kpm_v2_go.EngnbId_GNbId{
				GNbId: &asn1.BitString{
					Value: []byte{0xd4, 0xbc, 0x08, 0xfe},
					Len:   31,
				},
			},
		},
	}
}

func createGlobalenGnbIDlen32() *e2sm_kpm_v2_go.GlobalenGnbId {

	return &e2sm_kpm_v2_go.GlobalenGnbId{
		PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		GNbId: &e2sm_kpm_v2_go.EngnbId{
			EngnbId: &e2sm_kpm_v2_go.EngnbId_GNbId{
				GNbId: &asn1.BitString{
					Value: []byte{0xd4, 0xbc, 0x08, 0xff},
					Len:   32,
				},
			},
		},
	}
}

func Test_perEncodingGlobalenGnbID(t *testing.T) {

	gnbIDc := createGlobalenGnbID()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalenGnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GlobalenGnbID PER - decoded\n%v", result)
}

func Test_perGlobalenGnbIDCompareBytes(t *testing.T) {

	gnbIDc := createGlobalenGnbID()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalenGnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingGlobalenGnbIDlen31(t *testing.T) {

	gnbIDc := createGlobalenGnbIDlen31()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalenGnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GlobalenGnbID PER - decoded\n%v", result)
}

func Test_perGlobalenGnbIDlen31CompareBytes(t *testing.T) {

	gnbIDc := createGlobalenGnbIDlen31()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalenGnbIDlen31)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingGlobalenGnbIDlen32(t *testing.T) {

	gnbIDc := createGlobalenGnbIDlen32()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalenGnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GlobalenGnbID PER - decoded\n%v", result)
}

func Test_perGlobalenGnbIDlen32CompareBytes(t *testing.T) {

	gnbIDc := createGlobalenGnbIDlen32()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalenGnbIDlen32)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
