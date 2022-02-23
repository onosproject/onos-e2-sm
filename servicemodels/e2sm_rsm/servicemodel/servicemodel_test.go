// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/pdubuilder"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
	"testing"
)

var rsmv1TestSm RsmServiceModel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {
	plmnID := []byte{0x00, 0x01, 0x0F}
	nrCellID := &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x00, 0x10},
		Len:   36,
	}

	cgi, err := pdubuilder.CreateNrCGI(plmnID, nrCellID)
	assert.NilError(t, err)

	ih, err := pdubuilder.CreateE2SmRsmIndicationHeaderFormat1(cgi)
	assert.NilError(t, err)
	t.Logf("Created E2SM-RSM-IndicationHeader is \n%v", ih)

	protoBytes, err := proto.Marshal(ih)
	assert.NilError(t, err, "unexpected error marshalling E2SmRsmIndicationHeader to bytes")

	asn1Bytes, err := rsmv1TestSm.IndicationHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-RSM-IndicationHeader asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	indicationHeaderAsn1Bytes := []byte{0x00, 0x00, 0x01, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x10}

	protoBytes, err := rsmv1TestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	testIH := &e2sm_rsm_ies.E2SmRsmIndicationHeader{}
	err = proto.Unmarshal(protoBytes, testIH)
	assert.NilError(t, err)
	t.Logf("Decoded message is \n%v", testIH)
	assert.DeepEqual(t, []byte{0x00, 0x01, 0x0F}, testIH.GetIndicationHeaderFormat1().GetCgi().GetNRCgi().GetPLmnidentity().GetValue())
	assert.DeepEqual(t, []byte{0x00, 0x00, 0x00, 0x00, 0x10}, testIH.GetIndicationHeaderFormat1().GetCgi().GetNRCgi().GetNRcellIdentity().GetValue().GetValue())
}

