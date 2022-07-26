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

// Reference implementation

package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/voytechnology/protoc-gen-rest/testdata/simple/v1"
	restpb "github.com/voytechnology/protoc-gen-rest/testdata/simple/v1/simplerest"
)

func main() {
	http.Handle("/", restpb.NewSimpleServiceHandler(&Server{}))
	http.ListenAndServe(":8080", nil)
}

type Server struct {
	restpb.SimpleServiceServer
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Letter: "get",
	}, nil
}

func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	log.Println("Set", req)
	return &pb.SetResponse{}, nil
}
