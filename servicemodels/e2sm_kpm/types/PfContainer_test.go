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

func TestPfContainer_NewPfContainer(t *testing.T) {

	pfContainer := NewPfContainer()

	assert.Equal(t, reflect.TypeOf(PfContainer{}), reflect.TypeOf(*pfContainer), "PfContainer{} types are mismatched")
}

func TestPfContainer_SetOCuCp(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	pfContainer := NewPfContainer().SetOCuCp(NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs)))

	assert.Equal(t, pfContainer.OCuCp.GetGNbCuCpName().GetValue(), name, "TestPfContainer() GNbCuCpName values mismatch")
	assert.Equal(t, pfContainer.GetOCuCp().GetCuCpResourceStatus().GetNumberOfActiveUes(), nUEs, "TestPfContainer() NUEs values mismatch")
}

func TestPfContainer_GetOCuCp(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer := NewPfContainer().SetOCuCp(ocucp)

	assert.DeepEqual(t, pfContainer.GetOCuCp(), ocucp)
}

func TestPfContainer_GetPfContainer(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer1 := NewPfContainer().SetOCuCp(ocucp)
	pfContainer2 := pfContainer1.GetPfContainer()

	assert.DeepEqual(t, pfContainer1, pfContainer2)
}
