// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANparameterDef-Item.h"
import "C"

import (
	"fmt"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	//"unsafe"
)

//func xerEncodeRanparameterDefItem(ranparameterDefItem *e2sm_rc_pre_ies.RanparameterDefItem) ([]byte, error) {
//	ranparameterDefItemCP, err := newRanparameterDefItem(ranparameterDefItem)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterDefItem() %s", err.Error())
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RANparameterDef_Item, unsafe.Pointer(ranparameterDefItemCP))
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterDefItem() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func perEncodeRanparameterDefItem(ranparameterDefItem *e2sm_rc_pre_ies.RanparameterDefItem) ([]byte, error) {
//	ranparameterDefItemCP, err := newRanparameterDefItem(ranparameterDefItem)
//	if err != nil {
//		return nil, fmt.Errorf("xerEncodeRanparameterDefItem() %s", err.Error())
//	}
//
//	bytes, err := encodePerBuffer(&C.asn_DEF_RANparameterDef_Item, unsafe.Pointer(ranparameterDefItemCP))
//	if err != nil {
//		return nil, fmt.Errorf("perEncodeRanparameterDefItem() %s", err.Error())
//	}
//	return bytes, nil
//}
//
//func xerDecodeRanparameterDefItem(bytes []byte) (*e2sm_rc_pre_ies.RanparameterDefItem, error) {
//	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANparameterDef_Item)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from XER is nil")
//	}
//	return decodeRanparameterDefItem((*C.RANparameterDef_Item_t)(unsafePtr))
//}
//
//func perDecodeRanparameterDefItem(bytes []byte) (*e2sm_rc_pre_ies.RanparameterDefItem, error) {
//	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANparameterDef_Item)
//	if err != nil {
//		return nil, err
//	}
//	if unsafePtr == nil {
//		return nil, fmt.Errorf("pointer decoded from PER is nil")
//	}
//	return decodeRanparameterDefItem((*C.RANparameterDef_Item_t)(unsafePtr))
//}

func newRanparameterDefItem(ranparameterDefItem *e2sm_rc_pre_ies.RanparameterDefItem) (*C.RANparameterDef_Item_t, error) {

	ranparameterIDC, err := newRanparameterID(ranparameterDefItem.RanParameterId)
	if err != nil {
		return nil, fmt.Errorf("encodeRanparameterDefItem() %s", err.Error())
	}

	ranParameterNameC, err := newRanparameterName(ranparameterDefItem.RanParameterName)
	if err != nil {
		return nil, fmt.Errorf("encodeRanparameterDefItem() %s", err.Error())
	}

	ranParameterTypeC, err := newRanparameterType(&ranparameterDefItem.RanParameterType)
	if err != nil {
		return nil, fmt.Errorf("encodeRanparameterDefItem() %s", err.Error())
	}

	ranparameterDefItemC := C.RANparameterDef_Item_t{
		ranParameter_ID:   *ranparameterIDC,
		ranParameter_Name: *ranParameterNameC,
		ranParameter_Type: *ranParameterTypeC,
	}

	return &ranparameterDefItemC, nil
}

func decodeRanparameterDefItem(ranparameterDefItemC *C.RANparameterDef_Item_t) (*e2sm_rc_pre_ies.RanparameterDefItem, error) {

	ranparameterID, err := decodeRanparameterID(&ranparameterDefItemC.ranParameter_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeRanparameterDefItem() %s", err.Error())
	}

	ranParameterName, err := decodeRanparameterName(&ranparameterDefItemC.ranParameter_Name)
	if err != nil {
		return nil, fmt.Errorf("decodeRanparameterDefItem() %s", err.Error())
	}

	ranParameterType, err := decodeRanparameterType(&ranparameterDefItemC.ranParameter_Type)
	if err != nil {
		return nil, fmt.Errorf("decodeRanparameterDefItem() %s", err.Error())
	}

	ranparameterDefItem := e2sm_rc_pre_ies.RanparameterDefItem{
		RanParameterId:   ranparameterID,
		RanParameterName: ranParameterName,
		RanParameterType: *ranParameterType,
	}

	return &ranparameterDefItem, nil
}
