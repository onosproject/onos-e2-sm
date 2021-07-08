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

//ToDo - find out why PER bytes differ.. Not correct treatment of ENUMERATORs?
var refPerMeasLabel = "00000000  7f ff c0 01 02 03 40 40  01 02 03 00 17 68 18 00  |......@@.....h..|\n" +
	"00000010  1e 00 01 70 00 00 18 00  00 00 00 00 7a 00 01 c7  |...p........z...|\n" +
	"00000020  00 03 14 20                                       |... |"

func createMeasurementLabel() *e2sm_kpm_v2_go.MeasurementLabel {

	//var br int32 = 25
	//var lmm int32 = 1
	//sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	//var dbx int32 = 123
	//var dby int32 = 456
	//var dbz int32 = 789
	//plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	//seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	return &e2sm_kpm_v2_go.MeasurementLabel{
		PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x1, 0x2, 0x3},
		},
		SliceId: &e2sm_kpm_v2_go.Snssai{
			SD:  []byte{0x01, 0x02, 0x03},
			SSt: []byte{0x01},
		},
		//FiveQi: &e2sm_kpm_v2_go.FiveQi{
		//	Value: 23,
		//},
		//QFi: &e2sm_kpm_v2_go.Qfi{
		//	Value: 52,
		//},
		//QCi: &e2sm_kpm_v2_go.Qci{
		//	Value: 24,
		//},
		//QCimax: &e2sm_kpm_v2_go.Qci{
		//	Value: 30,
		//},
		//QCimin: &e2sm_kpm_v2_go.Qci{
		//	Value: 1,
		//},
		//ARpmax: &e2sm_kpm_v2_go.Arp{
		//	Value: 15,
		//},
		//ARpmin: &e2sm_kpm_v2_go.Arp{
		//	Value: 1,
		//},
		//BitrateRange:     &br,
		//LayerMuMimo:      &lmm,
		//SUm:              &sum,
		//DistBinX:         &dbx,
		//DistBinY:         &dby,
		//DistBinZ:         &dbz,
		//PreLabelOverride: &plo,
		//StartEndInd:      &seind,
	}
}

func Test_perEncodeMeasurementLabel(t *testing.T) {

	ml := createMeasurementLabel()
	t.Logf("MeasurementLabel message is\n%v", ml)

	per, err := aper.MarshalWithParams(*ml, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementLabel PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementLabel{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementLabel PER - decoded\n%v", result)
}

func Test_perMeasurementLabelCompareBytes(t *testing.T) {

	ml := createMeasurementLabel()
	t.Logf("MeasurementLabel message is\n%v", ml)

	per, err := aper.MarshalWithParams(*ml, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementLabel PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasLabel)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
