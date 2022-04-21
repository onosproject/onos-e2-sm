//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_E2SmRcRanfunctionDefinition(t *testing.T) {

	ranFunctionShortName := "E2SM-RC"
	ranFunctionOID := "1.3.6.1.4.1.53148.1.1.2.3"
	ranFunctionDescription1 := "empty E2SM-RC"

	msg1, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription1)
	assert.NilError(t, err)
	assert.Assert(t, msg1 != nil)

	ranFunctionDescription2 := "E2SM-RC EventTrigger"
	eventTriggerList := &e2smrcv1.RanfunctionDefinitionEventTrigger{
		RicEventTriggerStyleList: make([]*e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem, 0),
		// other parameters are optional
	}

	msg2, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription2)
	assert.NilError(t, err)
	assert.Assert(t, msg2 != nil)
	msg2.SetRanFunctionDefinitionEventTrigger(eventTriggerList)
}
