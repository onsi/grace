#!/usr/bin/env bash

source $HOME/.bashisms/s3_upload.bash

set -e

echo "Compiling for linux..."
GOOS=linux GOARCH=amd64 go build .
tar -zcf grace.tar.gz grace
echo "Uploading..."
upload_to_s3 grace.tar.gz
echo "Cleaning up..."
rm grace
rm grace.tar.gz

echo "Compiling for busybox..."
GOOS=linux GOARCH=amd64 go build -tags "busybox" .
echo "Constructing Dockerimage"
docker build -t="onsi/grace-busybox" .
docker push onsi/grace-busybox
echo "Cleaning up..."
rm grace
