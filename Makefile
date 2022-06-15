install:
	go install ./cmd/protoc-gen-rest

generate:
	buf generate

.PHONY: all

local: install generate
