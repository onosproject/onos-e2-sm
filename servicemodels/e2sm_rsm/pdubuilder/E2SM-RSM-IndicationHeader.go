// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdubuilder

import (
	"fmt"
	e2sm_rsm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies"
	e2sm_v2_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmRsmIndicationHeaderFormat1(cgi *e2sm_v2_ies.Cgi) *e2sm_rsm_ies.E2SmRsmIndicationHeader {

	return &e2sm_rsm_ies.E2SmRsmIndicationHeader{
		E2SmRsmIndicationHeader: &e2sm_rsm_ies.E2SmRsmIndicationHeader_IndicationHeaderFormat1{
			IndicationHeaderFormat1: &e2sm_rsm_ies.E2SmRsmIndicationHeaderFormat1{
				Cgi: cgi,
			},
		},
	}
}

func CreateNrCGI(plmnID []byte, nrCellID *asn1.BitString) (*e2sm_v2_ies.Cgi, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("length of Plmn ID is expected to be exactly 3 bytes")
	}

	if len(nrCellID.GetValue()) != 5 {
		return nil, fmt.Errorf("NRcellIdentity is expected to be 5 bytes long")
	}

	if nrCellID.GetValue()[4]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", nrCellID.GetValue()[3])
	}

	return &e2sm_v2_ies.Cgi{
		Cgi: &e2sm_v2_ies.Cgi_NRCgi{
			NRCgi: &e2sm_v2_ies.NrCgi{
				PLmnidentity: &e2sm_v2_ies.PlmnIdentity{
					Value: plmnID,
				},
				NRcellIdentity: &e2sm_v2_ies.NrcellIdentity{
					Value: nrCellID,
				},
			},
		},
	}, nil
}

func CreateEutraCGI(plmnID []byte, eutraCellID *asn1.BitString) (*e2sm_v2_ies.Cgi, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("length of Plmn ID is expected to be exactly 3 bytes")
	}

	if len(eutraCellID.GetValue()) != 4 {
		return nil, fmt.Errorf("EutraCellIdentity is expected to be 4 bytes long")
	}

	if eutraCellID.GetValue()[3]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", eutraCellID.GetValue()[3])
	}

	return &e2sm_v2_ies.Cgi{
		Cgi: &e2sm_v2_ies.Cgi_EUtraCgi{
			EUtraCgi: &e2sm_v2_ies.EutraCgi{
				PLmnidentity: &e2sm_v2_ies.PlmnIdentity{
					Value: plmnID,
				},
				EUtracellIdentity: &e2sm_v2_ies.EutracellIdentity{
					Value: eutraCellID,
				},
			},
		},
	}, nil
}
