// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdubuilder

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmIndicationHeader(t *testing.T) {
	bs := e2sm_kpm_v2.BitString{
		Value: 0x9bcd4,
		Len:   22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	var fileFormatVersion int32 = 2
	var senderName int32 = 123
	var vendorName int32 = 54

	globalKpmNodeID, err := CreateGlobalKpmnodeID_gNBID(bs, plmnID, gnbCuUpID, gnbDuID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmKpmIndicationHeader(timeStamp, fileFormatVersion, senderName, vendorName, *globalKpmNodeID)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