func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
	ueID, err := pdubuilder.CreateUeIDAmfUeNgapID(1)
	assert.NilError(t, err)

	ulSm := make([]*e2sm_rsm_ies.SliceMetrics, 0)
	ulM1, err := pdubuilder.CreateSliceMetrics(100, 100, 100, 15)
	assert.NilError(t, err)
	ulSm = append(ulSm, ulM1)
	ulM2, err := pdubuilder.CreateSliceMetrics(10, 1, 50, 7)
	assert.NilError(t, err)
	ulSm = append(ulSm, ulM2)
	ulM3, err := pdubuilder.CreateSliceMetrics(47, 16, 12, 1)
	assert.NilError(t, err)
	ulSm = append(ulSm, ulM3)

	dlSm := make([]*e2sm_rsm_ies.SliceMetrics, 0)
	dlM1, err := pdubuilder.CreateSliceMetrics(91, 31, 54, 11)
	assert.NilError(t, err)
	dlSm = append(dlSm, dlM1)
	dlM2, err := pdubuilder.CreateSliceMetrics(84, 37, 41, 6)
	assert.NilError(t, err)
	dlSm = append(dlSm, dlM2)

	im, err := pdubuilder.CreateE2SmRsmIndicationMessageFormat1(ueID, 27, 14, pdubuilder.CreateEmmCaseAttach(),
		ulSm, dlSm)
	assert.NilError(t, err)
	assert.Assert(t, im != nil)
	t.Logf("Created E2SM-RSM-IndicationMessage is \n%v", im)

	//err = newE2SmKpmPdu.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(im)
	assert.NilError(t, err, "unexpected error marshalling E2SmRsmIndicationMessage to bytes")

	asn1Bytes, err := rsmv1TestSm.IndicationMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-RSM-IndicationMessage asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {
	indicationMessageAsn1 := []byte{0x06, 0x00, 0x01, 0x00, 0x1b, 0x00, 0x0e, 0x00, 0x03, 0x64, 0x00, 0x64, 0xc9, 0xe1, 0x40, 0x01,
		0x64, 0xe5, 0xe0, 0x10, 0x18, 0x20, 0x02, 0x5b, 0x00, 0x1f, 0x6d, 0x6a, 0x80, 0x25, 0x52, 0xc0}

	protoBytes, err := rsmv1TestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	testIM := &e2sm_rsm_ies.E2SmRsmIndicationMessage{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	t.Logf("Decoded message is \n%v", testIM)
	assert.Equal(t, int64(1), testIM.GetIndicationMessageFormat1().GetUeId().GetAmfUeNgapId().GetValue())
	assert.Equal(t, int64(27), testIM.GetIndicationMessageFormat1().GetCuUeF1ApId().GetValue())
	assert.Equal(t, int64(14), testIM.GetIndicationMessageFormat1().GetDuUeF1ApId().GetValue())
	assert.Equal(t, pdubuilder.CreateEmmCaseAttach().Number(), testIM.GetIndicationMessageFormat1().GetEmmCase().Number())
	assert.Assert(t, len(testIM.GetIndicationMessageFormat1().GetUlSlicingMetrics()) == 3)
	assert.Assert(t, len(testIM.GetIndicationMessageFormat1().GetDlSlicingMetrics()) == 2)
}

func TestServicemodel_RanFuncDescriptionProtoToASN1(t *testing.T) {
	supportedConfigList := make([]*e2sm_rsm_ies.SupportedSlicingConfigItem, 0)
	su, err := pdubuilder.CreateSupportedSlicingConfigItem(pdubuilder.CreateE2SmRsmCommandSliceUpdate())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, su)
	sc, err := pdubuilder.CreateSupportedSlicingConfigItem(pdubuilder.CreateE2SmRsmCommandSliceCreate())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, sc)
	sd, err := pdubuilder.CreateSupportedSlicingConfigItem(pdubuilder.CreateE2SmRsmCommandSliceDelete())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, sd)
	ueAssoc, err := pdubuilder.CreateSupportedSlicingConfigItem(pdubuilder.CreateE2SmRsmCommandUeAssociate())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, ueAssoc)
	etd, err := pdubuilder.CreateSupportedSlicingConfigItem(pdubuilder.CreateE2SmRsmCommandEventTriggers())
	assert.NilError(t, err)
	supportedConfigList = append(supportedConfigList, etd)

	slicingCapability, err := pdubuilder.CreateSlicingCapabilityItem(71, 27, pdubuilder.CreateSlicingTypeDynamic(), 10,
		supportedConfigList)
	assert.NilError(t, err)

	slicingCapList := make([]*e2sm_rsm_ies.NodeSlicingCapabilityItem, 0)
	slicingCapList = append(slicingCapList, slicingCapability)

	rfd, err := pdubuilder.CreateE2SmRsmRanFunctionDescription("E2SM-RSM",
		"1.3.6.1.4.1.53148.1.1.2.102", "RAN Slicing Service Model", slicingCapList)
	assert.NilError(t, err)
	assert.Assert(t, rfd != nil)
	t.Logf("Created E2SM-RSM-RanFunctionDescription is \n%v", rfd)

	//err = newE2SmKpmPdu.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(rfd)
	assert.NilError(t, err, "unexpected error marshalling E2SmRsmRanfunctionDescription to bytes")

	asn1Bytes, err := rsmv1TestSm.RanFuncDescriptionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-RSM-RANfunctionDescription asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_RanFuncDescriptionASN1toProto(t *testing.T) {
	ranFuncDescriptionAsn1 := []byte{0x00, 0x70, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x52, 0x53, 0x4d, 0x00, 0x00, 0x1a, 0x31, 0x2e, 0x33,
		0x2e, 0x36, 0x2e, 0x31, 0x2e, 0x34, 0x2e, 0x31, 0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e, 0x31,
		0x2e, 0x31, 0x2e, 0x32, 0x2e, 0x31, 0x30, 0x32, 0x0c, 0x00, 0x52, 0x41, 0x4e, 0x20, 0x53, 0x6c,
		0x69, 0x63, 0x69, 0x6e, 0x67, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x4d, 0x6f,
		0x64, 0x65, 0x6c, 0x01, 0x00, 0x47, 0x00, 0x1b, 0x40, 0x01, 0x0a, 0x40, 0x80, 0x43, 0x20}

	protoBytes, err := rsmv1TestSm.RanFuncDescriptionASN1toProto(ranFuncDescriptionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	testRFD := &e2sm_rsm_ies.E2SmRsmRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, testRFD)
	t.Logf("Decoded message is \n%v", testRFD)
	assert.NilError(t, err)
	assert.Equal(t, "1.3.6.1.4.1.53148.1.1.2.102", testRFD.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, "RAN Slicing Service Model", testRFD.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "E2SM-RSM", testRFD.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, 1, len(testRFD.GetRicSlicingNodeCapabilityList()))
	assert.Equal(t, int32(10), testRFD.GetRicSlicingNodeCapabilityList()[0].GetMaxNumberOfUesPerSlice())
	assert.Equal(t, 5, len(testRFD.GetRicSlicingNodeCapabilityList()[0].GetSupportedConfig()))
}

func TestServicemodel_EventTriggerDefinitionProtoToASN1(t *testing.T) {

	etd, err := pdubuilder.CreateE2SmRsmEventTriggerDefinitionFormat1(pdubuilder.CreateRsmRicindicationTriggerTypeUponEmmEvent())
	assert.NilError(t, err)
	t.Logf("Created E2SM-RSM-EventTriggerDefinition is \n%v", etd)

	//err = e2SmKpmEventTriggerDefinition.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(etd)
	assert.NilError(t, err, "unexpected error marshalling E2SmRsmEventTriggerDefinition to bytes")

	asn1Bytes, err := rsmv1TestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-RSM-EventTriggerDefinition asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
	eventTriggerDefinitionAsn1 := []byte{0x00, 0x01, 0x00, 0x46, 0x00, 0x15, 0x00, 0x2b, 0x40, 0x01, 0x30, 0x82, 0xc0, 0x7f, 0x2f, 0xbe,
		0x04, 0x05, 0x3e, 0x6c, 0x80, 0x0b}

	protoBytes, err := rsmv1TestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	testETD := &e2sm_rsm_ies.E2SmRsmEventTriggerDefinition{}
	err = proto.Unmarshal(protoBytes, testETD)
	t.Logf("Decoded message is \n%v", testETD)
	assert.NilError(t, err)
	assert.Equal(t, pdubuilder.CreateRsmEmmTriggerTypeUeAttach().Number(), testETD.GetEventDefinitionFormats().GetEventDefinitionFormat1().GetTriggerType().Number())
}

