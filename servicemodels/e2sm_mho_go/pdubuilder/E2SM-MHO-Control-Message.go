// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smmhov2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2smv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMhoControlMessage(servingCgi *e2smv2.Cgi, uedID *e2smv2.Ueid, targetCgi *e2smv2.Cgi) (*e2smmhov2.E2SmMhoControlMessage, error) {

	e2smMhoMsgFormat1 := e2smmhov2.E2SmMhoControlMessageFormat1{
		ServingCgi: servingCgi,
		UedId:      uedID,
		TargetCgi:  targetCgi,
	}

	e2smMhoPdu := e2smmhov2.E2SmMhoControlMessage{
		E2SmMhoControlMessage: &e2smmhov2.E2SmMhoControlMessage_ControlMessageFormat1{
			ControlMessageFormat1: &e2smMhoMsgFormat1,
		},
	}

	if err := e2smMhoPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMhoControlMessage(): error validating E2SmPDU %s", err.Error())
	}
	return &e2smMhoPdu, nil
}

func CreateUeIDGNb(amf int64, plmnID []byte, amfRegionID []byte, amfSetID []byte, amfPointer []byte) (*e2smv2.Ueid, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("CreateUeIDGNb() PlmnID should contain only 3 bytes, got %v", len(plmnID))
	}
	if len(amfRegionID) != 1 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfRegionID should contain only 1 byte, got %v", len(amfRegionID))
	}
	if len(amfSetID) != 2 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfSetID should contain only 2 bytes, got %v", len(amfSetID))
	}
	if amfSetID[1]&0x3F > 0 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfSetID should contain only 10 bits, i.e., last 6 bits of last byte should be trailing zeros, got %v", amfSetID[1])
	}
	if len(amfPointer) != 1 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfPointer should contain only 1 byte, got %v", len(amfPointer))
	}
	if amfPointer[0]&0x03 > 0 {
		return nil, errors.NewInvalid("CreateUeIDGNb() AMfSetID should contain only 6 bits, i.e., last 2 bits should be trailing zeros, got %v", amfPointer)
	}

	ueID := &e2smv2.Ueid{
		Ueid: &e2smv2.Ueid_GNbUeid{
			GNbUeid: &e2smv2.UeidGnb{
				AmfUeNgapId: &e2smv2.AmfUeNgapId{
					Value: amf,
				},
				Guami: &e2smv2.Guami{
					PLmnidentity: &e2smv2.PlmnIdentity{
						Value: plmnID,
					},
					AMfregionId: &e2smv2.AmfregionId{
						Value: &asn1.BitString{
							Value: amfRegionID,
							Len:   8,
						},
					},
					AMfsetId: &e2smv2.AmfsetId{
						Value: &asn1.BitString{
							Value: amfSetID,
							Len:   10,
						},
					},
					AMfpointer: &e2smv2.Amfpointer{
						Value: &asn1.BitString{
							Value: amfPointer,
							Len:   6,
						},
					},
				},
			},
		},
	}

	if err := ueID.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateUeIDGNb() validation of UeID failed with %v", err)
	}

	return ueID, nil
}
