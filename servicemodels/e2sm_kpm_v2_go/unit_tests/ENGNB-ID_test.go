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

var refPerEnGnbID = "00000000  00 e4 cd 98                                       |....|"
var refPerEnGnbIDlen32 = "00000000  50 e4 cd 9b 00                                    |P....|"

func createEngnbID() *e2sm_kpm_v2_go.EngnbId {

	return &e2sm_kpm_v2_go.EngnbId{
		EngnbId: &e2sm_kpm_v2_go.EngnbId_GNbId{
			GNbId: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x09, 0x00},
				Len:   22,
			},
		},
	}
}

func createEngnbIDlen32() *e2sm_kpm_v2_go.EngnbId {

	return &e2sm_kpm_v2_go.EngnbId{
		EngnbId: &e2sm_kpm_v2_go.EngnbId_GNbId{
			GNbId: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x09, 0x00},
				Len:   32,
			},
		},
	}
}

func Test_perEncodingEnGnbID(t *testing.T) {

	gnbIDc := createEngnbID()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "sizeExt")
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.EngnbId{}
	err = aper.UnmarshalWithParams(per, &result, "sizeExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("enGnbID PER - decoded\n%v", result)
}

func Test_perEnGnbIDCompareBytes(t *testing.T) {

	gnbIDc := createEngnbID()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "sizeExt")
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnGnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingEnGnbIDlen32(t *testing.T) {

	gnbIDc := createEngnbIDlen32()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "sizeExt")
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.EngnbId{}
	//ToDo - Encodes correctly, but doesn't decode now.. Problems with finding choice index..
	err = aper.UnmarshalWithParams(per, &result, "sizeExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("enGnbID PER - decoded\n%v", result)
}

func Test_perEnGnbIDlen32CompareBytes(t *testing.T) {

	gnbIDc := createEngnbIDlen32()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "sizeExt")
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnGnbIDlen32)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

// Doesn't work - deals with incorrect choice index.. :(
func Test_stupidExperiment3(t *testing.T) {
	perRefBytes, err := hexlib.DumpToByte(refPerEnGnbIDlen32)
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(perRefBytes))

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	result := e2sm_kpm_v2_go.EngnbId{}
	err = aper.UnmarshalWithParams(perRefBytes, &result, "sizeExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("enGnbID PER - decoded\n%v", result)
}
