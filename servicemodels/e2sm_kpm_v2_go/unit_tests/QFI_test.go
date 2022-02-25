// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerQfi = "00000000  40                                                |@|"

func Test_perEncodingQfi(t *testing.T) {

	qfi := &e2smkpmv2.Qfi{
		Value: 32,
	}

	per, err := aper.Marshal(qfi, nil, nil)
	assert.NilError(t, err)
	t.Logf("QFI PER\n%v", hex.Dump(per))

	result := e2smkpmv2.Qfi{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("QFI PER - decoded\n%v", &result)
	assert.Equal(t, qfi.GetValue(), result.GetValue())
}

func Test_perQfiCompareBytes(t *testing.T) {

	qfi := &e2smkpmv2.Qfi{
		Value: 32,
	}

	per, err := aper.Marshal(qfi, nil, nil)
	assert.NilError(t, err)
	t.Logf("QFI PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerQfi)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
