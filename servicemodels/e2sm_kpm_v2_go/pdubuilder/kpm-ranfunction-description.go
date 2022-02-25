// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmKpmRanfunctionDescription(rfSn string, rfE2SMoid string, rfd string) (*e2smkpmv2.E2SmKpmRanfunctionDescription, error) {

	e2SmKpmPdu := e2smkpmv2.E2SmKpmRanfunctionDescription{
		RanFunctionName: &e2smkpmv2.RanfunctionName{
			RanFunctionShortName:   rfSn,
			RanFunctionE2SmOid:     rfE2SMoid,
			RanFunctionDescription: rfd,
		},
	}

	if err := e2SmKpmPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmKpmRanfunctionDescription(): error validating E2SmKpmRanfunctionDescription %s", err.Error())
	}
	return &e2SmKpmPdu, nil
}

func CreateRicKpmnodeItem(globalKpmnodeID *e2smkpmv2.GlobalKpmnodeId) *e2smkpmv2.RicKpmnodeItem {

	res := e2smkpmv2.RicKpmnodeItem{
		RicKpmnodeType: globalKpmnodeID,
	}

	return &res
}

func CreateCellMeasurementObjectItem(cellObjID string, cellGlobalID *e2smkpmv2.CellGlobalId) *e2smkpmv2.CellMeasurementObjectItem {

	return &e2smkpmv2.CellMeasurementObjectItem{
		CellObjectId: &e2smkpmv2.CellObjectId{
			Value: cellObjID,
		},
		CellGlobalId: cellGlobalID,
	}
}

func CreateCellGlobalIDNRCGI(plmnID []byte, cellIDBits36 []byte) (*e2smkpmv2.CellGlobalId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	if len(cellIDBits36) != 5 {
		return nil, errors.NewInvalid("expecting 5 bits to carry NRcellIdentity, got %d", len(cellIDBits36))
	}
	if cellIDBits36[4]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", cellIDBits36[4])
	}
	bs := asn1.BitString{
		Value: cellIDBits36,
		Len:   36,
	}

	return &e2smkpmv2.CellGlobalId{
		CellGlobalId: &e2smkpmv2.CellGlobalId_NrCgi{
			NrCgi: &e2smkpmv2.Nrcgi{
				PLmnIdentity: &e2smkpmv2.PlmnIdentity{
					Value: plmnID,
				},
				NRcellIdentity: &e2smkpmv2.NrcellIdentity{
					Value: &bs,
				},
			},
		},
	}, nil
}

func CreateCellGlobalIDEUTRACGI(plmnID []byte, bs *asn1.BitString) (*e2smkpmv2.CellGlobalId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	if len(bs.GetValue()) != 4 {
		return nil, errors.NewInvalid("expecting 4 bits to carry EutraCellIdentity, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[3]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[3])
	}

	return &e2smkpmv2.CellGlobalId{
		CellGlobalId: &e2smkpmv2.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2smkpmv2.Eutracgi{
				PLmnIdentity: &e2smkpmv2.PlmnIdentity{
					Value: plmnID,
				},
				EUtracellIdentity: &e2smkpmv2.EutracellIdentity{
					Value: bs,
				},
			},
		},
	}, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) *e2smkpmv2.RicEventTriggerStyleItem {

	return &e2smkpmv2.RicEventTriggerStyleItem{
		RicEventTriggerStyleType: &e2smkpmv2.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2smkpmv2.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2smkpmv2.RicFormatType{
			Value: ricFormatType,
		},
	}
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32,
	measInfoActionList *e2smkpmv2.MeasurementInfoActionList, indHdrFormatType int32,
	indMsgFormatType int32) *e2smkpmv2.RicReportStyleItem {

	return &e2smkpmv2.RicReportStyleItem{
		RicReportStyleType: &e2smkpmv2.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2smkpmv2.RicStyleName{
			Value: ricStyleName,
		},
		RicActionFormatType: &e2smkpmv2.RicFormatType{
			Value: ricFormatType,
		},
		MeasInfoActionList: measInfoActionList,
		RicIndicationHeaderFormatType: &e2smkpmv2.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2smkpmv2.RicFormatType{
			Value: indMsgFormatType,
		},
	}
}

func CreateMeasurementInfoActionItem(measTypeName string) *e2smkpmv2.MeasurementInfoActionItem {

	return &e2smkpmv2.MeasurementInfoActionItem{
		MeasName: &e2smkpmv2.MeasurementTypeName{
			Value: measTypeName,
		},
	}
}
