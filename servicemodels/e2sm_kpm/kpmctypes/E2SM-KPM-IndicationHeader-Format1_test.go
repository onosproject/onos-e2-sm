// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeE2SmKpmIndicationHeaderFormat1(t *testing.T) {

	var plmnID = "ONF"
	var gNbCuUpId int64 = 123
	var gNbDuId int64 = 456
	var plmnIDnrcgi = "onf"
	var sst = "1"
	var sd = "SD1"
	var fiveQi int32 = 0
	var qCi int32 = 0

	indicationHeaderFormat1 := &e2sm_kpm_ies.E2SmKpmIndicationHeaderFormat1{
		IdGlobalKpmnodeId: &e2sm_kpm_ies.GlobalKpmnodeId{
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
						Value: gNbCuUpId, //int64
					},
					GNbDuId: &e2sm_kpm_ies.GnbDuId{
						Value: gNbDuId, //int64
					},
				},
			},
		},
		NRcgi: &e2sm_kpm_ies.Nrcgi{
			PLmnIdentity: &e2sm_kpm_ies.PlmnIdentity{
				Value: []byte(plmnIDnrcgi),
			},
			NRcellIdentity: &e2sm_kpm_ies.NrcellIdentity{
				Value: &e2sm_kpm_ies.BitString{
					Value: 0x9bcd4, //uint64
					Len:   22,      //uint32
				},
			},
		},
		PLmnIdentity: &e2sm_kpm_ies.PlmnIdentity{
			Value: []byte(plmnID),
		},
		SliceId: &e2sm_kpm_ies.Snssai{
			SSt: []byte(sst),
			SD:  []byte(sd),
		},
		FiveQi: fiveQi, //int32
		Qci:    qCi,    //int32
	}

	xer, err := xerEncodeE2SmKpmIndicationHeaderFormat1(indicationHeaderFormat1)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnodeId XER\n%s", string(xer))
}
