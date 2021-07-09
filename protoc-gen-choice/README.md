# Protoc-Gen-Choice
This plugin takes as an input Protobuf file and generates a OneOf options map to treat CHOICE encoding.
> Plugin was written with `Protoc-Gen-Star (PG*)` tool. It generates CHOICE map for all CHOICEs defined in O-RAN **Service Models**.


## Build
To build this plugin, go inside folder `protoc-gen-choice` and run following commands:

`go build -v -o ./protoc-gen-choice`

`go install`

Installing this plugin in your environment would allow you to run it from anywhere in your OS. 

You can also stay in the top-level directory of the repo and run `make build_protoc_gen_choice`, which would build & install a `protoc-gen-choice` plugin.

## Usage
To use this plugin please run Protobuf compiler on `*.proto` files you want to process in the following way:

```bash
protoc -I="servicemodels/" --choice_out="servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2/" servicemodels/e2sm_kpm_v2/v2/e2sm_kpm_v2.proto
```

Here are the parameters you should pass:
- `-I="..."` specifies path to the imports included in `.proto` file(s)
- `--plugin="..."` specifies path to your custom plugin if it is not located in one of the folders in `$PATH`. Ignore it, if you've already run `make build_protoc_gen_choice`.
- `--choice_out="..."` specifies path where to store generated files
- `*.proto` is a path to the source `.proto` file(s) to process
