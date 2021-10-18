# onos-e2-sm
[O-RAN] defines the **E2 Service Models** in the form of "ASN.1" specifications. 

Each Service Model defines 5 top level objects:
1. **RAN Function Description** (from the E2Node, describes the supported actions and triggers)
1. **Event Trigger Defintion** (contained in a `RICSubscriptionRequest`, defines the conditions on which E2Node should send a `RICIndication`)
1. **Action Definition** (contained in a `RICSubscriptionRequest`, defines the actions on an E2Node)
1. **Indication Header** (contained in a `RICIndication`, describes general parameters of the source E2 Node)
1. **Indication Message** (contained in a `RICIndication`, describes specific parameters requested)

## Implementation in SD-RAN
The `onos-e2-sm` project provides, for each of these Service Models:
* A Protobuf translation of these ASN.1 specifications e.g. [e2sm_kpm_ies.proto](servicemodels/e2sm_kpm/v1beta1/e2sm_kpm_ies.proto)
* Go code mapping between Protobuf and ASN.1 PER encoding

The implementation can be accessed as either:
* a Go module e.g. `go get github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm@v0.7.0`
  (preferred for **xApps** who only need to access the Proto defintions) or
* as a Go **plugin** e.g. `e2sm_kpm.so.1.0.0` (allowing a loose coupling with `onos-e2t` and `ran-simulator`)

> Since dynamically loaded modules in Go require being compiled in the code of the target they will plugin to, there
> are 2 versions of the plugin docker file produced `onosproject/service-model-docker-e2sm_kpm-1.0.0`
> and `onosproject/service-model-ransim-e2sm_kpm-1.0.0`

> Third party vendors will be able to build their own Service Models and load them in to `onos-e2t` using the plugin method,
> and will be able to access the translated Protobuf in corresponding xApps

### Key Performance Metrics (E2SM_KPM)
This is the first E2 Service Model to be handled by SD-RAN - it is for extracting statistics from the E2Node.

Currently supported version is E2SM KPM v2 (with CGo approach). E2SM KPMv2 with purely Go approach is 
under maintenance and yet far away from the release. E2SM KPM v1 is partially implemented and not supported anymore.

There is also an experimental implementation of KPMv2 SM with Go-based APER library (produces APER bytes out of Protobuf). 
This is still under verification, bugs may be expected.


### Native Interface (E2SM_NI)
While the Proto definitions have been created for this Service Model, the Go mapping code has not been implemented in SD-RAN yet.

### RAN Control (E2SM_RC_PRE)
Pre-standard E2 Service model with PCI and Neighbor relation table information from E2 Nodes.

There is also an experimental implementation of RC-PRE SM with Go-based APER library. This is still under verification, 
bugs may be expected.

### Mobile HandOver (E2SM_MHO)
E2 Service model for handling Mobile HandOver use case.

There is also an experimental implementation of MHO SM with Go-based APER library. This is still under verification, 
bugs may be expected.

### RAN Slicing (E2SM_RSM)
E2 service model for handling RAN Slicing use case. It was implemented with Go-based APER library.

## Development
Service models are created from the ASN1 models stored at:
https://github.com/onosproject/openairinterface5g/tree/develop-onf/openair2/RIC_AGENT/MESSAGES/ASN1/R01

From these:

* Protobuf is generated with `asn1c -B` (requires the specially modified version of `asn1c` tool - [asn1c](https://github.com/onosproject/asn1c))
* Go code is generated from this Protobuf as an interface/reusable layer e.g. [e2sm_kpm_ie.pb.go](servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies/e2sm_kpm_ies.pb.go)
* C code is generated the version of the `asn1c` tool from the [O-RAN Software Community](https://gerrit.o-ran-sc.org/r/admin/repos/com/asn1c)
  with `asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.` e.g. [E2SM-KPM-IndicationHeader.h](servicemodels/e2sm_kpm/kpmctypes/E2SM-KPM-IndicationHeader.h)
* Then glue code is generated by hand (at first) using `CGO` (wrapping the C code in Go) e.g. [E2SM-KPM-IndicationHeader.go](servicemodels/e2sm_kpm/kpmctypes/E2SM-KPM-IndicationHeader.go)
  * It's also possible to use `protoc-gen-cgo`, a `protoc` plugin that prints CGo code out of Protobuf. Some hand tweaks are still needed to be done.  

> To generate the C code with the O-RAN Software Community version of the `asn1c` tool,
> it must be installed on your system with `sudo make install`.
> This is because it takes skeleton file from `/usr/local/share/asn1c`
> regardless of where it is run from.

The E2AP (E2 Application Protocol) is not a Service Model, and so is kept completely inside the `onos-e2t`.

[O-RAN]: https://www.o-ran.org/
