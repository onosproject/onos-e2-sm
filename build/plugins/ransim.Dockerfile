# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

ARG PLUGIN_BUILD_VERSION=v1.0

FROM onosproject/golang-build:$PLUGIN_BUILD_VERSION as pluginbuild
ENV GO111MODULE=on
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
ARG DUMMY_FILE_NAME=mod.ran-sim-dummy
COPY Makefile ${DUMMY_FILE_NAME} /go/src/github.com/onosproject/ran-simulator/
COPY servicemodels/${PLUGIN_MAKE_TARGET}/vendor/ \
/go/src/github.com/onosproject/ran-simulator/vendor/
COPY servicemodels/${PLUGIN_MAKE_TARGET}/ \
/go/src/github.com/onosproject/ran-simulator/vendor/github.com/onosproject/onos-e2-sm/servicemodels/${PLUGIN_MAKE_TARGET}/

RUN cd /go/src/github.com/onosproject/ran-simulator/ && mv ${DUMMY_FILE_NAME} go.mod
RUN cd /go/src/github.com/onosproject/ran-simulator && \
CGO_ENABLED=1 go build -o build/_output/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} \
-buildmode=plugin github.com/onosproject/onos-e2-sm/servicemodels/${PLUGIN_MAKE_TARGET}

FROM alpine:3.20
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
COPY --from=pluginbuild /go/src/github.com/onosproject/ran-simulator/build/_output/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} /

