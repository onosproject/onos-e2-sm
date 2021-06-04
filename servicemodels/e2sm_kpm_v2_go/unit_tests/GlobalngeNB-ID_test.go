// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerGlobalNgEnbID = "00000000  00 21 22 23 00 e4 cd 90  e4 cd 80 e4 cd 98        |.!\"#..........|"

func Test_perEncodingGlobalNgEnbID(t *testing.T) {

	bs := asn1.BitString{
		Value: 0x9bcde4,
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := asn1.BitString{
		Value: 0x9bcde4,
		Len:   18,
	}
	longMacroEnbID := asn1.BitString{
		Value: 0x9bcde4,
		Len:   21,
	}
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(&bs, plmnID, &shortMacroEnbID, &longMacroEnbID)
	ngeNbID.GetNgENb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(ngeNbID.GetNgENb().GetGlobalNgENbId(), "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalNgEnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalngeNbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GlobalNgEnbID PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalNgEnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
