// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type PmContainersList struct {
	PerformanceContainer *PfContainer
	TheRancontainer      *RanContainer
}

type PmContainersListBuilder interface {
	NewPmContainersList()
	SetPerformanceContainer(pfContainer *PfContainer)
	SetTheRancontainer(ranContainer *RanContainer)
	GetPerformanceContainer()
	GetTheRancontainer()
	GetPmContainersList()
}

func NewPmContainersList() *PmContainersList {
	return &PmContainersList{}
}

func (b *PmContainersList) SetPerformanceContainer(pfContainer *PfContainer) *PmContainersList {
	b.PerformanceContainer = pfContainer
	return b
}

func (b *PmContainersList) SetTheRancontainer(ranContainer *RanContainer) *PmContainersList {
	b.TheRancontainer = ranContainer
	return b
}

func (b *PmContainersList) GetPerformanceContainer() *PfContainer {
	return b.PerformanceContainer
}

func (b *PmContainersList) GetTheRancontainer() *RanContainer {
	return b.TheRancontainer
}

func (b *PmContainersList) GetPmContainersList() *PmContainersList {
	return b
}
