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

var refPerRanFunctionName = "00000000  41 80 4f 52 41 4e 2d 45  32 53 4d 2d 4b 50 4d 00  |A.ORAN-E2SM-KPM.|\n" +
	"00000010  00 18 31 2e 33 2e 36 2e  31 2e 34 2e 31 2e 35 33  |..1.3.6.1.4.1.53|\n" +
	"00000020  31 34 38 2e 31 2e 32 2e  32 2e 32 05 00 4b 50 4d  |148.1.2.2.2..KPM|\n" +
	"00000030  20 6d 6f 6e 69 74 6f 72  00 15                    | monitor..|\n"

var refPerRanFunctionNameNoOptional = "00000000  01 80 4f 52 41 4e 2d 45  32 53 4d 2d 4b 50 4d 00  |..ORAN-E2SM-KPM.|\n" +
	"00000010  00 18 31 2e 33 2e 36 2e  31 2e 34 2e 31 2e 35 33  |..1.3.6.1.4.1.53|\n" +
	"00000020  31 34 38 2e 31 2e 32 2e  32 2e 32 05 00 4b 50 4d  |148.1.2.2.2..KPM|\n" +
	"00000030  20 6d 6f 6e 69 74 6f 72                           | monitor|"

func createRanFunctionName() *e2sm_kpm_v2_go.RanfunctionName {

	var rfi int32 = 21
	return &e2sm_kpm_v2_go.RanfunctionName{
		RanFunctionShortName:   "ORAN-E2SM-KPM",
		RanFunctionE2SmOid:     "1.3.6.1.4.1.53148.1.2.2.2",
		RanFunctionDescription: "KPM monitor",
		RanFunctionInstance:    &rfi,
	}
}

func createRanFunctionNameNoOptional() *e2sm_kpm_v2_go.RanfunctionName {

	return &e2sm_kpm_v2_go.RanfunctionName{
		RanFunctionShortName:   "ORAN-E2SM-KPM",
		RanFunctionE2SmOid:     "1.3.6.1.4.1.53148.1.2.2.2",
		RanFunctionDescription: "KPM monitor",
	}
}

func Test_perEncodingRanFunctionName(t *testing.T) {

	rfn := createRanFunctionName()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*rfn, "valueExt")
	assert.NilError(t, err)
	t.Logf("RanFunctionName PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RanfunctionName{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("RanFunctionName PER - decoded\n%v", result)
}

func Test_perRanFunctionNameCompareBytes(t *testing.T) {

	rfn := createRanFunctionName()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*rfn, "valueExt")
	assert.NilError(t, err)
	t.Logf("RanFunctionName PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRanFunctionName)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingRanFunctionNameNoOptional(t *testing.T) {

	rfn := createRanFunctionNameNoOptional()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*rfn, "valueExt")
	assert.NilError(t, err)
	t.Logf("RanFunctionName PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RanfunctionName{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("RanFunctionName PER - decoded\n%v", result)
}

func Test_perRanFunctionNameCompareBytesNoOptional(t *testing.T) {

	rfn := createRanFunctionNameNoOptional()

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*rfn, "valueExt")
	assert.NilError(t, err)
	t.Logf("RanFunctionName PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRanFunctionNameNoOptional)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
