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

func createEngnbID() *e2sm_kpm_v2_go.EngnbId {

	return &e2sm_kpm_v2_go.EngnbId{
		EngnbId: &e2sm_kpm_v2_go.EngnbId_GNbId{
			GNbId: &asn1.BitString{
				Value: 0x9bcde4,
				Len:   22,
			},
		},
	}
}

func Test_perEncodingEnGnbID(t *testing.T) {

	gnbIDc := createEngnbID()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*gnbIDc, "valueExt")
	assert.NilError(t, err)
	t.Logf("enGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.EngnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("enGnbID PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEnGnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
