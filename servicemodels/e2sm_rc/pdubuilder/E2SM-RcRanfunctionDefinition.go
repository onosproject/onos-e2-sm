//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
)

func CreateE2SmRcRanfunctionDefinition(ranFunctionShortName string, ranFunctionOID string, ranFunctionDescription string) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,
			RanFunctionE2SmOid:     ranFunctionOID,
			RanFunctionDescription: ranFunctionDescription,
		},
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionEventTrigger(ranFunctionShortName string, ranFunctionOID string, ranFunctionDescription string, et *e2smrcv1.RanfunctionDefinitionEventTrigger) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,
			RanFunctionE2SmOid:     ranFunctionOID,
			RanFunctionDescription: ranFunctionDescription,
		},
		RanFunctionDefinitionEventTrigger: et,
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionReport(ranFunctionShortName string, ranFunctionOID string, ranFunctionDescription string, report *e2smrcv1.RanfunctionDefinitionReport) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,
			RanFunctionE2SmOid:     ranFunctionOID,
			RanFunctionDescription: ranFunctionDescription,
		},
		RanFunctionDefinitionReport: report,
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionInsert(ranFunctionShortName string, ranFunctionOID string, ranFunctionDescription string, insert *e2smrcv1.RanfunctionDefinitionInsert) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,
			RanFunctionE2SmOid:     ranFunctionOID,
			RanFunctionDescription: ranFunctionDescription,
		},
		RanFunctionDefinitionInsert: insert,
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionControl(ranFunctionShortName string, ranFunctionOID string, ranFunctionDescription string, control *e2smrcv1.RanfunctionDefinitionControl) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,
			RanFunctionE2SmOid:     ranFunctionOID,
			RanFunctionDescription: ranFunctionDescription,
		},
		RanFunctionDefinitionControl: control,
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinitionPolicy(ranFunctionShortName string, ranFunctionOID string, ranFunctionDescription string, policy *e2smrcv1.RanfunctionDefinitionPolicy) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

	msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
		RanFunctionName: &e2smcommonies.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,
			RanFunctionE2SmOid:     ranFunctionOID,
			RanFunctionDescription: ranFunctionDescription,
		},
		RanFunctionDefinitionPolicy: policy,
	}

	return msg, nil
}
