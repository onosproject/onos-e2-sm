// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2
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
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_GNB_ID_Choice, unsafe.Pointer(gnbIDChoiceC))
	if err != nil {
		return nil, err
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
		fmt.Printf("gNB ID %v %v %v %v\n", bsC, unsafe.Sizeof(bsC.size), unsafe.Sizeof(bsC.bits_unused), *bsC.buf)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	default:
		return nil, fmt.Errorf("unexpected type %T not yet implemented", choice)
	}

	gnbidcC := C.GNB_ID_Choice_t{
		present: pr,
		choice: choiceC,
	}

	return &gnbidcC, nil
}

func decodeGnbIDChoice(gnbIDcC *C.GNB_ID_Choice_t) (*e2sm_kpm_ies.GnbIdChoice, error) {
	return nil, fmt.Errorf("decodeGnbIDChoice() not yet implemented")
}
