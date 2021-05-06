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

func createRsrpMsg(rsrpVal int32) (*e2sm_mho.Rsrp, error) {

	// rsrp := pdubuilder.CreateRsrp() //ToDo - fill in arguments here(if this function exists

	rsrp := e2sm_mho.Rsrp{
		Value: rsrpVal,
	}

	if err := rsrp.Validate(); err != nil {
		return nil, fmt.Errorf("error validating Rsrp %s", err.Error())
	}
	return &rsrp, nil
}

func Test_xerEncodingRsrp(t *testing.T) {

	rsrp, err := createRsrpMsg(1234)
	assert.NilError(t, err, "Error creating Rsrp PDU")

	xer, err := xerEncodeRsrp(rsrp)
	assert.NilError(t, err)
	assert.Equal(t, 18, len(xer))
	t.Logf("Rsrp XER\n%s", string(xer))

	result, err := xerDecodeRsrp(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Rsrp XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, rsrp.GetValue(), result.GetValue())

}

func Test_perEncodingRsrp(t *testing.T) {

	rsrp, err := createRsrpMsg(1234)
	assert.NilError(t, err, "Error creating Rsrp PDU")

	per, err := perEncodeRsrp(rsrp)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("Rsrp PER\n%v", hex.Dump(per))

	result, err := perDecodeRsrp(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Rsrp PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, rsrp.GetValue(), result.GetValue())

}
