// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodePCIRange(t *testing.T) {

	var lowerPCI int32 = 10
	var upperPCI int32 = 11

	pciRange := &e2sm_rc_pre_ies.PciRange{
		LowerPci: &e2sm_rc_pre_ies.Pci{
			Value: lowerPCI,
		},
		UpperPci: &e2sm_rc_pre_ies.Pci{
			Value: upperPCI,
		},
	}

	xer, err := xerEncodePCIRange(pciRange)
	assert.NilError(t, err)
	t.Logf("PCIRange XER\n%s", string(xer))
}
