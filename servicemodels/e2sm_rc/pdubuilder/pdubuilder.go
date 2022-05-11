// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
    
e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"

e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"

    "github.com/onosproject/onos-lib-go/pkg/errors"
)


func CreateCgi(nRCgi *e2smcommoniesv1.NrCgi, eUtraCgi *e2smcommoniesv1.EutraCgi, ) (*e2smcommoniesv1.Cgi, error) {

    msg := &e2smcommoniesv1.Cgi{
    
    NRCgi: nRCgi,
    
    EUtraCgi: eUtraCgi,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCgi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateCoreCpID(fiveGc *e2smcommoniesv1.Guami, ePc *e2smcommoniesv1.Gummei, ) (*e2smcommoniesv1.CoreCpid, error) {

    msg := &e2smcommoniesv1.CoreCpid{
    
    FiveGc: fiveGc,
    
    EPc: ePc,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCoreCpID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDentifier(nG *e2smcommoniesv1.InterfaceIdNg, xN *e2smcommoniesv1.InterfaceIdXn, f1 *e2smcommoniesv1.InterfaceIdF1, e1 *e2smcommoniesv1.InterfaceIdE1, s1 *e2smcommoniesv1.InterfaceIdS1, x2 *e2smcommoniesv1.InterfaceIdX2, w1 *e2smcommoniesv1.InterfaceIdW1, ) (*e2smcommoniesv1.InterfaceIdentifier, error) {

    msg := &e2smcommoniesv1.InterfaceIdentifier{
    
    NG: nG,
    
    XN: xN,
    
    F1: f1,
    
    E1: e1,
    
    S1: s1,
    
    X2: x2,
    
    W1: w1,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDentifier() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDNg(guami *e2smcommoniesv1.Guami, ) (*e2smcommoniesv1.InterfaceIdNg, error) {

    msg := &e2smcommoniesv1.InterfaceIdNg{
    
    Guami: guami,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDNg() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDXn(globalNgRanId *e2smcommoniesv1.GlobalNgrannodeId, ) (*e2smcommoniesv1.InterfaceIdXn, error) {

    msg := &e2smcommoniesv1.InterfaceIdXn{
    
    GlobalNgRanId: globalNgRanId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDXn() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDF1(globalGnbId *e2smcommoniesv1.GlobalGnbId, gNbDuId *e2smcommoniesv1.GnbDuId, ) (*e2smcommoniesv1.InterfaceIdF1, error) {

    msg := &e2smcommoniesv1.InterfaceIdF1{
    
    GlobalGnbId: globalGnbId,
    
    GNbDuId: gNbDuId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDF1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDE1(globalGnbId *e2smcommoniesv1.GlobalGnbId, gNbCuUpId *e2smcommoniesv1.GnbCuUpId, ) (*e2smcommoniesv1.InterfaceIdE1, error) {

    msg := &e2smcommoniesv1.InterfaceIdE1{
    
    GlobalGnbId: globalGnbId,
    
    GNbCuUpId: gNbCuUpId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDE1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDS1(gUmmei *e2smcommoniesv1.Gummei, ) (*e2smcommoniesv1.InterfaceIdS1, error) {

    msg := &e2smcommoniesv1.InterfaceIdS1{
    
    GUmmei: gUmmei,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDS1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDX2(nodeType *e2smcommoniesv1.NodeType, ) (*e2smcommoniesv1.InterfaceIdX2, error) {

    msg := &e2smcommoniesv1.InterfaceIdX2{
    
    NodeType: nodeType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDX2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNodeType(globalEnbId *e2smcommoniesv1.GlobalEnbId, globalEnGnbId *e2smcommoniesv1.GlobalenGnbId, ) (*e2smcommoniesv1.NodeType, error) {

    msg := &e2smcommoniesv1.NodeType{
    
    GlobalEnbId: globalEnbId,
    
    GlobalEnGnbId: globalEnGnbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNodeType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDW1(globalNgENbId *e2smcommoniesv1.GlobalNgEnbId, ngENbDuId *e2smcommoniesv1.NgenbDuId, ) (*e2smcommoniesv1.InterfaceIdW1, error) {

    msg := &e2smcommoniesv1.InterfaceIdW1{
    
    GlobalNgENbId: globalNgENbId,
    
    NgENbDuId: ngENbDuId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDW1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceMessageID(interfaceProcedureId int32, messageType *e2smcommoniesv1.MessageType, ) (*e2smcommoniesv1.InterfaceMessageId, error) {

    msg := &e2smcommoniesv1.InterfaceMessageId{
    
    InterfaceProcedureId: interfaceProcedureId,
    
    MessageType: messageType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceMessageID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGroupID(fiveGc *e2smcommoniesv1.FiveQi, ePc *e2smcommoniesv1.Qci, ) (*e2smcommoniesv1.GroupId, error) {

    msg := &e2smcommoniesv1.GroupId{
    
    FiveGc: fiveGc,
    
    EPc: ePc,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGroupID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateQoSID(fiveGc *e2smcommoniesv1.FiveQi, ePc *e2smcommoniesv1.Qci, ) (*e2smcommoniesv1.QoSid, error) {

    msg := &e2smcommoniesv1.QoSid{
    
    FiveGc: fiveGc,
    
    EPc: ePc,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateQoSID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionName(ranFunctionShortName string, ranFunctionE2SmOid string, ranFunctionDescription string, ranFunctionInstance int32, ) (*e2smcommoniesv1.RanfunctionName, error) {

    msg := &e2smcommoniesv1.RanfunctionName{
    
    RanFunctionShortName: ranFunctionShortName,
    
    RanFunctionE2SmOid: ranFunctionE2SmOid,
    
    RanFunctionDescription: ranFunctionDescription,
    
    RanFunctionInstance: ranFunctionInstance,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicFormatType(value int32, ) (*e2smcommoniesv1.RicFormatType, error) {

    msg := &e2smcommoniesv1.RicFormatType{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicFormatType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicStyleType(value int32, ) (*e2smcommoniesv1.RicStyleType, error) {

    msg := &e2smcommoniesv1.RicStyleType{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicStyleType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicStyleName(value string, ) (*e2smcommoniesv1.RicStyleName, error) {

    msg := &e2smcommoniesv1.RicStyleName{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicStyleName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRrcMessageID(rrcType *e2smcommoniesv1.RrcType, messageId int64, ) (*e2smcommoniesv1.RrcMessageId, error) {

    msg := &e2smcommoniesv1.RrcMessageId{
    
    RrcType: rrcType,
    
    MessageId: messageId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRrcMessageID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRrcType(lte *e2smcommoniesv1.RrcclassLte, nr *e2smcommoniesv1.RrcclassNr, ) (*e2smcommoniesv1.RrcType, error) {

    msg := &e2smcommoniesv1.RrcType{
    
    Lte: lte,
    
    Nr: nr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRrcType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateServingCellArfcn(nR *e2smcommoniesv1.NrArfcn, eUtra *e2smcommoniesv1.EUtraArfcn, ) (*e2smcommoniesv1.ServingCellArfcn, error) {

    msg := &e2smcommoniesv1.ServingCellArfcn{
    
    NR: nR,
    
    EUtra: eUtra,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellArfcn() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateServingCellPci(nR *e2smcommoniesv1.NrPci, eUtra *e2smcommoniesv1.EUtraPci, ) (*e2smcommoniesv1.ServingCellPci, error) {

    msg := &e2smcommoniesv1.ServingCellPci{
    
    NR: nR,
    
    EUtra: eUtra,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellPci() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeID(gNbUeid *e2smcommoniesv1.UeidGnb, gNbDuUeid *e2smcommoniesv1.UeidGnbDu, gNbCuUpUeid *e2smcommoniesv1.UeidGnbCuUp, ngENbUeid *e2smcommoniesv1.UeidNgEnb, ngENbDuUeid *e2smcommoniesv1.UeidNgEnbDu, enGNbUeid *e2smcommoniesv1.UeidEnGnb, eNbUeid *e2smcommoniesv1.UeidEnb, ) (*e2smcommoniesv1.Ueid, error) {

    msg := &e2smcommoniesv1.Ueid{
    
    GNbUeid: gNbUeid,
    
    GNbDuUeid: gNbDuUeid,
    
    GNbCuUpUeid: gNbCuUpUeid,
    
    NgENbUeid: ngENbUeid,
    
    NgENbDuUeid: ngENbDuUeid,
    
    EnGNbUeid: enGNbUeid,
    
    ENbUeid: eNbUeid,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGnb(amfUeNgapId *e2smcommoniesv1.AmfUeNgapId, guami *e2smcommoniesv1.Guami, gNbCuUeF1ApIdList *e2smcommoniesv1.UeidGnbCuF1ApIdList, gNbCuCpUeE1ApIdList *e2smcommoniesv1.UeidGnbCuCpE1ApIdList, ranUeid *e2smcommoniesv1.Ranueid, mNgRanUeXnApId *e2smcommoniesv1.NgRannodeUexnApid, globalGnbId *e2smcommoniesv1.GlobalGnbId, globalNgRannodeId *e2smcommoniesv1.GlobalNgrannodeId, ) (*e2smcommoniesv1.UeidGnb, error) {

    msg := &e2smcommoniesv1.UeidGnb{
    
    AmfUeNgapId: amfUeNgapId,
    
    Guami: guami,
    
    GNbCuUeF1ApIdList: gNbCuUeF1ApIdList,
    
    GNbCuCpUeE1ApIdList: gNbCuCpUeE1ApIdList,
    
    RanUeid: ranUeid,
    
    MNgRanUeXnApId: mNgRanUeXnApId,
    
    GlobalGnbId: globalGnbId,
    
    GlobalNgRannodeId: globalNgRannodeId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGnb() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGnbCuCpE1ApIDList(value *e2smcommoniesv1.UeidGnbCuCpE1ApIdItem, ) (*e2smcommoniesv1.UeidGnbCuCpE1ApIdList, error) {

    msg := &e2smcommoniesv1.UeidGnbCuCpE1ApIdList{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGnbCuCpE1ApIDList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGnbCuCpE1ApIDItem(gNbCuCpUeE1ApId *e2smcommoniesv1.GnbCuCpUeE1ApId, ) (*e2smcommoniesv1.UeidGnbCuCpE1ApIdItem, error) {

    msg := &e2smcommoniesv1.UeidGnbCuCpE1ApIdItem{
    
    GNbCuCpUeE1ApId: gNbCuCpUeE1ApId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGnbCuCpE1ApIDItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGnbCuF1ApIDList(value *e2smcommoniesv1.UeidGnbCuCpF1ApIdItem, ) (*e2smcommoniesv1.UeidGnbCuF1ApIdList, error) {

    msg := &e2smcommoniesv1.UeidGnbCuF1ApIdList{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGnbCuF1ApIDList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGnbCuCpF1ApIDItem(gNbCuUeF1ApId *e2smcommoniesv1.GnbCuUeF1ApId, ) (*e2smcommoniesv1.UeidGnbCuCpF1ApIdItem, error) {

    msg := &e2smcommoniesv1.UeidGnbCuCpF1ApIdItem{
    
    GNbCuUeF1ApId: gNbCuUeF1ApId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGnbCuCpF1ApIDItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGnbDu(gNbCuUeF1ApId *e2smcommoniesv1.GnbCuUeF1ApId, ranUeid *e2smcommoniesv1.Ranueid, ) (*e2smcommoniesv1.UeidGnbDu, error) {

    msg := &e2smcommoniesv1.UeidGnbDu{
    
    GNbCuUeF1ApId: gNbCuUeF1ApId,
    
    RanUeid: ranUeid,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGnbDu() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGnbCuUp(gNbCuCpUeE1ApId *e2smcommoniesv1.GnbCuCpUeE1ApId, ranUeid *e2smcommoniesv1.Ranueid, ) (*e2smcommoniesv1.UeidGnbCuUp, error) {

    msg := &e2smcommoniesv1.UeidGnbCuUp{
    
    GNbCuCpUeE1ApId: gNbCuCpUeE1ApId,
    
    RanUeid: ranUeid,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGnbCuUp() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDNgEnb(amfUeNgapId *e2smcommoniesv1.AmfUeNgapId, guami *e2smcommoniesv1.Guami, ngENbCuUeW1ApId *e2smcommoniesv1.NgenbCuUeW1ApId, mNgRanUeXnApId *e2smcommoniesv1.NgRannodeUexnApid, globalNgEnbId *e2smcommoniesv1.GlobalNgEnbId, globalNgRannodeId *e2smcommoniesv1.GlobalNgrannodeId, ) (*e2smcommoniesv1.UeidNgEnb, error) {

    msg := &e2smcommoniesv1.UeidNgEnb{
    
    AmfUeNgapId: amfUeNgapId,
    
    Guami: guami,
    
    NgENbCuUeW1ApId: ngENbCuUeW1ApId,
    
    MNgRanUeXnApId: mNgRanUeXnApId,
    
    GlobalNgEnbId: globalNgEnbId,
    
    GlobalNgRannodeId: globalNgRannodeId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDNgEnb() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDNgEnbDu(ngENbCuUeW1ApId *e2smcommoniesv1.NgenbCuUeW1ApId, ) (*e2smcommoniesv1.UeidNgEnbDu, error) {

    msg := &e2smcommoniesv1.UeidNgEnbDu{
    
    NgENbCuUeW1ApId: ngENbCuUeW1ApId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDNgEnbDu() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDEnGnb(mENbUeX2ApId *e2smcommoniesv1.EnbUeX2ApId, mENbUeX2ApIdExtension *e2smcommoniesv1.EnbUeX2ApIdExtension, globalEnbId *e2smcommoniesv1.GlobalEnbId, gNbCuUeF1ApId *e2smcommoniesv1.GnbCuUeF1ApId, gNbCuCpUeE1ApIdList *e2smcommoniesv1.UeidGnbCuCpE1ApIdList, ranUeid *e2smcommoniesv1.Ranueid, ) (*e2smcommoniesv1.UeidEnGnb, error) {

    msg := &e2smcommoniesv1.UeidEnGnb{
    
    MENbUeX2ApId: mENbUeX2ApId,
    
    MENbUeX2ApIdExtension: mENbUeX2ApIdExtension,
    
    GlobalEnbId: globalEnbId,
    
    GNbCuUeF1ApId: gNbCuUeF1ApId,
    
    GNbCuCpUeE1ApIdList: gNbCuCpUeE1ApIdList,
    
    RanUeid: ranUeid,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDEnGnb() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDEnb(mMeUeS1ApId *e2smcommoniesv1.MmeUeS1ApId, gUmmei *e2smcommoniesv1.Gummei, mENbUeX2ApId *e2smcommoniesv1.EnbUeX2ApId, mENbUeX2ApIdExtension *e2smcommoniesv1.EnbUeX2ApIdExtension, globalEnbId *e2smcommoniesv1.GlobalEnbId, ) (*e2smcommoniesv1.UeidEnb, error) {

    msg := &e2smcommoniesv1.UeidEnb{
    
    MMeUeS1ApId: mMeUeS1ApId,
    
    GUmmei: gUmmei,
    
    MENbUeX2ApId: mENbUeX2ApId,
    
    MENbUeX2ApIdExtension: mENbUeX2ApIdExtension,
    
    GlobalEnbId: globalEnbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDEnb() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEnbID(macroENbId *e2smcommoniesv1.BitString, homeENbId *e2smcommoniesv1.BitString, shortMacroENbId *e2smcommoniesv1.BitString, longMacroENbId *e2smcommoniesv1.BitString, ) (*e2smcommoniesv1.EnbId, error) {

    msg := &e2smcommoniesv1.EnbId{
    
    MacroENbId: macroENbId,
    
    HomeENbId: homeENbId,
    
    ShortMacroENbId: shortMacroENbId,
    
    LongMacroENbId: longMacroENbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGlobalEnbID(pLmnidentity *e2smcommoniesv1.Plmnidentity, eNbId *e2smcommoniesv1.EnbId, ) (*e2smcommoniesv1.GlobalEnbId, error) {

    msg := &e2smcommoniesv1.GlobalEnbId{
    
    PLmnidentity: pLmnidentity,
    
    ENbId: eNbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalEnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGummei(pLmnIdentity *e2smcommoniesv1.Plmnidentity, mMeGroupId *e2smcommoniesv1.MmeGroupId, mMeCode *e2smcommoniesv1.MmeCode, ) (*e2smcommoniesv1.Gummei, error) {

    msg := &e2smcommoniesv1.Gummei{
    
    PLmnIdentity: pLmnIdentity,
    
    MMeGroupId: mMeGroupId,
    
    MMeCode: mMeCode,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGummei() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateMmeUeS1ApID(value int64, ) (*e2smcommoniesv1.MmeUeS1ApId, error) {

    msg := &e2smcommoniesv1.MmeUeS1ApId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMmeUeS1ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateQci(value int32, ) (*e2smcommoniesv1.Qci, error) {

    msg := &e2smcommoniesv1.Qci{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateQci() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSubscriberProfileIDforRfp(value int32, ) (*e2smcommoniesv1.SubscriberProfileIdforRfp, error) {

    msg := &e2smcommoniesv1.SubscriberProfileIdforRfp{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSubscriberProfileIDforRfp() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEnGnbID(enGNbId *e2smcommoniesv1.BitString, ) (*e2smcommoniesv1.EnGnbId, error) {

    msg := &e2smcommoniesv1.EnGnbId{
    
    EnGNbId: enGNbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnGnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEnbUeX2ApID(value int32, ) (*e2smcommoniesv1.EnbUeX2ApId, error) {

    msg := &e2smcommoniesv1.EnbUeX2ApId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbUeX2ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEnbUeX2ApIDExtension(value int32, ) (*e2smcommoniesv1.EnbUeX2ApIdExtension, error) {

    msg := &e2smcommoniesv1.EnbUeX2ApIdExtension{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbUeX2ApIDExtension() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEUtraArfcn(value int32, ) (*e2smcommoniesv1.EUtraArfcn, error) {

    msg := &e2smcommoniesv1.EUtraArfcn{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEUtraArfcn() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEUtraPci(value int32, ) (*e2smcommoniesv1.EUtraPci, error) {

    msg := &e2smcommoniesv1.EUtraPci{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEUtraPci() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEUtraTac(value []byte, ) (*e2smcommoniesv1.EUtraTac, error) {

    msg := &e2smcommoniesv1.EUtraTac{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEUtraTac() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGlobalenGnbID(pLmnIdentity *e2smcommoniesv1.Plmnidentity, enGNbId *e2smcommoniesv1.EnGnbId, ) (*e2smcommoniesv1.GlobalenGnbId, error) {

    msg := &e2smcommoniesv1.GlobalenGnbId{
    
    PLmnIdentity: pLmnIdentity,
    
    EnGNbId: enGNbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalenGnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNgenbCuUeW1ApID(value int64, ) (*e2smcommoniesv1.NgenbCuUeW1ApId, error) {

    msg := &e2smcommoniesv1.NgenbCuUeW1ApId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgenbCuUeW1ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNgenbDuID(value int64, ) (*e2smcommoniesv1.NgenbDuId, error) {

    msg := &e2smcommoniesv1.NgenbDuId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgenbDuID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAmfpointer(value *e2smcommoniesv1.BitString, ) (*e2smcommoniesv1.Amfpointer, error) {

    msg := &e2smcommoniesv1.Amfpointer{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAmfpointer() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAmfregionID(value *e2smcommoniesv1.BitString, ) (*e2smcommoniesv1.AmfregionId, error) {

    msg := &e2smcommoniesv1.AmfregionId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAmfregionID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAmfsetID(value *e2smcommoniesv1.BitString, ) (*e2smcommoniesv1.AmfsetId, error) {

    msg := &e2smcommoniesv1.AmfsetId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAmfsetID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAmfUeNgapID(value int64, ) (*e2smcommoniesv1.AmfUeNgapId, error) {

    msg := &e2smcommoniesv1.AmfUeNgapId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAmfUeNgapID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEutracellIDentity(value *e2smcommoniesv1.BitString, ) (*e2smcommoniesv1.EutracellIdentity, error) {

    msg := &e2smcommoniesv1.EutracellIdentity{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEutracellIDentity() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEutraCgi(pLmnidentity *e2smcommoniesv1.Plmnidentity, eUtracellIdentity *e2smcommoniesv1.EutracellIdentity, ) (*e2smcommoniesv1.EutraCgi, error) {

    msg := &e2smcommoniesv1.EutraCgi{
    
    PLmnidentity: pLmnidentity,
    
    EUtracellIdentity: eUtracellIdentity,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEutraCgi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateFiveQi(value int32, ) (*e2smcommoniesv1.FiveQi, error) {

    msg := &e2smcommoniesv1.FiveQi{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateFiveQi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGlobalGnbID(pLmnidentity *e2smcommoniesv1.Plmnidentity, gNbId *e2smcommoniesv1.GnbId, ) (*e2smcommoniesv1.GlobalGnbId, error) {

    msg := &e2smcommoniesv1.GlobalGnbId{
    
    PLmnidentity: pLmnidentity,
    
    GNbId: gNbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalGnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGlobalNgEnbID(pLmnidentity *e2smcommoniesv1.Plmnidentity, ngEnbId *e2smcommoniesv1.NgEnbId, ) (*e2smcommoniesv1.GlobalNgEnbId, error) {

    msg := &e2smcommoniesv1.GlobalNgEnbId{
    
    PLmnidentity: pLmnidentity,
    
    NgEnbId: ngEnbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalNgEnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbID(gNbId *e2smcommoniesv1.BitString, ) (*e2smcommoniesv1.GnbId, error) {

    msg := &e2smcommoniesv1.GnbId{
    
    GNbId: gNbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGuami(pLmnidentity *e2smcommoniesv1.Plmnidentity, aMfregionId *e2smcommoniesv1.AmfregionId, aMfsetId *e2smcommoniesv1.AmfsetId, aMfpointer *e2smcommoniesv1.Amfpointer, ) (*e2smcommoniesv1.Guami, error) {

    msg := &e2smcommoniesv1.Guami{
    
    PLmnidentity: pLmnidentity,
    
    AMfregionId: aMfregionId,
    
    AMfsetId: aMfsetId,
    
    AMfpointer: aMfpointer,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGuami() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateIndexToRfsp(value int32, ) (*e2smcommoniesv1.IndexToRfsp, error) {

    msg := &e2smcommoniesv1.IndexToRfsp{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateIndexToRfsp() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNgEnbID(macroNgEnbId *e2smcommoniesv1.BitString, shortMacroNgEnbId *e2smcommoniesv1.BitString, longMacroNgEnbId *e2smcommoniesv1.BitString, ) (*e2smcommoniesv1.NgEnbId, error) {

    msg := &e2smcommoniesv1.NgEnbId{
    
    MacroNgEnbId: macroNgEnbId,
    
    ShortMacroNgEnbId: shortMacroNgEnbId,
    
    LongMacroNgEnbId: longMacroNgEnbId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgEnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrcellIDentity(value *e2smcommoniesv1.BitString, ) (*e2smcommoniesv1.NrcellIdentity, error) {

    msg := &e2smcommoniesv1.NrcellIdentity{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrcellIDentity() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrCgi(pLmnidentity *e2smcommoniesv1.Plmnidentity, nRcellIdentity *e2smcommoniesv1.NrcellIdentity, ) (*e2smcommoniesv1.NrCgi, error) {

    msg := &e2smcommoniesv1.NrCgi{
    
    PLmnidentity: pLmnidentity,
    
    NRcellIdentity: nRcellIdentity,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrCgi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateQosFlowIDentifier(value int32, ) (*e2smcommoniesv1.QosFlowIdentifier, error) {

    msg := &e2smcommoniesv1.QosFlowIdentifier{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateQosFlowIDentifier() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSd(value []byte, ) (*e2smcommoniesv1.Sd, error) {

    msg := &e2smcommoniesv1.Sd{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSd() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSNssai(sSt *e2smcommoniesv1.Sst, sD *e2smcommoniesv1.Sd, ) (*e2smcommoniesv1.SNssai, error) {

    msg := &e2smcommoniesv1.SNssai{
    
    SSt: sSt,
    
    SD: sD,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSNssai() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSst(value []byte, ) (*e2smcommoniesv1.Sst, error) {

    msg := &e2smcommoniesv1.Sst{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSst() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNgRannodeUexnApID(value int64, ) (*e2smcommoniesv1.NgRannodeUexnApid, error) {

    msg := &e2smcommoniesv1.NgRannodeUexnApid{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgRannodeUexnApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGlobalNgrannodeID(gNb *e2smcommoniesv1.GlobalGnbId, ngENb *e2smcommoniesv1.GlobalNgEnbId, ) (*e2smcommoniesv1.GlobalNgrannodeId, error) {

    msg := &e2smcommoniesv1.GlobalNgrannodeId{
    
    GNb: gNb,
    
    NgENb: ngENb,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalNgrannodeID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbCuCpUeE1ApID(value int64, ) (*e2smcommoniesv1.GnbCuCpUeE1ApId, error) {

    msg := &e2smcommoniesv1.GnbCuCpUeE1ApId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbCuCpUeE1ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbCuUpID(value int64, ) (*e2smcommoniesv1.GnbCuUpId, error) {

    msg := &e2smcommoniesv1.GnbCuUpId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbCuUpID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateFiveGsTac(value []byte, ) (*e2smcommoniesv1.FiveGsTac, error) {

    msg := &e2smcommoniesv1.FiveGsTac{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateFiveGsTac() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateFreqBandNrItem(freqBandIndicatorNr int32, ) (*e2smcommoniesv1.FreqBandNrItem, error) {

    msg := &e2smcommoniesv1.FreqBandNrItem{
    
    FreqBandIndicatorNr: freqBandIndicatorNr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateFreqBandNrItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbCuUeF1ApID(value int64, ) (*e2smcommoniesv1.GnbCuUeF1ApId, error) {

    msg := &e2smcommoniesv1.GnbCuUeF1ApId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbCuUeF1ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbDuID(value int64, ) (*e2smcommoniesv1.GnbDuId, error) {

    msg := &e2smcommoniesv1.GnbDuId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbDuID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrPci(value int32, ) (*e2smcommoniesv1.NrPci, error) {

    msg := &e2smcommoniesv1.NrPci{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrPci() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrArfcn(nRarfcn int32, ) (*e2smcommoniesv1.NrArfcn, error) {

    msg := &e2smcommoniesv1.NrArfcn{
    
    NRarfcn: nRarfcn,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrArfcn() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrfrequencyBandList(value *e2smcommoniesv1.NrfrequencyBandItem, ) (*e2smcommoniesv1.NrfrequencyBandList, error) {

    msg := &e2smcommoniesv1.NrfrequencyBandList{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrfrequencyBandList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrfrequencyBandItem(freqBandIndicatorNr int32, supportedSulbandList *e2smcommoniesv1.SupportedSulbandList, ) (*e2smcommoniesv1.NrfrequencyBandItem, error) {

    msg := &e2smcommoniesv1.NrfrequencyBandItem{
    
    FreqBandIndicatorNr: freqBandIndicatorNr,
    
    SupportedSulbandList: supportedSulbandList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrfrequencyBandItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrfrequencyInfo(nrArfcn *e2smcommoniesv1.NrArfcn, frequencyBandList *e2smcommoniesv1.NrfrequencyBandList, frequencyShift7P5Khz *e2smcommoniesv1.NrfrequencyShift7P5Khz, ) (*e2smcommoniesv1.NrfrequencyInfo, error) {

    msg := &e2smcommoniesv1.NrfrequencyInfo{
    
    NrArfcn: nrArfcn,
    
    FrequencyBandList: frequencyBandList,
    
    FrequencyShift7P5Khz: frequencyShift7P5Khz,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrfrequencyInfo() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSupportedSulbandList(value *e2smcommoniesv1.SupportedSulfreqBandItem, ) (*e2smcommoniesv1.SupportedSulbandList, error) {

    msg := &e2smcommoniesv1.SupportedSulbandList{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSupportedSulbandList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSupportedSulfreqBandItem(freqBandIndicatorNr int32, ) (*e2smcommoniesv1.SupportedSulfreqBandItem, error) {

    msg := &e2smcommoniesv1.SupportedSulfreqBandItem{
    
    FreqBandIndicatorNr: freqBandIndicatorNr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSupportedSulfreqBandItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanueID(value []byte, ) (*e2smcommoniesv1.Ranueid, error) {

    msg := &e2smcommoniesv1.Ranueid{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanueID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreatePlmnIDentity(value []byte, ) (*e2smcommoniesv1.Plmnidentity, error) {

    msg := &e2smcommoniesv1.Plmnidentity{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreatePlmnIDentity() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateMmeGroupID(value []byte, ) (*e2smcommoniesv1.MmeGroupId, error) {

    msg := &e2smcommoniesv1.MmeGroupId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMmeGroupID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateMmeCode(value []byte, ) (*e2smcommoniesv1.MmeCode, error) {

    msg := &e2smcommoniesv1.MmeCode{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMmeCode() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNeighborCellList(value *e2smrcv1.NeighborCellItem, ) (*e2smrcv1.NeighborCellList, error) {

    msg := &e2smrcv1.NeighborCellList{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNeighborCellList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNeighborCellItem(ranTypeChoiceNr *e2smrcv1.NeighborCellItemChoiceNr, ranTypeChoiceEutra *e2smrcv1.NeighborCellItemChoiceEUtra, ) (*e2smrcv1.NeighborCellItem, error) {

    msg := &e2smrcv1.NeighborCellItem{
    
    RanTypeChoiceNr: ranTypeChoiceNr,
    
    RanTypeChoiceEutra: ranTypeChoiceEutra,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNeighborCellItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNeighborCellItemChoiceNr(nRCgi *e2smcommoniesv1.NrCgi, nRPci *e2smcommoniesv1.NrPci, fiveGsTac *e2smcommoniesv1.FiveGsTac, nRModeInfo *e2smrcv1.NRModeInfo, nRFreqInfo *e2smcommoniesv1.NrfrequencyInfo, x2XnEstablished *e2smrcv1.X2XNEstablished, hOValidated *e2smrcv1.HOValidated, version int32, ) (*e2smrcv1.NeighborCellItemChoiceNr, error) {

    msg := &e2smrcv1.NeighborCellItemChoiceNr{
    
    NRCgi: nRCgi,
    
    NRPci: nRPci,
    
    FiveGsTac: fiveGsTac,
    
    NRModeInfo: nRModeInfo,
    
    NRFreqInfo: nRFreqInfo,
    
    X2XnEstablished: x2XnEstablished,
    
    HOValidated: hOValidated,
    
    Version: version,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNeighborCellItemChoiceNr() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNeighborCellItemChoiceEUtra(eUtraCgi *e2smcommoniesv1.EutraCgi, eUtraPci *e2smcommoniesv1.EUtraPci, eUtraArfcn *e2smcommoniesv1.EUtraArfcn, eUtraTac *e2smcommoniesv1.EUtraTac, x2XnEstablished *e2smrcv1.X2XNEstablished, hOValidated *e2smrcv1.HOValidated, version int32, ) (*e2smrcv1.NeighborCellItemChoiceEUtra, error) {

    msg := &e2smrcv1.NeighborCellItemChoiceEUtra{
    
    EUtraCgi: eUtraCgi,
    
    EUtraPci: eUtraPci,
    
    EUtraArfcn: eUtraArfcn,
    
    EUtraTac: eUtraTac,
    
    X2XnEstablished: x2XnEstablished,
    
    HOValidated: hOValidated,
    
    Version: version,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNeighborCellItemChoiceEUtra() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNeighborRelationInfo(servingCellPci *e2smcommoniesv1.ServingCellPci, servingCellArfcn *e2smcommoniesv1.ServingCellArfcn, neighborCellList *e2smrcv1.NeighborCellList, ) (*e2smrcv1.NeighborRelationInfo, error) {

    msg := &e2smrcv1.NeighborRelationInfo{
    
    ServingCellPci: servingCellPci,
    
    ServingCellArfcn: servingCellArfcn,
    
    NeighborCellList: neighborCellList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNeighborRelationInfo() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerCellInfo(cellInfoList *e2smrcv1.EventTriggerCellInfoItem, ) (*e2smrcv1.EventTriggerCellInfo, error) {

    msg := &e2smrcv1.EventTriggerCellInfo{
    
    CellInfoList: cellInfoList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerCellInfo() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerCellInfoItem(eventTriggerCellId *e2smrcv1.RicEventTriggerCellId, cellType *e2smrcv1.CellType, logicalOr *e2smrcv1.LogicalOr, ) (*e2smrcv1.EventTriggerCellInfoItem, error) {

    msg := &e2smrcv1.EventTriggerCellInfoItem{
    
    EventTriggerCellId: eventTriggerCellId,
    
    CellType: cellType,
    
    LogicalOr: logicalOr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerCellInfoItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateCellType(cellTypeChoiceIndividual *e2smrcv1.EventTriggerCellInfoItemChoiceIndividual, cellTypeChoiceGroup *e2smrcv1.EventTriggerCellInfoItemChoiceGroup, ) (*e2smrcv1.CellType, error) {

    msg := &e2smrcv1.CellType{
    
    CellTypeChoiceIndividual: cellTypeChoiceIndividual,
    
    CellTypeChoiceGroup: cellTypeChoiceGroup,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerCellInfoItemChoiceIndivIDual(cellGlobalId *e2smcommoniesv1.Cgi, ) (*e2smrcv1.EventTriggerCellInfoItemChoiceIndividual, error) {

    msg := &e2smrcv1.EventTriggerCellInfoItemChoiceIndividual{
    
    CellGlobalId: cellGlobalId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerCellInfoItemChoiceIndivIDual() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerCellInfoItemChoiceGroup(ranParameterTesting *e2smrcv1.RanparameterTesting, ) (*e2smrcv1.EventTriggerCellInfoItemChoiceGroup, error) {

    msg := &e2smrcv1.EventTriggerCellInfoItemChoiceGroup{
    
    RanParameterTesting: ranParameterTesting,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerCellInfoItemChoiceGroup() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerUeInfo(ueInfoList *e2smrcv1.EventTriggerUeInfoItem, ) (*e2smrcv1.EventTriggerUeInfo, error) {

    msg := &e2smrcv1.EventTriggerUeInfo{
    
    UeInfoList: ueInfoList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerUeInfo() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerUeInfoItem(eventTriggerUeid *e2smrcv1.RicEventTriggerUeId, ueType *e2smrcv1.UeType, logicalOr *e2smrcv1.LogicalOr, ) (*e2smrcv1.EventTriggerUeInfoItem, error) {

    msg := &e2smrcv1.EventTriggerUeInfoItem{
    
    EventTriggerUeid: eventTriggerUeid,
    
    UeType: ueType,
    
    LogicalOr: logicalOr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerUeInfoItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeType(ueTypeChoiceIndividual *e2smrcv1.EventTriggerUeInfoItemChoiceIndividual, ueTypeChoiceGroup *e2smrcv1.EventTriggerUeInfoItemChoiceGroup, ) (*e2smrcv1.UeType, error) {

    msg := &e2smrcv1.UeType{
    
    UeTypeChoiceIndividual: ueTypeChoiceIndividual,
    
    UeTypeChoiceGroup: ueTypeChoiceGroup,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerUeInfoItemChoiceIndivIDual(ueId *e2smcommoniesv1.Ueid, ranParameterTesting *e2smrcv1.RanparameterTesting, ) (*e2smrcv1.EventTriggerUeInfoItemChoiceIndividual, error) {

    msg := &e2smrcv1.EventTriggerUeInfoItemChoiceIndividual{
    
    UeId: ueId,
    
    RanParameterTesting: ranParameterTesting,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerUeInfoItemChoiceIndivIDual() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerUeInfoItemChoiceGroup(ranParameterTesting *e2smrcv1.RanparameterTesting, ) (*e2smrcv1.EventTriggerUeInfoItemChoiceGroup, error) {

    msg := &e2smrcv1.EventTriggerUeInfoItemChoiceGroup{
    
    RanParameterTesting: ranParameterTesting,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerUeInfoItemChoiceGroup() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerUeeventInfo(ueEventList *e2smrcv1.EventTriggerUeeventInfoItem, ) (*e2smrcv1.EventTriggerUeeventInfo, error) {

    msg := &e2smrcv1.EventTriggerUeeventInfo{
    
    UeEventList: ueEventList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerUeeventInfo() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerUeeventInfoItem(ueEventId *e2smrcv1.RicEventTriggerUeeventId, logicalOr *e2smrcv1.LogicalOr, ) (*e2smrcv1.EventTriggerUeeventInfoItem, error) {

    msg := &e2smrcv1.EventTriggerUeeventInfoItem{
    
    UeEventId: ueEventId,
    
    LogicalOr: logicalOr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerUeeventInfoItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterID(value int64, ) (*e2smrcv1.RanparameterId, error) {

    msg := &e2smrcv1.RanparameterId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterName(value string, ) (*e2smrcv1.RanparameterName, error) {

    msg := &e2smrcv1.RanparameterName{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterDefinition(ranParameterDefinitionChoice *e2smrcv1.RanparameterDefinitionChoice, ) (*e2smrcv1.RanparameterDefinition, error) {

    msg := &e2smrcv1.RanparameterDefinition{
    
    RanParameterDefinitionChoice: ranParameterDefinitionChoice,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterDefinition() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterDefinitionChoice(choiceList *e2smrcv1.RanparameterDefinitionChoiceList, choiceStructure *e2smrcv1.RanparameterDefinitionChoiceStructure, ) (*e2smrcv1.RanparameterDefinitionChoice, error) {

    msg := &e2smrcv1.RanparameterDefinitionChoice{
    
    ChoiceList: choiceList,
    
    ChoiceStructure: choiceStructure,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterDefinitionChoice() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterDefinitionChoiceList(ranParameterList *e2smrcv1.RanparameterDefinitionChoiceListItem, ) (*e2smrcv1.RanparameterDefinitionChoiceList, error) {

    msg := &e2smrcv1.RanparameterDefinitionChoiceList{
    
    RanParameterList: ranParameterList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterDefinitionChoiceList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterDefinitionChoiceListItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.RanparameterDefinitionChoiceListItem, error) {

    msg := &e2smrcv1.RanparameterDefinitionChoiceListItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterDefinitionChoiceListItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterDefinitionChoiceStructure(ranParameterStructure *e2smrcv1.RanparameterDefinitionChoiceStructureItem, ) (*e2smrcv1.RanparameterDefinitionChoiceStructure, error) {

    msg := &e2smrcv1.RanparameterDefinitionChoiceStructure{
    
    RanParameterStructure: ranParameterStructure,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterDefinitionChoiceStructure() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterDefinitionChoiceStructureItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.RanparameterDefinitionChoiceStructureItem, error) {

    msg := &e2smrcv1.RanparameterDefinitionChoiceStructureItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterDefinitionChoiceStructureItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterValue(valueBoolean bool, valueInt int64, valueReal *e2smrcv1.ErrorInParsing:float, valueBitS *e2smrcv1.BitString, valueOctS []byte, valuePrintableString string, ) (*e2smrcv1.RanparameterValue, error) {

    msg := &e2smrcv1.RanparameterValue{
    
    ValueBoolean: valueBoolean,
    
    ValueInt: valueInt,
    
    ValueReal: valueReal,
    
    ValueBitS: valueBitS,
    
    ValueOctS: valueOctS,
    
    ValuePrintableString: valuePrintableString,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterValue() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterValueType(ranPChoiceElementTrue *e2smrcv1.RanparameterValueTypeChoiceElementTrue, ranPChoiceElementFalse *e2smrcv1.RanparameterValueTypeChoiceElementFalse, ranPChoiceStructure *e2smrcv1.RanparameterValueTypeChoiceStructure, ranPChoiceList *e2smrcv1.RanparameterValueTypeChoiceList, ) (*e2smrcv1.RanparameterValueType, error) {

    msg := &e2smrcv1.RanparameterValueType{
    
    RanPChoiceElementTrue: ranPChoiceElementTrue,
    
    RanPChoiceElementFalse: ranPChoiceElementFalse,
    
    RanPChoiceStructure: ranPChoiceStructure,
    
    RanPChoiceList: ranPChoiceList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterValueType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterValueTypeChoiceElementTrue(ranParameterValue *e2smrcv1.RanparameterValue, ) (*e2smrcv1.RanparameterValueTypeChoiceElementTrue, error) {

    msg := &e2smrcv1.RanparameterValueTypeChoiceElementTrue{
    
    RanParameterValue: ranParameterValue,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterValueTypeChoiceElementTrue() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterValueTypeChoiceElementFalse(ranParameterValue *e2smrcv1.RanparameterValue, ) (*e2smrcv1.RanparameterValueTypeChoiceElementFalse, error) {

    msg := &e2smrcv1.RanparameterValueTypeChoiceElementFalse{
    
    RanParameterValue: ranParameterValue,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterValueTypeChoiceElementFalse() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterValueTypeChoiceStructure(ranParameterStructure *e2smrcv1.RanparameterStructure, ) (*e2smrcv1.RanparameterValueTypeChoiceStructure, error) {

    msg := &e2smrcv1.RanparameterValueTypeChoiceStructure{
    
    RanParameterStructure: ranParameterStructure,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterValueTypeChoiceStructure() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterValueTypeChoiceList(ranParameterList *e2smrcv1.RanparameterList, ) (*e2smrcv1.RanparameterValueTypeChoiceList, error) {

    msg := &e2smrcv1.RanparameterValueTypeChoiceList{
    
    RanParameterList: ranParameterList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterValueTypeChoiceList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterStructure(sequenceOfRanParameters *e2smrcv1.RanparameterStructureItem, ) (*e2smrcv1.RanparameterStructure, error) {

    msg := &e2smrcv1.RanparameterStructure{
    
    SequenceOfRanParameters: sequenceOfRanParameters,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterStructure() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterStructureItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterValueType *e2smrcv1.RanparameterValueType, ) (*e2smrcv1.RanparameterStructureItem, error) {

    msg := &e2smrcv1.RanparameterStructureItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterValueType: ranParameterValueType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterStructureItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterList(listOfRanParameter *e2smrcv1.RanparameterStructure, ) (*e2smrcv1.RanparameterList, error) {

    msg := &e2smrcv1.RanparameterList{
    
    ListOfRanParameter: listOfRanParameter,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterTesting(value *e2smrcv1.RanparameterTestingItem, ) (*e2smrcv1.RanparameterTesting, error) {

    msg := &e2smrcv1.RanparameterTesting{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterTesting() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterTestingCondition(ranPChoiceComparison *e2smrcv1.RanPChoiceComparison, ranPChoicePresence *e2smrcv1.RanPChoicePresence, ) (*e2smrcv1.RanparameterTestingCondition, error) {

    msg := &e2smrcv1.RanparameterTestingCondition{
    
    RanPChoiceComparison: ranPChoiceComparison,
    
    RanPChoicePresence: ranPChoicePresence,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterTestingCondition() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterTestingItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterType *e2smrcv1.RanParameterType, ) (*e2smrcv1.RanparameterTestingItem, error) {

    msg := &e2smrcv1.RanparameterTestingItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterType: ranParameterType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterTestingItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanParameterType(ranPChoiceList *e2smrcv1.RanparameterTestingItemChoiceList, ranPChoiceStructure *e2smrcv1.RanparameterTestingItemChoiceStructure, ranPChoiceElementTrue *e2smrcv1.RanparameterTestingItemChoiceElementTrue, ranPChoiceElementFalse *e2smrcv1.RanparameterTestingItemChoiceElementFalse, ) (*e2smrcv1.RanParameterType, error) {

    msg := &e2smrcv1.RanParameterType{
    
    RanPChoiceList: ranPChoiceList,
    
    RanPChoiceStructure: ranPChoiceStructure,
    
    RanPChoiceElementTrue: ranPChoiceElementTrue,
    
    RanPChoiceElementFalse: ranPChoiceElementFalse,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanParameterType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterTestingItemChoiceList(ranParameterList *e2smrcv1.RanparameterTestingList, ) (*e2smrcv1.RanparameterTestingItemChoiceList, error) {

    msg := &e2smrcv1.RanparameterTestingItemChoiceList{
    
    RanParameterList: ranParameterList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterTestingItemChoiceList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterTestingItemChoiceStructure(ranParameterStructure *e2smrcv1.RanparameterTestingStructure, ) (*e2smrcv1.RanparameterTestingItemChoiceStructure, error) {

    msg := &e2smrcv1.RanparameterTestingItemChoiceStructure{
    
    RanParameterStructure: ranParameterStructure,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterTestingItemChoiceStructure() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterTestingItemChoiceElementTrue(ranParameterValue *e2smrcv1.RanparameterValue, ) (*e2smrcv1.RanparameterTestingItemChoiceElementTrue, error) {

    msg := &e2smrcv1.RanparameterTestingItemChoiceElementTrue{
    
    RanParameterValue: ranParameterValue,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterTestingItemChoiceElementTrue() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterTestingItemChoiceElementFalse(ranParameterTestCondition *e2smrcv1.RanparameterTestingCondition, ranParameterValue *e2smrcv1.RanparameterValue, logicalOr *e2smrcv1.LogicalOr, ) (*e2smrcv1.RanparameterTestingItemChoiceElementFalse, error) {

    msg := &e2smrcv1.RanparameterTestingItemChoiceElementFalse{
    
    RanParameterTestCondition: ranParameterTestCondition,
    
    RanParameterValue: ranParameterValue,
    
    LogicalOr: logicalOr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterTestingItemChoiceElementFalse() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterTestingList(value *e2smrcv1.RanparameterTestingItem, ) (*e2smrcv1.RanparameterTestingList, error) {

    msg := &e2smrcv1.RanparameterTestingList{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterTestingList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanparameterTestingStructure(value *e2smrcv1.RanparameterTestingItem, ) (*e2smrcv1.RanparameterTestingStructure, error) {

    msg := &e2smrcv1.RanparameterTestingStructure{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanparameterTestingStructure() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanCallProcessID(value int32, ) (*e2smrcv1.RanCallProcessId, error) {

    msg := &e2smrcv1.RanCallProcessId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanCallProcessID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicCallProcessTypeID(value int32, ) (*e2smrcv1.RicCallProcessTypeId, error) {

    msg := &e2smrcv1.RicCallProcessTypeId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicCallProcessTypeID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicCallProcessTypeName(value string, ) (*e2smrcv1.RicCallProcessTypeName, error) {

    msg := &e2smrcv1.RicCallProcessTypeName{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicCallProcessTypeName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicCallProcessBreakpointID(value int32, ) (*e2smrcv1.RicCallProcessBreakpointId, error) {

    msg := &e2smrcv1.RicCallProcessBreakpointId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicCallProcessBreakpointID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicCallProcessBreakpointName(value string, ) (*e2smrcv1.RicCallProcessBreakpointName, error) {

    msg := &e2smrcv1.RicCallProcessBreakpointName{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicCallProcessBreakpointName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicControlActionID(value int32, ) (*e2smrcv1.RicControlActionId, error) {

    msg := &e2smrcv1.RicControlActionId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicControlActionID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicControlActionName(value string, ) (*e2smrcv1.RicControlActionName, error) {

    msg := &e2smrcv1.RicControlActionName{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicControlActionName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicEventTriggerConditionID(value int32, ) (*e2smrcv1.RicEventTriggerConditionId, error) {

    msg := &e2smrcv1.RicEventTriggerConditionId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicEventTriggerConditionID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicEventTriggerUeID(value int32, ) (*e2smrcv1.RicEventTriggerUeId, error) {

    msg := &e2smrcv1.RicEventTriggerUeId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicEventTriggerUeID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicEventTriggerUeeventID(value int32, ) (*e2smrcv1.RicEventTriggerUeeventId, error) {

    msg := &e2smrcv1.RicEventTriggerUeeventId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicEventTriggerUeeventID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicEventTriggerCellID(value int32, ) (*e2smrcv1.RicEventTriggerCellId, error) {

    msg := &e2smrcv1.RicEventTriggerCellId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicEventTriggerCellID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicInsertIndicationID(value int32, ) (*e2smrcv1.RicInsertIndicationId, error) {

    msg := &e2smrcv1.RicInsertIndicationId{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicInsertIndicationID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicInsertIndicationName(value string, ) (*e2smrcv1.RicInsertIndicationName, error) {

    msg := &e2smrcv1.RicInsertIndicationName{
    
    Value: value,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicInsertIndicationName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicPolicyAction(ricPolicyActionId *e2smrcv1.RicControlActionId, ranParametersList *e2smrcv1.RicPolicyActionRanparameterItem, ) (*e2smrcv1.RicPolicyAction, error) {

    msg := &e2smrcv1.RicPolicyAction{
    
    RicPolicyActionId: ricPolicyActionId,
    
    RanParametersList: ranParametersList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicPolicyAction() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicPolicyActionRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterValueType *e2smrcv1.RanparameterValueType, ) (*e2smrcv1.RicPolicyActionRanparameterItem, error) {

    msg := &e2smrcv1.RicPolicyActionRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterValueType: ranParameterValueType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicPolicyActionRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcEventTrigger(ricEventTriggerFormats *e2smrcv1.RicEventTriggerFormats, ) (*e2smrcv1.E2SmRcEventTrigger, error) {

    msg := &e2smrcv1.E2SmRcEventTrigger{
    
    RicEventTriggerFormats: ricEventTriggerFormats,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTrigger() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicEventTriggerFormats(eventTriggerFormat1 *e2smrcv1.E2SmRcEventTriggerFormat1, eventTriggerFormat2 *e2smrcv1.E2SmRcEventTriggerFormat2, eventTriggerFormat3 *e2smrcv1.E2SmRcEventTriggerFormat3, eventTriggerFormat4 *e2smrcv1.E2SmRcEventTriggerFormat4, eventTriggerFormat5 *e2smrcv1.E2SmRcEventTriggerFormat5, ) (*e2smrcv1.RicEventTriggerFormats, error) {

    msg := &e2smrcv1.RicEventTriggerFormats{
    
    EventTriggerFormat1: eventTriggerFormat1,
    
    EventTriggerFormat2: eventTriggerFormat2,
    
    EventTriggerFormat3: eventTriggerFormat3,
    
    EventTriggerFormat4: eventTriggerFormat4,
    
    EventTriggerFormat5: eventTriggerFormat5,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicEventTriggerFormats() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcEventTriggerFormat1(messageList *e2smrcv1.E2SmRcEventTriggerFormat1Item, globalAssociatedUeinfo *e2smrcv1.EventTriggerUeInfo, ) (*e2smrcv1.E2SmRcEventTriggerFormat1, error) {

    msg := &e2smrcv1.E2SmRcEventTriggerFormat1{
    
    MessageList: messageList,
    
    GlobalAssociatedUeinfo: globalAssociatedUeinfo,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcEventTriggerFormat1Item(ricEventTriggerConditionId *e2smrcv1.RicEventTriggerConditionId, messageType *e2smrcv1.MessageTypeChoice, messageDirection *e2smrcv1.MessageDirection, associatedUeinfo *e2smrcv1.EventTriggerUeInfo, associatedUeevent *e2smrcv1.EventTriggerUeeventInfo, logicalOr *e2smrcv1.LogicalOr, ) (*e2smrcv1.E2SmRcEventTriggerFormat1Item, error) {

    msg := &e2smrcv1.E2SmRcEventTriggerFormat1Item{
    
    RicEventTriggerConditionId: ricEventTriggerConditionId,
    
    MessageType: messageType,
    
    MessageDirection: messageDirection,
    
    AssociatedUeinfo: associatedUeinfo,
    
    AssociatedUeevent: associatedUeevent,
    
    LogicalOr: logicalOr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat1Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateMessageTypeChoice(messageTypeChoiceNi *e2smrcv1.MessageTypeChoiceNi, messageTypeChoiceRrc *e2smrcv1.MessageTypeChoiceRrc, ) (*e2smrcv1.MessageTypeChoice, error) {

    msg := &e2smrcv1.MessageTypeChoice{
    
    MessageTypeChoiceNi: messageTypeChoiceNi,
    
    MessageTypeChoiceRrc: messageTypeChoiceRrc,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMessageTypeChoice() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateMessageTypeChoiceNi(nIType *e2smcommoniesv1.InterfaceType, nIIdentifier *e2smcommoniesv1.InterfaceIdentifier, nIMessage *e2smcommoniesv1.InterfaceMessageId, ) (*e2smrcv1.MessageTypeChoiceNi, error) {

    msg := &e2smrcv1.MessageTypeChoiceNi{
    
    NIType: nIType,
    
    NIIdentifier: nIIdentifier,
    
    NIMessage: nIMessage,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMessageTypeChoiceNi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateMessageTypeChoiceRrc(rRcMessage *e2smcommoniesv1.RrcMessageId, ) (*e2smrcv1.MessageTypeChoiceRrc, error) {

    msg := &e2smrcv1.MessageTypeChoiceRrc{
    
    RRcMessage: rRcMessage,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMessageTypeChoiceRrc() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcEventTriggerFormat2(ricCallProcessTypeId *e2smrcv1.RicCallProcessTypeId, ricCallProcessBreakpointId *e2smrcv1.RicCallProcessBreakpointId, associatedE2NodeInfo *e2smrcv1.RanparameterTesting, associatedUeinfo *e2smrcv1.EventTriggerUeInfo, ) (*e2smrcv1.E2SmRcEventTriggerFormat2, error) {

    msg := &e2smrcv1.E2SmRcEventTriggerFormat2{
    
    RicCallProcessTypeId: ricCallProcessTypeId,
    
    RicCallProcessBreakpointId: ricCallProcessBreakpointId,
    
    AssociatedE2NodeInfo: associatedE2NodeInfo,
    
    AssociatedUeinfo: associatedUeinfo,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcEventTriggerFormat3(e2NodeInfoChangeList *e2smrcv1.E2SmRcEventTriggerFormat3Item, ) (*e2smrcv1.E2SmRcEventTriggerFormat3, error) {

    msg := &e2smrcv1.E2SmRcEventTriggerFormat3{
    
    E2NodeInfoChangeList: e2NodeInfoChangeList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat3() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcEventTriggerFormat3Item(ricEventTriggerConditionId *e2smrcv1.RicEventTriggerConditionId, e2NodeInfoChangeId int32, associatedCellInfo *e2smrcv1.EventTriggerCellInfo, logicalOr *e2smrcv1.LogicalOr, ) (*e2smrcv1.E2SmRcEventTriggerFormat3Item, error) {

    msg := &e2smrcv1.E2SmRcEventTriggerFormat3Item{
    
    RicEventTriggerConditionId: ricEventTriggerConditionId,
    
    E2NodeInfoChangeId: e2NodeInfoChangeId,
    
    AssociatedCellInfo: associatedCellInfo,
    
    LogicalOr: logicalOr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat3Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcEventTriggerFormat4(uEinfoChangeList *e2smrcv1.E2SmRcEventTriggerFormat4Item, ) (*e2smrcv1.E2SmRcEventTriggerFormat4, error) {

    msg := &e2smrcv1.E2SmRcEventTriggerFormat4{
    
    UEinfoChangeList: uEinfoChangeList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat4() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcEventTriggerFormat4Item(ricEventTriggerConditionId *e2smrcv1.RicEventTriggerConditionId, triggerType *e2smrcv1.TriggerTypeChoice, associatedUeinfo *e2smrcv1.EventTriggerUeInfo, logicalOr *e2smrcv1.LogicalOr, ) (*e2smrcv1.E2SmRcEventTriggerFormat4Item, error) {

    msg := &e2smrcv1.E2SmRcEventTriggerFormat4Item{
    
    RicEventTriggerConditionId: ricEventTriggerConditionId,
    
    TriggerType: triggerType,
    
    AssociatedUeinfo: associatedUeinfo,
    
    LogicalOr: logicalOr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat4Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateTriggerTypeChoice(triggerTypeChoiceRrcstate *e2smrcv1.TriggerTypeChoiceRrcstate, triggerTypeChoiceUeid *e2smrcv1.TriggerTypeChoiceUeid, triggerTypeChoiceL2State *e2smrcv1.TriggerTypeChoiceL2State, ) (*e2smrcv1.TriggerTypeChoice, error) {

    msg := &e2smrcv1.TriggerTypeChoice{
    
    TriggerTypeChoiceRrcstate: triggerTypeChoiceRrcstate,
    
    TriggerTypeChoiceUeid: triggerTypeChoiceUeid,
    
    TriggerTypeChoiceL2State: triggerTypeChoiceL2State,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTriggerTypeChoice() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateTriggerTypeChoiceRrcstate(rrcStateList *e2smrcv1.TriggerTypeChoiceRrcstateItem, ) (*e2smrcv1.TriggerTypeChoiceRrcstate, error) {

    msg := &e2smrcv1.TriggerTypeChoiceRrcstate{
    
    RrcStateList: rrcStateList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTriggerTypeChoiceRrcstate() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateTriggerTypeChoiceRrcstateItem(stateChangedTo *e2smrcv1.RrcState, logicalOr *e2smrcv1.LogicalOr, ) (*e2smrcv1.TriggerTypeChoiceRrcstateItem, error) {

    msg := &e2smrcv1.TriggerTypeChoiceRrcstateItem{
    
    StateChangedTo: stateChangedTo,
    
    LogicalOr: logicalOr,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTriggerTypeChoiceRrcstateItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateTriggerTypeChoiceUeID(ueIdchangeId int32, ) (*e2smrcv1.TriggerTypeChoiceUeid, error) {

    msg := &e2smrcv1.TriggerTypeChoiceUeid{
    
    UeIdchangeId: ueIdchangeId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTriggerTypeChoiceUeID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateTriggerTypeChoiceL2state(associatedL2Variables *e2smrcv1.RanparameterTesting, ) (*e2smrcv1.TriggerTypeChoiceL2state, error) {

    msg := &e2smrcv1.TriggerTypeChoiceL2state{
    
    AssociatedL2Variables: associatedL2Variables,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTriggerTypeChoiceL2state() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcEventTriggerFormat5(onDemand *e2smrcv1.OnDemand, associatedUeinfo *e2smrcv1.EventTriggerUeInfo, associatedCellInfo *e2smrcv1.EventTriggerCellInfo, ) (*e2smrcv1.E2SmRcEventTriggerFormat5, error) {

    msg := &e2smrcv1.E2SmRcEventTriggerFormat5{
    
    OnDemand: onDemand,
    
    AssociatedUeinfo: associatedUeinfo,
    
    AssociatedCellInfo: associatedCellInfo,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcEventTriggerFormat5() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcActionDefinition(ricStyleType *e2smcommoniesv1.RicStyleType, ricActionDefinitionFormats *e2smrcv1.RicActionDefinitionFormats, ) (*e2smrcv1.E2SmRcActionDefinition, error) {

    msg := &e2smrcv1.E2SmRcActionDefinition{
    
    RicStyleType: ricStyleType,
    
    RicActionDefinitionFormats: ricActionDefinitionFormats,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcActionDefinition() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicActionDefinitionFormats(actionDefinitionFormat1 *e2smrcv1.E2SmRcActionDefinitionFormat1, actionDefinitionFormat2 *e2smrcv1.E2SmRcActionDefinitionFormat2, actionDefinitionFormat3 *e2smrcv1.E2SmRcActionDefinitionFormat3, ) (*e2smrcv1.RicActionDefinitionFormats, error) {

    msg := &e2smrcv1.RicActionDefinitionFormats{
    
    ActionDefinitionFormat1: actionDefinitionFormat1,
    
    ActionDefinitionFormat2: actionDefinitionFormat2,
    
    ActionDefinitionFormat3: actionDefinitionFormat3,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicActionDefinitionFormats() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcActionDefinitionFormat1(ranPToBeReportedList *e2smrcv1.E2SmRcActionDefinitionFormat1Item, ) (*e2smrcv1.E2SmRcActionDefinitionFormat1, error) {

    msg := &e2smrcv1.E2SmRcActionDefinitionFormat1{
    
    RanPToBeReportedList: ranPToBeReportedList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcActionDefinitionFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcActionDefinitionFormat1Item(ranParameterId *e2smrcv1.RanparameterId, ) (*e2smrcv1.E2SmRcActionDefinitionFormat1Item, error) {

    msg := &e2smrcv1.E2SmRcActionDefinitionFormat1Item{
    
    RanParameterId: ranParameterId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcActionDefinitionFormat1Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcActionDefinitionFormat2(ricPolicyConditionsList *e2smrcv1.E2SmRcActionDefinitionFormat2Item, ) (*e2smrcv1.E2SmRcActionDefinitionFormat2, error) {

    msg := &e2smrcv1.E2SmRcActionDefinitionFormat2{
    
    RicPolicyConditionsList: ricPolicyConditionsList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcActionDefinitionFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcActionDefinitionFormat2Item(ricPolicyAction *e2smrcv1.RicPolicyAction, ricPolicyConditionDefinition *e2smrcv1.RanparameterTesting, ) (*e2smrcv1.E2SmRcActionDefinitionFormat2Item, error) {

    msg := &e2smrcv1.E2SmRcActionDefinitionFormat2Item{
    
    RicPolicyAction: ricPolicyAction,
    
    RicPolicyConditionDefinition: ricPolicyConditionDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcActionDefinitionFormat2Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcActionDefinitionFormat3(ricInsertIndicationId *e2smrcv1.RicInsertIndicationId, ranPInsertIndicationList *e2smrcv1.E2SmRcActionDefinitionFormat3Item, ueId *e2smcommoniesv1.Ueid, ) (*e2smrcv1.E2SmRcActionDefinitionFormat3, error) {

    msg := &e2smrcv1.E2SmRcActionDefinitionFormat3{
    
    RicInsertIndicationId: ricInsertIndicationId,
    
    RanPInsertIndicationList: ranPInsertIndicationList,
    
    UeId: ueId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcActionDefinitionFormat3() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcActionDefinitionFormat3Item(ranParameterId *e2smrcv1.RanparameterId, ) (*e2smrcv1.E2SmRcActionDefinitionFormat3Item, error) {

    msg := &e2smrcv1.E2SmRcActionDefinitionFormat3Item{
    
    RanParameterId: ranParameterId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcActionDefinitionFormat3Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationHeader(ricIndicationHeaderFormats *e2smrcv1.RicIndicationHeaderFormats, ) (*e2smrcv1.E2SmRcIndicationHeader, error) {

    msg := &e2smrcv1.E2SmRcIndicationHeader{
    
    RicIndicationHeaderFormats: ricIndicationHeaderFormats,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationHeader() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicIndicationHeaderFormats(indicationHeaderFormat1 *e2smrcv1.E2SmRcIndicationHeaderFormat1, indicationHeaderFormat2 *e2smrcv1.E2SmRcIndicationHeaderFormat2, ) (*e2smrcv1.RicIndicationHeaderFormats, error) {

    msg := &e2smrcv1.RicIndicationHeaderFormats{
    
    IndicationHeaderFormat1: indicationHeaderFormat1,
    
    IndicationHeaderFormat2: indicationHeaderFormat2,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicIndicationHeaderFormats() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationHeaderFormat1(ricEventTriggerConditionId *e2smrcv1.RicEventTriggerConditionId, ) (*e2smrcv1.E2SmRcIndicationHeaderFormat1, error) {

    msg := &e2smrcv1.E2SmRcIndicationHeaderFormat1{
    
    RicEventTriggerConditionId: ricEventTriggerConditionId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationHeaderFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationHeaderFormat2(ueId *e2smcommoniesv1.Ueid, ricInsertStyleType *e2smcommoniesv1.RicStyleType, ricInsertIndicationId *e2smrcv1.RicInsertIndicationId, ) (*e2smrcv1.E2SmRcIndicationHeaderFormat2, error) {

    msg := &e2smrcv1.E2SmRcIndicationHeaderFormat2{
    
    UeId: ueId,
    
    RicInsertStyleType: ricInsertStyleType,
    
    RicInsertIndicationId: ricInsertIndicationId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationHeaderFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessage(ricIndicationMessageFormats *e2smrcv1.RicIndicationMessageFormats, ) (*e2smrcv1.E2SmRcIndicationMessage, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessage{
    
    RicIndicationMessageFormats: ricIndicationMessageFormats,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessage() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicIndicationMessageFormats(indicationMessageFormat1 *e2smrcv1.E2SmRcIndicationMessageFormat1, indicationMessageFormat2 *e2smrcv1.E2SmRcIndicationMessageFormat2, indicationMessageFormat3 *e2smrcv1.E2SmRcIndicationMessageFormat3, indicationMessageFormat4 *e2smrcv1.E2SmRcIndicationMessageFormat4, indicationMessageFormat5 *e2smrcv1.E2SmRcIndicationMessageFormat5, ) (*e2smrcv1.RicIndicationMessageFormats, error) {

    msg := &e2smrcv1.RicIndicationMessageFormats{
    
    IndicationMessageFormat1: indicationMessageFormat1,
    
    IndicationMessageFormat2: indicationMessageFormat2,
    
    IndicationMessageFormat3: indicationMessageFormat3,
    
    IndicationMessageFormat4: indicationMessageFormat4,
    
    IndicationMessageFormat5: indicationMessageFormat5,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicIndicationMessageFormats() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat1(ranPReportedList *e2smrcv1.E2SmRcIndicationMessageFormat1Item, ) (*e2smrcv1.E2SmRcIndicationMessageFormat1, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat1{
    
    RanPReportedList: ranPReportedList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat1Item(ranParameterId *e2smrcv1.RanparameterId, ranParameterValueType *e2smrcv1.RanparameterValueType, ) (*e2smrcv1.E2SmRcIndicationMessageFormat1Item, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat1Item{
    
    RanParameterId: ranParameterId,
    
    RanParameterValueType: ranParameterValueType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat1Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat2(ueParameterList *e2smrcv1.E2SmRcIndicationMessageFormat2Item, ) (*e2smrcv1.E2SmRcIndicationMessageFormat2, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat2{
    
    UeParameterList: ueParameterList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat2Item(ueId *e2smcommoniesv1.Ueid, ranPList *e2smrcv1.E2SmRcIndicationMessageFormat2RanparameterItem, ) (*e2smrcv1.E2SmRcIndicationMessageFormat2Item, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat2Item{
    
    UeId: ueId,
    
    RanPList: ranPList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat2Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat2RanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterValueType *e2smrcv1.RanparameterValueType, ) (*e2smrcv1.E2SmRcIndicationMessageFormat2RanparameterItem, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat2RanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterValueType: ranParameterValueType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat2RanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat3(cellInfoList *e2smrcv1.E2SmRcIndicationMessageFormat3Item, ) (*e2smrcv1.E2SmRcIndicationMessageFormat3, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat3{
    
    CellInfoList: cellInfoList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat3() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat3Item(cellGlobalId *e2smcommoniesv1.Cgi, cellContextInfo []byte, cellDeleted bool, neighborRelationTable *e2smrcv1.NeighborRelationInfo, ) (*e2smrcv1.E2SmRcIndicationMessageFormat3Item, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat3Item{
    
    CellGlobalId: cellGlobalId,
    
    CellContextInfo: cellContextInfo,
    
    CellDeleted: cellDeleted,
    
    NeighborRelationTable: neighborRelationTable,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat3Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat4(ueInfoList *e2smrcv1.E2SmRcIndicationMessageFormat4ItemUe, cellInfoList *e2smrcv1.E2SmRcIndicationMessageFormat4ItemCell, ) (*e2smrcv1.E2SmRcIndicationMessageFormat4, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat4{
    
    UeInfoList: ueInfoList,
    
    CellInfoList: cellInfoList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat4() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat4ItemUe(ueId *e2smcommoniesv1.Ueid, ueContextInfo []byte, cellGlobalId *e2smcommoniesv1.Cgi, ) (*e2smrcv1.E2SmRcIndicationMessageFormat4ItemUe, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat4ItemUe{
    
    UeId: ueId,
    
    UeContextInfo: ueContextInfo,
    
    CellGlobalId: cellGlobalId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat4ItemUe() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat4ItemCell(cellGlobalId *e2smcommoniesv1.Cgi, cellContextInfo []byte, neighborRelationTable *e2smrcv1.NeighborRelationInfo, ) (*e2smrcv1.E2SmRcIndicationMessageFormat4ItemCell, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat4ItemCell{
    
    CellGlobalId: cellGlobalId,
    
    CellContextInfo: cellContextInfo,
    
    NeighborRelationTable: neighborRelationTable,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat4ItemCell() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat5(ranPRequestedList *e2smrcv1.E2SmRcIndicationMessageFormat5Item, ) (*e2smrcv1.E2SmRcIndicationMessageFormat5, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat5{
    
    RanPRequestedList: ranPRequestedList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat5() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcIndicationMessageFormat5Item(ranParameterId *e2smrcv1.RanparameterId, ranParameterValueType *e2smrcv1.RanparameterValueType, ) (*e2smrcv1.E2SmRcIndicationMessageFormat5Item, error) {

    msg := &e2smrcv1.E2SmRcIndicationMessageFormat5Item{
    
    RanParameterId: ranParameterId,
    
    RanParameterValueType: ranParameterValueType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcIndicationMessageFormat5Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcCallProcessID(ricCallProcessIdFormats *e2smrcv1.RicCallProcessIdFormats, ) (*e2smrcv1.E2SmRcCallProcessId, error) {

    msg := &e2smrcv1.E2SmRcCallProcessId{
    
    RicCallProcessIdFormats: ricCallProcessIdFormats,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcCallProcessID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicCallProcessIDFormats(callProcessIdFormat1 *e2smrcv1.E2SmRcCallProcessIdFormat1, ) (*e2smrcv1.RicCallProcessIdFormats, error) {

    msg := &e2smrcv1.RicCallProcessIdFormats{
    
    CallProcessIdFormat1: callProcessIdFormat1,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicCallProcessIDFormats() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcCallProcessIDFormat1(ricCallProcessId *e2smrcv1.RanCallProcessId, ) (*e2smrcv1.E2SmRcCallProcessIdFormat1, error) {

    msg := &e2smrcv1.E2SmRcCallProcessIdFormat1{
    
    RicCallProcessId: ricCallProcessId,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcCallProcessIDFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcControlHeader(ricControlHeaderFormats *e2smrcv1.RicControlHeaderFormats, ) (*e2smrcv1.E2SmRcControlHeader, error) {

    msg := &e2smrcv1.E2SmRcControlHeader{
    
    RicControlHeaderFormats: ricControlHeaderFormats,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlHeader() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicControlHeaderFormats(controlHeaderFormat1 *e2smrcv1.E2SmRcControlHeaderFormat1, ) (*e2smrcv1.RicControlHeaderFormats, error) {

    msg := &e2smrcv1.RicControlHeaderFormats{
    
    ControlHeaderFormat1: controlHeaderFormat1,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicControlHeaderFormats() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcControlHeaderFormat1(ueId *e2smcommoniesv1.Ueid, ricStyleType *e2smcommoniesv1.RicStyleType, ricControlActionId *e2smrcv1.RicControlActionId, ricControlDecision *e2smrcv1.RicControlDecision, ) (*e2smrcv1.E2SmRcControlHeaderFormat1, error) {

    msg := &e2smrcv1.E2SmRcControlHeaderFormat1{
    
    UeId: ueId,
    
    RicStyleType: ricStyleType,
    
    RicControlActionId: ricControlActionId,
    
    RicControlDecision: ricControlDecision,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlHeaderFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcControlMessage(ricControlMessageFormats *e2smrcv1.RicControlMessageFormats, ) (*e2smrcv1.E2SmRcControlMessage, error) {

    msg := &e2smrcv1.E2SmRcControlMessage{
    
    RicControlMessageFormats: ricControlMessageFormats,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlMessage() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicControlMessageFormats(controlMessageFormat1 *e2smrcv1.E2SmRcControlMessageFormat1, ) (*e2smrcv1.RicControlMessageFormats, error) {

    msg := &e2smrcv1.RicControlMessageFormats{
    
    ControlMessageFormat1: controlMessageFormat1,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicControlMessageFormats() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcControlMessageFormat1(ranPList *e2smrcv1.E2SmRcControlMessageFormat1Item, ) (*e2smrcv1.E2SmRcControlMessageFormat1, error) {

    msg := &e2smrcv1.E2SmRcControlMessageFormat1{
    
    RanPList: ranPList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlMessageFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcControlMessageFormat1Item(ranParameterId *e2smrcv1.RanparameterId, ranParameterValueType *e2smrcv1.RanparameterValueType, ) (*e2smrcv1.E2SmRcControlMessageFormat1Item, error) {

    msg := &e2smrcv1.E2SmRcControlMessageFormat1Item{
    
    RanParameterId: ranParameterId,
    
    RanParameterValueType: ranParameterValueType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlMessageFormat1Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcControlOutcome(ricControlOutcomeFormats *e2smrcv1.RicControlOutcomeFormats, ) (*e2smrcv1.E2SmRcControlOutcome, error) {

    msg := &e2smrcv1.E2SmRcControlOutcome{
    
    RicControlOutcomeFormats: ricControlOutcomeFormats,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcome() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicControlOutcomeFormats(controlOutcomeFormat1 *e2smrcv1.E2SmRcControlOutcomeFormat1, ) (*e2smrcv1.RicControlOutcomeFormats, error) {

    msg := &e2smrcv1.RicControlOutcomeFormats{
    
    ControlOutcomeFormat1: controlOutcomeFormat1,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicControlOutcomeFormats() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcControlOutcomeFormat1(ranPList *e2smrcv1.E2SmRcControlOutcomeFormat1Item, ) (*e2smrcv1.E2SmRcControlOutcomeFormat1, error) {

    msg := &e2smrcv1.E2SmRcControlOutcomeFormat1{
    
    RanPList: ranPList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcControlOutcomeFormat1Item(ranParameterId *e2smrcv1.RanparameterId, ranParameterValue *e2smrcv1.RanparameterValue, ) (*e2smrcv1.E2SmRcControlOutcomeFormat1Item, error) {

    msg := &e2smrcv1.E2SmRcControlOutcomeFormat1Item{
    
    RanParameterId: ranParameterId,
    
    RanParameterValue: ranParameterValue,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcControlOutcomeFormat1Item() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmRcRanfunctionDefinition(ranFunctionName *e2smcommoniesv1.RanfunctionName, ranFunctionDefinitionEventTrigger *e2smrcv1.RanfunctionDefinitionEventTrigger, ranFunctionDefinitionReport *e2smrcv1.RanfunctionDefinitionReport, ranFunctionDefinitionInsert *e2smrcv1.RanfunctionDefinitionInsert, ranFunctionDefinitionControl *e2smrcv1.RanfunctionDefinitionControl, ranFunctionDefinitionPolicy *e2smrcv1.RanfunctionDefinitionPolicy, ) (*e2smrcv1.E2SmRcRanfunctionDefinition, error) {

    msg := &e2smrcv1.E2SmRcRanfunctionDefinition{
    
    RanFunctionName: ranFunctionName,
    
    RanFunctionDefinitionEventTrigger: ranFunctionDefinitionEventTrigger,
    
    RanFunctionDefinitionReport: ranFunctionDefinitionReport,
    
    RanFunctionDefinitionInsert: ranFunctionDefinitionInsert,
    
    RanFunctionDefinitionControl: ranFunctionDefinitionControl,
    
    RanFunctionDefinitionPolicy: ranFunctionDefinitionPolicy,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRcRanfunctionDefinition() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionEventTrigger(ricEventTriggerStyleList *e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem, ranL2ParametersList *e2smrcv1.L2ParametersRanparameterItem, ranCallProcessTypesList *e2smrcv1.RanfunctionDefinitionEventTriggerCallProcessItem, ranUeidentificationParametersList *e2smrcv1.UeidentificationRanparameterItem, ranCellIdentificationParametersList *e2smrcv1.CellIdentificationRanparameterItem, ) (*e2smrcv1.RanfunctionDefinitionEventTrigger, error) {

    msg := &e2smrcv1.RanfunctionDefinitionEventTrigger{
    
    RicEventTriggerStyleList: ricEventTriggerStyleList,
    
    RanL2ParametersList: ranL2ParametersList,
    
    RanCallProcessTypesList: ranCallProcessTypesList,
    
    RanUeidentificationParametersList: ranUeidentificationParametersList,
    
    RanCellIdentificationParametersList: ranCellIdentificationParametersList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionEventTrigger() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionEventTriggerStyleItem(ricEventTriggerStyleType *e2smcommoniesv1.RicStyleType, ricEventTriggerStyleName *e2smcommoniesv1.RicStyleName, ricEventTriggerFormatType *e2smcommoniesv1.RicFormatType, ) (*e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionEventTriggerStyleItem{
    
    RicEventTriggerStyleType: ricEventTriggerStyleType,
    
    RicEventTriggerStyleName: ricEventTriggerStyleName,
    
    RicEventTriggerFormatType: ricEventTriggerFormatType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionEventTriggerStyleItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateL2ParametersRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.L2ParametersRanparameterItem, error) {

    msg := &e2smrcv1.L2ParametersRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateL2ParametersRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDentificationRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.UeidentificationRanparameterItem, error) {

    msg := &e2smrcv1.UeidentificationRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDentificationRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateCellIDentificationRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.CellIdentificationRanparameterItem, error) {

    msg := &e2smrcv1.CellIdentificationRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellIDentificationRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionEventTriggerCallProcessItem(callProcessTypeId *e2smrcv1.RicCallProcessTypeId, callProcessTypeName *e2smrcv1.RicCallProcessTypeName, callProcessBreakpointsList *e2smrcv1.RanfunctionDefinitionEventTriggerBreakpointItem, ) (*e2smrcv1.RanfunctionDefinitionEventTriggerCallProcessItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionEventTriggerCallProcessItem{
    
    CallProcessTypeId: callProcessTypeId,
    
    CallProcessTypeName: callProcessTypeName,
    
    CallProcessBreakpointsList: callProcessBreakpointsList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionEventTriggerCallProcessItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionEventTriggerBreakpointItem(callProcessBreakpointId *e2smrcv1.RicCallProcessBreakpointId, callProcessBreakpointName *e2smrcv1.RicCallProcessBreakpointName, ranCallProcessBreakpointParametersList *e2smrcv1.CallProcessBreakpointRanparameterItem, ) (*e2smrcv1.RanfunctionDefinitionEventTriggerBreakpointItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionEventTriggerBreakpointItem{
    
    CallProcessBreakpointId: callProcessBreakpointId,
    
    CallProcessBreakpointName: callProcessBreakpointName,
    
    RanCallProcessBreakpointParametersList: ranCallProcessBreakpointParametersList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionEventTriggerBreakpointItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateCallProcessBreakpointRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.CallProcessBreakpointRanparameterItem, error) {

    msg := &e2smrcv1.CallProcessBreakpointRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCallProcessBreakpointRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionReport(ricReportStyleList *e2smrcv1.RanfunctionDefinitionReportItem, ) (*e2smrcv1.RanfunctionDefinitionReport, error) {

    msg := &e2smrcv1.RanfunctionDefinitionReport{
    
    RicReportStyleList: ricReportStyleList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionReport() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionReportItem(ricReportStyleType *e2smcommoniesv1.RicStyleType, ricReportStyleName *e2smcommoniesv1.RicStyleName, ricSupportedEventTriggerStyleType *e2smcommoniesv1.RicStyleType, ricReportActionFormatType *e2smcommoniesv1.RicFormatType, ricIndicationHeaderFormatType *e2smcommoniesv1.RicFormatType, ricIndicationMessageFormatType *e2smcommoniesv1.RicFormatType, ranReportParametersList *e2smrcv1.ReportRanparameterItem, ) (*e2smrcv1.RanfunctionDefinitionReportItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionReportItem{
    
    RicReportStyleType: ricReportStyleType,
    
    RicReportStyleName: ricReportStyleName,
    
    RicSupportedEventTriggerStyleType: ricSupportedEventTriggerStyleType,
    
    RicReportActionFormatType: ricReportActionFormatType,
    
    RicIndicationHeaderFormatType: ricIndicationHeaderFormatType,
    
    RicIndicationMessageFormatType: ricIndicationMessageFormatType,
    
    RanReportParametersList: ranReportParametersList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionReportItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateReportRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.ReportRanparameterItem, error) {

    msg := &e2smrcv1.ReportRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateReportRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionInsert(ricInsertStyleList *e2smrcv1.RanfunctionDefinitionInsertItem, ) (*e2smrcv1.RanfunctionDefinitionInsert, error) {

    msg := &e2smrcv1.RanfunctionDefinitionInsert{
    
    RicInsertStyleList: ricInsertStyleList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionInsert() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionInsertItem(ricInsertStyleType *e2smcommoniesv1.RicStyleType, ricInsertStyleName *e2smcommoniesv1.RicStyleName, ricSupportedEventTriggerStyleType *e2smcommoniesv1.RicStyleType, ricActionDefinitionFormatType *e2smcommoniesv1.RicFormatType, ricInsertIndicationList *e2smrcv1.RanfunctionDefinitionInsertIndicationItem, ricIndicationHeaderFormatType *e2smcommoniesv1.RicFormatType, ricIndicationMessageFormatType *e2smcommoniesv1.RicFormatType, ricCallProcessIdformatType *e2smcommoniesv1.RicFormatType, ) (*e2smrcv1.RanfunctionDefinitionInsertItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionInsertItem{
    
    RicInsertStyleType: ricInsertStyleType,
    
    RicInsertStyleName: ricInsertStyleName,
    
    RicSupportedEventTriggerStyleType: ricSupportedEventTriggerStyleType,
    
    RicActionDefinitionFormatType: ricActionDefinitionFormatType,
    
    RicInsertIndicationList: ricInsertIndicationList,
    
    RicIndicationHeaderFormatType: ricIndicationHeaderFormatType,
    
    RicIndicationMessageFormatType: ricIndicationMessageFormatType,
    
    RicCallProcessIdformatType: ricCallProcessIdformatType,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionInsertItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionInsertIndicationItem(ricInsertIndicationId *e2smrcv1.RicInsertIndicationId, ricInsertIndicationName *e2smrcv1.RicInsertIndicationName, ranInsertIndicationParametersList *e2smrcv1.InsertIndicationRanparameterItem, ) (*e2smrcv1.RanfunctionDefinitionInsertIndicationItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionInsertIndicationItem{
    
    RicInsertIndicationId: ricInsertIndicationId,
    
    RicInsertIndicationName: ricInsertIndicationName,
    
    RanInsertIndicationParametersList: ranInsertIndicationParametersList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionInsertIndicationItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInsertIndicationRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.InsertIndicationRanparameterItem, error) {

    msg := &e2smrcv1.InsertIndicationRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInsertIndicationRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionControl(ricControlStyleList *e2smrcv1.RanfunctionDefinitionControlItem, ) (*e2smrcv1.RanfunctionDefinitionControl, error) {

    msg := &e2smrcv1.RanfunctionDefinitionControl{
    
    RicControlStyleList: ricControlStyleList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionControl() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionControlItem(ricControlStyleType *e2smcommoniesv1.RicStyleType, ricControlStyleName *e2smcommoniesv1.RicStyleName, ricControlActionList *e2smrcv1.RanfunctionDefinitionControlActionItem, ricControlHeaderFormatType *e2smcommoniesv1.RicFormatType, ricControlMessageFormatType *e2smcommoniesv1.RicFormatType, ricCallProcessIdformatType *e2smcommoniesv1.RicFormatType, ricControlOutcomeFormatType *e2smcommoniesv1.RicFormatType, ranControlOutcomeParametersList *e2smrcv1.ControlOutcomeRanparameterItem, ) (*e2smrcv1.RanfunctionDefinitionControlItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionControlItem{
    
    RicControlStyleType: ricControlStyleType,
    
    RicControlStyleName: ricControlStyleName,
    
    RicControlActionList: ricControlActionList,
    
    RicControlHeaderFormatType: ricControlHeaderFormatType,
    
    RicControlMessageFormatType: ricControlMessageFormatType,
    
    RicCallProcessIdformatType: ricCallProcessIdformatType,
    
    RicControlOutcomeFormatType: ricControlOutcomeFormatType,
    
    RanControlOutcomeParametersList: ranControlOutcomeParametersList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionControlItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateControlOutcomeRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.ControlOutcomeRanparameterItem, error) {

    msg := &e2smrcv1.ControlOutcomeRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateControlOutcomeRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionControlActionItem(ricControlActionId *e2smrcv1.RicControlActionId, ricControlActionName *e2smrcv1.RicControlActionName, ranControlActionParametersList *e2smrcv1.ControlActionRanparameterItem, ) (*e2smrcv1.RanfunctionDefinitionControlActionItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionControlActionItem{
    
    RicControlActionId: ricControlActionId,
    
    RicControlActionName: ricControlActionName,
    
    RanControlActionParametersList: ranControlActionParametersList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionControlActionItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateControlActionRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.ControlActionRanparameterItem, error) {

    msg := &e2smrcv1.ControlActionRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateControlActionRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionPolicy(ricPolicyStyleList *e2smrcv1.RanfunctionDefinitionPolicyItem, ) (*e2smrcv1.RanfunctionDefinitionPolicy, error) {

    msg := &e2smrcv1.RanfunctionDefinitionPolicy{
    
    RicPolicyStyleList: ricPolicyStyleList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionPolicy() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionPolicyItem(ricPolicyStyleType *e2smcommoniesv1.RicStyleType, ricPolicyStyleName *e2smcommoniesv1.RicStyleName, ricSupportedEventTriggerStyleType *e2smcommoniesv1.RicStyleType, ricPolicyActionList *e2smrcv1.RanfunctionDefinitionPolicyActionItem, ) (*e2smrcv1.RanfunctionDefinitionPolicyItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionPolicyItem{
    
    RicPolicyStyleType: ricPolicyStyleType,
    
    RicPolicyStyleName: ricPolicyStyleName,
    
    RicSupportedEventTriggerStyleType: ricSupportedEventTriggerStyleType,
    
    RicPolicyActionList: ricPolicyActionList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionPolicyItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanfunctionDefinitionPolicyActionItem(ricPolicyActionId *e2smrcv1.RicControlActionId, ricPolicyActionName *e2smrcv1.RicControlActionName, ricActionDefinitionFormatType *e2smcommoniesv1.RicFormatType, ranPolicyActionParametersList *e2smrcv1.PolicyActionRanparameterItem, ranPolicyConditionParametersList *e2smrcv1.PolicyConditionRanparameterItem, ) (*e2smrcv1.RanfunctionDefinitionPolicyActionItem, error) {

    msg := &e2smrcv1.RanfunctionDefinitionPolicyActionItem{
    
    RicPolicyActionId: ricPolicyActionId,
    
    RicPolicyActionName: ricPolicyActionName,
    
    RicActionDefinitionFormatType: ricActionDefinitionFormatType,
    
    RanPolicyActionParametersList: ranPolicyActionParametersList,
    
    RanPolicyConditionParametersList: ranPolicyConditionParametersList,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanfunctionDefinitionPolicyActionItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreatePolicyActionRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.PolicyActionRanparameterItem, error) {

    msg := &e2smrcv1.PolicyActionRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreatePolicyActionRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreatePolicyConditionRanparameterItem(ranParameterId *e2smrcv1.RanparameterId, ranParameterName *e2smrcv1.RanparameterName, ranParameterDefinition *e2smrcv1.RanparameterDefinition, ) (*e2smrcv1.PolicyConditionRanparameterItem, error) {

    msg := &e2smrcv1.PolicyConditionRanparameterItem{
    
    RanParameterId: ranParameterId,
    
    RanParameterName: ranParameterName,
    
    RanParameterDefinition: ranParameterDefinition,
     }

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreatePolicyConditionRanparameterItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

