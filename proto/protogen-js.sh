#!/usr/bin/env bash

protoc \
	-I. \
  -I ${GOPATH}/src \
  --grpc-web_out=import_style=commonjs,mode=grpcweb:. *.proto