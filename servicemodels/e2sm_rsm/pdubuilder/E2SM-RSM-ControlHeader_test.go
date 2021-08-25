// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	"testing"
)

func Test_E2SmRsmControlHeader(t *testing.T) {

	ch1 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandSliceCreate())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch1)

	//ToDo - embed APER validation

	ch2 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandSliceUpdate())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch2)

	//ToDo - embed APER validation

	ch3 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandSliceDelete())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch3)

	//ToDo - embed APER validation

	ch4 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandUeAssociate())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch4)

	//ToDo - embed APER validation

	ch5 := CreateE2SmRsmControlHeader(CreateE2SmRsmCommandEventTriggers())
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", ch5)

	//ToDo - embed APER validation
}
