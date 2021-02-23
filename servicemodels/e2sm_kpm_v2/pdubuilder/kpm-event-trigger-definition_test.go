// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmEventTriggerDefinition(t *testing.T) {
	var rtPeriod int32 = 15

	newE2SmKpmPdu, err := CreateE2SmKpmEventTriggerDefinition(rtPeriod)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
