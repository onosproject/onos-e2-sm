// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctypes
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
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeGnbIdChoice(gnbIdChoice *e2sm_kpm_ies.GnbIdChoice) ([]byte, error) {
	gnbIDChoiceC, err := newGnbIDChoice(gnbIdChoice)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbIdChoice() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GNB_ID_Choice, unsafe.Pointer(gnbIDChoiceC))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbIdChoice() %s", err.Error())
	}
	return bytes, nil
}

func newGnbIDChoice(gnbIDc *e2sm_kpm_ies.GnbIdChoice) (*C.GNB_ID_Choice_t, error) {
	var pr C.GNB_ID_Choice_PR
	choiceC := [48]byte{}

	switch choice := gnbIDc.GetGnbIdChoice().(type) {
	case *e2sm_kpm_ies.GnbIdChoice_GnbId:
		pr = C.GNB_ID_Choice_PR_gnb_ID
		bsC := newBitString(choice.GnbId)
		fmt.Printf("newGnbIDChoice() gNB ID %v %v %v %v\n", bsC, unsafe.Sizeof(bsC.size), unsafe.Sizeof(bsC.bits_unused), *bsC.buf)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	default:
		return nil, fmt.Errorf("newGnbIDChoice() unexpected type %T not yet implemented", choice)
	}

	gnbidcC := C.GNB_ID_Choice_t{
		present: pr,
		choice: choiceC,
	}

	return &gnbidcC, nil
}

func decodeGnbIDChoice(gnbIDC *C.GNB_ID_Choice_t) (*e2sm_kpm_ies.GnbIdChoice, error) {
	result := new(e2sm_kpm_ies.GnbIdChoice)

	switch gnbIDC.present {
	case C.GNB_ID_Choice_PR_gnb_ID:
		gnbIDstructC := newBitStringFromArray(gnbIDC.choice)

		bitString, err := decodeBitString(gnbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeGnbIDChoice() %s", err.Error())
		}
		result.GnbIdChoice = &e2sm_kpm_ies.GnbIdChoice_GnbId{
			GnbId: bitString,
		}
	default:
		return nil, fmt.Errorf("decodeGnbIDChoice() %v not yet implemented", gnbIDC.present)
	}

	return result, nil
}
