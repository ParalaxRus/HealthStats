#!/bin/bash

PROTO_DIR=./proto
WEB_APP_GEN_DIR=./app/src/gen
STORAGE_GEN_DIR=./storage/internal/gen

protoc \
  -I $PROTO_DIR \
  --go_out=$STORAGE_GEN_DIR \
  --go-grpc_out=$STORAGE_GEN_DIR \
  service.proto

protoc -I $PROTO_DIR \
  --plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
  --ts_proto_out=$WEB_APP_GEN_DIR \
  --ts_proto_opt=outputClientImpl=grpc-web,esModuleInterop=true \
  service.proto