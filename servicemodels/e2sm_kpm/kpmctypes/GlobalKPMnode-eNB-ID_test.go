// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerDecodeGlobalKPMnodeENbID(t *testing.T) {

	plmnID := []byte{0x21, 0x18}

	globalKPMnodeENbID := &e2sm_kpm_ies.GlobalKpmnodeEnbId{
		GlobalENbId: &e2sm_kpm_ies.GlobalEnbId{
			PLmnIdentity: &e2sm_kpm_ies.PlmnIdentity{
				Value: plmnID,
			},
			ENbId: &e2sm_kpm_ies.EnbId{
				EnbId: &e2sm_kpm_ies.EnbId_HomeENbId{
					HomeENbId: &e2sm_kpm_ies.BitString{
						Value: 0x9bcd4, //uint64
						Len:   22,      //uint32
					},
				},
			},
		},
	}

	xer, err := xerEncodeGlobalKPMnodeENbID(globalKPMnodeENbID)
	assert.NilError(t, err)
	t.Logf("GlobalENb XER\n%s", string(xer))

	result, err := xerDecodeGlobalKPMnodeENbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
