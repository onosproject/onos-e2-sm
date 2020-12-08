// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestPmContainersList_NewPmContainersList(t *testing.T) {

	pmContainersList := NewPmContainersList()

	assert.Equal(t, reflect.TypeOf(PmContainersList{}), reflect.TypeOf(*pmContainersList), "PmContainersList{} types are mismatched")
}

func TestPmContainersList_SetPerformanceContainer(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer := NewPfContainer().SetOCuCp(ocucp)
	pmContainersList := NewPmContainersList().SetPerformanceContainer(pfContainer)

	assert.DeepEqual(t, pmContainersList.PerformanceContainer, pfContainer)
}

func TestPmContainersList_SetTheRancontainer(t *testing.T) {

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)
	pmContainersList := NewPmContainersList().SetTheRancontainer(ranContainer)

	assert.DeepEqual(t, pmContainersList.TheRancontainer, ranContainer)
}

func TestPmContainersList_GetPerformanceContainer(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer := NewPfContainer().SetOCuCp(ocucp)
	pmContainersList := NewPmContainersList().SetPerformanceContainer(pfContainer)

	assert.DeepEqual(t, pmContainersList.GetPerformanceContainer(), pfContainer)
}

func TestPmContainersList_GetTheRancontainer(t *testing.T) {

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)
	pmContainersList := NewPmContainersList().SetTheRancontainer(ranContainer)

	assert.DeepEqual(t, pmContainersList.GetTheRancontainer(), ranContainer)
}

func TestPmContainersList_GetPmContainersList(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer := NewPfContainer().SetOCuCp(ocucp)

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)
	pmContainersList1 := NewPmContainersList().SetPerformanceContainer(pfContainer).SetTheRancontainer(ranContainer)
	pmContainersList2 := pmContainersList1.GetPmContainersList()

	assert.DeepEqual(t, pmContainersList1, pmContainersList2)
}
