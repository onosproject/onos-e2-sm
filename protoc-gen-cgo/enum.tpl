// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package {{.PackageName}}ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "{{underscoreToDash .CstructName}}.h"
import "C"
import (
    "fmt"
    {{.ProtoFileName}} "github.com/onosproject/onos-e2-sm/servicemodels/{{cutIES .ProtoFileName}}/v1beta1/{{underscoreToDash .ProtoFileName}}" //ToDo - Make imports more dynamic
    "unsafe"
)

func xerEncode{{.MessageName}}({{lowCaseFirstLetter .MessageName}} *{{.ProtoFileName}}.{{.MessageName}}) ([]byte, error) {
    {{lowCaseFirstLetter .MessageName}}CP, err := new{{.MessageName}}({{lowCaseFirstLetter .MessageName}})
    if err != nil {
        return nil, err
    }

    bytes, err := encodeXer(&C.asn_DEF_{{dashToUnderscore .CstructName}}, unsafe.Pointer({{lowCaseFirstLetter .MessageName}}CP)) //ToDo - change name of C-encoder tag
    if err != nil {
        return nil, fmt.Errorf("xerEncode{{.MessageName}}() %s", err.Error())
    }
    return bytes, nil
}

func perEncode{{.MessageName}}({{lowCaseFirstLetter .MessageName}} *{{.ProtoFileName}}.{{.MessageName}}) ([]byte, error) {
    {{lowCaseFirstLetter .MessageName}}CP, err := new{{.MessageName}}({{lowCaseFirstLetter .MessageName}})
    if err != nil {
        return nil, err
    }

    bytes, err := encodePerBuffer(&C.asn_DEF_{{dashToUnderscore .CstructName}}, unsafe.Pointer({{lowCaseFirstLetter .MessageName}}CP)) //ToDo - change name of C-encoder tag
    if err != nil {
        return nil, fmt.Errorf("perEncode{{.MessageName}}() %s", err.Error())
    }
    return bytes, nil
}

func xerDecode{{.MessageName}}(bytes []byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
    unsafePtr, err := decodeXer(bytes, &C.asn_DEF_{{dashToUnderscore .CstructName}}) //ToDo - change name of C-encoder struct
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
        return nil, fmt.Errorf("pointer decoded from XER is nil")
    }
    return decode{{.MessageName}}((*C.{{dashToUnderscore .CstructName}}_t)(unsafePtr)) //ToDo - change name of C-struct
}

func perDecode{{.MessageName}}(bytes []byte) (*{{.ProtoFileName}}.{{.MessageName}}, error) {
    unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_{{dashToUnderscore .CstructName}}) //ToDo - change name of C-encoder struct
    if err != nil {
        return nil, err
    }
    if unsafePtr == nil {
       return nil, fmt.Errorf("pointer decoded from PER is nil")
    }
    return decode{{.MessageName}}((*C.{{dashToUnderscore .CstructName}}_t)(unsafePtr)) //ToDo - change name of C-struct
}

func new{{.MessageName}}({{lowCaseFirstLetter .MessageName}} *{{.ProtoFileName}}.{{.MessageName}}) (*C.{{dashToUnderscore .CstructName}}_t, error) { //ToDo - change name of C-struct
    var ret C.{{dashToUnderscore .CstructName}}_t //ToDo: change name of C-struct
    switch {{lowCaseFirstLetter .MessageName}} { {{with .FieldList}}{{ range . }}
            case {{.ProtoFileName}}.{{.MessageName}}_{{.FieldName}}:
                ret = C.{{.CstructLeafName}} {{ end }}{{end}} //ToDo: change name of C-variable
    default:
        return 0, fmt.Errorf("unexpected {{underscoreToDash .CstructName}} %v", {{lowCaseFirstLetter .MessageName}})
    }


    return &ret, nil
}

func decode{{.MessageName}}({{lowCaseFirstLetter .MessageName}}C *C.{{dashToUnderscore .CstructName}}_t) (*{{.ProtoFileName}}.{{.MessageName}}, error) { //ToDo - change name of C-struct

            //ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
    {{lowCaseFirstLetter .MessageName}} := {{.ProtoFileName}}.{{.MessageName}}(int32(*{{lowCaseFirstLetter .MessageName}}C))

    return &{{lowCaseFirstLetter .MessageName}}, nil
}