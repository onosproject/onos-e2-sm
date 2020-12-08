// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestOcucpPfContainer_NewOcucpPfContainer(t *testing.T) {

	ocucppf := NewOCuCpPfContainer()

	assert.Equal(t, reflect.TypeOf(OCuCpPfContainer{}), reflect.TypeOf(*ocucppf), "OCuCpPfContainer{} types are mismatched")
}

func TestOcucpPfContainer_SetGNbCuCpName(t *testing.T) {

	var name string = "ONF"
	ocucppf := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name))

	assert.Equal(t, ocucppf.GNbCuCpName.GetValue(), name, "Test_OcucpPfContainer SetGnbCuCpName values mismatch")
}

func TestOcucpPfContainer_SetCuCpResourceStatus(t *testing.T) {

	var nUEs int32 = 12
	ocucppf := NewOCuCpPfContainer().SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))

	assert.Equal(t, ocucppf.CuCpResourceStatus.GetNumberOfActiveUes(), nUEs, "Test_OcucpPfContainer SetCuCpResourceStatus values mismatch")
}

func TestOcucpPfContainer_GetGNbCuCpName(t *testing.T) {

	var name string = "ONF"
	ocucppf := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name))

	assert.Equal(t, ocucppf.GetGNbCuCpName().GetValue(), name, "Test_OcucpPfContainer GetGnbCuCpName values mismatch")
}

func TestOcucpPfContainer_GetCuCpResourceStatus(t *testing.T) {

	var nUEs int32 = 12
	ocucppf := NewOCuCpPfContainer().SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))

	assert.Equal(t, ocucppf.GetCuCpResourceStatus().GetNumberOfActiveUes(), nUEs, "Test_OcucpPfContainer GetCuCpResourceStatus values mismatch")
}

func TestOcucpPfContainer_GetOcucpPfContainer(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucppf1 := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	ocucppf2 := ocucppf1.GetOCuCpPfContainer()

	assert.DeepEqual(t, ocucppf1, ocucppf2)
}
