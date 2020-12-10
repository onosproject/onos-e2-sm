// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestRanfunctionName_NewRanfunctionName(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName := NewRanfunctionName(shortName, oid, description, instance)

	assert.Equal(t, reflect.TypeOf(RanfunctionName{}), reflect.TypeOf(*ranfunctionName), "RanfunctionName{} types are mismatched")
}

func TestRanfunctionName_GetRanFunctionShortName(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName := NewRanfunctionName(shortName, oid, description, instance)

	assert.Equal(t, ranfunctionName.GetRanFunctionShortName(), shortName, "Test_RanfunctionName ShortName mismatch")
}

func TestRanfunctionName_GetRanFunctionE2SmOid(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName := NewRanfunctionName(shortName, oid, description, instance)

	assert.Equal(t, ranfunctionName.GetRanFunctionE2SmOid(), oid, "Test_RanfunctionName Oid mismatch")
}

func TestRanfunctionName_GetRanFunctionDescription(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName := NewRanfunctionName(shortName, oid, description, instance)

	assert.Equal(t, ranfunctionName.GetRanFunctionDescription(), description, "Test_RanfunctionName Description mismatch")
}

func TestRanfunctionName_GetRanFunctionInstance(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName := NewRanfunctionName(shortName, oid, description, instance)

	assert.Equal(t, ranfunctionName.GetRanFunctionInstance(), instance, "Test_RanfunctionName Instance mismatch")
}

func TestRanfunctionName_GetRanfunctionName(t *testing.T) {

	var shortName string = "MenloPark"
	var oid string = "ONF"
	var description string = "OpenNetworkingFoundation"
	var instance int32 = 1
	ranfunctionName1 := NewRanfunctionName(shortName, oid, description, instance)
	ranfunctionName2 := ranfunctionName1.GetRanfunctionName()

	assert.DeepEqual(t, ranfunctionName1, ranfunctionName2)
}
