// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerRicStyleName1 = "00000000  00 00 4f                                          |..O|"
var refPerRicStyleName2 = "00000000  03 80 53 6f 6d 65 4e 61  6d 65                    |..SomeName|"
var refPerRicStyleName3 = "00000000  4a 00 54 68 69 73 49 73  4f 6e 65 48 75 6e 64 72  |J.ThisIsOneHundr|\n" +
	"00000010  65 64 46 69 66 74 79 43  68 61 72 73 4c 6f 6e 67  |edFiftyCharsLong|\n" +
	"00000020  53 74 72 69 6e 67 57 68  69 63 68 53 75 70 70 6f  |StringWhichSuppo|\n" +
	"00000030  73 65 64 54 6f 54 65 73  74 55 70 70 65 72 42 6f  |sedToTestUpperBo|\n" +
	"00000040  75 6e 64 4f 66 52 69 63  53 74 79 6c 65 4e 61 6d  |undOfRicStyleNam|\n" +
	"00000050  65 53 74 72 75 63 74 75  72 65 2e 49 74 49 73 49  |eStructure.ItIsI|\n" +
	"00000060  6d 70 6f 72 74 61 6e 74  46 6f 72 46 75 72 74 68  |mportantForFurth|\n" +
	"00000070  65 72 41 50 45 52 47 6f  4c 69 62 72 61 72 79 56  |erAPERGoLibraryV|\n" +
	"00000080  65 72 69 66 69 63 61 74  69 6f 6e 2c 69 73 69 74  |erification,isit|\n" +
	"00000090  3f 3f 3f 3f 3f 3f 3f                              |???????|"
var refPerRicStyleName4 = "00000000  80 80 9d 54 68 69 73 49  73 4f 6e 65 48 75 6e 64  |...ThisIsOneHund|\n" +
	"00000010  72 65 64 46 69 66 74 79  43 68 61 72 73 4c 6f 6e  |redFiftyCharsLon|\n" +
	"00000020  67 53 74 72 69 6e 67 57  68 69 63 68 53 75 70 70  |gStringWhichSupp|\n" +
	"00000030  6f 73 65 64 54 6f 54 65  73 74 55 70 70 65 72 42  |osedToTestUpperB|\n" +
	"00000040  6f 75 6e 64 4f 66 52 69  63 53 74 79 6c 65 4e 61  |oundOfRicStyleNa|\n" +
	"00000050  6d 65 53 74 72 75 63 74  75 72 65 2e 49 74 49 73  |meStructure.ItIs|\n" +
	"00000060  49 6d 70 6f 72 74 61 6e  74 46 6f 72 46 75 72 74  |ImportantForFurt|\n" +
	"00000070  68 65 72 41 50 45 52 47  6f 4c 69 62 72 61 72 79  |herAPERGoLibrary|\n" +
	"00000080  56 65 72 69 66 69 63 61  74 69 6f 6e 2c 69 73 69  |Verification,isi|\n" +
	"00000090  74 3f 3f 3f 3f 3f 3f 3f  45 78 74 65 6e 64 65 64  |t???????Extended|\n"

func Test_perEncodingRicStyleName1(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2_go.RicStyleName{
		Value: "O",
	}

	per, err := aper.MarshalWithParams(ricStyleName, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RicStyleName{}
	err = aper.UnmarshalWithParams(per, &result, "", nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("RIC-Style-Name - decoded\n%v", &result)
	assert.Equal(t, ricStyleName.GetValue(), result.GetValue())
}

func Test_perRicStyleName1CompareBytes(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2_go.RicStyleName{
		Value: "O",
	}

	per, err := aper.MarshalWithParams(ricStyleName, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicStyleName1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingRicStyleName2(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2_go.RicStyleName{
		Value: "SomeName",
	}

	per, err := aper.MarshalWithParams(ricStyleName, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RicStyleName{}
	err = aper.UnmarshalWithParams(per, &result, "", nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("RIC-Style-Name - decoded\n%v", &result)
	assert.Equal(t, ricStyleName.GetValue(), result.GetValue())
}

func Test_perRicStyleNameCompareBytes2(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2_go.RicStyleName{
		Value: "SomeName",
	}

	per, err := aper.MarshalWithParams(ricStyleName, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicStyleName2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingRicStyleName3(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2_go.RicStyleName{
		Value: "ThisIsOneHundredFiftyCharsLongStringWhichSupposedToTestUpperBoundOfRicStyleNameStructure.ItIsImportantForFurtherAPERGoLibraryVerification,isit???????",
	}

	per, err := aper.MarshalWithParams(ricStyleName, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RicStyleName{}
	err = aper.UnmarshalWithParams(per, &result, "", nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("RIC-Style-Name - decoded\n%v", &result)
	assert.Equal(t, ricStyleName.GetValue(), result.GetValue())
}

func Test_perRicStyleNameCompareBytes3(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2_go.RicStyleName{
		Value: "ThisIsOneHundredFiftyCharsLongStringWhichSupposedToTestUpperBoundOfRicStyleNameStructure.ItIsImportantForFurtherAPERGoLibraryVerification,isit???????",
	}

	per, err := aper.MarshalWithParams(ricStyleName, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicStyleName3)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingRicStyleName4(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2_go.RicStyleName{
		Value: "ThisIsOneHundredFiftyCharsLongStringWhichSupposedToTestUpperBoundOfRicStyleNameStructure.ItIsImportantForFurtherAPERGoLibraryVerification,isit???????Extended",
	}

	per, err := aper.MarshalWithParams(ricStyleName, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.RicStyleName{}
	err = aper.UnmarshalWithParams(per, &result, "", nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("RIC-Style-Name - decoded\n%v", &result)
	assert.Equal(t, ricStyleName.GetValue(), result.GetValue())
}

func Test_perRicStyleNameCompareBytes4(t *testing.T) {

	ricStyleName := &e2sm_kpm_v2_go.RicStyleName{
		Value: "ThisIsOneHundredFiftyCharsLongStringWhichSupposedToTestUpperBoundOfRicStyleNameStructure.ItIsImportantForFurtherAPERGoLibraryVerification,isit???????Extended",
	}

	per, err := aper.MarshalWithParams(ricStyleName, "", nil, nil)
	assert.NilError(t, err)
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerRicStyleName4)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
