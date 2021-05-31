// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

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

var refPer = "00000000  40 4f 4e 46 d4 bc 09 00                           |@ONF....|"

func TestMain(m *testing.M) {
	log := logging.GetLogger("asn1")
	log.SetLevel(logging.DebugLevel)
	os.Exit(m.Run())
}

func createCellGlobalID() *e2sm_kpm_v2_go.CellGlobalId {

	return &e2sm_kpm_v2_go.CellGlobalId{
		CellGlobalId: &e2sm_kpm_v2_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_kpm_v2_go.Eutracgi{
				EUtracellIdentity: &e2sm_kpm_v2_go.EutracellIdentity{
					Value: &asn1.BitString{
						Value: 0x9bcd4,
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

func Test_perDecodeCellGlobalID(t *testing.T) {

	cellGlobalID := createCellGlobalID()

	per, err := aper.MarshalWithParams(*cellGlobalID, "valueExt")
	assert.NilError(t, err)
	//assert.Equal(t, 8, len(per))
	t.Logf("CellGlobalID PER\n%s", hex.Dump(per))

	result := aper.UnmarshalWithParams(per, e2sm_kpm_v2_go.CellGlobalId{}, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CellGlobalID PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPer)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

