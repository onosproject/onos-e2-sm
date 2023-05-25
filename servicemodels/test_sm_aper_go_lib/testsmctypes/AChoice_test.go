package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createAChoice1() *test_sm_ies.Achoice {

	alist := createAList()

	return &test_sm_ies.Achoice{
		Achoice: &test_sm_ies.Achoice_Ch1{
			Ch1: alist,
		},
	}
}

func createAChoice2() *test_sm_ies.Achoice {

	astruct := createAStruct()

	return &test_sm_ies.Achoice{
		Achoice: &test_sm_ies.Achoice_Ch2{
			Ch2: astruct,
		},
	}
}

func Test_xerEncodingAChoice(t *testing.T) {

	achoice1 := createAChoice1()

	xer, err := xerEncodeAChoice(achoice1)
	assert.NilError(t, err)
	t.Logf("AChoice (Option1) XER\n%s", string(xer))

	result, err := xerDecodeAChoice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("AChoice (Option1) XER - decoded\n%v", result)
	assert.Equal(t, achoice1.GetCh1().GetItem()[0].Item1, result.GetCh1().GetItem()[0].Item1)
	assert.DeepEqual(t, achoice1.GetCh1().GetItem()[0].Item2, result.GetCh1().GetItem()[0].Item2)
	assert.Equal(t, achoice1.GetCh1().GetItem()[1].Item1, result.GetCh1().GetItem()[1].Item1)
	assert.DeepEqual(t, achoice1.GetCh1().GetItem()[1].Item2, result.GetCh1().GetItem()[1].Item2)

	achoice2 := createAChoice2()

	xer2, err := xerEncodeAChoice(achoice2)
	assert.NilError(t, err)
	t.Logf("AChoice (Option2) XER\n%s", string(xer2))

	result2, err := xerDecodeAChoice(xer2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("AChoice (Option2) XER - decoded\n%v", result2)
	assert.Equal(t, achoice2.GetCh2().GetItem()[0].Item1, result2.GetCh2().GetItem()[0].Item1)
	assert.DeepEqual(t, achoice2.GetCh2().GetItem()[0].Item2, result2.GetCh2().GetItem()[0].Item2)
	assert.Equal(t, achoice2.GetCh2().GetItem()[1].Item1, result2.GetCh2().GetItem()[1].Item1)
	assert.DeepEqual(t, achoice2.GetCh2().GetItem()[1].Item2, result2.GetCh2().GetItem()[1].Item2)
}

func Test_perEncodingAChoice(t *testing.T) {

	achoice1 := createAChoice1()

	per, err := perEncodeAChoice(achoice1)
	assert.NilError(t, err)
	t.Logf("AChoice (Option1) PER\n%v", hex.Dump(per))

	// Generating APER bytes with Go APER lib
	//perNew, err := aper.MarshalWithParams(achoice1, "choiceExt", test_sm_ies.Choicemap, nil)
	//assert.NilError(t, err)

	//Comparing bytes against each other
	//assert.DeepEqual(t, per, perNew)

	result, err := perDecodeAChoice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("AChoice (Option1) PER - decoded\n%v", result)
	assert.Equal(t, achoice1.GetCh1().GetItem()[0].Item1, result.GetCh1().GetItem()[0].Item1)
	assert.DeepEqual(t, achoice1.GetCh1().GetItem()[0].Item2, result.GetCh1().GetItem()[0].Item2)
	assert.Equal(t, achoice1.GetCh1().GetItem()[1].Item1, result.GetCh1().GetItem()[1].Item1)
	assert.DeepEqual(t, achoice1.GetCh1().GetItem()[1].Item2, result.GetCh1().GetItem()[1].Item2)

	achoice2 := createAChoice2()

	per2, err := perEncodeAChoice(achoice2)
	assert.NilError(t, err)
	t.Logf("AChoice (Option2) PER\n%v", hex.Dump(per2))

	// Generating APER bytes with Go APER lib
	//perNew2, err := aper.MarshalWithParams(achoice2, "choiceExt", test_sm_ies.Choicemap, nil)
	//assert.NilError(t, err)

	//Comparing bytes against each other
	//assert.DeepEqual(t, per2, perNew2)

	result2, err := perDecodeAChoice(per2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("AChoice (Option2) PER - decoded\n%v", result2)
	assert.Equal(t, achoice2.GetCh2().GetItem()[0].Item1, result2.GetCh2().GetItem()[0].Item1)
	assert.DeepEqual(t, achoice2.GetCh2().GetItem()[0].Item2, result2.GetCh2().GetItem()[0].Item2)
	assert.Equal(t, achoice2.GetCh2().GetItem()[1].Item1, result2.GetCh2().GetItem()[1].Item1)
	assert.DeepEqual(t, achoice2.GetCh2().GetItem()[1].Item2, result2.GetCh2().GetItem()[1].Item2)
}
