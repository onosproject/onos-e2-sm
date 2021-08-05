module github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/onosproject/onos-lib-go v0.7.14
	github.com/prometheus/common v0.4.0
	google.golang.org/protobuf v1.26.0
	gotest.tools v2.2.0+incompatible
)

replace github.com/onosproject/onos-lib-go => ../../../onos-lib-go
