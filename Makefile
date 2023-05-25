# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

export CGO_ENABLED=1
export GO111MODULE=on

E2T_MOD ?= github.com/onosproject/onos-e2t@master

ONOS_E2_SM_VERSION := latest
ONOS_BUILD_VERSION := v1.0
ONOS_PROTOC_VERSION := v1.0.2

BUF_VERSION := 1.0.0

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

PHONY:build
build: # @HELP build all libraries
build: build/_output/e2sm_kpm.so.1.0.0 build/_output/e2sm_kpm_v2.so.1.0.0 build/_output/e2sm_kpm_v2_go.so.1.0.0 build/_output/e2sm_ni.so.1.0.0 build/_output/e2sm_rc_pre.so.1.0.0 build/_output/e2sm_mho.so.1.0.0 build/_output/e2sm_rsm.so.1.0.0 build/_output/e2sm_rc_pre_go.so.1.0.0 build/_output/e2sm_mho_go.so.1.0.0 build/_output/e2sm_rc.so.1.0.0

build/_output/e2sm_kpm.so.1.0.0: # @HELP build the e2sm_kpm.so.1.0.0
	cd servicemodels/e2sm_kpm && CGO_ENABLED=1 go build -o build/_output/e2sm_kpm.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_kpm_v2.so.1.0.0: # @HELP build the e2sm_kpm_v2.so.1.0.0
	cd servicemodels/e2sm_kpm_v2 && CGO_ENABLED=1 go build -o build/_output/e2sm_kpm_v2.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_kpm_v2_go.so.1.0.0: # @HELP build the e2sm_kpm_v2.so.1.0.0
	cd servicemodels/e2sm_kpm_v2_go && go build -o build/_output/e2sm_kpm_v2_go.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_rc_pre_go.so.1.0.0: # @HELP build the e2sm_rc_pre_go.so.1.0.0
	cd servicemodels/e2sm_rc_pre_go && go build -o build/_output/e2sm_rc_pre_go.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_mho_go.so.1.0.0: # @HELP build the e2sm_mho_go.so.1.0.0
	cd servicemodels/e2sm_mho_go && go build -o build/_output/e2sm_mho_go.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_rsm.so.1.0.0: # @HELP build the e2sm_rsm.so.1.0.0
	cd servicemodels/e2sm_rsm && go build -o build/_output/e2sm_rsm.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_ni.so.1.0.0: # @HELP build the e2sm_ni.so.1.0.1
	cd servicemodels/e2sm_ni && CGO_ENABLED=1 go build -o build/_output/e2sm_ni.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_rc_pre.so.1.0.0: # @HELP build the e2sm_rc_pre.so.1.0.1
	cd servicemodels/e2sm_rc_pre && CGO_ENABLED=1 go build -o build/_output/e2sm_rc_pre.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_mho.so.1.0.0: # @HELP build the e2sm_mho.so.1.0.1
	cd servicemodels/e2sm_mho && CGO_ENABLED=1 go build -o build/_output/e2sm_mho.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_rc.so.1.0.0: # @HELP build the e2sm_rc.so.1.0.1
	cd servicemodels/e2sm_rc && CGO_ENABLED=1 go build -o build/_output/e2sm_rc.so.1.0.0 -buildmode=plugin .

build_protoc_gen_cgo:
	cd protoc-gen-cgo/ && go build -v -o ./protoc-gen-cgo && cd ..

build_protoc_gen_choice:
	cd protoc-gen-choice/ && go build -v -o ./protoc-gen-choice && go install && cd ..

build_protoc_gen_builder:
	cd protoc-gen-builder/ && go build -v -o ./protoc-gen-builder && go install && cd ..

test: # @HELP run the unit tests and source code validation
test: license build build_protoc_gen_cgo build_protoc_gen_choice build_protoc_gen_builder
	cd servicemodels/e2sm_rc && go test -race ./...
	cd servicemodels/e2sm_kpm && GODEBUG=cgocheck=0 go test -race ./...
	cd servicemodels/e2sm_rc_pre && GODEBUG=cgocheck=0 go test -race ./...
	cd servicemodels/e2sm_rc_pre_go && go test -race ./...
	#cd servicemodels/e2sm_kpm_v2 && GODEBUG=cgocheck=0 go test -race ./...
	cd servicemodels/e2sm_kpm_v2_go && go test -race ./...
	#cd servicemodels/e2sm_mho && GODEBUG=cgocheck=0 go test -race ./...
	cd servicemodels/e2sm_mho_go && go test -race ./...
	cd servicemodels/e2sm_rsm && go test -race ./...
	#cd servicemodels/test_sm_aper_go_lib/testsmctypes && GODEBUG=cgocheck=0 go test -race ./...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: license
