// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import "fmt"

type Qci int32
type FiveQi int32
type SSt [1]byte
type SD [3]byte
type NRcellIdentityValue uint64
type NRcellIdentityLen uint32

func SdFromSlice(sdBytes []byte) (SD, error) {
	if len(sdBytes) != 3 {
		return [3]byte{}, fmt.Errorf("SD must be 3 bytes long")
	}
	sd := [3]byte{sdBytes[0], sdBytes[1], sdBytes[2]}
	fmt.Printf("SD -- %v\n", sd)
	return sd, nil
}

func SstFromSlice(sstBytes []byte) (SSt, error) {
	if len(sstBytes) != 1 {
		return [1]byte{}, fmt.Errorf("SST must be 1 bytes long")
	}
	sst := [1]byte{sstBytes[0]}
	fmt.Printf("SST -- %v\n", sst)
	return sst, nil
}

func NewNrCellPlmnId(nrPlmnIDSlice []byte) (*NRcgi, error) {
	plmnID, err := PlmnIDFromSlice(nrPlmnIDSlice)
	if err != nil {
		return nil, err
	}

	return &NRcgi{
		PLmnIdentity: plmnID,
	}, nil
}

type NRcellIdentity struct {
	Value NRcellIdentityValue
	Len NRcellIdentityLen
}

type SliceId struct {
	SSt SSt
	SD SD
}

type NRcgi struct {
	PLmnIdentity PlmnID
	NRcellIdentity NRcellIdentity
}

type IndicationHeader1 struct {
	NRcgi NRcgi
	PLmnIdentity PlmnID
	SliceId SliceId
	FiveQi FiveQi
	Qci Qci
}

