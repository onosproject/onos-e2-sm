// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGlobalKpmnodeGnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len: 22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	var gNbCuUp int64 = 31
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeID_gNBID(bs, plmnID, gNbCuUp, gnbDuID)
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeGnbID(ngeNbID.GetGNb())
	assert.NilError(t, err)
	assert.Equal(t, 307, len(xer))
	t.Logf("GlobalKpmnodeGnbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalKpmnodeGnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len: 22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	var gNbCuUp int64 = 31
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeID_gNBID(bs, plmnID, gNbCuUp, gnbDuID)
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeGnbID(ngeNbID.GetGNb())
	assert.NilError(t, err)
	assert.Equal(t, 307, len(xer))
	t.Logf("GlobalKpmnodeGnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalKpmnodeGnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeGnbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalKpmnodeGnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len: 22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	var gNbCuUp int64 = 31
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeID_gNBID(bs, plmnID, gNbCuUp, gnbDuID)
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeGnbID(ngeNbID.GetGNb())
	assert.NilError(t, err)
	assert.Equal(t, 12, len(per))
	t.Logf("GlobalKpmnodeGnbID PER\n%s", string(per))
}

func Test_perDecodeGlobalKpmnodeGnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len: 22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	var gNbCuUp int64 = 31
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeID_gNBID(bs, plmnID, gNbCuUp, gnbDuID)
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeGnbID(ngeNbID.GetGNb())
	assert.NilError(t, err)
	assert.Equal(t, 12, len(per))
	t.Logf("GlobalKpmnodeGnbID PER\n%s", string(per))

	result, err := perDecodeGlobalKpmnodeGnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeGnbID PER - decoded\n%v", result)
}
