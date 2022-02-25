// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmKpmIndicationHeader(timeStamp []byte) (*e2smkpmv2.E2SmKpmIndicationHeader, error) {

	if len(timeStamp) != 4 {
		return nil, errors.NewInvalid("TimeStamp should be 4 chars")
	}

	e2SmKpmPdu := e2smkpmv2.E2SmKpmIndicationHeader{
		IndicationHeaderFormats: &e2smkpmv2.IndicationHeaderFormats{
			E2SmKpmIndicationHeader: &e2smkpmv2.IndicationHeaderFormats_IndicationHeaderFormat1{
				IndicationHeaderFormat1: &e2smkpmv2.E2SmKpmIndicationHeaderFormat1{
					ColletStartTime: &e2smkpmv2.TimeStamp{
						Value: timeStamp,
					},
				},
			},
		},
	}

	if err := e2SmKpmPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmKpmIndicationHeader(): error validating E2SmKpmPDU %s", err.Error())
	}
	return &e2SmKpmPdu, nil
}

func CreateGlobalKpmnodeIDgNBID(bs *asn1.BitString, plmnID []byte) (*e2smkpmv2.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	if bs.GetLen() < 22 || bs.GetLen() > 32 {
		return nil, errors.NewInvalid("expecting GNbID length in range 22 to 32 bits, got %d", bs.GetLen())
	}

	return &e2smkpmv2.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2smkpmv2.GlobalKpmnodeId_GNb{
			GNb: &e2smkpmv2.GlobalKpmnodeGnbId{
				GlobalGNbId: &e2smkpmv2.GlobalgNbId{
					GnbId: &e2smkpmv2.GnbIdChoice{
						GnbIdChoice: &e2smkpmv2.GnbIdChoice_GnbId{
							GnbId: bs,
						},
					},
					PlmnId: &e2smkpmv2.PlmnIdentity{
						Value: plmnID,
					},
				},
			},
		},
	}, nil
}

func CreateGlobalKpmnodeIDenGNbID(bsValue []byte, bsLen uint32, plmnID []byte) (*e2smkpmv2.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	if bsLen < 22 || bsLen > 32 {
		return nil, errors.NewInvalid("expecting GNbID length in range 22 to 32 bits, got %d", bsLen)
	}

	return &e2smkpmv2.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2smkpmv2.GlobalKpmnodeId_EnGNb{
			EnGNb: &e2smkpmv2.GlobalKpmnodeEnGnbId{
				GlobalGNbId: &e2smkpmv2.GlobalenGnbId{
					GNbId: &e2smkpmv2.EngnbId{
						EngnbId: &e2smkpmv2.EngnbId_GNbId{
							GNbId: &asn1.BitString{
								Value: bsValue,
								Len:   bsLen, // should be 22 to 32
							},
						},
					},
					PLmnIdentity: &e2smkpmv2.PlmnIdentity{
						Value: plmnID,
					},
				},
			},
		},
	}, nil
}

func CreateGlobalKpmnodeIDngENbID(enbID *e2smkpmv2.EnbIdChoice, plmnID []byte, shortMacroEnbID *asn1.BitString,
	longMacroEnbID *asn1.BitString) (*e2smkpmv2.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	return &e2smkpmv2.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2smkpmv2.GlobalKpmnodeId_NgENb{
			NgENb: &e2smkpmv2.GlobalKpmnodeNgEnbId{
				GlobalNgENbId: &e2smkpmv2.GlobalngeNbId{
					EnbId: enbID,
					PlmnId: &e2smkpmv2.PlmnIdentity{
						Value: plmnID,
					},
					ShortMacroENbId: shortMacroEnbID,
					LongMacroENbId:  longMacroEnbID,
				},
			},
		},
	}, nil
}

func CreateEnbIDchoiceMacro(bs *asn1.BitString) (*e2smkpmv2.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, errors.NewInvalid("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2smkpmv2.EnbIdChoice{
		EnbIdChoice: &e2smkpmv2.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: bs,
		},
	}, nil
}

func CreateEnbIDchoiceShortMacro(bs *asn1.BitString) (*e2smkpmv2.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, errors.NewInvalid("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x3f > 0 {
		return nil, errors.NewInvalid("expected last 6 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2smkpmv2.EnbIdChoice{
		EnbIdChoice: &e2smkpmv2.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: bs,
		},
	}, nil
}

func CreateEnbIDchoiceLongMacro(bs *asn1.BitString) (*e2smkpmv2.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, errors.NewInvalid("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x07 > 0 {
		return nil, errors.NewInvalid("expected last 3 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2smkpmv2.EnbIdChoice{
		EnbIdChoice: &e2smkpmv2.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: bs,
		},
	}, nil
}

func CreateGlobalKpmnodeIDeNBID(enbID *e2smkpmv2.EnbId, plmnID []byte) (*e2smkpmv2.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	return &e2smkpmv2.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2smkpmv2.GlobalKpmnodeId_ENb{
			ENb: &e2smkpmv2.GlobalKpmnodeEnbId{
				GlobalENbId: &e2smkpmv2.GlobalEnbId{
					ENbId: enbID,
					PLmnIdentity: &e2smkpmv2.PlmnIdentity{
						Value: plmnID,
					},
				},
			},
		},
	}, nil
}

func CreateHomeEnbID(bs *asn1.BitString) (*e2smkpmv2.EnbId, error) {

	if len(bs.GetValue()) != 4 {
		return nil, errors.NewInvalid("expecting length to be exactly 4 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[3]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[3])
	}

	return &e2smkpmv2.EnbId{
		EnbId: &e2smkpmv2.EnbId_HomeENbId{
			HomeENbId: bs,
		},
	}, nil
}

func CreateMacroEnbID(bs *asn1.BitString) (*e2smkpmv2.EnbId, error) {

	if len(bs.GetValue()) != 3 {
		return nil, errors.NewInvalid("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}

	return &e2smkpmv2.EnbId{
		EnbId: &e2smkpmv2.EnbId_MacroENbId{
			MacroENbId: bs,
		},
	}, nil
}
