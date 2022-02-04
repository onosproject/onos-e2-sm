// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerDecodeENGNbID(t *testing.T) {

	enGNbID := &e2sm_kpm_ies.EngnbId{
		EngnbId: &e2sm_kpm_ies.EngnbId_GNbId{
			GNbId: &e2sm_kpm_ies.BitString{
				Value: 0x9bcd4, //uint64
				Len:   22,      //uint32
			},
		},
	}

	xer, err := xerEncodeENGNbID(enGNbID)
	assert.NilError(t, err)
	t.Logf("ENGNb XER\n%s", string(xer))

	result, err := xerDecodeENGNbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
