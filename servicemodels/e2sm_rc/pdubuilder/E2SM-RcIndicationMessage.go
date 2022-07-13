//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat1() error validating E2SM-RC PDU %s", err.Error())
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat2() error validating E2SM-RC PDU %s", err.Error())
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat3() error validating E2SM-RC PDU %s", err.Error())
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat4() error validating E2SM-RC PDU %s", err.Error())
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat5() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat6(ricInsertStyleList []*e2smrcv1.E2SmRcIndicationMessageFormat6StyleItem) (*e2smrcv1.E2SmRcIndicationMessage, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessage{
		RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats{
			RicIndicationMessageFormats: &e2smrcv1.RicIndicationMessageFormats_IndicationMessageFormat6{
				IndicationMessageFormat6: &e2smrcv1.E2SmRcIndicationMessageFormat6{
					RicInsertStyleList: ricInsertStyleList,
				},
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat6() error validating E2SM-RC PDU %s", err.Error())
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat1Item() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat2Item(ueID *e2smcommonies.Ueid, ranPList []*e2smrcv1.E2SmRcIndicationMessageFormat2RanparameterItem) (*e2smrcv1.E2SmRcIndicationMessageFormat2Item, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat2Item{
		UeId:     ueID,
		RanPList: ranPList,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat2Item() error validating E2SM-RC PDU %s", err.Error())
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat2RanparameterItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat3Item(cgi *e2smcommonies.Cgi) (*e2smrcv1.E2SmRcIndicationMessageFormat3Item, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat3Item{
		CellGlobalId: cgi,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat3Item() error validating E2SM-RC PDU %s", err.Error())
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNeighborRelationInfo() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateServingCellPciNR(nrPci int32) (*e2smcommonies.ServingCellPci, error) {

	msg := &e2smcommonies.ServingCellPci{
		ServingCellPci: &e2smcommonies.ServingCellPci_NR{
			NR: &e2smcommonies.NrPci{
				Value: nrPci,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellPciNR() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateServingCellPciEUtraPci(eutraPci int32) (*e2smcommonies.ServingCellPci, error) {

	msg := &e2smcommonies.ServingCellPci{
		ServingCellPci: &e2smcommonies.ServingCellPci_EUtra{
			EUtra: &e2smcommonies.EUtraPci{
				Value: eutraPci,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellPciEUtraPci() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateServingCellArfcnNR(nrArfcn int32) (*e2smcommonies.ServingCellArfcn, error) {

	msg := &e2smcommonies.ServingCellArfcn{
		ServingCellArfcn: &e2smcommonies.ServingCellArfcn_NR{
			NR: &e2smcommonies.NrArfcn{
				NRarfcn: nrArfcn,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellArfcnNR() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateServingCellArfcnEUtraPci(eutraArfcn int32) (*e2smcommonies.ServingCellArfcn, error) {

	msg := &e2smcommonies.ServingCellArfcn{
		ServingCellArfcn: &e2smcommonies.ServingCellArfcn_EUtra{
			EUtra: &e2smcommonies.EUtraArfcn{
				Value: eutraArfcn,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellArfcnEUtraPci() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNeighborCellItemRanTypeChoiceNr(nrCgi *e2smcommonies.NrCgi, nrPci int32, fiveGsTac []byte,
	nrModeInfo e2smrcv1.NRModeInfo, nrFreqInfo *e2smcommonies.NrfrequencyInfo, x2XnEstablished e2smrcv1.X2XNEstablished,
	hOValidated e2smrcv1.HOValidated, version int32) (*e2smrcv1.NeighborCellItem, error) {

	if len(fiveGsTac) != 3 {
		return nil, errors.NewInvalid("CreateNeighborCellItemRanTypeChoiceNr() FiveGsTac should be of length 3 bytes, got %v", fiveGsTac)
	}

	msg := &e2smrcv1.NeighborCellItem{
		NeighborCellItem: &e2smrcv1.NeighborCellItem_RanTypeChoiceNr{
			RanTypeChoiceNr: &e2smrcv1.NeighborCellItemChoiceNr{
				NRCgi: nrCgi,
				NRPci: &e2smcommonies.NrPci{
					Value: nrPci,
				},
				FiveGsTac: &e2smcommonies.FiveGsTac{
					Value: fiveGsTac,
				},
				NRModeInfo:      nrModeInfo,
				NRFreqInfo:      nrFreqInfo,
				X2XnEstablished: x2XnEstablished,
				HOValidated:     hOValidated,
				Version:         version,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNeighborCellItemRanTypeChoiceNr() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNeighborCellItemRanTypeChoiceEutra(eutraCgi *e2smcommonies.EutraCgi, eutraPci int32,
	eutraArfcn int32, eutraTac []byte, x2XnEstablished e2smrcv1.X2XNEstablished,
	hOValidated e2smrcv1.HOValidated, version int32) (*e2smrcv1.NeighborCellItem, error) {

	if len(eutraTac) != 2 {
		return nil, errors.NewInvalid("CreateNeighborCellItemRanTypeChoiceEutra() EutraTac should be of length 2 bytes, got %v", eutraTac)
	}

	msg := &e2smrcv1.NeighborCellItem{
		NeighborCellItem: &e2smrcv1.NeighborCellItem_RanTypeChoiceEutra{
			RanTypeChoiceEutra: &e2smrcv1.NeighborCellItemChoiceEUtra{
				EUtraCgi: eutraCgi,
				EUtraPci: &e2smcommonies.EUtraPci{
					Value: eutraPci,
				},
				EUtraArfcn: &e2smcommonies.EUtraArfcn{
					Value: eutraArfcn,
				},
				EUtraTac: &e2smcommonies.EUtraTac{
					Value: eutraTac,
				},
				X2XnEstablished: x2XnEstablished,
				HOValidated:     hOValidated,
				Version:         version,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNeighborCellItemRanTypeChoiceEutra() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrCgi() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrPci(nrPci int32) (*e2smcommonies.NrPci, error) {

	msg := &e2smcommonies.NrPci{
		Value: nrPci,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrPci() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateFiveGsTac(fiveGsTac []byte) (*e2smcommonies.FiveGsTac, error) {

	msg := &e2smcommonies.FiveGsTac{
		Value: fiveGsTac,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateFiveGsTac() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNRModeInfoFDD() e2smrcv1.NRModeInfo {
	return e2smrcv1.NRModeInfo_NR_MODE_INFO_FDD
}

func CreateNRModeInfoTDD() e2smrcv1.NRModeInfo {
	return e2smrcv1.NRModeInfo_NR_MODE_INFO_TDD
}

func CreateNrfrequencyInfo(nrArfcn *e2smcommonies.NrArfcn, nrfrequencyBandList *e2smcommonies.NrfrequencyBandList) (*e2smcommonies.NrfrequencyInfo, error) {

	msg := &e2smcommonies.NrfrequencyInfo{
		NrArfcn:           nrArfcn,
		FrequencyBandList: nrfrequencyBandList,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrfrequencyInfo() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrArfcn(nrArfcn int32) (*e2smcommonies.NrArfcn, error) {

	msg := &e2smcommonies.NrArfcn{
		NRarfcn: nrArfcn,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrArfcn() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrfrequencyBandItem(freqBandIndicatorNr int32, supportedSulbandList *e2smcommonies.SupportedSulbandList) (*e2smcommonies.NrfrequencyBandItem, error) {

	msg := &e2smcommonies.NrfrequencyBandItem{
		FreqBandIndicatorNr:  freqBandIndicatorNr,
		SupportedSulbandList: supportedSulbandList,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrfrequencyBandItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSupportedSulfreqBandItem(freqBandIndicatorNr int32) (*e2smcommonies.SupportedSulfreqBandItem, error) {

	msg := &e2smcommonies.SupportedSulfreqBandItem{
		FreqBandIndicatorNr: freqBandIndicatorNr,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSupportedSulfreqBandItem() error validating E2SM-RC PDU %s", err.Error())
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEutraCgi() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEutraPci(eutraPci int32) (*e2smcommonies.EUtraPci, error) {

	msg := &e2smcommonies.EUtraPci{
		Value: eutraPci,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEutraPci() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEutraArfcn(eutraArfcn int32) (*e2smcommonies.EUtraArfcn, error) {

	msg := &e2smcommonies.EUtraArfcn{
		Value: eutraArfcn,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEutraArfcn() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEutraTac(eutraTac []byte) (*e2smcommonies.EUtraTac, error) {

	msg := &e2smcommonies.EUtraTac{
		Value: eutraTac,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEutraTac() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat4ItemUe(ueID *e2smcommonies.Ueid, cgi *e2smcommonies.Cgi) (*e2smrcv1.E2SmRcIndicationMessageFormat4ItemUe, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat4ItemUe{
		UeId:         ueID,
		CellGlobalId: cgi,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat4ItemUe() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat4ItemCell(cgi *e2smcommonies.Cgi) (*e2smrcv1.E2SmRcIndicationMessageFormat4ItemCell, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat4ItemCell{
		CellGlobalId: cgi,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat4ItemCell() error validating E2SM-RC PDU %s", err.Error())
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

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat5Item() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGNb(amf int64, plmnID []byte, amfRegionID []byte, amfSetID []byte, amfPointer []byte) (*e2smcommonies.Ueid, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("CreateUeIDGNb() PlmnID should contain only 3 bytes, got %v", len(plmnID))
	}
	if len(amfRegionID) != 1 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfRegionID should contain only 1 byte, got %v", len(amfRegionID))
	}
	if len(amfSetID) != 2 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfSetID should contain only 2 bytes, got %v", len(amfSetID))
	}
	if amfSetID[1]&0x3F > 0 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfSetID should contain only 10 bits, i.e., last 6 bits of last byte should be trailing zeros, got %v", amfSetID[1])
	}
	if len(amfPointer) != 1 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfPointer should contain only 1 byte, got %v", len(amfPointer))
	}
	if amfPointer[0]&0x03 > 0 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfSetID should contain only 6 bits, i.e., last 2 bits should be trailing zeros, got %v", amfPointer)
	}

	ueID := &e2smcommonies.Ueid{
		Ueid: &e2smcommonies.Ueid_GNbUeid{
			GNbUeid: &e2smcommonies.UeidGnb{
				AmfUeNgapId: &e2smcommonies.AmfUeNgapId{
					Value: amf,
				},
				Guami: &e2smcommonies.Guami{
					PLmnidentity: &e2smcommonies.Plmnidentity{
						Value: plmnID,
					},
					AMfregionId: &e2smcommonies.AmfregionId{
						Value: &asn1.BitString{
							Value: amfRegionID,
							Len:   8,
						},
					},
					AMfsetId: &e2smcommonies.AmfsetId{
						Value: &asn1.BitString{
							Value: amfSetID,
							Len:   10,
						},
					},
					AMfpointer: &e2smcommonies.Amfpointer{
						Value: &asn1.BitString{
							Value: amfPointer,
							Len:   6,
						},
					},
				},
			},
		},
	}

	if err := ueID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGNb() validation of UeID failed with %v", err)
	}

	return ueID, nil
}

func CreateUeIDGnbDu(gNbCuUeF1ApID int64) (*e2smcommonies.Ueid, error) {

	msg := &e2smcommonies.Ueid{
		Ueid: &e2smcommonies.Ueid_GNbDuUeid{
			GNbDuUeid: &e2smcommonies.UeidGnbDu{
				GNbCuUeF1ApId: &e2smcommonies.GnbCuUeF1ApId{
					Value: gNbCuUeF1ApID,
				},
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGnbDu() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateUeIDGnbCuUp(gNbCuCpUeE1ApID int64) (*e2smcommonies.Ueid, error) {

	msg := &e2smcommonies.Ueid{
		Ueid: &e2smcommonies.Ueid_GNbCuUpUeid{
			GNbCuUpUeid: &e2smcommonies.UeidGnbCuUp{
				GNbCuCpUeE1ApId: &e2smcommonies.GnbCuCpUeE1ApId{
					Value: gNbCuCpUeE1ApID,
				},
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGnbCuUp() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateUeIDNgEnb(amf int64, plmnID []byte, amfRegionID []byte, amfSetID []byte, amfPointer []byte) (*e2smcommonies.Ueid, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("CreateUeIDNgEnb() PlmnID should contain only 3 bytes, got %v", len(plmnID))
	}
	if len(amfRegionID) != 1 {
		return nil, errors.NewInvalid("CreateUeIDNgEnb() AMfRegionID should contain only 1 byte, got %v", len(amfRegionID))
	}
	if len(amfSetID) != 2 {
		return nil, errors.NewInvalid("CreateUeIDNgEnb() AMfSetID should contain only 2 bytes, got %v", len(amfSetID))
	}
	if amfSetID[1]&0x3F > 0 {
		return nil, errors.NewInvalid("CreateUeIDNgEnb() AMfSetID should contain only 10 bits, i.e., last 6 bits of last byte should be trailing zeros, got %v", amfSetID[1])
	}
	if len(amfPointer) != 1 {
		return nil, errors.NewInvalid("CreateUeIDNgEnb() AMfPointer should contain only 1 byte, got %v", len(amfPointer))
	}
	if amfPointer[0]&0x03 > 0 {
		return nil, errors.NewInvalid("CreateUeIDNgEnb() AMfSetID should contain only 6 bits, i.e., last 2 bits should be trailing zeros, got %v", amfPointer)
	}

	ueID := &e2smcommonies.Ueid{
		Ueid: &e2smcommonies.Ueid_NgENbUeid{
			NgENbUeid: &e2smcommonies.UeidNgEnb{
				AmfUeNgapId: &e2smcommonies.AmfUeNgapId{
					Value: amf,
				},
				Guami: &e2smcommonies.Guami{
					PLmnidentity: &e2smcommonies.Plmnidentity{
						Value: plmnID,
					},
					AMfregionId: &e2smcommonies.AmfregionId{
						Value: &asn1.BitString{
							Value: amfRegionID,
							Len:   8,
						},
					},
					AMfsetId: &e2smcommonies.AmfsetId{
						Value: &asn1.BitString{
							Value: amfSetID,
							Len:   10,
						},
					},
					AMfpointer: &e2smcommonies.Amfpointer{
						Value: &asn1.BitString{
							Value: amfPointer,
							Len:   6,
						},
					},
				},
			},
		},
	}

	if err := ueID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDNgEnb() validation of UeID failed with %v", err)
	}

	return ueID, nil
}

func CreateUeIDNgEnbDu(ngENbCuUeW1ApID int64) (*e2smcommonies.Ueid, error) {

	msg := &e2smcommonies.Ueid{
		Ueid: &e2smcommonies.Ueid_NgENbDuUeid{
			NgENbDuUeid: &e2smcommonies.UeidNgEnbDu{
				NgENbCuUeW1ApId: &e2smcommonies.NgenbCuUeW1ApId{
					Value: ngENbCuUeW1ApID,
				},
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDNgEnbDu() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateUeIDEnbGnb(mENbUeX2ApID int32, plmnID []byte, enbID *e2smcommonies.EnbId) (*e2smcommonies.Ueid, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("CreateUeIDEnbGnb() PlmnID should contain only 3 bytes, got %v", len(plmnID))
	}

	msg := &e2smcommonies.Ueid{
		Ueid: &e2smcommonies.Ueid_EnGNbUeid{
			EnGNbUeid: &e2smcommonies.UeidEnGnb{
				MENbUeX2ApId: &e2smcommonies.EnbUeX2ApId{
					Value: mENbUeX2ApID,
				},
				GlobalEnbId: &e2smcommonies.GlobalEnbId{
					PLmnidentity: &e2smcommonies.Plmnidentity{
						Value: plmnID,
					},
					ENbId: enbID,
				},
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDEnbGnb() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateUeIDEnb(mMeUeS1ApID int64, plmnID []byte, mMeGroupID []byte, mMeCode []byte) (*e2smcommonies.Ueid, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("CreateUeIDEnb() PlmnID should contain only 3 bytes, got %v", len(plmnID))
	}
	if len(mMeGroupID) != 2 {
		return nil, errors.NewInvalid("CreateUeIDEnb() MMeGroupID should contain only 2 byte, got %v", len(mMeGroupID))
	}
	if len(mMeCode) != 1 {
		return nil, errors.NewInvalid("CreateUeIDEnb() MMeCode should contain only 1 byte, got %v", len(mMeCode))
	}

	ueID := &e2smcommonies.Ueid{
		Ueid: &e2smcommonies.Ueid_ENbUeid{
			ENbUeid: &e2smcommonies.UeidEnb{
				MMeUeS1ApId: &e2smcommonies.MmeUeS1ApId{
					Value: mMeUeS1ApID,
				},
				GUmmei: &e2smcommonies.Gummei{
					PLmnIdentity: &e2smcommonies.Plmnidentity{
						Value: plmnID,
					},
					MMeGroupId: &e2smcommonies.MmeGroupId{
						Value: mMeGroupID,
					},
					MMeCode: &e2smcommonies.MmeCode{
						Value: mMeCode,
					},
				},
			},
		},
	}

	if err := ueID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDEnb() validation of UeID failed with %v", err)
	}

	return ueID, nil
}

func CreateEnbIDMacro(macroENbID []byte) (*e2smcommonies.EnbId, error) {

	if macroENbID[2]&0x0F > 0 {
		return nil, errors.NewInvalid("CreateEnbIDMacro() MacroENbID should contain only 20 bits, i.e., last 4 bits should be trailing zeros, got %v", macroENbID)
	}

	msg := &e2smcommonies.EnbId{
		EnbId: &e2smcommonies.EnbId_MacroENbId{
			MacroENbId: &asn1.BitString{
				Value: nil,
				Len:   20,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbIDMacro() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateEnbIDHome(homeENbID []byte) (*e2smcommonies.EnbId, error) {

	if homeENbID[3]&0x0F > 0 {
		return nil, errors.NewInvalid("CreateEnbIDHome() HomeENbID should contain only 28 bits, i.e., last 4 bits should be trailing zeros, got %v", homeENbID)
	}

	msg := &e2smcommonies.EnbId{
		EnbId: &e2smcommonies.EnbId_HomeENbId{
			HomeENbId: &asn1.BitString{
				Value: homeENbID,
				Len:   28,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbIDHome() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateEnbIDShortMacro(shortMacroENbID []byte) (*e2smcommonies.EnbId, error) {

	if shortMacroENbID[2]&0x3F > 0 {
		return nil, errors.NewInvalid("CreateEnbIDShortMacro() ShortMacroENbID should contain only 18 bits, i.e., last 6 bits should be trailing zeros, got %v", shortMacroENbID)
	}

	msg := &e2smcommonies.EnbId{
		EnbId: &e2smcommonies.EnbId_ShortMacroENbId{
			ShortMacroENbId: &asn1.BitString{
				Value: shortMacroENbID,
				Len:   18,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbIDShortMacro() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateEnbIDLongMacro(longMacroENbID []byte) (*e2smcommonies.EnbId, error) {

	if longMacroENbID[2]&0x07 > 0 {
		return nil, errors.NewInvalid("CreateEnbIDLongMacro() LongMacroENbID should contain only 21 bits, i.e., last 3 bits should be trailing zeros, got %v", longMacroENbID)
	}

	msg := &e2smcommonies.EnbId{
		EnbId: &e2smcommonies.EnbId_LongMacroENbId{
			LongMacroENbId: &asn1.BitString{
				Value: longMacroENbID,
				Len:   21,
			},
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbIDLongMacro() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateCgiNRCgi(nRCgi *e2smcommonies.NrCgi) (*e2smcommonies.Cgi, error) {

	msg := &e2smcommonies.Cgi{
		Cgi: &e2smcommonies.Cgi_NRCgi{
			NRCgi: nRCgi,
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCgiNRCgi() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateCgiEutraCgi(eUtraCgi *e2smcommonies.EutraCgi) (*e2smcommonies.Cgi, error) {

	msg := &e2smcommonies.Cgi{
		Cgi: &e2smcommonies.Cgi_EUtraCgi{
			EUtraCgi: eUtraCgi,
		},
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCgiNRCgi() validation of UeID failed with %v", err)
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat6StyleItem(indicatedInsertStyleType int32, ricInsertIndicationList []*e2smrcv1.E2SmRcIndicationMessageFormat6IndicationItem) (*e2smrcv1.E2SmRcIndicationMessageFormat6StyleItem, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat6StyleItem{
		IndicatedInsertStyleType: &e2smcommonies.RicStyleType{
			Value: indicatedInsertStyleType,
		},
		RicInsertIndicationList: ricInsertIndicationList,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat6StyleItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat6IndicationItem(ricInsertIndicationID int32, ranPInsertIndicationList []*e2smrcv1.E2SmRcIndicationMessageFormat6RanpItem) (*e2smrcv1.E2SmRcIndicationMessageFormat6IndicationItem, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat6IndicationItem{
		RicInsertIndicationId: &e2smrcv1.RicInsertIndicationId{
			Value: ricInsertIndicationID,
		},
		RanPInsertIndicationList: ranPInsertIndicationList,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat6IndicationItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat6RanpItem(ranParameterID int64, ranParameterValueType *e2smrcv1.RanparameterValueType) (*e2smrcv1.E2SmRcIndicationMessageFormat6RanpItem, error) {

	msg := &e2smrcv1.E2SmRcIndicationMessageFormat6RanpItem{
		RanParameterId: &e2smrcv1.RanparameterId{
			Value: ranParameterID,
		},
		RanParameterValueType: ranParameterValueType,
	}

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat6RanpItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return msg, nil
}
