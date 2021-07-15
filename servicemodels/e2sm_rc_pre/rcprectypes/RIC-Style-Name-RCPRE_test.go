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

func Test_xerEncodeRicStyleName(t *testing.T) {

	name := "RicStyleName"
	ricStyleName := &e2sm_rc_pre_v2.RicStyleName{
		Value: name,
	}

	xer, err := xerEncodeRicStyleName(ricStyleName)
	assert.NilError(t, err)
	//assert.Equal(t, 46, len(xer))
	t.Logf("RIC-Style-Name XER\n%s", string(xer))
}

func Test_xerDecodeRicStyleName(t *testing.T) {

	name := "RicStyleName"
	ricStyleName := &e2sm_rc_pre_v2.RicStyleName{
		Value: name,
	}

	xer, err := xerEncodeRicStyleName(ricStyleName)
	assert.NilError(t, err)
	//assert.Equal(t, 46, len(xer))
	t.Logf("RIC-Style-Name XER\n%s", string(xer))

	result, err := xerDecodeRicStyleName(xer)
	assert.NilError(t, err)
	assert.Equal(t, ricStyleName.Value, result.Value, "Encoded and decoded values are not the same")
	t.Logf("RIC-Style-Name XER - decoded\n%v", result)
}

func Test_perEncodeRicStyleName(t *testing.T) {

	name := "RicStyleName"
	ricStyleName := &e2sm_rc_pre_v2.RicStyleName{
		Value: name,
	}

	per, err := perEncodeRicStyleName(ricStyleName)
	assert.NilError(t, err)
	assert.Equal(t, 14, len(per))
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))
}

func Test_perDecodeRicStyleName(t *testing.T) {

	name := "RicStyleName"
	ricStyleName := &e2sm_rc_pre_v2.RicStyleName{
		Value: name,
	}

	per, err := perEncodeRicStyleName(ricStyleName)
	assert.NilError(t, err)
	assert.Equal(t, 14, len(per))
	t.Logf("RIC-Style-Name PER\n%v", hex.Dump(per))

	result, err := perDecodeRicStyleName(per)
	assert.NilError(t, err)
	assert.Equal(t, ricStyleName.Value, result.Value, "Encoded and decoded values are not the same")
	t.Logf("RIC-Style-Name PER - decoded\n%v", result)
}
