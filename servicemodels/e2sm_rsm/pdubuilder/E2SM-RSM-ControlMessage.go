// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRsmControlMessageSliceCreate(config *e2sm_rsm_ies.SliceConfig) (*e2sm_rsm_ies.E2SmRsmControlMessage, error) {

	cm := &e2sm_rsm_ies.E2SmRsmControlMessage{
		E2SmRsmControlMessage: &e2sm_rsm_ies.E2SmRsmControlMessage_SliceCreate{
			SliceCreate: config,
		},
	}

	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmControlMessageSliceCreate(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return cm, nil
}

func CreateE2SmRsmControlMessageSliceUpdate(config *e2sm_rsm_ies.SliceConfig) (*e2sm_rsm_ies.E2SmRsmControlMessage, error) {

	cm := &e2sm_rsm_ies.E2SmRsmControlMessage{
		E2SmRsmControlMessage: &e2sm_rsm_ies.E2SmRsmControlMessage_SliceUpdate{
			SliceUpdate: config,
		},
	}

	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmControlMessageSliceUpdate(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return cm, nil
}

func CreateE2SmRsmControlMessageSliceDelete(sliceID int64, sliceType e2sm_rsm_ies.SliceType) (*e2sm_rsm_ies.E2SmRsmControlMessage, error) {

	cm := &e2sm_rsm_ies.E2SmRsmControlMessage{
		E2SmRsmControlMessage: &e2sm_rsm_ies.E2SmRsmControlMessage_SliceDelete{
			SliceDelete: &e2sm_rsm_ies.SliceDelete{
				SliceId: &e2sm_rsm_ies.SliceId{
					Value: sliceID,
				},
				SliceType: sliceType,
			},
		},
	}

	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmControlMessageSliceDelete(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return cm, nil
}

func CreateE2SmRsmControlMessageSliceAssociate(config *e2sm_rsm_ies.SliceAssociate) (*e2sm_rsm_ies.E2SmRsmControlMessage, error) {

	cm := &e2sm_rsm_ies.E2SmRsmControlMessage{
		E2SmRsmControlMessage: &e2sm_rsm_ies.E2SmRsmControlMessage_SliceAssociate{
			SliceAssociate: config,
		},
	}

	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmControlMessageSliceAssociate(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return cm, nil
}

func CreateSliceConfig(sliceID int64, parameters *e2sm_rsm_ies.SliceParameters, sliceType e2sm_rsm_ies.SliceType) (*e2sm_rsm_ies.SliceConfig, error) {

	sc := &e2sm_rsm_ies.SliceConfig{
		SliceId: &e2sm_rsm_ies.SliceId{
			Value: sliceID,
		},
		SliceConfigParameters: parameters,
		SliceType:             sliceType,
	}

	if err := sc.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSliceConfig(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return sc, nil
}

func CreateSliceParameters(schType e2sm_rsm_ies.SchedulerType) (*e2sm_rsm_ies.SliceParameters, error) {

	sp := &e2sm_rsm_ies.SliceParameters{
		SchedulerType: schType,
	}

	if err := sp.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSliceParameters(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return sp, nil
}

func CreateScheduleConfigEmpty() *e2sm_rsm_ies.ScheduleConfig {
	return new(e2sm_rsm_ies.ScheduleConfig)
}

func CreateUlPowerControlEmpty() *e2sm_rsm_ies.UlpowerControl {
	return new(e2sm_rsm_ies.UlpowerControl)
}

func CreateLinkAdaptationEmpty() *e2sm_rsm_ies.LinkAdaptation {
	return new(e2sm_rsm_ies.LinkAdaptation)
}

func CreateHarqRextCapEmpty() *e2sm_rsm_ies.HarqrextCap {
	return new(e2sm_rsm_ies.HarqrextCap)
}

func CreateSchedulerTypeRoundRobin() e2sm_rsm_ies.SchedulerType {
	return e2sm_rsm_ies.SchedulerType_SCHEDULER_TYPE_ROUND_ROBIN
}

func CreateSchedulerTypeProportionallyFair() e2sm_rsm_ies.SchedulerType {
	return e2sm_rsm_ies.SchedulerType_SCHEDULER_TYPE_PROPORTIONALLY_FAIR
}

func CreateSchedulerTypeQosBased() e2sm_rsm_ies.SchedulerType {
	return e2sm_rsm_ies.SchedulerType_SCHEDULER_TYPE_QOS_BASED
}

func CreateAggregationLevelCapOne() e2sm_rsm_ies.AggregationLevelCap {
	return e2sm_rsm_ies.AggregationLevelCap_AGGREGATION_LEVEL_CAP_ONE
}

func CreateAggregationLevelCapTwo() e2sm_rsm_ies.AggregationLevelCap {
	return e2sm_rsm_ies.AggregationLevelCap_AGGREGATION_LEVEL_CAP_TWO
}

func CreateAggregationLevelCapFour() e2sm_rsm_ies.AggregationLevelCap {
	return e2sm_rsm_ies.AggregationLevelCap_AGGREGATION_LEVEL_CAP_FOUR
}

func CreateAggregationLevelCapEight() e2sm_rsm_ies.AggregationLevelCap {
	return e2sm_rsm_ies.AggregationLevelCap_AGGREGATION_LEVEL_CAP_EIGHT
}

func CreateAggregationLevelCapSixteen() e2sm_rsm_ies.AggregationLevelCap {
	return e2sm_rsm_ies.AggregationLevelCap_AGGREGATION_LEVEL_CAP_SIXTEEN
}

func CreateFeatureStatusEnable() e2sm_rsm_ies.FeatureStatus {
	return e2sm_rsm_ies.FeatureStatus_FEATURE_STATUS_ENABLE
}

func CreateFeatureStatusDisable() e2sm_rsm_ies.FeatureStatus {
	return e2sm_rsm_ies.FeatureStatus_FEATURE_STATUS_DISABLE
}

func CreateCarrierAggregationLevelCapOne() e2sm_rsm_ies.CarrierAggregationLevelCap {
	return e2sm_rsm_ies.CarrierAggregationLevelCap_CARRIER_AGGREGATION_LEVEL_CAP_ONE
}

func CreateCarrierAggregationLevelCapTwo() e2sm_rsm_ies.CarrierAggregationLevelCap {
	return e2sm_rsm_ies.CarrierAggregationLevelCap_CARRIER_AGGREGATION_LEVEL_CAP_TWO
}

func CreateCarrierAggregationLevelCapThree() e2sm_rsm_ies.CarrierAggregationLevelCap {
	return e2sm_rsm_ies.CarrierAggregationLevelCap_CARRIER_AGGREGATION_LEVEL_CAP_THREE
}

func CreateCarrierAggregationLevelCapFour() e2sm_rsm_ies.CarrierAggregationLevelCap {
	return e2sm_rsm_ies.CarrierAggregationLevelCap_CARRIER_AGGREGATION_LEVEL_CAP_FOUR
}

func CreateSliceAssociate(ueID *e2sm_rsm_ies.UeIdentity, bearerIDlist []*e2sm_rsm_ies.BearerId,
	dlSliceID int64) (*e2sm_rsm_ies.SliceAssociate, error) {

	if len(bearerIDlist) < 1 || len(bearerIDlist) > 32 {
		return nil, errors.NewInvalid("BearerID list should have 1 to 32 items")
	}

	sa := &e2sm_rsm_ies.SliceAssociate{
		UeId:     ueID,
		BearerId: bearerIDlist,
		DownLinkSliceId: &e2sm_rsm_ies.SliceIdassoc{
			Value: dlSliceID,
		},
	}

	if err := sa.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSliceAssociate(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return sa, nil
}

func CreateRiCapOne() e2sm_rsm_ies.RiCap {
	return e2sm_rsm_ies.RiCap_RI_CAP_ONE
}

func CreateRiCapTwo() e2sm_rsm_ies.RiCap {
	return e2sm_rsm_ies.RiCap_RI_CAP_TWO
}

func CreateTransmissionModeOne() e2sm_rsm_ies.TransmissionMode {
	return e2sm_rsm_ies.TransmissionMode_TRANSMISSION_MODE_ONE
}

func CreateTransmissionModeTwo() e2sm_rsm_ies.TransmissionMode {
	return e2sm_rsm_ies.TransmissionMode_TRANSMISSION_MODE_TWO
}

func CreateTransmissionModeThree() e2sm_rsm_ies.TransmissionMode {
	return e2sm_rsm_ies.TransmissionMode_TRANSMISSION_MODE_THREE
}

func CreateSliceTypeDL() e2sm_rsm_ies.SliceType {
	return e2sm_rsm_ies.SliceType_SLICE_TYPE_DL_SLICE
}

func CreateSliceTypeUL() e2sm_rsm_ies.SliceType {
	return e2sm_rsm_ies.SliceType_SLICE_TYPE_UL_SLICE
}
