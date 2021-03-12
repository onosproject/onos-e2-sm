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

func Test_xerEncodeGlobalKpmnodeNgEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   18,
	}
	longMacroEnbID := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   21,
	}
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(&bs, plmnID, &shortMacroEnbID, &longMacroEnbID, gnbDuID)
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeNgEnbID(ngeNbID.GetNgENb())
	assert.NilError(t, err)
	assert.Equal(t, 474, len(xer))
	t.Logf("GlobalKpmnodeNgEnbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalKpmnodeNgEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   18,
	}
	longMacroEnbID := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   21,
	}
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(&bs, plmnID, &shortMacroEnbID, &longMacroEnbID, gnbDuID)
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeNgEnbID(ngeNbID.GetNgENb())
	assert.NilError(t, err)
	assert.Equal(t, 474, len(xer))
	t.Logf("GlobalKpmnodeNgEnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalKpmnodeNgEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeNgEnbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalKpmnodeNgEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   18,
	}
	longMacroEnbID := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   21,
	}
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(&bs, plmnID, &shortMacroEnbID, &longMacroEnbID, gnbDuID)
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeNgEnbID(ngeNbID.GetNgENb())
	assert.NilError(t, err)
	assert.Equal(t, 15, len(per))
	t.Logf("GlobalKpmnodeNgEnbID PER\n%s", string(per))
}

func Test_perDecodeGlobalKpmnodeNgEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   18,
	}
	longMacroEnbID := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   21,
	}
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(&bs, plmnID, &shortMacroEnbID, &longMacroEnbID, gnbDuID)
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeNgEnbID(ngeNbID.GetNgENb())
	assert.NilError(t, err)
	assert.Equal(t, 15, len(per))
	t.Logf("GlobalKpmnodeNgEnbID PER\n%s", string(per))

	result, err := perDecodeGlobalKpmnodeNgEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeNgEnbID PER - decode\n%v", result)
}
