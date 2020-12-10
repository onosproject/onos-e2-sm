// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	//"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	"gotest.tools/assert"
	//"io/ioutil"
	"reflect"
	"testing"
)

func TestE2SmKpmActionDefinition_NewE2SmKpmActionDefinition(t *testing.T) {

	var value int32 = 14
	e2SmKpmActionDefinition := NewE2SmKpmActionDefinition(NewRicStyleType(value))

	assert.Equal(t, reflect.TypeOf(E2SmKpmActionDefinition{}), reflect.TypeOf(*e2SmKpmActionDefinition), "E2SmKpmActionDefinition{} types are mismatched")
	assert.Equal(t, e2SmKpmActionDefinition.RicStyleType.Value, value, "E2SmKpmActionDefinition{} values are mismatched")
}

func TestE2SmKpmActionDefinition_GetRicStyleType(t *testing.T) {

	var value int32 = 14
	ricStyleType := NewRicStyleType(value)
	e2SmKpmActionDefinition := NewE2SmKpmActionDefinition(ricStyleType)

	assert.DeepEqual(t, e2SmKpmActionDefinition.GetRicStyleType(), ricStyleType)
}

func TestE2SmKpmActionDefinition_GetE2SmKpmActionDefinition(t *testing.T) {

	var value int32 = 14
	e2SmKpmActionDefinition1 := NewE2SmKpmActionDefinition(NewRicStyleType(value))
	e2SmKpmActionDefinition2 := e2SmKpmActionDefinition1.GetE2SmKpmActionDefinition()

	assert.DeepEqual(t, e2SmKpmActionDefinition1, e2SmKpmActionDefinition2)
}
