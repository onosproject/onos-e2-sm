// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/encoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerEventTriggerDefinition = "00000000  00 0e                                             |..|"

func createE2SMKPMEventTriggerDefinition() (*e2sm_kpm_v2_go.E2SmKpmEventTriggerDefinition, error) {

	var rtPeriod int64 = 15

	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmEventTriggerDefinition(rtPeriod)
	if err != nil {
		return nil, err
	}

	return newE2SmKpmPdu, nil
}

func Test_perEncodingE2SmKpmEventTriggerDefinition(t *testing.T) {

	etd, err := createE2SMKPMEventTriggerDefinition()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmEventTriggerDefinition(etd)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-EventTriggerDefinition PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmKpmEventTriggerDefinition(per)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("E2SM-KPM-EventTriggerDefinition PER - decoded\n%v", result)
}

func Test_perE2SmKpmEventTriggerDefinitionCompareBytes(t *testing.T) {

	etd, err := createE2SMKPMEventTriggerDefinition()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmKpmEventTriggerDefinition(etd)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-EventTriggerDefinition PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerEventTriggerDefinition)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
