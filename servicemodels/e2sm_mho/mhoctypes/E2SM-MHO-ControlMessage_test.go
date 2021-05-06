// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdubuilder"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func Test_XerEncodeE2SmMhoControlMessage(t *testing.T) {
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

	e2SmMhoControlMessage, err := pdubuilder.CreateE2SmMhoControlMessage(servingCgi, ueID, targetCgi)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmMhoControlMessage(e2SmMhoControlMessage)
	assert.NilError(t, err)
	assert.Equal(t, 689, len(xer))
	t.Logf("E2SM-MHO-ControlMessage XER\n%s", string(xer))
}

func Test_XerDecodeE2SmMhoControlMessage(t *testing.T) {
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

	e2SmMhoControlMessage, err := pdubuilder.CreateE2SmMhoControlMessage(servingCgi, ueID, targetCgi)
	assert.NilError(t, err)

	xer, err := XerEncodeE2SmMhoControlMessage(e2SmMhoControlMessage)
	assert.NilError(t, err)
	assert.Equal(t, 689, len(xer))
	t.Logf("E2SM-MHO-ControlMessage XER\n%s", string(xer))

	result, err := XerDecodeE2SmMhoControlMessage(xer)
	assert.NilError(t, err)
	t.Logf("E2SM-MHO-ControlMessage decoded XER is\n%v", result)

	//assert.Equal(t, e2SmMhoControlMessage.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}

func Test_PerDecodeE2SmMhoControlMessage(t *testing.T) {
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

	e2SmMhoControlMessage, err := pdubuilder.CreateE2SmMhoControlMessage(servingCgi, ueID, targetCgi)
	assert.NilError(t, err)

	per, err := PerEncodeE2SmMhoControlMessage(e2SmMhoControlMessage)
	assert.NilError(t, err)
	assert.Equal(t, 21, len(per))
	t.Logf("E2SM-MHO-ControlMessage PER\n%v", hex.Dump(per))

	result, err := PerDecodeE2SmMhoControlMessage(per)
	assert.NilError(t, err)
	t.Logf("E2SM-MHO-ControlMessage decoded XER is\n%v", result)

	//assert.Equal(t, e2SmMhoControlMessage.GetRicStyleType().GetValue(), result.GetRicStyleType().GetValue(), "Encoded and decoded values are not the same")
}
