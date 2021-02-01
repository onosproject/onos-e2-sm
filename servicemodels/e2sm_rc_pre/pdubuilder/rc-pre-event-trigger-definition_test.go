// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreEventTriggerDefinition(t *testing.T) {
	var rtPeriod int32 = 12 // range is from 0 to 19

	newE2SmRcPrePdu, err := CreateE2SmRcPreEventTriggerDefinition(rtPeriod)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
}
