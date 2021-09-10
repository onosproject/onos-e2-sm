#!/bin/sh

# TODO: refactor all the compile-protos-*.sh and use command line paramters to distinguish

set -e

echo "** generate bindings"
proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src"

OUTPUTPATH="python"
rm -rf "$OUTPUTPATH"
mkdir -p "$OUTPUTPATH"

echo "* generate onos bindings"
cd servicemodels

# betterproto client bindings into $OUTPUTPATH
protoc "-I=$proto_imports" \
    "--python_betterproto_out=../$OUTPUTPATH" \
    e2sm_ni/v1beta1/e2sm_ni_ies.proto \
    validate/v1/validate.proto \
    e2sm_kpm_v2/v2/e2sm_kpm_v2.proto \
    e2sm_rc_pre/v2/e2sm_rc_pre_v2.proto \
    e2sm_mho/v1/e2sm_mho.proto \
    e2sm_kpm/v1beta1/e2sm_kpm_ies.proto
