// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmRcPreControlHeader(t *testing.T) {
	e2SmRcPreControlHeader, err := pdubuilder.CreateE2SmRcPreControlHeader()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlHeader(e2SmRcPreControlHeader)
	assert.NilError(t, err)
	assert.Equal(t, 235, len(xer))
	t.Logf("E2SM-RC-PRE-ControlHeader XER\n%s", string(xer))
}

func Test_XerDecodeE2SmRcPreControlHeader(t *testing.T) {

	e2SmRcPreControlHeader, err := pdubuilder.CreateE2SmRcPreControlHeader()
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmRcPreControlHeader(e2SmRcPreControlHeader)
	assert.NilError(t, err)
	assert.Equal(t, 235, len(xer))
	t.Logf("E2SM-RC-PRE-ControlHeader XER\n%s", string(xer))

	result, err := XerDecodeE2SmRcPreControlHeader(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-RC-PRE-ControlHeader XER\n%s", result)

	//assert.Equal(t, e2SmRcPreControlHeader.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}
