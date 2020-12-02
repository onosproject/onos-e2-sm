export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

ONOS_E2_SM_VERSION := latest
ONOS_BUILD_VERSION := v0.6.6
ONOS_PROTOC_VERSION := v0.6.6
BUF_VERSION := 0.27.1

build/_output/e2sm_kpm.so.1.0.0: # @HELP build the e2sm_kpm.so.1.0.0
	cd servicemodels/e2sm_kpm && CGO_ENABLED=1 go build -o build/_output/e2sm_kpm.so.1.0.0 -buildmode=plugin .

build/_output/e2sm_ni.so.1.0.0: # @HELP build the e2sm_ni.so.1.0.1
	cd servicemodels/e2sm_ni && CGO_ENABLED=1 go build -o build/_output/e2sm_ni.so.1.0.0 -buildmode=plugin .

PHONY:build
build: # @HELP build all libraries
build: build/_output/e2sm_kpm.so.1.0.0

test: # @HELP run the unit tests and source code validation
test: license_check build
# linters
	cd servicemodels/e2sm_kpm && GODEBUG=cgocheck=0 go test -race ./...

deps: # @HELP ensure that the required dependencies are in place
	cd servicemodels/e2sm_kpm
	go build -v -buildmode=plugin ./modelmain.go
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: # @HELP examines Go source code and reports coding problems
	cd servicemodels/e2sm_kpm && golangci-lint run --timeout 30m && cd ..

license_check: # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR} --boilerplate LicenseRef-ONF-Member-1.0

buflint: #@HELP run the "buf check lint" command on the proto files in 'api'
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2-sm \
		-w /go/src/github.com/onosproject/onos-e2-sm/servicemodels \
		bufbuild/buf:${BUF_VERSION} check lint

protos: # @HELP compile the protobuf files (using protoc-go Docker)
protos: buflint
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-e2-sm \
		-w /go/src/github.com/onosproject/onos-e2-sm \
		--entrypoint /go/src/github.com/onosproject/onos-e2-sm/build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

onos-e2-sm-base-docker: # @HELP build onos-e2-sm base Docker image
	@go mod vendor
	docker build . -f build/base/Dockerfile \
		--build-arg ONOS_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		--build-arg ONOS_MAKE_TARGET=build \
		-t onosproject/onos-e2-sm-base:${ONOS_E2_SM_VERSION}
	@rm -rf vendor

onos-e2-sm-docker: # @HELP build onos-e2-sm Docker image
onos-e2-sm-docker: onos-e2-sm-base-docker
	docker build . -f build/onos-e2-sm/Dockerfile \
		--build-arg ONOS_E2_SM_BASE_VERSION=${ONOS_E2_SM_VERSION} \
		-t onosproject/onos-e2-sm:${ONOS_E2_SM_VERSION}

images: # @HELP build all Docker images
images: build onos-e2-sm-docker

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/onos-e2-sm:${ONOS_E2_SM_VERSION}

all: build images

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION} onosproject/onos-e2-sm

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
