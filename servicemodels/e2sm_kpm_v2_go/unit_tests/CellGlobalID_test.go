// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"gotest.tools/assert"
	"os"
	"testing"
)

var refPerCellGlobalID1 = "00000000  40 4f 4e 46 d4 bc 09 00                           |@ONF....|"
var refPerCellGlobalID2 = "00000000  00 4f 4e 46 d4 bc 09 00  f0                       |.ONF.....|"

func TestMain(m *testing.M) {
	log := logging.GetLogger("asn1")
	log.SetLevel(logging.DebugLevel)
	os.Exit(m.Run())
}

func createCellGlobalID1() *e2sm_kpm_v2_go.CellGlobalId {

	return &e2sm_kpm_v2_go.CellGlobalId{
		CellGlobalId: &e2sm_kpm_v2_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_kpm_v2_go.Eutracgi{
				EUtracellIdentity: &e2sm_kpm_v2_go.EutracellIdentity{
					Value: &asn1.BitString{
						Value: []byte{0xd4, 0xbc, 0x09, 0x00},
						Len:   28,
					},
				},
				PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
					Value: []byte("ONF"),
				},
			},
		},
	}
}

func createCellGlobalID2() *e2sm_kpm_v2_go.CellGlobalId {

	return &e2sm_kpm_v2_go.CellGlobalId{
		CellGlobalId: &e2sm_kpm_v2_go.CellGlobalId_NrCgi{
			NrCgi: &e2sm_kpm_v2_go.Nrcgi{
				NRcellIdentity: &e2sm_kpm_v2_go.NrcellIdentity{
					Value: &asn1.BitString{
						Value: []byte{0xd4, 0xbc, 0x09, 0x00, 0xf0},
						Len:   36,
					},
				},
				PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
					Value: []byte("ONF"),
				},
			},
		},
	}
}

func Test_perEncodingCellGlobalID1(t *testing.T) {

	cellGlobalID := createCellGlobalID1()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(cellGlobalID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("CellGlobalID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.CellGlobalId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("CellGlobalID PER - decoded\n%v", &result)
	assert.DeepEqual(t, cellGlobalID.GetEUtraCgi().GetPLmnIdentity().GetValue(), result.GetEUtraCgi().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, cellGlobalID.GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue(), result.GetEUtraCgi().GetEUtracellIdentity().GetValue().GetValue())
	assert.Equal(t, cellGlobalID.GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen(), result.GetEUtraCgi().GetEUtracellIdentity().GetValue().GetLen())
}

func Test_perCellGlobalID1CompareBytes(t *testing.T) {

	cellGlobalID := createCellGlobalID1()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(cellGlobalID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("CellGlobalID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerCellGlobalID1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

func Test_perEncodingCellGlobalID2(t *testing.T) {

	cellGlobalID := createCellGlobalID2()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(cellGlobalID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("CellGlobalID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.CellGlobalId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("CellGlobalID PER - decoded\n%v", &result)
	assert.DeepEqual(t, cellGlobalID.GetEUtraCgi().GetPLmnIdentity().GetValue(), result.GetEUtraCgi().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, cellGlobalID.GetNrCgi().GetNRcellIdentity().GetValue().GetValue(), result.GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, cellGlobalID.GetNrCgi().GetNRcellIdentity().GetValue().GetLen(), result.GetNrCgi().GetNRcellIdentity().GetValue().GetLen())
}

func Test_perCellGlobalID2CompareBytes(t *testing.T) {

	cellGlobalID := createCellGlobalID2()

	//aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(cellGlobalID, "valueExt", e2sm_kpm_v2_go.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("CellGlobalID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerCellGlobalID2)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
