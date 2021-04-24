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
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
)

//func xerEncodeRcPreCommand(rcPreCommand *e2sm_rc_pre_v2.RcPreCommand) ([]byte, error) {
//	rcPreCommandCP, err := newRcPreCommand(rcPreCommand)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RC_PRE_Command, unsafe.Pointer(rcPreCommandCP)) //ToDo - change name of C-encoder tag
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRcPreCommand() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeRcPreCommand(rcPreCommand *e2sm_rc_pre_v2.RcPreCommand) ([]byte, error) {
//	rcPreCommandCP, err := newRcPreCommand(rcPreCommand)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_RC_PRE_Command, unsafe.Pointer(rcPreCommandCP)) //ToDo - change name of C-encoder tag
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeRcPreCommand() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerDecodeRcPreCommand(bytes []byte) (*e2sm_rc_pre_v2.RcPreCommand, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RC_PRE_Command)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeRcPreCommand((*C.RC_PRE_Command_t)(unsafePtr)) //ToDo - change name of C-struct
//}
//
//func perDecodeRcPreCommand(bytes []byte) (*e2sm_rc_pre_v2.RcPreCommand, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RC_PRE_Command)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRcPreCommand((*C.RC_PRE_Command_t)(unsafePtr))
//}

func newMhoCommand(mhoCommand *e2sm_mho.MhoCommand) (*C.MHO_Command_t, error) { //ToDo - change name of C-struct
	var ret C.MHO_Command_t
	// shad TODO
	//switch *mhoCommand {
	//case e2sm_mho.RcPreCommand_RC_PRE_COMMAND_SET_PARAMETERS:
	//	ret = C.RC_PRE_Command_setParameters
	//default:
	//	return nil, fmt.Errorf("unexpected RC-PRE-Command %v", rcPreCommand)
	//}

	return &ret, nil
}

func decodeMhoCommand(mhoCommandC *C.MHO_Command_t) (*e2sm_mho.MhoCommand, error) { //ToDo - change name of C-struct

	//ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
	mhoCommand := e2sm_mho.MhoCommand(int32(*mhoCommandC))

	return &mhoCommand, nil
}
