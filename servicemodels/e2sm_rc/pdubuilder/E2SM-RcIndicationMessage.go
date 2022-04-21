//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmRcIndicationMessageFormat1(ranPReportedList []*e2smrcv1.E2SmRcIndicationMessageFormat1Item) (*e2smrcv1.E2SmRcIndicationMessage, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessage{
		RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats{
			RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat1{
				IndicationMessageFormat1: &e2smrcv1.E2SmRcIndicationMessageFormat1{
					RanPReportedList: ranPReportedList,
				},
			},
		},
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat2(ueParameterList []*e2smrcv1.E2SmRcIndicationMessageFormat2Item) (*e2smrcv1.E2SmRcIndicationMessage, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessage{
		RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats{
			RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat2{
				IndicationMessageFormat2: &e2smrcv1.E2SmRcIndicationMessageFormat2{
					UeParameterList: ueParameterList,
				},
			},
		},
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat3(cellInfoList []*e2smrcv1.E2SmRcIndicationMessageFormat3Item) (*e2smrcv1.E2SmRcIndicationMessage, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessage{
		RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats{
			RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat3{
				IndicationMessageFormat3: &e2smrcv1.E2SmRcIndicationMessageFormat3{
					CellInfoList: cellInfoList,
				},
			},
		},
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat4(ueInfoList []*e2smrcv1.E2SmRcIndicationMessageFormat4ItemUe, cellInfoList []*e2smrcv1.E2SmRcIndicationMessageFormat4ItemCell) (*e2smrcv1.E2SmRcIndicationMessage, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessage{
		RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats{
			RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat4{
				IndicationMessageFormat4: &e2smrcv1.E2SmRcIndicationMessageFormat4{
					UeInfoList:   ueInfoList,
					CellInfoList: cellInfoList,
				},
			},
		},
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat5(ranPRequestedList []*e2smrcv1.E2SmRcIndicationMessageFormat5Item) (*e2smrcv1.E2SmRcIndicationMessage, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessage{
		RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats{
			RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat5{
				IndicationMessageFormat5: &e2smrcv1.E2SmRcIndicationMessageFormat5{
					RanPRequestedList: ranPRequestedList,
				},
			},
		},
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat1Item(ranParameterID int64, ranParameterValueType *e2smrcv1.RanparameterValueType) (*e2smrcv1.E2SmRcIndicationMessageFormat1Item, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat1Item{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterValueType: ranParameterValueType,
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat2Item(ueID *e2smcommonies.Ueid, ranPList []*e2smrcv1.E2SmRcIndicationMessageFormat2RanparameterItem) (*e2smrcv1.E2SmRcIndicationMessageFormat2Item, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat2Item{
		UeId:     ueID,
		RanPList: ranPList,
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat2RanparameterItem(ranParameterID int64, ranParameterValueType *e2smrcv1.RanparameterValueType) (*e2smrcv1.E2SmRcIndicationMessageFormat2RanparameterItem, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat2RanparameterItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterValueType: ranParameterValueType,
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat3Item(cgi *e2smcommonies.Cgi) (*e2smrcv1.E2SmRcIndicationMessageFormat3Item, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat3Item{
		CellGlobalId: cgi,
	}

	return msg, nil
}

func CreateNeighborRelationInfo(servingCellPci *e2smcommonies.ServingCellPci, servingCellArfcn *e2smcommonies.ServingCellArfcn,
	neighborCellList *e2smrcv1.NeighborCellList) (*e2smrcv1.NeighborRelationInfo, error) {

	msg := &e2smrcv1.NeighborRelationInfo{
		ServingCellPci:   servingCellPci,
		ServingCellArfcn: servingCellArfcn,
		NeighborCellList: neighborCellList,
	}

	return msg, nil
}

func CreateServingCellPciNR(nrPci int32) (*e2smcommonies.ServingCellPci, error) {

	return &e2smcommonies.ServingCellPci{
		ServingCellPci: &e2smcommonies.ServingCellPci_NR{
			NR: &e2smcommonies.NrPci{
				Value: nrPci,
			},
		},
	}, nil
}

func CreateServingCellPciEUtraPci(eutraPci int32) (*e2smcommonies.ServingCellPci, error) {

	return &e2smcommonies.ServingCellPci{
		ServingCellPci: &e2smcommonies.ServingCellPci_EUtra{
			EUtra: &e2smcommonies.EUtraPci{
				Value: eutraPci,
			},
		},
	}, nil
}

func CreateServingCellArfcnNR(nrArfcn int32) (*e2smcommonies.ServingCellArfcn, error) {

	return &e2smcommonies.ServingCellArfcn{
		ServingCellArfcn: &e2smcommonies.ServingCellArfcn_NR{
			NR: &e2smcommonies.NrArfcn{
				NRarfcn: nrArfcn,
			},
		},
	}, nil
}

func CreateServingCellArfcnEUtraPci(eutraArfcn int32) (*e2smcommonies.ServingCellArfcn, error) {

	return &e2smcommonies.ServingCellArfcn{
		ServingCellArfcn: &e2smcommonies.ServingCellArfcn_EUtra{
			EUtra: &e2smcommonies.EUtraArfcn{
				Value: eutraArfcn,
			},
		},
	}, nil
}

func CreateNeighborCellItemRanTypeChoiceNr(nrCgi *e2smcommonies.NrCgi, nrPci *e2smcommonies.NrPci, fiveGsTac *e2smcommonies.FiveGsTac,
	nrModeInfo e2smrcv1.NRModeInfo, nrFreqInfo *e2smcommonies.NrfrequencyInfo, x2XnEstablished e2smrcv1.X2XNEstablished,
	hOValidated e2smrcv1.HOValidated, version int32) (*e2smrcv1.NeighborCellItem, error) {

	return &e2smrcv1.NeighborCellItem{
		NeighborCellItem: &e2smrcv1.NeighborCellItem_RanTypeChoiceNr{
			RanTypeChoiceNr: &e2smrcv1.NeighborCellItemChoiceNr{
				NRCgi:           nrCgi,
				NRPci:           nrPci,
				FiveGsTac:       fiveGsTac,
				NRModeInfo:      nrModeInfo,
				NRFreqInfo:      nrFreqInfo,
				X2XnEstablished: x2XnEstablished,
				HOValidated:     hOValidated,
				Version:         version,
			},
		},
	}, nil
}

func CreateNeighborCellItemRanTypeChoiceEutra(eutraCgi *e2smcommonies.EutraCgi, eutraPci *e2smcommonies.EUtraPci,
	eutraArfcn *e2smcommonies.EUtraArfcn, eutraTac *e2smcommonies.EUtraTac, x2XnEstablished e2smrcv1.X2XNEstablished,
	hOValidated e2smrcv1.HOValidated, version int32) (*e2smrcv1.NeighborCellItem, error) {

	return &e2smrcv1.NeighborCellItem{
		NeighborCellItem: &e2smrcv1.NeighborCellItem_RanTypeChoiceEutra{
			RanTypeChoiceEutra: &e2smrcv1.NeighborCellItemChoiceEUtra{
				EUtraCgi:        eutraCgi,
				EUtraPci:        eutraPci,
				EUtraArfcn:      eutraArfcn,
				EUtraTac:        eutraTac,
				X2XnEstablished: x2XnEstablished,
				HOValidated:     hOValidated,
				Version:         version,
			},
		},
	}, nil
}

func CreateNrCgi(plmnID []byte, nrcellIdentity *asn1.BitString) (*e2smcommonies.NrCgi, error) {

	msg := &e2smcommonies.NrCgi{
		PLmnidentity: &e2smcommonies.Plmnidentity{
			Value: plmnID,
		},
		NRcellIdentity: &e2smcommonies.NrcellIdentity{
			Value: nrcellIdentity,
		},
	}

	return msg, nil
}

func CreateNrPci(nrPci int32) (*e2smcommonies.NrPci, error) {

	msg := &e2smcommonies.NrPci{
		Value: nrPci,
	}

	return msg, nil
}

func CreateFiveGsTac(fiveGsTac []byte) (*e2smcommonies.FiveGsTac, error) {

	msg := &e2smcommonies.FiveGsTac{
		Value: fiveGsTac,
	}

	return msg, nil
}

func CreateNRModeInfoFDD() e2smrcv1.NRModeInfo {
	return e2smrcv1.NRModeInfo_NR_MODE_INFO_FDD
}

func CreateNRModeInfoTDD() e2smrcv1.NRModeInfo {
	return e2smrcv1.NRModeInfo_NR_MODE_INFO_TDD
}

func CreateNrfrequencyInfo(nrArfcn *e2smcommonies.NrArfcn, nrfrequencyBandList *e2smcommonies.NrfrequencyBandList, nrfrequencyShift7P5Khz e2smcommonies.NrfrequencyShift7P5Khz) (*e2smcommonies.NrfrequencyInfo, error) {

	msg := &e2smcommonies.NrfrequencyInfo{
		NrArfcn:              nrArfcn,
		FrequencyBandList:    nrfrequencyBandList,
		FrequencyShift7P5Khz: &nrfrequencyShift7P5Khz,
	}

	return msg, nil
}

func CreateNrArfcn(nrArfcn int32) (*e2smcommonies.NrArfcn, error) {

	msg := &e2smcommonies.NrArfcn{
		NRarfcn: nrArfcn,
	}

	return msg, nil
}

func CreateNrfrequencyBandItem(freqBandIndicatorNr int32, supportedSulbandList *e2smcommonies.SupportedSulbandList) (*e2smcommonies.NrfrequencyBandItem, error) {

	msg := &e2smcommonies.NrfrequencyBandItem{
		FreqBandIndicatorNr:  freqBandIndicatorNr,
		SupportedSulbandList: supportedSulbandList,
	}

	return msg, nil
}

func CreateSupportedSulfreqBandItem(freqBandIndicatorNr int32) (*e2smcommonies.SupportedSulfreqBandItem, error) {

	msg := &e2smcommonies.SupportedSulfreqBandItem{
		FreqBandIndicatorNr: freqBandIndicatorNr,
	}

	return msg, nil
}

func CreateNrfrequencyShift7P5KhzFalse() e2smcommonies.NrfrequencyShift7P5Khz {
	return e2smcommonies.NrfrequencyShift7P5Khz_NRFREQUENCY_SHIFT7P5KHZ_FALSE
}

func CreateNrfrequencyShift7P5KhzTrue() e2smcommonies.NrfrequencyShift7P5Khz {
	return e2smcommonies.NrfrequencyShift7P5Khz_NRFREQUENCY_SHIFT7P5KHZ_TRUE
}

func CreateX2XNEstablishedFalse() e2smrcv1.X2XNEstablished {
	return e2smrcv1.X2XNEstablished_X2XN_ESTABLISHED_FALSE
}

func CreateX2XNEstablishedTrue() e2smrcv1.X2XNEstablished {
	return e2smrcv1.X2XNEstablished_X2XN_ESTABLISHED_TRUE
}

func CreateHOValidatedFalse() e2smrcv1.HOValidated {
	return e2smrcv1.HOValidated_HO_VALIDATED_FALSE
}

func CreateHOValidatedTrue() e2smrcv1.HOValidated {
	return e2smrcv1.HOValidated_HO_VALIDATED_TRUE
}

func CreateEutraCgi(plmnID []byte, eutraCellIdentity *asn1.BitString) (*e2smcommonies.EutraCgi, error) {

	msg := &e2smcommonies.EutraCgi{
		PLmnidentity: &e2smcommonies.Plmnidentity{
			Value: plmnID,
		},
		EUtracellIdentity: &e2smcommonies.EutracellIdentity{
			Value: eutraCellIdentity,
		},
	}

	return msg, nil
}

func CreateEutraPci(eutraPci int32) (*e2smcommonies.EUtraPci, error) {

	msg := &e2smcommonies.EUtraPci{
		Value: eutraPci,
	}

	return msg, nil
}

func CreateEutraArfcn(eutraArfcn int32) (*e2smcommonies.EUtraArfcn, error) {

	msg := &e2smcommonies.EUtraArfcn{
		Value: eutraArfcn,
	}

	return msg, nil
}

func CreateEutraTac(eutraTac []byte) (*e2smcommonies.EUtraTac, error) {

	msg := &e2smcommonies.EUtraTac{
		Value: eutraTac,
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat4ItemUe(ueID *e2smcommonies.Ueid, cgi *e2smcommonies.Cgi) (*e2smrcv1.E2SmRcIndicationMessageFormat4ItemUe, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat4ItemUe{
		UeId:         ueID,
		CellGlobalId: cgi,
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat4ItemCell(cgi *e2smcommonies.Cgi, cellContextInfo []byte, neighborRelationTable *e2smrcv1.NeighborRelationInfo) (*e2smrcv1.E2SmRcIndicationMessageFormat4ItemCell, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat4ItemCell{
		CellGlobalId: cgi,
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat5Item(ranParameterID int64, ranParameterValueType *e2smrcv1.RanparameterValueType) (*e2smrcv1.E2SmRcIndicationMessageFormat5Item, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat5Item{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterValueType: ranParameterValueType,
	}

	return msg, nil
}
