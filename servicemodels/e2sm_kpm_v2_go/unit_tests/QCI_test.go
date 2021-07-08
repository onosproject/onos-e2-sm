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

var refPerQci = "00000000  00 20                                             |. |"

func Test_perEncodingQci(t *testing.T) {

	qci := &e2sm_kpm_v2_go.Qci{
		Value: 32,
	}

	per, err := aper.Marshal(*qci)
	assert.NilError(t, err)
	t.Logf("QCI PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.Qci{}
	err = aper.Unmarshal(per, &result)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("QCI PER - decoded\n%v", result)
}

func Test_perQciCompareBytes(t *testing.T) {

	qci := &e2sm_kpm_v2_go.Qci{
		Value: 32,
	}

	per, err := aper.Marshal(*qci)
	assert.NilError(t, err)
	t.Logf("QCI PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerQci)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
