// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeSNSSAI(t *testing.T) {

	var sst = "ONF"
	var sd = "ONF"

	sliceID := &e2sm_kpm_ies.Snssai{
		SSt: []byte(sst),
		SD:  []byte(sd),
	}

	xer, err := xerEncodeSnssai(sliceID)
	assert.NilError(t, err)
	t.Logf("SNSSAI XER\n%s", string(xer))
}
