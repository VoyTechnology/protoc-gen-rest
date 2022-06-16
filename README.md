# protoc-gen-rest

Create easy to consume REST APIs from protobufs

## Features

This project is still in its infancy so nothing is yet actually developed.

- [ ] Server Generated Code
- [ ] Client Generated Code
- [ ] Unary RPCs
- [ ] Server Streaming RPCs
- [ ] Client Streaming RPCs

## Usage

The `protoc-gen-rest` is a plugin of the Google protobuf compiler. It reads
the protobuf service definion and generates a REST API using the annotations.

It is designed to solve the simple problem of creating REST APIs from protobufs.

### 1. Define your gRPC service using protocol buffers

`package/service/v1/service.proto`

```protobuf
syntax = "proto3";

package your.service.v1;

option go_package = "github.com/org/repo/proto/your/service/v1;service";

import "google/api/annotations.proto";

service CapitalizeService {
    rpc Capitalize(CapitalizeRequest) returns (CapitalizeResponse) {
        // These are the annotations that are used to generate the REST API.
        option (google.api.http) = {
            post: "/capitalize"
            body: "*"
        };
    }
}

message CapitalizeRequest {
    string message = 1;
}

message CapitalizeResponse {
    string message = 1;
}
```

### 2. Generate your REST API

`buf.yaml`

```yaml
version: v1
deps:
- buf.build/googleapis/googleapis
```

`buf.gen.yaml`

```yaml
version: v1
plugins:
    - name: go
      out: proto
    - name: protoc-gen-rest
      out: proto
```

### 3. Implement your service

```go
package main

import (
    pb "github.com/org/repo/proto/your/service/v1"
)

func run() error {
    h := pb.NewCapitalizeServiceHandler(&server{})
    http.Handle("/api", h)

    return http.ListenAndServe(":8080", nil)
}

type server struct {}

// Capitalize is implemented like any other gRPC service.
func (s *server) Capitalize(
    ctx context.Context, req *pb.CapitalizeRequest,
) (*pb.CapitalizeResponse, error) {
    return &pb.CapitalizeResponse{
        Message: strings.ToUpper(req.Message),
    }, nil
}

func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }
}
```

## More Examples

Coming soon

## Developing

There is no formalized process just yet, so for now just run

```sh
make local
```

And look at the generated files that the binary produces.
