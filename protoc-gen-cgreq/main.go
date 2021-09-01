package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)

	opts := protogen.Options{}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	testdata := filepath.Dir(ex) + "/../protoc-gen-openapiv2/internal/genopenapi/testdata/"

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		// cgreq stands for CodeGeneratorRequest
		f := testdata + path.Base(file.GeneratedFilenamePrefix+".cgreq")
		err = ioutil.WriteFile(f, input, 0664)
		if err != nil {
			panic(fmt.Sprintf("error on writing data to %q file: %v", f, err))
		}
	}
}
