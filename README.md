# sdran-e2-sm
SD-RAN E2 service model plugins

Service models are created from the ASN1 models stored at:
https://github.com/onosproject/openairinterface5g/tree/develop-onf/openair2/RIC_AGENT/MESSAGES/ASN1/R01

From these:

* Protobuf is generated with `asn1c -B`
* Go code is generated from this Protobuf as an interface/reusable layer
* C code is generated with `asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`

Then glue code is generated by hand (at first) using `CGO` (wrapping the C code in Go)

The E2AP (E2 Application Protocol) is not a Service Model, and so is kept completely inside the `onos-e2t`.

## Note
Current implementation of E2SM KPM model in `CGO` gives an error related to the `GO` pointers referencing bunch of `C` files. 
You can verify it running `go test -v ./..` in `servicemodels/e2sm_kpm/kpmctypes/`.

**Work is in progress.**