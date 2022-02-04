// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerDecodeENbID(t *testing.T) {

	enbID := &e2sm_kpm_ies.EnbId{
		EnbId: &e2sm_kpm_ies.EnbId_HomeENbId{
			HomeENbId: &e2sm_kpm_ies.BitString{
				Value: 0x9bcd4, //uint64
				Len:   22,      //uint32
			},
		},
	}

	xer, err := xerEncodeENbID(enbID)
	assert.NilError(t, err)
	t.Logf("ENb XER\n%s", string(xer))

	result, err := xerDecodeENbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
