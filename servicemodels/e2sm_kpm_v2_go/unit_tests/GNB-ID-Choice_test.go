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

var refPerGnbIDchoice = "00000000  00 d4 bc 0c                                       |....|"
var refPerGnbIDchoiceLen30 = "00000000  40 d4 bc 0c fc                                    |@....|"

func createGnbIDChoice() *e2sm_kpm_v2_go.GnbIdChoice {

	return &e2sm_kpm_v2_go.GnbIdChoice{
		GnbIdChoice: &e2sm_kpm_v2_go.GnbIdChoice_GnbId{
			GnbId: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x0c},
				Len:   22,
			},
		},
	}
}

func createGnbIDChoiceLen30() *e2sm_kpm_v2_go.GnbIdChoice {

	return &e2sm_kpm_v2_go.GnbIdChoice{
		GnbIdChoice: &e2sm_kpm_v2_go.GnbIdChoice_GnbId{
			GnbId: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x0c, 0xfc},
				Len:   30,
			},
		},
	}
}

func Test_perEncodingGnbIDChoice(t *testing.T) {

	gnbIDc := createGnbIDChoice()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GnbIDchoice PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GnbIdChoice{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GnbIDchoice PER - decoded\n%v", result)
}

func Test_perGnbIDChoiceCompareBytes(t *testing.T) {

	gnbIDc := createGnbIDChoice()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GnbIDchoice PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGnbIDchoice)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingGnbIDChoiceLen30(t *testing.T) {

	gnbIDc := createGnbIDChoiceLen30()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GnbIDchoice PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GnbIdChoice{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GnbIDchoice PER - decoded\n%v", result)
}

func Test_perGnbIDChoiceLen30CompareBytes(t *testing.T) {

	gnbIDc := createGnbIDChoiceLen30()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GnbIDchoice PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGnbIDchoiceLen30)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
