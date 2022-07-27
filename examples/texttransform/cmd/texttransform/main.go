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
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	pb "github.com/voytechnology/protoc-gen-rest/examples/texttransform/v1"
	rpb "github.com/voytechnology/protoc-gen-rest/examples/texttransform/v1/texttransformrest"
)

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	http.Handle("/", rpb.NewTextTransformServiceHandler(&TextTransformServer{}))
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

// TextTransformServer implements the TextTransformService
type TextTransformServer struct{}

// Capitalize the input text
func (s *TextTransformServer) Capitalize(
	ctx context.Context, req *pb.CapitalizeRequest) (*pb.CapitalizeResponse, error) {

	return &pb.CapitalizeResponse{
		Text: strings.ToUpper(req.Text),
	}, nil
}
