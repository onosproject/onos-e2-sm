module github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/onosproject/onos-api/go v0.7.80
	github.com/onosproject/onos-lib-go v0.7.13
	github.com/pkg/errors v0.9.1 // indirect
	google.golang.org/protobuf v1.26.0
	gotest.tools v2.2.0+incompatible
)

replace github.com/onosproject/onos-lib-go => ../../../onos-lib-go
