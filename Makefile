export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

ONOS_E2_SM_VERSION := latest
ONOS_BUILD_VERSION := v0.6.7
ONOS_PROTOC_VERSION := v0.6.7
BUF_VERSION := 0.36.0

build/_output/e2sm_kpm.so.1.0.0: # @HELP build the e2sm_kpm.so.1.0.0
	cd servicemodels/e2sm_kpm && CGO_ENABLED=1 go build -o build/_output/e2sm_kpm.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_kpm_v2.so.1.0.0: # @HELP build the e2sm_kpm_v2.so.1.0.0
	cd servicemodels/e2sm_kpm_v2 && CGO_ENABLED=1 go build -o build/_output/e2sm_kpm_v2.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_ni.so.1.0.0: # @HELP build the e2sm_ni.so.1.0.1
	cd servicemodels/e2sm_ni && CGO_ENABLED=1 go build -o build/_output/e2sm_ni.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_rc_pre.so.1.0.0: # @HELP build the e2sm_rc_pre.so.1.0.1
	cd servicemodels/e2sm_rc_pre && CGO_ENABLED=1 go build -o build/_output/e2sm_rc_pre.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_mho.so.1.0.0: # @HELP build the e2sm_mho.so.1.0.1
	cd servicemodels/e2sm_mho && CGO_ENABLED=1 go build -o build/_output/e2sm_mho.so.1.0.0 -buildmode=plugin .

PHONY:build
build: # @HELP build all libraries
build: build/_output/e2sm_kpm.so.1.0.0 build/_output/e2sm_kpm_v2.so.1.0.0 build/_output/e2sm_ni.so.1.0.0 build/_output/e2sm_rc_pre.so.1.0.0 build/_output/e2sm_mho.so.1.0.0

build_protoc_gen_cgo:
	cd protoc-gen-cgo/ && go build -v -o ./protoc-gen-cgo && cd ..

test: # @HELP run the unit tests and source code validation
test: license_check build build_protoc_gen_cgo linters
	cd servicemodels/e2sm_kpm && GODEBUG=cgocheck=0 go test -race ./...
	cd servicemodels/e2sm_rc_pre && GODEBUG=cgocheck=0 go test -race ./...
	cd servicemodels/e2sm_kpm_v2 && GODEBUG=cgocheck=0 go test -race ./...
	cd servicemodels/e2sm_mho && GODEBUG=cgocheck=0 go test -race ./...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: build-tools license_check linters
	cd servicemodels/e2sm_kpm && GODEBUG=cgocheck=0 TEST_PACKAGES=./... ./../../../build-tools/build/jenkins/make-unit
	cd servicemodels/e2sm_kpm_v2 && GODEBUG=cgocheck=0 TEST_PACKAGES=./... ./../../../build-tools/build/jenkins/make-unit
	cd servicemodels/e2sm_rc_pre && GODEBUG=cgocheck=0 TEST_PACKAGES=./... ./../../../build-tools/build/jenkins/make-unit
	cd servicemodels/e2sm_mho && GODEBUG=cgocheck=0 TEST_PACKAGES=./... ./../../../build-tools/build/jenkins/make-unit

deps_kpm: # @HELP ensure that the required dependencies are in place
	cd servicemodels/e2sm_kpm
	go build -v -buildmode=plugin ./modelmain.go
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

deps_rc: # @HELP ensure that the required dependencies are in place
	cd servicemodels/e2sm_rc_pre
	go build -v -buildmode=plugin ./modelmain.go
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

deps_mho: # @HELP ensure that the required dependencies are in place
	cd servicemodels/e2sm_mho && go build -v -buildmode=plugin ./modelmain.go && bash -c "diff -u <(echo -n) <(git diff go.mod)" && bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: golang-ci # @HELP examines Go source code and reports coding problems
	cd servicemodels/e2sm_kpm && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_kpm_v2 && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_ni && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_rc_pre && golangci-lint run --timeout 5m && cd ..
	cd servicemodels/e2sm_mho && golangci-lint run --timeout 5m && cd ..
	cd protoc-gen-cgo/ && golangci-lint run --timeout 5m && cd ..

