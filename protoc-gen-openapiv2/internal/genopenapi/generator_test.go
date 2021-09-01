package genopenapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/codegenerator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	genpkg "github.com/grpc-ecosystem/grpc-gateway/v2/internal/generator"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	reset   = "\033[0m"
)

// emitFiles is a modified function taken from main package
// where Generate method is called from. The modification is about
// returning an error that might occur when calling emitResp function.
func emitFiles(out []*descriptor.ResponseFile) error {
	files := make([]*pluginpb.CodeGeneratorResponse_File, len(out))
	for idx, item := range out {
		files[idx] = item.CodeGeneratorResponse_File
	}
	resp := &pluginpb.CodeGeneratorResponse{File: files}
	codegenerator.SetSupportedFeaturesOnCodeGeneratorResponse(resp)
	return emitResp(resp)
}

// emitResp is a modified version of similar function from
// main package where it is used for writing responses to Stdout.
// Instead it is modified to write responses to files further
// inspected for equality against wanted result.
func emitResp(resp *pluginpb.CodeGeneratorResponse) error {
	for _, f := range resp.File {
		bs := filepath.Base(f.GetName())
		tn := bs[:strings.IndexRune(bs, '.')]
		fn := "./testdata/" + tn + ".emitted.swagger.json"

		err := ioutil.WriteFile(fn, []byte(f.GetContent()), 0664)
		if err != nil {
			return fmt.Errorf("error on writing CodeGeneratorResponse content to %q file: %v", fn, err)
		}
	}
	return nil
}

func deepCompare(file1, file2 string) (bool, error) {
	const chunkSize = 64000

	f1, err := os.Open(file1)
	if err != nil {
		return false, fmt.Errorf("cannot open %q file, error: %v", file1, err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return false, fmt.Errorf("cannot open %q file, error: %v", file2, err)
	}
	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil
			} else if err1 == io.EOF || err2 == io.EOF {
				return false, nil
			} else {
				return false, fmt.Errorf("unexpected (not EOF) error "+
					"when reading files (1, 2): %v, %v", err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			return false, nil
		}
	}
}

func Test_generator_Generate(t *testing.T) {
	const (
		// testAtlaspatch tests all modification made by atlas patch
		testAtlaspatch = "atlaspatch"
	)

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd failed with error: %v", err)
	}

	testdata := pwd + "/testdata"
	t.Logf("Testdata folder: %s", testdata)

	type testwith struct {
		registry  *descriptor.Registry
		generator genpkg.Generator
		request   *pluginpb.CodeGeneratorRequest
		targets   []*descriptor.File
	}

	setup := func(t *testing.T, tname string) (testwith, []func()) {
		// remove previously emitted swagger JSON file if it exists
		swag := testdata + "/" + tname + ".emitted.swagger.json"
		if err := os.Remove(swag); err != nil && !os.IsNotExist(err) {
			t.Fatalf("Failed to remove %q file: %v", swag, err)
		}

		// setup
		var (
			with  testwith
			close []func()
		)

		proto := testdata + "/" + tname + ".cgreq"
		f, err := os.Open(proto)
		if err != nil {
			t.Fatalf("%s file open error: %v", proto, err)
			return with, close
		}
		close = append(close, func() {
			if err := f.Close(); err != nil {
				t.Errorf("Closing %s file failed: %v", proto, err)
			}
		})

		with.registry = descriptor.NewRegistry()

		// tune registry up to a certain test
		switch tname {
		case testAtlaspatch:
			with.registry.SetUseJSONNamesForFields(true)
			with.registry.SetMergeFileName("apidocs")
			// Atlas specific flags
			with.registry.SetAtlasPatch(true)
			with.registry.SetPrivateOperations(true)
			with.registry.SetCustomAnnotations(true)
		default:
			t.Fatalf("Unimplemented test %q", tname)
			return with, nil
		}

		with.request, err = codegenerator.ParseRequest(f)
		if err != nil {
			t.Fatalf("Parsing %s failed: %v", proto, err)
			return with, close
		}

		if err := with.registry.Load(with.request); err != nil {
			t.Fatalf("Loading data from request failed: %v", err)
			return with, close
		}

		//t.Logf("Registrty: %s", with.registry.String())

		with.generator = New(with.registry)

		for _, target := range with.request.FileToGenerate {
			f, err := with.registry.LookupFile(target)
			if err != nil {
				t.Fatalf("File lookup failure: %v", err)
				return with, close
			}
			with.targets = append(with.targets, f)
		}

		return with, close
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    testAtlaspatch,
			wantErr: false,
		},

		// TODO
		// test private methods
		// comments extraction

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			with, close := setup(t, tt.name)
			for _, fn := range close {
				defer fn()
			}

			gen, err := with.generator.Generate(with.targets)

			switch tt.wantErr {
			case true:
				if err == nil {
					genJSON, err := json.MarshalIndent(gen, "", "    ")
					if err != nil {
						t.Errorf("JSON marshal error %v", err)
						return
					}

					t.Errorf("\t%s No error but one is expected"+
						"\nGenerated: "+red+" \n\n%s\n\n "+reset, failed, string(genJSON))
					return
				}

				t.Logf("\t%s %s test is passed", succeed, tt.name)

			case false:
				if err != nil {
					t.Errorf("\t%s Unexpected error: "+red+" \n\n%s\n\n "+reset, failed, err.Error())
					return
				}

				genJSON, err := json.MarshalIndent(gen, "", "    ")
				if err != nil {
					t.Errorf("JSON marshal error %v", err)
					return
				}

				//t.Logf("\nGenerated: "+green+" \n\n%s\n\n "+reset, string(genJSON))

				if err := emitFiles(gen); err != nil {
					t.Errorf("\t%s Error attempting to emit generated result: "+red+" \n\n%s\n\n "+reset, failed, err.Error())
					t.Errorf("\nGenerated: "+red+" \n\n%s\n\n "+reset, string(genJSON))
					return
				}

				file1 := "./testdata/" + tt.name + ".emitted.swagger.json"
				file2 := "./testdata/" + tt.name + ".wanted.swagger.json"

				eq, err := deepCompare(file1, file2)
				if err != nil {
					t.Errorf("Emitted vs wanted files content comparison error: %v", err)
					return
				}

				if !eq {
					t.Errorf("\t%s Emitted swagger JSON file is not equal to wanted one, compare with: "+
						"\n"+yellow+" \n\ndiff %s %s\n\n "+reset, failed, file1, file2)
					return
				}
				t.Logf("\t%s %s test is passed", succeed, tt.name)
			}
		})
	}
}
