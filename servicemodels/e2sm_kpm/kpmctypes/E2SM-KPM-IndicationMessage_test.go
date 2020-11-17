// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeE2SmKpmIndicationMessage(t *testing.T) {

	var plmnID = "ONF"
	var cellIdentityValue uint64 = 0
	var cellIdentityLen uint32 = 0
	var ulTotalAvlblProbs int32 = 0
	var dlTotalAvlblProbs int32 = 0
	var fiveQi int32 = 0
	var dlPrbusage int32 = 0
	var ulPrbusage int32 = 0
	var qCi int32 = 0
	var qciDlPrbusage int32 = 0
	var qciUlPrbusage int32 = 0
	var sst = "1"
	var sd = "SD1"
	var gNbCuName = "OpenNetworking"
	var cuCpNumberActvts int32 = 0
	var ranContainer = "ranContainer"
	newE2SmKpmPdu, err := pdubuilder.CreateE2SmKpmIndicationMsg(plmnID, cellIdentityValue, cellIdentityLen, dlTotalAvlblProbs,
		ulTotalAvlblProbs, fiveQi, dlPrbusage, ulPrbusage, qCi, qciDlPrbusage, qciUlPrbusage, sst,
		sd, gNbCuName, cuCpNumberActvts, ranContainer)

	xer, err := xerEncodeE2SmKpmIndicationMessage(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationMessage XER\n%s", string(xer))
}
