// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementLabel.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeMeasurementLabel(measurementLabel *e2sm_kpm_v2.MeasurementLabel) ([]byte, error) {
	measurementLabelCP, err := newMeasurementLabel(measurementLabel)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementLabel() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementLabel, unsafe.Pointer(measurementLabelCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementLabel() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementLabel(measurementLabel *e2sm_kpm_v2.MeasurementLabel) ([]byte, error) {
	measurementLabelCP, err := newMeasurementLabel(measurementLabel)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementLabel() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementLabel, unsafe.Pointer(measurementLabelCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementLabel() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementLabel(bytes []byte) (*e2sm_kpm_v2.MeasurementLabel, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementLabel)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementLabel((*C.MeasurementLabel_t)(unsafePtr))
}

func perDecodeMeasurementLabel(bytes []byte) (*e2sm_kpm_v2.MeasurementLabel, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementLabel)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementLabel((*C.MeasurementLabel_t)(unsafePtr))
}

func newMeasurementLabel(measurementLabel *e2sm_kpm_v2.MeasurementLabel) (*C.MeasurementLabel_t, error) {

	plmnIDC, err := newPlmnIdentity(measurementLabel.PlmnId)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	sliceIDC, err := newSnssai(measurementLabel.SliceId)
	if err != nil {
		return nil, fmt.Errorf("newSnssai() %s", err.Error())
	}

	fiveQiC, err := newFiveQi(measurementLabel.FiveQi)
	if err != nil {
		return nil, fmt.Errorf("newFiveQi() %s", err.Error())
	}

	qCiC, err := newQci(measurementLabel.QCi)
	if err != nil {
		return nil, fmt.Errorf("newQci() %s", err.Error())
	}

	qCimaxC, err := newQci(measurementLabel.QCimax)
	if err != nil {
		return nil, fmt.Errorf("newQci() %s", err.Error())
	}

	qCiminC, err := newQci(measurementLabel.QCimin)
	if err != nil {
		return nil, fmt.Errorf("newQci() %s", err.Error())
	}

	aRpmaxC, err := newArp(measurementLabel.ARpmax)
	if err != nil {
		return nil, fmt.Errorf("newArp() %s", err.Error())
	}

	aRpminC, err := newArp(measurementLabel.ARpmin)
	if err != nil {
		return nil, fmt.Errorf("newArp() %s", err.Error())
	}

	bitrateRangeC := C.long(measurementLabel.BitrateRange)
	layerMuMimoC := C.long(measurementLabel.LayerMuMimo)
	sUmC := C.long(measurementLabel.SUm)
	distBinXC := C.long(measurementLabel.DistBinX)
	distBinYC := C.long(measurementLabel.DistBinY)
	distBinZC := C.long(measurementLabel.DistBinZ)
	preLabelOverrideC := C.long(measurementLabel.PreLabelOverride)
	startEndIndC := C.long(measurementLabel.StartEndInd)
	measurementLabelC := C.MeasurementLabel_t{
		//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
		plmnID:           plmnIDC,
		sliceID:          sliceIDC,
		fiveQI:           fiveQiC,
		qCI:              qCiC,
		qCImax:           qCimaxC,
		qCImin:           qCiminC,
		aRPmax:           aRpmaxC,
		aRPmin:           aRpminC,
		bitrateRange:     &bitrateRangeC,
		layerMU_MIMO:     &layerMuMimoC,
		sUM:              &sUmC,
		distBinX:         &distBinXC,
		distBinY:         &distBinYC,
		distBinZ:         &distBinZC,
		preLabelOverride: &preLabelOverrideC,
		startEndInd:      &startEndIndC,
	}

	return &measurementLabelC, nil
}

func decodeMeasurementLabel(measurementLabelC *C.MeasurementLabel_t) (*e2sm_kpm_v2.MeasurementLabel, error) {

	plmnID, err := decodePlmnIdentity(measurementLabelC.plmnID)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	sliceID, err := decodeSnssai(measurementLabelC.sliceID)
	if err != nil {
		return nil, fmt.Errorf("decodeSnssai() %s", err.Error())
	}

	fiveQi, err := decodeFiveQi(measurementLabelC.fiveQI)
	if err != nil {
		return nil, fmt.Errorf("decodeFiveQi() %s", err.Error())
	}

	qCi, err := decodeQci(measurementLabelC.qCI)
	if err != nil {
		return nil, fmt.Errorf("decodeQci() %s", err.Error())
	}

	qCimax, err := decodeQci(measurementLabelC.qCImax)
	if err != nil {
		return nil, fmt.Errorf("decodeQci() %s", err.Error())
	}

	qCimin, err := decodeQci(measurementLabelC.qCImin)
	if err != nil {
		return nil, fmt.Errorf("decodeQci() %s", err.Error())
	}

	aRpmax, err := decodeArp(measurementLabelC.aRPmax)
	if err != nil {
		return nil, fmt.Errorf("decodeArp() %s", err.Error())
	}

	aRpmin, err := decodeArp(measurementLabelC.aRPmin)
	if err != nil {
		return nil, fmt.Errorf("decodeArp() %s", err.Error())
	}

	bitrateRange := int32(*measurementLabelC.bitrateRange)
	layerMuMimo := int32(*measurementLabelC.layerMU_MIMO)
	sUm := int32(*measurementLabelC.sUM)
	distBinX := int32(*measurementLabelC.distBinX)
	distBinY := int32(*measurementLabelC.distBinY)
	distBinZ := int32(*measurementLabelC.distBinZ)
	preLabelOverride := int32(*measurementLabelC.preLabelOverride)
	startEndInd := int32(*measurementLabelC.startEndInd)
	measurementLabel := e2sm_kpm_v2.MeasurementLabel{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		PlmnId:           plmnID,
		SliceId:          sliceID,
		FiveQi:           fiveQi,
		QCi:              qCi,
		QCimax:           qCimax,
		QCimin:           qCimin,
		ARpmax:           aRpmax,
		ARpmin:           aRpmin,
		BitrateRange:     bitrateRange,
		LayerMuMimo:      layerMuMimo,
		SUm:              sUm,
		DistBinX:         distBinX,
		DistBinY:         distBinY,
		DistBinZ:         distBinZ,
		PreLabelOverride: preLabelOverride,
		StartEndInd:      startEndInd,
	}

	return &measurementLabel, nil
}

func decodeMeasurementLabelBytes(array [8]byte) (*e2sm_kpm_v2.MeasurementLabel, error) {
	measurementLabelC := (*C.MeasurementLabel_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeMeasurementLabel(measurementLabelC)
}
