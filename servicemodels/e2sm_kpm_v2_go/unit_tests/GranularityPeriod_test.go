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

var refPerGP = "00000000  00 14                                             |..|"

func Test_perEncodingGranularityPeriod(t *testing.T) {

	gp := &e2sm_kpm_v2_go.GranularityPeriod{
		Value: 21,
	}

	per, err := aper.Marshal(gp)
	assert.NilError(t, err)
	t.Logf("GranularityPeriod PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GranularityPeriod{}
	err = aper.Unmarshal(per, &result)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GranularityPeriod PER - decoded\n%v", result)
}

func Test_perGranularityPeriodCompareBytes(t *testing.T) {

	gp := &e2sm_kpm_v2_go.GranularityPeriod{
		Value: 21,
	}

	per, err := aper.Marshal(gp)
	assert.NilError(t, err)
	t.Logf("GranularityPeriod PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGP)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
