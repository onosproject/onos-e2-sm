// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestE2SmKpmIndicationMessageFormat1_NewE2SmKpmIndicationMessageFormat1(t *testing.T) {

	e2SmKpmIndicationMessageFormat1 := NewE2SmKpmIndicationMessageFormat1()

	assert.Equal(t, reflect.TypeOf(E2SmKpmIndicationMessageFormat1{}), reflect.TypeOf(*e2SmKpmIndicationMessageFormat1), "E2SmKpmIndicationMessageFormat1{} types are mismatched")
}

func TestE2SmKpmIndicationMessageFormat1_AddPmContainersListItem(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer := NewPfContainer().SetOCuCp(ocucp)

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)
	pmContainersListItem := NewPmContainersList().SetPerformanceContainer(pfContainer).SetTheRancontainer(ranContainer)

	e2SmKpmIndicationMessageFormat1 := NewE2SmKpmIndicationMessageFormat1().AddPmContainersListItem(pmContainersListItem)

	assert.Assert(t, e2SmKpmIndicationMessageFormat1.PmContainers != nil)
}

func TestE2SmKpmIndicationMessageFormat1_RetrievePmContainersListItemOCuCpByCuCpName(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer := NewPfContainer().SetOCuCp(ocucp)

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)
	pmContainersListItem := NewPmContainersList().SetPerformanceContainer(pfContainer).SetTheRancontainer(ranContainer)

	e2SmKpmIndicationMessageFormat1 := NewE2SmKpmIndicationMessageFormat1().AddPmContainersListItem(pmContainersListItem)

	assert.DeepEqual(t, e2SmKpmIndicationMessageFormat1.RetrievePmContainersListItemOCuCpByCuCpName(name), pmContainersListItem)
}

func TestE2SmKpmIndicationMessageFormat1_RetrievePmContainersListItemOCuCpByNUEs(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer := NewPfContainer().SetOCuCp(ocucp)

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)
	pmContainersListItem := NewPmContainersList().SetPerformanceContainer(pfContainer).SetTheRancontainer(ranContainer)

	e2SmKpmIndicationMessageFormat1 := NewE2SmKpmIndicationMessageFormat1().AddPmContainersListItem(pmContainersListItem)

	assert.DeepEqual(t, e2SmKpmIndicationMessageFormat1.RetrievePmContainersListItemOCuCpByNUEs(nUEs), pmContainersListItem)
}

func TestE2SmKpmIndicationMessageFormat1_RetrievePmContainersListItemOCuCpByRanContainer(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer := NewPfContainer().SetOCuCp(ocucp)

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)
	pmContainersListItem := NewPmContainersList().SetPerformanceContainer(pfContainer).SetTheRancontainer(ranContainer)

	e2SmKpmIndicationMessageFormat1 := NewE2SmKpmIndicationMessageFormat1().AddPmContainersListItem(pmContainersListItem)

	assert.DeepEqual(t, e2SmKpmIndicationMessageFormat1.RetrievePmContainersListItemByRanContainer(value), pmContainersListItem)
}

func TestE2SmKpmIndicationMessageFormat1_GetE2SmKpmIndicationMessageFormat1(t *testing.T) {

	var name string = "ONF"
	var nUEs int32 = 12
	ocucp := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs))
	pfContainer := NewPfContainer().SetOCuCp(ocucp)

	value := []byte{0x22, 0x21}
	ranContainer := NewRanContainer(value)
	pmContainersListItem1 := NewPmContainersList().SetPerformanceContainer(pfContainer).SetTheRancontainer(ranContainer)

	var name2 string = "FNO"
	var nUEs2 int32 = 25
	ocucp2 := NewOCuCpPfContainer().SetGNbCuCpName(NewGnbCuCpName(name2)).
		SetCuCpResourceStatus(NewOcucpPfContainerCuCpResourceStatus001(nUEs2))
	pfContainer2 := NewPfContainer().SetOCuCp(ocucp2)

	value2 := []byte{0x22, 0x21}
	ranContainer2 := NewRanContainer(value2)
	pmContainersListItem2 := NewPmContainersList().SetPerformanceContainer(pfContainer2).SetTheRancontainer(ranContainer2)

	e2SmKpmIndicationMessageFormat1_1 := NewE2SmKpmIndicationMessageFormat1().AddPmContainersListItem(pmContainersListItem1).AddPmContainersListItem(pmContainersListItem2)
	e2SmKpmIndicationMessageFormat1_2 := e2SmKpmIndicationMessageFormat1_1.GetE2SmKpmIndicationMessageFormat1()

	assert.DeepEqual(t, e2SmKpmIndicationMessageFormat1_1, e2SmKpmIndicationMessageFormat1_2)
	assert.DeepEqual(t, e2SmKpmIndicationMessageFormat1_2.RetrievePmContainersListItemOCuCpByCuCpName(name), pmContainersListItem1)
	assert.DeepEqual(t, e2SmKpmIndicationMessageFormat1_2.RetrievePmContainersListItemByRanContainer(value2), pmContainersListItem1)
}
