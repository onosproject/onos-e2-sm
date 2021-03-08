// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGlobalKpmnodeEnGnbID(t *testing.T) {

	var bsValue uint64 = 0x9bcde4
	var bsLen uint32 = 32
	plmnID := []byte{0x21, 0x22, 0x23}
	var gnbDuID int64 = 32
	var gnbCuUpID int64 = 42

	enbID, err := pdubuilder.CreateGlobalKpmnodeID_enGNbID(bsValue, bsLen, plmnID, gnbCuUpID, gnbDuID)
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeEnGnbID(enbID.GetEnGNb())
	assert.NilError(t, err)
	assert.Equal(t, 335, len(xer))
	t.Logf("GlobalKpmnodeEnGnbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalKpmnodeEnGnbID(t *testing.T) {

	var bsValue uint64 = 0x9bcde4
	var bsLen uint32 = 32
	plmnID := []byte{0x21, 0x22, 0x23}
	var gnbDuID int64 = 32
	var gnbCuUpID int64 = 42

	enbID, err := pdubuilder.CreateGlobalKpmnodeID_enGNbID(bsValue, bsLen, plmnID, gnbCuUpID, gnbDuID)
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeEnGnbID(enbID.GetEnGNb())
	assert.NilError(t, err)
	assert.Equal(t, 335, len(xer))
	t.Logf("GlobalKpmnodeEnGnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalKpmnodeEnGnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeEnGnbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalKpmnodeEnGnbID(t *testing.T) {

	var bsValue uint64 = 0x9bcde4
	var bsLen uint32 = 32
	plmnID := []byte{0x21, 0x22, 0x23}
	var gnbDuID int64 = 32
	var gnbCuUpID int64 = 42

	enbID, err := pdubuilder.CreateGlobalKpmnodeID_enGNbID(bsValue, bsLen, plmnID, gnbCuUpID, gnbDuID)
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeEnGnbID(enbID.GetEnGNb())
	assert.NilError(t, err)
	assert.Equal(t, 13, len(per))
	t.Logf("GlobalKpmnodeEnGnbID PER\n%s", string(per))
}

func Test_perDecodeGlobalKpmnodeEnGnbID(t *testing.T) {

	var bsValue uint64 = 0x9bcde4
	var bsLen uint32 = 32
	plmnID := []byte{0x21, 0x22, 0x23}
	var gnbDuID int64 = 32
	var gnbCuUpID int64 = 42

	enbID, err := pdubuilder.CreateGlobalKpmnodeID_enGNbID(bsValue, bsLen, plmnID, gnbCuUpID, gnbDuID)
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeEnGnbID(enbID.GetEnGNb())
	assert.NilError(t, err)
	assert.Equal(t, 13, len(per))
	t.Logf("GlobalKpmnodeEnGnbID PER\n%s", string(per))

	result, err := perDecodeGlobalKpmnodeEnGnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeEnGnbID PER - decoded\n%v", result)
}
