// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmctype

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "{{underscoreToDash .CstructName}}.h"
import "C"
import (
    "fmt"
    e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies" //ToDo: Make imports more dynamic
    "unsafe"
)

func xerEncode{{.MessageName}}({{lowCaseFirstLetter .MessageName}} *{{.ProtoFileName}}.{{.MessageName}}) ([]byte, error) {
    {{lowCaseFirstLetter .MessageName}}CP := new{{.MessageName}}({{lowCaseFirstLetter .MessageName}})

    bytes, err := encodeXer(&C.asn_DEF_{{dashToUnderscore .CstructName}}, unsafe.Pointer({{lowCaseFirstLetter .MessageName}}CP))
    if err != nil {
        return nil, fmt.Errorf("xerEncode{{.MessageName}}() %s", err.Error())
    }
    return bytes, nil
}

func perEncode{{.MessageName}}({{lowCaseFirstLetter .MessageName}} *{{.ProtoFileName}}.{{.MessageName}}) ([]byte, error) {
    {{lowCaseFirstLetter .MessageName}}CP := new{{.MessageName}}({{lowCaseFirstLetter .MessageName}})

    bytes, err := encodePerBuffer(&C.asn_DEF_{{dashToUnderscore .CstructName}}, unsafe.Pointer({{lowCaseFirstLetter .MessageName}}CP))
    if err != nil {
        return nil, fmt.Errorf("perEncode{{.MessageName}}() %s", err.Error())
    }
    return bytes, nil
}

func xerDecode{{.MessageName}}(bytes []byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
    unsafePtr, err := decodeXer(bytes, &C.asn_DEF_{{dashToUnderscore .CstructName}})
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
        return nil, fmt.Errorf("pointer decoded from XER is nil")
    }
    return decode{{.MessageName}}((*C.{{dashToUnderscore .CstructName}}_t)(unsafePtr)), nil
}

func perDecode{{.MessageName}}(bytes []byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
    unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_{{dashToUnderscore .CstructName}})
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
       return nil, fmt.Errorf("pointer decoded from PER is nil")
    }
    return decode{{.MessageName}}((*C.{{dashToUnderscore .CstructName}}_t)(unsafePtr)), nil
}

func new{{.MessageName}}({{lowCaseFirstLetter .MessageName}} *{{.ProtoFileName}}.{{.MessageName}}) *C.{{dashToUnderscore .CstructName}}_t {

    //TODO: automate picking of C.type data type
    {{lowCaseFirstLetter .MessageName}}C := C.long({{lowCaseFirstLetter .MessageName}}.{{.FieldName}})

    return &{{lowCaseFirstLetter .MessageName}}C
}

func decode{{.MessageName}}({{lowCaseFirstLetter .MessageName}}C *C.{{dashToUnderscore .CstructName}}_t) *{{.ProtoFileName}}.{{.MessageName}} {

    {{lowCaseFirstLetter .MessageName}} := {{.ProtoFileName}}.{{.MessageName}}{
        {{upperCaseFirstLetter .FieldTypeName}}: {{convertDataType .DataType}}(*{{lowCaseFirstLetter .MessageName}}C),
    }
    return &{{lowCaseFirstLetter .MessageName}}
}