// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreActionDefinition(t *testing.T) {
	var ActionDef int32 = 12

	newE2SmRcPrePdu, err := CreateE2SmRcPreActionDefinition(ActionDef)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
}
