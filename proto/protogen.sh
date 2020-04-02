#!/usr/bin/env bash

# Protocol Buffers for Go with Gadgets
#
# Copyright (c) 2013, The GoGo Authors. All rights reserved.
# http://github.com/gogo/protobuf

protoc \
	-I. \
	-I ${GOPATH}/src \
	--grpc-gateway_out=logtostderr=true:. \
	--go_out=plugins=grpc:. *.proto