package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createAStruct() *test_sm_ies.Astruct {

	astruct := &test_sm_ies.Astruct{}

	item1 := createTestNestedExtensionNoExt()
	astruct.Item = append(astruct.Item, item1)
	astruct.Item = append(astruct.Item, item1)

	return astruct
}

func Test_xerEncodingAStruct(t *testing.T) {

	astruct := createAStruct()

	xer, err := xerEncodeAStruct(astruct)
	assert.NilError(t, err)
	t.Logf("AStruct XER\n%s", string(xer))

	result, err := xerDecodeAStruct(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("AStruct XER - decoded\n%v", result)
	assert.Equal(t, astruct.GetItem()[0].Item1, result.GetItem()[0].Item1)
	assert.DeepEqual(t, astruct.GetItem()[0].Item2, result.GetItem()[0].Item2)
	assert.Equal(t, astruct.GetItem()[1].Item1, result.GetItem()[1].Item1)
	assert.DeepEqual(t, astruct.GetItem()[1].Item2, result.GetItem()[1].Item2)
}

func Test_perEncodingAStruct(t *testing.T) {

	astruct := createAStruct()

	per, err := perEncodeAStruct(astruct)
	assert.NilError(t, err)
	t.Logf("AStruct PER\n%v", hex.Dump(per))

	// Generating APER bytes with Go APER lib
	//perNew, err := aper.MarshalWithParams(astruct, "valueExt", test_sm_ies.Choicemap, nil)
	//assert.NilError(t, err)

	//Comparing bytes against each other
	//assert.DeepEqual(t, per, perNew)

	result, err := perDecodeAStruct(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("AStruct PER - decoded\n%v", result)
	assert.Equal(t, astruct.GetItem()[0].Item1, result.GetItem()[0].Item1)
	assert.DeepEqual(t, astruct.GetItem()[0].Item2, result.GetItem()[0].Item2)
	assert.Equal(t, astruct.GetItem()[1].Item1, result.GetItem()[1].Item1)
	assert.DeepEqual(t, astruct.GetItem()[1].Item2, result.GetItem()[1].Item2)
}
