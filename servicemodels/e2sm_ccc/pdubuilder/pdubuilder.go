// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smcommoniesv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-common-ies"

	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

	e2smcccv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_ccc/v1/e2sm-ccc-ies"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateInterfaceIDNG(guami *e2smcommoniesv1.Guami) (*e2smcommoniesv1.InterfaceIdNG, error) {

	msg := &e2smcommoniesv1.InterfaceIdNG{}
	msg.Guami = guami

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDNG() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDXn(globalNgRanID *e2smcommoniesv1.GlobalNgrannodeId) (*e2smcommoniesv1.InterfaceIdXn, error) {

	msg := &e2smcommoniesv1.InterfaceIdXn{}
	msg.GlobalNgRanId = globalNgRanID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDXn() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDF1(globalGnbID *e2smcommoniesv1.GlobalGnbID, gNbDuID *e2smcommoniesv1.GnbDUID) (*e2smcommoniesv1.InterfaceIdF1, error) {

	msg := &e2smcommoniesv1.InterfaceIdF1{}
	msg.GlobalGnbId = globalGnbID
	msg.GNbDuId = gNbDuID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDF1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDE1(globalGnbID *e2smcommoniesv1.GlobalGnbID, gNbCuUpID *e2smcommoniesv1.GnbCUUPID) (*e2smcommoniesv1.InterfaceIdE1, error) {

	msg := &e2smcommoniesv1.InterfaceIdE1{}
	msg.GlobalGnbId = globalGnbID
	msg.GNbCuUpId = gNbCuUpID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDE1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDS1(gUmmei *e2smcommoniesv1.Gummei) (*e2smcommoniesv1.InterfaceIdS1, error) {

	msg := &e2smcommoniesv1.InterfaceIdS1{}
	msg.GUmmei = gUmmei

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDS1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDX2(nodeType *e2smcommoniesv1.NodeTypeInterfaceIdX2) (*e2smcommoniesv1.InterfaceIdX2, error) {

	msg := &e2smcommoniesv1.InterfaceIdX2{}
	msg.NodeType = nodeType

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDX2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceIDW1(globalNgENbID *e2smcommoniesv1.GlobalNgEnbID, ngENbDuID *e2smcommoniesv1.NgenbDUID) (*e2smcommoniesv1.InterfaceIdW1, error) {

	msg := &e2smcommoniesv1.InterfaceIdW1{}
	msg.GlobalNgENbId = globalNgENbID
	msg.NgENbDuId = ngENbDuID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDW1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInterfaceMessageID(interfaceProcedureID int32, messageType e2smcommoniesv1.MessageTypeInterfaceMessageId) (*e2smcommoniesv1.InterfaceMessageId, error) {

	msg := &e2smcommoniesv1.InterfaceMessageId{}
	msg.InterfaceProcedureId = interfaceProcedureID
	msg.MessageType = messageType

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceMessageID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicFormatType(value int32) (*e2smcommoniesv1.RicFormatType, error) {

	msg := &e2smcommoniesv1.RicFormatType{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicFormatType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicStyleType(value int32) (*e2smcommoniesv1.RicStyleType, error) {

	msg := &e2smcommoniesv1.RicStyleType{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicStyleType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRicStyleName(value string) (*e2smcommoniesv1.RicStyleName, error) {

	msg := &e2smcommoniesv1.RicStyleName{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRicStyleName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRrcMessageID(rrcType *e2smcommoniesv1.RrcTypeRrcMessageId, messageID int32) (*e2smcommoniesv1.RrcMessageId, error) {

	msg := &e2smcommoniesv1.RrcMessageId{}
	msg.RrcType = rrcType
	msg.MessageId = messageID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRrcMessageID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGNbCUCPE1ApIDList(value []*e2smcommoniesv1.UeidGNbCUCPE1ApIDItem) (*e2smcommoniesv1.UeidGNbCUCPE1ApIDList, error) {

	msg := &e2smcommoniesv1.UeidGNbCUCPE1ApIDList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGNbCUCPE1ApIDList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGNbCUCPE1ApIDItem(gNbCuCpUeE1ApID *e2smcommoniesv1.GnbCUCPUEE1ApID) (*e2smcommoniesv1.UeidGNbCUCPE1ApIDItem, error) {

	msg := &e2smcommoniesv1.UeidGNbCUCPE1ApIDItem{}
	msg.GNbCuCpUeE1ApId = gNbCuCpUeE1ApID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGNbCUCPE1ApIDItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGNbCUF1ApIDList(value []*e2smcommoniesv1.UeidGNbCUCPF1ApIDItem) (*e2smcommoniesv1.UeidGNbCUF1ApIDList, error) {

	msg := &e2smcommoniesv1.UeidGNbCUF1ApIDList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGNbCUF1ApIDList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDGNbCUCPF1ApIDItem(gNbCuUeF1ApID *e2smcommoniesv1.GnbCUUEF1ApID) (*e2smcommoniesv1.UeidGNbCUCPF1ApIDItem, error) {

	msg := &e2smcommoniesv1.UeidGNbCUCPF1ApIDItem{}
	msg.GNbCuUeF1ApId = gNbCuUeF1ApID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGNbCUCPF1ApIDItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateUeIDNGENbDU(ngENbCuUeW1ApID *e2smcommoniesv1.NgenbCUUEW1ApID) (*e2smcommoniesv1.UeidNGENbDU, error) {

	msg := &e2smcommoniesv1.UeidNGENbDU{}
	msg.NgENbCuUeW1ApId = ngENbCuUeW1ApID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDNGENbDU() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGlobalEnbID(pLmnidentity *e2smcommoniesv1.Plmnidentity, eNbID *e2smcommoniesv1.EnbID) (*e2smcommoniesv1.GlobalEnbID, error) {

	msg := &e2smcommoniesv1.GlobalEnbID{}
	msg.PLmnidentity = pLmnidentity
	msg.ENbId = eNbID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalEnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGummei(pLmnIdentity *e2smcommoniesv1.Plmnidentity, mMeGroupID *e2smcommoniesv1.MmeGroupID, mMeCode *e2smcommoniesv1.MmeCode) (*e2smcommoniesv1.Gummei, error) {

	msg := &e2smcommoniesv1.Gummei{}
	msg.PLmnIdentity = pLmnIdentity
	msg.MMeGroupId = mMeGroupID
	msg.MMeCode = mMeCode

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGummei() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateMmeGroupID(value []byte) (*e2smcommoniesv1.MmeGroupID, error) {

	msg := &e2smcommoniesv1.MmeGroupID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMmeGroupID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateMmeCode(value []byte) (*e2smcommoniesv1.MmeCode, error) {

	msg := &e2smcommoniesv1.MmeCode{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMmeCode() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateMmeUES1ApID(value int64) (*e2smcommoniesv1.MmeUES1ApID, error) {

	msg := &e2smcommoniesv1.MmeUES1ApID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMmeUES1ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateQci(value int32) (*e2smcommoniesv1.Qci, error) {

	msg := &e2smcommoniesv1.Qci{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateQci() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSubscriberProfileIDforRfp(value int32) (*e2smcommoniesv1.SubscriberProfileIdforRfp, error) {

	msg := &e2smcommoniesv1.SubscriberProfileIdforRfp{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSubscriberProfileIDforRfp() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEnbUEX2ApID(value int32) (*e2smcommoniesv1.EnbUEX2ApID, error) {

	msg := &e2smcommoniesv1.EnbUEX2ApID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbUEX2ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEnbUEX2ApIDExtension(value int32) (*e2smcommoniesv1.EnbUEX2ApIDExtension, error) {

	msg := &e2smcommoniesv1.EnbUEX2ApIDExtension{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbUEX2ApIDExtension() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEUTraARfcn(value int32) (*e2smcommoniesv1.EUTraARfcn, error) {

	msg := &e2smcommoniesv1.EUTraARfcn{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEUTraARfcn() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEUTraPCi(value int32) (*e2smcommoniesv1.EUTraPCi, error) {

	msg := &e2smcommoniesv1.EUTraPCi{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEUTraPCi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEUTraTAc(value []byte) (*e2smcommoniesv1.EUTraTAc, error) {

	msg := &e2smcommoniesv1.EUTraTAc{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEUTraTAc() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGlobalenGnbID(pLmnIdentity *e2smcommoniesv1.Plmnidentity, enGNbID *e2smcommoniesv1.EnGNbID) (*e2smcommoniesv1.GlobalenGnbID, error) {

	msg := &e2smcommoniesv1.GlobalenGnbID{}
	msg.PLmnIdentity = pLmnIdentity
	msg.EnGNbId = enGNbID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalenGnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNgenbCUUEW1ApID(value int64) (*e2smcommoniesv1.NgenbCUUEW1ApID, error) {

	msg := &e2smcommoniesv1.NgenbCUUEW1ApID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgenbCUUEW1ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNgenbDUID(value int64) (*e2smcommoniesv1.NgenbDUID, error) {

	msg := &e2smcommoniesv1.NgenbDUID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgenbDUID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAmfpointer(value *asn1.BitString) (*e2smcommoniesv1.Amfpointer, error) {

	msg := &e2smcommoniesv1.Amfpointer{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAmfpointer() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAmfregionID(value *asn1.BitString) (*e2smcommoniesv1.AmfregionId, error) {

	msg := &e2smcommoniesv1.AmfregionId{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAmfregionID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAmfsetID(value *asn1.BitString) (*e2smcommoniesv1.AmfsetId, error) {

	msg := &e2smcommoniesv1.AmfsetId{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAmfsetID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAmfUENGapID(value int64) (*e2smcommoniesv1.AmfUENGapID, error) {

	msg := &e2smcommoniesv1.AmfUENGapID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAmfUENGapID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEutracellIdentity(value *asn1.BitString) (*e2smcommoniesv1.EutracellIdentity, error) {

	msg := &e2smcommoniesv1.EutracellIdentity{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEutracellIdentity() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEutraCGi(pLmnidentity *e2smcommoniesv1.Plmnidentity, eUtracellIdentity *e2smcommoniesv1.EutracellIdentity) (*e2smcommoniesv1.EutraCGi, error) {

	msg := &e2smcommoniesv1.EutraCGi{}
	msg.PLmnidentity = pLmnidentity
	msg.EUtracellIdentity = eUtracellIdentity

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEutraCGi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateFiveQi(value int32) (*e2smcommoniesv1.FiveQi, error) {

	msg := &e2smcommoniesv1.FiveQi{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateFiveQi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGlobalGnbID(pLmnidentity *e2smcommoniesv1.Plmnidentity, gNbID *e2smcommoniesv1.GnbID) (*e2smcommoniesv1.GlobalGnbID, error) {

	msg := &e2smcommoniesv1.GlobalGnbID{}
	msg.PLmnidentity = pLmnidentity
	msg.GNbId = gNbID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalGnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGlobalNgEnbID(pLmnidentity *e2smcommoniesv1.Plmnidentity, ngEnbID *e2smcommoniesv1.NgEnbID) (*e2smcommoniesv1.GlobalNgEnbID, error) {

	msg := &e2smcommoniesv1.GlobalNgEnbID{}
	msg.PLmnidentity = pLmnidentity
	msg.NgEnbId = ngEnbID

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalNgEnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGuami(pLmnidentity *e2smcommoniesv1.Plmnidentity, aMfregionID *e2smcommoniesv1.AmfregionId, aMfsetID *e2smcommoniesv1.AmfsetId, aMfpointer *e2smcommoniesv1.Amfpointer) (*e2smcommoniesv1.Guami, error) {

	msg := &e2smcommoniesv1.Guami{}
	msg.PLmnidentity = pLmnidentity
	msg.AMfregionId = aMfregionID
	msg.AMfsetId = aMfsetID
	msg.AMfpointer = aMfpointer

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGuami() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateIndexToRfsp(value int32) (*e2smcommoniesv1.IndexToRfsp, error) {

	msg := &e2smcommoniesv1.IndexToRfsp{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateIndexToRfsp() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrcellIdentity(value *asn1.BitString) (*e2smcommoniesv1.NrcellIdentity, error) {

	msg := &e2smcommoniesv1.NrcellIdentity{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrcellIdentity() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrCGi(pLmnidentity *e2smcommoniesv1.Plmnidentity, nRcellIdentity *e2smcommoniesv1.NrcellIdentity) (*e2smcommoniesv1.NrCGi, error) {

	msg := &e2smcommoniesv1.NrCGi{}
	msg.PLmnidentity = pLmnidentity
	msg.NRcellIdentity = nRcellIdentity

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrCGi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreatePlmnidentity(value []byte) (*e2smcommoniesv1.Plmnidentity, error) {

	msg := &e2smcommoniesv1.Plmnidentity{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreatePlmnidentity() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateQosFlowIDentifier(value int32) (*e2smcommoniesv1.QosFlowIdentifier, error) {

	msg := &e2smcommoniesv1.QosFlowIdentifier{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateQosFlowIDentifier() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSd(value []byte) (*e2smcommoniesv1.Sd, error) {

	msg := &e2smcommoniesv1.Sd{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSd() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSst(value []byte) (*e2smcommoniesv1.Sst, error) {

	msg := &e2smcommoniesv1.Sst{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSst() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNgRAnnodeUexnApID(value int64) (*e2smcommoniesv1.NgRAnnodeUexnApid, error) {

	msg := &e2smcommoniesv1.NgRAnnodeUexnApid{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgRAnnodeUexnApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbCUCPUEE1ApID(value int64) (*e2smcommoniesv1.GnbCUCPUEE1ApID, error) {

	msg := &e2smcommoniesv1.GnbCUCPUEE1ApID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbCUCPUEE1ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbCUUPID(value int64) (*e2smcommoniesv1.GnbCUUPID, error) {

	msg := &e2smcommoniesv1.GnbCUUPID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbCUUPID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateFiveGsTAc(value []byte) (*e2smcommoniesv1.FiveGsTAc, error) {

	msg := &e2smcommoniesv1.FiveGsTAc{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateFiveGsTAc() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateFreqBandNrItem(freqBandIndicatorNr int32) (*e2smcommoniesv1.FreqBandNrItem, error) {

	msg := &e2smcommoniesv1.FreqBandNrItem{}
	msg.FreqBandIndicatorNr = freqBandIndicatorNr

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateFreqBandNrItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbCUUEF1ApID(value int64) (*e2smcommoniesv1.GnbCUUEF1ApID, error) {

	msg := &e2smcommoniesv1.GnbCUUEF1ApID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbCUUEF1ApID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbDUID(value int64) (*e2smcommoniesv1.GnbDUID, error) {

	msg := &e2smcommoniesv1.GnbDUID{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbDUID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrPCi(value int32) (*e2smcommoniesv1.NrPCi, error) {

	msg := &e2smcommoniesv1.NrPCi{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrPCi() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrARfcn(nRarfcn int32) (*e2smcommoniesv1.NrARfcn, error) {

	msg := &e2smcommoniesv1.NrARfcn{}
	msg.NRarfcn = nRarfcn

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrARfcn() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrfrequencyBandList(value []*e2smcommoniesv1.NrfrequencyBandItem) (*e2smcommoniesv1.NrfrequencyBandList, error) {

	msg := &e2smcommoniesv1.NrfrequencyBandList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrfrequencyBandList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrfrequencyBandItem(freqBandIndicatorNr int32, supportedSulbandList *e2smcommoniesv1.SupportedSulbandList) (*e2smcommoniesv1.NrfrequencyBandItem, error) {

	msg := &e2smcommoniesv1.NrfrequencyBandItem{}
	msg.FreqBandIndicatorNr = freqBandIndicatorNr
	msg.SupportedSulbandList = supportedSulbandList

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrfrequencyBandItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanueID(value []byte) (*e2smcommoniesv1.Ranueid, error) {

	msg := &e2smcommoniesv1.Ranueid{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanueID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSupportedSulbandList(value []*e2smcommoniesv1.SupportedSulfreqBandItem) (*e2smcommoniesv1.SupportedSulbandList, error) {

	msg := &e2smcommoniesv1.SupportedSulbandList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSupportedSulbandList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSupportedSulfreqBandItem(freqBandIndicatorNr int32) (*e2smcommoniesv1.SupportedSulfreqBandItem, error) {

	msg := &e2smcommoniesv1.SupportedSulfreqBandItem{}
	msg.FreqBandIndicatorNr = freqBandIndicatorNr

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSupportedSulfreqBandItem() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbID(value int64) (*e2smcccv1.GnbId, error) {

	msg := &e2smcccv1.GnbId{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbIDLength(value int32) (*e2smcccv1.GnbIdLength, error) {

	msg := &e2smcccv1.GnbIdLength{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbIDLength() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbName(value string) (*e2smcccv1.GnbName, error) {

	msg := &e2smcccv1.GnbName{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbDuID(value int64) (*e2smcccv1.GnbDuId, error) {

	msg := &e2smcccv1.GnbDuId{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbDuID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGnbCuUpID(value int64) (*e2smcccv1.GnbCuUpId, error) {

	msg := &e2smcccv1.GnbCuUpId{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbCuUpID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateSnssaiList(value []*e2smcommoniesv1.SNSsai) (*e2smcccv1.SnssaiList, error) {

	msg := &e2smcccv1.SnssaiList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateSnssaiList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreatePlmnIDList(value []*e2smcommoniesv1.Plmnidentity) (*e2smcccv1.PlmnIdList, error) {

	msg := &e2smcccv1.PlmnIdList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreatePlmnIDList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreatePlmnInfoList(value []*e2smcccv1.PlmnInfo) (*e2smcccv1.PlmnInfoList, error) {

	msg := &e2smcccv1.PlmnInfoList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreatePlmnInfoList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGgnbID(value []byte) (*e2smcccv1.GgnbId, error) {

	msg := &e2smcccv1.GgnbId{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGgnbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGenbID(value []byte) (*e2smcccv1.GenbId, error) {

	msg := &e2smcccv1.GenbId{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGenbID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGgnbIDList(value []*e2smcccv1.GgnbId) (*e2smcccv1.GgnbIdList, error) {

	msg := &e2smcccv1.GgnbIdList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGgnbIDList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateGenbIDList(value []*e2smcccv1.GenbId) (*e2smcccv1.GenbIdList, error) {

	msg := &e2smcccv1.GenbIdList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGenbIDList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrPci(value int32) (*e2smcccv1.NrPci, error) {

	msg := &e2smcccv1.NrPci{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrPci() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateNrTac(value int32) (*e2smcccv1.NrTac, error) {

	msg := &e2smcccv1.NrTac{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNrTac() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateTai(plmnID *e2smcommoniesv1.Plmnidentity, nrTac *e2smcccv1.NrTac) (*e2smcccv1.Tai, error) {

	msg := &e2smcccv1.Tai{}
	msg.PlmnId = plmnID
	msg.NrTac = nrTac

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateTai() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRrmPolicyMember(plmnID *e2smcommoniesv1.Plmnidentity, snssai *e2smcommoniesv1.SNSsai) (*e2smcccv1.RrmPolicyMember, error) {

	msg := &e2smcccv1.RrmPolicyMember{}
	msg.PlmnId = plmnID
	msg.Snssai = snssai

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRrmPolicyMember() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRrmPolicyMemberList(value []*e2smcccv1.RrmPolicyMember) (*e2smcccv1.RrmPolicyMemberList, error) {

	msg := &e2smcccv1.RrmPolicyMemberList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRrmPolicyMemberList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateOGnbCuCpFunction(gnbID *e2smcccv1.GnbId, gnbIDLength *e2smcccv1.GnbIdLength, gnbCuName *e2smcccv1.GnbName, plmnID *e2smcommoniesv1.Plmnidentity, x2ExcludeList *e2smcccv1.GenbIdList, xnExcludeList *e2smcccv1.GgnbIdList, x2IncludeList *e2smcccv1.GenbIdList, xnIncludeList *e2smcccv1.GgnbIdList, x2XnHoexcludeList *e2smcccv1.GenbIdList) (*e2smcccv1.OGnbCuCpFunction, error) {

	msg := &e2smcccv1.OGnbCuCpFunction{}
	msg.GnbId = gnbID
	msg.GnbIdLength = gnbIDLength
	msg.GnbCuName = gnbCuName
	msg.PlmnId = plmnID
	msg.X2ExcludeList = x2ExcludeList
	msg.XnExcludeList = xnExcludeList
	msg.X2IncludeList = x2IncludeList
	msg.XnIncludeList = xnIncludeList
	msg.X2XnHoexcludeList = x2XnHoexcludeList

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateOGnbCuCpFunction() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateOGnbCuUpFunction(gnbID *e2smcccv1.GnbId, gnbIDLength *e2smcccv1.GnbIdLength, gnbCuUpID *e2smcccv1.GnbCuUpId, plmnInfoList *e2smcccv1.PlmnInfoList) (*e2smcccv1.OGnbCuUpFunction, error) {

	msg := &e2smcccv1.OGnbCuUpFunction{}
	msg.GnbId = gnbID
	msg.GnbIdLength = gnbIDLength
	msg.GnbCuUpId = gnbCuUpID
	msg.PlmnInfoList = plmnInfoList

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateOGnbCuUpFunction() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateOGnbDuFunction(gnbDuID *e2smcccv1.GnbDuId, gnbDuName *e2smcccv1.GnbName, gnbID *e2smcccv1.GnbId, gnbIDLength *e2smcccv1.GnbIdLength) (*e2smcccv1.OGnbDuFunction, error) {

	msg := &e2smcccv1.OGnbDuFunction{}
	msg.GnbDuId = gnbDuID
	msg.GnbDuName = gnbDuName
	msg.GnbId = gnbID
	msg.GnbIdLength = gnbIDLength

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateOGnbDuFunction() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateCellLocalID(value int32) (*e2smcccv1.CellLocalId, error) {

	msg := &e2smcccv1.CellLocalId{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellLocalID() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateONrCellCu(cellLocalID *e2smcccv1.CellLocalId, plmnInfoList *e2smcccv1.PlmnInfoList) (*e2smcccv1.ONrCellCu, error) {

	msg := &e2smcccv1.ONrCellCu{}
	msg.CellLocalId = cellLocalID
	msg.PlmnInfoList = plmnInfoList

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateONrCellCu() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateBwpList(value []*e2smcccv1.OBwp) (*e2smcccv1.BwpList, error) {

	msg := &e2smcccv1.BwpList{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateBwpList() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateONrCellDu(cellLocalID *e2smcccv1.CellLocalId, cellState e2smcccv1.CellState, plmnInfoList *e2smcccv1.PlmnInfoList, nrPci *e2smcccv1.NrPci, nrTac *e2smcccv1.NrTac, arfcnDl int32, arfcnUl int32, arfcnSul int32, bSchannelBwDl int32, ssbFrequency int32, ssbPeriodicity e2smcccv1.SsbPeriodicity, ssbSubCarrierSpacing e2smcccv1.SsbSubCarrierSpacing, ssbOffset int32, ssbDuration e2smcccv1.SsbDuration, bSchannelBwUl int32, bSchannelBwSul int32, bwpList *e2smcccv1.BwpList) (*e2smcccv1.ONrCellDu, error) {

	msg := &e2smcccv1.ONrCellDu{}
	msg.CellLocalId = cellLocalID
	msg.CellState = cellState
	msg.PlmnInfoList = plmnInfoList
	msg.NrPci = nrPci
	msg.NrTac = nrTac
	msg.ArfcnDl = arfcnDl
	msg.ArfcnUl = arfcnUl
	msg.ArfcnSul = arfcnSul
	msg.BSchannelBwDl = bSchannelBwDl
	msg.SsbFrequency = ssbFrequency
	msg.SsbPeriodicity = ssbPeriodicity
	msg.SsbSubCarrierSpacing = ssbSubCarrierSpacing
	msg.SsbOffset = ssbOffset
	msg.SsbDuration = ssbDuration
	msg.BSchannelBwUl = bSchannelBwUl
	msg.BSchannelBwSul = bSchannelBwSul
	msg.BwpList = bwpList

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateONrCellDu() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateORRmpolicyRatio(resourceType e2smcccv1.ResourceType, schedulerType e2smcccv1.SchedulerType, rRmpolicyMemberList *e2smcccv1.RrmPolicyMemberList, rRmpolicyMaxRatio int32, rRmpolicyMinRatio int32, rRmpolicyDedicatedRatio int32) (*e2smcccv1.ORRmpolicyRatio, error) {

	msg := &e2smcccv1.ORRmpolicyRatio{}
	msg.ResourceType = resourceType
	msg.SchedulerType = schedulerType
	msg.RRmpolicyMemberList = rRmpolicyMemberList
	msg.RRmpolicyMaxRatio = rRmpolicyMaxRatio
	msg.RRmpolicyMinRatio = rRmpolicyMinRatio
	msg.RRmpolicyDedicatedRatio = rRmpolicyDedicatedRatio

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateORRmpolicyRatio() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateOBwp(bwpContext e2smcccv1.BwpContext, isInitialBwp e2smcccv1.IsInitialBwp, subCarrierSpacing e2smcccv1.SubCarrierSpacing, cyclicPrefix e2smcccv1.CyclicPrefix, startRb int32, numberOfRbs int32) (*e2smcccv1.OBwp, error) {

	msg := &e2smcccv1.OBwp{}
	msg.BwpContext = bwpContext
	msg.IsInitialBwp = isInitialBwp
	msg.SubCarrierSpacing = subCarrierSpacing
	msg.CyclicPrefix = cyclicPrefix
	msg.StartRb = startRb
	msg.NumberOfRbs = numberOfRbs

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateOBwp() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcRIcIndicationHeader(indicationHeaderFormat *e2smcccv1.IndicationHeaderFormat) (*e2smcccv1.E2SmCCcRIcIndicationHeader, error) {

	msg := &e2smcccv1.E2SmCCcRIcIndicationHeader{}
	msg.IndicationHeaderFormat = indicationHeaderFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcIndicationHeader() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcIndicationHeaderFormat1(indicationReason e2smcccv1.IndicationReason, eventTime []byte) (*e2smcccv1.E2SmCCcIndicationHeaderFormat1, error) {

	msg := &e2smcccv1.E2SmCCcIndicationHeaderFormat1{}
	msg.IndicationReason = indicationReason
	msg.EventTime = eventTime

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcIndicationHeaderFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcRIcIndicationMessage(indicationMessageFormat *e2smcccv1.IndicationMessageFormat) (*e2smcccv1.E2SmCCcRIcIndicationMessage, error) {

	msg := &e2smcccv1.E2SmCCcRIcIndicationMessage{}
	msg.IndicationMessageFormat = indicationMessageFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcIndicationMessage() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcIndicationMessageFormat1(listOfConfigurationStructuresReported *e2smcccv1.ListOfConfigurationsReported) (*e2smcccv1.E2SmCCcIndicationMessageFormat1, error) {

	msg := &e2smcccv1.E2SmCCcIndicationMessageFormat1{}
	msg.ListOfConfigurationStructuresReported = listOfConfigurationStructuresReported

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcIndicationMessageFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfConfigurationsReported(value []*e2smcccv1.ConfigurationStructure) (*e2smcccv1.ListOfConfigurationsReported, error) {

	msg := &e2smcccv1.ListOfConfigurationsReported{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfConfigurationsReported() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateRanConfigurationStructureName(value []byte) (*e2smcccv1.RanConfigurationStructureName, error) {

	msg := &e2smcccv1.RanConfigurationStructureName{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRanConfigurationStructureName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateValuesOfAttributes(ranConfigurationStructure *e2smcccv1.E2SmCCcRAnConfigurationStructure) (*e2smcccv1.ValuesOfAttributes, error) {

	msg := &e2smcccv1.ValuesOfAttributes{}
	msg.RanConfigurationStructure = ranConfigurationStructure

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateValuesOfAttributes() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcIndicationMessageFormat2(listOfCellsReported *e2smcccv1.ListOfCellsReported) (*e2smcccv1.E2SmCCcIndicationMessageFormat2, error) {

	msg := &e2smcccv1.E2SmCCcIndicationMessageFormat2{}
	msg.ListOfCellsReported = listOfCellsReported

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcIndicationMessageFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfCellsReported(value []*e2smcccv1.CellReported) (*e2smcccv1.ListOfCellsReported, error) {

	msg := &e2smcccv1.ListOfCellsReported{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfCellsReported() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateCellReported(cellGlobalID *e2smcccv1.CellGlobalId, listOfConfigurationStructuresReported *e2smcccv1.ListOfConfigurationsReported) (*e2smcccv1.CellReported, error) {

	msg := &e2smcccv1.CellReported{}
	msg.CellGlobalId = cellGlobalID
	msg.ListOfConfigurationStructuresReported = listOfConfigurationStructuresReported

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellReported() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcRIcControlHeader(controlHeaderFormat *e2smcccv1.ControlHeaderFormat) (*e2smcccv1.E2SmCCcRIcControlHeader, error) {

	msg := &e2smcccv1.E2SmCCcRIcControlHeader{}
	msg.ControlHeaderFormat = controlHeaderFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcControlHeader() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcControlHeaderFormat1(ricStyleType *e2smcommoniesv1.RicStyleType) (*e2smcccv1.E2SmCCcControlHeaderFormat1, error) {

	msg := &e2smcccv1.E2SmCCcControlHeaderFormat1{}
	msg.RicStyleType = ricStyleType

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcControlHeaderFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcRIcControlMessage(controlMessageFormat *e2smcccv1.ControlMessageFormat) (*e2smcccv1.E2SmCCcRIcControlMessage, error) {

	msg := &e2smcccv1.E2SmCCcRIcControlMessage{}
	msg.ControlMessageFormat = controlMessageFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcControlMessage() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcControlMessageFormat1(listOfConfigurationStructures *e2smcccv1.ListOfConfigurationStructures) (*e2smcccv1.E2SmCCcControlMessageFormat1, error) {

	msg := &e2smcccv1.E2SmCCcControlMessageFormat1{}
	msg.ListOfConfigurationStructures = listOfConfigurationStructures

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcControlMessageFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfConfigurationStructures(value []*e2smcccv1.ConfigurationStructureWrite) (*e2smcccv1.ListOfConfigurationStructures, error) {

	msg := &e2smcccv1.ListOfConfigurationStructures{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfConfigurationStructures() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateConfigurationStructureWrite(ranConfigurationStructureName *e2smcccv1.RanConfigurationStructureName, oldValuesOfAttributes *e2smcccv1.ValuesOfAttributes, newValuesOfAttributes *e2smcccv1.ValuesOfAttributes) (*e2smcccv1.ConfigurationStructureWrite, error) {

	msg := &e2smcccv1.ConfigurationStructureWrite{}
	msg.RanConfigurationStructureName = ranConfigurationStructureName
	msg.OldValuesOfAttributes = oldValuesOfAttributes
	msg.NewValuesOfAttributes = newValuesOfAttributes

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateConfigurationStructureWrite() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcControlMessageFormat2(listOfCellsReported *e2smcccv1.ListOfCells) (*e2smcccv1.E2SmCCcControlMessageFormat2, error) {

	msg := &e2smcccv1.E2SmCCcControlMessageFormat2{}
	msg.ListOfCellsReported = listOfCellsReported

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcControlMessageFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfCells(value []*e2smcccv1.CellControlled) (*e2smcccv1.ListOfCells, error) {

	msg := &e2smcccv1.ListOfCells{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfCells() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateCellControlled(cellGlobalID *e2smcccv1.CellGlobalId, listOfConfigurationStructures *e2smcccv1.ListOfConfigurationStructures) (*e2smcccv1.CellControlled, error) {

	msg := &e2smcccv1.CellControlled{}
	msg.CellGlobalId = cellGlobalID
	msg.ListOfConfigurationStructures = listOfConfigurationStructures

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellControlled() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfSupportedRanconfigurationStructures(value []*e2smcccv1.RanconfigurationStructure) (*e2smcccv1.ListOfSupportedRanconfigurationStructures, error) {

	msg := &e2smcccv1.ListOfSupportedRanconfigurationStructures{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfSupportedRanconfigurationStructures() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfSupportedAttributes(value []*e2smcccv1.Attribute) (*e2smcccv1.ListOfSupportedAttributes, error) {

	msg := &e2smcccv1.ListOfSupportedAttributes{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfSupportedAttributes() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAttribute(attributeName *e2smcccv1.AttributeName, supportedServices *e2smcccv1.Ricservices) (*e2smcccv1.Attribute, error) {

	msg := &e2smcccv1.Attribute{}
	msg.AttributeName = attributeName
	msg.SupportedServices = supportedServices

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAttribute() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTrigger(listOfSupportedEventTriggerStyles *e2smcccv1.ListOfSupportedEventTriggerStyles) (*e2smcccv1.EventTrigger, error) {

	msg := &e2smcccv1.EventTrigger{}
	msg.ListOfSupportedEventTriggerStyles = listOfSupportedEventTriggerStyles

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTrigger() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfSupportedEventTriggerStyles(value []*e2smcccv1.EventTriggerStyle) (*e2smcccv1.ListOfSupportedEventTriggerStyles, error) {

	msg := &e2smcccv1.ListOfSupportedEventTriggerStyles{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfSupportedEventTriggerStyles() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerStyle(eventTriggerStyleType *e2smcommoniesv1.RicStyleType, eventTriggerStyleName *e2smcommoniesv1.RicStyleName, eventTriggerFormatType *e2smcommoniesv1.RicFormatType) (*e2smcccv1.EventTriggerStyle, error) {

	msg := &e2smcccv1.EventTriggerStyle{}
	msg.EventTriggerStyleType = eventTriggerStyleType
	msg.EventTriggerStyleName = eventTriggerStyleName
	msg.EventTriggerFormatType = eventTriggerFormatType

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerStyle() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateReportService(listOfSupportedReportStyles *e2smcccv1.ListOfSupportedReportStyles) (*e2smcccv1.ReportService, error) {

	msg := &e2smcccv1.ReportService{}
	msg.ListOfSupportedReportStyles = listOfSupportedReportStyles

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateReportService() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfSupportedReportStyles(value []*e2smcccv1.ReportStyle) (*e2smcccv1.ListOfSupportedReportStyles, error) {

	msg := &e2smcccv1.ListOfSupportedReportStyles{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfSupportedReportStyles() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateReportStyle(reportServiceStyleType *e2smcommoniesv1.RicStyleType, reportServiceStyleName *e2smcommoniesv1.RicStyleName, listOfSupportedEventTriggerStylesForReportStyle *e2smcccv1.ListOfSupportedEventTriggerStylesForReportStyle, reportServiceActionDefinitionFormatType *e2smcommoniesv1.RicFormatType, reportServiceIndicationHeaderFormatType *e2smcommoniesv1.RicFormatType, reportServiceIndicationMessageFormatType *e2smcommoniesv1.RicFormatType) (*e2smcccv1.ReportStyle, error) {

	msg := &e2smcccv1.ReportStyle{}
	msg.ReportServiceStyleType = reportServiceStyleType
	msg.ReportServiceStyleName = reportServiceStyleName
	msg.ListOfSupportedEventTriggerStylesForReportStyle = listOfSupportedEventTriggerStylesForReportStyle
	msg.ReportServiceActionDefinitionFormatType = reportServiceActionDefinitionFormatType
	msg.ReportServiceIndicationHeaderFormatType = reportServiceIndicationHeaderFormatType
	msg.ReportServiceIndicationMessageFormatType = reportServiceIndicationMessageFormatType

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateReportStyle() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfSupportedEventTriggerStylesForReportStyle(value []*e2smcccv1.EventTriggerStyleType) (*e2smcccv1.ListOfSupportedEventTriggerStylesForReportStyle, error) {

	msg := &e2smcccv1.ListOfSupportedEventTriggerStylesForReportStyle{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfSupportedEventTriggerStylesForReportStyle() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateEventTriggerStyleType(eventTriggerStyleType *e2smcommoniesv1.RicStyleType) (*e2smcccv1.EventTriggerStyleType, error) {

	msg := &e2smcccv1.EventTriggerStyleType{}
	msg.EventTriggerStyleType = eventTriggerStyleType

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerStyleType() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateInsertService(value Empty) (*e2smcccv1.InsertService, error) {

	msg := &e2smcccv1.InsertService{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInsertService() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateControlService(listOfSupportedControlStyles *e2smcccv1.ListOfSupportedControlStyles) (*e2smcccv1.ControlService, error) {

	msg := &e2smcccv1.ControlService{}
	msg.ListOfSupportedControlStyles = listOfSupportedControlStyles

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateControlService() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfSupportedControlStyles(value []*e2smcccv1.ControlStyle) (*e2smcccv1.ListOfSupportedControlStyles, error) {

	msg := &e2smcccv1.ListOfSupportedControlStyles{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfSupportedControlStyles() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreatePolicyService(value Empty) (*e2smcccv1.PolicyService, error) {

	msg := &e2smcccv1.PolicyService{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreatePolicyService() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfCellsForRanfunctionDefinition(value []*e2smcccv1.CellForRanfunctionDefinition) (*e2smcccv1.ListOfCellsForRanfunctionDefinition, error) {

	msg := &e2smcccv1.ListOfCellsForRanfunctionDefinition{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfCellsForRanfunctionDefinition() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcRIceventTriggerDefinition(eventTriggerDefinitionFormat *e2smcccv1.EventTriggerDefinitionFormat) (*e2smcccv1.E2SmCCcRIceventTriggerDefinition, error) {

	msg := &e2smcccv1.E2SmCCcRIceventTriggerDefinition{}
	msg.EventTriggerDefinitionFormat = eventTriggerDefinitionFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIceventTriggerDefinition() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcEventTriggerDefinitionFormat1(listOfNodeLevelConfigurationStructuresForEventTrigger *e2smcccv1.ListOfRanconfigurationStructuresForEventTrigger) (*e2smcccv1.E2SmCCcEventTriggerDefinitionFormat1, error) {

	msg := &e2smcccv1.E2SmCCcEventTriggerDefinitionFormat1{}
	msg.ListOfNodeLevelConfigurationStructuresForEventTrigger = listOfNodeLevelConfigurationStructuresForEventTrigger

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcEventTriggerDefinitionFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfRanconfigurationStructuresForEventTrigger(value []*e2smcccv1.RanconfigurationStructureForEventTrigger) (*e2smcccv1.ListOfRanconfigurationStructuresForEventTrigger, error) {

	msg := &e2smcccv1.ListOfRanconfigurationStructuresForEventTrigger{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfRanconfigurationStructuresForEventTrigger() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfAttributes(value []*e2smcccv1.AttributeName) (*e2smcccv1.ListOfAttributes, error) {

	msg := &e2smcccv1.ListOfAttributes{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfAttributes() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcEventTriggerDefinitionFormat2(listOfCellLevelConfigurationStructuresForEventTrigger *e2smcccv1.ListOfCellLevelConfigurationStructuresForEventTrigger) (*e2smcccv1.E2SmCCcEventTriggerDefinitionFormat2, error) {

	msg := &e2smcccv1.E2SmCCcEventTriggerDefinitionFormat2{}
	msg.ListOfCellLevelConfigurationStructuresForEventTrigger = listOfCellLevelConfigurationStructuresForEventTrigger

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcEventTriggerDefinitionFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfCellLevelConfigurationStructuresForEventTrigger(value []*e2smcccv1.CellLevelConfigurationStructureForEventTrigger) (*e2smcccv1.ListOfCellLevelConfigurationStructuresForEventTrigger, error) {

	msg := &e2smcccv1.ListOfCellLevelConfigurationStructuresForEventTrigger{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfCellLevelConfigurationStructuresForEventTrigger() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateCellLevelConfigurationStructureForEventTrigger(cellGlobalID *e2smcccv1.CellGlobalId, listOfRanconfigurationStructuresForEventTrigger *e2smcccv1.ListOfRanconfigurationStructuresForEventTrigger) (*e2smcccv1.CellLevelConfigurationStructureForEventTrigger, error) {

	msg := &e2smcccv1.CellLevelConfigurationStructureForEventTrigger{}
	msg.CellGlobalId = cellGlobalID
	msg.ListOfRanconfigurationStructuresForEventTrigger = listOfRanconfigurationStructuresForEventTrigger

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellLevelConfigurationStructureForEventTrigger() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateAttributeName(value []byte) (*e2smcccv1.AttributeName, error) {

	msg := &e2smcccv1.AttributeName{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateAttributeName() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcEventTriggerDefinitionFormat3(period int32) (*e2smcccv1.E2SmCCcEventTriggerDefinitionFormat3, error) {

	msg := &e2smcccv1.E2SmCCcEventTriggerDefinitionFormat3{}
	msg.Period = period

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcEventTriggerDefinitionFormat3() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcRIcactionDefinition(ricStyleType *e2smcommoniesv1.RicStyleType, actionDefinitionFormat *e2smcccv1.ActionDefinitionFormat) (*e2smcccv1.E2SmCCcRIcactionDefinition, error) {

	msg := &e2smcccv1.E2SmCCcRIcactionDefinition{}
	msg.RicStyleType = ricStyleType
	msg.ActionDefinitionFormat = actionDefinitionFormat

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRIcactionDefinition() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcActionDefinitionFormat1(listOfNodeLevelRanconfigurationStructuresForAdf *e2smcccv1.ListOfRanconfigurationStructuresForAdf) (*e2smcccv1.E2SmCCcActionDefinitionFormat1, error) {

	msg := &e2smcccv1.E2SmCCcActionDefinitionFormat1{}
	msg.ListOfNodeLevelRanconfigurationStructuresForAdf = listOfNodeLevelRanconfigurationStructuresForAdf

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcActionDefinitionFormat1() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfRanconfigurationStructuresForAdf(value []*e2smcccv1.RanconfigurationStructureForAdf) (*e2smcccv1.ListOfRanconfigurationStructuresForAdf, error) {

	msg := &e2smcccv1.ListOfRanconfigurationStructuresForAdf{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfRanconfigurationStructuresForAdf() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateE2SmCCcActionDefinitionFormat2(listOfCellConfigurationsToBeReportedForAdf *e2smcccv1.ListOfCellConfigurationsToBeReportedForAdf) (*e2smcccv1.E2SmCCcActionDefinitionFormat2, error) {

	msg := &e2smcccv1.E2SmCCcActionDefinitionFormat2{}
	msg.ListOfCellConfigurationsToBeReportedForAdf = listOfCellConfigurationsToBeReportedForAdf

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcActionDefinitionFormat2() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateListOfCellConfigurationsToBeReportedForAdf(value []*e2smcccv1.CellConfigurationToBeReportedForAdf) (*e2smcccv1.ListOfCellConfigurationsToBeReportedForAdf, error) {

	msg := &e2smcccv1.ListOfCellConfigurationsToBeReportedForAdf{}
	msg.Value = value

	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateListOfCellConfigurationsToBeReportedForAdf() error validating PDU %s", err.Error())
	}

	return msg, nil
}

func CreateCgiNRCgi(nRCgi *e2smcommoniesv1.NrCGi) (*e2smcommoniesv1.Cgi, error) {

	item := &e2smcommoniesv1.Cgi{
		Cgi: &e2smcommoniesv1.Cgi_NRCgi{
			NRCgi: nRCgi,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCgiNRCgi() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateCgiEUtraCgi(eUtraCgi *e2smcommoniesv1.EutraCGi) (*e2smcommoniesv1.Cgi, error) {

	item := &e2smcommoniesv1.Cgi{
		Cgi: &e2smcommoniesv1.Cgi_EUtraCgi{
			EUtraCgi: eUtraCgi,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCgiEUtraCgi() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateCoreCpIDFiveGc(fiveGc *e2smcommoniesv1.Guami) (*e2smcommoniesv1.CoreCpid, error) {

	item := &e2smcommoniesv1.CoreCpid{
		CoreCpid: &e2smcommoniesv1.CoreCpid_FiveGc{
			FiveGc: fiveGc,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCoreCpIDFiveGc() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateCoreCpIDEPc(ePc *e2smcommoniesv1.Gummei) (*e2smcommoniesv1.CoreCpid, error) {

	item := &e2smcommoniesv1.CoreCpid{
		CoreCpid: &e2smcommoniesv1.CoreCpid_EPc{
			EPc: ePc,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCoreCpIDEPc() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateInterfaceIDentifierNG(nG *e2smcommoniesv1.InterfaceIdNG) (*e2smcommoniesv1.InterfaceIdentifier, error) {

	item := &e2smcommoniesv1.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommoniesv1.InterfaceIdentifier_NG{
			NG: nG,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDentifierNG() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateInterfaceIDentifierXN(xN *e2smcommoniesv1.InterfaceIdXn) (*e2smcommoniesv1.InterfaceIdentifier, error) {

	item := &e2smcommoniesv1.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommoniesv1.InterfaceIdentifier_XN{
			XN: xN,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDentifierXN() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateInterfaceIDentifierF1(f1 *e2smcommoniesv1.InterfaceIdF1) (*e2smcommoniesv1.InterfaceIdentifier, error) {

	item := &e2smcommoniesv1.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommoniesv1.InterfaceIdentifier_F1{
			F1: f1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDentifierF1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateInterfaceIDentifierE1(e1 *e2smcommoniesv1.InterfaceIdE1) (*e2smcommoniesv1.InterfaceIdentifier, error) {

	item := &e2smcommoniesv1.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommoniesv1.InterfaceIdentifier_E1{
			E1: e1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDentifierE1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateInterfaceIDentifierS1(s1 *e2smcommoniesv1.InterfaceIdS1) (*e2smcommoniesv1.InterfaceIdentifier, error) {

	item := &e2smcommoniesv1.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommoniesv1.InterfaceIdentifier_S1{
			S1: s1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDentifierS1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateInterfaceIDentifierX2(x2 *e2smcommoniesv1.InterfaceIdX2) (*e2smcommoniesv1.InterfaceIdentifier, error) {

	item := &e2smcommoniesv1.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommoniesv1.InterfaceIdentifier_X2{
			X2: x2,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDentifierX2() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateInterfaceIDentifierW1(w1 *e2smcommoniesv1.InterfaceIdW1) (*e2smcommoniesv1.InterfaceIdentifier, error) {

	item := &e2smcommoniesv1.InterfaceIdentifier{
		InterfaceIdentifier: &e2smcommoniesv1.InterfaceIdentifier_W1{
			W1: w1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateInterfaceIDentifierW1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateNodeTypeInterfaceIDX2GlobalENbID(globalENbID *e2smcommoniesv1.GlobalEnbID) (*e2smcommoniesv1.NodeTypeInterfaceIdX2, error) {

	item := &e2smcommoniesv1.NodeTypeInterfaceIdX2{
		NodeTypeInterfaceIdX2: &e2smcommoniesv1.NodeTypeInterfaceIdX2_GlobalENbId{
			GlobalENbId: globalENbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNodeTypeInterfaceIDX2GlobalENbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateNodeTypeInterfaceIDX2GlobalEnGNbID(globalEnGNbID *e2smcommoniesv1.GlobalenGnbID) (*e2smcommoniesv1.NodeTypeInterfaceIdX2, error) {

	item := &e2smcommoniesv1.NodeTypeInterfaceIdX2{
		NodeTypeInterfaceIdX2: &e2smcommoniesv1.NodeTypeInterfaceIdX2_GlobalEnGNbId{
			GlobalEnGNbId: globalEnGNbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNodeTypeInterfaceIDX2GlobalEnGNbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateGroupIDFiveGc(fiveGc *e2smcommoniesv1.FiveQi) (*e2smcommoniesv1.GroupId, error) {

	item := &e2smcommoniesv1.GroupId{
		GroupId: &e2smcommoniesv1.GroupId_FiveGc{
			FiveGc: fiveGc,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGroupIDFiveGc() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateGroupIDEPc(ePc *e2smcommoniesv1.Qci) (*e2smcommoniesv1.GroupId, error) {

	item := &e2smcommoniesv1.GroupId{
		GroupId: &e2smcommoniesv1.GroupId_EPc{
			EPc: ePc,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGroupIDEPc() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateQoSIDFiveGc(fiveGc *e2smcommoniesv1.FiveQi) (*e2smcommoniesv1.QoSid, error) {

	item := &e2smcommoniesv1.QoSid{
		QoSid: &e2smcommoniesv1.QoSid_FiveGc{
			FiveGc: fiveGc,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateQoSIDFiveGc() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateQoSIDEPc(ePc *e2smcommoniesv1.Qci) (*e2smcommoniesv1.QoSid, error) {

	item := &e2smcommoniesv1.QoSid{
		QoSid: &e2smcommoniesv1.QoSid_EPc{
			EPc: ePc,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateQoSIDEPc() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateRrcTypeRrcMessageIDLTe(lTe e2smcommoniesv1.RrcclassLte) (*e2smcommoniesv1.RrcTypeRrcMessageId, error) {

	item := &e2smcommoniesv1.RrcTypeRrcMessageId{
		RrcTypeRrcMessageId: &e2smcommoniesv1.RrcTypeRrcMessageId_LTe{
			LTe: lTe,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRrcTypeRrcMessageIDLTe() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateRrcTypeRrcMessageIDNR(nR e2smcommoniesv1.RrcclassNR) (*e2smcommoniesv1.RrcTypeRrcMessageId, error) {

	item := &e2smcommoniesv1.RrcTypeRrcMessageId{
		RrcTypeRrcMessageId: &e2smcommoniesv1.RrcTypeRrcMessageId_NR{
			NR: nR,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateRrcTypeRrcMessageIDNR() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateServingCellARfcnNR(nR *e2smcommoniesv1.NrARfcn) (*e2smcommoniesv1.ServingCellARfcn, error) {

	item := &e2smcommoniesv1.ServingCellARfcn{
		ServingCellARfcn: &e2smcommoniesv1.ServingCellARfcn_NR{
			NR: nR,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellARfcnNR() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateServingCellARfcnEUtra(eUtra *e2smcommoniesv1.EUTraARfcn) (*e2smcommoniesv1.ServingCellARfcn, error) {

	item := &e2smcommoniesv1.ServingCellARfcn{
		ServingCellARfcn: &e2smcommoniesv1.ServingCellARfcn_EUtra{
			EUtra: eUtra,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellARfcnEUtra() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateServingCellPCiNR(nR *e2smcommoniesv1.NrPCi) (*e2smcommoniesv1.ServingCellPCi, error) {

	item := &e2smcommoniesv1.ServingCellPCi{
		ServingCellPCi: &e2smcommoniesv1.ServingCellPCi_NR{
			NR: nR,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellPCiNR() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateServingCellPCiEUtra(eUtra *e2smcommoniesv1.EUTraPCi) (*e2smcommoniesv1.ServingCellPCi, error) {

	item := &e2smcommoniesv1.ServingCellPCi{
		ServingCellPCi: &e2smcommoniesv1.ServingCellPCi_EUtra{
			EUtra: eUtra,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateServingCellPCiEUtra() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateUeIDGNbUeID(gNbUeID *e2smcommoniesv1.UeidGNb) (*e2smcommoniesv1.Ueid, error) {

	item := &e2smcommoniesv1.Ueid{
		Ueid: &e2smcommoniesv1.Ueid_GNbUeid{
			GNbUeid: gNbUeID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGNbUeID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateUeIDGNbDuUeID(gNbDuUeID *e2smcommoniesv1.UeidGNbDU) (*e2smcommoniesv1.Ueid, error) {

	item := &e2smcommoniesv1.Ueid{
		Ueid: &e2smcommoniesv1.Ueid_GNbDuUeid{
			GNbDuUeid: gNbDuUeID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGNbDuUeID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateUeIDGNbCuUpUeID(gNbCuUpUeID *e2smcommoniesv1.UeidGNbCUUP) (*e2smcommoniesv1.Ueid, error) {

	item := &e2smcommoniesv1.Ueid{
		Ueid: &e2smcommoniesv1.Ueid_GNbCuUpUeid{
			GNbCuUpUeid: gNbCuUpUeID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGNbCuUpUeID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateUeIDNgENbUeID(ngENbUeID *e2smcommoniesv1.UeidNGENb) (*e2smcommoniesv1.Ueid, error) {

	item := &e2smcommoniesv1.Ueid{
		Ueid: &e2smcommoniesv1.Ueid_NgENbUeid{
			NgENbUeid: ngENbUeID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDNgENbUeID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateUeIDNgENbDuUeID(ngENbDuUeID *e2smcommoniesv1.UeidNGENbDU) (*e2smcommoniesv1.Ueid, error) {

	item := &e2smcommoniesv1.Ueid{
		Ueid: &e2smcommoniesv1.Ueid_NgENbDuUeid{
			NgENbDuUeid: ngENbDuUeID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDNgENbDuUeID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateUeIDEnGNbUeID(enGNbUeID *e2smcommoniesv1.UeidENGNb) (*e2smcommoniesv1.Ueid, error) {

	item := &e2smcommoniesv1.Ueid{
		Ueid: &e2smcommoniesv1.Ueid_EnGNbUeid{
			EnGNbUeid: enGNbUeID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDEnGNbUeID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateUeIDENbUeID(eNbUeID *e2smcommoniesv1.UeidENb) (*e2smcommoniesv1.Ueid, error) {

	item := &e2smcommoniesv1.Ueid{
		Ueid: &e2smcommoniesv1.Ueid_ENbUeid{
			ENbUeid: eNbUeID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDENbUeID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateEnbIDMacroENbID(macroENbID *asn1.BitString) (*e2smcommoniesv1.EnbID, error) {

	item := &e2smcommoniesv1.EnbID{
		EnbID: &e2smcommoniesv1.EnbID_MacroENbId{
			MacroENbId: macroENbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbIDMacroENbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateEnbIDHomeENbID(homeENbID *asn1.BitString) (*e2smcommoniesv1.EnbID, error) {

	item := &e2smcommoniesv1.EnbID{
		EnbID: &e2smcommoniesv1.EnbID_HomeENbId{
			HomeENbId: homeENbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbIDHomeENbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateEnbIDShortMacroENbID(shortMacroENbID *asn1.BitString) (*e2smcommoniesv1.EnbID, error) {

	item := &e2smcommoniesv1.EnbID{
		EnbID: &e2smcommoniesv1.EnbID_ShortMacroENbId{
			ShortMacroENbId: shortMacroENbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbIDShortMacroENbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateEnbIDLongMacroENbID(longMacroENbID *asn1.BitString) (*e2smcommoniesv1.EnbID, error) {

	item := &e2smcommoniesv1.EnbID{
		EnbID: &e2smcommoniesv1.EnbID_LongMacroENbId{
			LongMacroENbId: longMacroENbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnbIDLongMacroENbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateEnGNbIDEnGNbID(enGNbID *asn1.BitString) (*e2smcommoniesv1.EnGNbID, error) {

	item := &e2smcommoniesv1.EnGNbID{
		EnGNbID: &e2smcommoniesv1.EnGNbID_EnGNbId{
			EnGNbId: enGNbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEnGNbIDEnGNbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateGnbIDGNbID(gNbID *asn1.BitString) (*e2smcommoniesv1.GnbID, error) {

	item := &e2smcommoniesv1.GnbID{
		GnbID: &e2smcommoniesv1.GnbID_GNbId{
			GNbId: gNbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGnbIDGNbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateNgEnbIDMacroNgEnbID(macroNgEnbID *asn1.BitString) (*e2smcommoniesv1.NgEnbID, error) {

	item := &e2smcommoniesv1.NgEnbID{
		NgEnbID: &e2smcommoniesv1.NgEnbID_MacroNgEnbId{
			MacroNgEnbId: macroNgEnbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgEnbIDMacroNgEnbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateNgEnbIDShortMacroNgEnbID(shortMacroNgEnbID *asn1.BitString) (*e2smcommoniesv1.NgEnbID, error) {

	item := &e2smcommoniesv1.NgEnbID{
		NgEnbID: &e2smcommoniesv1.NgEnbID_ShortMacroNgEnbId{
			ShortMacroNgEnbId: shortMacroNgEnbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgEnbIDShortMacroNgEnbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateNgEnbIDLongMacroNgEnbID(longMacroNgEnbID *asn1.BitString) (*e2smcommoniesv1.NgEnbID, error) {

	item := &e2smcommoniesv1.NgEnbID{
		NgEnbID: &e2smcommoniesv1.NgEnbID_LongMacroNgEnbId{
			LongMacroNgEnbId: longMacroNgEnbID,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateNgEnbIDLongMacroNgEnbID() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateGlobalNgrannodeIDGNb(gNb *e2smcommoniesv1.GlobalGnbID) (*e2smcommoniesv1.GlobalNgrannodeId, error) {

	item := &e2smcommoniesv1.GlobalNgrannodeId{
		GlobalNgrannodeId: &e2smcommoniesv1.GlobalNgrannodeId_GNb{
			GNb: gNb,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalNgrannodeIDGNb() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateGlobalNgrannodeIDNgENb(ngENb *e2smcommoniesv1.GlobalNgEnbID) (*e2smcommoniesv1.GlobalNgrannodeId, error) {

	item := &e2smcommoniesv1.GlobalNgrannodeId{
		GlobalNgrannodeId: &e2smcommoniesv1.GlobalNgrannodeId_NgENb{
			NgENb: ngENb,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateGlobalNgrannodeIDNgENb() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateE2SmCCcRAnConfigurationStructureOGnbCuCpFunction(oGnbCuCpFunction *e2smcccv1.OGnbCuCpFunction) (*e2smcccv1.E2SmCCcRAnConfigurationStructure, error) {

	item := &e2smcccv1.E2SmCCcRAnConfigurationStructure{
		E2SmCCcRAnConfigurationStructure: &e2smcccv1.E2SmCCcRAnConfigurationStructure_OGnbCuCpFunction{
			OGnbCuCpFunction: oGnbCuCpFunction,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRAnConfigurationStructureOGnbCuCpFunction() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateE2SmCCcRAnConfigurationStructureOGnbCuUpFunction(oGnbCuUpFunction *e2smcccv1.OGnbCuUpFunction) (*e2smcccv1.E2SmCCcRAnConfigurationStructure, error) {

	item := &e2smcccv1.E2SmCCcRAnConfigurationStructure{
		E2SmCCcRAnConfigurationStructure: &e2smcccv1.E2SmCCcRAnConfigurationStructure_OGnbCuUpFunction{
			OGnbCuUpFunction: oGnbCuUpFunction,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRAnConfigurationStructureOGnbCuUpFunction() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateE2SmCCcRAnConfigurationStructureOGnbDuFunction(oGnbDuFunction *e2smcccv1.OGnbDuFunction) (*e2smcccv1.E2SmCCcRAnConfigurationStructure, error) {

	item := &e2smcccv1.E2SmCCcRAnConfigurationStructure{
		E2SmCCcRAnConfigurationStructure: &e2smcccv1.E2SmCCcRAnConfigurationStructure_OGnbDuFunction{
			OGnbDuFunction: oGnbDuFunction,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRAnConfigurationStructureOGnbDuFunction() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateE2SmCCcRAnConfigurationStructureONrCellCu(oNrCellCu *e2smcccv1.ONrCellCu) (*e2smcccv1.E2SmCCcRAnConfigurationStructure, error) {

	item := &e2smcccv1.E2SmCCcRAnConfigurationStructure{
		E2SmCCcRAnConfigurationStructure: &e2smcccv1.E2SmCCcRAnConfigurationStructure_ONrCellCu{
			ONrCellCu: oNrCellCu,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRAnConfigurationStructureONrCellCu() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateE2SmCCcRAnConfigurationStructureONrCellDu(oNrCellDu *e2smcccv1.ONrCellDu) (*e2smcccv1.E2SmCCcRAnConfigurationStructure, error) {

	item := &e2smcccv1.E2SmCCcRAnConfigurationStructure{
		E2SmCCcRAnConfigurationStructure: &e2smcccv1.E2SmCCcRAnConfigurationStructure_ONrCellDu{
			ONrCellDu: oNrCellDu,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRAnConfigurationStructureONrCellDu() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateE2SmCCcRAnConfigurationStructureORrmpolicyRatio(oRrmpolicyRatio *e2smcccv1.ORRmpolicyRatio) (*e2smcccv1.E2SmCCcRAnConfigurationStructure, error) {

	item := &e2smcccv1.E2SmCCcRAnConfigurationStructure{
		E2SmCCcRAnConfigurationStructure: &e2smcccv1.E2SmCCcRAnConfigurationStructure_ORrmpolicyRatio{
			ORrmpolicyRatio: oRrmpolicyRatio,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRAnConfigurationStructureORrmpolicyRatio() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateE2SmCCcRAnConfigurationStructureOBwp(oBwp *e2smcccv1.OBwp) (*e2smcccv1.E2SmCCcRAnConfigurationStructure, error) {

	item := &e2smcccv1.E2SmCCcRAnConfigurationStructure{
		E2SmCCcRAnConfigurationStructure: &e2smcccv1.E2SmCCcRAnConfigurationStructure_OBwp{
			OBwp: oBwp,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmCCcRAnConfigurationStructureOBwp() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateIndicationHeaderFormatE2SmCccIndicationHeaderFormat1(e2SmCccIndicationHeaderFormat1 *e2smcccv1.E2SmCCcIndicationHeaderFormat1) (*e2smcccv1.IndicationHeaderFormat, error) {

	item := &e2smcccv1.IndicationHeaderFormat{
		IndicationHeaderFormat: &e2smcccv1.IndicationHeaderFormat_E2SmCccIndicationHeaderFormat1{
			E2SmCccIndicationHeaderFormat1: e2SmCccIndicationHeaderFormat1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateIndicationHeaderFormatE2SmCccIndicationHeaderFormat1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateIndicationMessageFormatE2SmCccIndicationMessageFormat1(e2SmCccIndicationMessageFormat1 *e2smcccv1.E2SmCCcIndicationMessageFormat1) (*e2smcccv1.IndicationMessageFormat, error) {

	item := &e2smcccv1.IndicationMessageFormat{
		IndicationMessageFormat: &e2smcccv1.IndicationMessageFormat_E2SmCccIndicationMessageFormat1{
			E2SmCccIndicationMessageFormat1: e2SmCccIndicationMessageFormat1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateIndicationMessageFormatE2SmCccIndicationMessageFormat1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateIndicationMessageFormatE2SmCccIndicationMessageFormat2(e2SmCccIndicationMessageFormat2 *e2smcccv1.E2SmCCcIndicationMessageFormat2) (*e2smcccv1.IndicationMessageFormat, error) {

	item := &e2smcccv1.IndicationMessageFormat{
		IndicationMessageFormat: &e2smcccv1.IndicationMessageFormat_E2SmCccIndicationMessageFormat2{
			E2SmCccIndicationMessageFormat2: e2SmCccIndicationMessageFormat2,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateIndicationMessageFormatE2SmCccIndicationMessageFormat2() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateCellGlobalIDNRCgi(nRCgi *e2smcommoniesv1.NrCGi) (*e2smcccv1.CellGlobalId, error) {

	item := &e2smcccv1.CellGlobalId{
		CellGlobalId: &e2smcccv1.CellGlobalId_NRCgi{
			NRCgi: nRCgi,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellGlobalIDNRCgi() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateCellGlobalIDEUtraCgi(eUtraCgi *e2smcommoniesv1.EutraCGi) (*e2smcccv1.CellGlobalId, error) {

	item := &e2smcccv1.CellGlobalId{
		CellGlobalId: &e2smcccv1.CellGlobalId_EUtraCgi{
			EUtraCgi: eUtraCgi,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateCellGlobalIDEUtraCgi() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateControlHeaderFormatE2SmCccControlHeaderFormat1(e2SmCccControlHeaderFormat1 *e2smcccv1.E2SmCCcControlHeaderFormat1) (*e2smcccv1.ControlHeaderFormat, error) {

	item := &e2smcccv1.ControlHeaderFormat{
		ControlHeaderFormat: &e2smcccv1.ControlHeaderFormat_E2SmCccControlHeaderFormat1{
			E2SmCccControlHeaderFormat1: e2SmCccControlHeaderFormat1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateControlHeaderFormatE2SmCccControlHeaderFormat1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateControlMessageFormatE2SmCccControlMessageFormat1(e2SmCccControlMessageFormat1 *e2smcccv1.E2SmCCcControlMessageFormat1) (*e2smcccv1.ControlMessageFormat, error) {

	item := &e2smcccv1.ControlMessageFormat{
		ControlMessageFormat: &e2smcccv1.ControlMessageFormat_E2SmCccControlMessageFormat1{
			E2SmCccControlMessageFormat1: e2SmCccControlMessageFormat1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateControlMessageFormatE2SmCccControlMessageFormat1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateControlMessageFormatE2SmCccControlMessageFormat2(e2SmCccControlMessageFormat2 *e2smcccv1.E2SmCCcControlMessageFormat2) (*e2smcccv1.ControlMessageFormat, error) {

	item := &e2smcccv1.ControlMessageFormat{
		ControlMessageFormat: &e2smcccv1.ControlMessageFormat_E2SmCccControlMessageFormat2{
			E2SmCccControlMessageFormat2: e2SmCccControlMessageFormat2,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateControlMessageFormatE2SmCccControlMessageFormat2() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat1(e2SmCccEventTriggerDefinitionFormat1 *e2smcccv1.E2SmCCcEventTriggerDefinitionFormat1) (*e2smcccv1.EventTriggerDefinitionFormat, error) {

	item := &e2smcccv1.EventTriggerDefinitionFormat{
		EventTriggerDefinitionFormat: &e2smcccv1.EventTriggerDefinitionFormat_E2SmCccEventTriggerDefinitionFormat1{
			E2SmCccEventTriggerDefinitionFormat1: e2SmCccEventTriggerDefinitionFormat1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat2(e2SmCccEventTriggerDefinitionFormat2 *e2smcccv1.E2SmCCcEventTriggerDefinitionFormat2) (*e2smcccv1.EventTriggerDefinitionFormat, error) {

	item := &e2smcccv1.EventTriggerDefinitionFormat{
		EventTriggerDefinitionFormat: &e2smcccv1.EventTriggerDefinitionFormat_E2SmCccEventTriggerDefinitionFormat2{
			E2SmCccEventTriggerDefinitionFormat2: e2SmCccEventTriggerDefinitionFormat2,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat2() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat3(e2SmCccEventTriggerDefinitionFormat3 *e2smcccv1.E2SmCCcEventTriggerDefinitionFormat3) (*e2smcccv1.EventTriggerDefinitionFormat, error) {

	item := &e2smcccv1.EventTriggerDefinitionFormat{
		EventTriggerDefinitionFormat: &e2smcccv1.EventTriggerDefinitionFormat_E2SmCccEventTriggerDefinitionFormat3{
			E2SmCccEventTriggerDefinitionFormat3: e2SmCccEventTriggerDefinitionFormat3,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateEventTriggerDefinitionFormatE2SmCccEventTriggerDefinitionFormat3() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateActionDefinitionFormatE2SmCccActionDefinitionFormat1(e2SmCccActionDefinitionFormat1 *e2smcccv1.E2SmCCcActionDefinitionFormat1) (*e2smcccv1.ActionDefinitionFormat, error) {

	item := &e2smcccv1.ActionDefinitionFormat{
		ActionDefinitionFormat: &e2smcccv1.ActionDefinitionFormat_E2SmCccActionDefinitionFormat1{
			E2SmCccActionDefinitionFormat1: e2SmCccActionDefinitionFormat1,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateActionDefinitionFormatE2SmCccActionDefinitionFormat1() error validating PDU %s", err.Error())
	}

	return item, nil
}
func CreateActionDefinitionFormatE2SmCccActionDefinitionFormat2(e2SmCccActionDefinitionFormat2 *e2smcccv1.E2SmCCcActionDefinitionFormat2) (*e2smcccv1.ActionDefinitionFormat, error) {

	item := &e2smcccv1.ActionDefinitionFormat{
		ActionDefinitionFormat: &e2smcccv1.ActionDefinitionFormat_E2SmCccActionDefinitionFormat2{
			E2SmCccActionDefinitionFormat2: e2SmCccActionDefinitionFormat2,
		},
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateActionDefinitionFormatE2SmCccActionDefinitionFormat2() error validating PDU %s", err.Error())
	}

	return item, nil
}

func CreateMessageTypeInterfaceMessageIdInitiatingMessage() e2smcommoniesv1.MessageTypeInterfaceMessageId {
	return e2smcommoniesv1.MessageTypeInterfaceMessageId_MESSAGE_TYPE_INTERFACE_MESSAGE_ID_INITIATING_MESSAGE
}
func CreateMessageTypeInterfaceMessageIdSuccessfulOutcome() e2smcommoniesv1.MessageTypeInterfaceMessageId {
	return e2smcommoniesv1.MessageTypeInterfaceMessageId_MESSAGE_TYPE_INTERFACE_MESSAGE_ID_SUCCESSFUL_OUTCOME
}
func CreateMessageTypeInterfaceMessageIdUnsuccessfulOutcome() e2smcommoniesv1.MessageTypeInterfaceMessageId {
	return e2smcommoniesv1.MessageTypeInterfaceMessageId_MESSAGE_TYPE_INTERFACE_MESSAGE_ID_UNSUCCESSFUL_OUTCOME
}
func CreateInterfaceTypeNG() e2smcommoniesv1.InterfaceType {
	return e2smcommoniesv1.InterfaceType_INTERFACE_TYPE_N_G
}
func CreateInterfaceTypeXn() e2smcommoniesv1.InterfaceType {
	return e2smcommoniesv1.InterfaceType_INTERFACE_TYPE_XN
}
func CreateInterfaceTypeF1() e2smcommoniesv1.InterfaceType {
	return e2smcommoniesv1.InterfaceType_INTERFACE_TYPE_F1
}
func CreateInterfaceTypeE1() e2smcommoniesv1.InterfaceType {
	return e2smcommoniesv1.InterfaceType_INTERFACE_TYPE_E1
}
func CreateInterfaceTypeS1() e2smcommoniesv1.InterfaceType {
	return e2smcommoniesv1.InterfaceType_INTERFACE_TYPE_S1
}
func CreateInterfaceTypeX2() e2smcommoniesv1.InterfaceType {
	return e2smcommoniesv1.InterfaceType_INTERFACE_TYPE_X2
}
func CreateInterfaceTypeW1() e2smcommoniesv1.InterfaceType {
	return e2smcommoniesv1.InterfaceType_INTERFACE_TYPE_W1
}
func CreateRrcclassLteBCchBch() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_B_CCH_BCH
}
func CreateRrcclassLteBCchBchMbms() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_B_CCH_BCH_MBMS
}
func CreateRrcclassLteBCchDlSch() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_B_CCH_DL_SCH
}
func CreateRrcclassLteBCchDlSchBr() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_B_CCH_DL_SCH_BR
}
func CreateRrcclassLteBCchDlSchMbms() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_B_CCH_DL_SCH_MBMS
}
func CreateRrcclassLteMCch() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_M_CCH
}
func CreateRrcclassLtePCch() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_P_CCH
}
func CreateRrcclassLteDLCcch() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_D_L_CCCH
}
func CreateRrcclassLteDLDcch() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_D_L_DCCH
}
func CreateRrcclassLteULCcch() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_U_L_CCCH
}
func CreateRrcclassLteULDcch() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_U_L_DCCH
}
func CreateRrcclassLteSCMcch() e2smcommoniesv1.RrcclassLte {
	return e2smcommoniesv1.RrcclassLte_RRCCLASS_LTE_S_C_MCCH
}
func CreateRrcclassNrBCchBch() e2smcommoniesv1.RrcclassNR {
	return e2smcommoniesv1.RrcclassNR_RRCCLASS_NR_B_CCH_BCH
}
func CreateRrcclassNrBCchDlSch() e2smcommoniesv1.RrcclassNR {
	return e2smcommoniesv1.RrcclassNR_RRCCLASS_NR_B_CCH_DL_SCH
}
func CreateRrcclassNrDLCcch() e2smcommoniesv1.RrcclassNR {
	return e2smcommoniesv1.RrcclassNR_RRCCLASS_NR_D_L_CCCH
}
func CreateRrcclassNrDLDcch() e2smcommoniesv1.RrcclassNR {
	return e2smcommoniesv1.RrcclassNR_RRCCLASS_NR_D_L_DCCH
}
func CreateRrcclassNrPCch() e2smcommoniesv1.RrcclassNR {
	return e2smcommoniesv1.RrcclassNR_RRCCLASS_NR_P_CCH
}
func CreateRrcclassNrULCcch() e2smcommoniesv1.RrcclassNR {
	return e2smcommoniesv1.RrcclassNR_RRCCLASS_NR_U_L_CCCH
}
func CreateRrcclassNrULCcch1() e2smcommoniesv1.RrcclassNR {
	return e2smcommoniesv1.RrcclassNR_RRCCLASS_NR_U_L_CCCH1
}
func CreateRrcclassNrULDcch() e2smcommoniesv1.RrcclassNR {
	return e2smcommoniesv1.RrcclassNR_RRCCLASS_NR_U_L_DCCH
}
func CreateNrfrequencyShift7p5khzFalse() e2smcommoniesv1.NrfrequencyShift7P5Khz {
	return e2smcommoniesv1.NrfrequencyShift7P5Khz_NRFREQUENCY_SHIFT7P5KHZ_FALSE
}
func CreateNrfrequencyShift7p5khzTrue() e2smcommoniesv1.NrfrequencyShift7P5Khz {
	return e2smcommoniesv1.NrfrequencyShift7P5Khz_NRFREQUENCY_SHIFT7P5KHZ_TRUE
}
func CreateCellStateIdle() e2smcccv1.CellState {
	return e2smcccv1.CellState_CELL_STATE_IDLE
}
func CreateCellStateInactive() e2smcccv1.CellState {
	return e2smcccv1.CellState_CELL_STATE_INACTIVE
}
func CreateCellStateActive() e2smcccv1.CellState {
	return e2smcccv1.CellState_CELL_STATE_ACTIVE
}
func CreateCyclicPrefixNormal() e2smcccv1.CyclicPrefix {
	return e2smcccv1.CyclicPrefix_CYCLIC_PREFIX_NORMAL
}
func CreateCyclicPrefixExtended() e2smcccv1.CyclicPrefix {
	return e2smcccv1.CyclicPrefix_CYCLIC_PREFIX_EXTENDED
}
func CreateBwpContextDl() e2smcccv1.BwpContext {
	return e2smcccv1.BwpContext_BWP_CONTEXT_DL
}
func CreateBwpContextUl() e2smcccv1.BwpContext {
	return e2smcccv1.BwpContext_BWP_CONTEXT_UL
}
func CreateBwpContextSul() e2smcccv1.BwpContext {
	return e2smcccv1.BwpContext_BWP_CONTEXT_SUL
}
func CreateIsInitialBwpInitial() e2smcccv1.IsInitialBwp {
	return e2smcccv1.IsInitialBwp_IS_INITIAL_BWP_INITIAL
}
func CreateIsInitialBwpOther() e2smcccv1.IsInitialBwp {
	return e2smcccv1.IsInitialBwp_IS_INITIAL_BWP_OTHER
}
func CreateIndicationReasonUponSubscription() e2smcccv1.IndicationReason {
	return e2smcccv1.IndicationReason_INDICATION_REASON_UPON_SUBSCRIPTION
}
func CreateIndicationReasonUponChange() e2smcccv1.IndicationReason {
	return e2smcccv1.IndicationReason_INDICATION_REASON_UPON_CHANGE
}
func CreateIndicationReasonPeriodic() e2smcccv1.IndicationReason {
	return e2smcccv1.IndicationReason_INDICATION_REASON_PERIODIC
}
func CreateChangeTypeNone() e2smcccv1.ChangeType {
	return e2smcccv1.ChangeType_CHANGE_TYPE_NONE
}
func CreateChangeTypeModification() e2smcccv1.ChangeType {
	return e2smcccv1.ChangeType_CHANGE_TYPE_MODIFICATION
}
func CreateChangeTypeAddition() e2smcccv1.ChangeType {
	return e2smcccv1.ChangeType_CHANGE_TYPE_ADDITION
}
func CreateChangeTypeDeletion() e2smcccv1.ChangeType {
	return e2smcccv1.ChangeType_CHANGE_TYPE_DELETION
}
func CreateReportTypeAll() e2smcccv1.ReportType {
	return e2smcccv1.ReportType_REPORT_TYPE_ALL
}
func CreateReportTypeChange() e2smcccv1.ReportType {
	return e2smcccv1.ReportType_REPORT_TYPE_CHANGE
}
func CreateResourceTypePrbUl() e2smcccv1.ResourceType {
	return e2smcccv1.ResourceType_RESOURCE_TYPE_PRB_UL
}
func CreateResourceTypePrbDl() e2smcccv1.ResourceType {
	return e2smcccv1.ResourceType_RESOURCE_TYPE_PRB_DL
}
func CreateResourceTypeDrb() e2smcccv1.ResourceType {
	return e2smcccv1.ResourceType_RESOURCE_TYPE_DRB
}
func CreateResourceTypeRrc() e2smcccv1.ResourceType {
	return e2smcccv1.ResourceType_RESOURCE_TYPE_RRC
}
func CreateSchedulerTypeRoundRobin() e2smcccv1.SchedulerType {
	return e2smcccv1.SchedulerType_SCHEDULER_TYPE_ROUND_ROBIN
}
func CreateSchedulerTypeProportionallyFair() e2smcccv1.SchedulerType {
	return e2smcccv1.SchedulerType_SCHEDULER_TYPE_PROPORTIONALLY_FAIR
}
func CreateSchedulerTypeQosBased() e2smcccv1.SchedulerType {
	return e2smcccv1.SchedulerType_SCHEDULER_TYPE_QOS_BASED
}
func CreateSsbPeriodicityN5() e2smcccv1.SsbPeriodicity {
	return e2smcccv1.SsbPeriodicity_SSB_PERIODICITY_N5
}
func CreateSsbPeriodicityN10() e2smcccv1.SsbPeriodicity {
	return e2smcccv1.SsbPeriodicity_SSB_PERIODICITY_N10
}
func CreateSsbPeriodicityN20() e2smcccv1.SsbPeriodicity {
	return e2smcccv1.SsbPeriodicity_SSB_PERIODICITY_N20
}
func CreateSsbPeriodicityN40() e2smcccv1.SsbPeriodicity {
	return e2smcccv1.SsbPeriodicity_SSB_PERIODICITY_N40
}
func CreateSsbPeriodicityN80() e2smcccv1.SsbPeriodicity {
	return e2smcccv1.SsbPeriodicity_SSB_PERIODICITY_N80
}
func CreateSsbPeriodicityN160() e2smcccv1.SsbPeriodicity {
	return e2smcccv1.SsbPeriodicity_SSB_PERIODICITY_N160
}
func CreateSsbDurationN1() e2smcccv1.SsbDuration {
	return e2smcccv1.SsbDuration_SSB_DURATION_N1
}
func CreateSsbDurationN2() e2smcccv1.SsbDuration {
	return e2smcccv1.SsbDuration_SSB_DURATION_N2
}
func CreateSsbDurationN3() e2smcccv1.SsbDuration {
	return e2smcccv1.SsbDuration_SSB_DURATION_N3
}
func CreateSsbDurationN4() e2smcccv1.SsbDuration {
	return e2smcccv1.SsbDuration_SSB_DURATION_N4
}
func CreateSsbDurationN5() e2smcccv1.SsbDuration {
	return e2smcccv1.SsbDuration_SSB_DURATION_N5
}
func CreateSsbSubCarrierSpacingN15() e2smcccv1.SsbSubCarrierSpacing {
	return e2smcccv1.SsbSubCarrierSpacing_SSB_SUB_CARRIER_SPACING_N15
}
func CreateSsbSubCarrierSpacingN30() e2smcccv1.SsbSubCarrierSpacing {
	return e2smcccv1.SsbSubCarrierSpacing_SSB_SUB_CARRIER_SPACING_N30
}
func CreateSsbSubCarrierSpacingN120() e2smcccv1.SsbSubCarrierSpacing {
	return e2smcccv1.SsbSubCarrierSpacing_SSB_SUB_CARRIER_SPACING_N120
}
func CreateSsbSubCarrierSpacingN240() e2smcccv1.SsbSubCarrierSpacing {
	return e2smcccv1.SsbSubCarrierSpacing_SSB_SUB_CARRIER_SPACING_N240
}
func CreateSubCarrierSpacingN15() e2smcccv1.SubCarrierSpacing {
	return e2smcccv1.SubCarrierSpacing_SUB_CARRIER_SPACING_N15
}
func CreateSubCarrierSpacingN30() e2smcccv1.SubCarrierSpacing {
	return e2smcccv1.SubCarrierSpacing_SUB_CARRIER_SPACING_N30
}
func CreateSubCarrierSpacingN60() e2smcccv1.SubCarrierSpacing {
	return e2smcccv1.SubCarrierSpacing_SUB_CARRIER_SPACING_N60
}
func CreateSubCarrierSpacingN120() e2smcccv1.SubCarrierSpacing {
	return e2smcccv1.SubCarrierSpacing_SUB_CARRIER_SPACING_N120
}
