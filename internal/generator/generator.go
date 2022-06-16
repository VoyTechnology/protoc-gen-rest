// Copyright 2022 VoyTechnology
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"bytes"
	"fmt"
	"go/format"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/voytechnology/protoc-gen-rest/internal/descriptor"
)

type generator struct {
}

func New() *generator {
	return &generator{}
}

func (g *generator) Generate(targets []*descriptor.SourceFile) ([]*descriptor.TargetFile, error) {
	var files []*descriptor.TargetFile
	for _, file := range targets {
		code, err := g.generateFile(file)
		if err != nil {
			return nil, err
		}

		formatted, err := format.Source([]byte(code))
		if err != nil {
			return nil, fmt.Errorf("error formatting generated code: %v", err)
		}

		files = append(files, &descriptor.TargetFile{
			CodeGeneratorResponse_File: &pluginpb.CodeGeneratorResponse_File{
				Name:    proto.String(file.GeneratedFilenamePrefix + ".pb.rest.go"),
				Content: proto.String(string(formatted)),
			},
		})
	}

	return files, nil
}

func (g *generator) generateFile(file *descriptor.SourceFile) (string, error) {
	b := new(bytes.Buffer)
	if err := tmpl.Execute(b, file); err != nil {
		return "", err
	}

	return b.String(), nil

}
