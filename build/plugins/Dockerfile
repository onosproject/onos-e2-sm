FROM onosproject/golang-build:v0.6.7 as pluginbuild
ENV GO111MODULE=on
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION

COPY build/_input/${PLUGIN_MAKE_TARGET} /build

WORKDIR /build

RUN GOFLAGS=-mod=vendor CGO_ENABLED=1 go build -o ${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} \
    -buildmode=plugin github.com/onosproject/onos-e2-sm/servicemodels/${PLUGIN_MAKE_TARGET}

FROM alpine:3.12
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
COPY --from=pluginbuild /build/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} /

