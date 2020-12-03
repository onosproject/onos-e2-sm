ARG PLUGIN_BUILD_VERSION=stable

FROM onosproject/golang-build:$PLUGIN_BUILD_VERSION as pluginbuild
ENV GO111MODULE=on
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
COPY Makefile servicemodels/${PLUGIN_MAKE_TARGET} /go/src/github.com/onosproject/onos-e2-sm/

RUN cd /go/src/github.com/onosproject/onos-e2-sm && GOFLAGS=-mod=vendor make build/_output/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION}

FROM alpine:3.12
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
COPY --from=pluginbuild /go/src/github.com/onosproject/onos-e2-sm/build/_output/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} /

