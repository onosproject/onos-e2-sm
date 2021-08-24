// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"gotest.tools/assert"
	"testing"
)

func TestCreateE2SmRsmRanFunctionDescription(t *testing.T) {

	supportedConfigList := make([]*e2sm_rsm_ies.SupportedSlicingConfigItem, 0)
	supportedConfigList = append(supportedConfigList, CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandSliceUpdate()))
	supportedConfigList = append(supportedConfigList, CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandSliceCreate()))
	supportedConfigList = append(supportedConfigList, CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandSliceDelete()))
	supportedConfigList = append(supportedConfigList, CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandUeAssociate()))
	supportedConfigList = append(supportedConfigList, CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandEventTriggers()))

	slicingCapability, err := CreateSlicingCapabilityItem(71, 27, CreateSlicingTypeDynamic(), 10,
		supportedConfigList)
	assert.NilError(t, err)

	slicingCapList := make([]*e2sm_rsm_ies.NodeSlicingCapabilityItem, 0)
	slicingCapList = append(slicingCapList, slicingCapability)

	rfd := CreateE2SmRsmRanFunctionDescription("E2SM-RSM",
		"1.3.6.1.4.1.53148.1.1.2.102", "RAN Slicing Service Model", slicingCapList)
	assert.Assert(t, rfd != nil)
	t.Logf("Created E2SM-RSM-RanFunctionDescription is \n%s", rfd.String())

	//ToDo - embed APER encoding validation
}
