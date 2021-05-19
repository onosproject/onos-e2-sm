// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"encoding/hex"
	rcprectypes "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var rfdBytes = "00f04f52414e2d4532534d2d52432d50524500001a312e332e362e312e342e312e35333134382e312e322e322e3130300280" +
	"52432d505245600001068052432d5052452d74726967676572000100010c8050434920616e64204e52542075706461746520666f7220674e420" +
	"0010001"
var etdBytes = "00"
var ihBytes = "2013f1840000000010"
var imBytes = "580aef7a00000100010013f184000000001affff000001"
var chBytes = "2013f1840000000010"
var cmBytes = "0000010100706369000006"
var coBytes = "200000000001"

func Test_DecodeRanFunctionDescriptionBytesRadisys(t *testing.T) {
	rfdPer, err := hexlib.Asn1BytesToByte(rfdBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for RanFunctionDescription in HEX format are \n%v\n", hex.Dump(rfdPer))
	rfd, err := rcprectypes.PerDecodeE2SmRcPreRanfunctionDescription(rfdPer)
	assert.NilError(t, err)
	err = rfd.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", rfd)
	assert.Equal(t, "ORAN-E2SM-RC-PRE", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "RC-PRE", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "1.3.6.1.4.1.53148.1.2.2.100", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, int32(1), rfd.GetE2SmRcPreRanfunctionItem().GetRicEventTriggerStyleList()[0].GetRicEventTriggerFormatType().GetValue())
	assert.Equal(t, "RC-PRE-trigger", rfd.GetE2SmRcPreRanfunctionItem().GetRicEventTriggerStyleList()[0].GetRicEventTriggerStyleName().GetValue())
	assert.Equal(t, int32(1), rfd.GetE2SmRcPreRanfunctionItem().GetRicReportStyleList()[0].GetRicReportStyleType().GetValue())
	assert.Equal(t, "PCI and NRT update for gNB", rfd.GetE2SmRcPreRanfunctionItem().GetRicReportStyleList()[0].GetRicReportStyleName().GetValue())

	// Printing message in XER form
	xer, err := rcprectypes.XerEncodeE2SmRcPreRanfunctionDescription(rfd)
	assert.NilError(t, err)
	t.Logf("RanFunctionDescription XER is \n%s", xer)
}

func Test_DecodeEventTriggerDefinitionBytesRadisys(t *testing.T) {
	etdPer, err := hexlib.Asn1BytesToByte(etdBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for EventTriggerDefinition in HEX format are \n%v\n", hex.Dump(etdPer))
	etd, err := rcprectypes.PerDecodeE2SmRcPreEventTriggerDefinition(etdPer)
	assert.NilError(t, err)
	err = etd.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", etd)
	assert.Equal(t, e2sm_rc_pre_v2.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_UPON_CHANGE, etd.GetEventDefinitionFormat1().GetTriggerType())

	// Printing message in XER form
	xer, err := rcprectypes.XerEncodeE2SmRcPreEventTriggerDefinition(etd)
	assert.NilError(t, err)
	t.Logf("EventTriggerDefinition XER is \n%s", xer)
}

func Test_DecodeIndicationHeaderBytesRadisys(t *testing.T) {
	ihPer, err := hexlib.Asn1BytesToByte(ihBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for IndicationHeader in HEX format are \n%v\n", hex.Dump(ihPer))
	ih, err := rcprectypes.PerDecodeE2SmRcPreIndicationHeader(ihPer)
	assert.NilError(t, err)
	err = ih.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", ih)
	assert.DeepEqual(t, []byte{0x13, 0xF1, 0x84}, ih.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetPLmnIdentity().GetValue())
	//assert.Equal(t, uint64(1), ih.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, uint32(36), ih.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())

	// Printing message in XER form
	xer, err := rcprectypes.XerEncodeE2SmRcPreIndicationHeader(ih)
	assert.NilError(t, err)
	t.Logf("IndicationHeader XER is \n%s", xer)
}

func Test_DecodeIndicationMessageBytesRadisys(t *testing.T) {
	imPer, err := hexlib.Asn1BytesToByte(imBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for IndicationMessage in HEX format are \n%v\n", hex.Dump(imPer))
	im, err := rcprectypes.PerDecodeE2SmRcPreIndicationMessage(imPer)
	assert.NilError(t, err)
	err = im.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", im)
	assert.Equal(t, int32(716666), im.GetIndicationMessageFormat1().GetDlArfcn().GetNrArfcn().GetValue())
	assert.Equal(t, e2sm_rc_pre_v2.CellSize_CELL_SIZE_FEMTO, im.GetIndicationMessageFormat1().GetCellSize())
	assert.Equal(t, int32(1), im.GetIndicationMessageFormat1().GetPci().GetValue())
	assert.DeepEqual(t, []byte{0x13, 0xf1, 0x84}, im.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetNrCgi().GetPLmnIdentity().GetValue())
	//assert.Equal(t, uint64(1), im.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, uint32(36), im.GetIndicationMessageFormat1().GetNeighbors()[0].GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())
	assert.Equal(t, int32(65535), im.GetIndicationMessageFormat1().GetNeighbors()[0].GetDlArfcn().GetNrArfcn().GetValue())
	assert.Equal(t, e2sm_rc_pre_v2.CellSize_CELL_SIZE_FEMTO, im.GetIndicationMessageFormat1().GetNeighbors()[0].GetCellSize())
	assert.Equal(t, int32(1), im.GetIndicationMessageFormat1().GetNeighbors()[0].GetPci().GetValue())

	// Printing message in XER form
	xer, err := rcprectypes.XerEncodeE2SmRcPreIndicationMessage(im)
	assert.NilError(t, err)
	t.Logf("IndicationMessage XER is \n%s", xer)
}

func Test_DecodeControlHeaderBytesRadisys(t *testing.T) {
	chPer, err := hexlib.Asn1BytesToByte(chBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for ControlHeader in HEX format are \n%v\n", hex.Dump(chPer))
	ch, err := rcprectypes.PerDecodeE2SmRcPreControlHeader(chPer)
	assert.NilError(t, err)
	err = ch.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", ch)
	assert.DeepEqual(t, []byte{0x13, 0xF1, 0x84}, ch.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetPLmnIdentity().GetValue())
	//assert.Equal(t, uint64(1), ch.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetValue())
	assert.Equal(t, uint32(36), ch.GetControlHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().GetValue().GetLen())
	assert.Equal(t, e2sm_rc_pre_v2.RcPreCommand_RC_PRE_COMMAND_SET_PARAMETERS, ch.GetControlHeaderFormat1().GetRcCommand())

	// Printing message in XER form
	xer, err := rcprectypes.XerEncodeE2SmRcPreControlHeader(ch)
	assert.NilError(t, err)
	t.Logf("ControlHeader XER is \n%s", xer)
}

func Test_DecodeControlMessageBytesRadisys(t *testing.T) {
	cmPer, err := hexlib.Asn1BytesToByte(cmBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for ControlMessage in HEX format are \n%v\n", hex.Dump(cmPer))
	cm, err := rcprectypes.PerDecodeE2SmRcPreControlMessage(cmPer)
	assert.NilError(t, err)
	err = cm.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", cm)
	assert.Equal(t, int32(1), cm.GetControlMessage().GetParameterType().GetRanParameterId().GetValue())
	assert.Equal(t, "pci", cm.GetControlMessage().GetParameterType().GetRanParameterName().GetValue())
	assert.Equal(t, e2sm_rc_pre_v2.RanparameterType_RANPARAMETER_TYPE_INTEGER, cm.GetControlMessage().GetParameterType().GetRanParameterType())
	assert.Equal(t, uint32(6), cm.GetControlMessage().GetParameterVal().GetValueInt())

	// Printing message in XER form
	xer, err := rcprectypes.XerEncodeE2SmRcPreControlMessage(cm)
	assert.NilError(t, err)
	t.Logf("ControlMessage XER is \n%s", xer)
}

func Test_DecodeControlOutcomeBytesRadisys(t *testing.T) {
	coPer, err := hexlib.Asn1BytesToByte(coBytes)
	assert.NilError(t, err)
	t.Logf("Radisys Bytes for ControlOutcome in HEX format are \n%v\n", hex.Dump(coPer))
	co, err := rcprectypes.PerDecodeE2SmRcPreControlOutcome(coPer)
	assert.NilError(t, err)
	err = co.Validate()
	assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", co)
	assert.Equal(t, int32(1), co.GetControlOutcomeFormat1().GetOutcomeElementList()[0].GetRanParameterId().GetValue())

	// Printing message in XER form
	xer, err := rcprectypes.XerEncodeE2SmRcPreControlOutcome(co)
	assert.NilError(t, err)
	t.Logf("ControlOutcome XER is \n%s", xer)
}
