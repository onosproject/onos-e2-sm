<!--
SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->

# Protoc-Gen-Builder
This plugin takes as an input Protobuf file and generates mandatory packages for SM:
* `encoder` - placeholder for top-level PDUs encoding and decoding functions.
* `pdubuilder` - placeholder for top-level PDU builders (and all supplementary functions)
  * A `builder` file for each `.proto` file would be generated as well - it holds builders for optional setting optional fields in the message.
* `servicemodel` - placeholder for main servicemodel interface.
  * This is optional package to generate
  * If you need to generate it, you should explicitly specify it in command line, see below

**A mandatory prerequisite is to have set $ONOS_ROOT variable in your environment!** 

## Build
To build this plugin, stay in the top-level directory, `onos-e2-sm`, and run:

```bash
make build_protoc_gen_builder
```

It would build & install a `protoc-gen-builder` plugin.

You can also go inside the folder `protoc-gen-builder` and run following commands to build plugin manually:

`go build -v -o ./protoc-gen-builder`

`go install`

Installing this plugin in your environment would allow you to run it from anywhere in your OS.

## Usage
To use this plugin please run Protobuf compiler on `*.proto` files you want to process in the following way:

```bash
protoc -I="$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api" --proto_path="sm=true:servicemodels/" --builder_out="servicemodels/e2sm_kpm_v2_go/" servicemodels/e2sm_kpm_v2_go/v2/e2sm_kpm_v2_go.proto
```

Here are the parameters you should pass:
- `-I="..."` specifies path to the imports included in `.proto` file(s)
- `--plugin="..."` specifies path to your custom plugin if it is not located in one of the folders in `$PATH`. Ignore it, if you've already run `make build_protoc_gen_choice`.
- `--builder_out="..."` specifies path where to store generated files
- `*.proto` is a path to the source `.proto` file(s) to process

Here is a set of parameters you could pass to the plugin:
- `sm` - generates `servicemodel` package.

Inside `$proto_imports` you should store path to the imported files in the proto - usually absolute path to the repo - 
otherwise it would through a warning that your `--proto_path` is not relative. 
Example of `$proto_imports`
```bash
proto_imports=${GOPATH}/src/github.com/onosproject/onos-e2-sm/
```

For E2AP, you should run following command under `onos-e2t/` folder:
```bash
protoc -I="$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2t/api" --proto_path="api/" --builder_out="sm=false:pkg/southbound/e2ap/" e2ap/v2/e2ap_pdu_descriptions.proto e2ap/v2/e2ap_pdu_contents.proto e2ap/v2/e2ap_ies.proto e2ap/v2/e2ap_containers.proto e2ap/v2/e2ap_constants.proto e2ap/v2/e2ap_commondatatypes.proto
```
## Known issues
* This plugin doesn't generate `servicemodel` package (this is to be done in the future)
* This plugin doesn't generate a `pdubuilder` package (this is to be done in the future)
  * It also doesn't generate builders for handling optional items of the message (this is to be done in the future)
