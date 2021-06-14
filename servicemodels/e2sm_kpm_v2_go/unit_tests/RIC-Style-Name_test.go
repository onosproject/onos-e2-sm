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

var refPerRicStyleName = "00000000  03 80 53 6f 6d 65 4e 61  6d 65                    |..SomeName|"

func Test_perEncodingRicStyleName(t *testing.T) {

	ricStyleName := e2sm_kpm_v2_go.RicStyleName{
		Value: "SomeName",
	}

	per, err := aper.MarshalWithParams(ricStyleName, "sizeExt")
	assert.NilError(t, err)
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RicStyleName{}
	err = aper.UnmarshalWithParams(per, &result, "sizeExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("RIC-Style-Name - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicStyleName)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
