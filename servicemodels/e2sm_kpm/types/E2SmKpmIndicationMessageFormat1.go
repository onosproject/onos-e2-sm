// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import "reflect"

type E2SmKpmIndicationMessageFormat1 struct {
	PmContainers []*PmContainersList
}

type E2SmKpmIndicationMessageFormat1Builder interface {
	NewE2SmKpmIndicationMessageFormat1()
	AddPmContainersListItem(pmContainersListItem *PmContainersList)
	RetrievePmContainersListItemOCuCpByCuCpName(gnbCuCpName string)
	RetrievePmContainersListItemOCuCpByNUEs(numberOfActiveUes int32)
	RetrievePmContainersListItemOCuCpByRanContainer(ranContainer []byte)
	GetPmContainersList()
	GetE2SmKpmIndicationMessageFormat1()
}

func NewE2SmKpmIndicationMessageFormat1() *E2SmKpmIndicationMessageFormat1 {
	return &E2SmKpmIndicationMessageFormat1{}
}

func (b *E2SmKpmIndicationMessageFormat1) AddPmContainersListItem(pmContainersListItem *PmContainersList) *E2SmKpmIndicationMessageFormat1 {
	b.PmContainers = append(b.PmContainers, pmContainersListItem)
	return b
}

//TODO: Probably would be redundant once ODu and OCuUp containers would come -- somehow, could be extended with container type in the future
// Also may return list of containers which satisfies this condition
func (b *E2SmKpmIndicationMessageFormat1) RetrievePmContainersListItemOCuCpByCuCpName(gnbCuCpName string) *PmContainersList {
	for _, ocucpItem := range b.PmContainers {
		if ocucpItem.PerformanceContainer.OCuCp != nil && ocucpItem.PerformanceContainer.OCuCp.GNbCuCpName.Value == gnbCuCpName {
			return ocucpItem
		}
	}
	return nil
}

//TODO: Probably would be redundant once ODu and OCuUp containers would come -- somehow, could be extended with container type in the future
// Also may return list of containers which satisfies this condition
func (b *E2SmKpmIndicationMessageFormat1) RetrievePmContainersListItemOCuCpByNUEs(numberOfActiveUes int32) *PmContainersList {
	for _, ocucpItem := range b.PmContainers {
		if ocucpItem.PerformanceContainer.OCuCp != nil && ocucpItem.PerformanceContainer.OCuCp.CuCpResourceStatus.NumberOfActiveUes == numberOfActiveUes {
			return ocucpItem
		}
	}
	return nil
}

//TODO: May return list of all matched PmContainersList items related to the certain RanContainer
func (b *E2SmKpmIndicationMessageFormat1) RetrievePmContainersListItemByRanContainer(ranContainer []byte) *PmContainersList {
	for _, pmContainerItem := range b.PmContainers {
		if reflect.DeepEqual(pmContainerItem.TheRancontainer.Value, ranContainer) {
			return pmContainerItem
		}
	}
	return nil
}

func (b *E2SmKpmIndicationMessageFormat1) GetE2SmKpmIndicationMessageFormat1() *E2SmKpmIndicationMessageFormat1 {
	return b
}
