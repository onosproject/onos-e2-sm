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

var refPerCellObjectID = "00000000  00 00 03 31 32 33                                 |...123|"

func createCellObjectID() *e2sm_kpm_v2_go.CellObjectId {

	return &e2sm_kpm_v2_go.CellObjectId{
		Value: "123",
	}
}

func Test_perEncodingCellObjectID(t *testing.T) {

	coID := createCellObjectID()

	per, err := aper.MarshalWithParams(*coID, "valueExt")
	assert.NilError(t, err)
	t.Logf("CellObjectID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.CellObjectId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("CellObjectID PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerCellObjectID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
