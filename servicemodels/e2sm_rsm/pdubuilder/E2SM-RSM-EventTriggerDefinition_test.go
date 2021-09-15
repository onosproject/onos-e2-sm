// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/encoder"
	"gotest.tools/assert"
	"testing"
)

func Test_E2SmRsmEventTriggerDefinition(t *testing.T) {

	etd, err := CreateE2SmRsmEventTriggerDefinitionFormat1(CreateRsmRicindicationTriggerTypeUponEmmEvent())
	assert.NilError(t, err)
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", etd)

	// APER validation
	per, err := encoder.PerEncodeE2SmRsmEventTriggerDefinition(etd)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-EventTriggerDefinition PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmEventTriggerDefinition(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-EventTriggerDefinition PER - decoded\n%v", result)
	assert.DeepEqual(t, etd.String(), result.String())
}
