# Protoc-Gen-CGo
This plugin takes as an input Protobuf file and generates a CGo code appropriately.
> Plugin was written with `Protoc-Gen-Star (PG*)` tool and it handles CGo code generation for O-RAN defined **Service Models**.

Please note that this plugin requires additional changes in Protobuf file (more in **Additional changes** section).

## Build
To build this plugin, go inside folder `protoc-gen-cgo` and run following command:

`go build -v -o ./protoc-gen-cgo`

It would build a `protoc-gen-cgo` plugin.

## Usage
To use this plugin please run Protobuf compiler on `*.proto` files you want to process in the following way:
```bash
protoc -I="." --plugin="./protoc-gen-cgo" --cgo_out="./generated" *.proto
```

Here are the parameters you should pass:

- `-I="..."` specifies path to the imports included in `.proto` file(s),
- `--plugin="..."` specifies path to your custom plugin if it is not located in one of the folders in `$PATH`,
- `--cgo_out="..."` specifies path where to store newly generated CGo code. Also specifies input parameters to the plugin (see example above),
- `*.proto` is a path to the source `.proto` file(s) to process.

Here is a set of parameters you could pass to the plugin:
- `cgo` - generates CGo glue code (additional changes by hand are needed),
- `ut` - generates unit test to corresponding CGo code (additional changes by hand are needed),
- `t` - generates types (custom structures, wrapper to ease Protobuf usage),

A valid example of plugin usage:
```bash
protoc -I="$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api" --plugin="./protoc-gen-cgo" --cgo_out="cgo=true,ut=false,t=false:./generated" ../servicemodels/e2sm_kpm_v2/v2/e2sm_kpm_v2.proto
```

## Additonal changes
To ensure that this plugin works correctly you should make some changes to source `.proto` file(s). 
#### Message
You need to add in the header of the message a name of the corresponding `.h` (which was generated with `ASN1C` tool) file enclosed with `{}` parenthesis. 
If there is no appropriate C-file for this message, or this message is a constant, put either `{-}` in the header or the word `constant` (you can also put both) - it would 
tell the tool to avoid generating code. If you put `{-}`, don''t forget to delete'      
Also add `[json_name="..."]` tag to each leaf of the message 
and put in as an argument appropriate `C-name` taken from corresponding `.h` file. Do it for all data structures in each `.proto` file
you want to process. If one of the message substructures is optional, consider adding `:OPTIONAL` to the end of the `json_name` tag - it would push tool 
to handle them as an optional instances.    
In the end it should look similar to the following example:

```protobuf
//{SNSSAI}
message Snssai {
  bytes s_st = 1 [(validate.v1.rules).bytes.len = 1, json_name = "sST"];
  bytes s_d = 2 [(validate.v1.rules).bytes.len = 3, json_name = "sD:OPTIONAL"];
};
```

#### Enum
In the header of an `Enum` type should be added a name of the corresponding `.h` file enclosed with `{}` parenthesis. See example below:
```protobuf
// {RT-Period-IE}
enum RtPeriodIe {
  ...
}
```

## Additional changes by hand
Generated code covers only some part of the work (thankfully big one). Somehow some changes by hand should be done. 
Majority of them has been covered with `//ToDo` comments inserted into generated code. In the first step please address all of them.

#### CGo code
The most tricky part is to verify that all basic structures are encoded correctly (like SNSSAI, TimeStamp and others). Please go through each
generated message and verify that the encoding and decoding is handled properly and due to the one defined in corresponding `.h` file.    
Pay attention to the messages which have `OneOf` substructure and the other substructures not included in the OneOf structure. 
This case is not handled properly and couldn't be distinguished by the tool (yet).

#### Unit tests
You will need to fill in and adjust `createXXX` functions in each unit test. Somehow, it provides basic structure for composing the message.


> Once all changes are done, please use `go build -gcflags="-e" ./...` command to build the code. It would list you all errors compiler could raise during the building process.

You can also notice that within generated files are present files ending with `_types.go`. These files were generated as a wrappers translating abstract structures to the `Protobuf` to ease the usage of it. 
In the future releases this feature would be addressed with the flag necessary to pass to the plugin (see `SDRAN-621`). 

## Troubleshooting

#### Post-processing errors caused by `GoFmt()`
In some cases, plugin could raise an error during post-processing of generated files. Unfortunately it doesn't provide much information to debug it. 
In this case please comment `line 21` in `main.go`, recompile the plugin and run it again to generate files.
Once it's done, run following command on generated files to get all errors GoFmt() raises:
`gofmt -w generated/*.go`

Possible causes:
1. `Protobuf` file misses some fields you should provide in `[json_name]` tag (see section **Additional changes**)
2. Template file generates some peaces of code which could be hard to distinguish (two lines were not split, etc.)

#### Other errors occured during code generation
If you have found an error which was not covered here, a good place to track the root cause could be looking into the log output generated by the plugin. 
In this case please refer to the file `/tmp/report.txt`.