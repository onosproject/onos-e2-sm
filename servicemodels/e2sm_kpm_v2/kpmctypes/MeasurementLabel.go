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
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
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

	measurementLabelC := C.MeasurementLabel_t{}
	var err error

	if measurementLabel.PlmnId != nil {
		measurementLabelC.plmnID, err = newPlmnIdentity(measurementLabel.PlmnId)
		if err != nil {
			return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
		}
	}

	if measurementLabel.SliceId != nil {
		measurementLabelC.sliceID, err = newSnssai(measurementLabel.SliceId)
		if err != nil {
			return nil, fmt.Errorf("newSnssai() %s", err.Error())
		}
	}

	if measurementLabel.FiveQi != nil {
		measurementLabelC.fiveQI, err = newFiveQi(measurementLabel.FiveQi)
		if err != nil {
			return nil, fmt.Errorf("newFiveQi() %s", err.Error())
		}
	}

	if measurementLabel.QFi != nil {
		measurementLabelC.qFI, err = newQfi(measurementLabel.QFi)
		if err != nil {
			return nil, fmt.Errorf("newQfi() %s", err.Error())
		}
	}

	if measurementLabel.QCi != nil {
		measurementLabelC.qCI, err = newQci(measurementLabel.QCi)
		if err != nil {
			return nil, fmt.Errorf("newQci() %s", err.Error())
		}
	}

	if measurementLabel.QCimax != nil {
		measurementLabelC.qCImax, err = newQci(measurementLabel.QCimax)
		if err != nil {
			return nil, fmt.Errorf("newQci() %s", err.Error())
		}
	}

	if measurementLabel.QCimin != nil {
		measurementLabelC.qCImin, err = newQci(measurementLabel.QCimin)
		if err != nil {
			return nil, fmt.Errorf("newQci() %s", err.Error())
		}
	}

	if measurementLabel.ARpmax != nil {
		measurementLabelC.aRPmax, err = newArp(measurementLabel.ARpmax)
		if err != nil {
			return nil, fmt.Errorf("newArp() %s", err.Error())
		}
	}

	if measurementLabel.ARpmin != nil {
		measurementLabelC.aRPmin, err = newArp(measurementLabel.ARpmin)
		if err != nil {
			return nil, fmt.Errorf("newArp() %s", err.Error())
		}
	}

	if measurementLabel.BitrateRange != -1 {
		bitrateRangeC := C.long(measurementLabel.BitrateRange)
		measurementLabelC.bitrateRange = &bitrateRangeC
	}
	if measurementLabel.LayerMuMimo != -1 {
		layerMuMimoC := C.long(measurementLabel.LayerMuMimo)
		measurementLabelC.layerMU_MIMO = &layerMuMimoC
	}

	if measurementLabel.SUm != -1 {
		var sUmC C.MeasurementLabel__sUM_t
		switch measurementLabel.SUm {
		case e2sm_kpm_v2.SUM_SUM_TRUE:
			sUmC = C.MeasurementLabel__sUM_true
		default:
			return nil, fmt.Errorf("unexpected MeasurementLabel SUm %v", measurementLabel.SUm)
		}
		measurementLabelC.sUM = &sUmC
	}

	if measurementLabel.DistBinX != -1 {
		distBinXC := C.long(measurementLabel.DistBinX)
		measurementLabelC.distBinX = &distBinXC
	}

	if measurementLabel.DistBinY != -1 {
		distBinYC := C.long(measurementLabel.DistBinY)
		measurementLabelC.distBinY = &distBinYC
	}

	if measurementLabel.DistBinZ != -1 {
		distBinZC := C.long(measurementLabel.DistBinZ)
		measurementLabelC.distBinZ = &distBinZC
	}

	if measurementLabel.PreLabelOverride != -1 {
		var preLabelOverrideC C.MeasurementLabel__preLabelOverride_t
		switch measurementLabel.PreLabelOverride {
		case e2sm_kpm_v2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE:
			preLabelOverrideC = C.MeasurementLabel__preLabelOverride_true
		default:
			return nil, fmt.Errorf("unexpected MeasurementLabel PreLabelOverride %v", measurementLabel.PreLabelOverride)
		}
		measurementLabelC.preLabelOverride = &preLabelOverrideC
	}

	if measurementLabel.StartEndInd != -1 {
		var startEndIndC C.MeasurementLabel__startEndInd_t
		switch measurementLabel.StartEndInd {
		case e2sm_kpm_v2.StartEndInd_START_END_IND_START:
			startEndIndC = C.MeasurementLabel__startEndInd_start
		case e2sm_kpm_v2.StartEndInd_START_END_IND_END:
			startEndIndC = C.MeasurementLabel__startEndInd_end
		default:
			return nil, fmt.Errorf("unexpected MeasurementLabel StartEndInd %v", measurementLabel.StartEndInd)
		}
		measurementLabelC.startEndInd = &startEndIndC
	}

	return &measurementLabelC, nil
}

