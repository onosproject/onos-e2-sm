# Protoc-Gen-CGo
This plugin takes as an input Protobuf file and generates a CGo code appropriately.
> Plugin was written with `Protoc-Gen-Star (PG*)` tool.

Please note that this plugin requires additoinal changes in Protobuf file (more in **Additional changes** section).

## Build
To build this plugin, go inside folder `protoc-gen-cgo` and run following command:

`go build -v -o ./protoc-gen-cgo`

It would build a `protoc-gen-cgo` plugin.

## Usage
To use this plugin please run Protobuf compiler on `*.proto` files you want to process in the following way:

`protoc -I="." --plugin="./protoc-gen-cgo" --cgo_out="./generated" *.proto`

Here are the parameters you should pass:

- `-I="..."` specifies path to the imports in `.proto` file(s)

- `--plugin="..."` specifies path to your custom plugin if it is not located in one of the folders in `$PATH`

- `--cgo_out="..."` specifies path where to store generated files

- `*.proto` is a path to the source `.proto` file(s) to process

## Additonal changes
To ensure that this plugin works correctly you should make some changes to source `.proto` file(s). Please add `[json_name="..."]` tag 
and put in as an argument appropriate `C-struct` name (which was generated with `ASN1C` tool) for all data structures in each `.proto` file
you want to process.   
In the end it should look similar to following example:

```protobuf
message RicStyleType {
    int32 value = 1 [json_name = "RIC-Style-Type"];
};
```

<span style="color:red"> This is a temporary solution and would be reconsidered with the future updates. </span>