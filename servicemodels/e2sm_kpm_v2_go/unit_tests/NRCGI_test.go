// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerNrCGI = "00000000  00 21 22 23 d4 bc 09 00  00                       |.!\"#.....|"

func createNrcgi() *e2sm_kpm_v2_go.Nrcgi {
	return &e2sm_kpm_v2_go.Nrcgi{
		PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x21, 0x22, 0x23},
		},
		NRcellIdentity: &e2sm_kpm_v2_go.NrcellIdentity{
			Value: &asn1.BitString{
				Value: []byte{0xd4, 0xbc, 0x09, 0x00},
				Len:   36,
			},
		},
	}
}

func Test_perEncodingNrCGI(t *testing.T) {

	nrCgi := createNrcgi()

	per, err := aper.MarshalWithParams(*nrCgi, "valueExt")
	assert.NilError(t, err)
	t.Logf("NrCGI PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.Nrcgi{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("NrCGI PER - decoded\n%v", result)
}

func Test_perNrCGICompareBytes(t *testing.T) {

	nrCgi := createNrcgi()

	per, err := aper.MarshalWithParams(*nrCgi, "valueExt")
	assert.NilError(t, err)
	t.Logf("NrCGI PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerNrCGI)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
