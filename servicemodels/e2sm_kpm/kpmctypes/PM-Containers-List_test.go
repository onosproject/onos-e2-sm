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
func Test_xerEncodePmContainersListItem(t *testing.T) {

	var testGnbCuCpName = "0xAB"
	var testNumberOfActiveUes int32 = 1

	pmContainersList := &e2sm_kpm_ies.PmContainersList{
		PerformanceContainer: &e2sm_kpm_ies.PfContainer{
			PfContainer: &e2sm_kpm_ies.PfContainer_OCuCp{
				OCuCp: &e2sm_kpm_ies.OcucpPfContainer{
					GNbCuCpName: &e2sm_kpm_ies.GnbCuCpName{
						Value: testGnbCuCpName, //string -- PrintableString(SIZE(1..150,...))
					},
					CuCpResourceStatus: &e2sm_kpm_ies.OcucpPfContainer_CuCpResourceStatus001{
						NumberOfActiveUes: testNumberOfActiveUes, //int32 -- INTEGER (1..65536, ...)
					},
				},
				//TODO: Implement other types of containers
			},
		},
		//TODO: Implement TheRancontainer
	}

	xer, err := xerEncodePmContainersListItem(pmContainersList)
	assert.NilError(t, err)
	t.Logf("PF-Container XER\n%s", string(xer))
}
