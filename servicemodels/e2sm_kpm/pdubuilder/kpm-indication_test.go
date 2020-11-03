// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package sandbox

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmIndicationMsg(t *testing.T) {
	newE2SmKpmPdu, err := CreateE2SmKpmIndicationMsg("ONF", "1", "SD1", "ranContainer")
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
