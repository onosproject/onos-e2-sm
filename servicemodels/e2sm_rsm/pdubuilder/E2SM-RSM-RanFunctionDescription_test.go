// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/encoder"
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
	t.Logf("Created E2SM-RSM-RanFunctionDescription is \n%v", rfd)

	//ToDo - embed APER encoding validation
	per, err := encoder.PerEncodeE2SmRsmRanFunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-RanFunctionDescription PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-RanFunctionDescription PER - decoded\n%v", result)
	assert.DeepEqual(t, rfd.String(), result.String())
}
