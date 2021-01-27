#!/usr/bin/env bash
mkdir -p "gen"

go build -o ./hrpc/hrpc-bin ./hrpc/main.go

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc --experimental_allow_proto3_optional \
    --proto_path=protocol \
    --plugin=protoc-gen-custom=./hrpc/hrpc-bin \
    --custom_out=./gen \
    --custom_opt="/templates/hrpc-server-go.htmpl" \
    --go_out=./gen \
    --validate_out="lang=go:gen" \
    $(find "${dir}" -name '*.proto')
done

rsync -a -v --remove-source-files gen/github.com/harmony-development/legato/gen/ ./gen

go fmt ./gen/./...