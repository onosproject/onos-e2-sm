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

func Test_xerEncodeTimeStamp(t *testing.T) {

	stamp := []byte{0x01, 0x02, 0x03, 0x04}
	timeStamp := &e2sm_kpm_v2.TimeStamp{
		Value: stamp,
	}
	xer, err := xerEncodeTimeStamp(timeStamp)
	assert.NilError(t, err)
	//assert.Equal(t, 35, len(xer))
	t.Logf("TimeStamp XER\n%s", string(xer))
}

func Test_xerDecodeTimeStamp(t *testing.T) {

	stamp := []byte{0x01, 0x02, 0x03, 0x04}
	timeStamp := &e2sm_kpm_v2.TimeStamp{
		Value: stamp,
	}
	xer, err := xerEncodeTimeStamp(timeStamp)
	assert.NilError(t, err)
	//assert.Equal(t, 35, len(xer))
	t.Logf("TimeStamp XER\n%s", string(xer))

	result, err := xerDecodeTimeStamp(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TimeStamp XER - decoded\n%s", result)
}

func Test_perEncodeTimeStamp(t *testing.T) {

	stamp := []byte{0x01, 0x02, 0x03, 0x04}
	timeStamp := &e2sm_kpm_v2.TimeStamp{
		Value: stamp,
	}
	per, err := perEncodeTimeStamp(timeStamp)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("TimeStamp PER\n%v", hex.Dump(per))
}

func Test_perDecodeTimeStamp(t *testing.T) {

	stamp := []byte{0x01, 0x02, 0x03, 0x04}
	timeStamp := &e2sm_kpm_v2.TimeStamp{
		Value: stamp,
	}
	per, err := perEncodeTimeStamp(timeStamp)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("TimeStamp PER\n%v", hex.Dump(per))

	result, err := perDecodeTimeStamp(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TimeStamp PER - decoded\n%v", result)
}
