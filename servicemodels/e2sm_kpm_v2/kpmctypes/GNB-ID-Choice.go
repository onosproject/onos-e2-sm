// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GNB-ID-Choice.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeGnbIDChoice(gnbIDChoice *e2sm_kpm_v2.GnbIdChoice) ([]byte, error) {
	gnbIDChoiceCP, err := newGnbIDChoice(gnbIDChoice)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbIDChoice() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GNB_ID_Choice, unsafe.Pointer(gnbIDChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbIDChoice() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGnbIDChoice(gnbIDChoice *e2sm_kpm_v2.GnbIdChoice) ([]byte, error) {
	gnbIDChoiceCP, err := newGnbIDChoice(gnbIDChoice)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGnbIDChoice() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GNB_ID_Choice, unsafe.Pointer(gnbIDChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGnbIDChoice() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGnbIDChoice(bytes []byte) (*e2sm_kpm_v2.GnbIdChoice, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GNB_ID_Choice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGnbIDChoice((*C.GNB_ID_Choice_t)(unsafePtr))
}

func perDecodeGnbIDChoice(bytes []byte) (*e2sm_kpm_v2.GnbIdChoice, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GNB_ID_Choice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGnbIDChoice((*C.GNB_ID_Choice_t)(unsafePtr))
}

func newGnbIDChoice(gnbIDChoice *e2sm_kpm_v2.GnbIdChoice) (*C.GNB_ID_Choice_t, error) {

	var pr C.GNB_ID_Choice_PR
	choiceC := [48]byte{}

	switch choice := gnbIDChoice.GetGnbIdChoice().(type) {
	case *e2sm_kpm_v2.GnbIdChoice_GnbId:
		pr = C.GNB_ID_Choice_PR_gnb_ID

		bsC, err := newBitString(choice.GnbId)
		if err != nil {
			return nil, fmt.Errorf("newBitString() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	default:
		return nil, fmt.Errorf("newGnbIdChoice() %T not yet implemented", choice)
	}

	gnbIDChoiceC := C.GNB_ID_Choice_t{
		present: pr,
		choice:  choiceC,
	}

	return &gnbIDChoiceC, nil
}

func decodeGnbIDChoice(gnbIDchoiceC *C.GNB_ID_Choice_t) (*e2sm_kpm_v2.GnbIdChoice, error) {

	gnbIDchoice := new(e2sm_kpm_v2.GnbIdChoice)

	switch gnbIDchoiceC.present {
	case C.GNB_ID_Choice_PR_gnb_ID:
		gnbIDstructC := newBitStringFromArray(gnbIDchoiceC.choice)

		gnbID, err := decodeBitString(gnbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeBitStringBytes() %s", err.Error())
		}
		gnbIDchoice.GnbIdChoice = &e2sm_kpm_v2.GnbIdChoice_GnbId{
			GnbId: gnbID,
		}
	default:
		return nil, fmt.Errorf("decodeGnbIDChoice() %v not yet implemented", gnbIDchoiceC.present)
	}

	return gnbIDchoice, nil
}
