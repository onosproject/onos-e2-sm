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

func CreateE2SmRcEventTriggerFormat1(msgl []*e2smrcv1.E2SmRcEventTriggerFormat1Item) (*e2smrcv1.E2SmRcEventTrigger, error) {

	ch := &e2smrcv1.E2SmRcEventTrigger{
		RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats{
			RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats_EventTriggerFormat1{
				EventTriggerFormat1: &e2smrcv1.E2SmRcEventTriggerFormat1{
					MessageList: msgl,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat1() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcEventTriggerFormat2(rcptID int32, rcpbID int32) (*e2smrcv1.E2SmRcEventTrigger, error) {

	ch := &e2smrcv1.E2SmRcEventTrigger{
		RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats{
			RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats_EventTriggerFormat2{
				EventTriggerFormat2: &e2smrcv1.E2SmRcEventTriggerFormat2{
					RicCallProcessTypeId: &e2smrcv1.RicCallProcessTypeId{
						Value: rcptID,
					},
					RicCallProcessBreakpointId: &e2smrcv1.RicCallProcessBreakpointId{
						Value: rcpbID,
					},
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat2() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcEventTriggerFormat3(e2nichl []*e2smrcv1.E2SmRcEventTriggerFormat3Item) (*e2smrcv1.E2SmRcEventTrigger, error) {

	ch := &e2smrcv1.E2SmRcEventTrigger{
		RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats{
			RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats_EventTriggerFormat3{
				EventTriggerFormat3: &e2smrcv1.E2SmRcEventTriggerFormat3{
					E2NodeInfoChangeList: e2nichl,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat3() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcEventTriggerFormat4(ueInfoChangeList []*e2smrcv1.E2SmRcEventTriggerFormat4Item) (*e2smrcv1.E2SmRcEventTrigger, error) {

	ch := &e2smrcv1.E2SmRcEventTrigger{
		RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats{
			RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats_EventTriggerFormat4{
				EventTriggerFormat4: &e2smrcv1.E2SmRcEventTriggerFormat4{
					UEinfoChangeList: ueInfoChangeList,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat4() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcEventTriggerFormat5(onDemand e2smrcv1.OnDemand) (*e2smrcv1.E2SmRcEventTrigger, error) {

	ch := &e2smrcv1.E2SmRcEventTrigger{
		RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats{
			RicEventTriggerFormats: &e2smrcv1.RicEventTriggerFormats_EventTriggerFormat5{
				EventTriggerFormat5: &e2smrcv1.E2SmRcEventTriggerFormat5{
					OnDemand: onDemand,
				},
			},
		},
	}

	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat5() error validating E2SM-RC PDU %s", err.Error())
	}

	return ch, nil
}

func CreateE2SmRcEventTriggerFormat1Item(recID int32, msgt *e2smrcv1.MessageTypeChoice) (*e2smrcv1.E2SmRcEventTriggerFormat1Item, error) {

	item := &e2smrcv1.E2SmRcEventTriggerFormat1Item{
		RicEventTriggerConditionId: &e2smrcv1.RicEventTriggerConditionId{
			Value: recID,
		},
		MessageType: msgt,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat1Item() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateMessageTypeChoiceNi(niType e2smcommonies.InterfaceType, niID *e2smcommonies.InterfaceIdentifier, ipID int32,
	msgt e2smcommonies.MessageType) (*e2smrcv1.MessageTypeChoice, error) {

	ni := &e2smrcv1.MessageTypeChoiceNi{
		NIType: niType,
	}

	item := &e2smrcv1.MessageTypeChoice{
		MessageTypeChoice: &e2smrcv1.MessageTypeChoice_MessageTypeChoiceNi{
			MessageTypeChoiceNi: ni,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMessageTypeChoiceNi() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateMessageTypeChoiceRrc(rrcType *e2smcommonies.RrcType, msgID int64) (*e2smrcv1.MessageTypeChoice, error) {

	rrc := &e2smrcv1.MessageTypeChoiceRrc{
		RRcMessage: &e2smcommonies.RrcMessageId{
			RrcType:   rrcType,
			MessageId: msgID,
		},
	}

	item := &e2smrcv1.MessageTypeChoice{
		MessageTypeChoice: &e2smrcv1.MessageTypeChoice_MessageTypeChoiceRrc{
			MessageTypeChoiceRrc: rrc,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMessageTypeChoiceRrc() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateInterfaceTypeNG() e2smcommonies.InterfaceType {
	return e2smcommonies.InterfaceType_INTERFACE_TYPE_N_G
}

func CreateInterfaceTypeXN() e2smcommonies.InterfaceType {
	return e2smcommonies.InterfaceType_INTERFACE_TYPE_XN
}

func CreateInterfaceTypeF1() e2smcommonies.InterfaceType {
	return e2smcommonies.InterfaceType_INTERFACE_TYPE_F1
}

func CreateInterfaceTypeE1() e2smcommonies.InterfaceType {
	return e2smcommonies.InterfaceType_INTERFACE_TYPE_E1
}

func CreateInterfaceTypeS1() e2smcommonies.InterfaceType {
	return e2smcommonies.InterfaceType_INTERFACE_TYPE_S1
}

func CreateInterfaceTypeX2() e2smcommonies.InterfaceType {
	return e2smcommonies.InterfaceType_INTERFACE_TYPE_X2
}

func CreateInterfaceTypeW1() e2smcommonies.InterfaceType {
	return e2smcommonies.InterfaceType_INTERFACE_TYPE_W1
}

func CreateInterfaceIdentifierNG(plmnID []byte, regionID *asn1.BitString, setID *asn1.BitString, pointer *asn1.BitString) (*e2smcommonies.InterfaceIdentifier, error) {

	item := &e2smcommonies.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommonies.InterfaceIdentifier_NG{
			NG: &e2smcommonies.InterfaceIdNg{
				Guami: &e2smcommonies.Guami{
					PLmnidentity: &e2smcommonies.Plmnidentity{
						Value: plmnID,
					},
					AMfregionId: &e2smcommonies.AmfregionId{
						Value: regionID,
					},
					AMfsetId: &e2smcommonies.AmfsetId{
						Value: setID,
					},
					AMfpointer: &e2smcommonies.Amfpointer{
						Value: pointer,
					},
				},
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIdentifierNG() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateInterfaceIdentifierXN(gngrnID *e2smcommonies.GlobalNgrannodeId) (*e2smcommonies.InterfaceIdentifier, error) {

	item := &e2smcommonies.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommonies.InterfaceIdentifier_XN{
			XN: &e2smcommonies.InterfaceIdXn{
				GlobalNgRanId: gngrnID,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIdentifierXN() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateInterfaceIdentifierF1(plmnID []byte, gnbID *asn1.BitString, duID int64) (*e2smcommonies.InterfaceIdentifier, error) {

	item := &e2smcommonies.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommonies.InterfaceIdentifier_F1{
			F1: &e2smcommonies.InterfaceIdF1{
				GlobalGnbId: &e2smcommonies.GlobalGnbId{
					PLmnidentity: &e2smcommonies.Plmnidentity{
						Value: plmnID,
					},
					GNbId: &e2smcommonies.GnbId{
						GnbId: &e2smcommonies.GnbId_GNbId{
							GNbId: gnbID,
						},
					},
				},
				GNbDuId: &e2smcommonies.GnbDuId{
					Value: duID,
				},
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIdentifierF1() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateInterfaceIdentifierE1(plmnID []byte, gnbID *asn1.BitString, cuupID int64) (*e2smcommonies.InterfaceIdentifier, error) {

	item := &e2smcommonies.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommonies.InterfaceIdentifier_E1{
			E1: &e2smcommonies.InterfaceIdE1{
				GlobalGnbId: &e2smcommonies.GlobalGnbId{
					PLmnidentity: &e2smcommonies.Plmnidentity{
						Value: plmnID,
					},
					GNbId: &e2smcommonies.GnbId{
						GnbId: &e2smcommonies.GnbId_GNbId{
							GNbId: gnbID,
						},
					},
				},
				GNbCuUpId: &e2smcommonies.GnbCuUpId{
					Value: cuupID,
				},
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIdentifierE1() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateInterfaceIdentifierS1(plmnID []byte, groupID []byte, code []byte) (*e2smcommonies.InterfaceIdentifier, error) {

	item := &e2smcommonies.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommonies.InterfaceIdentifier_S1{
			S1: &e2smcommonies.InterfaceIdS1{
				GUmmei: &e2smcommonies.Gummei{
					PLmnIdentity: &e2smcommonies.Plmnidentity{
						Value: plmnID,
					},
					MMeGroupId: &e2smcommonies.MmeGroupId{
						Value: groupID,
					},
					MMeCode: &e2smcommonies.MmeCode{
						Value: code,
					},
				},
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIdentifierS1() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateInterfaceIdentifierX2(nodeType *e2smcommonies.NodeType) (*e2smcommonies.InterfaceIdentifier, error) {

	item := &e2smcommonies.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommonies.InterfaceIdentifier_X2{
			X2: &e2smcommonies.InterfaceIdX2{
				NodeType: nodeType,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIdentifierX2() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateInterfaceIdentifierW1(plmnID []byte, ngEnbID *e2smcommonies.NgEnbId, duID int64) (*e2smcommonies.InterfaceIdentifier, error) {

	item := &e2smcommonies.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommonies.InterfaceIdentifier_W1{
			W1: &e2smcommonies.InterfaceIdW1{
				GlobalNgENbId: &e2smcommonies.GlobalNgEnbId{
					PLmnidentity: &e2smcommonies.Plmnidentity{
						Value: plmnID,
					},
					NgEnbId: ngEnbID,
				},
				NgENbDuId: &e2smcommonies.NgenbDuId{
					Value: duID,
				},
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIdentifierW1() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateGlobalNgrannodeIDGnb(plmnID []byte, gnbID *asn1.BitString) (*e2smcommonies.GlobalNgrannodeId, error) {

	item := &e2smcommonies.GlobalNgrannodeId{
		GlobalNgrannodeId: &e2smcommonies.GlobalNgrannodeId_GNb{
			GNb: &e2smcommonies.GlobalGnbId{
				PLmnidentity: &e2smcommonies.Plmnidentity{
					Value: plmnID,
				},
				GNbId: &e2smcommonies.GnbId{
					GnbId: &e2smcommonies.GnbId_GNbId{
						GNbId: gnbID,
					},
				},
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalNgrannodeIDGnb() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateGlobalNgrannodeIDNgEnb(plmnID []byte, ngEnbID *e2smcommonies.NgEnbId) (*e2smcommonies.GlobalNgrannodeId, error) {

	item := &e2smcommonies.GlobalNgrannodeId{
		GlobalNgrannodeId: &e2smcommonies.GlobalNgrannodeId_NgENb{
			NgENb: &e2smcommonies.GlobalNgEnbId{
				PLmnidentity: &e2smcommonies.Plmnidentity{
					Value: plmnID,
				},
				NgEnbId: ngEnbID,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalNgrannodeIDNgEnb() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateNgEnbIDMacro(macro *asn1.BitString) (*e2smcommonies.NgEnbId, error) {

	item := &e2smcommonies.NgEnbId{
		NgEnbId: &e2smcommonies.NgEnbId_MacroNgEnbId{
			MacroNgEnbId: macro,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgEnbIDMacro() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateNgEnbIDShortMacro(shMacro *asn1.BitString) (*e2smcommonies.NgEnbId, error) {

	item := &e2smcommonies.NgEnbId{
		NgEnbId: &e2smcommonies.NgEnbId_ShortMacroNgEnbId{
			ShortMacroNgEnbId: shMacro,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgEnbIDShortMacro() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateNgEnbIDLongMacro(lMacro *asn1.BitString) (*e2smcommonies.NgEnbId, error) {

	item := &e2smcommonies.NgEnbId{
		NgEnbId: &e2smcommonies.NgEnbId_LongMacroNgEnbId{
			LongMacroNgEnbId: lMacro,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgEnbIDLongMacro() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateNodeTypeEnbID(plmnID []byte, enbID *e2smcommonies.EnbId) (*e2smcommonies.NodeType, error) {

	item := &e2smcommonies.NodeType{
		NodeType: &e2smcommonies.NodeType_GlobalEnbId{
			GlobalEnbId: &e2smcommonies.GlobalEnbId{
				PLmnidentity: &e2smcommonies.Plmnidentity{
					Value: plmnID,
				},
				ENbId: enbID,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNodeTypeEnbID() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateNodeTypeEnGnbID(plmnID []byte, enGnbID *asn1.BitString) (*e2smcommonies.NodeType, error) {

	item := &e2smcommonies.NodeType{
		NodeType: &e2smcommonies.NodeType_GlobalEnGnbId{
			GlobalEnGnbId: &e2smcommonies.GlobalenGnbId{
				PLmnIdentity: &e2smcommonies.Plmnidentity{
					Value: plmnID,
				},
				EnGNbId: &e2smcommonies.EnGnbId{
					EnGnbId: &e2smcommonies.EnGnbId_EnGNbId{
						EnGNbId: enGnbID,
					},
				},
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNodeTypeEnGnbID() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateMessageTypeInitiatingMessage() e2smcommonies.MessageType {
	return e2smcommonies.MessageType_MESSAGE_TYPE_INITIATING_MESSAGE
}

func CreateMessageTypeSuccessfulOutcome() e2smcommonies.MessageType {
	return e2smcommonies.MessageType_MESSAGE_TYPE_SUCCESSFUL_OUTCOME
}

func CreateMessageTypeUnsuccessfulOutcome() e2smcommonies.MessageType {
	return e2smcommonies.MessageType_MESSAGE_TYPE_UNSUCCESSFUL_OUTCOME
}

func CreateRrcTypeLte(lte e2smcommonies.RrcclassLte) (*e2smcommonies.RrcType, error) {

	item := &e2smcommonies.RrcType{
		RrcType: &e2smcommonies.RrcType_Lte{
			Lte: lte,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRrcTypeLte() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateRrcTypeNr(nr e2smcommonies.RrcclassNr) (*e2smcommonies.RrcType, error) {

	item := &e2smcommonies.RrcType{
		RrcType: &e2smcommonies.RrcType_Nr{
			Nr: nr,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRrcTypeNr() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateMessageDirectionIncoming() e2smrcv1.MessageDirection {
	return e2smrcv1.MessageDirection_MESSAGE_DIRECTION_INCOMING
}

func CreateMessageDirectionOutgoing() e2smrcv1.MessageDirection {
	return e2smrcv1.MessageDirection_MESSAGE_DIRECTION_OUTGOING
}

func CreateEventTriggerUeInfoItem(retrueID int32, ueType *e2smrcv1.UeType) (*e2smrcv1.EventTriggerUeInfoItem, error) {

	item := &e2smrcv1.EventTriggerUeInfoItem{
		EventTriggerUeid: &e2smrcv1.RicEventTriggerUeId{
			Value: retrueID,
		},
		UeType: ueType,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerUeInfoItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateUeTypeIndividual(ueID *e2smcommonies.Ueid, rpt *e2smrcv1.RanparameterTesting) (*e2smrcv1.UeType, error) {

	item := &e2smrcv1.UeType{
		UeType: &e2smrcv1.UeType_UeTypeChoiceIndividual{
			UeTypeChoiceIndividual: &e2smrcv1.EventTriggerUeInfoItemChoiceIndividual{
				UeId:                ueID,
				RanParameterTesting: rpt,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeTypeIndividual() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateUeTypeGroup(rpt *e2smrcv1.RanparameterTesting) (*e2smrcv1.UeType, error) {

	item := &e2smrcv1.UeType{
		UeType: &e2smrcv1.UeType_UeTypeChoiceGroup{
			UeTypeChoiceGroup: &e2smrcv1.EventTriggerUeInfoItemChoiceGroup{
				RanParameterTesting: rpt,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeTypeGroup() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateEventTriggerUeeventInfoItem(retrueeID int32) (*e2smrcv1.EventTriggerUeeventInfoItem, error) {

	item := &e2smrcv1.EventTriggerUeeventInfoItem{
		UeEventId: &e2smrcv1.RicEventTriggerUeeventId{
			Value: retrueeID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerUeeventInfoItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateE2SmRcEventTriggerFormat3Item(recID int32, changeID int32) (*e2smrcv1.E2SmRcEventTriggerFormat3Item, error) {

	item := &e2smrcv1.E2SmRcEventTriggerFormat3Item{
		RicEventTriggerConditionId: &e2smrcv1.RicEventTriggerConditionId{
			Value: recID,
		},
		E2NodeInfoChangeId: changeID,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat3Item() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateEventTriggerCellInfoItem(retcID int32, cellType *e2smrcv1.CellType) (*e2smrcv1.EventTriggerCellInfoItem, error) {

	item := &e2smrcv1.EventTriggerCellInfoItem{
		EventTriggerCellId: &e2smrcv1.RicEventTriggerCellId{
			Value: retcID,
		},
		CellType: cellType,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerCellInfoItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateCellTypeChoiceIndividual(cgi *e2smcommonies.Cgi) (*e2smrcv1.CellType, error) {

	item := &e2smrcv1.CellType{
		CellType: &e2smrcv1.CellType_CellTypeChoiceIndividual{
			CellTypeChoiceIndividual: &e2smrcv1.EventTriggerCellInfoItemChoiceIndividual{
				CellGlobalId: cgi,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellTypeChoiceIndividual() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateCellTypeChoiceGroup(ranParameterTesting *e2smrcv1.RanparameterTesting) (*e2smrcv1.CellType, error) {

	item := &e2smrcv1.CellType{
		CellType: &e2smrcv1.CellType_CellTypeChoiceGroup{
			CellTypeChoiceGroup: &e2smrcv1.EventTriggerCellInfoItemChoiceGroup{
				RanParameterTesting: ranParameterTesting,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellTypeChoiceGroup() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateE2SmRcEventTriggerFormat4Item(recID int32, triggerType *e2smrcv1.TriggerTypeChoice) (*e2smrcv1.E2SmRcEventTriggerFormat4Item, error) {

	item := &e2smrcv1.E2SmRcEventTriggerFormat4Item{
		RicEventTriggerConditionId: &e2smrcv1.RicEventTriggerConditionId{
			Value: recID,
		},
		TriggerType: triggerType,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat4Item() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateTriggerTypeChoiceRrcstate(rrcStateList *e2smrcv1.TriggerTypeChoiceRrcstate) (*e2smrcv1.TriggerTypeChoice, error) {

	item := &e2smrcv1.TriggerTypeChoice{
		TriggerTypeChoice: &e2smrcv1.TriggerTypeChoice_TriggerTypeChoiceRrcstate{
			TriggerTypeChoiceRrcstate: rrcStateList,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTriggerTypeChoiceRrcstate() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateTriggerTypeChoiceUeID(ueIDchange int32) (*e2smrcv1.TriggerTypeChoice, error) {

	item := &e2smrcv1.TriggerTypeChoice{
		TriggerTypeChoice: &e2smrcv1.TriggerTypeChoice_TriggerTypeChoiceUeid{
			TriggerTypeChoiceUeid: &e2smrcv1.TriggerTypeChoiceUeid{
				UeIdchangeId: ueIDchange,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTriggerTypeChoiceUeID() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateTriggerTypeChoiceL2State(ranParameterTesting *e2smrcv1.RanparameterTesting) (*e2smrcv1.TriggerTypeChoice, error) {

	item := &e2smrcv1.TriggerTypeChoice{
		TriggerTypeChoice: &e2smrcv1.TriggerTypeChoice_TriggerTypeChoiceL2State{
			TriggerTypeChoiceL2State: &e2smrcv1.TriggerTypeChoiceL2State{
				AssociatedL2Variables: ranParameterTesting,
			},
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTriggerTypeChoiceL2State() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateTriggerTypeChoiceRrcstateItem(rrcState e2smrcv1.RrcState) (*e2smrcv1.TriggerTypeChoiceRrcstateItem, error) {

	item := &e2smrcv1.TriggerTypeChoiceRrcstateItem{
		StateChangedTo: rrcState,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTriggerTypeChoiceRrcstateItem() error validating E2SM-RC PDU %s", err.Error())
	}

	return item, nil
}

func CreateOnDemandTrue() e2smrcv1.OnDemand {
	return e2smrcv1.OnDemand_ON_DEMAND_TRUE
}
