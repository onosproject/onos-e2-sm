// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

import (
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodePCI(t *testing.T) {

	var pci int32 = 252

	c := &e2sm_rc_pre_v2.Pci{
		Value: pci,
	}

	xer, err := xerEncodePCI(c)
	assert.NilError(t, err)
	t.Logf("PCI XER\n%s", string(xer))
}

func Test_decodePCI(t *testing.T) {

	var pci int32 = 252

	c := &e2sm_rc_pre_v2.Pci{
		Value: pci,
	}

	pciC := newPCI(c)
	result := decodePCI(pciC)
	assert.Equal(t, pci, result.Value, "Something went wrong, comparison is incorrect")
}

func Test_perEncodePCI(t *testing.T) {

	var pci int32 = 252

	c := &e2sm_rc_pre_v2.Pci{
		Value: pci,
	}

	pciPer, err := perEncodePCI(c)
	assert.NilError(t, err)
	assert.Assert(t, pciPer != nil)
	t.Logf("PCI PER\n%v", pciPer)
}

func Test_perDecodePCI(t *testing.T) {

	var pci int32 = 252

	c := &e2sm_rc_pre_v2.Pci{
		Value: pci,
	}

	pciPer, err := perEncodePCI(c)
	assert.NilError(t, err)
	assert.Assert(t, pciPer != nil)
	t.Logf("PCI PER\n%v", pciPer)

	result, err := perDecodePCI(pciPer)
	assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	t.Logf("PCI PER decoded is \n%v", result)

	assert.Equal(t, pci, result.Value, "Encoded and decoded values are not the same")
}
