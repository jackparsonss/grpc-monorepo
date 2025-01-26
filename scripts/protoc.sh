#!/bin/bash

# Fetch and checkout the branch explicitly
git fetch origin $GITHUB_HEAD_REF
git checkout $GITHUB_HEAD_REF

sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

cd proto

for SERVICE_NAME in "order" "inventory" "payment"
do
protoc --go_out=./proto/golang --go_opt=paths=source_relative \
  --go-grpc_out=./proto/golang --go-grpc_opt=paths=source_relative \
 ./proto/${SERVICE_NAME}/*.proto
done

git config --local user.name "github-actions[bot]"
git config --local user.email "github-actions[bot]@users.noreply.github.com"
git add . && git commit -am "proto update" || true
git push origin $GITHUB_HEAD_REF
