// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package unit_test

import (
	"encoding/hex"
	"github.com/google/martian/log"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

// Set of reference bytes obtained from MHO SM with CGo
// Set of negative values
var rsrpN65536 string = "00000000  00 00                                             |..|"
var rsrpN32000 string = "00000000  20 83 00                                          | ..|"
var rsrpN1234 string = "00000000  20 fb 2e                                          | ..|"
var rsrpN156 string = "00000000  20 ff 64                                          | .d|"
var rsrpN1 string = "00000000  20 ff ff                                          | ..|"

//Set of positive values
var rsrp0 string = "00000000  40 01 00 00                                       |@...|"
var rsrp1 string = "00000000  40 01 00 01                                       |@...|"
var rsrp156 string = "00000000  40 01 00 9c                                       |@...|"
var rsrp1234 string = "00000000  40 01 04 d2                                       |@...|"
var rsrp32000 string = "00000000  40 01 7d 00                                       |@.}.|"
var rsrp65536 string = "00000000  40 02 00 00                                       |@...|"

func init() {
	log.SetLevel(log.Debug)
}

func TestRsrpSample(t *testing.T) {

	//refPer, err := hexlib.DumpToByte(rsrpN65536)
	//assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: -1234,
	}

	//t.Logf("Casting value to uint64 results in \n%v", uint64(-156))
	t.Logf("Casting value to uint64 results in \n%v", uint64(rsrp.GetValue()))

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	//assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrpN65536(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrpN65536)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: -65536,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrpN32000(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrpN32000)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: -32000,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrpN1234(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrpN1234)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: -1234,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrpN156(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrpN156)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: -156,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrpN1(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrpN1)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: -1,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrp0(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrp0)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: 0,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrp1(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrp1)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: 1,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrp156(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrp156)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: 156,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrp1234(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrp1234)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: 1234,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrp32000(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrp32000)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: 32000,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}

func TestRsrp65536(t *testing.T) {

	refPer, err := hexlib.DumpToByte(rsrp65536)
	assert.NilError(t, err)

	rsrp := &e2sm_mho_go.Rsrp{
		Value: 65536,
	}

	aper.ChoiceMap = e2sm_mho_go.MhoChoicemap
	per, err := aper.MarshalWithParams(rsrp, "")
	assert.NilError(t, err)
	t.Logf("Encoded RSRP PER bytes are\n%v", hex.Dump(per))
	assert.DeepEqual(t, per, refPer)

	result := e2sm_mho_go.Rsrp{}
	err = aper.UnmarshalWithParams(per, &result, "")
	assert.NilError(t, err)
	t.Logf("Decoded RSRP from PER is\n%v", &result)

	assert.Equal(t, rsrp.GetValue(), result.GetValue())
	assert.Equal(t, rsrp.String(), result.String())
}
