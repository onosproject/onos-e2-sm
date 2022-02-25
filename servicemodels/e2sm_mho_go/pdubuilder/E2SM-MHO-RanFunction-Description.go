// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smmhov2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2smv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMhoRanfunctionDescriptionMsg(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string) (*e2smmhov2.E2SmMhoRanfunctionDescription, error) {

	e2smMhoPdu := e2smmhov2.E2SmMhoRanfunctionDescription{
		RanFunctionName: &e2smv2.RanfunctionName{
			RanFunctionShortName:   ranFunctionShortName,   //string
			RanFunctionE2SmOid:     ranFunctionE2SmOid,     //sting
			RanFunctionDescription: ranFunctionDescription, //string
		},
		E2SmMhoRanfunctionItem: &e2smmhov2.E2SmMhoRanfunctionDescription_E2SmMhoRanfunctionItem001{},
	}

	if err := e2smMhoPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMhoRanfunctionDescriptionMsg(): error validating E2SmPDU %s", err.Error())
	}
	return &e2smMhoPdu, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) (*e2smmhov2.RicEventTriggerStyleList, error) {

	res := &e2smmhov2.RicEventTriggerStyleList{
		RicEventTriggerStyleType: &e2smv2.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2smv2.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2smv2.RicFormatType{
			Value: ricFormatType,
		},
	}

	if err := res.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicEventTriggerStyleItem(): error validationg E2SmPDU %s", err)
	}

	return res, nil
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, indHdrFormatType int32, indMsgFormatType int32) (*e2smmhov2.RicReportStyleList, error) {

	res := &e2smmhov2.RicReportStyleList{
		RicReportStyleType: &e2smv2.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2smv2.RicStyleName{
			Value: ricStyleName,
		},
		RicIndicationHeaderFormatType: &e2smv2.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2smv2.RicFormatType{
			Value: indMsgFormatType,
		},
	}

	if err := res.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicReportStyleItem(): error validationg E2SmPDU %s", err)
	}

	return res, nil
}
