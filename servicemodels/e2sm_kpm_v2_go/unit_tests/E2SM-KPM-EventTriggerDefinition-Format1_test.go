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

func createE2SMKPMEventTriggerDefinitionFormat1() (*e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinitionFormat1, error) {

	var rtPeriod int64 = 15

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmEventTriggerDefinition(rtPeriod)
	if err != nil {
		return nil, err
	}

	return newE2SmKpmPdu.GetEventDefinitionFormats().GetEventDefinitionFormat1(), nil
}

func Test_perEncodingE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	etdf1, err := createE2SMKPMEventTriggerDefinitionFormat1()
	assert.NilError(t, err)

	per, err := aper.MarshalWithParams(etdf1, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-EventTriggerDefinition-Format1 PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinitionFormat1{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-EventTriggerDefinition-Format1 PER - decoded\n%v", result)
}

func Test_perE2SmKpmEventTriggerDefinitionFormat1CompareBytes(t *testing.T) {

	etdf1, err := createE2SMKPMEventTriggerDefinitionFormat1()
	assert.NilError(t, err)

	per, err := aper.MarshalWithParams(etdf1, "valueExt")
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-EventTriggerDefinition-Format1 PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
