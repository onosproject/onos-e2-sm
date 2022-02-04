// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodePFContainer(t *testing.T) {

	var testGnbCuCpName = "0xAB"

	pfContainer := &e2sm_kpm_ies.PfContainer{
		PfContainer: &e2sm_kpm_ies.PfContainer_OCuCp{
			OCuCp: &e2sm_kpm_ies.OcucpPfContainer{
				GNbCuCpName: &e2sm_kpm_ies.GnbCuCpName{
					Value: testGnbCuCpName, //string -- PrintableString(SIZE(1..150,...))
				},
				CuCpResourceStatus: &e2sm_kpm_ies.OcucpPfContainer_CuCpResourceStatus001{
					NumberOfActiveUes: 1, //int32 -- INTEGER (1..65536, ...)
				},
			},
			//TODO: Implement other types of containers
		},
	}

	xer, err := xerEncodePFContainer(pfContainer)
	assert.NilError(t, err)
	t.Logf("PF-Container XER\n%s", string(xer))
}
