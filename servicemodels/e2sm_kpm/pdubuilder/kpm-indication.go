// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
)

func CreateE2SmKpmIndicationMsg(plmnID string, SSt string, SD string, ranContainer string) (*e2sm_kpm_ies.E2SmKpmIndicationMessage, error) {
	if len(plmnID) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}
	if len(SSt) != 1 {
		return nil, fmt.Errorf("error: SSt should be 1 char")
	}
	if len(SD) != 3 {
		return nil, fmt.Errorf("error: SD should be 3 chars")
	}

	cellResourceReport := e2sm_kpm_ies.CellResourceReportListItem{
		NRcgi: &e2sm_kpm_ies.Nrcgi{
			PLmnIdentity: &e2sm_kpm_ies.PlmnIdentity{
				Value: []byte(plmnID),
			},
			NRcellIdentity: &e2sm_kpm_ies.NrcellIdentity{
				Value: &e2sm_kpm_ies.BitString{
					Value: 0, //uint64
					Len:   0, //uint32
				},
			},
		},
		DlTotalofAvailablePrbs: 0, //int32
		UlTotalofAvailablePrbs: 0, //int32
		ServedPlmnPerCellList:  make([]*e2sm_kpm_ies.ServedPlmnPerCellListItem, 0),
	}

	serverPlmCells := e2sm_kpm_ies.ServedPlmnPerCellListItem{
		PLmnIdentity: &e2sm_kpm_ies.PlmnIdentity{
			Value: []byte(plmnID),
		},
		DuPm_5Gc: &e2sm_kpm_ies.FgcDuPmContainer{
			SlicePerPlmnPerCellList: make([]*e2sm_kpm_ies.SlicePerPlmnPerCellListItem, 0),
		},
		DuPmEpc: &e2sm_kpm_ies.EpcDuPmContainer{
			PerQcireportList: make([]*e2sm_kpm_ies.PerQcireportListItem, 0),
		},
	}

	slicePerPlmn := e2sm_kpm_ies.SlicePerPlmnPerCellListItem{
		SliceId: &e2sm_kpm_ies.Snssai{
			SSt: []byte(SSt),
			SD:  []byte(SD),
		},
		FQiperslicesPerPlmnPerCellList: make([]*e2sm_kpm_ies.FqiperslicesPerPlmnPerCellListItem, 0),
	}

	fQuipSlPerPlmnPerCell := e2sm_kpm_ies.FqiperslicesPerPlmnPerCellListItem{
		FiveQi:     0, //int32
		DlPrbusage: 0, //int32
		UlPrbusage: 0, //int32
	}
	slicePerPlmn.FQiperslicesPerPlmnPerCellList = append(slicePerPlmn.FQiperslicesPerPlmnPerCellList, &fQuipSlPerPlmnPerCell)

	perQcireportItem := e2sm_kpm_ies.PerQcireportListItem{
		Qci:        0, //int32
		DlPrbusage: 0, //int32
		UlPrbusage: 0, //int32
	}
	serverPlmCells.DuPmEpc.PerQcireportList = append(serverPlmCells.DuPmEpc.PerQcireportList, &perQcireportItem)
	serverPlmCells.DuPm_5Gc.SlicePerPlmnPerCellList = append(serverPlmCells.DuPm_5Gc.SlicePerPlmnPerCellList, &slicePerPlmn)
	cellResourceReport.ServedPlmnPerCellList = append(cellResourceReport.ServedPlmnPerCellList, &serverPlmCells)

	oduContainer := e2sm_kpm_ies.OduPfContainer{
		CellResourceReportList: make([]*e2sm_kpm_ies.CellResourceReportListItem, 0),
	}
	oduContainer.CellResourceReportList = append(oduContainer.CellResourceReportList, &cellResourceReport)

	containers := e2sm_kpm_ies.PmContainersList{
		PerformanceContainer: &e2sm_kpm_ies.PfContainer{
			PfContainer: &e2sm_kpm_ies.PfContainer_ODu{
				ODu: &oduContainer,
			},
		},
		TheRancontainer: &e2sm_kpm_ies.RanContainer{
			Value: []byte(ranContainer),
		},
	}

	e2SmIindicationMsg := e2sm_kpm_ies.E2SmKpmIndicationMessage_IndicationMessageFormat1{
		IndicationMessageFormat1: &e2sm_kpm_ies.E2SmKpmIndicationMessageFormat1{
			PmContainers: make([]*e2sm_kpm_ies.PmContainersList, 0),
		},
	}
	e2SmIindicationMsg.IndicationMessageFormat1.PmContainers = append(e2SmIindicationMsg.IndicationMessageFormat1.PmContainers, &containers)

	e2smKpmPdu := e2sm_kpm_ies.E2SmKpmIndicationMessage{
		E2SmKpmIndicationMessage: &e2SmIindicationMsg,
	}

	if err := e2smKpmPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2smKpmPdu, nil
}
