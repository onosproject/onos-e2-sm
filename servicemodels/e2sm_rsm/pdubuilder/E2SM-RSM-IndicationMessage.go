// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"fmt"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies"
)

func CreateE2SmRsmIndicationMessageFormat1(ueID *e2sm_rsm_ies.UeIdentity, cuUeID int64, duUeID int64,
	emmcase e2sm_rsm_ies.Emmcase, ulSm []*e2sm_rsm_ies.SliceMetrics, dlSm []*e2sm_rsm_ies.SliceMetrics) (*e2sm_rsm_ies.E2SmRsmIndicationMessage, error) {

	if len(ulSm) == 0 {
		return nil, fmt.Errorf("UL Slicing Metrics list should have at least 1 item")
	}

	if len(dlSm) == 0 {
		return nil, fmt.Errorf("DL Slicing Metrics list should have at least 1 item")
	}

	im := e2sm_rsm_ies.E2SmRsmIndicationMessage{
		E2SmRsmIndicationMessage: &e2sm_rsm_ies.E2SmRsmIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: &e2sm_rsm_ies.E2SmRsmIndicationMessageFormat1{
				UeId: ueID,
				CuUeF1ApId: &e2sm_rsm_ies.CuUeF1ApId{
					Value: cuUeID,
				},
				DuUeF1ApId: &e2sm_rsm_ies.DuUeF1ApId{
					Value: duUeID,
				},
				EmmCase:          emmcase,
				UlSlicingMetrics: ulSm,
				DlSlicingMetrics: dlSm,
			},
		},
	}

	return &im, nil
}

func CreateE2SmRsmIndicationMessageFormat2(tt e2sm_rsm_ies.RsmEmmTriggerType, ueIDlist []*e2sm_rsm_ies.UeIdentity,
	pt e2sm_rsm_ies.UeIdType, bearerList []*e2sm_rsm_ies.BearerId) (*e2sm_rsm_ies.E2SmRsmIndicationMessage, error) {

	if len(ueIDlist) == 0 {
		return nil, fmt.Errorf("UE ID list should contain at least 1 element. All elements of the list should be different UE identifiers")
	}

	if len(bearerList) < 1 || len(bearerList) > 32 {
		return nil, fmt.Errorf("BearerID list should have 1 to 32 items")
	}

	return &e2sm_rsm_ies.E2SmRsmIndicationMessage{
		E2SmRsmIndicationMessage: &e2sm_rsm_ies.E2SmRsmIndicationMessage_IndicationMessageFormat2{
			IndicationMessageFormat2: &e2sm_rsm_ies.E2SmRsmIndicationMessageFormat2{
				TriggerType:       tt,
				UeIdlist:          ueIDlist,
				PrefferedUeIdtype: pt,
				BearerId:          bearerList,
			},
		},
	}, nil
}

func CreateUeIDCuUeF1ApID(val int64) *e2sm_rsm_ies.UeIdentity {

	return &e2sm_rsm_ies.UeIdentity{
		UeIdentity: &e2sm_rsm_ies.UeIdentity_CuUeF1ApId{
			CuUeF1ApId: &e2sm_rsm_ies.CuUeF1ApId{
				Value: val,
			},
		},
	}
}

func CreateUeIDDuUeF1ApID(val int64) *e2sm_rsm_ies.UeIdentity {

	return &e2sm_rsm_ies.UeIdentity{
		UeIdentity: &e2sm_rsm_ies.UeIdentity_DuUeF1ApId{
			DuUeF1ApId: &e2sm_rsm_ies.DuUeF1ApId{
				Value: val,
			},
		},
	}
}

func CreateUeIDRanUeNgapID(val int64) *e2sm_rsm_ies.UeIdentity {

	return &e2sm_rsm_ies.UeIdentity{
		UeIdentity: &e2sm_rsm_ies.UeIdentity_RanUeNgapId{
			RanUeNgapId: &e2sm_rsm_ies.RanUeNgapId{
				Value: val,
			},
		},
	}
}

func CreateUeIDAmfUeNgapID(val int64) *e2sm_rsm_ies.UeIdentity {

	return &e2sm_rsm_ies.UeIdentity{
		UeIdentity: &e2sm_rsm_ies.UeIdentity_AmfUeNgapId{
			AmfUeNgapId: &e2sm_v2_ies.AmfUeNgapId{
				Value: val,
			},
		},
	}
}

func CreateUeIDEnbUeS1ApID(val int32) *e2sm_rsm_ies.UeIdentity {

	return &e2sm_rsm_ies.UeIdentity{
		UeIdentity: &e2sm_rsm_ies.UeIdentity_EnbUeS1ApId{
			EnbUeS1ApId: &e2sm_rsm_ies.EnbUeS1ApId{
				Value: val,
			},
		},
	}
}

func CreateEmmCaseAttach() e2sm_rsm_ies.Emmcase {
	return e2sm_rsm_ies.Emmcase_EMMCASE_ATTACHED
}

func CreateEmmCaseDetach() e2sm_rsm_ies.Emmcase {
	return e2sm_rsm_ies.Emmcase_EMMCASE_DETACHED
}

func CreateSliceMetrics(prb int32, ueToSlice int32, bler int32, cqi int32) (*e2sm_rsm_ies.SliceMetrics, error) {

	if prb < 0 || prb > 100 {
		return nil, fmt.Errorf("PRB utilization value should be within range 0 to 100")
	}

	if bler < 0 || bler > 100 {
		return nil, fmt.Errorf("BLER value should be within range 0 to 100")
	}

	if cqi < 0 || cqi > 15 {
		return nil, fmt.Errorf("average CQI value should be within range 0 to 15")
	}

	return &e2sm_rsm_ies.SliceMetrics{
		PrbUtilization:    prb,
		NumUeAssocToSlice: ueToSlice,
		SliceLevelBler:    bler,
		AvgCqi:            cqi,
	}, nil
}
