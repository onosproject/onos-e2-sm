// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGlobalKpmNodeId(t *testing.T) {

	var plmnID = "ONF"
	var gNbCuUpId int64 = 123
	var gNbDuId int64 = 456

	idGlobalKpmnodeId := &e2sm_kpm_ies.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_ies.GlobalKpmnodeId_GNb{
			GNb: &e2sm_kpm_ies.GlobalKpmnodeGnbId{
				GlobalGNbId: &e2sm_kpm_ies.GlobalgNbId{
					PlmnId: &e2sm_kpm_ies.PlmnIdentity{
						Value: []byte(plmnID),
					},
					GnbId: &e2sm_kpm_ies.GnbIdChoice{
						GnbIdChoice: &e2sm_kpm_ies.GnbIdChoice_GnbId{
							GnbId: &e2sm_kpm_ies.BitString{
								Value: 0x9bcd4, //uint64
								Len:   22,      //uint32
							},
						},
					},
				},
				GNbCuUpId: &e2sm_kpm_ies.GnbCuUpId{
					Value: gNbCuUpId,
				},
				GNbDuId: &e2sm_kpm_ies.GnbDuId{
					Value: gNbDuId,
				},
			},
		},
	}

	xer, err := xerEncodeGlobalKpmNodeId(idGlobalKpmnodeId)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnode-ID XER\n%s", string(xer))
}
