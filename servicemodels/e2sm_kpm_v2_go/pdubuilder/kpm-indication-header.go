// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SmKpmIndicationHeader(timeStamp []byte, fileFormatVersion *string, senderName *string, senderType *string,
	vendorName *string, globalKpmNodeID *e2sm_kpm_v2_go.GlobalKpmnodeId) (*e2sm_kpm_v2_go.E2SmKpmIndicationHeader, error) {

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

	// optional instance
	if fileFormatVersion != nil {
		e2SmKpmPdu.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().FileFormatversion = fileFormatVersion
	}
	// optional instance
	if senderName != nil {
		e2SmKpmPdu.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().SenderName = senderName
	}
	// optional instance
	if senderType != nil {
		e2SmKpmPdu.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().SenderType = senderType
	}
	// optional instance
	if vendorName != nil {
		e2SmKpmPdu.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().VendorName = vendorName
	}
	// optional instance
	if globalKpmNodeID != nil {
		e2SmKpmPdu.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().KpmNodeId = globalKpmNodeID
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

func CreateGlobalKpmnodeIDngENbID(bs *asn1.BitString, plmnID []byte, shortMacroEnbID *asn1.BitString,
	longMacroEnbID *asn1.BitString) (*e2sm_kpm_v2_go.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 chars")
	}

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_NgENb{
			NgENb: &e2sm_kpm_v2_go.GlobalKpmnodeNgEnbId{
				GlobalNgENbId: &e2sm_kpm_v2_go.GlobalngeNbId{
					EnbId: &e2sm_kpm_v2_go.EnbIdChoice{
						EnbIdChoice: &e2sm_kpm_v2_go.EnbIdChoice_EnbIdMacro{
							EnbIdMacro: bs,
						},
					},
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

func CreateGlobalKpmnodeIDeNBID(bs *asn1.BitString, plmnID []byte) (*e2sm_kpm_v2_go.GlobalKpmnodeId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 chars")
	}

	return &e2sm_kpm_v2_go.GlobalKpmnodeId{
		GlobalKpmnodeId: &e2sm_kpm_v2_go.GlobalKpmnodeId_ENb{
			ENb: &e2sm_kpm_v2_go.GlobalKpmnodeEnbId{
				GlobalENbId: &e2sm_kpm_v2_go.GlobalEnbId{
					ENbId: &e2sm_kpm_v2_go.EnbId{
						EnbId: &e2sm_kpm_v2_go.EnbId_HomeENbId{
							HomeENbId: bs,
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