#	cd servicemodels/e2sm_kpm && GODEBUG=cgocheck=0 TEST_PACKAGES=./... ./../../build/build-tools/build/jenkins/make-unit
#	cd servicemodels/e2sm_kpm_v2 && GODEBUG=cgocheck=0 TEST_PACKAGES=./... ./../../build/build-tools/build/jenkins/make-unit
#	cd servicemodels/e2sm_kpm_v2_go && TEST_PACKAGES=./... ./../../build/build-tools/build/jenkins/make-unit
#	cd servicemodels/e2sm_rc_pre && GODEBUG=cgocheck=0 TEST_PACKAGES=./... ./../../build/build-tools/build/jenkins/make-unit
#	cd servicemodels/e2sm_rc_pre_go && TEST_PACKAGES=./... ./../../build/build-tools/build/jenkins/make-unit
#	cd servicemodels/e2sm_mho && GODEBUG=cgocheck=0 TEST_PACKAGES=./... ./../../build/build-tools/build/jenkins/make-unit
#	cd servicemodels/e2sm_mho_go && TEST_PACKAGES=./... ./../../build/build-tools/build/jenkins/make-unit
#	cd servicemodels/e2sm_rc && TEST_PACKAGES=./... ./../../build/build-tools/build/jenkins/make-unit

sm-linters: golang-ci # @HELP examines Go source code and reports coding problems
	cd servicemodels/e2sm_kpm && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_rc && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_kpm_v2 && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_kpm_v2_go && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_ni && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_rc_pre && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_rc_pre_go && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_mho && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_rsm && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/test_sm_aper_go_lib && golangci-lint run --timeout 5m && cd ..
	cd protoc-gen-cgo/ && golangci-lint run --timeout 5m && cd ..
	cd protoc-gen-choice/ && golangci-lint run --timeout 5m && cd ..
	cd protoc-gen-builder/ && golangci-lint run --timeout 5m && cd ..

buflint: #@HELP run the "buf check lint" command on the proto files in 'api'
	docker run -it \
		-v `pwd`:/go/src/github.com/onosproject/onos-e2-sm \
		-v `pwd`/../onos-lib-go/api/asn1:/go/src/github.com/onosproject/onos-e2-sm/servicemodels/asn1 \
		-w /go/src/github.com/onosproject/onos-e2-sm/servicemodels \
		bufbuild/buf:${BUF_VERSION} lint

protos: # @HELP compile the protobuf files (using protoc-go Docker)
protos: buflint
	docker run -it \
		-v `pwd`:/go/src/github.com/onosproject/onos-e2-sm \
		-v `pwd`/../onos-lib-go:/go/src/github.com/onosproject/onos-lib-go \
		-w /go/src/github.com/onosproject/onos-e2-sm \
		--entrypoint /go/src/github.com/onosproject/onos-e2-sm/build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

protos-py: # @HELP compile the protobuf files for python (using protoc-go Docker)
protos-py:
	docker run -it \
		-v `pwd`:/go/src/github.com/onosproject/onos-e2-sm \
		-v `pwd`/../onos-lib-go:/go/src/github.com/onosproject/onos-lib-go \
		-w /go/src/github.com/onosproject/onos-e2-sm \
		--entrypoint /go/src/github.com/onosproject/onos-e2-sm/build/bin/compile-protos-py.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

PHONY: service-model-docker-e2sm_kpm-1.0.0
service-model-docker-e2sm_kpm-1.0.0: # @HELP build e2sm_kpm 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_kpm ${E2T_MOD}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_kpm" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_kpm-1.0.0:${ONOS_E2_SM_VERSION}

PHONY: service-model-docker-e2sm_kpm_v2-1.0.0
service-model-docker-e2sm_kpm_v2-1.0.0: # @HELP build e2sm_kpm_v2 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_kpm_v2 ${E2T_MOD} onosproject/service-model-docker-e2sm_kpm_v2-1.0.0:${ONOS_E2_SM_VERSION}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_kpm_v2" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_kpm_v2-1.0.0:${ONOS_E2_SM_VERSION}

PHONY: service-model-docker-e2sm_kpm_v2_go-1.0.0
service-model-docker-e2sm_kpm_v2_go-1.0.0: # @HELP build e2sm_kpm_v2 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_kpm_v2_go ${E2T_MOD} onosproject/service-model-docker-e2sm_kpm_v2_go-1.0.0:${ONOS_E2_SM_VERSION}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_kpm_v2_go" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_kpm_v2_go-1.0.0:${ONOS_E2_SM_VERSION}

