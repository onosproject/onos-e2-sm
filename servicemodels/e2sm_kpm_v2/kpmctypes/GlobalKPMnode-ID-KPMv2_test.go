// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"encoding/hex"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func createGlobalKpmnodeID1() *e2sm_kpm_v2.GlobalKpmnodeId {

	return &e2sm_kpm_v2.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2.GlobalKpmnodeId_GNb{
			GNb: &e2sm_kpm_v2.GlobalKpmnodeGnbId{
				GlobalGNbId: &e2sm_kpm_v2.GlobalgNbId{
					GnbId: &e2sm_kpm_v2.GnbIdChoice{
						GnbIdChoice: &e2sm_kpm_v2.GnbIdChoice_GnbId{
							GnbId: &e2sm_kpm_v2.BitString{
								Value: 0x9bcd4,
								Len:   28,
							},
						},
					},
					PlmnId: &e2sm_kpm_v2.PlmnIdentity{
						Value: []byte{0x21, 0x22, 0x23},
					},
				},
				GNbCuUpId: &e2sm_kpm_v2.GnbCuUpId{
					Value: 21,
				},
				GNbDuId: &e2sm_kpm_v2.GnbDuId{
					Value: 32,
				},
			},
		},
	}
}

func createGlobalKpmnodeID2() *e2sm_kpm_v2.GlobalKpmnodeId {

	return &e2sm_kpm_v2.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2.GlobalKpmnodeId_EnGNb{
			EnGNb: &e2sm_kpm_v2.GlobalKpmnodeEnGnbId{
				GlobalGNbId: &e2sm_kpm_v2.GlobalenGnbId{
					GNbId: &e2sm_kpm_v2.EngnbId{
						EngnbId: &e2sm_kpm_v2.EngnbId_GNbId{
							GNbId: &e2sm_kpm_v2.BitString{
								Value: 0x9bcd4,
								Len:   28,
							},
						},
					},
					PLmnIdentity: &e2sm_kpm_v2.PlmnIdentity{
						Value: []byte{0x21, 0x22, 0x23},
					},
				},
				GNbCuUpId: &e2sm_kpm_v2.GnbCuUpId{
					Value: 21,
				},
				GNbDuId: &e2sm_kpm_v2.GnbDuId{
					Value: 32,
				},
			},
		},
	}
}

func createGlobalKpmnodeID3() *e2sm_kpm_v2.GlobalKpmnodeId {

	return &e2sm_kpm_v2.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2.GlobalKpmnodeId_NgENb{
			NgENb: &e2sm_kpm_v2.GlobalKpmnodeNgEnbId{
				GlobalNgENbId: &e2sm_kpm_v2.GlobalngeNbId{
					EnbId: &e2sm_kpm_v2.EnbIdChoice{
						EnbIdChoice: &e2sm_kpm_v2.EnbIdChoice_EnbIdMacro{
							EnbIdMacro: &e2sm_kpm_v2.BitString{
								Value: 0x9bcd4,
								Len:   20,
							},
						},
					},
					PlmnId: &e2sm_kpm_v2.PlmnIdentity{
						Value: []byte{0x21, 0x22, 0x23},
					},
					ShortMacroENbId: &e2sm_kpm_v2.BitString{
						Value: 0x9bcd4,
						Len:   18,
					},
					LongMacroENbId: &e2sm_kpm_v2.BitString{
						Value: 0x9bcd4,
						Len:   21,
					},
				},
				GNbDuId: &e2sm_kpm_v2.GnbDuId{
					Value: 32,
				},
			},
		},
	}
}

func createGlobalKpmnodeID4() *e2sm_kpm_v2.GlobalKpmnodeId {

	return &e2sm_kpm_v2.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2.GlobalKpmnodeId_ENb{
			ENb: &e2sm_kpm_v2.GlobalKpmnodeEnbId{
				GlobalENbId: &e2sm_kpm_v2.GlobalEnbId{
					ENbId: &e2sm_kpm_v2.EnbId{
						EnbId: &e2sm_kpm_v2.EnbId_HomeENbId{
							HomeENbId: &e2sm_kpm_v2.BitString{
								Value: 0x9bcd4,
								Len:   28,
							},
						},
					},
					PLmnIdentity: &e2sm_kpm_v2.PlmnIdentity{
						Value: []byte{0x21, 0x22, 0x23},
					},
				},
			},
		},
	}
}

