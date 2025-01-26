#!/bin/bash

# Fetch and checkout the branch explicitly
git fetch origin $GITHUB_HEAD_REF
git checkout $GITHUB_HEAD_REF

echo "Current Directory: $(pwd)"
ls -la
git status

cd proto

for SERVICE_NAME in "order" "inventory" "payment"
do
protoc --go_out=./golang --go_opt=paths=source_relative \
  --go-grpc_out=./golang --go-grpc_opt=paths=source_relative \
 ./${SERVICE_NAME}/*.proto
done

git config --global --add safe.directory "$GITHUB_WORKSPACE"
git config --local user.name "github-actions[bot]"
git config --local user.email "github-actions[bot]@users.noreply.github.com"
git add . && git commit -am "proto update" || true
git push origin $GITHUB_HEAD_REF
