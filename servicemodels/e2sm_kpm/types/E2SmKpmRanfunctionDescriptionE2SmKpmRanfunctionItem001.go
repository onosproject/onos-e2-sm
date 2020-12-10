// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

type E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001 struct {
	RicEventTriggerStyleList []*RicEventTriggerStyleList
	RicReportStyleList       []*RicReportStyleList
}

type E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001Builder interface {
	NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001()
	AddRicEventTriggerStyleListItem(ricEventTriggerStyleList *RicEventTriggerStyleList)
	RetrieveRicEventTriggerStyleListItemByRicStyleType(ricEventTriggerStyleList []*RicEventTriggerStyleList, ricStyleType int32)
	RetrieveRicEventTriggerStyleListItemByRicStyleName(ricEventTriggerStyleList []*RicEventTriggerStyleList, ricStyleName string)
	RetrieveRicEventTriggerStyleListItemByRicFormatType(ricEventTriggerStyleList []*RicEventTriggerStyleList, ricFormatType int32)
	AddRicReportStyleListItem(ricReportStyleList *RicReportStyleList)
	RetrieveRicReportStyleListItemByRicStyleType(ricReportStyleList []*RicReportStyleList, ricStyleType int32)
	RetrieveRicReportStyleListItemByRicStyleName(ricReportStyleList []*RicReportStyleList, ricStyleName string)
	RetrieveRicReportStyleListItemByIndicationHeader(ricReportStyleList []*RicReportStyleList, indHdrFormatType int32)
	RetrieveRicReportStyleListItemByIndicationMessage(ricReportStyleList []*RicReportStyleList, indMsgFormatType int32)
	GetRicEventTriggerStyleList()
	GetRicReportStyleList()
	GetE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001()
}

func NewE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001() *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001 {
	return &E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001{}
}

func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) AddRicEventTriggerStyleListItem(ricEventTriggerStyleList *RicEventTriggerStyleList) *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001 {
	b.RicEventTriggerStyleList = append(b.RicEventTriggerStyleList, ricEventTriggerStyleList)
	return b
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicEventTriggerStyleListItemByRicStyleType(ricEventTriggerStyleList []*RicEventTriggerStyleList, ricStyleType int32) *RicEventTriggerStyleList {
	for _, rEtSlIe := range ricEventTriggerStyleList {
		if rEtSlIe.RicEventTriggerStyleType.Value == ricStyleType {
			return rEtSlIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicEventTriggerStyleListItemByRicStyleName(ricEventTriggerStyleList []*RicEventTriggerStyleList, ricStyleName string) *RicEventTriggerStyleList {
	for _, rEtSlIe := range ricEventTriggerStyleList {
		if rEtSlIe.RicEventTriggerStyleName.Value == ricStyleName {
			return rEtSlIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicEventTriggerStyleListItemByRicFormatType(ricEventTriggerStyleList []*RicEventTriggerStyleList, ricFormatType int32) *RicEventTriggerStyleList {
	for _, rEtSlIe := range ricEventTriggerStyleList {
		if rEtSlIe.RicEventTriggerFormatType.Value == ricFormatType {
			return rEtSlIe
		}
	}
	return nil
}

func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) AddRicReportStyleListItem(ricReportStyleList *RicReportStyleList) *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001 {
	b.RicReportStyleList = append(b.RicReportStyleList, ricReportStyleList)
	return b
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicReportStyleListItemByRicStyleType(ricReportStyleList []*RicReportStyleList, ricStyleType int32) *RicReportStyleList {
	for _, rRsLIe := range ricReportStyleList {
		if rRsLIe.RicReportStyleType.Value == ricStyleType {
			return rRsLIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicReportStyleListItemByRicStyleName(ricReportStyleList []*RicReportStyleList, ricStyleName string) *RicReportStyleList {
	for _, rRsLIe := range ricReportStyleList {
		if rRsLIe.RicReportStyleName.Value == ricStyleName {
			return rRsLIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicReportStyleListItemByIndicationHeader(ricReportStyleList []*RicReportStyleList, indHdrFormatType int32) *RicReportStyleList {
	for _, rRsLIe := range ricReportStyleList {
		if rRsLIe.RicIndicationHeaderFormatType.Value == indHdrFormatType {
			return rRsLIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicReportStyleListItemByIndicationMessage(ricReportStyleList []*RicReportStyleList, indMsgFormatType int32) *RicReportStyleList {
	for _, rRsLIe := range ricReportStyleList {
		if rRsLIe.RicIndicationMessageFormatType.Value == indMsgFormatType {
			return rRsLIe
		}
	}
	return nil
}

func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) GetRicEventTriggerStyleList() []*RicEventTriggerStyleList {
	return b.RicEventTriggerStyleList
}

func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) GetRicReportStyleList() []*RicReportStyleList {
	return b.RicReportStyleList
}

func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) GetE2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001() *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001 {
	return b
}
