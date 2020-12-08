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

func TestE2SmKpmIndicationMessage_NewE2SmKpmIndicationMessage(t *testing.T) {

	e2SmKpmIndicationMessage := NewE2SmKpmIndicationMessage()

	assert.Equal(t, reflect.TypeOf(E2SmKpmIndicationMessage{}), reflect.TypeOf(*e2SmKpmIndicationMessage), "E2SmKpmIndicationMessage{} types are mismatched")
}

func TestE2SmKpmIndicationMessage_SetE2SmKpmIndicationMessageFormat1(t *testing.T) {

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

	e2SmKpmIndicationMessageFormat1 := NewE2SmKpmIndicationMessageFormat1().AddPmContainersListItem(pmContainersListItem1).AddPmContainersListItem(pmContainersListItem2)
	e2SmKpmIndicationMessage := NewE2SmKpmIndicationMessage().SetE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1)

	assert.DeepEqual(t, e2SmKpmIndicationMessage.IndicationMessageFormat1, e2SmKpmIndicationMessageFormat1)
}

func TestE2SmKpmIndicationMessage_GetE2SmKpmIndicationMessageFormat1(t *testing.T) {

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

	e2SmKpmIndicationMessageFormat1 := NewE2SmKpmIndicationMessageFormat1().AddPmContainersListItem(pmContainersListItem1).AddPmContainersListItem(pmContainersListItem2)
	e2SmKpmIndicationMessage := NewE2SmKpmIndicationMessage().SetE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1)

	assert.DeepEqual(t, e2SmKpmIndicationMessage.GetE2SmKpmIndicationMessageFormat1(), e2SmKpmIndicationMessageFormat1)
}

func TestE2SmKpmIndicationMessage_GetE2SmKpmIndicationMessage(t *testing.T) {

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

	e2SmKpmIndicationMessageFormat1 := NewE2SmKpmIndicationMessageFormat1().AddPmContainersListItem(pmContainersListItem1).AddPmContainersListItem(pmContainersListItem2)
	e2SmKpmIndicationMessage1 := NewE2SmKpmIndicationMessage().SetE2SmKpmIndicationMessageFormat1(e2SmKpmIndicationMessageFormat1)
	e2SmKpmIndicationMessage2 := e2SmKpmIndicationMessage1.GetE2SmKpmIndicationMessage()

	assert.DeepEqual(t, e2SmKpmIndicationMessage1, e2SmKpmIndicationMessage2)
}
