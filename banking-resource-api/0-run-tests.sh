#!/bin/bash
set -eo pipefail
REGION=$(aws configure get region)
echo "Using aws region: $REGION..."
AWS_REGION=$REGION go test -coverpkg=./... -coverprofile=c.out ./...


