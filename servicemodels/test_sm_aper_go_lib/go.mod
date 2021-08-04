module github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib

go 1.16

require (
	github.com/onosproject/onos-lib-go v0.7.14
	google.golang.org/genproto v0.0.0-20201113130914-ce600e9a6f9e // indirect
	google.golang.org/protobuf v1.26.0
	gotest.tools v2.2.0+incompatible
)

//replace github.com/onosproject/onos-lib-go => ../../../onos-lib-go
