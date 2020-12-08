// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type NRcgi struct {
	PlmnID   *PlmnIdentity
	NRcellID *NrcellIdentity
}

type NRcgiBuilder interface {
	NewNRcgi()
	SetPlmnID(plmnID PlmnIdentity)
	SetNRcellID(nrCellID NrcellIdentity)
	GetPlmnID()
	GetNRcellID()
	GetNRcgi()
}

func NewNRcgi() *NRcgi {
	return &NRcgi{}
}

func (b *NRcgi) SetPlmnID(plmnID *PlmnIdentity) *NRcgi {
	b.PlmnID = plmnID
	return b
}

func (b *NRcgi) SetNRcellID(nrCellID *NrcellIdentity) *NRcgi {
	b.NRcellID = nrCellID
	return b
}

func (b *NRcgi) GetPlmnID() *PlmnIdentity {
	return b.PlmnID
}

func (b *NRcgi) GetNRcellID() *NrcellIdentity {
	return b.NRcellID
}

func (b *NRcgi) GetNRcgi() *NRcgi {
	return b
}
