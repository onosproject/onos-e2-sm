// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RT-Period-IE.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"unsafe"
)

//ToDo: Solve "Cannot convert rtPeriodIeC (type _Ctype_long) to type unsafe.Pointer"
func xerEncodeRtPeriodIe(rtPeriodIe e2sm_rc_pre_ies.RtPeriodIe) ([]byte, error) {
	rtPeriodIeC, err := newRtPeriodIe(rtPeriodIe)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RT_Period_IE, unsafe.Pointer(&rtPeriodIeC))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRtPeriodIe() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRtPeriodIe(rtPeriodIe e2sm_rc_pre_ies.RtPeriodIe) ([]byte, error) {
	rtPeriodIeC, err := newRtPeriodIe(rtPeriodIe)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RT_Period_IE, unsafe.Pointer(&rtPeriodIeC))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRtPeriodIe() %s", err.Error())
	}
	return bytes, nil
}

//ToDo: Decide what to return instead of nil, which cannot be returned since return value is not a pointer anymore
func xerDecodeRtPeriodIe(bytes []byte) (e2sm_rc_pre_ies.RtPeriodIe, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RT_Period_IE)
	if err != nil {
		return 0, err
	}
	if unsafePtr == nil {
		return 0, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRtPeriodIe((*C.RT_Period_IE_t)(unsafePtr)), nil
}

func perDecodeRtPeriodIe(bytes []byte) (e2sm_rc_pre_ies.RtPeriodIe, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RT_Period_IE)
	if err != nil {
		return 0, err
	}
	if unsafePtr == nil {
		return 0, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRtPeriodIe((*C.RT_Period_IE_t)(unsafePtr)), nil
}

func newRtPeriodIe(rtPeriodIe e2sm_rc_pre_ies.RtPeriodIe) (C.RT_Period_IE_t, error) {
	var ret C.RT_Period_IE_t
	switch rtPeriodIe {
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS10:
		ret = C.RT_Period_IE_ms10
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS20:
		ret = C.RT_Period_IE_ms20
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS32:
		ret = C.RT_Period_IE_ms32
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS40:
		ret = C.RT_Period_IE_ms40
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS60:
		ret = C.RT_Period_IE_ms60
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS64:
		ret = C.RT_Period_IE_ms64
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS70:
		ret = C.RT_Period_IE_ms70
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS80:
		ret = C.RT_Period_IE_ms80
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS128:
		ret = C.RT_Period_IE_ms128
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS160:
		ret = C.RT_Period_IE_ms160
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS256:
		ret = C.RT_Period_IE_ms256
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS320:
		ret = C.RT_Period_IE_ms320
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS512:
		ret = C.RT_Period_IE_ms512
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS640:
		ret = C.RT_Period_IE_ms640
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS1024:
		ret = C.RT_Period_IE_ms1024
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS1280:
		ret = C.RT_Period_IE_ms1280
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS2048:
		ret = C.RT_Period_IE_ms2048
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS2560:
		ret = C.RT_Period_IE_ms2560
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS5120:
		ret = C.RT_Period_IE_ms5120
	case e2sm_rc_pre_ies.RtPeriodIe_RT_PERIOD_IE_MS10240:
		ret = C.RT_Period_IE_ms10240
	default:
		return 0, fmt.Errorf("unexpected RT-Period-IE %v", rtPeriodIe)
	}
	return ret, nil
}

func decodeRtPeriodIe(rtPeriodIeC *C.RT_Period_IE_t) e2sm_rc_pre_ies.RtPeriodIe {

	return e2sm_rc_pre_ies.RtPeriodIe(int32(*rtPeriodIeC))
}
