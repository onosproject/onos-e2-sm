// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MHO-Command.h"
import "C"
import (
    "encoding/binary"
	"fmt"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho" //ToDo - Make imports more dynamic
	"unsafe"
)

func xerEncodeMhoCommand(mhoCommand *e2sm_mho.MhoCommand) ([]byte, error) {
    mhoCommandCP, err := newMhoCommand(mhoCommand)
    if err != nil {
        return nil, err
    }
    bytes, err := encodeXer(&C.asn_DEF_MHO_Command, unsafe.Pointer(mhoCommandCP)) //ToDo - change name of C-encoder tag
    if err != nil {
        return nil, fmt.Errorf("xerEncodeMhoCommand() %s", err.Error())
    }
    return bytes, nil
}

func perEncodeMhoCommand(mhoCommand *e2sm_mho.MhoCommand) ([]byte, error) {
    mhoCommandCP, err := newMhoCommand(mhoCommand)
    if err != nil {
        return nil, err
    }

    bytes, err := encodePerBuffer(&C.asn_DEF_MHO_Command, unsafe.Pointer(mhoCommandCP))
    if err != nil {
        return nil, fmt.Errorf("perEncodeMhoCommand() %s", err.Error())
    }
    return bytes, nil
}

func xerDecodeMhoCommand(bytes []byte) (*e2sm_mho.MhoCommand, error) {
    unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MHO_Command)
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
        return nil, fmt.Errorf("pointer decoded from XER is nil")
    }
    return decodeMhoCommand((*C.MHO_Command_t)(unsafePtr)) //ToDo - change name of C-struct
}

func perDecodeMhoCommand(bytes []byte) (*e2sm_mho.MhoCommand, error) {
    unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MHO_Command)
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
        return nil, fmt.Errorf("pointer decoded from PER is nil")
    }
    return decodeMhoCommand((*C.MHO_Command_t)(unsafePtr))
}

func newMhoCommand(mhoCommand *e2sm_mho.MhoCommand) (*C.MHO_Command_t, error) {
	var ret C.MHO_Command_t
    switch *mhoCommand {
    case e2sm_mho.MhoCommand_MHO_COMMAND_INITIATE_HANDOVER:
        ret = C.MHO_Command_initiateHandover //ToDo - double-check correctness of the name
    default:
        return nil, fmt.Errorf("unexpected MhoCommand %v", mhoCommand)
    }

	return &ret, nil
}

func decodeMhoCommand(mhoCommandC *C.MHO_Command_t) (*e2sm_mho.MhoCommand, error) {

	//ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
	mhoCommand := e2sm_mho.MhoCommand(int32(*mhoCommandC))

	return &mhoCommand, nil
}

func decodeMhoCommandBytes(array [8]byte) (*e2sm_mho.MhoCommand, error) { //ToDo - Check addressing correct structure in Protobuf
    mhoCommandC := (*C.MHO_Command_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))

    return decodeMhoCommand(mhoCommandC)
}
