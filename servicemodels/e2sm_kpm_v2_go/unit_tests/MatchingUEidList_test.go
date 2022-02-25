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

var refPerMUeIDL = "00000000  00 00 00 06 53 6f 6d 65  55 45                    |....SomeUE|"

func Test_perEncodeMatchingUeIDList(t *testing.T) {

	muei := &e2smkpmv2.MatchingUeidItem{
		UeId: &e2smkpmv2.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}
	muel := &e2smkpmv2.MatchingUeidList{
		Value: make([]*e2smkpmv2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)
	//muel.Value = append(muel.Value, muei)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(muel, "", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MatchingUeIDList PER\n%v", hex.Dump(per))

	result := e2smkpmv2.MatchingUeidList{}
	err = aper.UnmarshalWithParams(per, &result, "", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MatchingUeIDList PER - decoded\n%v", &result)
	assert.DeepEqual(t, muel.GetValue()[0].GetUeId().GetValue(), result.GetValue()[0].GetUeId().GetValue())
}

func Test_perMatchingUeIDListCompareBytes(t *testing.T) {

	muei := &e2smkpmv2.MatchingUeidItem{
		UeId: &e2smkpmv2.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}
	muel := &e2smkpmv2.MatchingUeidList{
		Value: make([]*e2smkpmv2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(muel, "", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MatchingUeIDList PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMUeIDL)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_stupidExperiment1(t *testing.T) {
	perRefBytes, err := hexlib.DumpToByte(refPerMUeIDL)
	assert.NilError(t, err)
	t.Logf("MatchingUeIDList PER\n%v", hex.Dump(perRefBytes))

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	result := e2smkpmv2.MatchingUeidList{}
	err = aper.UnmarshalWithParams(perRefBytes, &result, "", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MatchingUeIDList PER - decoded\n%v", &result)
}
