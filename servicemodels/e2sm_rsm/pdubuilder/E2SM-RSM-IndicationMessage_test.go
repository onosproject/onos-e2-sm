// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"gotest.tools/assert"
	"testing"
)

func TestCreateE2SmRsmIndicationMessage(t *testing.T) {

	ueID := CreateUeIDAmfUeNgapID(1)

	ulSm := make([]*e2sm_rsm_ies.SliceMetrics, 0)
	ulM1, err := CreateSliceMetrics(100, 100, 100, 15)
	ulSm = append(ulSm, ulM1)
	ulM2, err := CreateSliceMetrics(10, 1, 50, 7)
	ulSm = append(ulSm, ulM2)
	ulM3, err := CreateSliceMetrics(47, 16, 12, 1)
	ulSm = append(ulSm, ulM3)

	dlSm := make([]*e2sm_rsm_ies.SliceMetrics, 0)
	dlM1, err := CreateSliceMetrics(91, 31, 54, 11)
	dlSm = append(ulSm, dlM1)
	dlM2, err := CreateSliceMetrics(84, 37, 41, 6)
	dlSm = append(ulSm, dlM2)

	im, err := CreateE2SmRsmIndicationMessageFormat1(ueID, 27, 14, CreateEmmCaseAttach(),
		ulSm, dlSm)
	assert.NilError(t, err)
	assert.Assert(t, im != nil)
	t.Logf("Created E2SM-RSM-IndicationMessage is \n%v", im)

	//ToDo - embed APER encoding validation
}

func TestCreateE2SmRsmIndicationMessageErrors(t *testing.T) {

	ueID := CreateUeIDAmfUeNgapID(1)

	_, err := CreateSliceMetrics(101, 100, 100, 15)
	assert.ErrorContains(t, err, "PRB utilization value should be within range 0 to 100")

	_, err = CreateSliceMetrics(99, 100, -10, 15)
	assert.ErrorContains(t, err, "BLER value should be within range 0 to 100")

	_, err = CreateSliceMetrics(99, 100, 11, 17)
	assert.ErrorContains(t, err, "CQI value should be within range 0 to 15")

	_, err = CreateE2SmRsmIndicationMessageFormat1(ueID, 27, 14, CreateEmmCaseAttach(),
		nil, nil)
	assert.ErrorContains(t, err, "Slicing Metrics list should have at least 1 item.")
}
