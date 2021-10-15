// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoEventTriggerDefinition(t *testing.T) {
	var rtPeriod int32 = 12

	newE2SmMhoPdu, err := CreateE2SmMhoEventTriggerDefinitionPeriodic(rtPeriod)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)

	per, err := encoder.PerEncodeE2SmMhoEventTriggerDefinition(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded E2SM-MHO-EventTriggerDefinition: \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMhoEventTriggerDefinition(per)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-MHO-EventTriggerDefinition is \n%v", result)
	assert.Equal(t, newE2SmMhoPdu.String(), result.String())

	newE2SmMhoPduUponChange, err := CreateE2SmMhoEventTriggerDefinitionUponRcvMeasReport()
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPduUponChange != nil)

	per1, err := encoder.PerEncodeE2SmMhoEventTriggerDefinition(newE2SmMhoPduUponChange)
	assert.NilError(t, err)
	t.Logf("PER Encoded E2SM-MHO-EventTriggerDefinition: \n%v", hex.Dump(per1))

	result1, err := encoder.PerDecodeE2SmMhoEventTriggerDefinition(per1)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-MHO-EventTriggerDefinition is \n%v", result)
	assert.Equal(t, newE2SmMhoPduUponChange.String(), result1.String())
}
