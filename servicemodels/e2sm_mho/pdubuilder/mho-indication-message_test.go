// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/mhoctypes"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoIndicationMsg(t *testing.T) {

	ueID := &e2sm_mho.UeIdentity{
		Value: "1234",
	}
	rsrp := &e2sm_mho.Rsrp{
		Value: 1234,
	}
	newE2SmMhoPdu, err := CreateE2SmMhoIndicationMsgFormat1(ueID, rsrp)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)

	xer, err := mhoctypes.XerEncodeE2SmMhoIndicationMessage(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Message: %s", string(xer))

	per, err := mhoctypes.PerEncodeE2SmMhoIndicationMessage(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Indication Message: %v", hex.Dump(per))

}
