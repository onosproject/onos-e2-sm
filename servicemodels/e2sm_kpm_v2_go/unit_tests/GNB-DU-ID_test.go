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

var refPerGnbDuID = "00000000  20 04 d2                                          | ..|"
var refPerGnbDuIDhigh = "00000000  40 01 00 01                                          | ..|"
var refPerGnbDuIDub = "00000000  80 0f ff ff ff ff                                 |......|"

func createGnbDuID() *e2sm_kpm_v2_go.GnbDuId {

	return &e2sm_kpm_v2_go.GnbDuId{
		Value: 1234,
	}
}

func createGnbDuIDhigh() *e2sm_kpm_v2_go.GnbDuId {

	return &e2sm_kpm_v2_go.GnbDuId{
		Value: 65537,
	}
}

func createGnbDuIDub() *e2sm_kpm_v2_go.GnbDuId {

	return &e2sm_kpm_v2_go.GnbDuId{
		Value: 68719476735,
	}
}

func Test_perEncodingGnbDuID(t *testing.T) {

	gnbDuID := createGnbDuID()

	per, err := aper.Marshal(*gnbDuID)
	assert.NilError(t, err)
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GnbDuId{}
	err = aper.Unmarshal(per, &result)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GnbDuID PER - decoded\n%v", result)
}

func Test_perGnbDuIDCompareBytes(t *testing.T) {

	gnbDuID := createGnbDuID()

	per, err := aper.Marshal(*gnbDuID)
	assert.NilError(t, err)
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGnbDuID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingGnbDuIDhigh(t *testing.T) {

	gnbDuID := createGnbDuIDhigh()

	per, err := aper.Marshal(*gnbDuID)
	assert.NilError(t, err)
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GnbDuId{}
	err = aper.Unmarshal(per, &result)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GnbDuID PER - decoded\n%v", result)
}

func Test_perGnbDuIDhighCompareBytes(t *testing.T) {

	gnbDuID := createGnbDuIDhigh()

	per, err := aper.Marshal(*gnbDuID)
	assert.NilError(t, err)
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGnbDuIDhigh)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingGnbDuIDub(t *testing.T) {

	gnbDuID := createGnbDuIDub()

	per, err := aper.Marshal(*gnbDuID)
	assert.NilError(t, err)
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GnbDuId{}
	err = aper.Unmarshal(per, &result)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GnbDuID PER - decoded\n%v", result)
}

func Test_perGnbDuIDubCompareBytes(t *testing.T) {

	gnbDuID := createGnbDuIDub()

	per, err := aper.Marshal(*gnbDuID)
	assert.NilError(t, err)
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGnbDuIDub)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
