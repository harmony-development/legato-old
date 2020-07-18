#!/usr/bin/env bash
./protoc_gen_go.sh --proto_path=protocol --plugin_name=go --plugin_out=gen --plugin_opt=plugins=grpc