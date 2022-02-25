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

var refPerPlmnID = "00000000  21 22 23                                          |!\"#|"

func Test_perEncodingPlmnID(t *testing.T) {

	plmnID := &e2smkpmv2.PlmnIdentity{
		Value: []byte{0x21, 0x22, 0x23},
	}

	per, err := aper.Marshal(plmnID, nil, nil)
	assert.NilError(t, err)
	t.Logf("PLMN-Identity PER\n%v", hex.Dump(per))

	result := e2smkpmv2.PlmnIdentity{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("PLMN-Identity PER - decoded\n%v", &result)
	assert.DeepEqual(t, plmnID.GetValue(), result.GetValue())
}

func Test_perPlmnIDCompareBytes(t *testing.T) {

	plmnID := &e2smkpmv2.PlmnIdentity{
		Value: []byte{0x21, 0x22, 0x23},
	}

	per, err := aper.Marshal(plmnID, nil, nil)
	assert.NilError(t, err)
	t.Logf("PLMN-Identity PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerPlmnID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
