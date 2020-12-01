// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/types"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
)

func DecodeE2SmKpmIndicationHeader(e2smKpmPdu *e2sm_kpm_ies.E2SmKpmIndicationHeader) (*types.GlobalKpmNodeIdentity, *types.IndicationHeader1, error) {
	var nodeIdentity *types.GlobalKpmNodeIdentity
	var indicationHeader *types.IndicationHeader1
	var err error

	if err := e2smKpmPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2SmKpmPdu %s", err.Error())
	}

	fmt.Printf("Hooray?\n")
	identifierIe := e2smKpmPdu.GetIndicationHeaderFormat1().GetIdGlobalKpmnodeId()
	if identifierIe == nil {
		return nil, nil, fmt.Errorf("error E2SmKpmPdu does not have Id-GlobalKpmnode-ID")
	}
	fmt.Printf("Hooray?\n")
	switch kpmNodeId := identifierIe.GetGlobalKpmnodeId().(type) {
	case *e2sm_kpm_ies.GlobalKpmnodeId_GNb:
		nodeIdentity, err = types.NewKpmNodeIdentity(kpmNodeId.GNb.GetGlobalGNbId().GetPlmnId().GetValue())
		if err != nil {
			return nil, nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.KpmNodeTypeGnbId
		choice, ok := kpmNodeId.GNb.GetGlobalGNbId().GetGnbId().GetGnbIdChoice().(*e2sm_kpm_ies.GnbIdChoice_GnbId)
		if !ok {
			return nil, nil, fmt.Errorf("expected dNBId")
		}
		nodeIdentity.NodeIdentifier = make([]byte, 8)
		binary.LittleEndian.PutUint64(nodeIdentity.NodeIdentifier, choice.GnbId.GetValue())
		// TODO: investigate GNB-CU-UP-ID and GNB-DU-ID

	case *e2sm_kpm_ies.GlobalKpmnodeId_EnGNb:
			nodeIdentity, err:= types.NewKpmNodeIdentity(kpmNodeId.EnGNb.GetGlobalGNbId().GetPLmnIdentity().GetValue())
			if err != nil{
				return nil, nil, fmt.Errorf("error extracting node identifier")
			}
			nodeIdentity.NodeType = types.KpmNodeTypeEnGnbId
			return nil, nil, fmt.Errorf("getting identifier of EnGnb not yet implemented")

	case *e2sm_kpm_ies.GlobalKpmnodeId_ENb:
		nodeIdentity, err:= types.NewKpmNodeIdentity(kpmNodeId.ENb.GetGlobalENbId().GetPLmnIdentity().GetValue())
		if err != nil{
			return nil, nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.KpmNodeTypeEnbId
		return nil, nil, fmt.Errorf("getting identifier of ENb not yet implemented")

	case *e2sm_kpm_ies.GlobalKpmnodeId_NgENb:
		nodeIdentity, err:= types.NewKpmNodeIdentity(kpmNodeId.NgENb.GetGlobalNgENbId().GetPlmnId().GetValue())
		if err != nil{
			return nil, nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.KpmNodeTypeNgEnbId
		return nil, nil, fmt.Errorf("getting identifier of NgENb not yet implemented")
	}
	fmt.Printf("Hooray?\n")

	//ToDo: fix Go panic error..
	plmnIdNrCgi, err := types.PlmnIDFromSlice(e2smKpmPdu.GetIndicationHeaderFormat1().GetNRcgi().GetPLmnIdentity().GetValue())
	if err != nil {
		return nodeIdentity, nil, fmt.Errorf("error extracting to plmnId for NrCgi")
	}
	fmt.Printf("Am i getting there?\n")
	indicationHeader.NRcgi.PLmnIdentity = plmnIdNrCgi
	//indicationHeader.NRcgi, err = types.NewNrCellPlmnId(e2smKpmPdu.GetIndicationHeaderFormat1().GetNRcgi().GetPLmnIdentity().GetValue())

	fmt.Printf("Hooray?\n")

	indicationHeader.NRcgi.NRcellIdentity.Value = types.NRcellIdentityValue(e2smKpmPdu.GetIndicationHeaderFormat1().GetNRcgi().GetNRcellIdentity().GetValue().GetValue())
	indicationHeader.NRcgi.NRcellIdentity.Len = types.NRcellIdentityLen(e2smKpmPdu.GetIndicationHeaderFormat1().GetNRcgi().GetNRcellIdentity().GetValue().GetLen())
	fmt.Printf("Hooray!\n")

	plmnId, err := types.PlmnIDFromSlice(e2smKpmPdu.GetIndicationHeaderFormat1().GetPLmnIdentity().GetValue())
	if err != nil {
		return nodeIdentity, nil, fmt.Errorf("error extracting to plmnId")
	}
	indicationHeader.PLmnIdentity = plmnId

	sst, err := types.SstFromSlice(e2smKpmPdu.GetIndicationHeaderFormat1().GetSliceId().GetSSt())
	if err != nil {
		return nodeIdentity, nil, fmt.Errorf("error extracting to sst")
	}
	indicationHeader.SliceId.SSt = sst

	sd, err := types.SdFromSlice(e2smKpmPdu.GetIndicationHeaderFormat1().GetSliceId().GetSD())
	if err != nil {
		return nodeIdentity, nil, fmt.Errorf("error extracting to sd")
	}
	indicationHeader.SliceId.SD = sd

	indicationHeader.FiveQi = types.FiveQi(e2smKpmPdu.GetIndicationHeaderFormat1().GetFiveQi())
	indicationHeader.Qci = types.Qci(e2smKpmPdu.GetIndicationHeaderFormat1().GetQci())

	fmt.Printf("indicationHeader -- %v\n", indicationHeader)

	return nodeIdentity, indicationHeader, nil
}
