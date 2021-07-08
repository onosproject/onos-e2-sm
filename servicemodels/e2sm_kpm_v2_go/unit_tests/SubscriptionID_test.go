// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerSubID = "00000000  40 30 38                                          |@08|"

func Test_perEncodingSubscriptionID(t *testing.T) {

	subID := &e2sm_kpm_v2_go.SubscriptionId{
		Value: 12345,
	}

	per, err := aper.Marshal(subID)
	assert.NilError(t, err)
	t.Logf("SubscriptionID PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.SubscriptionId{}
	err = aper.Unmarshal(per, &result)
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("SubscriptionID PER - decoded\n%v", result)
}

func Test_perSubscriptionIDCompareBytes(t *testing.T) {

	subID := &e2sm_kpm_v2_go.SubscriptionId{
		Value: 12345,
	}

	per, err := aper.Marshal(subID)
	assert.NilError(t, err)
	t.Logf("SubscriptionID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerSubID)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
