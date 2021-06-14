# Protoc-Gen-Choice
This plugin takes as an input Protobuf file and generates a OneOf options map to treat CHOICE encoding.
> Plugin was written with `Protoc-Gen-Star (PG*)` tool. It generates CHOICE map for all CHOICEs defined in O-RAN **Service Models**.


## Build
To build this plugin, go inside folder `protoc-gen-choice` and run following command:

`go build -v -o ./protoc-gen-choice`

Or stay in the top-level directory of the repo and run `make build_protoc_gen_choice`.

It would build a `protoc-gen-choice` plugin.

## Usage
To use this plugin please run Protobuf compiler on `*.proto` files you want to process in the following way:

```bash
protoc -I="../servicemodels/" --plugin="./protoc-gen-choice" --choice_out="../servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go/" ../servicemodels/e2sm_kpm_v2_go/v2/e2sm_kpm_v2_go.proto
```

Here are the parameters you should pass:
- `-I="..."` specifies path to the imports included in `.proto` file(s)
- `--plugin="..."` specifies path to your custom plugin if it is not located in one of the folders in `$PATH`
- `--choice_out="..."` specifies path where to store generated files
- `*.proto` is a path to the source `.proto` file(s) to process
