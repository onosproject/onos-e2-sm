module github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/gogo/protobuf v1.3.2
	github.com/onosproject/onos-api/go v0.9.8
	github.com/onosproject/onos-lib-go v0.8.15
	google.golang.org/protobuf v1.28.0
	gotest.tools v2.2.0+incompatible
)

replace github.com/onosproject/onos-api/go => github.com/adibrastegarnia/onos-api/go v0.9.9-0.20220426000032-3afe72d227b6
