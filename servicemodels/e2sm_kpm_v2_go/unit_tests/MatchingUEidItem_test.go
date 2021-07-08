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

var refPerMUeIDI string = "00000000  00 06 53 6f 6d 65 55 45                           |..SomeUE|"

func Test_perEncodeMatchingUeIDItem(t *testing.T) {

	muei := &e2sm_kpm_v2_go.MatchingUeidItem{
		UeId: &e2sm_kpm_v2_go.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*muei, "valueExt")
	assert.NilError(t, err)
	t.Logf("MatchingUeIDitem PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MatchingUeidItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MatchingUeIDItem PER - decoded\n%v", result)
}

func Test_perMatchingUeIDItemCompareBytes(t *testing.T) {

	muei := &e2sm_kpm_v2_go.MatchingUeidItem{
		UeId: &e2sm_kpm_v2_go.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*muei, "valueExt")
	assert.NilError(t, err)
	t.Logf("MatchingUeIDitem PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMUeIDI)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
