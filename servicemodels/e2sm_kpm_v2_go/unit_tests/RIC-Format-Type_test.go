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

var refPerRicFormatType = "00000000  00 20                                             |. |"

func Test_perEncodingRicFormatType(t *testing.T) {

	ricFormatType := &e2smkpmv2.RicFormatType{
		Value: 32,
	}

	per, err := aper.MarshalWithParams(ricFormatType, "valueExt", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Format-Type PER\n%v", hex.Dump(per))

	result := e2smkpmv2.RicFormatType{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("RIC-Format-Type - decoded\n%v", &result)
	assert.Equal(t, ricFormatType.GetValue(), result.GetValue())
}

func Test_perRicFormatTypeCompareBytes(t *testing.T) {

	ricFormatType := &e2smkpmv2.RicFormatType{
		Value: 32,
	}

	per, err := aper.MarshalWithParams(ricFormatType, "valueExt", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Format-Type PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicFormatType)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
