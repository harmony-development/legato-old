#!/usr/bin/env bash
mkdir -p "gen"
for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
  echo $(find "${dir}" -name '*.proto')
  protoc --experimental_allow_proto3_optional \
  --proto_path=protocol \
  --proto_path=${GOPATH}/src/github.com/google/protobuf/src \
  --proto_path=${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
  --go_out=gen \
  --go_opt="plugins=grpc" \
  --validate_out="lang=go:gen" \
  $(find "${dir}" -name '*.proto')
done