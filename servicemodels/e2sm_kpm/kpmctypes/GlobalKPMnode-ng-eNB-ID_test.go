// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerDecodeGlobalKPMnodengeNBID(t *testing.T) {

	plmnID := []byte{0x21, 0x18}

	globalKPMnodeNgeNbID := &e2sm_kpm_ies.GlobalKpmnodeNgEnbId{
		GlobalNgENbId: &e2sm_kpm_ies.GlobalngeNbId{
			PlmnId: &e2sm_kpm_ies.PlmnIdentity{
				Value: plmnID,
			},
			EnbId: &e2sm_kpm_ies.EnbIdChoice{
				EnbIdChoice: &e2sm_kpm_ies.EnbIdChoice_EnbIdLongmacro{
					EnbIdLongmacro: &e2sm_kpm_ies.BitString{
						Value: 0x9bcd4, //uint64
						Len:   22,      //uint32
					},
				},
			},
		},
	}

	xer, err := xerEncodeGlobalKPMnodengeNBID(globalKPMnodeNgeNbID)
	assert.NilError(t, err)
	t.Logf("ENb XER\n%s", string(xer))

	result, err := xerDecodeGlobalKPMnodengeNBID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
