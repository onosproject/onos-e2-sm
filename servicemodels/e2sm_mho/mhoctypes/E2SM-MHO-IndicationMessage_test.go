// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

import (
	"encoding/hex"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" // ToDo - Make imports more dynamic
	"gotest.tools/assert"
	"testing"
)

func createE2SmMhoIndicationMessageMsg() (*e2sm_mho.E2SmMhoIndicationMessage, error) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	measReport := []*e2sm_mho.E2SmMhoMeasurementReportItem{
		{
			Cgi: &e2sm_mho.CellGlobalId{
				CellGlobalId: &e2sm_mho.CellGlobalId_EUtraCgi{
					EUtraCgi: &e2sm_mho.Eutracgi{
						PLmnIdentity: &e2sm_mho.PlmnIdentity{
							Value: plmnIDBytes,
						},
						EUtracellIdentity: &e2sm_mho.EutracellIdentity{
							Value: &e2sm_mho.BitString{
								Value: 0x9bcd4ab, //uint64
								Len:   28,        //uint32
							},
						},
					},
				},
			},
			Rsrp: &e2sm_mho.Rsrp{
				Value: 1234,
			},
		},
	}
	e2SmMhoIndicationMessage := e2sm_mho.E2SmMhoIndicationMessage{
		E2SmMhoIndicationMessage: &e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat1{
			IndicationMessageFormat1: &e2sm_mho.E2SmMhoIndicationMessageFormat1{
				UeId:       &e2sm_mho.UeIdentity{Value: "1234"},
				MeasReport: measReport,
			},
		},
	}

	if err := e2SmMhoIndicationMessage.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmMhoIndicationMessage %s", err.Error())
	}
	return &e2SmMhoIndicationMessage, nil
}

func Test_xerEncodingE2SmMhoIndicationMessage(t *testing.T) {

	e2SmMhoIndicationMessage, err := createE2SmMhoIndicationMessageMsg()
	assert.NilError(t, err, "Error creating E2SmMhoIndicationMessage PDU")

	xer, err := XerEncodeE2SmMhoIndicationMessage(e2SmMhoIndicationMessage)
	assert.NilError(t, err)
	assert.Equal(t, 645, len(xer))
	t.Logf("E2SmMhoIndicationMessage XER\n%s", string(xer))

	result, err := XerDecodeE2SmMhoIndicationMessage(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoIndicationMessage XER - decoded\n%v", result)
	// ToDo - adjust field's verification
	//assert.DeepEqual(t, e2SmMhoIndicationMessage.GetIndicationMessageFormat1().GetValue(), result.GetIndicationMessageFormat1().GetValue())
	//assert.DeepEqual(t, e2SmMhoIndicationMessage.GetIndicationMessageFormat2().GetValue(), result.GetIndicationMessageFormat2().GetValue())

}

func Test_perEncodingE2SmMhoIndicationMessage(t *testing.T) {

	e2SmMhoIndicationMessage, err := createE2SmMhoIndicationMessageMsg()
	assert.NilError(t, err, "Error creating E2SmMhoIndicationMessage PDU")

	per, err := PerEncodeE2SmMhoIndicationMessage(e2SmMhoIndicationMessage)
	assert.NilError(t, err)
	assert.Equal(t, 18, len(per))
	t.Logf("E2SmMhoIndicationMessage PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmMhoIndicationMessage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoIndicationMessage PER - decoded\n%v", result)
	// ToDo - adjust field's verification

	//This one if for OneOf fields

	//assert.DeepEqual(t, e2SmMhoIndicationMessage.GetIndicationMessageFormat1().GetValue(), result.GetIndicationMessageFormat1().GetValue())
	//assert.DeepEqual(t, e2SmMhoIndicationMessage.GetIndicationMessageFormat2().GetValue(), result.GetIndicationMessageFormat2().GetValue())

}
