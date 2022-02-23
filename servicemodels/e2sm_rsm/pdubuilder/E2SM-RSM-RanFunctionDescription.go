// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRsmRanFunctionDescription(rfSn string, rfE2SMoid string, rfd string,
	slicingCapability []*e2sm_rsm_ies.NodeSlicingCapabilityItem) (*e2sm_rsm_ies.E2SmRsmRanfunctionDescription, error) {

	e2SmRsmPdu := e2sm_rsm_ies.E2SmRsmRanfunctionDescription{
		RanFunctionName: &e2sm_v2_ies.RanfunctionName{
			RanFunctionShortName:   rfSn,
			RanFunctionE2SmOid:     rfE2SMoid,
			RanFunctionDescription: rfd,
		},
		RicSlicingNodeCapabilityList: slicingCapability,
	}

	if err := e2SmRsmPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRsmRanFunctionDescription(): error validating E2SmRsmRanfunctionDescription %s", err.Error())
	}
	return &e2SmRsmPdu, nil
}

func CreateSlicingCapabilityItem(maxDlSlice int32, maxUlSlice int32, slicingType e2sm_rsm_ies.SlicingType,
	maxUeSlice int32, supportedConfig []*e2sm_rsm_ies.SupportedSlicingConfigItem) (*e2sm_rsm_ies.NodeSlicingCapabilityItem, error) {

	if len(supportedConfig) < 1 || len(supportedConfig) > 5 {
		return nil, errors.NewInvalid("this least should contain up to 5 E2SM-RSM-Commands. Minimum requirement is to have at least 1 E2SM-RSM-Command")
	}

	item := e2sm_rsm_ies.NodeSlicingCapabilityItem{
		MaxNumberOfSlicesDl:    maxDlSlice,
		MaxNumberOfSlicesUl:    maxUlSlice,
		SlicingType:            slicingType,
		MaxNumberOfUesPerSlice: maxUeSlice,
		SupportedConfig:        supportedConfig,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSlicingCapabilityItem(): error validating E2SmRsmRanfunctionDescription %s", err.Error())
	}
	return &item, nil
}

func CreateSlicingTypeDynamic() e2sm_rsm_ies.SlicingType {
	return e2sm_rsm_ies.SlicingType_SLICING_TYPE_DYNAMIC
}

func CreateSlicingTypeStatic() e2sm_rsm_ies.SlicingType {
	return e2sm_rsm_ies.SlicingType_SLICING_TYPE_STATIC
}

func CreateSupportedSlicingConfigItem(configType e2sm_rsm_ies.E2SmRsmCommand) (*e2sm_rsm_ies.SupportedSlicingConfigItem, error) {

	item := e2sm_rsm_ies.SupportedSlicingConfigItem{
		SlicingConfigType: configType,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSupportedSlicingConfigItem(): error validating E2SmRsmRanfunctionDescription %s", err.Error())
	}
	return &item, nil
}

func CreateE2SmRsmCommandSliceCreate() e2sm_rsm_ies.E2SmRsmCommand {
	return e2sm_rsm_ies.E2SmRsmCommand_E2_SM_RSM_COMMAND_SLICE_CREATE
}

func CreateE2SmRsmCommandSliceUpdate() e2sm_rsm_ies.E2SmRsmCommand {
	return e2sm_rsm_ies.E2SmRsmCommand_E2_SM_RSM_COMMAND_SLICE_UPDATE
}

func CreateE2SmRsmCommandSliceDelete() e2sm_rsm_ies.E2SmRsmCommand {
	return e2sm_rsm_ies.E2SmRsmCommand_E2_SM_RSM_COMMAND_SLICE_DELETE
}

func CreateE2SmRsmCommandUeAssociate() e2sm_rsm_ies.E2SmRsmCommand {
	return e2sm_rsm_ies.E2SmRsmCommand_E2_SM_RSM_COMMAND_UE_ASSOCIATE
}

func CreateE2SmRsmCommandEventTriggers() e2sm_rsm_ies.E2SmRsmCommand {
	return e2sm_rsm_ies.E2SmRsmCommand_E2_SM_RSM_COMMAND_EVENT_TRIGGERS
}
