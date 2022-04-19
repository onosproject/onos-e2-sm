// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2smcommonies

func (m *RanfunctionName) SetRanFunctionInstance(rfi int32) *RanfunctionName {
	m.RanFunctionInstance = &rfi
	return m
}
