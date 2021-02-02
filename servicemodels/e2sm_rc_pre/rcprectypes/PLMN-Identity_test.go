// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodePlmnIdentity(t *testing.T) {

	var plmnID = "ONF"

	plmnIdentity := &e2sm_rc_pre_ies.PlmnIdentity{
		Value: []byte(plmnID),
	}

	xer, err := xerEncodePlmnIdentity(plmnIdentity)
	assert.NilError(t, err)
	t.Logf("PLMN-Identity XER\n%s", string(xer))
}
