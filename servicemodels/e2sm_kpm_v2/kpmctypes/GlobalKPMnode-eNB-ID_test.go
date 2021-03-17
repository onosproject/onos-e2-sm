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

func Test_xerEncodeGlobalKpmnodeEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   28,
	}
	plmnID := []byte{0x21, 0x22, 0x23}

	enbID, err := pdubuilder.CreateGlobalKpmnodeIDeNBID(&bs, plmnID)
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeEnbID(enbID.GetENb())
	assert.NilError(t, err)
	assert.Equal(t, 269, len(xer))
	t.Logf("GlobalKpmnodeEnbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalKpmnodeEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   28,
	}
	plmnID := []byte{0x21, 0x22, 0x23}

	enbID, err := pdubuilder.CreateGlobalKpmnodeIDeNBID(&bs, plmnID)
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeEnbID(enbID.GetENb())
	assert.NilError(t, err)
	assert.Equal(t, 269, len(xer))
	t.Logf("GlobalKpmnodeEnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalKpmnodeEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeEnbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalKpmnodeEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   28,
	}
	plmnID := []byte{0x21, 0x22, 0x23}

	enbID, err := pdubuilder.CreateGlobalKpmnodeIDeNBID(&bs, plmnID)
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeEnbID(enbID.GetENb())
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per))
	t.Logf("GlobalKpmnodeEnbID PER\n%s", string(per))
}

func Test_perDecodeGlobalKpmnodeEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcde4,
		Len:   28,
	}
	plmnID := []byte{0x21, 0x22, 0x23}

	enbID, err := pdubuilder.CreateGlobalKpmnodeIDeNBID(&bs, plmnID)
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeEnbID(enbID.GetENb())
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per))
	t.Logf("GlobalKpmnodeEnbID PER\n%s", string(per))

	result, err := perDecodeGlobalKpmnodeEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeEnbID PER - decoded\n%s", result)
}
