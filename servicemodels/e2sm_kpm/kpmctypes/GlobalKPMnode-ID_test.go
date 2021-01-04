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

	plmnID := []byte{0x21, 0x18}
	var gNbCuUpID int64 = 123
	var gNbDuID int64 = 456

	idGlobalKpmnodeID1 := &e2sm_kpm_ies.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_ies.GlobalKpmnodeId_GNb{
			GNb: &e2sm_kpm_ies.GlobalKpmnodeGnbId{
				GlobalGNbId: &e2sm_kpm_ies.GlobalgNbId{
					PlmnId: &e2sm_kpm_ies.PlmnIdentity{
						Value: plmnID,
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
					Value: gNbCuUpID,
				},
				GNbDuId: &e2sm_kpm_ies.GnbDuId{
					Value: gNbDuID,
				},
			},
		},
	}

	idGlobalKpmnodeID2 := &e2sm_kpm_ies.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_ies.GlobalKpmnodeId_ENb{
			ENb: &e2sm_kpm_ies.GlobalKpmnodeEnbId{
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
			},
		},
	}

	idGlobalKpmnodeID3 := &e2sm_kpm_ies.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_ies.GlobalKpmnodeId_NgENb{
			NgENb: &e2sm_kpm_ies.GlobalKpmnodeNgEnbId{
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
			},
		},
	}

	idGlobalKpmnodeID4 := &e2sm_kpm_ies.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_ies.GlobalKpmnodeId_EnGNb{
			EnGNb: &e2sm_kpm_ies.GlobalKpmnodeEnGnbId{
				GlobalGNbId: &e2sm_kpm_ies.GlobalenGnbId{
					PLmnIdentity: &e2sm_kpm_ies.PlmnIdentity{
						Value: plmnID,
					},
					GNbId: &e2sm_kpm_ies.EngnbId{
						EngnbId: &e2sm_kpm_ies.EngnbId_GNbId{
							GNbId: &e2sm_kpm_ies.BitString{
								Value: 0x9bcd4, //uint64
								Len:   22,      //uint32
							},
						},
					},
				},
			},
		},
	}

	xer, err := xerEncodeGlobalKpmNodeID(idGlobalKpmnodeID1)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnode-ID XER\n%s", string(xer))

	xer, err = xerEncodeGlobalKpmNodeID(idGlobalKpmnodeID2)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnode-ID XER\n%s", string(xer))

	xer, err = xerEncodeGlobalKpmNodeID(idGlobalKpmnodeID3)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnode-ID XER\n%s", string(xer))

	xer, err = xerEncodeGlobalKpmNodeID(idGlobalKpmnodeID4)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnode-ID XER\n%s", string(xer))
}
