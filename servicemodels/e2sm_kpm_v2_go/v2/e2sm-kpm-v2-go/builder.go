// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2smkpmv2go

func (mii *MeasurementInfoItem) SetLabelInfoList(labelInfoList *LabelInfoList) *MeasurementInfoItem {
	mii.LabelInfoList = labelInfoList
	return mii
}

func (ml *MeasurementLabel) SetPlmnID(plmnID *PlmnIdentity) *MeasurementLabel {
	ml.PlmnId = plmnID
	return ml
}

func (s *Snssai) SetSliceID(sd []byte) *Snssai {
	s.SD = sd
	return s
}

func (ml *MeasurementLabel) SetSliceID(sliceID *Snssai) *MeasurementLabel {
	ml.SliceId = sliceID
	return ml
}

func (ml *MeasurementLabel) SetFiveQi(fiveQi int32) *MeasurementLabel {
	ml.FiveQi = &FiveQi{
		Value: fiveQi,
	}
	return ml
}

func (ml *MeasurementLabel) SetQfi(qfi int32) *MeasurementLabel {
	ml.QFi = &Qfi{
		Value: qfi,
	}
	return ml
}

func (ml *MeasurementLabel) SetQci(qci int32) *MeasurementLabel {
	ml.QCi = &Qci{
		Value: qci,
	}
	return ml
}

func (ml *MeasurementLabel) SetQciMax(qci int32) *MeasurementLabel {
	ml.QCimax = &Qci{
		Value: qci,
	}
	return ml
}

func (ml *MeasurementLabel) SetQciMin(qci int32) *MeasurementLabel {
	ml.QCimin = &Qci{
		Value: qci,
	}
	return ml
}

func (ml *MeasurementLabel) SetArpMax(arp int32) *MeasurementLabel {
	ml.ARpmax = &Arp{
		Value: arp,
	}
	return ml
}

func (ml *MeasurementLabel) SetArpMin(arp int32) *MeasurementLabel {
	ml.ARpmin = &Arp{
		Value: arp,
	}
	return ml
}

func (ml *MeasurementLabel) SetBitRange(br int32) *MeasurementLabel {
	ml.BitrateRange = &br
	return ml
}

func (ml *MeasurementLabel) SetLayerMuMIMO(lmm int32) *MeasurementLabel {
	ml.LayerMuMimo = &lmm
	return ml
}

func (ml *MeasurementLabel) SetSUm() *MeasurementLabel {
	sum := SUM_SUM_TRUE
	ml.SUm = &sum
	return ml
}

func (ml *MeasurementLabel) SetDistBinX(distBinX int32) *MeasurementLabel {
	ml.DistBinX = &distBinX
	return ml
}

func (ml *MeasurementLabel) SetDistBinY(distBinY int32) *MeasurementLabel {
	ml.DistBinY = &distBinY
	return ml
}

func (ml *MeasurementLabel) SetDistBinZ(distBinZ int32) *MeasurementLabel {
	ml.DistBinZ = &distBinZ
	return ml
}

func (ml *MeasurementLabel) SetPreLabelOverride() *MeasurementLabel {
	plo := PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	ml.PreLabelOverride = &plo
	return ml
}

func (ml *MeasurementLabel) SetStartEndIndication(sei StartEndInd) *MeasurementLabel {
	ml.StartEndInd = &sei
	return ml
}

func (ih *E2SmKpmIndicationHeader) SetFileFormatVersion(ffv string) *E2SmKpmIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmKpmIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().FileFormatversion = &ffv
	default:
		return ih
	}
	return ih
}

func (ih *E2SmKpmIndicationHeader) SetSenderName(sn string) *E2SmKpmIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmKpmIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().SenderName = &sn
	default:
		return ih
	}
	return ih
}

func (ih *E2SmKpmIndicationHeader) SetSenderType(st string) *E2SmKpmIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmKpmIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().SenderType = &st
	default:
		return ih
	}
	return ih
}

func (ih *E2SmKpmIndicationHeader) SetVendorName(vn string) *E2SmKpmIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmKpmIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().VendorName = &vn
	default:
		return ih
	}
	return ih
}

func (ih *E2SmKpmIndicationHeader) SetGlobalKPMnodeID(gknID *GlobalKpmnodeId) *E2SmKpmIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmKpmIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().KpmNodeId = gknID
	default:
		return ih
	}
	return ih
}

func (im *E2SmKpmIndicationMessage) SetMeasInfoList(measInfoList *MeasurementInfoList) *E2SmKpmIndicationMessage {
	switch im.IndicationMessageFormats.E2SmKpmIndicationMessage.(type) {
	case *IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().MeasInfoList = measInfoList
	// Left for future extensions
	//case *IndicationMessageFormats_IndicationMessageFormat2:
	//	im.GetIndicationMessageFormats().GetIndicationMessageFormat2().MeasInfoList = measInfoList
	default:
		return im
	}

	return im
}

func (im *E2SmKpmIndicationMessage) SetGranularityPeriod(gp int64) *E2SmKpmIndicationMessage {
	switch im.IndicationMessageFormats.E2SmKpmIndicationMessage.(type) {
	case *IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GranulPeriod = &GranularityPeriod{
			Value: gp,
		}
	case *IndicationMessageFormats_IndicationMessageFormat2:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GranulPeriod = &GranularityPeriod{
			Value: gp,
		}
	default:
		return im
	}

	return im
}

func (im *E2SmKpmIndicationMessage) SetCellObjectID(cellObjID string) *E2SmKpmIndicationMessage {
	switch im.IndicationMessageFormats.E2SmKpmIndicationMessage.(type) {
	case *IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().CellObjId = &CellObjectId{
			Value: cellObjID,
		}
	case *IndicationMessageFormats_IndicationMessageFormat2:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat2().CellObjId = &CellObjectId{
			Value: cellObjID,
		}
	default:
		return im
	}

	return im
}

func (mci *MeasurementCondUeidItem) SetMatchingUEUDlist(ml *MatchingUeidList) *MeasurementCondUeidItem {
	mci.MatchingUeidList = ml
	return mci
}

func (mdi *MeasurementDataItem) SetIncompleteFlag() *MeasurementDataItem {
	incf := IncompleteFlag_INCOMPLETE_FLAG_TRUE
	mdi.IncompleteFlag = &incf
	return mdi
}

func (rfd *E2SmKpmRanfunctionDescription) SetRanFunctionInstance(rfi int32) *E2SmKpmRanfunctionDescription {
	rfd.GetRanFunctionName().RanFunctionInstance = &rfi
	return rfd
}

func (rfd *E2SmKpmRanfunctionDescription) SetRicKpmNodeList(rknl []*RicKpmnodeItem) *E2SmKpmRanfunctionDescription {
	rfd.RicKpmNodeList = rknl
	return rfd
}

func (rfd *E2SmKpmRanfunctionDescription) SetRicEventTriggerStyleList(retsl []*RicEventTriggerStyleItem) *E2SmKpmRanfunctionDescription {
	rfd.RicEventTriggerStyleList = retsl
	return rfd
}

func (rfd *E2SmKpmRanfunctionDescription) SetRicReportStyleList(rrsl []*RicReportStyleItem) *E2SmKpmRanfunctionDescription {
	rfd.RicReportStyleList = rrsl
	return rfd
}

func (rkni *RicKpmnodeItem) SetCellMeasurementObjectList(cmol []*CellMeasurementObjectItem) *RicKpmnodeItem {
	rkni.CellMeasurementObjectList = cmol
	return rkni
}
