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

var refPerMII = "00000000  50 00 14 01 3f ff e0 01  02 03 40 40 01 02 03 00  |P...?.....@@....|\n" +
	"00000010  17 7c 18 00 1e 00 01 70  00 00 18 00 00 00 00 00  |.|.....p........|\n" +
	"00000020  7a 00 01 c7 00 03 14 00                           |z.......|"

func createMeasurementInfoItem() (*e2sm_kpm_v2_go.MeasurementInfoItem, error) {

	labelInfoList := &e2sm_kpm_v2_go.LabelInfoList{
		Value: make([]*e2sm_kpm_v2_go.LabelInfoItem, 0),
	}

	var br int32 = 25
	var lmm int32 = 1
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_START

	labelInfoItem := &e2sm_kpm_v2_go.LabelInfoItem{
		MeasLabel: &e2sm_kpm_v2_go.MeasurementLabel{
			PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
				Value: []byte{0x1, 0x2, 0x3},
			},
			SliceId: &e2sm_kpm_v2_go.Snssai{
				SD:  []byte{0x01, 0x02, 0x03},
				SSt: []byte{0x01},
			},
			FiveQi: &e2sm_kpm_v2_go.FiveQi{
				Value: 23,
			},
			QFi: &e2sm_kpm_v2_go.Qfi{
				Value: 62,
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
		},
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	res := &e2sm_kpm_v2_go.MeasurementInfoItem{
		MeasType: &e2sm_kpm_v2_go.MeasurementType{
			MeasurementType: &e2sm_kpm_v2_go.MeasurementType_MeasId{
				MeasId: &e2sm_kpm_v2_go.MeasurementTypeId{
					Value: 21,
				},
			},
		},
		LabelInfoList: labelInfoList,
	}
	//if err := res.Validate(); err != nil {
	//	return nil, err
	//}

	return res, nil
}

func Test_perEncodeMeasurementInfoItem(t *testing.T) {

	mii, err := createMeasurementInfoItem()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-Item PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementInfoItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	assert.Assert(t, &result != nil)
	t.Logf("MeasurementInfo-Item PER - decoded\n%v", result)
}

func Test_perMeasurementInfoItemCompareBytes(t *testing.T) {

	mii, err := createMeasurementInfoItem()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-Item PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMII)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
