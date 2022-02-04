// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ENB-ID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeENbID(enbID *e2sm_kpm_ies.EnbId) ([]byte, error) {

	enbIDCP, err := newENbID(enbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeENbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ENB_ID, unsafe.Pointer(enbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeENbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeENbID(bytes []byte) (*e2sm_kpm_ies.EnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ENB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("xerDecodeENbID() pointer decoded from XER is nil")
	}
	return decodeENbID((*C.ENB_ID_t)(unsafePtr))
}

func newENbID(enbID *e2sm_kpm_ies.EnbId) (*C.ENB_ID_t, error) {
	var pr C.ENB_ID_PR
	choiceC := [48]byte{}

	switch choice := enbID.EnbId.(type) {
	case *e2sm_kpm_ies.EnbId_MacroENbId:
		pr = C.ENB_ID_PR_macro_eNB_ID
		macroENbIDC := newBitString(choice.MacroENbId)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(macroENbIDC))))
	case *e2sm_kpm_ies.EnbId_HomeENbId:
		pr = C.ENB_ID_PR_home_eNB_ID
		homeEnbIDC := newBitString(choice.HomeENbId)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(homeEnbIDC))))
	case *e2sm_kpm_ies.EnbId_ShortMacroENbId:
		pr = C.ENB_ID_PR_short_Macro_eNB_ID
		shortMacroEnbIDC := newBitString(choice.ShortMacroENbId)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(shortMacroEnbIDC))))
	case *e2sm_kpm_ies.EnbId_LongMacroENbId:
		pr = C.ENB_ID_PR_long_Macro_eNB_ID
		longMacroEnbIDC := newBitString(choice.LongMacroENbId)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(longMacroEnbIDC))))
	default:
		return nil, fmt.Errorf("newENbID() unexpected type %T", choice)
	}

	enbIDC := C.ENB_ID_t{
		present: pr,
		choice:  choiceC,
	}

	return &enbIDC, nil
}

func decodeENbID(enbIDC *C.ENB_ID_t) (*e2sm_kpm_ies.EnbId, error) {
	result := new(e2sm_kpm_ies.EnbId)

	switch enbIDC.present {
	case C.ENB_ID_PR_macro_eNB_ID:

		enbIDstructC := newBitStringFromArray(enbIDC.choice)
		enb, err := decodeBitString(enbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeENbID() %s", err.Error())
		}

		result.EnbId = &e2sm_kpm_ies.EnbId_MacroENbId{
			MacroENbId: enb,
		}
	case C.ENB_ID_PR_home_eNB_ID:

		enbIDstructC := newBitStringFromArray(enbIDC.choice)
		enb, err := decodeBitString(enbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeENbID() %s", err.Error())
		}

		result.EnbId = &e2sm_kpm_ies.EnbId_HomeENbId{
			HomeENbId: enb,
		}
	case C.ENB_ID_PR_short_Macro_eNB_ID:

		enbIDstructC := newBitStringFromArray(enbIDC.choice)
		enb, err := decodeBitString(enbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeENbID() %s", err.Error())
		}

		result.EnbId = &e2sm_kpm_ies.EnbId_ShortMacroENbId{
			ShortMacroENbId: enb,
		}
	case C.ENB_ID_PR_long_Macro_eNB_ID:

		enbIDstructC := newBitStringFromArray(enbIDC.choice)
		enb, err := decodeBitString(enbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeENbID() %s", err.Error())
		}

		result.EnbId = &e2sm_kpm_ies.EnbId_LongMacroENbId{
			LongMacroENbId: enb,
		}
	default:
		return nil, fmt.Errorf("decodeENbID() %v unknown type", enbIDC.present)
	}

	return result, nil
}
