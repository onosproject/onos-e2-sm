// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	"encoding/hex"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeRicStyleType(t *testing.T) {

	var ricType int32 = 11
	ricStyleType := &e2sm_rc_pre_v2.RicStyleType{
		Value: ricType,
	}

	xer, err := xerEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	//assert.Equal(t, 36, len(xer))
	t.Logf("RIC-Style-Type XER\n%s", string(xer))
}

func Test_xerDecodeRicStyleType(t *testing.T) {

	var ricType int32 = 11
	ricStyleType := &e2sm_rc_pre_v2.RicStyleType{
		Value: ricType,
	}

	xer, err := xerEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	//assert.Equal(t, 36, len(xer))
	t.Logf("RIC-Style-Type XER\n%s", string(xer))

	result, err := xerDecodeRicStyleType(xer)
	assert.NilError(t, err)
	assert.Equal(t, ricStyleType.Value, result.Value, "Encoded and decoded values are not the same")
	t.Logf("RIC-Style-Type XER - decoded\n%v", result)
}

func Test_perEncodeRicStyleType(t *testing.T) {

	var ricType int32 = 11
	ricStyleType := &e2sm_rc_pre_v2.RicStyleType{
		Value: ricType,
	}

	per, err := perEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("RIC-Style-Type PER\n%v", hex.Dump(per))
}

func Test_perDecodeRicStyleType(t *testing.T) {

	var ricType int32 = 11
	ricStyleType := &e2sm_rc_pre_v2.RicStyleType{
		Value: ricType,
	}

	per, err := perEncodeRicStyleType(ricStyleType)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("RIC-Style-Type PER\n%s", hex.Dump(per))

	result, err := perDecodeRicStyleType(per)
	assert.NilError(t, err)
	t.Logf("RIC-Style-Type PER - decoded\n%v", result)
	assert.Equal(t, ricStyleType.Value, result.Value, "Encoded and decoded values are not the same")
}
