//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcEventTriggerFormat1(t *testing.T) {

	associatedUeInfo := make([]*e2smrcv1.EventTriggerUeInfoItem, 0)

	ranParameterTesting := make([]*e2smrcv1.RanparameterTestingItem, 0)

	ranParameterValue, err := CreateRanparameterValueInt(42)
	assert.NilError(t, err)

	ranParameterType, err := CreateRanParameterTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	ranParameterTestingItem, err := CreateRanparameterTestingItem(1, ranParameterType)
	assert.NilError(t, err)

	ranParameterTesting = append(ranParameterTesting, ranParameterTestingItem)

	ueType, err := CreateUeTypeGroup(&e2smrcv1.RanparameterTesting{
		Value: ranParameterTesting,
	})
	assert.NilError(t, err)

	ueInfo, err := CreateEventTriggerUeInfoItem(2, ueType)
	assert.NilError(t, err)
	ueInfo.SetLogicalOr(CreateLogicalOrTrue())

	associatedUeInfo = append(associatedUeInfo, ueInfo)

	messageList := make([]*e2smrcv1.E2SmRcEventTriggerFormat1Item, 0)

	rrcType, err := CreateRrcTypeLte(e2smcommonies.RrcclassLte_RRCCLASS_LTE_B_CCH_BCH)
	assert.NilError(t, err)

	msgTypeChoice, err := CreateMessageTypeChoiceRrc(rrcType, 32)
	assert.NilError(t, err)

	messageItem, err := CreateE2SmRcEventTriggerFormat1Item(1, msgTypeChoice, nil, nil, nil, nil)
	assert.NilError(t, err)
	messageItem.SetAssociatedUeinfo(&e2smrcv1.EventTriggerUeInfo{
		UeInfoList: associatedUeInfo,
	}).SetMessageDirection(CreateMessageDirectionOutgoing()).SetLogicalOr(CreateLogicalOrFalse())

	messageList = append(messageList, messageItem)

	msg, err := CreateE2SmRcEventTriggerFormat1(messageList)
	assert.NilError(t, err)
	msg.GetRicEventTriggerFormats().GetEventTriggerFormat1().SetGlobalAssociatedUeinfo(&e2smrcv1.EventTriggerUeInfo{
		UeInfoList: associatedUeInfo,
	})

	aper, err := encoder.PerEncodeE2SmRcEventTrigger(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcEventTrigger(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcEventTriggerFormat2(t *testing.T) {

	associatedUeInfo := make([]*e2smrcv1.EventTriggerUeInfoItem, 0)

	ranParameterTesting := make([]*e2smrcv1.RanparameterTestingItem, 0)

	ranParameterValue, err := CreateRanparameterValueInt(42)
	assert.NilError(t, err)

	ranParameterType, err := CreateRanParameterTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	ranParameterTestingItem, err := CreateRanparameterTestingItem(1, ranParameterType)
	assert.NilError(t, err)

	ranParameterTesting = append(ranParameterTesting, ranParameterTestingItem)

	ueType, err := CreateUeTypeGroup(&e2smrcv1.RanparameterTesting{
		Value: ranParameterTesting,
	})
	assert.NilError(t, err)

	ueInfo, err := CreateEventTriggerUeInfoItem(2, ueType)
	assert.NilError(t, err)
	ueInfo.SetLogicalOr(CreateLogicalOrTrue())

	associatedUeInfo = append(associatedUeInfo, ueInfo)

	msg, err := CreateE2SmRcEventTriggerFormat2(12, 15)
	assert.NilError(t, err)
	msg.GetRicEventTriggerFormats().GetEventTriggerFormat2().SetAssociatedUeinfo(&e2smrcv1.EventTriggerUeInfo{
		UeInfoList: associatedUeInfo,
	}).SetAssociatedE2NodeInfo(&e2smrcv1.RanparameterTesting{
		Value: ranParameterTesting,
	})

	aper, err := encoder.PerEncodeE2SmRcEventTrigger(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcEventTrigger(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcEventTriggerFormat3(t *testing.T) {

	associatedCellInfo := make([]*e2smrcv1.EventTriggerCellInfoItem, 0)

	eutraCgi, err := CreateEutraCgi([]byte{0xFF, 0x00, 0xFF}, &asn1.BitString{
		Value: []byte{0xFF, 0xFF, 0xFF, 0xF0},
		Len:   28,
	})
	assert.NilError(t, err)

	cgi, err := CreateCgiEutraCgi(eutraCgi)
	assert.NilError(t, err)

	cellType, err := CreateCellTypeChoiceIndividual(cgi)
	assert.NilError(t, err)

	cellInfoItem, err := CreateEventTriggerCellInfoItem(1, cellType)
	assert.NilError(t, err)
	cellInfoItem.SetLogicalOr(CreateLogicalOrTrue())
	associatedCellInfo = append(associatedCellInfo, cellInfoItem)

	item, err := CreateE2SmRcEventTriggerFormat3Item(91, 513)
	assert.NilError(t, err)
	item.SetAssociatedCellInfo(&e2smrcv1.EventTriggerCellInfo{
		CellInfoList: associatedCellInfo,
	}).SetLogicalOr(CreateLogicalOrFalse())

	e2NodeInfoChangeList := make([]*e2smrcv1.E2SmRcEventTriggerFormat3Item, 0)
	e2NodeInfoChangeList = append(e2NodeInfoChangeList, item)

	msg, err := CreateE2SmRcEventTriggerFormat3(e2NodeInfoChangeList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcEventTrigger(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcEventTrigger(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcEventTriggerFormat4(t *testing.T) {

	associatedUeInfo := make([]*e2smrcv1.EventTriggerUeInfoItem, 0)

	ranParameterTesting := make([]*e2smrcv1.RanparameterTestingItem, 0)

	ranParameterValue, err := CreateRanparameterValueOctS([]byte{0x33, 0xFF})
	assert.NilError(t, err)

	ranParameterType, err := CreateRanParameterTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	ranParameterTestingItem, err := CreateRanparameterTestingItem(1, ranParameterType)
	assert.NilError(t, err)

	ranParameterTesting = append(ranParameterTesting, ranParameterTestingItem)

	ueType, err := CreateUeTypeGroup(&e2smrcv1.RanparameterTesting{
		Value: ranParameterTesting,
	})
	assert.NilError(t, err)

	ueInfo, err := CreateEventTriggerUeInfoItem(2, ueType)
	assert.NilError(t, err)
	ueInfo.SetLogicalOr(CreateLogicalOrTrue())

	associatedUeInfo = append(associatedUeInfo, ueInfo)

	triggerTypeChoice, err := CreateTriggerTypeChoiceUeID(27)
	assert.NilError(t, err)

	item, err := CreateE2SmRcEventTriggerFormat4Item(91, triggerTypeChoice)
	assert.NilError(t, err)
	item.SetAssociatedUeinfo(&e2smrcv1.EventTriggerUeInfo{
		UeInfoList: associatedUeInfo,
	}).SetLogicalOr(CreateLogicalOrFalse())

	ueInfoChangeList := make([]*e2smrcv1.E2SmRcEventTriggerFormat4Item, 0)
	ueInfoChangeList = append(ueInfoChangeList, item)

	msg, err := CreateE2SmRcEventTriggerFormat4(ueInfoChangeList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcEventTrigger(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcEventTrigger(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcEventTriggerFormat5(t *testing.T) {

	associatedUeInfo := make([]*e2smrcv1.EventTriggerUeInfoItem, 0)

	ranParameterTesting := make([]*e2smrcv1.RanparameterTestingItem, 0)

	ranParameterValue, err := CreateRanparameterValueOctS([]byte{0x33, 0xFF})
	assert.NilError(t, err)

	ranParameterType, err := CreateRanParameterTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	ranParameterTestingItem, err := CreateRanparameterTestingItem(1, ranParameterType)
	assert.NilError(t, err)

	ranParameterTesting = append(ranParameterTesting, ranParameterTestingItem)

	ueType, err := CreateUeTypeGroup(&e2smrcv1.RanparameterTesting{
		Value: ranParameterTesting,
	})
	assert.NilError(t, err)

	ueInfo, err := CreateEventTriggerUeInfoItem(2, ueType)
	assert.NilError(t, err)
	ueInfo.SetLogicalOr(CreateLogicalOrTrue())

	associatedUeInfo = append(associatedUeInfo, ueInfo)

	associatedCellInfo := make([]*e2smrcv1.EventTriggerCellInfoItem, 0)

	eutraCgi, err := CreateEutraCgi([]byte{0xFF, 0x00, 0xFF}, &asn1.BitString{
		Value: []byte{0xFF, 0xFF, 0xFF, 0xF0},
		Len:   28,
	})
	assert.NilError(t, err)

	cgi, err := CreateCgiEutraCgi(eutraCgi)
	assert.NilError(t, err)

	cellType, err := CreateCellTypeChoiceIndividual(cgi)
	assert.NilError(t, err)

	cellInfoItem, err := CreateEventTriggerCellInfoItem(1, cellType)
	assert.NilError(t, err)
	cellInfoItem.SetLogicalOr(CreateLogicalOrTrue())
	associatedCellInfo = append(associatedCellInfo, cellInfoItem)

	msg, err := CreateE2SmRcEventTriggerFormat5(CreateOnDemandTrue())
	assert.NilError(t, err)
	msg.GetRicEventTriggerFormats().GetEventTriggerFormat5().SetAssociatedCellInfo(&e2smrcv1.EventTriggerCellInfo{
		CellInfoList: associatedCellInfo,
	}).SetAssociatedUeinfo(&e2smrcv1.EventTriggerUeInfo{
		UeInfoList: associatedUeInfo,
	})

	aper, err := encoder.PerEncodeE2SmRcEventTrigger(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcEventTrigger(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}
