// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

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

//func TestMain(m *testing.M) {
//	log := logging.GetLogger("asn1")
//	log.SetLevel(logging.DebugLevel)
//	os.Exit(m.Run())
//}

func createEutracellIdentity() *e2sm_kpm_v2_go.EutracellIdentity {

	return &e2sm_kpm_v2_go.EutracellIdentity{
		Value: &asn1.BitString{
			Value: 0x9bcd4,
			Len:   28,
		},
	}
}

func Test_perDecodeEutracellIdentity(t *testing.T) {

	eCellID := createEutracellIdentity()

	per, err := aper.MarshalWithParams(*eCellID, "valueExt")
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("EutracellIdentity PER\n%s", hex.Dump(per))

	result := aper.UnmarshalWithParams(per, e2sm_kpm_v2_go.EutracellIdentity{}, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EutracellIdentity PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEutraCellIdentity)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
