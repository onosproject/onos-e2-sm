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

var refPerGlobalGnbID = "00000000  00 21 22 23 00 d4 bc 08                           |.!\"#....|"

func createGlobalgNbID() *e2sm_kpm_v2_go.GlobalgNbId {

	return &e2sm_kpm_v2_go.GlobalgNbId{
		PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		GnbId: &e2sm_kpm_v2_go.GnbIdChoice{
			GnbIdChoice: &e2sm_kpm_v2_go.GnbIdChoice_GnbId{
				GnbId: &asn1.BitString{
					Value: []byte{0xd4, 0xbc, 0x09, 0x00},
					Len:   22,
				},
			},
		},
	}
}

func Test_perEncodingGlobalGnbID(t *testing.T) {

	gnbIDc := createGlobalgNbID()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalgNbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GlobalGnbID PER - decoded\n%v", result)
}

func Test_perGlobalGnbIDCompareBytes(t *testing.T) {

	gnbIDc := createGlobalgNbID()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalGnbID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalGnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
