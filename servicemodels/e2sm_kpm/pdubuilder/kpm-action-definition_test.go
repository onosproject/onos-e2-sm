// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmActionDefinition(t *testing.T) {
	var ActionDef int32 = 12

	newE2SmKpmPdu, err := CreateE2SmKpmActionDefinition(ActionDef)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
