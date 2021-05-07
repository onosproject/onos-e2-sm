// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	"encoding/hex"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"gotest.tools/assert"
	"testing"
)

func createE2SmMhoMeasurementReportItemMsg() (*e2sm_mho.E2SmMhoMeasurementReportItem, error) {

	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)
	e2SmMhoMeasurementReportItem := e2sm_mho.E2SmMhoMeasurementReportItem{
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
	}

	if err := e2SmMhoMeasurementReportItem.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmMhoMeasurementReportItem %s", err.Error())
	}
	return &e2SmMhoMeasurementReportItem, nil
}

func Test_xerEncodingE2SmMhoMeasurementReportItem(t *testing.T) {

	e2SmMhoMeasurementReportItem, err := createE2SmMhoMeasurementReportItemMsg()
	assert.NilError(t, err, "Error creating E2SmMhoMeasurementReportItem PDU")

	xer, err := xerEncodeE2SmMhoMeasurementReportItem(e2SmMhoMeasurementReportItem)
	assert.NilError(t, err)
	assert.Equal(t, 313, len(xer))
	t.Logf("E2SmMhoMeasurementReportItem XER\n%s", string(xer))

	result, err := xerDecodeE2SmMhoMeasurementReportItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoMeasurementReportItem XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoMeasurementReportItem.GetCgi(), result.GetCgi())
	assert.Equal(t, e2SmMhoMeasurementReportItem.GetRsrp().GetValue(), result.GetRsrp().GetValue())

}

func Test_perEncodingE2SmMhoMeasurementReportItem(t *testing.T) {

	e2SmMhoMeasurementReportItem, err := createE2SmMhoMeasurementReportItemMsg()
	assert.NilError(t, err, "Error creating E2SmMhoMeasurementReportItem PDU")

	per, err := perEncodeE2SmMhoMeasurementReportItem(e2SmMhoMeasurementReportItem)
	assert.NilError(t, err)
	assert.Equal(t, 11, len(per))
	t.Logf("E2SmMhoMeasurementReportItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2SmMhoMeasurementReportItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoMeasurementReportItem PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoMeasurementReportItem.GetCgi(), result.GetCgi())
	assert.Equal(t, e2SmMhoMeasurementReportItem.GetRsrp().GetValue(), result.GetRsrp().GetValue())

}