func decodeMeasurementLabel(measurementLabelC *C.MeasurementLabel_t) (*e2sm_kpm_v2.MeasurementLabel, error) {

	measurementLabel := new(e2sm_kpm_v2.MeasurementLabel)
	var err error

	if measurementLabelC.plmnID != nil {
		measurementLabel.PlmnId, err = decodePlmnIdentity(measurementLabelC.plmnID)
		if err != nil {
			return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
		}
	}

	if measurementLabelC.sliceID != nil {
		measurementLabel.SliceId, err = decodeSnssai(measurementLabelC.sliceID)
		if err != nil {
			return nil, fmt.Errorf("decodeSnssai() %s", err.Error())
		}
	}

	if measurementLabelC.fiveQI != nil {
		measurementLabel.FiveQi, err = decodeFiveQi(measurementLabelC.fiveQI)
		if err != nil {
			return nil, fmt.Errorf("decodeFiveQi() %s", err.Error())
		}
	}

	if measurementLabelC.qFI != nil {
		measurementLabel.QFi, err = decodeQfi(measurementLabelC.qFI)
		if err != nil {
			return nil, fmt.Errorf("decodeQfi() %s", err.Error())
		}
	}

	if measurementLabelC.qCI != nil {
		measurementLabel.QCi, err = decodeQci(measurementLabelC.qCI)
		if err != nil {
			return nil, fmt.Errorf("decodeQci() %s", err.Error())
		}
	}

	if measurementLabelC.qCImax != nil {
		measurementLabel.QCimax, err = decodeQci(measurementLabelC.qCImax)
		if err != nil {
			return nil, fmt.Errorf("decodeQci() %s", err.Error())
		}
	}

	if measurementLabelC.qCImin != nil {
		measurementLabel.QCimin, err = decodeQci(measurementLabelC.qCImin)
		if err != nil {
			return nil, fmt.Errorf("decodeQci() %s", err.Error())
		}
	}

	if measurementLabelC.aRPmax != nil {
		measurementLabel.ARpmax, err = decodeArp(measurementLabelC.aRPmax)
		if err != nil {
			return nil, fmt.Errorf("decodeArp() %s", err.Error())
		}
	}

	if measurementLabelC.aRPmin != nil {
		measurementLabel.ARpmin, err = decodeArp(measurementLabelC.aRPmin)
		if err != nil {
			return nil, fmt.Errorf("decodeArp() %s", err.Error())
		}
	}

	if measurementLabelC.bitrateRange != nil {
		measurementLabel.BitrateRange = int32(*measurementLabelC.bitrateRange)
	}

	if measurementLabelC.layerMU_MIMO != nil {
		measurementLabel.LayerMuMimo = int32(*measurementLabelC.layerMU_MIMO)
	}

	if measurementLabelC.sUM != nil {
		measurementLabel.SUm = e2sm_kpm_v2.SUM_SUM_TRUE
	}

	if measurementLabelC.distBinX != nil {
		measurementLabel.DistBinX = int32(*measurementLabelC.distBinX)
	}
	if measurementLabelC.distBinY != nil {
		measurementLabel.DistBinY = int32(*measurementLabelC.distBinY)
	}
	if measurementLabelC.distBinZ != nil {
		measurementLabel.DistBinZ = int32(*measurementLabelC.distBinZ)
	}

	if measurementLabelC.preLabelOverride != nil {
		measurementLabel.PreLabelOverride = e2sm_kpm_v2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	}
	if measurementLabelC.startEndInd != nil {
		measurementLabel.StartEndInd = e2sm_kpm_v2.StartEndInd(int32(*measurementLabelC.startEndInd))
	}

	return measurementLabel, nil
}

func decodeMeasurementLabelBytes(array [8]byte) (*e2sm_kpm_v2.MeasurementLabel, error) {
	mlC := (*C.MeasurementLabel_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))

	return decodeMeasurementLabel(mlC)
}
