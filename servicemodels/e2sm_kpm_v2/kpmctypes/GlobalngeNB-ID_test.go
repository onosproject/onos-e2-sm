// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGlobalngeNbID(t *testing.T) {

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

	xer, err := xerEncodeGlobalngeNbID(ngeNbID.GetNgENb().GetGlobalNgENbId())
	assert.NilError(t, err)
	assert.Equal(t, 331, len(xer))
	t.Logf("GlobalngeNbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalngeNbID(t *testing.T) {

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

	xer, err := xerEncodeGlobalngeNbID(ngeNbID.GetNgENb().GetGlobalNgENbId())
	assert.NilError(t, err)
	assert.Equal(t, 331, len(xer))
	t.Logf("GlobalngeNbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalngeNbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalngeNbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalngeNbID(t *testing.T) {

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

	per, err := perEncodeGlobalngeNbID(ngeNbID.GetNgENb().GetGlobalNgENbId())
	assert.NilError(t, err)
	assert.Equal(t, 14, len(per))
	t.Logf("GlobalngeNbID PER\n%s", string(per))
}

func Test_perDecodeGlobalngeNbID(t *testing.T) {

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

	per, err := perEncodeGlobalngeNbID(ngeNbID.GetNgENb().GetGlobalNgENbId())
	assert.NilError(t, err)
	assert.Equal(t, 14, len(per))
	t.Logf("GlobalngeNbID PER\n%s", string(per))

	result, err := perDecodeGlobalngeNbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalngeNbID PER - decode\n%v", result)
}
