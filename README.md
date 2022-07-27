# protoc-gen-rest

Create easy to consume REST APIs from protobufs

## Features

This project is still in its infancy so nothing is yet actually developed.

- [x] Server Generated Code [[#1](https://github.com/VoyTechnology/protoc-gen-rest/issues/1)]
- [ ] Client Generated Code [[#2](https://github.com/VoyTechnology/protoc-gen-rest/issues/2)]
- [ ] Streaming RPCs
- [ ] Accept multiple Content-Types
  - [x] `application/json` [[#3](https://github.com/VoyTechnology/protoc-gen-rest/issues/3)]
  - [ ] `application/grpc-web+proto` [[#6](https://github.com/VoyTechnology/protoc-gen-rest/issues/6)]]
- [ ] Publish to buf.build registry

## Usage

The `protoc-gen-rest` is a plugin of the Google protobuf compiler. It reads
the protobuf service definion and generates a REST API using the annotations.

It is designed to solve the simple problem of creating REST APIs from protobufs.

### 1. Define your gRPC service using protocol buffers

`package/service/v1/service.proto`

```protobuf
syntax = "proto3";

package examples.texttransform.v1;

option go_package = "github.com/voytechnology/protoc-gen-rest/examples/texttransform/v1;texttransform";

import "google/api/annotations.proto";

// TextTransformService is an example service used for transforming text.
service TextTransformService {
    rpc Capitalize(CapitalizeRequest) returns (CapitalizeResponse) {
        option (google.api.http) = {
            post: "/v1/texttransform/capitalize"
        };
    }
}

message CapitalizeRequest {
    string text = 1;
}

message CapitalizeResponse {
    string text = 1;
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
  - remote: buf.build/library/plugins/go:v1.27.1-1
    out: .
    opt:
      - paths=source_relative
  - name: rest
    out: .
    opt:
      - paths=source_relative
```

> Note: You must install the generator prior to generating the protobuf using
> `go install github.com/voytechnology/protoc-gen-rest/cmd/protoc-gen-rest`.
>
> Once the generator is published to buf.build this will no longer be
> required and you would be able to use a remote instead.

```sh
buf generate
```

### 3. Implement your service

```go
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
```

### 4. Use your newly generated API

Run your service and send a request to the endpoint picked with the
`google.api.http` option.

```sh
curl -X POST -d '{"text":"abc"}' localhost:8080/v1/texttransform/capitalize
{"text": "ABC"}
```

## More Examples

Coming soon

## Developing

There is no formalized process just yet, so for now just run

```sh
make local
```

And look at the generated files that the binary produces.
