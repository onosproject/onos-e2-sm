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
	RetrieveRicEventTriggerStyleListItemByRicStyleType(ricStyleType int32)
	RetrieveRicEventTriggerStyleListItemByRicStyleName(ricStyleName string)
	RetrieveRicEventTriggerStyleListItemByRicFormatType(ricFormatType int32)
	AddRicReportStyleListItem(ricReportStyleList *RicReportStyleList)
	RetrieveRicReportStyleListItemByRicStyleType(ricStyleType int32)
	RetrieveRicReportStyleListItemByRicStyleName(ricStyleName string)
	RetrieveRicReportStyleListItemByIndicationHeader(indHdrFormatType int32)
	RetrieveRicReportStyleListItemByIndicationMessage(indMsgFormatType int32)
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
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicEventTriggerStyleListItemByRicStyleType(ricStyleType int32) *RicEventTriggerStyleList {
	for _, rEtSlIe := range b.RicEventTriggerStyleList {
		if rEtSlIe.RicEventTriggerStyleType.Value == ricStyleType {
			return rEtSlIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicEventTriggerStyleListItemByRicStyleName(ricStyleName string) *RicEventTriggerStyleList {
	for _, rEtSlIe := range b.RicEventTriggerStyleList {
		if rEtSlIe.RicEventTriggerStyleName.Value == ricStyleName {
			return rEtSlIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicEventTriggerStyleListItemByRicFormatType(ricFormatType int32) *RicEventTriggerStyleList {
	for _, rEtSlIe := range b.RicEventTriggerStyleList {
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
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicReportStyleListItemByRicStyleType(ricStyleType int32) *RicReportStyleList {
	for _, rRsLIe := range b.RicReportStyleList {
		if rRsLIe.RicReportStyleType.Value == ricStyleType {
			return rRsLIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicReportStyleListItemByRicStyleName(ricStyleName string) *RicReportStyleList {
	for _, rRsLIe := range b.RicReportStyleList {
		if rRsLIe.RicReportStyleName.Value == ricStyleName {
			return rRsLIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicReportStyleListItemByIndicationHeader(indHdrFormatType int32) *RicReportStyleList {
	for _, rRsLIe := range b.RicReportStyleList {
		if rRsLIe.RicIndicationHeaderFormatType.Value == indHdrFormatType {
			return rRsLIe
		}
	}
	return nil
}

//TODO: In the future may return list of items which satisfy this condition
func (b *E2SmKpmRanfunctionDescriptionE2SmKpmRanfunctionItem001) RetrieveRicReportStyleListItemByIndicationMessage(indMsgFormatType int32) *RicReportStyleList {
	for _, rRsLIe := range b.RicReportStyleList {
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
