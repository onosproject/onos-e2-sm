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

var refPerEutraCellIdentity = "00000000  d4 bc 09 00                                       |....|"

func createEutracellIdentity() *e2sm_kpm_v2_go.EutracellIdentity {

	return &e2sm_kpm_v2_go.EutracellIdentity{
		Value: &asn1.BitString{
			Value: 0x9bcd4,
			Len:   28,
		},
	}
}

func Test_perEncodingEutracellIdentity(t *testing.T) {

	eCellID := createEutracellIdentity()

	per, err := aper.Marshal(*eCellID)
	assert.NilError(t, err)
	t.Logf("EutraCellIdentity PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.EutracellIdentity{}
	err = aper.Unmarshal(per, &result)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("EutraCellIdentity PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEutraCellIdentity)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
