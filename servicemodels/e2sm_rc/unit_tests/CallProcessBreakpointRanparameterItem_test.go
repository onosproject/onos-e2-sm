//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package unit_tests

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/choiceOptions"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	"testing"
)

func Test_CallProcessBreakpointRanparameterItem(t *testing.T) {

	structItem := &e2smrcv1.RanparameterDefinitionChoiceStructureItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: 43,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: "item",
		},
	}

	structList := make([]*e2smrcv1.RanparameterDefinitionChoiceStructureItem, 0)
	structList = append(structList, structItem)
	structList = append(structList, structItem)

	msg := &e2smrcv1.CallProcessBreakpointRanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: 32,
		},
		RanParameterName: &e2smrcv1.RanparameterName{
			Value: "Radisys",
		},
		RanParameterDefinition: &e2smrcv1.RanparameterDefinition{
			RanParameterDefinitionChoice: &e2smrcv1.RanparameterDefinitionChoice{
				RanparameterDefinitionChoice: &e2smrcv1.RanparameterDefinitionChoice_ChoiceStructure{
					ChoiceStructure: &e2smrcv1.RanparameterDefinitionChoiceStructure{
						RanParameterStructure: structList,
					},
				},
			},
		},
	}

	per, err := aper.MarshalWithParams(msg, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	assert.NilError(t, err)
	t.Logf("Encoded E2SM-RcActionDefinition PER bytes are\n%v", hex.Dump(per))

	result := &e2smrcv1.CallProcessBreakpointRanparameterItem{}
	err = aper.UnmarshalWithParams(per, result, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-RcActionDefinition from PER is\n%v", result)

	assert.Equal(t, msg.String(), result.String())
}
