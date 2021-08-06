FROM onosproject/golang-build:v0.6.7 as pluginbuild
ENV GO111MODULE=on
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
COPY Makefile mod.dummy /go/src/github.com/sdran/onos-e2-sm/
COPY servicemodels/${PLUGIN_MAKE_TARGET}/vendor/ \
/go/src/github.com/sdran/onos-e2-sm/vendor/
COPY servicemodels/${PLUGIN_MAKE_TARGET}/ \
/go/src/github.com/sdran/onos-e2-sm/vendor/github.com/sdran/onos-e2-sm/servicemodels/${PLUGIN_MAKE_TARGET}/

RUN cd /go/src/github.com/sdran/onos-e2-sm/ && mv mod.dummy go.mod
RUN cd /go/src/github.com/sdran/onos-e2-sm && \
CGO_ENABLED=1 go build -o build/_output/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} \
-buildmode=plugin github.com/sdran/onos-e2-sm/servicemodels/${PLUGIN_MAKE_TARGET}

FROM alpine:3.12
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
COPY --from=pluginbuild /go/src/github.com/sdran/onos-e2-sm/build/_output/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} /

