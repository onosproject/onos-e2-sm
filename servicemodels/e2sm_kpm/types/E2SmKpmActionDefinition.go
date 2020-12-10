// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type E2SmKpmActionDefinition struct {
	RicStyleType *RicStyleType
}

type E2SmKpmActionDefinitionBuilder interface {
	NewE2SmKpmActionDefinition(ricStyleType *RicStyleType)
	GetRicStyleType()
	GetE2SmKpmActionDefinition()
}

func NewE2SmKpmActionDefinition(ricStyleType *RicStyleType) *E2SmKpmActionDefinition {
	return &E2SmKpmActionDefinition{
		RicStyleType: ricStyleType,
	}
}

func (b *E2SmKpmActionDefinition) GetRicStyleType() *RicStyleType {
	return b.RicStyleType
}

func (b *E2SmKpmActionDefinition) GetE2SmKpmActionDefinition() *E2SmKpmActionDefinition {
	return b
}
