#!/bin/sh

proto_path=$1
protoc --go_out=./pb --go-grpc_out=./pb --go-grpc_opt=require_unimplemented_servers=false $proto_path
