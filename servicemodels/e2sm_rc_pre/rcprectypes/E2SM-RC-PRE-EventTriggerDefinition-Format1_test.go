// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2SmRcPreEventTriggerDefinitionFormat1() (*e2sm_rc_pre_ies.E2SmRcPreEventTriggerDefinitionFormat1, error) {

	var rtPeriod int32 = 12 // range is from 0 to 19

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinition(rtPeriod)
	if err != nil {
		return nil, err
	}

	return newE2SmRcPrePdu.GetEventDefinitionFormat1(), nil

}

func Test_xerEncodeE2SmRcPreEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmRcPreEventTriggerDefinitionFormat1, err := createE2SmRcPreEventTriggerDefinitionFormat1()
	assert.NilError(t, err)

	xer, err := xerEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 265, len(xer))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition-Format1 XER\n%s", string(xer))
}

func Test_xerDecodeE2SmRcPreEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmRcPreEventTriggerDefinitionFormat1, err := createE2SmRcPreEventTriggerDefinitionFormat1()
	assert.NilError(t, err)

	xer, err := xerEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 265, len(xer))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition-Format1 XER\n%s", string(xer))

	result, err := xerDecodeE2SmRcPreEventTriggerDefinitionFormat1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, E2SmRcPreEventTriggerDefinitionFormat1.GetPolicyTestList(), result.GetPolicyTestList(), "Encoded and decoded E2SM-RC-PRE-EventTriggerDefinition-Format1 -- PolicyTestList values are not the same")
}

func Test_perEncodeE2SmRcPreEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmRcPreEventTriggerDefinitionFormat1, err := createE2SmRcPreEventTriggerDefinitionFormat1()
	assert.NilError(t, err)
	per, err := perEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition-Format1 PER\n%s", string(per))
}

func Test_perDecodeE2SmRcPreEventTriggerDefinitionFormat1(t *testing.T) {

	E2SmRcPreEventTriggerDefinitionFormat1, err := createE2SmRcPreEventTriggerDefinitionFormat1()
	assert.NilError(t, err)
	per, err := perEncodeE2SmRcPreEventTriggerDefinitionFormat1(E2SmRcPreEventTriggerDefinitionFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2SM-RC-PRE-EventTriggerDefinition-Format1 PER\n%s", string(per))

	result, err := perDecodeE2SmRcPreEventTriggerDefinitionFormat1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	//assert.Equal(t, E2SmRcPreEventTriggerDefinitionFormat1, result, "Encoded and decoded E2SM-RC-PRE-EventTriggerDefinition-Format1 -- PolicyTestList values are not the same")
}
