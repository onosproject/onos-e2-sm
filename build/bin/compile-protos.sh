#!/bin/sh

proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src"

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2-sm/servicemodels:. \
  e2sm_kpm/v1beta1/e2sm_kpm_ies.proto

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2-sm/servicemodels:. \
  e2sm_ni/v1beta1/e2sm_ni_ies.proto

cp -r github.com/onosproject/onos-e2-sm/* .
rm -rf github.com
