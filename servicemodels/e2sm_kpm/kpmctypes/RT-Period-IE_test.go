// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
)

func createRtPeriodIe() e2sm_kpm_ies.RtPeriodIe {

	return e2sm_kpm_ies.RtPeriodIe_RT_PERIOD_IE_MS10

}

//func Test_xerEncodeRtPeriodIe(t *testing.T) {
//
//	rtPeriodIe := createRtPeriodIe()
//
//	xer, err := xerEncodeRtPeriodIe(rtPeriodIe)
//	assert.NilError(t, err)
//	assert.Equal(t, 308, len(xer))
//	t.Logf("RT-Period-IE XER\n%s", string(xer))
//}

//func Test_xerDecodeRtPeriodIe(t *testing.T) {
//
//	rtPeriodIe := createRtPeriodIe()
//
//	xer, err := xerEncodeRtPeriodIe(rtPeriodIe)
//	assert.NilError(t, err)
//	assert.Equal(t, 308, len(xer))
//	t.Logf("RT-Period-IE XER\n%s", string(xer))
//
//	result, err := xerDecodeRtPeriodIe(xer)
//	assert.NilError(t, err)
//	assert.Equal(t, rtPeriodIe, result, "Encoded and decoded RtPeriodIe values are not the same")
//}

//func Test_perEncodeRtPeriodIe(t *testing.T) {
//
//	rtPeriodIe := createRtPeriodIe()
//
//	per, err := perEncodeRtPeriodIe(rtPeriodIe)
//	assert.NilError(t, err)
//	assert.Equal(t, 18, len(per))
//	t.Logf("RT-Period-IE PER\n%s", string(per))
//}

//func Test_perDecodeRtPeriodIe(t *testing.T) {
//
//	rtPeriodIe := createRtPeriodIe()
//
//	per, err := perEncodeRtPeriodIe(rtPeriodIe)
//	assert.NilError(t, err)
//	assert.Equal(t, 18, len(per))
//	t.Logf("RT-Period-IE PER\n%s", string(per))
//
//	result, err := perDecodeRtPeriodIe(per)
//	assert.NilError(t, err)
//	assert.Equal(t, rtPeriodIe, result, "Encoded and decoded RtPeriodIe values are not the same")
//}