PHONY: service-model-docker-e2sm_rsm-1.0.0
service-model-docker-e2sm_rsm-1.0.0: # @HELP build e2sm_kpm_v2 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_rsm ${E2T_MOD} onosproject/service-model-docker-e2sm_rsm-1.0.0:${ONOS_E2_SM_VERSION}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_rsm" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_rsm-1.0.0:${ONOS_E2_SM_VERSION}

PHONY: service-model-docker-e2sm_ni-1.0.0
service-model-docker-e2sm_ni-1.0.0: # @HELP build e2sm_ni 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_ni ${E2T_MOD}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_ni" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_ni-1.0.0:${ONOS_E2_SM_VERSION}

PHONY: service-model-docker-e2sm_rc_pre-1.0.0
service-model-docker-e2sm_rc_pre-1.0.0: # @HELP build e2sm_rc_pre 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_rc_pre ${E2T_MOD}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_rc_pre" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_rc_pre-1.0.0:${ONOS_E2_SM_VERSION}

PHONY: service-model-docker-e2sm_rc_pre_go-1.0.0
service-model-docker-e2sm_rc_pre_go-1.0.0: # @HELP build e2sm_rc_pre_go 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_rc_pre_go ${E2T_MOD} onosproject/service-model-docker-e2sm_rc_pre_go-1.0.0:${ONOS_E2_SM_VERSION}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_rc_pre_go" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_rc_pre_go-1.0.0:${ONOS_E2_SM_VERSION}

PHONY: service-model-docker-e2sm_mho-1.0.0
service-model-docker-e2sm_mho-1.0.0: # @HELP build e2sm_mho 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_mho ${E2T_MOD}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_mho" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_mho-1.0.0:${ONOS_E2_SM_VERSION}


PHONY: service-model-docker-e2sm_mho_go-1.0.0
service-model-docker-e2sm_mho_go-1.0.0: # @HELP build e2sm_mho_go 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_mho_go ${E2T_MOD} onosproject/service-model-docker-e2sm_mho_go-1.0.0:${ONOS_E2_SM_VERSION}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_mho_go" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_mho_go-1.0.0:${ONOS_E2_SM_VERSION}

PHONY: service-model-docker-e2sm_rc-1.0.0
service-model-docker-e2sm_rc-1.0.0: # @HELP build e2sm_rc_pre_go 1.0.0 plugin Docker image
	./build/bin/build-deps e2sm_rc ${E2T_MOD} onosproject/service-model-docker-e2sm_rc-1.0.0:${ONOS_E2_SM_VERSION}
	docker build . -f build/plugins/Dockerfile \
			--build-arg PLUGIN_MAKE_TARGET="e2sm_rc" \
			--build-arg PLUGIN_MAKE_VERSION="1.0.0" \
			-t onosproject/service-model-docker-e2sm_rc-1.0.0:${ONOS_E2_SM_VERSION}

images: # @HELP build all Docker images
images: build service-model-docker-e2sm_kpm_v2_go-1.0.0 \
	service-model-docker-e2sm_rsm-1.0.0 \
	service-model-docker-e2sm_rc_pre_go-1.0.0 \
	service-model-docker-e2sm_mho_go-1.0.0 \
	service-model-docker-e2sm_rc-1.0.0


kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/service-model-docker-e2sm_kpm_v2_go-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-docker-e2sm_rsm-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-docker-e2sm_rc_pre_go-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-docker-e2sm_mho_go-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-docker-e2sm_rc-1.0.0:${ONOS_E2_SM_VERSION}




all: build images

publish: # @HELP publish version on github and dockerhub
	./build/build-tools/publish-version servicemodels/e2sm_kpm_v2_go/${VERSION} onosproject/service-model-docker-e2sm_kpm_v2_go-1.0.0
	./build/build-tools/publish-version servicemodels/e2sm_rc_pre_go/${VERSION} onosproject/service-model-docker-e2sm_rc_pre_go-1.0.0
	./build/build-tools/publish-version servicemodels/e2sm_mho_go/${VERSION} onosproject/service-model-docker-e2sm_mho_go-1.0.0
	./build/build-tools/publish-version servicemodels/e2sm_rsm/${VERSION} onosproject/service-model-docker-e2sm_rsm-1.0.0
	./build/build-tools/publish-version servicemodels/e2sm_rc/${VERSION} onosproject/service-model-docker-e2sm_rc-1.0.0


jenkins-publish: # @HELP Jenkins calls this to publish artifacts
	./build/bin/push-images
	./build/build-tools/release-merge-commit

clean:: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor ./cmd/onos-e2-sm/onos-e2-sm ./cmd/onos/onos
	rm -fr servicemodels/*/vendor
	go clean -testcache github.com/onosproject/onos-e2-sm/...

