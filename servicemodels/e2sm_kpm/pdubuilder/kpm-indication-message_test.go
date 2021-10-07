// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmIndicationMsg(t *testing.T) {
	var plmnID = "ONF"
	var cellIdentityValue uint64
	var cellIdentityLen uint32
	var ulTotalAvlblProbs int32
	var dlTotalAvlblProbs int32
	var fiveQi int32
	var dlPrbusage int32
	var ulPrbusage int32
	var qCi int32
	var qciDlPrbusage int32
	var qciUlPrbusage int32
	var sst = "1"
	var sd = "SD1"
	var gNbCuName = "OpenNetworking"
	var cuCpNumberActvts int32 = 1
	var ranContainer = "ranContainer"
	newE2SmKpmPdu, err := CreateE2SmKpmIndicationMsg(plmnID, cellIdentityValue, cellIdentityLen, dlTotalAvlblProbs,
		ulTotalAvlblProbs, fiveQi, dlPrbusage, ulPrbusage, qCi, qciDlPrbusage, qciUlPrbusage, sst,
		sd, gNbCuName, cuCpNumberActvts, ranContainer)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)

	xer, err := kpmctypes.XerEncodeE2SmKpmIndicationMessage(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("XER Encoded Ind Message: %s", string(xer))

	per, err := kpmctypes.PerEncodeE2SmKpmIndicationMessage(newE2SmKpmPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Ind Message: % x", per)

}
