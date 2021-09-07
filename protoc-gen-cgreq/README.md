# protoc-gen-cgreg

**protoc-gen-cgreg** is yet another plugin for protocol buffer compiler whose sole purpose is to help out with testing other plugins (protoc-gen-openapiv2, protoc-gen-grpc-gateway). Currently it has hardcoded path that limits its use to **_protoc-gen-openapiv2_** only. In the future it might be updated with options allowing necessary customization and wider appliance.

## How it works

Normally, a protocol compiler ([protoc](https://grpc.io/docs/protoc-installation/), [bufbuild/buf](https://github.com/bufbuild/buf)) is extandable via plugins where the code stubs are generated and output. A plugin should implement the workload which starts from getting binary representation of proto file which is then used to generate and output required stub (file). Both aforementioned protocol compilers pass down proto input as [CodeGeneratorRequest](https://pkg.go.dev/google.golang.org/protobuf).

https://developers.google.com/protocol-buffers/docs/reference/other
> A plugin is just a program which reads a CodeGeneratorRequest protocol buffer from standard input and then writes a CodeGeneratorResponse protocol buffer to standard output.

Basically, protoc serializes the message descriptions it parsed into a **CodeGenerationRequest** message and writes it to a plugin via stdin. The plugin parses the request, generates the source code that the user requested and writes it back as a **CodeGenerationResponse** serialized to binary representation via stdout.

Because **_protoc-gen-openapiv2_** takes an instance of CodeGeneratorRequest for an input, it is difficult to initialize one directly. Much simpler way is to parse certain proto and get it as the result of such parsing. This is what the plugin does. It simply writes the whole binary content of received CodeGeneratorRequest into a file with `cgreq` extension which is then can be feed into **_protoc-gen-openapiv2_** plugin in order to get and verify the output.

## How to use

**protoc-gen-cgreq** directory has the plugin source code in main.go. It should be built first and then used. There is a number of Make targets in the root of the project that do this:

- atlas-build-cgreq
- atlas-regen-tdata

**atlas-build-cgreq** target builds the executable out of the source code and puts it into the same directory. 
**atlas-regen-tdata** target rebuilds the the executable every time and generates corresponding *.cgreq files out of proto files found in

```
$(PROJ_DIR)/protoc-gen-openapiv2/internal/genopenapi/testdata
```

directory. So, for `atlaspatch.proto` it should produce alongside `atlaspatch.cgreq` which is then used by unit tests as an input to get whatever **_protoc-gen-openapiv2_** plugin generates out of it.

Note that **protoc-gen-cgreg** plugin was implemented as auxiliary tool to test **_protoc-gen-openapiv2_** plugin and might need adjusting to test other plugins.  
