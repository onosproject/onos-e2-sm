// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoEventTriggerDefinition(t *testing.T) {
	var rtPeriod int32 = 12

	E2SmMhoPdu, err := CreateE2SmMhoEventTriggerDefinitionPeriodic(rtPeriod)
	assert.NilError(t, err)
	assert.Assert(t, E2SmMhoPdu != nil)

	E2SmMhoPduUponChange, err := CreateE2SmMhoEventTriggerDefinitionUponRcvMeasReport()
	assert.NilError(t, err)
	assert.Assert(t, E2SmMhoPduUponChange != nil)

}
