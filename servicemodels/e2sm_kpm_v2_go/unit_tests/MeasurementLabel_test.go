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

var refPerMeasLabel = "00000000  7f ff c0 01 02 03 4f c0  01 02 03 00 17 68 18 00  |......@@.....h..|\n" +
	"00000010  1e 00 01 70 00 00 18 00  00 00 00 00 7a 00 01 c7  |...p........z...|\n" +
	"00000020  00 03 14 20                                       |... |"
var refPerMeasLabelXcldSomeOptnl = "00000000  66 7d 40 01 02 03 40 40  01 02 03 00 18 00 1e 00  |f}@...@@........|\n" +
	"00000010  00 18 00 00 00 00 00 7a  00 03 14 40              |.......z...@|"

func createMeasurementLabel() *e2sm_kpm_v2_go.MeasurementLabel {

	var br int32 = 25
	var lmm int32 = 1
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END

	return &e2sm_kpm_v2_go.MeasurementLabel{
		PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x01, 0x02, 0x03},
		},
		SliceId: &e2sm_kpm_v2_go.Snssai{
			SD:  []byte{0x01, 0x02, 0x03},
			SSt: []byte{0x3f},
		},
		FiveQi: &e2sm_kpm_v2_go.FiveQi{
			Value: 23,
		},
		QFi: &e2sm_kpm_v2_go.Qfi{
			Value: 52,
		},
		QCi: &e2sm_kpm_v2_go.Qci{
			Value: 24,
		},
		QCimax: &e2sm_kpm_v2_go.Qci{
			Value: 30,
		},
		QCimin: &e2sm_kpm_v2_go.Qci{
			Value: 1,
		},
		ARpmax: &e2sm_kpm_v2_go.Arp{
			Value: 15,
		},
		ARpmin: &e2sm_kpm_v2_go.Arp{
			Value: 1,
		},
		BitrateRange:     &br,
		LayerMuMimo:      &lmm,
		SUm:              &sum,
		DistBinX:         &dbx,
		DistBinY:         &dby,
		DistBinZ:         &dbz,
		PreLabelOverride: &plo,
		StartEndInd:      &seind,
	}
}

func createMeasurementLabelXcldSomeOptnl() *e2sm_kpm_v2_go.MeasurementLabel {

	var br int32 = 25
	var lmm int32 = 1
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	var dbx int32 = 123
	var dbz int32 = 789
	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END

	return &e2sm_kpm_v2_go.MeasurementLabel{
		PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
			Value: []byte{0x01, 0x02, 0x03},
		},
		SliceId: &e2sm_kpm_v2_go.Snssai{
			SD:  []byte{0x01, 0x02, 0x03},
			SSt: []byte{0x01},
		},
		QCi: &e2sm_kpm_v2_go.Qci{
			Value: 24,
		},
		QCimax: &e2sm_kpm_v2_go.Qci{
			Value: 30,
		},
		ARpmin: &e2sm_kpm_v2_go.Arp{
			Value: 1,
		},
		BitrateRange: &br,
		LayerMuMimo:  &lmm,
		SUm:          &sum,
		DistBinX:     &dbx,
		DistBinZ:     &dbz,
		StartEndInd:  &seind,
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
	t.Logf("MeasurementLabel PER - decoded\n%v", &result)
	assert.DeepEqual(t, ml.GetPlmnId().GetValue(), result.GetPlmnId().GetValue())
	assert.DeepEqual(t, ml.GetSliceId().GetSD(), result.GetSliceId().GetSD())
	assert.DeepEqual(t, ml.GetSliceId().GetSSt(), result.GetSliceId().GetSSt())
	assert.Equal(t, ml.GetFiveQi().GetValue(), result.GetFiveQi().GetValue())
	assert.Equal(t, ml.GetQFi().GetValue(), result.GetQFi().GetValue())
	assert.Equal(t, ml.GetQCi().GetValue(), result.GetQCi().GetValue())
	assert.Equal(t, ml.GetQCimax().GetValue(), result.GetQCimax().GetValue())
	assert.Equal(t, ml.GetQCimin().GetValue(), result.GetQCimin().GetValue())
	assert.Equal(t, ml.GetARpmax().GetValue(), result.GetARpmax().GetValue())
	assert.Equal(t, ml.GetARpmin().GetValue(), result.GetARpmin().GetValue())
	assert.Equal(t, ml.GetBitrateRange(), result.GetBitrateRange())
	assert.Equal(t, ml.GetLayerMuMimo(), result.GetLayerMuMimo())
	assert.Equal(t, ml.GetSUm().Number(), result.GetSUm().Number())
	assert.Equal(t, ml.GetDistBinX(), result.GetDistBinX())
	assert.Equal(t, ml.GetDistBinY(), result.GetDistBinY())
	assert.Equal(t, ml.GetDistBinZ(), result.GetDistBinZ())
	assert.Equal(t, ml.GetPreLabelOverride().Number(), result.GetPreLabelOverride().Number())
	assert.Equal(t, ml.GetStartEndInd().Number(), result.GetStartEndInd().Number())
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

func Test_perEncodeMeasurementLabelXcldSomeOptnl(t *testing.T) {

	ml := createMeasurementLabelXcldSomeOptnl()
	t.Logf("MeasurementLabel message is\n%v", ml)

	per, err := aper.MarshalWithParams(*ml, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementLabel PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementLabel{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementLabel PER - decoded\n%v", &result)
	assert.DeepEqual(t, ml.GetPlmnId().GetValue(), result.GetPlmnId().GetValue())
	assert.DeepEqual(t, ml.GetSliceId().GetSD(), result.GetSliceId().GetSD())
	assert.DeepEqual(t, ml.GetSliceId().GetSSt(), result.GetSliceId().GetSSt())
	assert.Equal(t, ml.GetQCi().GetValue(), result.GetQCi().GetValue())
	assert.Equal(t, ml.GetQCimax().GetValue(), result.GetQCimax().GetValue())
	assert.Equal(t, ml.GetQCimin().GetValue(), result.GetQCimin().GetValue())
	assert.Equal(t, ml.GetARpmin().GetValue(), result.GetARpmin().GetValue())
	assert.Equal(t, ml.GetBitrateRange(), result.GetBitrateRange())
	assert.Equal(t, ml.GetLayerMuMimo(), result.GetLayerMuMimo())
	assert.Equal(t, ml.GetSUm().Number(), result.GetSUm().Number())
	assert.Equal(t, ml.GetDistBinX(), result.GetDistBinX())
	assert.Equal(t, ml.GetDistBinZ(), result.GetDistBinZ())
	assert.Equal(t, ml.GetStartEndInd().Number(), result.GetStartEndInd().Number())
}

func Test_perMeasurementLabelXcldSomeOptnlCompareBytes(t *testing.T) {

	ml := createMeasurementLabelXcldSomeOptnl()
	t.Logf("MeasurementLabel message is\n%v", ml)

	per, err := aper.MarshalWithParams(*ml, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementLabel PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasLabelXcldSomeOptnl)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
