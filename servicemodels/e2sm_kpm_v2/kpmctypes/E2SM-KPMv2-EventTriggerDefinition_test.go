// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createE2SMKPMEventTriggerDefinition() *e2sm_kpm_v2.E2SmKpmEventTriggerDefinition {

	var rtPeriod uint32 = 15

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmEventTriggerDefinition(rtPeriod)

	return newE2SmKpmPdu
}

//func Test_xerEncodeE2SmKpmEventTriggerDefinition(t *testing.T) {
//
//	etd := createE2SMKPMEventTriggerDefinition()
//
//	xer, err := XerEncodeE2SmKpmEventTriggerDefinition(etd)
//	assert.NilError(t, err)
//	assert.Equal(t, 263, len(xer))
//	t.Logf("E2SmKpmEventTriggerDefinition XER\n%s", string(xer))
//}
//
//func Test_xerDecodeE2SmKpmEventTriggerDefinition(t *testing.T) {
//
//	etd := createE2SMKPMEventTriggerDefinition()
//
//	xer, err := XerEncodeE2SmKpmEventTriggerDefinition(etd)
//	assert.NilError(t, err)
//	assert.Equal(t, 263, len(xer))
//	t.Logf("E2SmKpmEventTriggerDefinition XER\n%s", string(xer))
//
//	result, err := XerDecodeE2SmKpmEventTriggerDefinition(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2SmKpmEventTriggerDefinition XER - decoded\n%s", result)
//}

func Test_perEncodeE2SmKpmEventTriggerDefinition(t *testing.T) {

	etd := createE2SMKPMEventTriggerDefinition()

	per, err := PerEncodeE2SmKpmEventTriggerDefinition(etd)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SmKpmEventTriggerDefinition PER\n%s", string(per))
}

func Test_perDecodeE2SmKpmEventTriggerDefinition(t *testing.T) {

	etd := createE2SMKPMEventTriggerDefinition()

	per, err := PerEncodeE2SmKpmEventTriggerDefinition(etd)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SmKpmEventTriggerDefinition PER\n%s", string(per))

	result, err := PerDecodeE2SmKpmEventTriggerDefinition(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmEventTriggerDefinition PER - decoded\n%s", result)
}
