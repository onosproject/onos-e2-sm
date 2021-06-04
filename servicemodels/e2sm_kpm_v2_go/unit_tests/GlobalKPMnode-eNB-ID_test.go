// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerGlobalKPMEnbID = "00000000  00 21 22 23 40 e4 cd 9b  00                       |.!\"#@....|"

func Test_perEncodingGlobalKPMnodeEnbID(t *testing.T) {

	bs := asn1.BitString{
		Value: 0x9bcde4,
		Len:   28,
	}
	plmnID := []byte{0x21, 0x22, 0x23}

	enbID, err := pdubuilder.CreateGlobalKpmnodeIDeNBID(&bs, plmnID)
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(*enbID.GetENb(), "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalKPMnodeEnbID (Home) PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalKpmnodeEnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GlobalKPMnodeEnbID (Home) PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalKPMEnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
