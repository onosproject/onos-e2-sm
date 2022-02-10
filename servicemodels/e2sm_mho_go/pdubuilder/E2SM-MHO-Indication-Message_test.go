// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/encoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmMhoIndicationMsgF1(t *testing.T) {

	ueID, err := CreateUeIDGNb(1, []byte{0x01, 0x02, 0x03}, []byte{0xFF}, []byte{0xFF, 0xC0}, []byte{0xFC})
	assert.NilError(t, err)

	cgi, err := CreateCgiNrCGI([]byte{0xAA, 0xFD, 0xD4}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x40, 0x00},
		Len:   36,
	})
	assert.NilError(t, err)
	rsrp := &e2sm_mho_go.Rsrp{
		Value: 1234,
	}
	measItem, err := CreateMeasurementRecordItem(cgi, rsrp)
	assert.NilError(t, err)
	measItem.SetFiveQi(21)

	measReport := make([]*e2sm_mho_go.E2SmMhoMeasurementReportItem, 0)
	measReport = append(measReport, measItem)

	newE2SmMhoPdu, err := CreateE2SmMhoIndicationMsgFormat1(ueID, measReport)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)
	t.Logf("E2SM-MHO-IndicationMessage is \n%v", newE2SmMhoPdu)

	per, err := encoder.PerEncodeE2SmMhoIndicationMessage(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded E2SM-MHO-IndicationMessage: \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMhoIndicationMessage(per)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-MHO-IndicationMessage is \n%v", result)
	assert.DeepEqual(t, newE2SmMhoPdu.String(), result.String())
}

func TestE2SmMhoIndicationMsgF2(t *testing.T) {

	ueID, err := CreateUeIDGNb(1, []byte{0x01, 0x02, 0x03}, []byte{0xFF}, []byte{0xFF, 0xC0}, []byte{0xFC})
	assert.NilError(t, err)

	newE2SmMhoPdu, err := CreateE2SmMhoIndicationMsgFormat2(ueID, CreateRrcStatusConnected())
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMhoPdu != nil)

	per, err := encoder.PerEncodeE2SmMhoIndicationMessage(newE2SmMhoPdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded E2SM-MHO-IndicationMessage: \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMhoIndicationMessage(per)
	assert.NilError(t, err)
	t.Logf("Decoded E2SM-MHO-IndicationMessage is \n%v", result)
	assert.Equal(t, newE2SmMhoPdu.String(), result.String())
}
