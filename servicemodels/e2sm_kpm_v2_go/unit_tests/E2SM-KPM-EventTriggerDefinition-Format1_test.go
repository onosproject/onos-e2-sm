// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerEventTriggerDefinitionFormat1 = "00000000  00 0e                                             |..|"

func createE2SMKPMEventTriggerDefinitionFormat1() *e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinitionFormat1 {

	var rtPeriod int32 = 15

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmEventTriggerDefinition(rtPeriod)

	return newE2SmKpmPdu.GetEventDefinitionFormat1()
}

func Test_perEncodingE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	etdf1 := createE2SMKPMEventTriggerDefinitionFormat1()

	per, err := aper.MarshalWithParams(etdf1, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-EventTriggerDefinition-Format1 PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinitionFormat1{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-EventTriggerDefinition-Format1 PER - decoded\n%v", result)

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
