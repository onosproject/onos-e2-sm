// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeArp(t *testing.T) {

	arp := &e2sm_kpm_v2.Arp{
		Value: 15,
	}
	xer, err := xerEncodeArp(arp)
	assert.NilError(t, err)
	//assert.Equal(t, 14, len(xer))
	t.Logf("ARP XER\n%s", string(xer))
}

func Test_xerDecodeArp(t *testing.T) {

	arp := &e2sm_kpm_v2.Arp{
		Value: 15,
	}
	xer, err := xerEncodeArp(arp)
	assert.NilError(t, err)
	//assert.Equal(t, 14, len(xer))
	t.Logf("ARP XER\n%s", string(xer))

	result, err := xerDecodeArp(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ARP XER - decoded\n%s", result)
}

func Test_perEncodeArp(t *testing.T) {

	arp := &e2sm_kpm_v2.Arp{
		Value: 15,
	}
	per, err := perEncodeArp(arp)
	assert.NilError(t, err)
	//assert.Equal(t, 1, len(per))
	t.Logf("ARP PER\n%s", string(per))
}

func Test_perDecodeArp(t *testing.T) {

	arp := &e2sm_kpm_v2.Arp{
		Value: 15,
	}
	per, err := perEncodeArp(arp)
	assert.NilError(t, err)
	//assert.Equal(t, 1, len(per))
	t.Logf("ARP PER\n%s", string(per))

	result, err := perDecodeArp(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ARP PER - decoded\n%s", result)
}
