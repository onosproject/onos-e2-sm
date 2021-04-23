#!/bin/sh

proto_imports=".:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src"

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_kpm/v1beta1/e2sm_kpm_ies.proto

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_kpm_v2/v2/e2sm_kpm_v2.proto

#protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
#  --go_out=. \
#  e2sm_ni/v1beta1/e2sm_ni_ies.proto

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_rc_pre/v1/e2sm_rc_pre_ies.proto

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_rc_pre/v2/e2sm_rc_pre_v2.proto

cp -r github.com/onosproject/onos-e2-sm/* .
rm -rf github.com