build-tools: # @HELP install the ONOS build tools if needed
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi

jenkins-tools: # @HELP installs tooling needed for Jenkins
	cd .. && go get -u github.com/jstemmer/go-junit-report && go get github.com/t-yuki/gocover-cobertura

golang-ci: # @HELP install golang-ci if not present
	golangci-lint --version || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b `go env GOPATH`/bin v1.36.0

license_check: build-tools # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR} --boilerplate LicenseRef-ONF-Member-1.0

buflint: #@HELP run the "buf check lint" command on the proto files in 'api'
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2-sm \
		-w /go/src/github.com/onosproject/onos-e2-sm/servicemodels \
		bufbuild/buf:${BUF_VERSION} lint

protos: # @HELP compile the protobuf files (using protoc-go Docker)
protos: buflint
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2-sm \
		-w /go/src/github.com/onosproject/onos-e2-sm \
		--entrypoint /go/src/github.com/onosproject/onos-e2-sm/build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

PHONY: service-model-docker-e2sm_kpm-1.0.0
service-model-docker-e2sm_kpm-1.0.0: # @HELP build e2sm_kpm 1.0.0 plugin Docker image
	@cd servicemodels/e2sm_kpm && go mod vendor && cd ../..
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2sm_kpm \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/service-model-docker-e2sm_kpm-1.0.0:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

PHONY: service-model-docker-e2sm_kpm_v2-1.0.0
service-model-docker-e2sm_kpm_v2-1.0.0: # @HELP build e2sm_kpm_v2 1.0.0 plugin Docker image
	@cd servicemodels/e2sm_kpm_v2 && go mod vendor && cd ../..
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2sm_kpm_v2 \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/service-model-docker-e2sm_kpm_v2-1.0.0:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

PHONY: service-model-ransim-docker-e2sm_kpm-1.0.0
service-model-ransim-docker-e2sm_kpm-1.0.0: # @HELP build e2sm_kpm 1.0.0 plugin Docker image for RAN Simulator
	@cd servicemodels/e2sm_kpm && go mod vendor && cd ../..
	docker build . -f build/plugins/ransim.Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2sm_kpm \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/service-model-ransim-e2sm_kpm-1.0.0:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

PHONY: service-model-ransim-docker-e2sm_kpm_v2-1.0.0
service-model-ransim-docker-e2sm_kpm_v2-1.0.0: # @HELP build e2sm_kpm_v2 1.0.0 plugin Docker image for RAN Simulator
	@cd servicemodels/e2sm_kpm_v2 && go mod vendor && cd ../..
	docker build . -f build/plugins/ransim.Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2sm_kpm_v2\
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/service-model-ransim-e2sm_kpm_v2-1.0.0:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

PHONY: service-model-docker-e2sm_ni-1.0.0
service-model-docker-e2sm_ni-1.0.0: # @HELP build e2sm_ni 1.0.0 plugin Docker image
	@cd servicemodels/e2sm_ni && go mod vendor && cd ../..
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2sm_ni \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/service-model-docker-e2sm_ni-1.0.0:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

PHONY: service-model-docker-e2sm_rc_pre-1.0.0
service-model-docker-e2sm_rc_pre-1.0.0: # @HELP build e2sm_rc_pre 1.0.0 plugin Docker image
	@cd servicemodels/e2sm_rc_pre && go mod vendor && cd ../..
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2sm_rc_pre \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/service-model-docker-e2sm_rc_pre-1.0.0:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

PHONY: service-model-ransim-docker-e2sm_rc_pre-1.0.0
service-model-ransim-docker-e2sm_rc_pre-1.0.0: # @HELP build e2sm_rc_pre 1.0.0 plugin Docker image for RAN Simulator
	@cd servicemodels/e2sm_rc_pre && go mod vendor && cd ../..
	docker build . -f build/plugins/ransim.Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2sm_rc_pre \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/service-model-ransim-e2sm_rc_pre-1.0.0:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

