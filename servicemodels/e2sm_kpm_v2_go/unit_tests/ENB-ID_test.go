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

var refPerEnbIDmacro = "00000000  00 d4 bc 00                                       |....|"
var refPerEnbIDhome = "00000000  40 d4 bc 09 00                                    |@....|"


func createEnbIDMacro() *e2sm_kpm_v2_go.EnbId {

	return &e2sm_kpm_v2_go.EnbId{
		EnbId: &e2sm_kpm_v2_go.EnbId_MacroENbId{
			MacroENbId: &asn1.BitString{
				Value: 0x9bcd4,
				Len:   20,
			},
		},
	}
}

func createEnbIDHome() *e2sm_kpm_v2_go.EnbId {

	return &e2sm_kpm_v2_go.EnbId{
		EnbId: &e2sm_kpm_v2_go.EnbId_HomeENbId{
			HomeENbId: &asn1.BitString{
				Value: 0x9bcd4,
				Len:   28,
			},
		},
	}
}

func Test_perDecodeEnbID(t *testing.T) {

	enbID := createEnbIDMacro()

	per, err := aper.MarshalWithParams(enbID, "valueExt")
	assert.NilError(t, err)
	t.Logf("EnbID (Macro) PER\n%v", hex.Dump(per))

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	result := e2sm_kpm_v2_go.EnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("EnbID (Macro) PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnbIDmacro)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perDecodeEnbIDhome(t *testing.T) {

	enbID := createEnbIDHome()

	per, err := aper.MarshalWithParams(enbID, "valueExt")
	assert.NilError(t, err)
	t.Logf("EnbID (Home) PER\n%v", hex.Dump(per))

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	result := e2sm_kpm_v2_go.EnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("EnbID (Home) PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnbIDhome)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
