#!/usr/bin/env bash

BASEDIR=$(dirname "$0") # where genprotos resides


mkdir -p "$BASEDIR/gen"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc --experimental_allow_proto3_optional \
    --proto_path=protocol \
    --hrpc_out=$BASEDIR/gen \
    --hrpc_opt=hrpc-server-echo-go:hrpc-client-go:hrpc-scanner \
    --go_out=$BASEDIR/gen \
    $(find "${dir}" -name '*.proto')
done

rsync -a -v --remove-source-files $BASEDIR/gen/github.com/harmony-development/legato/gen/ $BASEDIR/gen
rm -rf $BASEDIR/gen/google.golang.org

goimports -w $BASEDIR/gen/*/v1/*.go