PHONY: service-model-docker-e2sm_mho-1.0.0
service-model-docker-e2sm_mho-1.0.0: # @HELP build e2sm_mho 1.0.0 plugin Docker image
	@cd servicemodels/e2sm_mho && go mod vendor && cd ../..
	docker build . -f build/plugins/Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2sm_mho \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/service-model-docker-e2sm_mho-1.0.0:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

PHONY: service-model-ransim-docker-e2sm_mho-1.0.0
service-model-ransim-docker-e2sm_mho-1.0.0: # @HELP build e2sm_mho 1.0.0 plugin Docker image for RAN Simulator
	@cd servicemodels/e2sm_mho && go mod vendor && cd ../..
	docker build . -f build/plugins/ransim.Dockerfile \
		--build-arg PLUGIN_MAKE_TARGET=e2sm_mho \
		--build-arg PLUGIN_MAKE_VERSION=1.0.0 \
		--build-arg PLUGIN_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		-t onosproject/service-model-ransim-e2sm_mho-1.0.0:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

images: # @HELP build all Docker images
images: build service-model-docker-e2sm_kpm-1.0.0 service-model-ransim-docker-e2sm_kpm-1.0.0 \
	service-model-docker-e2sm_kpm_v2-1.0.0 service-model-ransim-docker-e2sm_kpm_v2-1.0.0 \
	service-model-docker-e2sm_ni-1.0.0 service-model-docker-e2sm_rc_pre-1.0.0 \
	service-model-ransim-docker-e2sm_rc_pre-1.0.0 service-model-docker-e2sm_mho-1.0.0 \
	service-model-ransim-e2sm_mho-1.0.0

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/service-model-docker-e2sm_kpm-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-ransim-e2sm_kpm-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-docker-e2sm_kpm_v2-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-ransim-e2sm_kpm_v2-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-docker-e2sm_ni-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-docker-e2sm_rc_pre-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-ransim-e2sm_rc_pre-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-docker-mho-1.0.0:${ONOS_E2_SM_VERSION}
	kind load docker-image onosproject/service-model-ransim-mho-1.0.0:${ONOS_E2_SM_VERSION}


all: build images

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version servicemodels/e2sm_kpm/${VERSION} onosproject/service-model-docker-e2sm_kpm-1.0.0 onosproject/service-model-ransim-e2sm_kpm-1.0.0
	./../build-tools/publish-version servicemodels/e2sm_kpm_v2/${VERSION} onosproject/service-model-docker-e2sm_kpm_v2-1.0.0 onosproject/service-model-ransim-e2sm_kpm_v2-1.0.0
	./../build-tools/publish-version servicemodels/e2sm_ni/${VERSION} onosproject/service-model-docker-e2sm_ni-1.0.0
	./../build-tools/publish-version servicemodels/e2sm_rc_pre/${VERSION} onosproject/service-model-docker-e2sm_rc_pre-1.0.0 onosproject/service-model-ransim-e2sm_rc_pre-1.0.0 servicemodels/mho/${VERSION} onosproject/service-model-docker-mho-1.0.0 onosproject/service-model-ransim-mho-1.0.0
	./../build-tools/publish-version


jenkins-publish: build-tools jenkins-tools # @HELP Jenkins calls this to publish artifacts
	./build/bin/push-images
	../build-tools/release-merge-commit

bumponosdeps: # @HELP update "onosproject" go dependencies and push patch to git. Add a version to dependency to make it different to $VERSION
	./../build-tools/bump-onos-deps ${VERSION}

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor ./cmd/onos-e2-sm/onos-e2-sm ./cmd/onos/onos
	go clean -testcache github.com/onosproject/onos-e2-sm/...

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
