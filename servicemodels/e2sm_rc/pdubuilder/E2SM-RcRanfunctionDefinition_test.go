//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
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
	msg1.GetRanFunctionName().SetRanFunctionInstance(1)

	aper1, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg1)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper1))

	result1, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper1)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result1)
	assert.Equal(t, msg1.String(), result1.String())

	ranFunctionDescription2 := "E2SM-RC EventTrigger"
	eventTriggerList := &e2smrcv1.RanfunctionDefinitionEventTrigger{
		RicEventTriggerStyleList: make([]*e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem, 0),
		// other parameters are optional
		RanL2ParametersList: make([]*e2smrcv1.L2ParametersRanparameterItem, 0),
	}

	ricEventTriggerStyleItem, err := CreateRanfunctionDefinitionEventTriggerStyleItem(1, "SomeName", 1)
	assert.NilError(t, err)
	eventTriggerList.RicEventTriggerStyleList = append(eventTriggerList.RicEventTriggerStyleList, ricEventTriggerStyleItem)

	// adding optional item as well
	l2RanParameterItem, err := CreateL2ParametersRanparameterItem(1, "L2 parameter")
	assert.NilError(t, err)
	eventTriggerList.RanL2ParametersList = append(eventTriggerList.RanL2ParametersList, l2RanParameterItem)

	msg2, err := CreateE2SmRcRanfunctionDefinition(ranFunctionShortName, ranFunctionOID, ranFunctionDescription2)
	assert.NilError(t, err)
	assert.Assert(t, msg2 != nil)
	msg2.SetRanFunctionDefinitionEventTrigger(eventTriggerList).GetRanFunctionName().SetRanFunctionInstance(2)

	aper2, err := encoder.PerEncodeE2SmRcRanfunctionDefinition(msg2)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper2))

	result2, err := encoder.PerDecodeE2SmRcRanfunctionDefinition(aper2)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result2)
	assert.Equal(t, msg2.String(), result2.String())
}
