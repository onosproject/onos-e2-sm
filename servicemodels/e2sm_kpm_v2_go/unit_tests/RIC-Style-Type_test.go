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

var refPerRicStyleType = "00000000  00 6f                                             |.o|"

func Test_perEncodingRicStyleType(t *testing.T) {

	ricStyleType := e2sm_kpm_v2_go.RicStyleType{
		Value: 111,
	}

	per, err := aper.MarshalWithParams(ricStyleType, "valueExt")
	assert.NilError(t, err)
	t.Logf("RIC-Style-Type PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RicStyleType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("RIC-Style-Type PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicStyleType)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
