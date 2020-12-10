// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

// This function encodes PmContainersList item.
// Containers can have different types. Value of PfContainer contains only one certain type of a container.
// It doesn't allow to assign more containers to the same structures, what leads us to encoding each of the
// container separately.
func Test_xerEncodeRanContainer(t *testing.T) {

	value := []byte{0x22, 0x21}

	ranContainer := &e2sm_kpm_ies.RanContainer{
		Value: value,
	}

	xer, err := xerEncodeRanContainer(ranContainer)
	assert.NilError(t, err)
	t.Logf("RAN-Container XER\n%s", string(xer))
}
