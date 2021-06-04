// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerGlobalKPMnodeEnGnbID = "00000000  60 21 22 23 50 e4 cd 9b  00 00 2a 00 20           |`!\"#P.....*. |"

func Test_perEncodeGlobalKpmnodeEnGnbID(t *testing.T) {

	var bsValue uint64 = 0x9bcde4
	var bsLen uint32 = 32
	plmnID := []byte{0x21, 0x22, 0x23}
	var gnbDuID int64 = 32
	var gnbCuUpID int64 = 42

	enbID, err := pdubuilder.CreateGlobalKpmnodeIDenGNbID(bsValue, bsLen, plmnID)
	enbID.GetEnGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	enbID.GetEnGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(enbID.GetEnGNb(), "valueExt")
	assert.NilError(t, err)
	t.Logf("GlobalKPMnodeEnGnbID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.GlobalKpmnodeGnbId{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("GlobalKPMnodeEnGnbID PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerGlobalKPMnodeEnGnbID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
