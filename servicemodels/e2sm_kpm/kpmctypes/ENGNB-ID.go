// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ENGNB-ID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeENGNbID(enGNbID *e2sm_kpm_ies.EngnbId) ([]byte, error) {

	enGNbIDCP, err := newENGNbID(enGNbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeENGNbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ENGNB_ID, unsafe.Pointer(enGNbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeENGNbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeENGNbID(bytes []byte) (*e2sm_kpm_ies.EngnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ENGNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("xerDecodeENGNbID() pointer decoded from XER is nil")
	}
	return decodeENGNbID((*C.ENGNB_ID_t)(unsafePtr))
}

func newENGNbID(enGNbID *e2sm_kpm_ies.EngnbId) (*C.ENGNB_ID_t, error) {
	var pr C.ENGNB_ID_PR
	choiceC := [48]byte{}

	switch choice := enGNbID.EngnbId.(type) {
	case *e2sm_kpm_ies.EngnbId_GNbId:
		pr = C.ENGNB_ID_PR_gNB_ID
		enGNbIDC := newBitString(choice.GNbId)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(enGNbIDC))))
	default:
		return nil, fmt.Errorf("newENGNbID() unexpected type %T", choice)
	}

	enGNbIDC := C.ENGNB_ID_t{
		present: pr,
		choice:  choiceC,
	}

	return &enGNbIDC, nil
}

func decodeENGNbID(enGNbIDC *C.ENGNB_ID_t) (*e2sm_kpm_ies.EngnbId, error) {
	result := new(e2sm_kpm_ies.EngnbId)

	switch enGNbIDC.present {
	case C.ENGNB_ID_PR_gNB_ID:
		enGNbIDstructC := newBitStringFromArray(enGNbIDC.choice)
		enGNb, err := decodeBitString(enGNbIDstructC)
		if err != nil {
			return nil, fmt.Errorf("decodeENbID() %s", err.Error())
		}

		result.EngnbId = &e2sm_kpm_ies.EngnbId_GNbId{
			GNbId: enGNb,
		}
	default:
		return nil, fmt.Errorf("decodeENGNbID() %v unknown type", enGNbIDC.present)
	}

	return result, nil
}
