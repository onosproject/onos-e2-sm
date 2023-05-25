package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createAList() *test_sm_ies.Alist {

	alist := &test_sm_ies.Alist{}

	item1 := createTestNestedExtensionNoExt()
	alist.Item = append(alist.Item, item1)
	alist.Item = append(alist.Item, item1)

	return alist
}

func Test_xerEncodingAList(t *testing.T) {

	alist := createAList()

	xer, err := xerEncodeAList(alist)
	assert.NilError(t, err)
	t.Logf("AList XER\n%s", string(xer))

	result, err := xerDecodeAList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("AList XER - decoded\n%v", result)
	assert.Equal(t, alist.GetItem()[0].Item1, result.GetItem()[0].Item1)
	assert.DeepEqual(t, alist.GetItem()[0].Item2, result.GetItem()[0].Item2)
	assert.Equal(t, alist.GetItem()[1].Item1, result.GetItem()[1].Item1)
	assert.DeepEqual(t, alist.GetItem()[1].Item2, result.GetItem()[1].Item2)
}

func Test_perEncodingAList(t *testing.T) {

	alist := createAList()

	per, err := perEncodeAList(alist)
	assert.NilError(t, err)
	t.Logf("AList PER\n%v", hex.Dump(per))

	// Generating APER bytes with Go APER lib
	//perNew, err := aper.MarshalWithParams(alist, "valueExt", test_sm_ies.Choicemap, nil)
	//assert.NilError(t, err)

	//Comparing bytes against each other
	//assert.DeepEqual(t, per, perNew)

	result, err := perDecodeAList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("AList PER - decoded\n%v", result)
	assert.Equal(t, alist.GetItem()[0].Item1, result.GetItem()[0].Item1)
	assert.DeepEqual(t, alist.GetItem()[0].Item2, result.GetItem()[0].Item2)
	assert.Equal(t, alist.GetItem()[1].Item1, result.GetItem()[1].Item1)
	assert.DeepEqual(t, alist.GetItem()[1].Item2, result.GetItem()[1].Item2)
}
