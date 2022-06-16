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

package main

import (
	"flag"
	"os"

	"github.com/voytechnology/protoc-gen-rest/internal/descriptor"
	"github.com/voytechnology/protoc-gen-rest/internal/registry"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/voytechnology/protoc-gen-rest/internal/generator"
)

func main() {
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		registry := registry.NewRegistry(gen)
		generator := generator.New()

		var sources []*descriptor.SourceFile
		for _, fileName := range gen.Request.FileToGenerate {
			source, err := registry.Lookup(fileName)
			if err != nil {
				return err
			}
			sources = append(sources, source)
		}

		targets, err := generator.Generate(sources)
		if err != nil {
			return err
		}
		for _, target := range targets {
			f, err := os.Create(*target.Name)
			if err != nil {
				return err
			}
			if _, err := f.WriteString(*target.Content); err != nil {
				return err
			}
		}
		return nil
	})
}
