#!/usr/bin/env bash

dt=$(date -u +%Y%m%d)
version=" BUILD @ $(date -u '+%Y%m%d %H:%M:%S') "
echo "$version"
sed -i "" -E "s/####.+####/#### $version ####/g" main.go
mkdir -p dist
env GOOS=linux GOARCH=amd64 go build -o "./dist/screenroom-ws.$dt" .

rsync -avx --progress "./dist/screenroom-ws.$dt" ir:/opt/bin/


