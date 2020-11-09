// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeOCUCPPFContainer(t *testing.T) {

	//c := &e2sm_kpm_ies.PfContainer_OCuCp{
		oCuCp := &e2sm_kpm_ies.OcucpPfContainer{
			GNbCuCpName: &e2sm_kpm_ies.GnbCuCpName{
				Value: "0xAB", //string -- PrintableString(SIZE(1..150,...))
			},
			CuCpResourceStatus: &e2sm_kpm_ies.OcucpPfContainer_CuCpResourceStatus001{
				NumberOfActiveUes: 1, //int32 -- INTEGER (1..65536, ...)
			},
		}
	//}

	xer, err := xerEncodeOCUCPPFContainer(oCuCp)
	assert.NilError(t, err)
	t.Logf("OCUCP-PF-Container XER\n%s", string(xer))
}