func Test_xerEncodeGlobalKpmnodeID(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID1()

	xer, err := xerEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 370, len(xer))
	t.Logf("GlobalKpmnodeID (GNb) XER\n%s", string(xer))

	globalKpmnodeID = createGlobalKpmnodeID2()

	xer, err = xerEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 388, len(xer))
	t.Logf("GlobalKpmnodeID (enGNb) XER\n%s", string(xer))

	globalKpmnodeID = createGlobalKpmnodeID3()

	xer, err = xerEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 551, len(xer))
	t.Logf("GlobalKpmnodeID (ngENb) XER\n%s", string(xer))

	globalKpmnodeID = createGlobalKpmnodeID4()

	xer, err = xerEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 318, len(xer))
	t.Logf("GlobalKpmnodeID (ENb) XER\n%s", string(xer))
}

func Test_xerDecodeGlobalKpmnodeID(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID1()

	xer, err := xerEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 370, len(xer))
	t.Logf("GlobalKpmnodeID (GNb) XER\n%s", string(xer))

	result, err := xerDecodeGlobalKpmnodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeID (GNb) XER - decoded\n%s", result)

	globalKpmnodeID = createGlobalKpmnodeID2()

	xer, err = xerEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 388, len(xer))
	t.Logf("GlobalKpmnodeID (enGNb) XER\n%s", string(xer))

	result, err = xerDecodeGlobalKpmnodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeID (enGNb) XER - decoded\n%s", result)

	globalKpmnodeID = createGlobalKpmnodeID3()

	xer, err = xerEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 551, len(xer))
	t.Logf("GlobalKpmnodeID (ngENb) XER\n%s", string(xer))

	result, err = xerDecodeGlobalKpmnodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeID (ngENb) XER - decoded\n%s", result)

	globalKpmnodeID = createGlobalKpmnodeID4()

	xer, err = xerEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 318, len(xer))
	t.Logf("GlobalKpmnodeID (ENb) XER\n%s", string(xer))

	result, err = xerDecodeGlobalKpmnodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeID (ENb) XER - decoded\n%s", result)
}

func Test_perEncodeGlobalKpmnodeID(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID1()

	per, err := perEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 12, len(per))
	t.Logf("GlobalKpmnodeID (GNb) PER\n%v", hex.Dump(per))

	globalKpmnodeID = createGlobalKpmnodeID2()

	per, err = perEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 12, len(per))
	t.Logf("GlobalKpmnodeID (enGNb) PER\n%v", hex.Dump(per))

	globalKpmnodeID = createGlobalKpmnodeID3()

	per, err = perEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 15, len(per))
	t.Logf("GlobalKpmnodeID (ngENb) PER\n%v", hex.Dump(per))

	globalKpmnodeID = createGlobalKpmnodeID4()

	per, err = perEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 9, len(per))
	t.Logf("GlobalKpmnodeID (ENb) PER\n%v", hex.Dump(per))
}

func Test_perDecodeGlobalKpmnodeID(t *testing.T) {

	globalKpmnodeID := createGlobalKpmnodeID1()

	per, err := perEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 12, len(per))
	t.Logf("GlobalKpmnodeID (GNb) PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalKpmnodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeID (GNb) PER - decoded\n%v", result)

	globalKpmnodeID = createGlobalKpmnodeID2()

	per, err = perEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 12, len(per))
	t.Logf("GlobalKpmnodeID (enGNb) PER\n%v", hex.Dump(per))

	result, err = perDecodeGlobalKpmnodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeID (enGNb) PER - decoded\n%v", result)

	globalKpmnodeID = createGlobalKpmnodeID3()

	per, err = perEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 15, len(per))
	t.Logf("GlobalKpmnodeID (ngENb) PER\n%v", hex.Dump(per))

	result, err = perDecodeGlobalKpmnodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeID (ngENb) PER - decoded\n%v", result)

	globalKpmnodeID = createGlobalKpmnodeID4()

	per, err = perEncodeGlobalKpmnodeID(globalKpmnodeID)
	assert.NilError(t, err)
	//assert.Equal(t, 9, len(per))
	t.Logf("GlobalKpmnodeID (ENb) PER\n%v", hex.Dump(per))

	result, err = perDecodeGlobalKpmnodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeID (ENb) PER - decoded\n%v", result)
}
