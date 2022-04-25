//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcCallProcessID(t *testing.T) {

	msg, err := CreateE2SmRcCallProcessIDFormat1(11)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcCallProcessId(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcCallProcessId(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}
