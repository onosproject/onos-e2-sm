// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrsm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	e2smv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRsmIndicationMessageFormat1(ueID *e2smrsm.UeIdentity, cuUeID int64, duUeID int64,
	emmcase e2smrsm.Emmcase, ulSm []*e2smrsm.SliceMetrics, dlSm []*e2smrsm.SliceMetrics) (*e2smrsm.E2SmRsmIndicationMessage, error) {

	if len(ulSm) == 0 {
		return nil, errors.NewInvalid("UL Slicing Metrics list should have at least 1 item")
	}

	if len(dlSm) == 0 {
		return nil, errors.NewInvalid("DL Slicing Metrics list should have at least 1 item")
	}

	im := e2smrsm.E2SmRsmIndicationMessage{
		E2SmRsmIndicationMessage: &e2smrsm.E2SmRsmIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: &e2smrsm.E2SmRsmIndicationMessageFormat1{
				UeId: ueID,
				CuUeF1ApId: &e2smrsm.CuUeF1ApId{
					Value: cuUeID,
				},
				DuUeF1ApId: &e2smrsm.DuUeF1ApId{
					Value: duUeID,
				},
				EmmCase:          emmcase,
				UlSlicingMetrics: ulSm,
				DlSlicingMetrics: dlSm,
			},
		},
	}

	if err := im.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmIndicationMessageFormat1(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return &im, nil
}

func CreateE2SmRsmIndicationMessageFormat2(tt e2smrsm.RsmEmmTriggerType, ueIDlist []*e2smrsm.UeIdentity,
	pt e2smrsm.UeIdType, bearerList []*e2smrsm.BearerId) (*e2smrsm.E2SmRsmIndicationMessage, error) {

	if len(ueIDlist) == 0 {
		return nil, errors.NewInvalid("UE ID list should contain at least 1 element. All elements of the list should be different UE identifiers")
	}

	if len(bearerList) < 1 || len(bearerList) > 32 {
		return nil, errors.NewInvalid("BearerID list should have 1 to 32 items")
	}

	im := &e2smrsm.E2SmRsmIndicationMessage{
		E2SmRsmIndicationMessage: &e2smrsm.E2SmRsmIndicationMessage_IndicationMessageFormat2{
			IndicationMessageFormat2: &e2smrsm.E2SmRsmIndicationMessageFormat2{
				TriggerType:       tt,
				UeIdlist:          ueIDlist,
				PrefferedUeIdtype: pt,
				BearerId:          bearerList,
			},
		},
	}

	if err := im.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmIndicationMessageFormat2(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return im, nil
}

func CreateUeIDCuUeF1ApID(val int64) (*e2smrsm.UeIdentity, error) {

	ueID := &e2smrsm.UeIdentity{
		UeIdentity: &e2smrsm.UeIdentity_CuUeF1ApId{
			CuUeF1ApId: &e2smrsm.CuUeF1ApId{
				Value: val,
			},
		},
	}

	if err := ueID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDCuUeF1ApID(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return ueID, nil
}

func CreateUeIDDuUeF1ApID(val int64) (*e2smrsm.UeIdentity, error) {

	ueID := &e2smrsm.UeIdentity{
		UeIdentity: &e2smrsm.UeIdentity_DuUeF1ApId{
			DuUeF1ApId: &e2smrsm.DuUeF1ApId{
				Value: val,
			},
		},
	}

	if err := ueID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDDuUeF1ApID(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return ueID, nil
}

func CreateUeIDRanUeNgapID(val int64) (*e2smrsm.UeIdentity, error) {

	ueID := &e2smrsm.UeIdentity{
		UeIdentity: &e2smrsm.UeIdentity_RanUeNgapId{
			RanUeNgapId: &e2smrsm.RanUeNgapId{
				Value: val,
			},
		},
	}

	if err := ueID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDRanUeNgapID(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return ueID, nil
}

func CreateUeIDAmfUeNgapID(val int64) (*e2smrsm.UeIdentity, error) {

	ueID := &e2smrsm.UeIdentity{
		UeIdentity: &e2smrsm.UeIdentity_AmfUeNgapId{
			AmfUeNgapId: &e2smv2.AmfUeNgapId{
				Value: val,
			},
		},
	}

	if err := ueID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDAmfUeNgapID(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return ueID, nil
}

func CreateUeIDEnbUeS1ApID(val int32) (*e2smrsm.UeIdentity, error) {

	ueID := &e2smrsm.UeIdentity{
		UeIdentity: &e2smrsm.UeIdentity_EnbUeS1ApId{
			EnbUeS1ApId: &e2smrsm.EnbUeS1ApId{
				Value: val,
			},
		},
	}

	if err := ueID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDEnbUeS1ApID(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return ueID, nil
}

func CreateEmmCaseAttach() e2smrsm.Emmcase {
	return e2smrsm.Emmcase_EMMCASE_ATTACHED
}

func CreateEmmCaseDetach() e2smrsm.Emmcase {
	return e2smrsm.Emmcase_EMMCASE_DETACHED
}

func CreateSliceMetrics(prb int32, ueToSlice int32, bler int32, cqi int32) (*e2smrsm.SliceMetrics, error) {

	if prb < 0 || prb > 100 {
		return nil, errors.NewInvalid("PRB utilization value should be within range 0 to 100")
	}

	if bler < 0 || bler > 100 {
		return nil, errors.NewInvalid("BLER value should be within range 0 to 100")
	}

	if cqi < 0 || cqi > 15 {
		return nil, errors.NewInvalid("average CQI value should be within range 0 to 15")
	}

	sm := &e2smrsm.SliceMetrics{
		PrbUtilization:    prb,
		NumUeAssocToSlice: ueToSlice,
		SliceLevelBler:    bler,
		AvgCqi:            cqi,
	}

	if err := sm.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSliceMetrics(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return sm, nil
}
