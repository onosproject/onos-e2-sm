# Protoc-Gen-Choice
This plugin takes as an input Protobuf file and generates a OneOf options map to treat CHOICE encoding.
> Plugin was written with `Protoc-Gen-Star (PG*)` tool. It generates CHOICE map for all CHOICEs defined in O-RAN **Service Models**.

**A mandatory prerequisite is to have set $ONOS_ROOT variable in your environment!** 

## Build
To build this plugin, stay in the top-level directory and run 

```bash
make build_protoc_gen_choice
```

It would build & install a `protoc-gen-choice` plugin.

You can also go inside the folder `protoc-gen-choice` and run following commands to build plugin manually:

`go build -v -o ./protoc-gen-choice`

`go install`

Installing this plugin in your environment would allow you to run it from anywhere in your OS.

## Usage
To use this plugin please run Protobuf compiler on `*.proto` files you want to process in the following way:

```bash
protoc -I="$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api" --proto_path="servicemodels/" --choice_out="servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go/" servicemodels/e2sm_kpm_v2_go/v2/e2sm_kpm_v2_go.proto
```

Here are the parameters you should pass:
- `-I="..."` specifies path to the imports included in `.proto` file(s)
- `--plugin="..."` specifies path to your custom plugin if it is not located in one of the folders in `$PATH`. Ignore it, if you've already run `make build_protoc_gen_choice`.
- `--choice_out="..."` specifies path where to store generated files
- `*.proto` is a path to the source `.proto` file(s) to process

Inside `$proto_imports` you should store path to the imported files in the proto - usually absolute path to the repo - 
otherwise it would through a warning that your `--proto_path` is not relative. 
Example of `$proto_imports`
```bash
proto_imports=${GOPATH}/src/github.com/onosproject/onos-e2-sm/
```
