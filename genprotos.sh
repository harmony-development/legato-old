#!/usr/bin/env bash
mkdir -p "gen"

cd ./hrpc
go build
cd ..

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc --experimental_allow_proto3_optional \
    --proto_path=protocol \
    --hrpc_out=./gen \
    --hrpc_opt=hrpc-server-echo-go:hrpc-scanner:hrpc-client-go \
    --go_out=./gen \
    --validate_out="lang=go:gen" \
    $(find "${dir}" -name '*.proto')
done

rsync -a -v --remove-source-files gen/github.com/harmony-development/legato/gen/ ./gen

goimports -w ./gen/*/v1/*.go
