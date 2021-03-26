// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RIC-KPMNode-Item.h"
//#include "Cell-Measurement-Object-Item.h"
import "C"
import (
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"unsafe"
)

func xerEncodeRicKpmnodeItem(ricKpmnodeItem *e2sm_kpm_v2.RicKpmnodeItem) ([]byte, error) {
	ricKpmnodeItemCP, err := newRicKpmnodeItem(ricKpmnodeItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicKpmnodeItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RIC_KPMNode_Item, unsafe.Pointer(ricKpmnodeItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicKpmnodeItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicKpmnodeItem(ricKpmnodeItem *e2sm_kpm_v2.RicKpmnodeItem) ([]byte, error) {
	ricKpmnodeItemCP, err := newRicKpmnodeItem(ricKpmnodeItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicKpmnodeItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RIC_KPMNode_Item, unsafe.Pointer(ricKpmnodeItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicKpmnodeItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicKpmnodeItem(bytes []byte) (*e2sm_kpm_v2.RicKpmnodeItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RIC_KPMNode_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicKpmnodeItem((*C.RIC_KPMNode_Item_t)(unsafePtr))
}

func perDecodeRicKpmnodeItem(bytes []byte) (*e2sm_kpm_v2.RicKpmnodeItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RIC_KPMNode_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicKpmnodeItem((*C.RIC_KPMNode_Item_t)(unsafePtr))
}

func newRicKpmnodeItem(ricKpmnodeItem *e2sm_kpm_v2.RicKpmnodeItem) (*C.RIC_KPMNode_Item_t, error) {

	//fmt.Printf("We're inside newRicKpmnodeItem(), starting encoding...")
	//fmt.Printf("Encoding CellMeasurementObjectList: \n %v \n", ricKpmnodeItem.GetCellMeasurementObjectList())
	cellMeasurementObjectListC := new(C.struct_RIC_KPMNode_Item__cell_Measurement_Object_List) //ToDo - verify correctness of the variable's name
	for _, cellMeasurementObjectListItem := range ricKpmnodeItem.GetCellMeasurementObjectList() {
		//fmt.Printf("Encoding CellMeasurementObjectItem: \n %v \n", cellMeasurementObjectListItem)
		cellMeasurementObjectListItemC, err := newCellMeasurementObjectItem(cellMeasurementObjectListItem)
		//fmt.Printf("That's the C version of encoded CellMeasurementObjectItem: \n %v \n", cellMeasurementObjectListItemC)
		if err != nil {
			return nil, fmt.Errorf("newCellMeasurementObjectItem() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(cellMeasurementObjectListC), unsafe.Pointer(cellMeasurementObjectListItemC)); err != nil {
			return nil, err
		}
	}
	//fmt.Printf("That's the C version of encoded CellMeasurementObjectList: \n %v \n", cellMeasurementObjectListC)

	//fmt.Printf("Encoding GlobalKpmnodeI: \n %v \n", ricKpmnodeItem.RicKpmnodeType)
	ricKpmnodeTypeC, err := newGlobalKpmnodeID(ricKpmnodeItem.RicKpmnodeType)
	if err != nil {
		return nil, fmt.Errorf("newGlobalKpmnodeID() %s", err.Error())
	}
	//fmt.Printf("That's the C version of encoded RicKpmnodeType: \n %v \n", ricKpmnodeTypeC)

	//fmt.Printf("Composing final C-structure...")
	ricKpmnodeItemC := C.RIC_KPMNode_Item_t{
		ric_KPMNode_Type:             *ricKpmnodeTypeC,
		cell_Measurement_Object_List: cellMeasurementObjectListC,
	}
	//fmt.Printf("This is final C-structure: \n %v \n", ricKpmnodeItemC)

	return &ricKpmnodeItemC, nil
}

func decodeRicKpmnodeItem(ricKpmnodeItemC *C.RIC_KPMNode_Item_t) (*e2sm_kpm_v2.RicKpmnodeItem, error) {

	//fmt.Printf("We're inside decodeRicKpmnodeItem(), starting decoding...")
	//fmt.Printf("Decoding RicKpmnodeType: \n %v \n", ricKpmnodeItemC.ric_KPMNode_Type)
	ricKpmnodeType, err := decodeGlobalKpmnodeID(&ricKpmnodeItemC.ric_KPMNode_Type)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalKpmnodeID() %s", err.Error())
	}
	//fmt.Printf("That's what was decoded from C: \n %v \n", ricKpmnodeType)


	ricKpmnodeItem := e2sm_kpm_v2.RicKpmnodeItem{
		RicKpmnodeType:            ricKpmnodeType,
		CellMeasurementObjectList: make([]*e2sm_kpm_v2.CellMeasurementObjectItem, 0),

	}

	//fmt.Printf("Decoding CellMeasurementObjectList: \n %v \n", ricKpmnodeItemC.cell_Measurement_Object_List)
	//fmt.Printf("We've got %v items to decode \n", int(ricKpmnodeItemC.cell_Measurement_Object_List.list.count))
	ieCount := int(ricKpmnodeItemC.cell_Measurement_Object_List.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(ricKpmnodeItemC.cell_Measurement_Object_List.list.array)) * uintptr(i)
		ieC := *(**C.Cell_Measurement_Object_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ricKpmnodeItemC.cell_Measurement_Object_List.list.array)) + offset))
		//fmt.Printf("Decoding CellMeasurementObjectItem: \n %v \n", ieC)
		ie, err := decodeCellMeasurementObjectItem(ieC)
		//fmt.Printf("Decoded CellMeasurementObjectItem is: \n %v \n", ie)
		if err != nil {
			return nil, fmt.Errorf("decodeCellMeasurementObjectItem() %s", err.Error())
		}
		ricKpmnodeItem.CellMeasurementObjectList = append(ricKpmnodeItem.CellMeasurementObjectList, ie)
	}
	//fmt.Printf("Decoded CellMeasurementObjectList is: \n %v \n", ricKpmnodeItem.CellMeasurementObjectList)

	//fmt.Printf("This is final decoded structure: \n %v \n", ricKpmnodeItem)

	return &ricKpmnodeItem, nil
}
