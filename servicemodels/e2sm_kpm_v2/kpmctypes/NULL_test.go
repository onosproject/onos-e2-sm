// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//func createNrcgi() *e2sm_kpm_v2.Nrcgi {
//	return &e2sm_kpm_v2.Nrcgi{
//		PLmnIdentity: &e2sm_kpm_v2.PlmnIdentity{
//			Value: []byte{21, 22, 23},
//		},
//		NRcellIdentity: &e2sm_kpm_v2.NrcellIdentity{
//			Value: &e2sm_kpm_v2.BitString{
//				Value: 0x9bcd4,
//				Len:   22,
//			},
//		},
//	}
//}
//
//func Test_xerEncodeNrcgi(t *testing.T) {
//
//	nrcgi := createNrcgi()
//
//	xer, err := xerEncodeNrcgi(nrcgi)
//	assert.NilError(t, err)
//	assert.Equal(t, 4, len(xer))
//	t.Logf("NRCGI XER\n%s", string(xer))
//}
//
//func Test_xerDecodeNrcgi(t *testing.T) {
//
//	nrcgi := createNrcgi()
//
//	xer, err := xerEncodeNrcgi(nrcgi)
//	assert.NilError(t, err)
//	assert.Equal(t, 4, len(xer))
//	t.Logf("NRCGI XER\n%s", string(xer))
//
//	result, err := xerDecodeNrcgi(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//}
//
//func Test_perEncodeNrcgi(t *testing.T) {
//
//	nrcgi := createNrcgi()
//
//	per, err := perEncodeNrcgi(nrcgi)
//	assert.NilError(t, err)
//	assert.Equal(t, 4, len(per))
//	t.Logf("NRCGI PER\n%s", string(per))
//}
//
//func Test_perDecodeNrcgi(t *testing.T) {
//
//	nrcgi := createNrcgi()
//
//	per, err := perEncodeNrcgi(nrcgi)
//	assert.NilError(t, err)
//	assert.Equal(t, 4, len(per))
//	t.Logf("NRCGI PER\n%s", string(per))
//
//	result, err := perDecodeNrcgi(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//}
