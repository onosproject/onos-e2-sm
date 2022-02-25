// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerEnGnbID = "00000000  00 d4 bc 0c                                       |....|"
var refPerEnGnbIDlen32 = "00000000  50 d4 bc 0c 98                                    |P....|"

func createEngnbID() *e2smkpmv2.EngnbId {

	return &e2smkpmv2.EngnbId{
		EngnbId: &e2smkpmv2.EngnbId_GNbId{
			GNbId: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x0c},
				Len:   22,
			},
		},
	}
}

func createEngnbIDlen32() *e2smkpmv2.EngnbId {

	return &e2smkpmv2.EngnbId{
		EngnbId: &e2smkpmv2.EngnbId_GNbId{
			GNbId: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x0c, 0x98},
				Len:   32,
			},
		},
	}
}

func Test_perEncodingEnGnbID(t *testing.T) {

	gnbIDc := createEngnbID()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(gnbIDc, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(per))

	result := e2smkpmv2.EngnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("enGnbID PER - decoded\n%v", &result)
	assert.DeepEqual(t, gnbIDc.GetGNbId().GetValue(), result.GetGNbId().GetValue())
	assert.Equal(t, gnbIDc.GetGNbId().GetLen(), result.GetGNbId().GetLen())
}

func Test_perEnGnbIDCompareBytes(t *testing.T) {

	gnbIDc := createEngnbID()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(gnbIDc, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnGnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingEnGnbIDlen32(t *testing.T) {

	gnbIDc := createEngnbIDlen32()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(gnbIDc, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(per))

	result := e2smkpmv2.EngnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("enGnbID PER - decoded\n%v", &result)
	assert.DeepEqual(t, gnbIDc.GetGNbId().GetValue(), result.GetGNbId().GetValue())
	assert.Equal(t, gnbIDc.GetGNbId().GetLen(), result.GetGNbId().GetLen())
}

func Test_perEnGnbIDlen32CompareBytes(t *testing.T) {

	gnbIDc := createEngnbIDlen32()

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(gnbIDc, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnGnbIDlen32)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_stupidExperiment3(t *testing.T) {
	perRefBytes, err := hexlib.DumpToByte(refPerEnGnbIDlen32)
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(perRefBytes))

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	result := e2smkpmv2.EngnbId{}
	err = aper.UnmarshalWithParams(perRefBytes, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("enGnbID PER - decoded\n%v", &result)
}
