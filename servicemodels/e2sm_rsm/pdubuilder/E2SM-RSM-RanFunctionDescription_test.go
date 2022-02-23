// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
	su, err := CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandSliceUpdate())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, su)
	sc, err := CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandSliceCreate())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, sc)
	sd, err := CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandSliceDelete())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, sd)
	ueassoc, err := CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandUeAssociate())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, ueassoc)
	etd, err := CreateSupportedSlicingConfigItem(CreateE2SmRsmCommandEventTriggers())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, etd)

	slicingCapability, err := CreateSlicingCapabilityItem(71, 27, CreateSlicingTypeDynamic(), 10,
		supportedConfigList)
	assert.NilError(t, err)

	slicingCapList := make([]*e2sm_rsm_ies.NodeSlicingCapabilityItem, 0)
	slicingCapList = append(slicingCapList, slicingCapability)

	rfd, err := CreateE2SmRsmRanFunctionDescription("E2SM-RSM",
		"1.3.6.1.4.1.53148.1.1.2.102", "RAN Slicing Service Model", slicingCapList)
	assert.NilError(t, err)
	assert.Assert(t, rfd != nil)
	rfd.GetRanFunctionName().SetRanFunctionInstance(2)
	t.Logf("Created E2SM-RSM-RanFunctionDescription is \n%v", rfd)

	// APER encoding validation
	per, err := encoder.PerEncodeE2SmRsmRanFunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("E2SM-RSM-RanFunctionDescription PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRsmRanFunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-RSM-RanFunctionDescription PER - decoded\n%v", result)
	assert.DeepEqual(t, rfd.String(), result.String())
}
