#!/bin/bash

PROTO_DIR=./api
GATEWAY_GEN_DIR=./gateway/internal/gen
STORAGE_GEN_DIR=./storage/internal/gen

# gateway proto client
protoc \
  -I $PROTO_DIR \
  --go_out=$GATEWAY_GEN_DIR \
  --go-grpc_out=$GATEWAY_GEN_DIR \
  $PROTO_DIR/*.proto


# storage proto server
protoc \
  -I $PROTO_DIR \
  --go_out=$STORAGE_GEN_DIR \
  --go-grpc_out=$STORAGE_GEN_DIR \
  $PROTO_DIR/*.proto