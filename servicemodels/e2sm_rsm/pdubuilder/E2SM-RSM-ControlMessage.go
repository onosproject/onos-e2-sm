// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrsm "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRsmControlMessageSliceCreate(config *e2smrsm.SliceConfig) (*e2smrsm.E2SmRsmControlMessage, error) {

	cm := &e2smrsm.E2SmRsmControlMessage{
		E2SmRsmControlMessage: &e2smrsm.E2SmRsmControlMessage_SliceCreate{
			SliceCreate: config,
		},
	}

	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmControlMessageSliceCreate(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return cm, nil
}

func CreateE2SmRsmControlMessageSliceUpdate(config *e2smrsm.SliceConfig) (*e2smrsm.E2SmRsmControlMessage, error) {

	cm := &e2smrsm.E2SmRsmControlMessage{
		E2SmRsmControlMessage: &e2smrsm.E2SmRsmControlMessage_SliceUpdate{
			SliceUpdate: config,
		},
	}

	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmControlMessageSliceUpdate(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return cm, nil
}

func CreateE2SmRsmControlMessageSliceDelete(sliceID int64, sliceType e2smrsm.SliceType) (*e2smrsm.E2SmRsmControlMessage, error) {

	cm := &e2smrsm.E2SmRsmControlMessage{
		E2SmRsmControlMessage: &e2smrsm.E2SmRsmControlMessage_SliceDelete{
			SliceDelete: &e2smrsm.SliceDelete{
				SliceId: &e2smrsm.SliceId{
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

func CreateE2SmRsmControlMessageSliceAssociate(config *e2smrsm.SliceAssociate) (*e2smrsm.E2SmRsmControlMessage, error) {

	cm := &e2smrsm.E2SmRsmControlMessage{
		E2SmRsmControlMessage: &e2smrsm.E2SmRsmControlMessage_SliceAssociate{
			SliceAssociate: config,
		},
	}

	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmControlMessageSliceAssociate(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return cm, nil
}

func CreateSliceConfig(sliceID int64, parameters *e2smrsm.SliceParameters, sliceType e2smrsm.SliceType) (*e2smrsm.SliceConfig, error) {

	sc := &e2smrsm.SliceConfig{
		SliceId: &e2smrsm.SliceId{
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

func CreateSliceParameters(schType e2smrsm.SchedulerType) (*e2smrsm.SliceParameters, error) {

	sp := &e2smrsm.SliceParameters{
		SchedulerType: schType,
	}

	if err := sp.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSliceParameters(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return sp, nil
}

func CreateScheduleConfigEmpty() *e2smrsm.ScheduleConfig {
	return new(e2smrsm.ScheduleConfig)
}

func CreateUlPowerControlEmpty() *e2smrsm.UlpowerControl {
	return new(e2smrsm.UlpowerControl)
}

func CreateLinkAdaptationEmpty() *e2smrsm.LinkAdaptation {
	return new(e2smrsm.LinkAdaptation)
}

func CreateHarqRextCapEmpty() *e2smrsm.HarqrextCap {
	return new(e2smrsm.HarqrextCap)
}

func CreateSchedulerTypeRoundRobin() e2smrsm.SchedulerType {
	return e2smrsm.SchedulerType_SCHEDULER_TYPE_ROUND_ROBIN
}

func CreateSchedulerTypeProportionallyFair() e2smrsm.SchedulerType {
	return e2smrsm.SchedulerType_SCHEDULER_TYPE_PROPORTIONALLY_FAIR
}

func CreateSchedulerTypeQosBased() e2smrsm.SchedulerType {
	return e2smrsm.SchedulerType_SCHEDULER_TYPE_QOS_BASED
}

func CreateAggregationLevelCapOne() e2smrsm.AggregationLevelCap {
	return e2smrsm.AggregationLevelCap_AGGREGATION_LEVEL_CAP_ONE
}

func CreateAggregationLevelCapTwo() e2smrsm.AggregationLevelCap {
	return e2smrsm.AggregationLevelCap_AGGREGATION_LEVEL_CAP_TWO
}

func CreateAggregationLevelCapFour() e2smrsm.AggregationLevelCap {
	return e2smrsm.AggregationLevelCap_AGGREGATION_LEVEL_CAP_FOUR
}

func CreateAggregationLevelCapEight() e2smrsm.AggregationLevelCap {
	return e2smrsm.AggregationLevelCap_AGGREGATION_LEVEL_CAP_EIGHT
}

func CreateAggregationLevelCapSixteen() e2smrsm.AggregationLevelCap {
	return e2smrsm.AggregationLevelCap_AGGREGATION_LEVEL_CAP_SIXTEEN
}

func CreateFeatureStatusEnable() e2smrsm.FeatureStatus {
	return e2smrsm.FeatureStatus_FEATURE_STATUS_ENABLE
}

func CreateFeatureStatusDisable() e2smrsm.FeatureStatus {
	return e2smrsm.FeatureStatus_FEATURE_STATUS_DISABLE
}

func CreateCarrierAggregationLevelCapOne() e2smrsm.CarrierAggregationLevelCap {
	return e2smrsm.CarrierAggregationLevelCap_CARRIER_AGGREGATION_LEVEL_CAP_ONE
}

func CreateCarrierAggregationLevelCapTwo() e2smrsm.CarrierAggregationLevelCap {
	return e2smrsm.CarrierAggregationLevelCap_CARRIER_AGGREGATION_LEVEL_CAP_TWO
}

func CreateCarrierAggregationLevelCapThree() e2smrsm.CarrierAggregationLevelCap {
	return e2smrsm.CarrierAggregationLevelCap_CARRIER_AGGREGATION_LEVEL_CAP_THREE
}

func CreateCarrierAggregationLevelCapFour() e2smrsm.CarrierAggregationLevelCap {
	return e2smrsm.CarrierAggregationLevelCap_CARRIER_AGGREGATION_LEVEL_CAP_FOUR
}

func CreateSliceAssociate(ueID *e2smrsm.UeIdentity, bearerIDlist []*e2smrsm.BearerId,
	dlSliceID int64) (*e2smrsm.SliceAssociate, error) {

	if len(bearerIDlist) < 1 || len(bearerIDlist) > 32 {
		return nil, errors.NewInvalid("BearerID list should have 1 to 32 items")
	}

	sa := &e2smrsm.SliceAssociate{
		UeId:     ueID,
		BearerId: bearerIDlist,
		DownLinkSliceId: &e2smrsm.SliceIdassoc{
			Value: dlSliceID,
		},
	}

	if err := sa.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSliceAssociate(): error validating E2SM-RSM PDU %s", err.Error())
	}
	return sa, nil
}

func CreateRiCapOne() e2smrsm.RiCap {
	return e2smrsm.RiCap_RI_CAP_ONE
}

func CreateRiCapTwo() e2smrsm.RiCap {
	return e2smrsm.RiCap_RI_CAP_TWO
}

func CreateTransmissionModeOne() e2smrsm.TransmissionMode {
	return e2smrsm.TransmissionMode_TRANSMISSION_MODE_ONE
}

func CreateTransmissionModeTwo() e2smrsm.TransmissionMode {
	return e2smrsm.TransmissionMode_TRANSMISSION_MODE_TWO
}

func CreateTransmissionModeThree() e2smrsm.TransmissionMode {
	return e2smrsm.TransmissionMode_TRANSMISSION_MODE_THREE
}

func CreateSliceTypeDL() e2smrsm.SliceType {
	return e2smrsm.SliceType_SLICE_TYPE_DL_SLICE
}

func CreateSliceTypeUL() e2smrsm.SliceType {
	return e2smrsm.SliceType_SLICE_TYPE_UL_SLICE
}
