// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"fmt"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmKpmIndicationHeader(timeStamp []byte) (*e2sm_kpm_v2_go.E2SmKpmIndicationHeader, error) {

	if len(timeStamp) != 4 {
		return nil, fmt.Errorf("TimeStamp should be 4 chars")
	}

	e2SmKpmPdu := e2sm_kpm_v2_go.E2SmKpmIndicationHeader{
		IndicationHeaderFormats: &e2sm_kpm_v2_go.IndicationHeaderFormats{
			E2SmKpmIndicationHeader: &e2sm_kpm_v2_go.IndicationHeaderFormats_IndicationHeaderFormat1{
				IndicationHeaderFormat1: &e2sm_kpm_v2_go.E2SmKpmIndicationHeaderFormat1{
					ColletStartTime: &e2sm_kpm_v2_go.TimeStamp{
						Value: timeStamp,
					},
				},
			},
		},
	}

	//if err := e2SmKpmPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2SmKpmPDU %s", err.Error())
	//}
	return &e2SmKpmPdu, nil
}

func CreateGlobalKpmnodeIDgNBID(bs *asn1.BitString, plmnID []byte) (*e2sm_kpm_v2_go.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 chars")
	}

	if bs.GetLen() < 22 || bs.GetLen() > 32 {
		return nil, fmt.Errorf("expecting GNbID length in range 22 to 32 bits, got %d", bs.GetLen())
	}

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_GNb{
			GNb: &e2sm_kpm_v2_go.GlobalKpmnodeGnbId{
				GlobalGNbId: &e2sm_kpm_v2_go.GlobalgNbId{
					GnbId: &e2sm_kpm_v2_go.GnbIdChoice{
						GnbIdChoice: &e2sm_kpm_v2_go.GnbIdChoice_GnbId{
							GnbId: bs,
						},
					},
					PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: plmnID,
					},
				},
			},
		},
	}, nil
}

func CreateGlobalKpmnodeIDenGNbID(bsValue []byte, bsLen uint32, plmnID []byte) (*e2sm_kpm_v2_go.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 chars")
	}

	if bsLen < 22 || bsLen > 32 {
		return nil, fmt.Errorf("expecting GNbID length in range 22 to 32 bits, got %d", bsLen)
	}

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_EnGNb{
			EnGNb: &e2sm_kpm_v2_go.GlobalKpmnodeEnGnbId{
				GlobalGNbId: &e2sm_kpm_v2_go.GlobalenGnbId{
					GNbId: &e2sm_kpm_v2_go.EngnbId{
						EngnbId: &e2sm_kpm_v2_go.EngnbId_GNbId{
							GNbId: &asn1.BitString{
								Value: bsValue,
								Len:   bsLen, // should be 22 to 32
							},
						},
					},
					PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: plmnID,
					},
				},
			},
		},
	}, nil
}

func CreateGlobalKpmnodeIDngENbID(enbID *e2sm_kpm_v2_go.EnbIdChoice, plmnID []byte, shortMacroEnbID *asn1.BitString,
	longMacroEnbID *asn1.BitString) (*e2sm_kpm_v2_go.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 chars")
	}

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_NgENb{
			NgENb: &e2sm_kpm_v2_go.GlobalKpmnodeNgEnbId{
				GlobalNgENbId: &e2sm_kpm_v2_go.GlobalngeNbId{
					EnbId: enbID,
					PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: plmnID,
					},
					ShortMacroENbId: shortMacroEnbID,
					LongMacroENbId:  longMacroEnbID,
				},
			},
		},
	}, nil
}

func CreateEnbIDchoiceMacro(bs *asn1.BitString) (*e2sm_kpm_v2_go.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2sm_kpm_v2_go.EnbIdChoice{
		EnbIdChoice: &e2sm_kpm_v2_go.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: bs,
		},
	}, nil
}

func CreateEnbIDchoiceShortMacro(bs *asn1.BitString) (*e2sm_kpm_v2_go.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x3f > 0 {
		return nil, fmt.Errorf("expected last 6 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2sm_kpm_v2_go.EnbIdChoice{
		EnbIdChoice: &e2sm_kpm_v2_go.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: bs,
		},
	}, nil
}

func CreateEnbIDchoiceLongMacro(bs *asn1.BitString) (*e2sm_kpm_v2_go.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x07 > 0 {
		return nil, fmt.Errorf("expected last 3 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2sm_kpm_v2_go.EnbIdChoice{
		EnbIdChoice: &e2sm_kpm_v2_go.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: bs,
		},
	}, nil
}

func CreateGlobalKpmnodeIDeNBID(enbID *e2sm_kpm_v2_go.EnbId, plmnID []byte) (*e2sm_kpm_v2_go.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 chars")
	}

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_ENb{
			ENb: &e2sm_kpm_v2_go.GlobalKpmnodeEnbId{
				GlobalENbId: &e2sm_kpm_v2_go.GlobalEnbId{
					ENbId: enbID,
					PLmnIdentity: &e2sm_kpm_v2_go.PlmnIdentity{
						Value: plmnID,
					},
				},
			},
		},
	}, nil
}

func CreateHomeEnbID(bs *asn1.BitString) (*e2sm_kpm_v2_go.EnbId, error) {

	if len(bs.GetValue()) != 4 {
		return nil, fmt.Errorf("expecting length to be exactly 4 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[3]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[3])
	}

	return &e2sm_kpm_v2_go.EnbId{
		EnbId: &e2sm_kpm_v2_go.EnbId_HomeENbId{
			HomeENbId: bs,
		},
	}, nil
}

func CreateMacroEnbID(bs *asn1.BitString) (*e2sm_kpm_v2_go.EnbId, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}

	return &e2sm_kpm_v2_go.EnbId{
		EnbId: &e2sm_kpm_v2_go.EnbId_MacroENbId{
			MacroENbId: bs,
		},
	}, nil
}