func TestServicemodel_ControlHeaderProtoToASN1(t *testing.T) {
	ch, err := pdubuilder.CreateE2SmRsmControlHeader(pdubuilder.CreateE2SmRsmCommandSliceUpdate())
	assert.NilError(t, err)
	t.Logf("Created E2SM-RSM-ControlHeader is \n%v", ch)

	//err = newE2SmKpmPdu.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	protoBytes, err := proto.Marshal(ch)
	assert.NilError(t, err, "unexpected error marshalling E2SmRsmControlHeader to bytes")

	asn1Bytes, err := rsmv1TestSm.ControlHeaderProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-RSM-ControlHeader asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlHeaderASN1toProto(t *testing.T) {
	chAsn1 := []byte{0x08}

	protoBytes, err := rsmv1TestSm.ControlHeaderASN1toProto(chAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	testCH := &e2sm_rsm_ies.E2SmRsmControlHeader{}
	err = proto.Unmarshal(protoBytes, testCH)
	assert.NilError(t, err)
	t.Logf("Decoded message is \n%v", testCH)
	assert.Equal(t, pdubuilder.CreateE2SmRsmCommandSliceUpdate().Number(), testCH.GetRsmCommand().Number())
}

func TestServicemodel_ControlMessageProtoToASN1(t *testing.T) {
	ueID, err := pdubuilder.CreateUeIDRanUeNgapID(7)
	assert.NilError(t, err)

	drbIDfourG, err := pdubuilder.CreateDrbIDfourG(12, 127)
	assert.NilError(t, err)
	bearerID1, err := pdubuilder.CreateBearerIDdrb(drbIDfourG)
	assert.NilError(t, err)

	flowMap := make([]*e2sm_rsm_ies.QoSflowLevelParameters, 0)
	dqos, err := pdubuilder.CreateQosFlowLevelParametersDynamic(10, 62, 54)
	assert.NilError(t, err)
	flowMap = append(flowMap, dqos)
	ndqos, err := pdubuilder.CreateQosFlowLevelParametersNonDynamic(11)
	assert.NilError(t, err)
	flowMap = append(flowMap, ndqos)

	drbIDfiveG, err := pdubuilder.CreateDrbIDfiveG(27, 62, flowMap)
	assert.NilError(t, err)
	bearerID2, err := pdubuilder.CreateBearerIDdrb(drbIDfiveG)
	assert.NilError(t, err)

	bearerList := make([]*e2sm_rsm_ies.BearerId, 0)
	bearerList = append(bearerList, bearerID1)
	bearerList = append(bearerList, bearerID2)

	config, err := pdubuilder.CreateSliceAssociate(ueID, bearerList, 7)
	assert.NilError(t, err)
	config.SetUplinkSliceID(19)

	cm, err := pdubuilder.CreateE2SmRsmControlMessageSliceAssociate(config)
	assert.NilError(t, err)
	t.Logf("Created E2SM-RSM-ControlMessage (Associate Slice) is \n%v", cm)

	//err = newE2SmKpmPdu.Validate()
	//assert.NilError(t, err, "error validating E2SmPDU")

	protoBytes, err := proto.Marshal(cm)
	assert.NilError(t, err, "unexpected error marshalling E2SmRsmControlMessage to bytes")

	asn1Bytes, err := rsmv1TestSm.ControlMessageProtoToASN1(protoBytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-RSM-ControlMessage asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_ControlMessageASN1toProto(t *testing.T) {
	cmAsn1 := []byte{0x69, 0x00, 0x07, 0x08, 0x2c, 0x7f, 0x2d, 0x3e, 0x04, 0x05, 0x3e, 0x6c, 0x80, 0x0b, 0x00, 0x07,
		0x00, 0x13}

	protoBytes, err := rsmv1TestSm.ControlMessageASN1toProto(cmAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	testCM := &e2sm_rsm_ies.E2SmRsmControlMessage{}
	err = proto.Unmarshal(protoBytes, testCM)
	assert.NilError(t, err)
	t.Logf("Decoded message is \n%v", testCM)
	assert.Equal(t, int64(7), testCM.GetSliceAssociate().GetUeId().GetRanUeNgapId().GetValue())
	assert.Equal(t, int64(7), testCM.GetSliceAssociate().GetDownLinkSliceId().GetValue())
	assert.Equal(t, 2, len(testCM.GetSliceAssociate().GetBearerId()))
	assert.Equal(t, int32(127), testCM.GetSliceAssociate().GetBearerId()[0].GetDrbId().GetFourGdrbId().GetQci().GetValue())
	assert.Equal(t, int32(62), testCM.GetSliceAssociate().GetBearerId()[1].GetDrbId().GetFiveGdrbId().GetQfi().GetValue())
}
