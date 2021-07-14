// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGlobalKpmnodeGnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	var gNbCuUpID int64 = 31
	var gnbDuID int64 = 42

	gNbID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	gNbID.GetGNb().GNbCuUpId = &e2sm_kpm_v2.GnbCuUpId{
		Value: gNbCuUpID,
	}
	gNbID.GetGNb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeGnbID(gNbID.GetGNb())
	assert.NilError(t, err)
	//assert.Equal(t, 311, len(xer))
	//assert.Equal(t, 241, len(xer)) // without GNbCuUpID and GNbDuID
	t.Logf("GlobalKpmnodeGnbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalKpmnodeGnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	var gNbCuUpID int64 = 31
	var gnbDuID int64 = 42

	gNbID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	gNbID.GetGNb().GNbCuUpId = &e2sm_kpm_v2.GnbCuUpId{
		Value: gNbCuUpID,
	}
	gNbID.GetGNb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeGnbID(gNbID.GetGNb())
	assert.NilError(t, err)
	//assert.Equal(t, 311, len(xer))
	//assert.Equal(t, 241, len(xer)) // without GNbCuUpID and GNbDuID
	t.Logf("GlobalKpmnodeGnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalKpmnodeGnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeGnbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalKpmnodeGnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	var gNbCuUpID int64 = 31
	var gnbDuID int64 = 42

	gNbID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	gNbID.GetGNb().GNbCuUpId = &e2sm_kpm_v2.GnbCuUpId{
		Value: gNbCuUpID,
	}
	gNbID.GetGNb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeGnbID(gNbID.GetGNb())
	assert.NilError(t, err)
	//assert.Equal(t, 12, len(per))
	//assert.Equal(t, 8, len(per)) // without GNbCuUpID and GNbDuID
	t.Logf("GlobalKpmnodeGnbID PER\n%v", hex.Dump(per))
}

func Test_perDecodeGlobalKpmnodeGnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	var gNbCuUpID int64 = 31
	var gnbDuID int64 = 42

	gNbID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	gNbID.GetGNb().GNbCuUpId = &e2sm_kpm_v2.GnbCuUpId{
		Value: gNbCuUpID,
	}
	gNbID.GetGNb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeGnbID(gNbID.GetGNb())
	assert.NilError(t, err)
	//assert.Equal(t, 12, len(per))
	//assert.Equal(t, 8, len(per)) // without GNbCuUpID and GNbDuID
	t.Logf("GlobalKpmnodeGnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalKpmnodeGnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeGnbID PER - decoded\n%v", result)
}
