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

//ToDo - find out why encoder prepends with three zero bytes..
var refPerMUeIDL string = "00000000  00 00 00 06 53 6f 6d 65  55 45                    |....SomeUE|"

func Test_perEncodeMatchingUeIDList(t *testing.T) {

	muei := &e2sm_kpm_v2_go.MatchingUeidItem{
		UeId: &e2sm_kpm_v2_go.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}
	muel := &e2sm_kpm_v2_go.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2_go.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*muei, "")
	assert.NilError(t, err)
	t.Logf("MatchingUeIDList PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MatchingUeidList{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MatchingUeIDList PER - decoded\n%v", result)
}

func Test_perMatchingUeIDListCompareBytes(t *testing.T) {

	muei := &e2sm_kpm_v2_go.MatchingUeidItem{
		UeId: &e2sm_kpm_v2_go.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}
	muel := &e2sm_kpm_v2_go.MatchingUeidList{
		Value: make([]*e2sm_kpm_v2_go.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*muei, "")
	assert.NilError(t, err)
	t.Logf("MatchingUeIDList PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMUeIDL)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

// It actually decodes valid set of bytes..
func Test_stupidExperiment1(t *testing.T) {
	perRefBytes, err := hexlib.DumpToByte(refPerMUeIDL)
	assert.NilError(t, err)
	t.Logf("MatchingUeIDList PER\n%v", hex.Dump(perRefBytes))

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	result := e2sm_kpm_v2_go.MatchingUeidList{}
	err = aper.UnmarshalWithParams(perRefBytes, &result, "")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MatchingUeIDList PER - decoded\n%v", result)
}
