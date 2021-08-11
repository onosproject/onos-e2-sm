// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmKpmRanfunctionDescription(rfSn string, rfE2SMoid string, rfd string) *e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription {

	e2SmKpmPdu := e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription{
		RanFunctionName: &e2sm_kpm_v2_go.RanfunctionName{
			RanFunctionShortName:   rfSn,
			RanFunctionE2SmOid:     rfE2SMoid,
			RanFunctionDescription: rfd,
		},
	}

	//if err := e2SmKpmPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmRanfunctionDescription %s", err.Error())
	//}
	return &e2SmKpmPdu
}

func CreateRicKpmnodeItem(globalKpmnodeID *e2sm_kpm_v2_go.GlobalKpmnodeId) *e2sm_kpm_v2_go.RicKpmnodeItem {

	res := e2sm_kpm_v2_go.RicKpmnodeItem{
		RicKpmnodeType: globalKpmnodeID,
	}

	return &res
}

func CreateCellMeasurementObjectItem(cellObjID string, cellGlobalID *e2sm_kpm_v2_go.CellGlobalId) *e2sm_kpm_v2_go.CellMeasurementObjectItem {

	return &e2sm_kpm_v2_go.CellMeasurementObjectItem{
		CellObjectId: &e2sm_kpm_v2_go.CellObjectId{
			Value: cellObjID,
		},
		CellGlobalId: cellGlobalID,
	}
}

func CreateCellGlobalIDNRCGI(plmnID []byte, cellIDBits36 []byte) (*e2sm_kpm_v2_go.CellGlobalId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 chars")
	}

	if len(cellIDBits36) != 5 {
		return nil, fmt.Errorf("expecting 5 bits to carry NRcellIdentity, got %d", len(cellIDBits36))
	}
	if cellIDBits36[4]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", cellIDBits36[4])
	}
	bs := asn1.BitString{
		Value: cellIDBits36,
		Len:   36,
	}

	return &e2sm_kpm_v2_go.CellGlobalId{
		CellGlobalId: &e2sm_kpm_v2_go.CellGlobalId_NrCgi{
			NrCgi: &e2sm_kpm_v2_go.Nrcgi{
				PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
					Value: plmnID,
				},
				NRcellIdentity: &e2sm_kpm_v2_go.NrcellIdentity{
					Value: &bs,
				},
			},
		},
	}, nil
}

func CreateCellGlobalIDEUTRACGI(plmnID []byte, bs *asn1.BitString) (*e2sm_kpm_v2_go.CellGlobalId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 chars")
	}

	if len(bs.GetValue()) != 4 {
		return nil, fmt.Errorf("expecting 4 bits to carry EutraCellIdentity, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[3]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[3])
	}

	return &e2sm_kpm_v2_go.CellGlobalId{
		CellGlobalId: &e2sm_kpm_v2_go.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2sm_kpm_v2_go.Eutracgi{
				PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
					Value: plmnID,
				},
				EUtracellIdentity: &e2sm_kpm_v2_go.EutracellIdentity{
					Value: bs,
				},
			},
		},
	}, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) *e2sm_kpm_v2_go.RicEventTriggerStyleItem {

	return &e2sm_kpm_v2_go.RicEventTriggerStyleItem{
		RicEventTriggerStyleType: &e2sm_kpm_v2_go.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2sm_kpm_v2_go.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2sm_kpm_v2_go.RicFormatType{
			Value: ricFormatType,
		},
	}
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32,
	measInfoActionList *e2sm_kpm_v2_go.MeasurementInfoActionList, indHdrFormatType int32,
	indMsgFormatType int32) *e2sm_kpm_v2_go.RicReportStyleItem {

	return &e2sm_kpm_v2_go.RicReportStyleItem{
		RicReportStyleType: &e2sm_kpm_v2_go.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2sm_kpm_v2_go.RicStyleName{
			Value: ricStyleName,
		},
		RicActionFormatType: &e2sm_kpm_v2_go.RicFormatType{
			Value: ricFormatType,
		},
		MeasInfoActionList: measInfoActionList,
		RicIndicationHeaderFormatType: &e2sm_kpm_v2_go.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2sm_kpm_v2_go.RicFormatType{
			Value: indMsgFormatType,
		},
	}
}

func CreateMeasurementInfoActionItem(measTypeName string) *e2sm_kpm_v2_go.MeasurementInfoActionItem {

	return &e2sm_kpm_v2_go.MeasurementInfoActionItem{
		MeasName: &e2sm_kpm_v2_go.MeasurementTypeName{
			Value: measTypeName,
		},
	}
}
