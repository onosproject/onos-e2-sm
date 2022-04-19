//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
)

func CreateE2SmRcRanfunctionDefinition(rfsn string, oID string, dscr string) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   rfsn,
			RanFunctionE2SmOid:     oID,
			RanFunctionDescription: dscr,
		},
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionEventTrigger(rfsn string, oID string, dscr string, et *e2smrcv1.RanfunctionDefinitionEventTrigger) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   rfsn,
			RanFunctionE2SmOid:     oID,
			RanFunctionDescription: dscr,
		},
		RanFunctionDefinitionEventTrigger: et,
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionReport(rfsn string, oID string, dscr string, report *e2smrcv1.RanfunctionDefinitionReport) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   rfsn,
			RanFunctionE2SmOid:     oID,
			RanFunctionDescription: dscr,
		},
		RanFunctionDefinitionReport: report,
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionInsert(rfsn string, oID string, dscr string, insert *e2smrcv1.RanfunctionDefinitionInsert) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   rfsn,
			RanFunctionE2SmOid:     oID,
			RanFunctionDescription: dscr,
		},
		RanFunctionDefinitionInsert: insert,
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionControl(rfsn string, oID string, dscr string, control *e2smrcv1.RanfunctionDefinitionControl) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   rfsn,
			RanFunctionE2SmOid:     oID,
			RanFunctionDescription: dscr,
		},
		RanFunctionDefinitionControl: control,
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionPolicy(rfsn string, oID string, dscr string, policy *e2smrcv1.RanfunctionDefinitionPolicy) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   rfsn,
			RanFunctionE2SmOid:     oID,
			RanFunctionDescription: dscr,
		},
		RanFunctionDefinitionPolicy: policy,
	}

	return msg, nil
}
