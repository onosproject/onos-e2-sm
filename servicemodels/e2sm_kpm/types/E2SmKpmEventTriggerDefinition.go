// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type E2SmKpmEventTriggerDefinition struct {
	E2SmKpmEventTriggerDefinition *E2SmKpmEventTriggerDefinitionFormat1
}

type E2SmKpmEventTriggerDefinitionBuilder interface {
	NewE2SmKpmEventTriggerDefinition()
	SetE2SmKpmEventTriggerDefinition(item *E2SmKpmEventTriggerDefinitionFormat1)
	GetE2SmKpmEventTriggerDefinition()
}

func NewE2SmKpmEventTriggerDefinition() *E2SmKpmEventTriggerDefinition {
	return &E2SmKpmEventTriggerDefinition{}
}

func (b *E2SmKpmEventTriggerDefinition) SetE2SmKpmEventTriggerDefinition(item *E2SmKpmEventTriggerDefinitionFormat1) *E2SmKpmEventTriggerDefinition {
	b.E2SmKpmEventTriggerDefinition = item
	return b
}

func (b *E2SmKpmEventTriggerDefinition) GetE2SmKpmEventTriggerDefinition() *E2SmKpmEventTriggerDefinitionFormat1 {
	return b.E2SmKpmEventTriggerDefinition
}
