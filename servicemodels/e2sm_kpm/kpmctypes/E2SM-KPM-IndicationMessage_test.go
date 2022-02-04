// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeE2SmKpmIndicationMessage(t *testing.T) {

	var testGnbCuCpName = "0xAB"
	var testNumberOfActiveUes int32 = 1
	value := []byte{0x22, 0x21}

	containerOcuCp1 := &e2sm_kpm_ies.PmContainersList{
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
		TheRancontainer: &e2sm_kpm_ies.RanContainer{
			Value: value,
		},
	}

	e2SmIindicationMsg := &e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2sm_kpm_ies.E2SmKpmIndicationMessageFormat1{
			PmContainers: make([]*e2sm_kpm_ies.PmContainersList, 0),
		},
	}
	e2SmIindicationMsg.IndicationMessageFormat1.PmContainers = append(e2SmIindicationMsg.IndicationMessageFormat1.PmContainers, containerOcuCp1)

	e2smKpmPdu := &e2sm_kpm_ies.E2SmKpmIndicationMessage{
		E2SmKpmIndicationMessage: e2SmIindicationMsg,
	}

	xer, err := XerEncodeE2SmKpmIndicationMessage(e2smKpmPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-KPM-IndicationMessage XER\n%s", string(xer))
}
