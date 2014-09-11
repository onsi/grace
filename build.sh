#!/usr/bin/env bash

source $HOME/.bashisms/s3_upload.bash

set -e

echo "Compiling for linux..."
GOOS=linux GOARCH=amd64 go build .
tar -zcf grace.tar.gz grace
rm grace
echo "Uploading..."
upload_to_s3 grace.tar.gz
rm grace.tar.gz
