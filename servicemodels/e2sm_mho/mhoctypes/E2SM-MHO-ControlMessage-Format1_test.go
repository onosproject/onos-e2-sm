// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mhoctypes

import (
	"encoding/hex"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func createE2SmMhoControlMessageFormat1Msg() (*e2sm_mho.E2SmMhoControlMessageFormat1, error) {
	var plmnID = "12f410"
	plmnIDBytes, _ := hex.DecodeString(plmnID)

	servingCgi := &e2sm_mho.CellGlobalId{
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
	}

	ueID := &e2sm_mho.UeIdentity{
		Value: "1234",
	}

	targetCgi := &e2sm_mho.CellGlobalId{
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
	}

	e2SmMhoControlMessageFormat1 := e2sm_mho.E2SmMhoControlMessageFormat1{
		ServingCgi: servingCgi,
		UedId:      ueID,
		TargetCgi:  targetCgi,
	}

	if err := e2SmMhoControlMessageFormat1.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmMhoControlMessageFormat1 %s", err.Error())
	}
	return &e2SmMhoControlMessageFormat1, nil
}

func Test_xerEncodingE2SmMhoControlMessageFormat1(t *testing.T) {

	e2SmMhoControlMessageFormat1, err := createE2SmMhoControlMessageFormat1Msg()
	assert.NilError(t, err, "Error creating E2SmMhoControlMessageFormat1 PDU")

	xer, err := xerEncodeE2SmMhoControlMessageFormat1(e2SmMhoControlMessageFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 578, len(xer))
	t.Logf("E2SmMhoControlMessageFormat1 XER\n%s", string(xer))

	result, err := xerDecodeE2SmMhoControlMessageFormat1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoControlMessageFormat1 XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoControlMessageFormat1.GetServingCgi(), result.GetServingCgi())
	//assert.Equal(t, e2SmMhoControlMessageFormat1.GetUedId(), result.GetUedId())
	//assert.Equal(t, e2SmMhoControlMessageFormat1.GetTargetCgi(), result.GetTargetCgi())

}

func Test_perEncodingE2SmMhoControlMessageFormat1(t *testing.T) {

	e2SmMhoControlMessageFormat1, err := createE2SmMhoControlMessageFormat1Msg()
	assert.NilError(t, err, "Error creating E2SmMhoControlMessageFormat1 PDU")

	per, err := perEncodeE2SmMhoControlMessageFormat1(e2SmMhoControlMessageFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 21, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("E2SmMhoControlMessageFormat1 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2SmMhoControlMessageFormat1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoControlMessageFormat1 PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoControlMessageFormat1.GetServingCgi(), result.GetServingCgi())
	//assert.Equal(t, e2SmMhoControlMessageFormat1.GetUedId(), result.GetUedId())
	//assert.Equal(t, e2SmMhoControlMessageFormat1.GetTargetCgi(), result.GetTargetCgi())

}
