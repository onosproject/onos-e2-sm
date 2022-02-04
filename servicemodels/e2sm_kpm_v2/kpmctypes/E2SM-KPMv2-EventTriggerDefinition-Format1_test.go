// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createE2SMKPMEventTriggerDefinitionFormat1() *e2sm_kpm_v2.E2SmKpmEventTriggerDefinitionFormat1 {

	var rtPeriod uint32 = 15

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmEventTriggerDefinition(rtPeriod)

	return newE2SmKpmPdu.GetEventDefinitionFormat1()
}

func Test_xerEncodeE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	etdf1 := createE2SMKPMEventTriggerDefinitionFormat1()

	xer, err := xerEncodeE2SmKpmEventTriggerDefinitionFormat1(etdf1)
	assert.NilError(t, err)
	//assert.Equal(t, 131, len(xer))
	t.Logf("E2SmKpmEventTriggerDefinitionFormat1 XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	etdf1 := createE2SMKPMEventTriggerDefinitionFormat1()

	xer, err := xerEncodeE2SmKpmEventTriggerDefinitionFormat1(etdf1)
	assert.NilError(t, err)
	//assert.Equal(t, 131, len(xer))
	t.Logf("E2SmKpmEventTriggerDefinitionFormat1 XER\n%s", string(xer))

	result, err := xerDecodeE2SmKpmEventTriggerDefinitionFormat1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmEventTriggerDefinitionFormat1 XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	etdf1 := createE2SMKPMEventTriggerDefinitionFormat1()

	per, err := perEncodeE2SmKpmEventTriggerDefinitionFormat1(etdf1)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SmKpmEventTriggerDefinitionFormat1 PER\n%v", hex.Dump(per))
}

func Test_perDecodeE2SmKpmEventTriggerDefinitionFormat1(t *testing.T) {

	etdf1 := createE2SMKPMEventTriggerDefinitionFormat1()

	per, err := perEncodeE2SmKpmEventTriggerDefinitionFormat1(etdf1)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SmKpmEventTriggerDefinitionFormat1 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2SmKpmEventTriggerDefinitionFormat1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmEventTriggerDefinitionFormat1 PER - decoded\n%s", result)
}
