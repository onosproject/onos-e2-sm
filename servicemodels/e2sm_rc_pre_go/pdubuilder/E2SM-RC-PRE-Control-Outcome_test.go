// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/encoder"
	e2sm_rc_pre_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreControlOutcome(t *testing.T) {
	var ranParameterID int32 = 20

	elementList := make([]*e2sm_rc_pre_go.RanparameterItem, 0)
	elementList = append(elementList, CreateRanParameterItem(ranParameterID))

	newE2SmRcPrePdu, err := CreateE2SmRcPreControlOutcome(elementList)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err := encoder.PerEncodeE2SmRcPreControlOutcome(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Control Outcome: \n%v", hex.Dump(per))

	resultPer, err := encoder.PerDecodeE2SmRcPreControlOutcome(per)
	assert.NilError(t, err)
	t.Logf("PER Control Outcome - decoded \n%v", resultPer)
	assert.Equal(t, newE2SmRcPrePdu.String(), resultPer.String())
}
