// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	"encoding/hex"
	"fmt"
	pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdubuilder"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"gotest.tools/assert"
	"testing"
)

func createRicControlMessagePriorityMsg() (*e2sm_mho.RicControlMessagePriority, error) {

	// ricControlMessagePriority := pdubuilder.CreateRicControlMessagePriority() //ToDo - fill in arguments here(if this function exists

	ricControlMessagePriority := e2sm_mho.RicControlMessagePriority{
		Value: nil,
	}

	if err := ricControlMessagePriority.Validate(); err != nil {
		return nil, fmt.Errorf("error validating RicControlMessagePriority %s", err.Error())
	}
	return &ricControlMessagePriority, nil
}

func Test_xerEncodingRicControlMessagePriority(t *testing.T) {

	ricControlMessagePriority, err := createRicControlMessagePriorityMsg()
	assert.NilError(t, err, "Error creating RicControlMessagePriority PDU")

	xer, err := xerEncodeRicControlMessagePriority(ricControlMessagePriority)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("RicControlMessagePriority XER\n%s", string(xer))

	result, err := xerDecodeRicControlMessagePriority(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicControlMessagePriority XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, ricControlMessagePriority.GetValue(), result.GetValue())

}

func Test_perEncodingRicControlMessagePriority(t *testing.T) {

	ricControlMessagePriority, err := createRicControlMessagePriorityMsg()
	assert.NilError(t, err, "Error creating RicControlMessagePriority PDU")

	per, err := perEncodeRicControlMessagePriority(ricControlMessagePriority)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("RicControlMessagePriority PER\n%v", hex.Dump(per))

	result, err := perDecodeRicControlMessagePriority(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicControlMessagePriority PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, ricControlMessagePriority.GetValue(), result.GetValue())

}
