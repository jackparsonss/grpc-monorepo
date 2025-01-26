#!/bin/bash

sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

for SERVICE_NAME in "order" "inventory" "payment"
do
protoc --go_out=./golang --go_opt=paths=source_relative \
  --go-grpc_out=./golang --go-grpc_opt=paths=source_relative \
 ./${SERVICE_NAME}/*.proto
cd proto/golang/${SERVICE_NAME}
go mod init \
  github.com/jackparsonss/go-grpc/golang/${SERVICE_NAME} || true
go mod tidy
done

cd ../../../
git config --local user.name "github-actions[bot]"
git config --local user.email "github-actions[bot]@users.noreply.github.com"
git add . && git commit -am "proto update" || true
git push origin
